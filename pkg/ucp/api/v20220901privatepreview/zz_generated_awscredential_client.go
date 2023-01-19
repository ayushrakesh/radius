//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package v20220901privatepreview

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AWSCredentialClient contains the methods for the AWSCredential group.
// Don't use this type directly, use NewAWSCredentialClient() instead.
type AWSCredentialClient struct {
	host string
	pl runtime.Pipeline
}

// NewAWSCredentialClient creates a new instance of AWSCredentialClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAWSCredentialClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*AWSCredentialClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &AWSCredentialClient{
		host: ep,
pl: pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a Credential.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-privatepreview
// planeType - The type of the plane
// planeName - The name of the plane
// credentialName - The name of the credential
// credential - Credential details
// options - AWSCredentialClientCreateOrUpdateOptions contains the optional parameters for the AWSCredentialClient.CreateOrUpdate
// method.
func (client *AWSCredentialClient) CreateOrUpdate(ctx context.Context, planeType string, planeName string, credentialName string, credential CredentialResource, options *AWSCredentialClientCreateOrUpdateOptions) (AWSCredentialClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, planeType, planeName, credentialName, credential, options)
	if err != nil {
		return AWSCredentialClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AWSCredentialClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return AWSCredentialClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AWSCredentialClient) createOrUpdateCreateRequest(ctx context.Context, planeType string, planeName string, credentialName string, credential CredentialResource, options *AWSCredentialClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/planes/{planeType}/{planeName}/providers/System.AWS/credentials/{credentialName}"
	if planeType == "" {
		return nil, errors.New("parameter planeType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planeType}", url.PathEscape(planeType))
	if planeName == "" {
		return nil, errors.New("parameter planeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planeName}", url.PathEscape(planeName))
	if credentialName == "" {
		return nil, errors.New("parameter credentialName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{credentialName}", url.PathEscape(credentialName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, credential)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AWSCredentialClient) createOrUpdateHandleResponse(resp *http.Response) (AWSCredentialClientCreateOrUpdateResponse, error) {
	result := AWSCredentialClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CredentialResource); err != nil {
		return AWSCredentialClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete credential resource for this plane instance
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-privatepreview
// planeType - The type of the plane
// planeName - The name of the plane
// credentialName - The name of the credential
// options - AWSCredentialClientDeleteOptions contains the optional parameters for the AWSCredentialClient.Delete method.
func (client *AWSCredentialClient) Delete(ctx context.Context, planeType string, planeName string, credentialName string, options *AWSCredentialClientDeleteOptions) (AWSCredentialClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, planeType, planeName, credentialName, options)
	if err != nil {
		return AWSCredentialClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AWSCredentialClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return AWSCredentialClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return AWSCredentialClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AWSCredentialClient) deleteCreateRequest(ctx context.Context, planeType string, planeName string, credentialName string, options *AWSCredentialClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/planes/{planeType}/{planeName}/providers/System.AWS/credentials/{credentialName}"
	if planeType == "" {
		return nil, errors.New("parameter planeType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planeType}", url.PathEscape(planeType))
	if planeName == "" {
		return nil, errors.New("parameter planeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planeName}", url.PathEscape(planeName))
	if credentialName == "" {
		return nil, errors.New("parameter credentialName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{credentialName}", url.PathEscape(credentialName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get name of the secret that is holding credentials for the plane instance
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-privatepreview
// planeType - The type of the plane
// planeName - The name of the plane
// credentialName - The name of the credential
// options - AWSCredentialClientGetOptions contains the optional parameters for the AWSCredentialClient.Get method.
func (client *AWSCredentialClient) Get(ctx context.Context, planeType string, planeName string, credentialName string, options *AWSCredentialClientGetOptions) (AWSCredentialClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, planeType, planeName, credentialName, options)
	if err != nil {
		return AWSCredentialClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AWSCredentialClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AWSCredentialClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AWSCredentialClient) getCreateRequest(ctx context.Context, planeType string, planeName string, credentialName string, options *AWSCredentialClientGetOptions) (*policy.Request, error) {
	urlPath := "/planes/{planeType}/{planeName}/providers/System.AWS/credentials/{credentialName}"
	if planeType == "" {
		return nil, errors.New("parameter planeType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planeType}", url.PathEscape(planeType))
	if planeName == "" {
		return nil, errors.New("parameter planeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planeName}", url.PathEscape(planeName))
	if credentialName == "" {
		return nil, errors.New("parameter credentialName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{credentialName}", url.PathEscape(credentialName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AWSCredentialClient) getHandleResponse(resp *http.Response) (AWSCredentialClientGetResponse, error) {
	result := AWSCredentialClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CredentialResource); err != nil {
		return AWSCredentialClientGetResponse{}, err
	}
	return result, nil
}

// List - List the credentials for this plane instance
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-privatepreview
// planeType - The type of the plane
// planeName - The name of the plane
// options - AWSCredentialClientListOptions contains the optional parameters for the AWSCredentialClient.List method.
func (client *AWSCredentialClient) List(ctx context.Context, planeType string, planeName string, options *AWSCredentialClientListOptions) (AWSCredentialClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, planeType, planeName, options)
	if err != nil {
		return AWSCredentialClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AWSCredentialClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AWSCredentialClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *AWSCredentialClient) listCreateRequest(ctx context.Context, planeType string, planeName string, options *AWSCredentialClientListOptions) (*policy.Request, error) {
	urlPath := "/planes/{planeType}/{planeName}/providers/System.AWS/credentials"
	if planeType == "" {
		return nil, errors.New("parameter planeType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planeType}", url.PathEscape(planeType))
	if planeName == "" {
		return nil, errors.New("parameter planeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planeName}", url.PathEscape(planeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AWSCredentialClient) listHandleResponse(resp *http.Response) (AWSCredentialClientListResponse, error) {
	result := AWSCredentialClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CredentialResourceList); err != nil {
		return AWSCredentialClientListResponse{}, err
	}
	return result, nil
}
