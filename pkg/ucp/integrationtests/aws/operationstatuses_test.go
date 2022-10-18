// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package aws

// Tests that test with Mock RP functionality and UCP Server

import (
	"context"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GetOperationStatuses(t *testing.T) {
	ucp, ucpClient, cloudcontrolClient, _ := initializeTest(t)

	cloudcontrolClient.EXPECT().GetResourceRequestStatus(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, params *cloudcontrol.GetResourceRequestStatusInput, optFns ...func(*cloudcontrol.Options)) (*cloudcontrol.GetResourceRequestStatusOutput, error) {
		output := cloudcontrol.GetResourceRequestStatusOutput{
			ProgressEvent: &types.ProgressEvent{
				RequestToken: to.StringPtr(testAWSRequestToken),
				EventTime:    &time.Time{},
			},
		}
		return &output, nil
	})

	operationResultsRequest, err := http.NewRequest(http.MethodGet, ucp.URL+basePath+testProxyRequestAWSAsyncPath+"/operationStatuses/"+strings.ToLower(testAWSRequestToken), nil)
	require.NoError(t, err)
	operationResultsResponse, err := ucpClient.httpClient.Do(operationResultsRequest)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, operationResultsResponse.StatusCode)
}
