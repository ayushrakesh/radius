//go:build go1.18
// +build go1.18

// Licensed under the Apache License, Version 2.0 . See LICENSE in the repository root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package v20220315privatepreview

import "time"

// BasicDaprResourceProperties - Basic properties of a Dapr component object.
type BasicDaprResourceProperties struct {
	// REQUIRED; Fully qualified resource ID for the environment that the link is linked to
	Environment *string `json:"environment,omitempty"`

	// Fully qualified resource ID for the application that the link is consumed by
	Application *string `json:"application,omitempty"`

	// READ-ONLY; The name of the Dapr component object. Use this value in your code when interacting with the Dapr client to
// use the Dapr component.
	ComponentName *string `json:"componentName,omitempty" azure:"ro"`

	// READ-ONLY; Status of a resource.
	Status *ResourceStatus `json:"status,omitempty" azure:"ro"`
}

// BasicResourceProperties - Basic properties of a Radius resource.
type BasicResourceProperties struct {
	// REQUIRED; Fully qualified resource ID for the environment that the link is linked to
	Environment *string `json:"environment,omitempty"`

	// Fully qualified resource ID for the application that the link is consumed by
	Application *string `json:"application,omitempty"`

	// READ-ONLY; Status of a resource.
	Status *ResourceStatus `json:"status,omitempty" azure:"ro"`
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info map[string]interface{} `json:"info,omitempty" azure:"ro"`

	// READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo `json:"additionalInfo,omitempty" azure:"ro"`

	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error details.
	Details []*ErrorDetail `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The error target.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail `json:"error,omitempty"`
}

// MongoDatabaseListSecretsResult - The secret values for the given MongoDatabase resource
type MongoDatabaseListSecretsResult struct {
	// Connection string used to connect to the target Mongo database
	ConnectionString *string `json:"connectionString,omitempty"`

	// Password to use when connecting to the target Mongo database
	Password *string `json:"password,omitempty"`

	// Username to use when connecting to the target Mongo database
	Username *string `json:"username,omitempty"`
}

// MongoDatabasePropertiesClassification provides polymorphic access to related types.
// Call the interface's GetMongoDatabaseProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *MongoDatabaseProperties, *RecipeMongoDatabaseProperties, *ResourceMongoDatabaseProperties, *ValuesMongoDatabaseProperties
type MongoDatabasePropertiesClassification interface {
	// GetMongoDatabaseProperties returns the MongoDatabaseProperties content of the underlying type.
	GetMongoDatabaseProperties() *MongoDatabaseProperties
}

// MongoDatabaseProperties - MongoDatabase portable resource properties
type MongoDatabaseProperties struct {
	// REQUIRED; Fully qualified resource ID for the environment that the link is linked to
	Environment *string `json:"environment,omitempty"`

	// REQUIRED; Discriminator property for MongoDatabaseProperties.
	Mode *string `json:"mode,omitempty"`

	// Fully qualified resource ID for the application that the link is consumed by
	Application *string `json:"application,omitempty"`

	// Secrets values provided for the resource
	Secrets *MongoDatabaseSecrets `json:"secrets,omitempty"`

	// READ-ONLY; Provisioning state of the mongo database portable resource at the time the operation was called
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Status of a resource.
	Status *ResourceStatus `json:"status,omitempty" azure:"ro"`
}

// GetMongoDatabaseProperties implements the MongoDatabasePropertiesClassification interface for type MongoDatabaseProperties.
func (m *MongoDatabaseProperties) GetMongoDatabaseProperties() *MongoDatabaseProperties { return m }

// MongoDatabaseResource - MongoDatabase portable resource
type MongoDatabaseResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// The resource-specific properties for this resource.
	Properties MongoDatabasePropertiesClassification `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// MongoDatabaseResourceListResult - The response of a MongoDatabaseResource list operation.
type MongoDatabaseResourceListResult struct {
	// REQUIRED; The MongoDatabaseResource items on this page
	Value []*MongoDatabaseResource `json:"value,omitempty"`

	// The link to the next page of items
	NextLink *string `json:"nextLink,omitempty"`
}

// MongoDatabaseSecrets - The secret values for the given MongoDatabase resource
type MongoDatabaseSecrets struct {
	// Connection string used to connect to the target Mongo database
	ConnectionString *string `json:"connectionString,omitempty"`

	// Password to use when connecting to the target Mongo database
	Password *string `json:"password,omitempty"`

	// Username to use when connecting to the target Mongo database
	Username *string `json:"username,omitempty"`
}

// MongoDatabasesClientBeginDeleteOptions contains the optional parameters for the MongoDatabasesClient.BeginDelete method.
type MongoDatabasesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MongoDatabasesClientCreateOrUpdateOptions contains the optional parameters for the MongoDatabasesClient.CreateOrUpdate
// method.
type MongoDatabasesClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// MongoDatabasesClientGetOptions contains the optional parameters for the MongoDatabasesClient.Get method.
type MongoDatabasesClientGetOptions struct {
	// placeholder for future optional parameters
}

// MongoDatabasesClientListByRootScopeOptions contains the optional parameters for the MongoDatabasesClient.ListByRootScope
// method.
type MongoDatabasesClientListByRootScopeOptions struct {
	// placeholder for future optional parameters
}

// MongoDatabasesClientListSecretsOptions contains the optional parameters for the MongoDatabasesClient.ListSecrets method.
type MongoDatabasesClientListSecretsOptions struct {
	// placeholder for future optional parameters
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType `json:"actionType,omitempty" azure:"ro"`

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
// operations.
	IsDataAction *bool `json:"isDataAction,omitempty" azure:"ro"`

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
	Origin *Origin `json:"origin,omitempty" azure:"ro"`
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
// Machine", "Restart Virtual Machine".
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
// Compute".
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
// Schedule Collections".
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// Recipe - The recipe used to automatically deploy underlying infrastructure for a link
type Recipe struct {
	// REQUIRED; The name of the recipe within the environment to use
	Name *string `json:"name,omitempty"`

	// Key/value parameters to pass into the recipe at deployment
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// RecipeMongoDatabaseProperties - MongoDatabase Properties for Mode Recipe
type RecipeMongoDatabaseProperties struct {
	// REQUIRED; Fully qualified resource ID for the environment that the link is linked to
	Environment *string `json:"environment,omitempty"`

	// REQUIRED; Discriminator property for MongoDatabaseProperties.
	Mode *string `json:"mode,omitempty"`

	// REQUIRED; The recipe used to automatically deploy underlying infrastructure for the mongodatabases portable resource
	Recipe *Recipe `json:"recipe,omitempty"`

	// Fully qualified resource ID for the application that the link is consumed by
	Application *string `json:"application,omitempty"`

	// Host name of the target Mongo database
	Host *string `json:"host,omitempty"`

	// Port value of the target Mongo database
	Port *int32 `json:"port,omitempty"`

	// Secrets values provided for the resource
	Secrets *MongoDatabaseSecrets `json:"secrets,omitempty"`

	// READ-ONLY; Database name of the target Mongo database
	Database *string `json:"database,omitempty" azure:"ro"`

	// READ-ONLY; Provisioning state of the mongo database portable resource at the time the operation was called
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Status of a resource.
	Status *ResourceStatus `json:"status,omitempty" azure:"ro"`
}

// GetMongoDatabaseProperties implements the MongoDatabasePropertiesClassification interface for type RecipeMongoDatabaseProperties.
func (r *RecipeMongoDatabaseProperties) GetMongoDatabaseProperties() *MongoDatabaseProperties {
	return &MongoDatabaseProperties{
		Mode: r.Mode,
		ProvisioningState: r.ProvisioningState,
		Secrets: r.Secrets,
		Status: r.Status,
		Environment: r.Environment,
		Application: r.Application,
	}
}

// RecipeRedisCacheProperties - RedisCache Properties for Mode Recipe
type RecipeRedisCacheProperties struct {
	// REQUIRED; Fully qualified resource ID for the environment that the link is linked to
	Environment *string `json:"environment,omitempty"`

	// REQUIRED; Discriminator property for RedisCacheProperties.
	Mode *string `json:"mode,omitempty"`

	// REQUIRED; The recipe used to automatically deploy underlying infrastructure for the rediscaches portable resource
	Recipe *Recipe `json:"recipe,omitempty"`

	// Fully qualified resource ID for the application that the link is consumed by
	Application *string `json:"application,omitempty"`

	// The host name of the target Redis cache
	Host *string `json:"host,omitempty"`

	// The port value of the target Redis cache
	Port *int32 `json:"port,omitempty"`

	// Secrets provided by resource
	Secrets *RedisCacheSecrets `json:"secrets,omitempty"`

	// READ-ONLY; Provisioning state of the redis cache portable resource at the time the operation was called
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Status of a resource.
	Status *ResourceStatus `json:"status,omitempty" azure:"ro"`

	// READ-ONLY; The username for Redis cache
	Username *string `json:"username,omitempty" azure:"ro"`
}

// GetRedisCacheProperties implements the RedisCachePropertiesClassification interface for type RecipeRedisCacheProperties.
func (r *RecipeRedisCacheProperties) GetRedisCacheProperties() *RedisCacheProperties {
	return &RedisCacheProperties{
		Mode: r.Mode,
		ProvisioningState: r.ProvisioningState,
		Secrets: r.Secrets,
		Status: r.Status,
		Environment: r.Environment,
		Application: r.Application,
	}
}

// RecipeSQLDatabaseProperties - SqlDatabase Properties for Mode Recipe
type RecipeSQLDatabaseProperties struct {
	// REQUIRED; Fully qualified resource ID for the environment that the link is linked to
	Environment *string `json:"environment,omitempty"`

	// REQUIRED; Discriminator property for SqlDatabaseProperties.
	Mode *string `json:"mode,omitempty"`

	// REQUIRED; The recipe used to automatically deploy underlying infrastructure for the sqldatabases portable resource
	Recipe *Recipe `json:"recipe,omitempty"`

	// Fully qualified resource ID for the application that the link is consumed by
	Application *string `json:"application,omitempty"`

	// The name of the Sql database.
	Database *string `json:"database,omitempty"`

	// The fully qualified domain name of the Sql database.
	Server *string `json:"server,omitempty"`

	// READ-ONLY; Provisioning state of the Sql database portable resource at the time the operation was called
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Status of a resource.
	Status *ResourceStatus `json:"status,omitempty" azure:"ro"`
}

// GetSQLDatabaseProperties implements the SQLDatabasePropertiesClassification interface for type RecipeSQLDatabaseProperties.
func (r *RecipeSQLDatabaseProperties) GetSQLDatabaseProperties() *SQLDatabaseProperties {
	return &SQLDatabaseProperties{
		Mode: r.Mode,
		ProvisioningState: r.ProvisioningState,
		Status: r.Status,
		Environment: r.Environment,
		Application: r.Application,
	}
}

// RedisCacheListSecretsResult - The secret values for the given RedisCache resource
type RedisCacheListSecretsResult struct {
	// The connection string used to connect to the Redis cache
	ConnectionString *string `json:"connectionString,omitempty"`

	// The password for this Redis cache instance
	Password *string `json:"password,omitempty"`
}

// RedisCachePropertiesClassification provides polymorphic access to related types.
// Call the interface's GetRedisCacheProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *RecipeRedisCacheProperties, *RedisCacheProperties, *ResourceRedisCacheProperties, *ValuesRedisCacheProperties
type RedisCachePropertiesClassification interface {
	// GetRedisCacheProperties returns the RedisCacheProperties content of the underlying type.
	GetRedisCacheProperties() *RedisCacheProperties
}

// RedisCacheProperties - RedisCache portable resource properties
type RedisCacheProperties struct {
	// REQUIRED; Fully qualified resource ID for the environment that the link is linked to
	Environment *string `json:"environment,omitempty"`

	// REQUIRED; Discriminator property for RedisCacheProperties.
	Mode *string `json:"mode,omitempty"`

	// Fully qualified resource ID for the application that the link is consumed by
	Application *string `json:"application,omitempty"`

	// Secrets provided by resource
	Secrets *RedisCacheSecrets `json:"secrets,omitempty"`

	// READ-ONLY; Provisioning state of the redis cache portable resource at the time the operation was called
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Status of a resource.
	Status *ResourceStatus `json:"status,omitempty" azure:"ro"`
}

// GetRedisCacheProperties implements the RedisCachePropertiesClassification interface for type RedisCacheProperties.
func (r *RedisCacheProperties) GetRedisCacheProperties() *RedisCacheProperties { return r }

// RedisCacheResource - RedisCache portable resource
type RedisCacheResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// The resource-specific properties for this resource.
	Properties RedisCachePropertiesClassification `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// RedisCacheResourceListResult - The response of a RedisCacheResource list operation.
type RedisCacheResourceListResult struct {
	// REQUIRED; The RedisCacheResource items on this page
	Value []*RedisCacheResource `json:"value,omitempty"`

	// The link to the next page of items
	NextLink *string `json:"nextLink,omitempty"`
}

// RedisCacheSecrets - The secret values for the given RedisCache resource
type RedisCacheSecrets struct {
	// The connection string used to connect to the Redis cache
	ConnectionString *string `json:"connectionString,omitempty"`

	// The password for this Redis cache instance
	Password *string `json:"password,omitempty"`
}

// RedisCachesClientCreateOrUpdateOptions contains the optional parameters for the RedisCachesClient.CreateOrUpdate method.
type RedisCachesClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// RedisCachesClientDeleteOptions contains the optional parameters for the RedisCachesClient.Delete method.
type RedisCachesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// RedisCachesClientGetOptions contains the optional parameters for the RedisCachesClient.Get method.
type RedisCachesClientGetOptions struct {
	// placeholder for future optional parameters
}

// RedisCachesClientListByRootScopeOptions contains the optional parameters for the RedisCachesClient.ListByRootScope method.
type RedisCachesClientListByRootScopeOptions struct {
	// placeholder for future optional parameters
}

// RedisCachesClientListSecretsOptions contains the optional parameters for the RedisCachesClient.ListSecrets method.
type RedisCachesClientListSecretsOptions struct {
	// placeholder for future optional parameters
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ResourceMongoDatabaseProperties - MongoDatabase Properties for Mode Resource
type ResourceMongoDatabaseProperties struct {
	// REQUIRED; Fully qualified resource ID for the environment that the link is linked to
	Environment *string `json:"environment,omitempty"`

	// REQUIRED; Discriminator property for MongoDatabaseProperties.
	Mode *string `json:"mode,omitempty"`

	// REQUIRED; Fully qualified resource ID of a supported resource with Mongo API to use for this portable resource
	Resource *string `json:"resource,omitempty"`

	// Fully qualified resource ID for the application that the link is consumed by
	Application *string `json:"application,omitempty"`

	// Host name of the target Mongo database
	Host *string `json:"host,omitempty"`

	// Port value of the target Mongo database
	Port *int32 `json:"port,omitempty"`

	// Secrets values provided for the resource
	Secrets *MongoDatabaseSecrets `json:"secrets,omitempty"`

	// READ-ONLY; Database name of the target Mongo database
	Database *string `json:"database,omitempty" azure:"ro"`

	// READ-ONLY; Provisioning state of the mongo database portable resource at the time the operation was called
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Status of a resource.
	Status *ResourceStatus `json:"status,omitempty" azure:"ro"`
}

// GetMongoDatabaseProperties implements the MongoDatabasePropertiesClassification interface for type ResourceMongoDatabaseProperties.
func (r *ResourceMongoDatabaseProperties) GetMongoDatabaseProperties() *MongoDatabaseProperties {
	return &MongoDatabaseProperties{
		Mode: r.Mode,
		ProvisioningState: r.ProvisioningState,
		Secrets: r.Secrets,
		Status: r.Status,
		Environment: r.Environment,
		Application: r.Application,
	}
}

// ResourceRedisCacheProperties - RedisCache Properties for Mode Resource
type ResourceRedisCacheProperties struct {
	// REQUIRED; Fully qualified resource ID for the environment that the link is linked to
	Environment *string `json:"environment,omitempty"`

	// REQUIRED; Discriminator property for RedisCacheProperties.
	Mode *string `json:"mode,omitempty"`

	// REQUIRED; Fully qualified resource ID of a supported resource with Redis API to use for this portable resource
	Resource *string `json:"resource,omitempty"`

	// Fully qualified resource ID for the application that the link is consumed by
	Application *string `json:"application,omitempty"`

	// The host name of the target Redis cache
	Host *string `json:"host,omitempty"`

	// The port value of the target Redis cache
	Port *int32 `json:"port,omitempty"`

	// Secrets provided by resource
	Secrets *RedisCacheSecrets `json:"secrets,omitempty"`

	// READ-ONLY; Provisioning state of the redis cache portable resource at the time the operation was called
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Status of a resource.
	Status *ResourceStatus `json:"status,omitempty" azure:"ro"`

	// READ-ONLY; The username for Redis cache
	Username *string `json:"username,omitempty" azure:"ro"`
}

// GetRedisCacheProperties implements the RedisCachePropertiesClassification interface for type ResourceRedisCacheProperties.
func (r *ResourceRedisCacheProperties) GetRedisCacheProperties() *RedisCacheProperties {
	return &RedisCacheProperties{
		Mode: r.Mode,
		ProvisioningState: r.ProvisioningState,
		Secrets: r.Secrets,
		Status: r.Status,
		Environment: r.Environment,
		Application: r.Application,
	}
}

// ResourceSQLDatabaseProperties - SqlDatabase Properties for Mode Resource
type ResourceSQLDatabaseProperties struct {
	// REQUIRED; Fully qualified resource ID for the environment that the link is linked to
	Environment *string `json:"environment,omitempty"`

	// REQUIRED; Discriminator property for SqlDatabaseProperties.
	Mode *string `json:"mode,omitempty"`

	// REQUIRED; Fully qualified resource ID of a supported resource with Sql API to use for this portable resource
	Resource *string `json:"resource,omitempty"`

	// Fully qualified resource ID for the application that the link is consumed by
	Application *string `json:"application,omitempty"`

	// The name of the Sql database.
	Database *string `json:"database,omitempty"`

	// The fully qualified domain name of the Sql database.
	Server *string `json:"server,omitempty"`

	// READ-ONLY; Provisioning state of the Sql database portable resource at the time the operation was called
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Status of a resource.
	Status *ResourceStatus `json:"status,omitempty" azure:"ro"`
}

// GetSQLDatabaseProperties implements the SQLDatabasePropertiesClassification interface for type ResourceSQLDatabaseProperties.
func (r *ResourceSQLDatabaseProperties) GetSQLDatabaseProperties() *SQLDatabaseProperties {
	return &SQLDatabaseProperties{
		Mode: r.Mode,
		ProvisioningState: r.ProvisioningState,
		Status: r.Status,
		Environment: r.Environment,
		Application: r.Application,
	}
}

// ResourceStatus - Status of a resource.
type ResourceStatus struct {
	// Properties of an output resource
	OutputResources []map[string]interface{} `json:"outputResources,omitempty"`
}

// SQLDatabasePropertiesClassification provides polymorphic access to related types.
// Call the interface's GetSQLDatabaseProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *RecipeSQLDatabaseProperties, *ResourceSQLDatabaseProperties, *SQLDatabaseProperties, *ValuesSQLDatabaseProperties
type SQLDatabasePropertiesClassification interface {
	// GetSQLDatabaseProperties returns the SQLDatabaseProperties content of the underlying type.
	GetSQLDatabaseProperties() *SQLDatabaseProperties
}

// SQLDatabaseProperties - SqlDatabase properties
type SQLDatabaseProperties struct {
	// REQUIRED; Fully qualified resource ID for the environment that the link is linked to
	Environment *string `json:"environment,omitempty"`

	// REQUIRED; Discriminator property for SqlDatabaseProperties.
	Mode *string `json:"mode,omitempty"`

	// Fully qualified resource ID for the application that the link is consumed by
	Application *string `json:"application,omitempty"`

	// READ-ONLY; Provisioning state of the Sql database portable resource at the time the operation was called
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Status of a resource.
	Status *ResourceStatus `json:"status,omitempty" azure:"ro"`
}

// GetSQLDatabaseProperties implements the SQLDatabasePropertiesClassification interface for type SQLDatabaseProperties.
func (s *SQLDatabaseProperties) GetSQLDatabaseProperties() *SQLDatabaseProperties { return s }

// SQLDatabaseResource - SqlDatabase portable resource
type SQLDatabaseResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// The resource-specific properties for this resource.
	Properties SQLDatabasePropertiesClassification `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// SQLDatabaseResourceListResult - The response of a SqlDatabaseResource list operation.
type SQLDatabaseResourceListResult struct {
	// REQUIRED; The SqlDatabaseResource items on this page
	Value []*SQLDatabaseResource `json:"value,omitempty"`

	// The link to the next page of items
	NextLink *string `json:"nextLink,omitempty"`
}

// SQLDatabasesClientCreateOrUpdateOptions contains the optional parameters for the SQLDatabasesClient.CreateOrUpdate method.
type SQLDatabasesClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// SQLDatabasesClientDeleteOptions contains the optional parameters for the SQLDatabasesClient.Delete method.
type SQLDatabasesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// SQLDatabasesClientGetOptions contains the optional parameters for the SQLDatabasesClient.Get method.
type SQLDatabasesClientGetOptions struct {
	// placeholder for future optional parameters
}

// SQLDatabasesClientListByRootScopeOptions contains the optional parameters for the SQLDatabasesClient.ListByRootScope method.
type SQLDatabasesClientListByRootScopeOptions struct {
	// placeholder for future optional parameters
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ValuesMongoDatabaseProperties - MongoDatabase Properties for Mode Values
type ValuesMongoDatabaseProperties struct {
	// REQUIRED; Fully qualified resource ID for the environment that the link is linked to
	Environment *string `json:"environment,omitempty"`

	// REQUIRED; Host name of the target Mongo database
	Host *string `json:"host,omitempty"`

	// REQUIRED; Discriminator property for MongoDatabaseProperties.
	Mode *string `json:"mode,omitempty"`

	// REQUIRED; Port value of the target Mongo database
	Port *int32 `json:"port,omitempty"`

	// Fully qualified resource ID for the application that the link is consumed by
	Application *string `json:"application,omitempty"`

	// Secrets values provided for the resource
	Secrets *MongoDatabaseSecrets `json:"secrets,omitempty"`

	// READ-ONLY; Database name of the target Mongo database
	Database *string `json:"database,omitempty" azure:"ro"`

	// READ-ONLY; Provisioning state of the mongo database portable resource at the time the operation was called
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Status of a resource.
	Status *ResourceStatus `json:"status,omitempty" azure:"ro"`
}

// GetMongoDatabaseProperties implements the MongoDatabasePropertiesClassification interface for type ValuesMongoDatabaseProperties.
func (v *ValuesMongoDatabaseProperties) GetMongoDatabaseProperties() *MongoDatabaseProperties {
	return &MongoDatabaseProperties{
		Mode: v.Mode,
		ProvisioningState: v.ProvisioningState,
		Secrets: v.Secrets,
		Status: v.Status,
		Environment: v.Environment,
		Application: v.Application,
	}
}

// ValuesRedisCacheProperties - RedisCache Properties for Mode Values
type ValuesRedisCacheProperties struct {
	// REQUIRED; Fully qualified resource ID for the environment that the link is linked to
	Environment *string `json:"environment,omitempty"`

	// REQUIRED; The host name of the target Redis cache
	Host *string `json:"host,omitempty"`

	// REQUIRED; Discriminator property for RedisCacheProperties.
	Mode *string `json:"mode,omitempty"`

	// REQUIRED; The port value of the target Redis cache
	Port *int32 `json:"port,omitempty"`

	// Fully qualified resource ID for the application that the link is consumed by
	Application *string `json:"application,omitempty"`

	// Secrets provided by resource
	Secrets *RedisCacheSecrets `json:"secrets,omitempty"`

	// READ-ONLY; Provisioning state of the redis cache portable resource at the time the operation was called
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Status of a resource.
	Status *ResourceStatus `json:"status,omitempty" azure:"ro"`

	// READ-ONLY; The username for Redis cache
	Username *string `json:"username,omitempty" azure:"ro"`
}

// GetRedisCacheProperties implements the RedisCachePropertiesClassification interface for type ValuesRedisCacheProperties.
func (v *ValuesRedisCacheProperties) GetRedisCacheProperties() *RedisCacheProperties {
	return &RedisCacheProperties{
		Mode: v.Mode,
		ProvisioningState: v.ProvisioningState,
		Secrets: v.Secrets,
		Status: v.Status,
		Environment: v.Environment,
		Application: v.Application,
	}
}

// ValuesSQLDatabaseProperties - SqlDatabase Properties for Mode Values
type ValuesSQLDatabaseProperties struct {
	// REQUIRED; The name of the Sql database.
	Database *string `json:"database,omitempty"`

	// REQUIRED; Fully qualified resource ID for the environment that the link is linked to
	Environment *string `json:"environment,omitempty"`

	// REQUIRED; Discriminator property for SqlDatabaseProperties.
	Mode *string `json:"mode,omitempty"`

	// REQUIRED; The fully qualified domain name of the Sql database.
	Server *string `json:"server,omitempty"`

	// Fully qualified resource ID for the application that the link is consumed by
	Application *string `json:"application,omitempty"`

	// READ-ONLY; Provisioning state of the Sql database portable resource at the time the operation was called
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Status of a resource.
	Status *ResourceStatus `json:"status,omitempty" azure:"ro"`
}

// GetSQLDatabaseProperties implements the SQLDatabasePropertiesClassification interface for type ValuesSQLDatabaseProperties.
func (v *ValuesSQLDatabaseProperties) GetSQLDatabaseProperties() *SQLDatabaseProperties {
	return &SQLDatabaseProperties{
		Mode: v.Mode,
		ProvisioningState: v.ProvisioningState,
		Status: v.Status,
		Environment: v.Environment,
		Application: v.Application,
	}
}
