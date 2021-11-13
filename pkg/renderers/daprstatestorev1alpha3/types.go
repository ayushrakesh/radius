// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package daprstatestorev1alpha3

import (
	"github.com/Azure/radius/pkg/azure/azresources"
)

const (
	ResourceType = "dapr.io.StateStoreComponent"
)

var StorageAccountResourceType = azresources.KnownType{
	Types: []azresources.ResourceType{
		{
			Type: azresources.StorageStorageAccounts,
			Name: "*",
		},
	},
}

type Properties struct {
	Kind     string `json:"kind"`
	Managed  bool   `json:"managed"`
	Resource string `json:"resource"`
}
