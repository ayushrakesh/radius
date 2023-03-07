//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package generated

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

// HTTPRoutesClient contains the methods for the HTTPRoutes group.
// Don't use this type directly, use NewHTTPRoutesClient() instead.
type HTTPRoutesClient struct {
	host      string
	rootScope string
	pl        runtime.Pipeline
}

// NewHTTPRoutesClient creates a new instance of HTTPRoutesClient with the specified values.
// rootScope - The scope in which the resource is present. For Azure resource this would be /subscriptions/{subscriptionID}/resourceGroup/{resourcegroupID}
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewHTTPRoutesClient(rootScope string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HTTPRoutesClient, error) {
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
	client := &HTTPRoutesClient{
		rootScope: rootScope,
		host:      ep,
		pl:        pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update an HTTP Route.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-15-privatepreview
// httpRouteName - The name of the HTTP Route.
// httpRouteResource - HTTP Route details
// options - HTTPRoutesClientCreateOrUpdateOptions contains the optional parameters for the HTTPRoutesClient.CreateOrUpdate
// method.
func (client *HTTPRoutesClient) CreateOrUpdate(ctx context.Context, httpRouteName string, httpRouteResource HTTPRouteResource, options *HTTPRoutesClientCreateOrUpdateOptions) (HTTPRoutesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, httpRouteName, httpRouteResource, options)
	if err != nil {
		return HTTPRoutesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRoutesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return HTTPRoutesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *HTTPRoutesClient) createOrUpdateCreateRequest(ctx context.Context, httpRouteName string, httpRouteResource HTTPRouteResource, options *HTTPRoutesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/httpRoutes/{httpRouteName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if httpRouteName == "" {
		return nil, errors.New("parameter httpRouteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{httpRouteName}", url.PathEscape(httpRouteName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, httpRouteResource)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *HTTPRoutesClient) createOrUpdateHandleResponse(resp *http.Response) (HTTPRoutesClientCreateOrUpdateResponse, error) {
	result := HTTPRoutesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HTTPRouteResource); err != nil {
		return HTTPRoutesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete an HTTP Route.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-15-privatepreview
// httpRouteName - The name of the HTTP Route.
// options - HTTPRoutesClientDeleteOptions contains the optional parameters for the HTTPRoutesClient.Delete method.
func (client *HTTPRoutesClient) Delete(ctx context.Context, httpRouteName string, options *HTTPRoutesClientDeleteOptions) (HTTPRoutesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, httpRouteName, options)
	if err != nil {
		return HTTPRoutesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRoutesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return HTTPRoutesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return HTTPRoutesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *HTTPRoutesClient) deleteCreateRequest(ctx context.Context, httpRouteName string, options *HTTPRoutesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/httpRoutes/{httpRouteName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if httpRouteName == "" {
		return nil, errors.New("parameter httpRouteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{httpRouteName}", url.PathEscape(httpRouteName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the properties of an HTTP Route.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-15-privatepreview
// httpRouteName - The name of the HTTP Route.
// options - HTTPRoutesClientGetOptions contains the optional parameters for the HTTPRoutesClient.Get method.
func (client *HTTPRoutesClient) Get(ctx context.Context, httpRouteName string, options *HTTPRoutesClientGetOptions) (HTTPRoutesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, httpRouteName, options)
	if err != nil {
		return HTTPRoutesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRoutesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPRoutesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *HTTPRoutesClient) getCreateRequest(ctx context.Context, httpRouteName string, options *HTTPRoutesClientGetOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/httpRoutes/{httpRouteName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if httpRouteName == "" {
		return nil, errors.New("parameter httpRouteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{httpRouteName}", url.PathEscape(httpRouteName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HTTPRoutesClient) getHandleResponse(resp *http.Response) (HTTPRoutesClientGetResponse, error) {
	result := HTTPRoutesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HTTPRouteResource); err != nil {
		return HTTPRoutesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByScopePager - List all HTTP Routes in the given scope.
// Generated from API version 2022-03-15-privatepreview
// options - HTTPRoutesClientListByScopeOptions contains the optional parameters for the HTTPRoutesClient.ListByScope method.
func (client *HTTPRoutesClient) NewListByScopePager(options *HTTPRoutesClientListByScopeOptions) *runtime.Pager[HTTPRoutesClientListByScopeResponse] {
	return runtime.NewPager(runtime.PagingHandler[HTTPRoutesClientListByScopeResponse]{
		More: func(page HTTPRoutesClientListByScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HTTPRoutesClientListByScopeResponse) (HTTPRoutesClientListByScopeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByScopeCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return HTTPRoutesClientListByScopeResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return HTTPRoutesClientListByScopeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return HTTPRoutesClientListByScopeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByScopeHandleResponse(resp)
		},
	})
}

// listByScopeCreateRequest creates the ListByScope request.
func (client *HTTPRoutesClient) listByScopeCreateRequest(ctx context.Context, options *HTTPRoutesClientListByScopeOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/httpRoutes"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByScopeHandleResponse handles the ListByScope response.
func (client *HTTPRoutesClient) listByScopeHandleResponse(resp *http.Response) (HTTPRoutesClientListByScopeResponse, error) {
	result := HTTPRoutesClientListByScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HTTPRouteResourceList); err != nil {
		return HTTPRoutesClientListByScopeResponse{}, err
	}
	return result, nil
}

// Update - Update the properties of an existing HTTP Route.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-15-privatepreview
// httpRouteName - The name of the HTTP Route.
// httpRouteResource - HTTP Route details
// options - HTTPRoutesClientUpdateOptions contains the optional parameters for the HTTPRoutesClient.Update method.
func (client *HTTPRoutesClient) Update(ctx context.Context, httpRouteName string, httpRouteResource HTTPRouteResource, options *HTTPRoutesClientUpdateOptions) (HTTPRoutesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, httpRouteName, httpRouteResource, options)
	if err != nil {
		return HTTPRoutesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRoutesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return HTTPRoutesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *HTTPRoutesClient) updateCreateRequest(ctx context.Context, httpRouteName string, httpRouteResource HTTPRouteResource, options *HTTPRoutesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/httpRoutes/{httpRouteName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if httpRouteName == "" {
		return nil, errors.New("parameter httpRouteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{httpRouteName}", url.PathEscape(httpRouteName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, httpRouteResource)
}

// updateHandleResponse handles the Update response.
func (client *HTTPRoutesClient) updateHandleResponse(resp *http.Response) (HTTPRoutesClientUpdateResponse, error) {
	result := HTTPRoutesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HTTPRouteResource); err != nil {
		return HTTPRoutesClientUpdateResponse{}, err
	}
	return result, nil
}