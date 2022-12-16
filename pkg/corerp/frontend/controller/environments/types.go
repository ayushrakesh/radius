// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package environments

const (
	ResourceTypeName = "Applications.Core/environments"
)

func supportedProviders() []string {
	return []string{"azure"}
}
