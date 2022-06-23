//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package v20220315privatepreview

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RedisCachesClient contains the methods for the RedisCaches group.
// Don't use this type directly, use NewRedisCachesClient() instead.
type RedisCachesClient struct {
	ep string
	pl runtime.Pipeline
	rootScope string
}

// NewRedisCachesClient creates a new instance of RedisCachesClient with the specified values.
func NewRedisCachesClient(con *arm.Connection, rootScope string) *RedisCachesClient {
	return &RedisCachesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), rootScope: rootScope}
}

// CreateOrUpdate - Creates or updates a RedisCache resource
// If the operation fails it returns the *ErrorResponse error type.
func (client *RedisCachesClient) CreateOrUpdate(ctx context.Context, redisCacheName string, redisCacheParameters RedisCacheResource, options *RedisCachesCreateOrUpdateOptions) (RedisCachesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, redisCacheName, redisCacheParameters, options)
	if err != nil {
		return RedisCachesCreateOrUpdateResponse{}, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return RedisCachesCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return RedisCachesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RedisCachesClient) createOrUpdateCreateRequest(ctx context.Context, redisCacheName string, redisCacheParameters RedisCacheResource, options *RedisCachesCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Connector/redisCaches/{redisCacheName}"
	if client.rootScope == "" {
		return nil, errors.New("parameter client.rootScope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", url.PathEscape(client.rootScope))
	if redisCacheName == "" {
		return nil, errors.New("parameter redisCacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{redisCacheName}", url.PathEscape(redisCacheName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, redisCacheParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *RedisCachesClient) createOrUpdateHandleResponse(resp *http.Response) (RedisCachesCreateOrUpdateResponse, error) {
	result := RedisCachesCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RedisCacheResource); err != nil {
		return RedisCachesCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *RedisCachesClient) createOrUpdateHandleError(resp *http.Response) error {
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

// Delete - Deletes an existing redisCache resource
// If the operation fails it returns the *ErrorResponse error type.
func (client *RedisCachesClient) Delete(ctx context.Context, redisCacheName string, options *RedisCachesDeleteOptions) (RedisCachesDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, redisCacheName, options)
	if err != nil {
		return RedisCachesDeleteResponse{}, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return RedisCachesDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return RedisCachesDeleteResponse{}, client.deleteHandleError(resp)
	}
	return RedisCachesDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RedisCachesClient) deleteCreateRequest(ctx context.Context, redisCacheName string, options *RedisCachesDeleteOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Connector/redisCaches/{redisCacheName}"
	if client.rootScope == "" {
		return nil, errors.New("parameter client.rootScope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", url.PathEscape(client.rootScope))
	if redisCacheName == "" {
		return nil, errors.New("parameter redisCacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{redisCacheName}", url.PathEscape(redisCacheName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *RedisCachesClient) deleteHandleError(resp *http.Response) error {
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

// Get - Retrieves information about a redisCache resource
// If the operation fails it returns the *ErrorResponse error type.
func (client *RedisCachesClient) Get(ctx context.Context, redisCacheName string, options *RedisCachesGetOptions) (RedisCachesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, redisCacheName, options)
	if err != nil {
		return RedisCachesGetResponse{}, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return RedisCachesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RedisCachesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RedisCachesClient) getCreateRequest(ctx context.Context, redisCacheName string, options *RedisCachesGetOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Connector/redisCaches/{redisCacheName}"
	if client.rootScope == "" {
		return nil, errors.New("parameter client.rootScope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", url.PathEscape(client.rootScope))
	if redisCacheName == "" {
		return nil, errors.New("parameter redisCacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{redisCacheName}", url.PathEscape(redisCacheName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RedisCachesClient) getHandleResponse(resp *http.Response) (RedisCachesGetResponse, error) {
	result := RedisCachesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RedisCacheResource); err != nil {
		return RedisCachesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *RedisCachesClient) getHandleError(resp *http.Response) error {
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

// ListByRootScope - Lists information about all redisCache resources in the given root scope
// If the operation fails it returns the *ErrorResponse error type.
func (client *RedisCachesClient) ListByRootScope(options *RedisCachesListByRootScopeOptions) (*RedisCachesListByRootScopePager) {
	return &RedisCachesListByRootScopePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByRootScopeCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp RedisCachesListByRootScopeResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RedisCacheList.NextLink)
		},
	}
}

// listByRootScopeCreateRequest creates the ListByRootScope request.
func (client *RedisCachesClient) listByRootScopeCreateRequest(ctx context.Context, options *RedisCachesListByRootScopeOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Connector/redisCaches"
	if client.rootScope == "" {
		return nil, errors.New("parameter client.rootScope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", url.PathEscape(client.rootScope))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByRootScopeHandleResponse handles the ListByRootScope response.
func (client *RedisCachesClient) listByRootScopeHandleResponse(resp *http.Response) (RedisCachesListByRootScopeResponse, error) {
	result := RedisCachesListByRootScopeResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RedisCacheList); err != nil {
		return RedisCachesListByRootScopeResponse{}, err
	}
	return result, nil
}

// listByRootScopeHandleError handles the ListByRootScope error response.
func (client *RedisCachesClient) listByRootScopeHandleError(resp *http.Response) error {
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

// ListSecrets - Lists secrets values for the specified RedisCache resource
// If the operation fails it returns the *ErrorResponse error type.
func (client *RedisCachesClient) ListSecrets(ctx context.Context, redisCacheName string, options *RedisCachesListSecretsOptions) (RedisCachesListSecretsResponse, error) {
	req, err := client.listSecretsCreateRequest(ctx, redisCacheName, options)
	if err != nil {
		return RedisCachesListSecretsResponse{}, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return RedisCachesListSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RedisCachesListSecretsResponse{}, client.listSecretsHandleError(resp)
	}
	return client.listSecretsHandleResponse(resp)
}

// listSecretsCreateRequest creates the ListSecrets request.
func (client *RedisCachesClient) listSecretsCreateRequest(ctx context.Context, redisCacheName string, options *RedisCachesListSecretsOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Connector/redisCaches/{redisCacheName}/listSecrets"
	if client.rootScope == "" {
		return nil, errors.New("parameter client.rootScope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", url.PathEscape(client.rootScope))
	if redisCacheName == "" {
		return nil, errors.New("parameter redisCacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{redisCacheName}", url.PathEscape(redisCacheName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listSecretsHandleResponse handles the ListSecrets response.
func (client *RedisCachesClient) listSecretsHandleResponse(resp *http.Response) (RedisCachesListSecretsResponse, error) {
	result := RedisCachesListSecretsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RedisCacheSecrets); err != nil {
		return RedisCachesListSecretsResponse{}, err
	}
	return result, nil
}

// listSecretsHandleError handles the ListSecrets error response.
func (client *RedisCachesClient) listSecretsHandleError(resp *http.Response) error {
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

