/*
Copyright 2023 The Radius Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package applications

import (
	"context"
	"reflect"
	"testing"

	"github.com/radius-project/radius/pkg/cli/clients_new/generated"
)

func Test_isResourceInEnvironment(t *testing.T) {
	id := "/planes/radius/local/resourceGroups/cool-group/providers/Applications.Core/Applications/myapp"

	type args struct {
		ctx             context.Context
		resource        generated.GenericResource
		environmentName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "resource is in environment",
			args: args{
				ctx: context.Background(),
				resource: generated.GenericResource{
					ID: &id,
					Properties: map[string]interface{}{
						"environment": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/applications.core/environments/env0",
					},
				},
				environmentName: "env0",
			},
			want: true,
		},
		{
			name: "resource is not in environment",
			args: args{
				ctx: context.Background(),
				resource: generated.GenericResource{
					ID: &id,
					Properties: map[string]interface{}{
						"environment": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/applications.core/environments/env0",
					},
				},
				environmentName: "env",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isResourceInEnvironment(tt.args.ctx, tt.args.resource, tt.args.environmentName); got != tt.want {
				t.Errorf("isResourceInEnvironment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isResourceInApplication(t *testing.T) {
	id := "/planes/radius/local/resourceGroups/cool-group/providers/Applications.Core/Applications/myapp"

	type args struct {
		ctx             context.Context
		resource        generated.GenericResource
		applicationName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "resource is in application",
			args: args{
				ctx: context.Background(),
				resource: generated.GenericResource{
					ID: &id,
					Properties: map[string]interface{}{
						"application": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/applications.core/applications/myapp",
					},
				},
				applicationName: "myapp",
			},
			want: true,
		},
		{
			name: "resource is not in application",
			args: args{
				ctx: context.Background(),
				resource: generated.GenericResource{
					ID: &id,
					Properties: map[string]interface{}{
						"application": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/applications.core/applications/myapp",
					},
				},
				applicationName: "myapp2",
			},
			want: false,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if got := isResourceInApplication(tt.args.ctx, tt.args.resource, tt.args.applicationName); got != tt.want {
				t.Errorf("isResourceInApplication() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compute(t *testing.T) {

	sqlRteID := "/planes/radius/local/resourcegroups/default/providers/Applications.Core/httpRoutes/sql-rte"
	sqlRteType := "Applications.Core/httpRoutes"
	sqlRteName := "sql-rte"

	sqlAppCntrID := "/planes/radius/local/resourcegroups/default/providers/Applications.Core/containers/sql-app-ctnr"
	sqlAppCntrName := "sql-app-ctnr"
	sqlAppCntrType := "Applications.Core/containers"

	sqlCntrID := "/planes/radius/local/resourcegroups/default/providers/Applications.Core/containers/sql-ctnr"
	sqlCntrName := "sql-ctnr"
	sqlCntrType := "Applications.Core/containers"

	sqlDbID := "/planes/radius/local/resourcegroups/default/providers/Applications.Datastores/sqlDatabases/sql-db"

	type args struct {
		applicationName      string
		applicationResources []generated.GenericResource
		environmentResources []generated.GenericResource
	}
	tests := []struct {
		name string
		args args
		want *ApplicationGraphResponse
	}{
		{
			name: "compute graph",
			args: args{
				applicationName: "myapp",
				applicationResources: []generated.GenericResource{
					{
						ID: &sqlRteID,
						Properties: map[string]interface{}{
							"application":       "/planes/radius/local/resourcegroups/default/providers/Applications.Core/Applications/myapp",
							"provisioningState": "Succeeded",
						},
						Name: &sqlRteName,
						Type: &sqlRteType,
					},
					{
						ID: &sqlAppCntrID,
						Properties: map[string]interface{}{
							"connections": map[string]interface{}{
								"sql": map[string]interface{}{
									"source": "/planes/radius/local/resourcegroups/default/providers/Applications.Datastores/sqlDatabases/sql-db",
								},
							},
							"application":       "/planes/radius/local/resourcegroups/default/providers/Applications.Core/Applications/myapp",
							"provisioningState": "Succeeded",
							"status": map[string]interface{}{
								"outputResources": map[string]interface{}{
									"localId": "something",
									"id":      "/some/thing/else",
								},
							},
						},
						Name: &sqlAppCntrName,
						Type: &sqlAppCntrType,
					},
					{
						ID: &sqlCntrID,
						Properties: map[string]interface{}{
							"container": map[string]interface{}{
								"ports": []interface{}{
									map[string]interface{}{
										"port":     8080,
										"protocol": "TCP",
										"provides": "/planes/radius/local/resourcegroups/default/providers/Applications.Core/httpRoutes/sql-rte",
									},
								},
							},
							"application":       "/planes/radius/local/resourcegroups/default/providers/Applications.Core/Applications/myapp",
							"provisioningState": "Succeeded",
							"status": map[string]interface{}{
								"outputResources": map[string]interface{}{
									"localId": "something",
									"id":      "/some/thing/else",
								},
							},
						},
						Name: &sqlCntrName,
						Type: &sqlCntrType,
					},
				},
				environmentResources: []generated.GenericResource{},
			},
			want: &ApplicationGraphResponse{
				Resources: []*ApplicationGraphResource{
					{
						ID:                sqlRteID,
						Name:              sqlRteName,
						Type:              sqlRteType,
						ProvisioningState: "Succeeded",
						Resources:         []ApplicationGraphOutputResource{},
						Connections: []ApplicationGraphConnection{
							{
								ID:        sqlCntrID,
								Direction: "Inbound",
							},
						},
					},
					{
						ID:                sqlAppCntrID,
						Name:              sqlAppCntrName,
						Type:              sqlAppCntrType,
						ProvisioningState: "Succeeded",
						Resources: []ApplicationGraphOutputResource{
							{},
						},
						Connections: []ApplicationGraphConnection{
							{
								Direction: "Inbound",
								ID:        sqlDbID,
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeGraph(tt.args.applicationName, tt.args.applicationResources, tt.args.environmentResources); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compute() = %v, want %v", got, tt.want)
			}
		})
	}
}
