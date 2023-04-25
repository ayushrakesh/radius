// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package processors

import (
	"context"
	"fmt"

	"github.com/project-radius/radius/pkg/recipes"
	rpv1 "github.com/project-radius/radius/pkg/rp/v1"
)

// ResourceProcessor is responsible for processing the results of recipe execution or any
// other change to the lifecycle of a link resource. Each resource processor supports a single
// Radius resource type (eg: RedisCache).
type ResourceProcessor[P interface {
	*T
	rpv1.RadiusResourceModel
}, T any] interface {
	// Process is called to process the results of recipe execution or any other changes to the resource
	// data model. Process should modify the datamodel in place to perform updates.
	Process(ctx context.Context, resource P, output *recipes.RecipeOutput) error
}

//go:generate mockgen -destination=./mock_resourceclient.go -package=processors -self_package github.com/project-radius/radius/pkg/linkrp/processors github.com/project-radius/radius/pkg/linkrp/processors ResourceClient

// ResourceClient is a client used by resource processors for interacting with UCP resources.
type ResourceClient interface {
	// Delete deletes a resource by id.
	//
	// If the API version is omitted, then an attempt will be made to look up the API version.
	Delete(ctx context.Context, id string, apiVersion string) error
}

// ResourceError represents an error that occured while processing a resource.
type ResourceError struct {
	ID    string
	Inner error
}

// Error formats the error as a string.
func (e *ResourceError) Error() string {
	return fmt.Sprintf("failed to delete resource %q: %v", e.ID, e.Inner)
}

// Unwrap gets the wrapper error of this ResourceDeletionErr.
func (e *ResourceError) Unwrap() error {
	return e.Inner
}

// ValidationError represents a user-facing validation message reported by the processor.
type ValidationError struct {
	Message string
}

// Error gets the string representation of the error.
func (e *ValidationError) Error() string {
	return e.Message
}
