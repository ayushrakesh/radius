// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package v1

// SystemData is the readonly metadata pertaining to creation and last modification of the resource.
// https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/common-api-contracts.md#system-metadata-for-all-azure-resources
type SystemData struct {
	// CreatedBy is a string identifier for the identity that created the resource.
	CreatedBy string `json:"createdBy,omitempty"`
	// CreatedByType is the type of identity that created the resource: user, application, managedIdentity.
	CreatedByType string `json:"createdByType,omitempty"`
	// CreatedAt is the timestamp of resource creation (UTC).
	CreatedAt string `json:"createdAt,omitempty"`
	// LastModifiedBy is a string identifier for the identity that last modified the resource.
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`
	// LastModifiedBy is the type of identity that last modified the resource: user, application, managedIdentity
	LastModifiedByType string `json:"lastModifiedByType,omitempty"`
	// LastModifiedBy is the timestamp of resource last modification (UTC).
	LastModifiedAt string `json:"lastModifiedAt,omitempty"`
}

// UpdateSystemData creates or updates new systemdata from old and new resources.
func UpdateSystemData(old *SystemData, new *SystemData) SystemData {
	if old == nil {
		old = &SystemData{}
	}
	if new == nil {
		new = &SystemData{}
	}

	newSystemData := *old

	if old.CreatedAt == "" && new.CreatedAt != "" {
		newSystemData.CreatedAt = new.CreatedAt
		newSystemData.CreatedBy = new.CreatedBy
		newSystemData.CreatedByType = new.CreatedByType
	}

	if new.LastModifiedAt != "" {
		newSystemData.LastModifiedAt = new.LastModifiedAt
		newSystemData.LastModifiedBy = new.LastModifiedBy
		newSystemData.LastModifiedByType = new.LastModifiedByType

		// backfill
		if newSystemData.CreatedAt == "" {
			newSystemData.CreatedAt = new.LastModifiedAt
			newSystemData.CreatedBy = new.LastModifiedBy
			newSystemData.CreatedByType = new.LastModifiedByType
		}
	}

	return newSystemData
}