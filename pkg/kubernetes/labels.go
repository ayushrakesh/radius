// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package kubernetes

import (
	"fmt"
	"strings"
)

// Commonly-used and Radius-Specific labels for Kubernetes
const (
	LabelRadiusApplication  = "radius.dev/application"
	LabelRadiusResource     = "radius.dev/resource"
	LabelRadiusDeployment   = "radius.dev/deployment"
	LabelRadiusRouteFmt     = "radius.dev/route-%s-%s"
	LabelRadiusRevision     = "radius.dev/revision"
	LabelRadiusResourceType = "radius.dev/resource-type"
	LabelPartOf             = "app.kubernetes.io/part-of"
	LabelName               = "app.kubernetes.io/name"
	LabelManagedBy          = "app.kubernetes.io/managed-by"
	LabelManagedByRadiusRP  = "radius-rp"
	LabelAADPodIdentity     = "aadpodidbinding"

	FieldManager      = "radius-rp"
	AnnotationLocalID = "radius.dev/local-id"
)

// NOTE: the difference between descriptive labels and selector labels
//
// descriptive labels:
// - intended for humans and human troubleshooting
// - we include both our own metadata and kubernetes *recommended* labels
// - we might (in the future) include non-determinisitic details because these are informational
//
// selector labels:
// - intended for programmatic matching (selecting a pod for logging)
// - no value in redundancy between our own metadata and *recommended* labels, simpler is better
// - MUST remain deterministic
// - MUST be a subset of descriptive labels
//
// In general, descriptive labels should be applied all to Kubernetes objects, and selector labels should be used
// when programmatically querying those objects.

// MakeDescriptiveLabels returns a map of the descriptive labels for a Kubernetes resource associated with a Radius resource.
// The descriptive labels are a superset of the selector labels.
func MakeDescriptiveLabels(application string, resource string) map[string]string {
	return map[string]string{
		LabelRadiusApplication: application,
		LabelRadiusResource:    resource,
		LabelName:              resource,
		LabelPartOf:            application,
		LabelManagedBy:         LabelManagedByRadiusRP,
	}
}

// MakeSelectorLablels returns a map of labels suitable for a Kubernetes selector to identify a labeled Radius-managed
// Kubernetes object.
//
// This function is used to generate the labels used by a Deployment to select its Pods. eg: the Deployment and Pods
// are the same resource.
func MakeSelectorLabels(application string, resource string) map[string]string {
	if resource != "" {
		return map[string]string{
			LabelRadiusApplication: application,
			LabelRadiusResource:    resource,
		}
	}
	return map[string]string{
		LabelRadiusApplication: application,
	}
}

// MakeRouteSelectorLabels returns a map of labels suitable for a Kubernetes selector to identify a labeled Radius-managed
// Kubernetes object.
//
// This function differs from MakeSelectorLablels in that it's intended to *cross* resources. eg: The Service created by
// an HttpRoute and the Deployment created by a Container.
func MakeRouteSelectorLabels(application string, resourceType string, route string) map[string]string {
	return map[string]string{
		LabelRadiusApplication: application,

		// NOTE: pods can serve multiple routes of different types. Therefore we need to encode the
		// the route's type and name in the *key* to support multiple matches.
		fmt.Sprintf(LabelRadiusRouteFmt, strings.ToLower(strings.TrimSuffix(resourceType, "Route")), strings.ToLower(route)): "true",
	}
}

// MakeAADPodIdentityBindingLabels returns a map binding the Pod Identity name to the pod
func MakeAADPodIdentityBindingLabels(podIdentityName string) map[string]string {
	return map[string]string{
		LabelAADPodIdentity: podIdentityName,
	}
}

// MakeRouteSelectorLabels returns a map of labels suitable for a Kubernetes selector to identify a labeled Radius-managed
// Kubernetes object.
//
// This function differs from MakeSelectorLablels in that it's intended to *cross* resources. eg: The Service created by
// an HttpRoute and the Deployment created by a Container.
func MakeResourceCRDLabels(application string, resourceType string, resource string) map[string]string {
	if resourceType != "" && resource != "" {
		return map[string]string{
			LabelRadiusApplication:  application,
			LabelRadiusResourceType: resourceType,
			LabelRadiusResource:     resource,
			LabelName:               resource,
			LabelPartOf:             application,
			LabelManagedBy:          LabelManagedByRadiusRP,
		}
	}

	return map[string]string{
		LabelRadiusApplication: application,
		LabelName:              application,
		LabelManagedBy:         LabelManagedByRadiusRP,
	}
}

func MakeResourceName(application string, resource string) string {
	if application != "" && resource != "" {
		return strings.ToLower(application + "-" + resource)
	}

	if application != "" && resource == "" {
		return strings.ToLower(application)
	}

	if application == "" && resource != "" {
		return strings.ToLower(resource)
	}

	// We should never have this case
	return "resource-name"
}