//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package radclient

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AzureConnectionClient contains the methods for the AzureConnection group.
// Don't use this type directly, use NewAzureConnectionClient() instead.
type AzureConnectionClient struct {
	ep string
	pl runtime.Pipeline
	subscriptionID string
}

// NewAzureConnectionClient creates a new instance of AzureConnectionClient with the specified values.
func NewAzureConnectionClient(con *arm.Connection, subscriptionID string) *AzureConnectionClient {
	return &AzureConnectionClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a AzureConnection resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *AzureConnectionClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, applicationName string, azureConnectionName string, parameters AzureConnectionResource, options *AzureConnectionBeginCreateOrUpdateOptions) (AzureConnectionCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, applicationName, azureConnectionName, parameters, options)
	if err != nil {
		return AzureConnectionCreateOrUpdatePollerResponse{}, err
	}
	result := AzureConnectionCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("AzureConnectionClient.CreateOrUpdate", "location", resp, 	client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return AzureConnectionCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &AzureConnectionCreateOrUpdatePoller {
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a AzureConnection resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *AzureConnectionClient) createOrUpdate(ctx context.Context, resourceGroupName string, applicationName string, azureConnectionName string, parameters AzureConnectionResource, options *AzureConnectionBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, applicationName, azureConnectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	 return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AzureConnectionClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, applicationName string, azureConnectionName string, parameters AzureConnectionResource, options *AzureConnectionBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/radiusv3/Application/{applicationName}/AzureConnection/{azureConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if azureConnectionName == "" {
		return nil, errors.New("parameter azureConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureConnectionName}", url.PathEscape(azureConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *AzureConnectionClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes a AzureConnection resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *AzureConnectionClient) BeginDelete(ctx context.Context, resourceGroupName string, applicationName string, azureConnectionName string, options *AzureConnectionBeginDeleteOptions) (AzureConnectionDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, applicationName, azureConnectionName, options)
	if err != nil {
		return AzureConnectionDeletePollerResponse{}, err
	}
	result := AzureConnectionDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("AzureConnectionClient.Delete", "location", resp, 	client.pl, client.deleteHandleError)
	if err != nil {
		return AzureConnectionDeletePollerResponse{}, err
	}
	result.Poller = &AzureConnectionDeletePoller {
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a AzureConnection resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *AzureConnectionClient) deleteOperation(ctx context.Context, resourceGroupName string, applicationName string, azureConnectionName string, options *AzureConnectionBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, applicationName, azureConnectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	 return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AzureConnectionClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, applicationName string, azureConnectionName string, options *AzureConnectionBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/radiusv3/Application/{applicationName}/AzureConnection/{azureConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if azureConnectionName == "" {
		return nil, errors.New("parameter azureConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureConnectionName}", url.PathEscape(azureConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *AzureConnectionClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets a AzureConnection resource by name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *AzureConnectionClient) Get(ctx context.Context, resourceGroupName string, applicationName string, azureConnectionName string, options *AzureConnectionGetOptions) (AzureConnectionGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, applicationName, azureConnectionName, options)
	if err != nil {
		return AzureConnectionGetResponse{}, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return AzureConnectionGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AzureConnectionGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AzureConnectionClient) getCreateRequest(ctx context.Context, resourceGroupName string, applicationName string, azureConnectionName string, options *AzureConnectionGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/radiusv3/Application/{applicationName}/AzureConnection/{azureConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if azureConnectionName == "" {
		return nil, errors.New("parameter azureConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureConnectionName}", url.PathEscape(azureConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AzureConnectionClient) getHandleResponse(resp *http.Response) (AzureConnectionGetResponse, error) {
	result := AzureConnectionGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureConnectionResource); err != nil {
		return AzureConnectionGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *AzureConnectionClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - List the AzureConnection resources deployed in the application.
// If the operation fails it returns the *ErrorResponse error type.
func (client *AzureConnectionClient) List(ctx context.Context, resourceGroupName string, applicationName string, options *AzureConnectionListOptions) (AzureConnectionListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, applicationName, options)
	if err != nil {
		return AzureConnectionListResponse{}, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return AzureConnectionListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AzureConnectionListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *AzureConnectionClient) listCreateRequest(ctx context.Context, resourceGroupName string, applicationName string, options *AzureConnectionListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/radiusv3/Application/{applicationName}/AzureConnection"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AzureConnectionClient) listHandleResponse(resp *http.Response) (AzureConnectionListResponse, error) {
	result := AzureConnectionListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureConnectionList); err != nil {
		return AzureConnectionListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *AzureConnectionClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

