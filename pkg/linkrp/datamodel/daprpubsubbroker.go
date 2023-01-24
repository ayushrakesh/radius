// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package datamodel

import (
	v1 "github.com/project-radius/radius/pkg/armrpc/api/v1"
	"github.com/project-radius/radius/pkg/linkrp"
	"github.com/project-radius/radius/pkg/rp"
	"github.com/project-radius/radius/pkg/rp/outputresource"
)

// DaprPubSubBroker represents DaprPubSubBroker link resource.
type DaprPubSubBroker struct {
	v1.BaseResource

	// Properties is the properties of the resource.
	Properties DaprPubSubBrokerProperties `json:"properties"`

	// LinkMetadata represents internal DataModel properties common to all link types.
	LinkMetadata
}

func (r *DaprPubSubBroker) Transform(outputResources []outputresource.OutputResource, computedValues map[string]any, secretValues map[string]rp.SecretValueReference) error {
	r.Properties.Status.OutputResources = outputResources
	r.ComputedValues = computedValues
	r.SecretValues = secretValues
	if topic, ok := computedValues[linkrp.TopicNameKey].(string); ok {
		r.Properties.Topic = topic
	}

	if componentName, ok := computedValues[linkrp.ComponentNameKey].(string); ok {
		r.Properties.ComponentName = componentName
	}

	return nil
}

// ApplyDeploymentOutput applies the properties changes based on the deployment output.
func (r *DaprPubSubBroker) ApplyDeploymentOutput(do rp.DeploymentOutput) {
	r.Properties.Status.OutputResources = do.DeployedOutputResources
}

// OutputResources returns the output resources array.
func (r *DaprPubSubBroker) OutputResources() []outputresource.OutputResource {
	return r.Properties.Status.OutputResources
}

// ResourceMetadata returns the application resource metadata.
func (r *DaprPubSubBroker) ResourceMetadata() *rp.BasicResourceProperties {
	return &r.Properties.BasicResourceProperties
}

// ComputedValues returns the computed values on the link.
func (r *DaprPubSubBroker) GetComputedValues() map[string]any {
	return r.LinkMetadata.ComputedValues
}

// SecretValues returns the secret values for the link.
func (r *DaprPubSubBroker) GetSecretValues() map[string]rp.SecretValueReference {
	return r.LinkMetadata.SecretValues
}

// RecipeData returns the recipe data for the link.
func (r *DaprPubSubBroker) GetRecipeData() RecipeData {
	return r.LinkMetadata.RecipeData
}

func (daprPubSub *DaprPubSubBroker) ResourceTypeName() string {
	return linkrp.DaprPubSubBrokersResourceType
}

// DaprPubSubBrokerProperties represents the properties of DaprPubSubBroker resource.
type DaprPubSubBrokerProperties struct {
	rp.BasicResourceProperties
	rp.BasicDaprResourceProperties
	Topic    string         `json:"topic,omitempty"` // Topic name of the Azure ServiceBus resource. Provided by the user.
	Mode     LinkMode       `json:"mode"`
	Metadata map[string]any `json:"metadata,omitempty"`
	Recipe   LinkRecipe     `json:"recipe"`
	Resource string         `json:"resource,omitempty"`
	Type     string         `json:"type,omitempty"`
	Version  string         `json:"version,omitempty"`
}
