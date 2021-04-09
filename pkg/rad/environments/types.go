// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package environments

import (
	"fmt"
)

const KindAzureCloud = "azure"

type Environment interface {
	GetName() string
	GetKind() string
}

// AzureCloudEnvironment represents an Azure Cloud Radius environment.
type AzureCloudEnvironment struct {
	Name           string `mapstructure:"name" validate:"required"`
	Kind           string `mapstructure:"kind" validate:"required"`
	SubscriptionID string `mapstructure:"subscriptionid" validate:"required"`
	ResourceGroup  string `mapstructure:"resourcegroup" validate:"required"`
	ClusterName    string `mapstructure:"clustername" validate:"required"`

	// We tolerate and allow extra fields - this helps with forwards compat.
	Properties map[string]interface{} `mapstructure:",remain"`
}

func (e *AzureCloudEnvironment) GetName() string {
	return e.Name
}

func (e *AzureCloudEnvironment) GetKind() string {
	return e.Kind
}

// GenericEnvironment represents an *unknown* kind of environment.
type GenericEnvironment struct {
	Name string `mapstructure:"name" validate:"required"`
	Kind string `mapstructure:"kind" validate:"required"`

	// Capture arbitrary other properties
	Properties map[string]interface{} `mapstructure:",remain"`
}

func (e *GenericEnvironment) GetName() string {
	return e.Name
}

func (e *GenericEnvironment) GetKind() string {
	return e.Kind
}

func RequireAzureCloud(e Environment) (*AzureCloudEnvironment, error) {
	az, ok := e.(*AzureCloudEnvironment)
	if !ok {
		return nil, fmt.Errorf("an '%v' environment is required but the kind was '%v'", KindAzureCloud, e.GetKind())
	}

	return az, nil
}
