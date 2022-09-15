//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package v20220315privatepreview

const (
	moduleName = "v20220315privatepreview"
	moduleVersion = "v0.0.1"
)

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication CreatedByType = "Application"
	CreatedByTypeKey CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{	
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// DaprPubSubBrokerPropertiesKind - The DaprPubSubProperties kind
type DaprPubSubBrokerPropertiesKind string

const (
	DaprPubSubBrokerPropertiesKindGeneric DaprPubSubBrokerPropertiesKind = "generic"
	DaprPubSubBrokerPropertiesKindPubsubAzureServicebus DaprPubSubBrokerPropertiesKind = "pubsub.azure.servicebus"
)

// PossibleDaprPubSubBrokerPropertiesKindValues returns the possible values for the DaprPubSubBrokerPropertiesKind const type.
func PossibleDaprPubSubBrokerPropertiesKindValues() []DaprPubSubBrokerPropertiesKind {
	return []DaprPubSubBrokerPropertiesKind{	
		DaprPubSubBrokerPropertiesKindGeneric,
		DaprPubSubBrokerPropertiesKindPubsubAzureServicebus,
	}
}

// DaprSecretStorePropertiesKind - Radius kind for Dapr Secret Store
type DaprSecretStorePropertiesKind string

const (
	DaprSecretStorePropertiesKindGeneric DaprSecretStorePropertiesKind = "generic"
)

// PossibleDaprSecretStorePropertiesKindValues returns the possible values for the DaprSecretStorePropertiesKind const type.
func PossibleDaprSecretStorePropertiesKindValues() []DaprSecretStorePropertiesKind {
	return []DaprSecretStorePropertiesKind{	
		DaprSecretStorePropertiesKindGeneric,
	}
}

// DaprStateStorePropertiesKind - The Dapr StateStore kind
type DaprStateStorePropertiesKind string

const (
	DaprStateStorePropertiesKindGeneric DaprStateStorePropertiesKind = "generic"
	DaprStateStorePropertiesKindStateAzureTablestorage DaprStateStorePropertiesKind = "state.azure.tablestorage"
	DaprStateStorePropertiesKindStateSqlserver DaprStateStorePropertiesKind = "state.sqlserver"
)

// PossibleDaprStateStorePropertiesKindValues returns the possible values for the DaprStateStorePropertiesKind const type.
func PossibleDaprStateStorePropertiesKindValues() []DaprStateStorePropertiesKind {
	return []DaprStateStorePropertiesKind{	
		DaprStateStorePropertiesKindGeneric,
		DaprStateStorePropertiesKindStateAzureTablestorage,
		DaprStateStorePropertiesKindStateSqlserver,
	}
}

// ProvisioningState - Provisioning state of the connector at the time the operation was called
type ProvisioningState string

const (
	ProvisioningStateAccepted ProvisioningState = "Accepted"
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	ProvisioningStateFailed ProvisioningState = "Failed"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{	
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateProvisioning,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}
