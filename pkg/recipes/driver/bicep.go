/*
Copyright 2023 The Radius Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package driver

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/go-logr/logr"
	coredm "github.com/project-radius/radius/pkg/corerp/datamodel"
	"github.com/project-radius/radius/pkg/linkrp/datamodel"
	"github.com/project-radius/radius/pkg/linkrp/processors"
	"github.com/project-radius/radius/pkg/recipes"
	"github.com/project-radius/radius/pkg/recipes/recipecontext"
	"github.com/project-radius/radius/pkg/resourcemodel"
	"github.com/project-radius/radius/pkg/rp/util"
	rpv1 "github.com/project-radius/radius/pkg/rp/v1"
	clients "github.com/project-radius/radius/pkg/sdk/clients"
	ucp_aws "github.com/project-radius/radius/pkg/ucp/aws"
	"github.com/project-radius/radius/pkg/ucp/resources"
	"github.com/project-radius/radius/pkg/ucp/ucplog"
	"golang.org/x/sync/errgroup"
)

//go:generate mockgen -destination=./mock_driver.go -package=driver -self_package github.com/project-radius/radius/pkg/recipes/driver github.com/project-radius/radius/pkg/recipes/driver Driver
const (
	deploymentPrefix      = "recipe"
	pollFrequency         = time.Second * 5
	ResultPropertyName    = "result"
	recipeParameters      = "parameters"
	deletionRetryInterval = time.Minute * 1
)

var _ Driver = (*bicepDriver)(nil)

// # Function Explanation
//
// NewBicepDriver creates a new bicep driver instance with the given ARM client options, deployment client and resource client.
func NewBicepDriver(armOptions *arm.ClientOptions, deploymentClient *clients.ResourceDeploymentsClient, client processors.ResourceClient) Driver {
	return &bicepDriver{ArmClientOptions: armOptions, DeploymentClient: deploymentClient, ResourceClient: client}
}

type bicepDriver struct {
	ArmClientOptions *arm.ClientOptions
	DeploymentClient *clients.ResourceDeploymentsClient
	ResourceClient   processors.ResourceClient
}

// # Function Explanation
//
// Execute fetches recipe contents from container registry, creates a deployment ID, a recipe context parameter, recipe parameters,
// a provider config, and deploys a bicep template for the recipe using UCP deployment client, then polls until the deployment
// is done and prepares the recipe response.
func (d *bicepDriver) Execute(ctx context.Context, configuration recipes.Configuration, recipe recipes.ResourceMetadata, definition recipes.EnvironmentDefinition) (*recipes.RecipeOutput, error) {
	logger := logr.FromContextOrDiscard(ctx)
	logger.Info(fmt.Sprintf("Deploying recipe: %q, template: %q", definition.Name, definition.TemplatePath))

	recipeData := make(map[string]any)
	err := util.ReadFromRegistry(ctx, definition.TemplatePath, &recipeData)
	if err != nil {
		return nil, err
	}
	// create the context object to be passed to the recipe deployment
	recipeContext, err := recipecontext.New(&recipe, &configuration)
	if err != nil {
		return nil, err
	}

	// get the parameters after resolving the conflict between developer and operator parameters
	// if the recipe template also has the context parameter defined then add it to the parameter for deployment
	_, isContextParameterDefined := recipeData[recipeParameters].(map[string]any)[datamodel.RecipeContextParameter]
	parameters := createRecipeParameters(recipe.Parameters, definition.Parameters, isContextParameterDefined, recipeContext)

	deploymentName := deploymentPrefix + strconv.FormatInt(time.Now().UnixNano(), 10)
	deploymentID, err := createDeploymentID(recipeContext.Resource.ID, deploymentName)
	if err != nil {
		return nil, err
	}

	// Provider config will specify the Azure and AWS scopes (if provided).
	providerConfig := newProviderConfig(deploymentID.FindScope(resources.ResourceGroupsSegment), configuration.Providers)

	logger.Info("deploying bicep template for recipe", "deploymentID", deploymentID)
	if providerConfig.AWS != nil {
		logger.Info("using AWS provider", "deploymentID", deploymentID, "scope", providerConfig.AWS.Value.Scope)
	}
	if providerConfig.Az != nil {
		logger.Info("using Azure provider", "deploymentID", deploymentID, "scope", providerConfig.Az.Value.Scope)
	}

	poller, err := d.DeploymentClient.CreateOrUpdate(
		ctx,
		clients.Deployment{
			Properties: &clients.DeploymentProperties{
				Mode:           armresources.DeploymentModeIncremental,
				ProviderConfig: &providerConfig,
				Parameters:     parameters,
				Template:       recipeData,
			},
		},
		deploymentID.String(),
		clients.DeploymentsClientAPIVersion,
	)
	if err != nil {
		return nil, err
	}

	resp, err := poller.PollUntilDone(ctx, &runtime.PollUntilDoneOptions{Frequency: pollFrequency})
	if err != nil {
		return nil, err
	}

	recipeResponse, err := prepareRecipeResponse(resp.Properties.Outputs, resp.Properties.OutputResources)
	if err != nil {
		return nil, fmt.Errorf("failed to read the recipe output %q: %w", ResultPropertyName, err)
	}

	return &recipeResponse, nil
}

// # Function Explanation
//
// Delete deletes the output resources created by the recipe deployment.
// This operation creates a goroutine for each output resource and deletes them concurrently.
func (d *bicepDriver) Delete(ctx context.Context, outputResources []rpv1.OutputResource) error {
	logger := ucplog.FromContextOrDiscard(ctx)

	// Create a waitgroup to track the deletion of each output resource
	g, groupCtx := errgroup.WithContext(ctx)

	for i := range outputResources {
		outputResource := outputResources[i]

		// Create a goroutine that handles the deletion of one resource
		g.Go(func() error {
			id := outputResource.Identity.GetID()
			logger.Info(fmt.Sprintf("Deleting output resource: %v, LocalID: %s, resource type: %s\n", outputResource.Identity, outputResource.LocalID, outputResource.ResourceType.Type))

			// If the resource is not managed by Radius, skip the deletion
			if outputResource.RadiusManaged == nil || !*outputResource.RadiusManaged {
				logger.Info(fmt.Sprintf("Skipping deletion of output resource: %q, not managed by Radius", id), id)
				return nil
			}

			// Only AWS resources are retried on deletion failure
			maxDeletionRetries := 1
			if strings.EqualFold(outputResource.ResourceType.Provider, "aws") {
				maxDeletionRetries = 5
			}

			for attempt := 1; attempt <= maxDeletionRetries; attempt++ {
				err := d.ResourceClient.Delete(groupCtx, id, resourcemodel.APIVersionUnknown)
				if err != nil && ucp_aws.IsAWSResourceNotFoundError(err) {
					// If the AWS resource is not found, then it is already deleted
					break
				} else if err != nil {
					if attempt == maxDeletionRetries {
						deletionErr := fmt.Errorf("failed to delete output resource: %q with error %v", id, err)
						logger.Error(err, fmt.Sprintf("Failed to delete output resource %q", id))
						return deletionErr
					} else {
						logger.Info(fmt.Sprintf("Failed to delete output resource: %q with error %v, retrying in %v\n", id, err, deletionRetryInterval))
						time.Sleep(deletionRetryInterval)
						continue
					}
				} else {
					// If the err is nil, then the resource is deleted successfully
					break
				}
			}

			logger.Info(fmt.Sprintf("Deleted output resource: %q", id), id)
			return nil
		})
	}

	// g.Wait waits for all goroutines to complete
	// and returns the first non-nil error returned
	// by one of the goroutines.
	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

// createRecipeParameters creates the parameters to be passed for recipe deployment after handling conflicts in parameters set by operator and developer.
// In case of conflict the developer parameter takes precedence. If recipe has context parameter defined adds the context information to the parameters list
func createRecipeParameters(devParams, operatorParams map[string]any, isCxtSet bool, recipeContext *recipecontext.Context) map[string]any {
	parameters := map[string]any{}
	for k, v := range operatorParams {
		parameters[k] = map[string]any{
			"value": v,
		}
	}
	for k, v := range devParams {
		parameters[k] = map[string]any{
			"value": v,
		}
	}
	if isCxtSet {
		parameters["context"] = map[string]any{
			"value": *recipeContext,
		}
	}
	return parameters
}

func createDeploymentID(resourceID string, deploymentName string) (resources.ID, error) {
	parsed, err := resources.ParseResource(resourceID)
	if err != nil {
		return resources.ID{}, err
	}

	resourceGroup := parsed.FindScope(resources.ResourceGroupsSegment)
	return resources.ParseResource(fmt.Sprintf("/planes/radius/local/resourceGroups/%s/providers/Microsoft.Resources/deployments/%s", resourceGroup, deploymentName))
}

func newProviderConfig(resourceGroup string, envProviders coredm.Providers) clients.ProviderConfig {
	config := clients.NewDefaultProviderConfig(resourceGroup)

	if envProviders.Azure != (coredm.ProvidersAzure{}) {
		config.Az = &clients.Az{
			Type: clients.ProviderTypeAzure,
			Value: clients.Value{
				Scope: envProviders.Azure.Scope,
			},
		}
	}

	if envProviders.AWS != (coredm.ProvidersAWS{}) {
		config.AWS = &clients.AWS{
			Type: clients.ProviderTypeAWS,
			Value: clients.Value{
				Scope: envProviders.AWS.Scope,
			},
		}
	}

	return config
}

// prepareRecipeResponse populates the recipe response from parsing the deployment output 'result' object and the
// resources created by the template.
func prepareRecipeResponse(outputs any, resources []*armresources.ResourceReference) (recipes.RecipeOutput, error) {
	// We populate the recipe response from the 'result' output (if set)
	// and the resources created by the template.
	//
	// Note that there are two ways a resource can be returned:
	// - Implicitly when it is created in the template (it will be in 'resources').
	// - Explicitly as part of the 'result' output.
	//
	// The latter is needed because non-ARM and non-UCP resources are not returned as part of the implicit 'resources'
	// collection. For us this mostly means Kubernetes resources - the user has to be explicit.
	recipeResponse := recipes.RecipeOutput{}

	out, ok := outputs.(map[string]any)
	if ok {
		recipeOutput, ok := out[ResultPropertyName].(map[string]any)
		if ok {
			output, ok := recipeOutput["value"].(map[string]any)
			if ok {
				b, err := json.Marshal(&output)
				if err != nil {
					return recipes.RecipeOutput{}, err
				}

				// Using a decoder to block unknown fields.
				decoder := json.NewDecoder(bytes.NewBuffer(b))
				decoder.DisallowUnknownFields()
				err = decoder.Decode(&recipeResponse)
				if err != nil {
					return recipes.RecipeOutput{}, err
				}
			}
		}
	}

	// process the 'resources' created by the template
	for _, id := range resources {
		recipeResponse.Resources = append(recipeResponse.Resources, *id.ID)
	}

	// Make sure our maps are non-nil (it's just friendly).
	if recipeResponse.Secrets == nil {
		recipeResponse.Secrets = map[string]any{}
	}
	if recipeResponse.Values == nil {
		recipeResponse.Values = map[string]any{}
	}

	return recipeResponse, nil
}
