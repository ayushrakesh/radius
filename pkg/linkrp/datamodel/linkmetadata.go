// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package datamodel

import (
	"github.com/project-radius/radius/pkg/rp"
)

// LinkMetadata represents internal DataModel properties common to all link types.
type LinkMetadata struct {
	// ComputedValues map is any resource values that will be needed for more operations.
	// For example; database name to generate secrets for cosmos DB.
	ComputedValues map[string]interface{} `json:"computedValues,omitempty"`

	// Stores action to retrieve secret values. For Azure, connectionstring is accessed through cosmos listConnectionString operation, if secrets are not provided as input
	SecretValues map[string]rp.SecretValueReference `json:"secretValues,omitempty"`

	RecipeData RecipeData `json:"recipeData,omitempty"`
}

type RecipeData struct {
	RecipeProperties

	Provider string

	// API version to use to perform operations on resources supported by the link.
	// For example for Azure resources, every service has different REST API version that must be specified in the request.
	APIVersion string

	// Resource ids of the resources deployed by the recipe
	Resources []string
}

type RecipeProperties struct {
	LinkRecipe
	LinkType     string
	TemplatePath string
}

// LinkRecipe is the recipe details used to automatically deploy underlying infrastructure for a link
type LinkRecipe struct {
	// Name of the recipe within the environment to use
	Name string `json:"name,omitempty"`
	// Parameters are key/value parameters to pass into the recipe at deployment
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}