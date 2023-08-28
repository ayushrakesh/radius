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

package engine

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/project-radius/radius/pkg/corerp/datamodel"
	"github.com/project-radius/radius/pkg/recipes"
	"github.com/project-radius/radius/pkg/recipes/configloader"
	recipedriver "github.com/project-radius/radius/pkg/recipes/driver"
	rpv1 "github.com/project-radius/radius/pkg/rp/v1"
	"github.com/project-radius/radius/pkg/to"
	"github.com/project-radius/radius/pkg/ucp/resources"
	resources_kubernetes "github.com/project-radius/radius/pkg/ucp/resources/kubernetes"
	"github.com/project-radius/radius/test/testcontext"
	"github.com/stretchr/testify/require"
)

func setup(t *testing.T) (engine, configloader.MockConfigurationLoader, recipedriver.MockDriver) {
	ctrl := gomock.NewController(t)
	configLoader := configloader.NewMockConfigurationLoader(ctrl)
	mDriver := recipedriver.NewMockDriver(ctrl)
	options := Options{
		ConfigurationLoader: configLoader,
		Drivers: map[string]recipedriver.Driver{
			recipes.TemplateKindBicep:     mDriver,
			recipes.TemplateKindTerraform: mDriver,
		},
	}
	engine := engine{
		options: options,
	}
	return engine, *configLoader, *mDriver
}

func Test_Engine_Execute_Success(t *testing.T) {
	recipeMetadata := recipes.ResourceMetadata{
		Name:          "mongo-azure",
		ApplicationID: "/planes/radius/local/resourcegroups/test-rg/providers/applications.core/applications/app1",
		EnvironmentID: "/planes/radius/local/resourcegroups/test-rg/providers/applications.core/environments/env1",
		ResourceID:    "/planes/radius/local/resourceGroups/test-rg/providers/Microsoft.Resources/deployments/recipe",
		Parameters: map[string]any{
			"resourceName": "resource1",
		},
	}
	envConfig := &recipes.Configuration{
		Runtime: recipes.RuntimeConfiguration{
			Kubernetes: &recipes.KubernetesRuntime{
				Namespace: "default",
			},
		},
		Providers: datamodel.Providers{
			Azure: datamodel.ProvidersAzure{
				Scope: "scope",
			},
		},
	}
	recipeResult := &recipes.RecipeOutput{
		Resources: []string{"mongoStorageAccount", "mongoDatabase"},
		Secrets: map[string]any{
			"connectionString": "mongodb://testUser:testPassword@testAccount1.mongo.cosmos.azure.com:10255",
		},
		Values: map[string]any{
			"host": "testAccount1.mongo.cosmos.azure.com",
			"port": 10255,
		},
	}
	recipeDefinition := &recipes.EnvironmentDefinition{
		Driver:       recipes.TemplateKindBicep,
		TemplatePath: "radiusdev.azurecr.io/recipes/functionaltest/basic/mongodatabases/azure:1.0",
		ResourceType: "Applications.Link/mongoDatabases",
	}
	ctx := testcontext.New(t)
	engine, configLoader, driver := setup(t)

	configLoader.EXPECT().
		LoadConfiguration(ctx, recipeMetadata).
		Times(1).
		Return(envConfig, nil)
	configLoader.EXPECT().
		LoadRecipe(ctx, &recipeMetadata).
		Times(1).
		Return(recipeDefinition, nil)
	driver.EXPECT().
		Execute(ctx, recipedriver.ExecuteOptions{
			BaseOptions: recipedriver.BaseOptions{
				Configuration: *envConfig,
				Recipe:        recipeMetadata,
				Definition:    *recipeDefinition,
			},
		}).
		Times(1).
		Return(recipeResult, nil)

	result, err := engine.Execute(ctx, recipeMetadata)
	require.NoError(t, err)
	require.Equal(t, result, recipeResult)
}

func Test_Engine_Execute_Failure(t *testing.T) {
	recipeMetadata := recipes.ResourceMetadata{
		Name:          "mongo-azure",
		ApplicationID: "/planes/radius/local/resourcegroups/test-rg/providers/applications.core/applications/app1",
		EnvironmentID: "/planes/radius/local/resourcegroups/test-rg/providers/applications.core/environments/env1",
		ResourceID:    "/planes/radius/local/resourceGroups/test-rg/providers/Microsoft.Resources/deployments/recipe",
		Parameters: map[string]any{
			"resourceName": "resource1",
		},
	}
	envConfig := &recipes.Configuration{
		Runtime: recipes.RuntimeConfiguration{
			Kubernetes: &recipes.KubernetesRuntime{
				Namespace: "default",
			},
		},
		Providers: datamodel.Providers{
			Azure: datamodel.ProvidersAzure{
				Scope: "scope",
			},
		},
	}
	recipeDefinition := &recipes.EnvironmentDefinition{
		Driver:       recipes.TemplateKindBicep,
		TemplatePath: "radiusdev.azurecr.io/recipes/functionaltest/basic/mongodatabases/azure:1.0",
		ResourceType: "Applications.Link/mongoDatabases",
	}
	ctx := testcontext.New(t)
	engine, configLoader, driver := setup(t)

	configLoader.EXPECT().
		LoadConfiguration(ctx, recipeMetadata).
		Times(1).
		Return(envConfig, nil)
	configLoader.EXPECT().
		LoadRecipe(ctx, &recipeMetadata).
		Times(1).
		Return(recipeDefinition, nil)
	driver.EXPECT().
		Execute(ctx, recipedriver.ExecuteOptions{
			BaseOptions: recipedriver.BaseOptions{
				Configuration: *envConfig,
				Recipe:        recipeMetadata,
				Definition:    *recipeDefinition,
			},
		}).
		Times(1).
		Return(nil, errors.New("failed to execute recipe"))

	result, err := engine.Execute(ctx, recipeMetadata)
	require.Nil(t, result)
	require.Error(t, err)
	require.Equal(t, err.Error(), "failed to execute recipe")
}

func Test_Engine_Terraform_Success(t *testing.T) {
	recipeMetadata := recipes.ResourceMetadata{
		Name:          "mongo-azure",
		ApplicationID: "/planes/radius/local/resourcegroups/test-rg/providers/applications.core/applications/app1",
		EnvironmentID: "/planes/radius/local/resourcegroups/test-rg/providers/applications.core/environments/env1",
		ResourceID:    "/planes/deployments/local/resourceGroups/test-rg/providers/Microsoft.Resources/deployments/recipe",
		Parameters: map[string]any{
			"resourceName": "resource1",
		},
	}
	envConfig := &recipes.Configuration{
		Runtime: recipes.RuntimeConfiguration{
			Kubernetes: &recipes.KubernetesRuntime{
				Namespace: "default",
			},
		},
		Providers: datamodel.Providers{
			Azure: datamodel.ProvidersAzure{
				Scope: "scope",
			},
		},
	}
	recipeResult := &recipes.RecipeOutput{
		Resources: []string{"mongoStorageAccount", "mongoDatabase"},
		Secrets: map[string]any{
			"connectionString": "mongodb://testUser:testPassword@testAccount1.mongo.cosmos.azure.com:10255",
		},
		Values: map[string]any{
			"host": "testAccount1.mongo.cosmos.azure.com",
			"port": 10255,
		},
	}
	recipeDefinition := &recipes.EnvironmentDefinition{
		Driver:          recipes.TemplateKindTerraform,
		TemplatePath:    "Azure/redis/azurerm",
		TemplateVersion: "1.1.0",
		ResourceType:    "Applications.Link/mongoDatabases",
	}
	ctx := testcontext.New(t)
	engine, configLoader, driver := setup(t)

	configLoader.EXPECT().
		LoadConfiguration(ctx, recipeMetadata).
		Times(1).
		Return(envConfig, nil)
	configLoader.EXPECT().
		LoadRecipe(ctx, &recipeMetadata).
		Times(1).
		Return(recipeDefinition, nil)
	driver.EXPECT().
		Execute(ctx, recipedriver.ExecuteOptions{
			BaseOptions: recipedriver.BaseOptions{
				Configuration: *envConfig,
				Recipe:        recipeMetadata,
				Definition:    *recipeDefinition,
			},
		}).
		Times(1).
		Return(recipeResult, nil)

	result, err := engine.Execute(ctx, recipeMetadata)
	require.NoError(t, err)
	require.Equal(t, result, recipeResult)
}

func Test_Engine_InvalidDriver(t *testing.T) {
	ctx := testcontext.New(t)
	engine, configLoader, _ := setup(t)

	recipeDefinition := &recipes.EnvironmentDefinition{
		Driver:       "invalid",
		TemplatePath: "radiusdev.azurecr.io/recipes/functionaltest/basic/mongodatabases/azure:1.0",
		ResourceType: "Applications.Link/mongoDatabases",
	}

	recipeMetadata := recipes.ResourceMetadata{
		Name:          "mongo-azure",
		ApplicationID: "/planes/radius/local/resourcegroups/test-rg/providers/applications.core/applications/app1",
		EnvironmentID: "/planes/radius/local/resourcegroups/test-rg/providers/applications.core/environments/env1",
		ResourceID:    "/planes/radius/local/resourceGroups/test-rg/providers/Microsoft.Resources/deployments/recipe",
		Parameters: map[string]any{
			"resourceName": "resource1",
		},
	}

	configLoader.EXPECT().
		LoadRecipe(ctx, &recipeMetadata).
		Times(1).
		Return(recipeDefinition, nil)
	_, err := engine.Execute(ctx, recipeMetadata)
	require.Error(t, err)
	require.Equal(t, err.Error(), "could not find driver invalid")
}

func Test_Engine_Lookup_Error(t *testing.T) {
	ctx := testcontext.New(t)
	engine, configLoader, _ := setup(t)
	recipeMetadata := recipes.ResourceMetadata{
		Name:          "mongo-azure",
		ApplicationID: "/planes/radius/local/resourcegroups/test-rg/providers/applications.core/applications/app1",
		EnvironmentID: "/planes/radius/local/resourcegroups/test-rg/providers/applications.core/environments/env1",
		ResourceID:    "/planes/radius/local/resourceGroups/test-rg/providers/Microsoft.Resources/deployments/recipe",
		Parameters: map[string]any{
			"resourceName": "resource1",
		},
	}
	configLoader.EXPECT().
		LoadRecipe(ctx, &recipeMetadata).
		Times(1).
		Return(nil, errors.New("could not find recipe mongo-azure in environment env1"))
	_, err := engine.Execute(ctx, recipeMetadata)
	require.Error(t, err)
}

func Test_Engine_Load_Error(t *testing.T) {
	ctx := testcontext.New(t)
	engine, configLoader, _ := setup(t)
	recipeMetadata := recipes.ResourceMetadata{
		Name:          "mongo-azure",
		ApplicationID: "/planes/radius/local/resourcegroups/test-rg/providers/applications.core/applications/app1",
		EnvironmentID: "/planes/radius/local/resourcegroups/test-rg/providers/applications.core/environments/env1",
		ResourceID:    "/planes/radius/local/resourceGroups/test-rg/providers/Microsoft.Resources/deployments/recipe",
		Parameters: map[string]any{
			"resourceName": "resource1",
		},
	}
	recipeDefinition := &recipes.EnvironmentDefinition{
		Driver:       recipes.TemplateKindBicep,
		TemplatePath: "radiusdev.azurecr.io/recipes/functionaltest/basic/mongodatabases/azure:1.0",
		ResourceType: "Applications.Link/mongoDatabases",
	}
	configLoader.EXPECT().
		LoadRecipe(ctx, &recipeMetadata).
		Times(1).
		Return(recipeDefinition, nil)
	configLoader.EXPECT().
		LoadConfiguration(ctx, recipeMetadata).
		Times(1).
		Return(nil, errors.New("unable to fetch namespace information"))
	_, err := engine.Execute(ctx, recipeMetadata)
	require.Error(t, err)
}

func Test_Engine_Delete_Success(t *testing.T) {
	recipeMetadata, recipeDefinition, outputResources := getDeleteInputs()

	envConfig := &recipes.Configuration{
		Runtime: recipes.RuntimeConfiguration{
			Kubernetes: &recipes.KubernetesRuntime{
				Namespace: "default",
			},
		},
		Providers: datamodel.Providers{
			Azure: datamodel.ProvidersAzure{
				Scope: "scope",
			},
		},
	}

	ctx := testcontext.New(t)
	engine, configLoader, driver := setup(t)

	configLoader.EXPECT().
		LoadRecipe(ctx, &recipeMetadata).
		Times(1).
		Return(&recipeDefinition, nil)

	configLoader.EXPECT().
		LoadConfiguration(ctx, recipeMetadata).
		Times(1).
		Return(envConfig, nil)

	driver.EXPECT().
		Delete(ctx, recipedriver.DeleteOptions{
			BaseOptions: recipedriver.BaseOptions{
				Configuration: *envConfig,
				Recipe:        recipeMetadata,
				Definition:    recipeDefinition,
			},
			OutputResources: outputResources,
		}).
		Times(1).
		Return(nil)

	err := engine.Delete(ctx, recipeMetadata, outputResources)
	require.NoError(t, err)
}

func Test_Engine_Delete_Error(t *testing.T) {
	recipeMetadata, recipeDefinition, outputResources := getDeleteInputs()

	envConfig := &recipes.Configuration{
		Runtime: recipes.RuntimeConfiguration{
			Kubernetes: &recipes.KubernetesRuntime{
				Namespace: "default",
			},
		},
		Providers: datamodel.Providers{
			Azure: datamodel.ProvidersAzure{
				Scope: "scope",
			},
		},
	}

	ctx := testcontext.New(t)
	engine, configLoader, driver := setup(t)

	configLoader.EXPECT().
		LoadRecipe(ctx, &recipeMetadata).
		Times(1).
		Return(&recipeDefinition, nil)

	configLoader.EXPECT().
		LoadConfiguration(ctx, recipeMetadata).
		Times(1).
		Return(envConfig, nil)

	driver.EXPECT().
		Delete(ctx, recipedriver.DeleteOptions{
			BaseOptions: recipedriver.BaseOptions{
				Configuration: *envConfig,
				Recipe:        recipeMetadata,
				Definition:    recipeDefinition,
			},
			OutputResources: outputResources,
		}).
		Times(1).
		Return(fmt.Errorf("could not find API version for type %q, no supported API versions",
			outputResources[0].ID))

	err := engine.Delete(ctx, recipeMetadata, outputResources)
	require.Error(t, err)
}

func Test_Delete_InvalidDriver(t *testing.T) {
	recipeMetadata, recipeDefinition, outputResources := getDeleteInputs()
	recipeDefinition.Driver = "invalid"

	ctx := testcontext.New(t)
	engine, configLoader, _ := setup(t)

	configLoader.EXPECT().
		LoadRecipe(ctx, &recipeMetadata).
		Times(1).
		Return(&recipeDefinition, nil)
	err := engine.Delete(ctx, recipeMetadata, outputResources)
	require.Error(t, err)
	require.Equal(t, err.Error(), "could not find driver invalid")
}

func Test_Delete_Lookup_Error(t *testing.T) {
	ctx := testcontext.New(t)
	engine, configLoader, _ := setup(t)
	recipeMetadata, _, outputResources := getDeleteInputs()

	configLoader.EXPECT().
		LoadRecipe(ctx, &recipeMetadata).
		Times(1).
		Return(nil, errors.New("could not find recipe mongo-azure in environment env1"))
	err := engine.Delete(ctx, recipeMetadata, outputResources)
	require.Error(t, err)
}

func getDeleteInputs() (recipes.ResourceMetadata, recipes.EnvironmentDefinition, []rpv1.OutputResource) {
	recipeMetadata := recipes.ResourceMetadata{
		Name:          "mongo-azure",
		ApplicationID: "/planes/radius/local/resourcegroups/test-rg/providers/applications.core/applications/app1",
		EnvironmentID: "/planes/radius/local/resourcegroups/test-rg/providers/applications.core/environments/env1",
		ResourceID:    "/planes/radius/local/resourceGroups/test-rg/providers/Applications.Link/mongoDatabases/test-db",
		Parameters: map[string]any{
			"resourceName": "resource1",
		},
	}

	recipeDefinition := recipes.EnvironmentDefinition{
		Driver:       recipes.TemplateKindBicep,
		TemplatePath: "radiusdev.azurecr.io/recipes/functionaltest/basic/mongodatabases/azure:1.0",
		ResourceType: "Applications.Link/mongoDatabases",
	}

	outputResources := []rpv1.OutputResource{
		{
			ID: resources.MustParse("/subscriptions/test-sub/resourcegroups/test-rg/providers/Microsoft.DocumentDB/accounts/test-account"),
		},
	}
	return recipeMetadata, recipeDefinition, outputResources
}

func Test_GarbageCollectResources(t *testing.T) {
	outputResources := []rpv1.OutputResource{
		{
			ID: resources_kubernetes.IDFromParts(
				resources_kubernetes.PlaneNameTODO,
				"core",
				"Deployment",
				"recipe-app",
				"redis"),
			RadiusManaged: to.Ptr(true),
		},
	}
	recipeMetadata := recipes.ResourceMetadata{
		Name:          "mongo-azure",
		ApplicationID: "/planes/radius/local/resourcegroups/test-rg/providers/applications.core/applications/app1",
		EnvironmentID: "/planes/radius/local/resourcegroups/test-rg/providers/applications.core/environments/env1",
		ResourceID:    "/planes/radius/local/resourceGroups/test-rg/providers/Microsoft.Resources/deployments/recipe",
		Parameters: map[string]any{
			"resourceName": "resource1",
		},
	}
	recipeDefinition := &recipes.EnvironmentDefinition{
		Driver:       recipes.TemplateKindBicep,
		TemplatePath: "radiusdev.azurecr.io/recipes/functionaltest/basic/mongodatabases/azure:1.0",
		ResourceType: "Applications.Link/mongoDatabases",
	}
	ctx := testcontext.New(t)
	engine, configLoader, driver := setup(t)

	tests := []struct {
		desc string
		err  error
	}{
		{
			desc: "success",
			err:  nil,
		},
		{
			desc: "deletion failed",
			err:  errors.New("test-error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			configLoader.EXPECT().
				LoadRecipe(ctx, &recipeMetadata).
				Times(1).
				Return(recipeDefinition, nil)
			driver.EXPECT().
				GarbageCollectResources(ctx, outputResources).
				Times(1).Return(tt.err)
			err := engine.GarbageCollectResources(ctx, recipeMetadata, outputResources)
			if tt.err != nil {
				require.Equal(t, tt.err.Error(), "test-error")
			} else {
				require.NoError(t, err)
			}
		})
	}
}
