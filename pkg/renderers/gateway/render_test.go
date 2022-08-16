// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/go-logr/logr"
	"github.com/project-radius/radius/pkg/azure/azresources"
	"github.com/project-radius/radius/pkg/azure/radclient"
	"github.com/project-radius/radius/pkg/kubernetes"
	"github.com/project-radius/radius/pkg/radlogger"
	"github.com/project-radius/radius/pkg/radrp/outputresource"
	"github.com/project-radius/radius/pkg/renderers"
	"github.com/project-radius/radius/pkg/renderers/httproutev1alpha3"
	"github.com/project-radius/radius/pkg/resourcekinds"
	contourv1 "github.com/projectcontour/contour/apis/projectcontour/v1"
	"github.com/stretchr/testify/require"
)

const (
	subscriptionID  = "default"
	resourceGroup   = "default"
	applicationName = "test-application"
	resourceName    = "test-gateway"
	publicIP        = "86.753.099.99"
)

func createContext(t *testing.T) context.Context {
	logger, err := radlogger.NewTestLogger(t)
	if err != nil {
		t.Log("Unable to initialize logger")
		return context.Background()
	}
	return logr.NewContext(context.Background(), logger)
}

func Test_GetDependencyIDs_Success(t *testing.T) {
	testRouteAResourceID := makeRouteResourceID("testroutea")
	testRouteBResourceID := makeRouteResourceID("testrouteb")
	properties := radclient.GatewayProperties{
		Routes: []*radclient.GatewayRoute{
			{
				Destination: &testRouteAResourceID,
			},
			{
				Destination: &testRouteBResourceID,
			},
		},
	}
	resource := makeResource(t, properties)

	renderer := Renderer{}
	radiusResourceIDs, azureResourceIDs, err := renderer.GetDependencyIDs(createContext(t), resource)
	require.NoError(t, err)
	require.Len(t, radiusResourceIDs, 2)
	require.Len(t, azureResourceIDs, 0)

	expectedRadiusResourceIDs := []azresources.ResourceID{
		makeResourceID(t, testRouteAResourceID),
		makeResourceID(t, testRouteBResourceID),
	}
	require.ElementsMatch(t, expectedRadiusResourceIDs, radiusResourceIDs)

	expectedAzureResourceIDs := []azresources.ResourceID{}
	require.ElementsMatch(t, expectedAzureResourceIDs, azureResourceIDs)
}

func Test_Render_WithNoHostname(t *testing.T) {
	r := &Renderer{}

	properties, expectedIncludes := makeTestGateway(radclient.GatewayProperties{})
	resource := makeResource(t, properties)
	dependencies := map[string]renderers.RendererDependency{}
	additionalProperties := GetRuntimeOptions()

	output, err := r.Render(context.Background(), renderers.RenderOptions{Resource: resource, Dependencies: dependencies, Runtime: additionalProperties})
	require.NoError(t, err)
	require.Len(t, output.Resources, 2)
	require.Empty(t, output.SecretValues)

	expectedHostname := fmt.Sprintf("%s.%s.%s.nip.io", resourceName, applicationName, publicIP)
	expectedURL := "http://" + expectedHostname
	require.Equal(t, expectedURL, output.ComputedValues["url"].Value)

	validateGateway(t, output.Resources, expectedHostname, expectedIncludes)
}

func Test_Render_WithPrefix(t *testing.T) {
	r := &Renderer{}

	prefix := "prefix"
	properties, expectedIncludes := makeTestGateway(radclient.GatewayProperties{
		Hostname: &radclient.GatewayPropertiesHostname{
			Prefix: &prefix,
		},
	})
	resource := makeResource(t, properties)
	dependencies := map[string]renderers.RendererDependency{}
	additionalProperties := GetRuntimeOptions()

	output, err := r.Render(context.Background(), renderers.RenderOptions{Resource: resource, Dependencies: dependencies, Runtime: additionalProperties})
	require.NoError(t, err)
	require.Len(t, output.Resources, 2)
	require.Empty(t, output.SecretValues)

	expectedHostname := fmt.Sprintf("%s.%s.%s.nip.io", prefix, applicationName, publicIP)
	expectedURL := "http://" + expectedHostname
	require.Equal(t, expectedURL, output.ComputedValues["url"].Value)

	validateGateway(t, output.Resources, expectedHostname, expectedIncludes)
}

func Test_Render_WithFQHostname(t *testing.T) {
	r := &Renderer{}

	expectedHostname := "test-fqdn.contoso.com"
	expectedURL := "http://" + expectedHostname
	properties, expectedIncludes := makeTestGateway(radclient.GatewayProperties{
		Hostname: &radclient.GatewayPropertiesHostname{
			FullyQualifiedHostname: &expectedHostname,
		},
	})
	resource := makeResource(t, properties)
	dependencies := map[string]renderers.RendererDependency{}
	additionalProperties := GetRuntimeOptions()

	output, err := r.Render(context.Background(), renderers.RenderOptions{Resource: resource, Dependencies: dependencies, Runtime: additionalProperties})
	require.NoError(t, err)
	require.Len(t, output.Resources, 2)
	require.Empty(t, output.SecretValues)
	require.Equal(t, expectedURL, output.ComputedValues["url"].Value)

	validateGateway(t, output.Resources, expectedHostname, expectedIncludes)
}

func Test_Render_WithFQHostname_OverridesPrefix(t *testing.T) {
	r := &Renderer{}

	expectedHostname := "http://test-fqdn.contoso.com"
	expectedURL := "http://" + expectedHostname
	prefix := "test-prefix"
	properties, expectedIncludes := makeTestGateway(radclient.GatewayProperties{
		Hostname: &radclient.GatewayPropertiesHostname{
			Prefix:                 &prefix,
			FullyQualifiedHostname: &expectedHostname,
		},
	})
	resource := makeResource(t, properties)
	dependencies := map[string]renderers.RendererDependency{}
	additionalProperties := GetRuntimeOptions()

	output, err := r.Render(context.Background(), renderers.RenderOptions{Resource: resource, Dependencies: dependencies, Runtime: additionalProperties})
	require.NoError(t, err)
	require.Len(t, output.Resources, 2)
	require.Empty(t, output.SecretValues)
	require.Equal(t, expectedURL, output.ComputedValues["url"].Value)

	validateGateway(t, output.Resources, expectedHostname, expectedIncludes)
}

func Test_Render_DevEnvironment(t *testing.T) {
	r := &Renderer{}

	publicIP := "http://localhost:32323"
	expectedFqdn := "localhost"
	properties, expectedIncludes := makeTestGateway(radclient.GatewayProperties{})
	resource := makeResource(t, properties)
	dependencies := map[string]renderers.RendererDependency{}
	additionalProperties := renderers.RuntimeOptions{
		Gateway: renderers.GatewayOptions{
			PublicEndpointOverride: true,
			PublicIP:               publicIP,
		},
	}

	output, err := r.Render(context.Background(), renderers.RenderOptions{Resource: resource, Dependencies: dependencies, Runtime: additionalProperties})
	require.NoError(t, err)
	require.Len(t, output.Resources, 2)
	require.Empty(t, output.SecretValues)
	require.Equal(t, publicIP, output.ComputedValues["url"].Value)

	validateGateway(t, output.Resources, expectedFqdn, expectedIncludes)
}

func Test_Render_PublicEndpointOverride(t *testing.T) {
	r := &Renderer{}

	publicIP := "http://www.contoso.com:32323"
	expectedFqdn := "www.contoso.com"
	properties, expectedIncludes := makeTestGateway(radclient.GatewayProperties{})
	resource := makeResource(t, properties)
	dependencies := map[string]renderers.RendererDependency{}
	additionalProperties := renderers.RuntimeOptions{
		Gateway: renderers.GatewayOptions{
			PublicEndpointOverride: true,
			PublicIP:               publicIP,
		},
	}

	output, err := r.Render(context.Background(), renderers.RenderOptions{Resource: resource, Dependencies: dependencies, Runtime: additionalProperties})
	require.NoError(t, err)
	require.Len(t, output.Resources, 2)
	require.Empty(t, output.SecretValues)
	require.Equal(t, publicIP, output.ComputedValues["url"].Value)

	validateGateway(t, output.Resources, expectedFqdn, expectedIncludes)
}

func Test_Render_WithMissingPublicIP(t *testing.T) {
	r := &Renderer{}

	properties, expectedIncludes := makeTestGateway(radclient.GatewayProperties{})
	resource := makeResource(t, properties)
	dependencies := map[string]renderers.RendererDependency{}
	additionalProperties := renderers.RuntimeOptions{
		Gateway: renderers.GatewayOptions{
			PublicEndpointOverride: false,
			PublicIP:               "",
		},
	}

	output, err := r.Render(context.Background(), renderers.RenderOptions{Resource: resource, Dependencies: dependencies, Runtime: additionalProperties})
	require.NoError(t, err)
	require.Len(t, output.Resources, 2)
	require.Empty(t, output.SecretValues)
	require.Equal(t, "unknown", output.ComputedValues["url"].Value)

	validateGateway(t, output.Resources, resource.ApplicationName, expectedIncludes)
}

func Test_Render_Fails_WithNoRoute(t *testing.T) {
	r := &Renderer{}

	properties := radclient.GatewayProperties{}
	resource := makeResource(t, properties)
	dependencies := map[string]renderers.RendererDependency{}
	additionalProperties := GetRuntimeOptions()

	output, err := r.Render(context.Background(), renderers.RenderOptions{Resource: resource, Dependencies: dependencies, Runtime: additionalProperties})
	require.Error(t, err)
	require.Equal(t, err.Error(), "must have at least one route when declaring a Gateway resource")
	require.Len(t, output.Resources, 0)
	require.Empty(t, output.SecretValues)
	require.Empty(t, output.ComputedValues)
}

func Test_Render_Fails_WithoutFQHostnameOrPrefix(t *testing.T) {
	r := &Renderer{}

	properties := radclient.GatewayProperties{
		Hostname: &radclient.GatewayPropertiesHostname{},
	}
	resource := makeResource(t, properties)
	dependencies := map[string]renderers.RendererDependency{}
	additionalProperties := GetRuntimeOptions()

	output, err := r.Render(context.Background(), renderers.RenderOptions{Resource: resource, Dependencies: dependencies, Runtime: additionalProperties})
	require.Error(t, err)
	require.Equal(t, err.Error(), "getting hostname failed with error: must provide either prefix or fullyQualifiedHostname if hostname is specified")
	require.Len(t, output.Resources, 0)
	require.Empty(t, output.SecretValues)
	require.Empty(t, output.ComputedValues)
}

func Test_Render_Single_Route(t *testing.T) {
	r := &Renderer{}

	var routes []*radclient.GatewayRoute
	routeName := "routename"
	destination := makeRouteResourceID(routeName)
	path := "/"
	route := radclient.GatewayRoute{
		Destination: &destination,
		Path:        &path,
	}
	routes = append(routes, &route)
	properties := radclient.GatewayProperties{
		Routes: routes,
	}
	resource := makeResource(t, properties)
	dependencies := map[string]renderers.RendererDependency{}
	additionalProperties := GetRuntimeOptions()
	expectedHostname := fmt.Sprintf("%s.%s.%s.nip.io", resourceName, applicationName, publicIP)
	expectedURL := "http://" + expectedHostname

	output, err := r.Render(context.Background(), renderers.RenderOptions{Resource: resource, Dependencies: dependencies, Runtime: additionalProperties})
	require.NoError(t, err)
	require.Len(t, output.Resources, 2)
	require.Empty(t, output.SecretValues)
	require.Equal(t, expectedURL, output.ComputedValues["url"].Value)

	expectedIncludes := []contourv1.Include{
		{
			Name: kubernetes.MakeResourceName(applicationName, routeName),
			Conditions: []contourv1.MatchCondition{
				{
					Prefix: path,
				},
			},
		},
	}

	validateGateway(t, output.Resources, expectedHostname, expectedIncludes)
	validateHttpRoute(t, output.Resources, routeName, 80, nil)
}

func Test_Render_Multiple_Routes(t *testing.T) {
	r := &Renderer{}

	var routes []*radclient.GatewayRoute
	routeAName := "routeaname"
	routeADestination := makeRouteResourceID(routeAName)
	routeAPath := "/routea"
	routeA := radclient.GatewayRoute{
		Destination: &routeADestination,
		Path:        &routeAPath,
	}
	routeBName := "routenbname"
	routeBDestination := makeRouteResourceID(routeBName)
	routeBPath := "/routeb"
	routeB := radclient.GatewayRoute{
		Destination: &routeBDestination,
		Path:        &routeBPath,
	}
	routes = append(routes, &routeA)
	routes = append(routes, &routeB)
	properties := radclient.GatewayProperties{
		Routes: routes,
	}
	resource := makeResource(t, properties)
	dependencies := map[string]renderers.RendererDependency{}
	additionalProperties := GetRuntimeOptions()
	expectedHostname := fmt.Sprintf("%s.%s.%s.nip.io", resourceName, applicationName, publicIP)
	expectedURL := "http://" + expectedHostname

	output, err := r.Render(context.Background(), renderers.RenderOptions{Resource: resource, Dependencies: dependencies, Runtime: additionalProperties})
	require.NoError(t, err)
	require.Len(t, output.Resources, 3)
	require.Empty(t, output.SecretValues)
	require.Equal(t, expectedURL, output.ComputedValues["url"].Value)

	expectedIncludes := []contourv1.Include{
		{
			Name: kubernetes.MakeResourceName(applicationName, routeAName),
			Conditions: []contourv1.MatchCondition{
				{
					Prefix: routeAPath,
				},
			},
		},
		{
			Name: kubernetes.MakeResourceName(applicationName, routeBName),
			Conditions: []contourv1.MatchCondition{
				{
					Prefix: routeBPath,
				},
			},
		},
	}

	validateGateway(t, output.Resources, expectedHostname, expectedIncludes)
	validateHttpRoute(t, output.Resources, routeAName, 80, nil)
	validateHttpRoute(t, output.Resources, routeBName, 80, nil)
}

func Test_Render_Route_WithPrefixRewrite(t *testing.T) {
	r := &Renderer{}

	var routes []*radclient.GatewayRoute
	routeName := "routename"
	destination := makeRouteResourceID(routeName)
	path := "/backend"
	rewrite := "/rewrite"
	route := radclient.GatewayRoute{
		Destination:   &destination,
		Path:          &path,
		ReplacePrefix: &rewrite,
	}
	routes = append(routes, &route)
	properties := radclient.GatewayProperties{
		Routes: routes,
	}
	resource := makeResource(t, properties)
	dependencies := map[string]renderers.RendererDependency{}
	additionalProperties := GetRuntimeOptions()
	expectedHostname := fmt.Sprintf("%s.%s.%s.nip.io", resourceName, applicationName, publicIP)
	expectedURL := "http://" + expectedHostname

	output, err := r.Render(context.Background(), renderers.RenderOptions{Resource: resource, Dependencies: dependencies, Runtime: additionalProperties})
	require.NoError(t, err)
	require.Len(t, output.Resources, 2)
	require.Empty(t, output.SecretValues)
	require.Equal(t, expectedURL, output.ComputedValues["url"].Value)

	expectedIncludes := []contourv1.Include{
		{
			Name: kubernetes.MakeResourceName(applicationName, routeName),
			Conditions: []contourv1.MatchCondition{
				{
					Prefix: path,
				},
			},
		},
	}
	validateGateway(t, output.Resources, expectedHostname, expectedIncludes)

	expectedPathRewritePolicy := &contourv1.PathRewritePolicy{
		ReplacePrefix: []contourv1.ReplacePrefix{
			{
				Prefix:      path,
				Replacement: rewrite,
			},
		},
	}
	validateHttpRoute(t, output.Resources, routeName, 80, expectedPathRewritePolicy)
}

func Test_Render_Route_WithMultiplePrefixRewrite(t *testing.T) {
	r := &Renderer{}

	var routes []*radclient.GatewayRoute
	routeAName := "routeaname"
	routeBName := "routebname"
	destinationA := makeRouteResourceID(routeAName)
	destinationB := makeRouteResourceID(routeBName)
	routeAPath := "/routea"
	routeA := radclient.GatewayRoute{
		Destination: &destinationA,
		Path:        &routeAPath,
	}
	routeBPath := "/routeb"
	routeBRewrite := "routebrewrite"
	routeB := radclient.GatewayRoute{
		Destination:   &destinationB,
		Path:          &routeBPath,
		ReplacePrefix: &routeBRewrite,
	}
	routeCPath := "/routec"
	routeCRewrite := "routecrewrite"
	routeC := radclient.GatewayRoute{
		Destination:   &destinationB,
		Path:          &routeCPath,
		ReplacePrefix: &routeCRewrite,
	}
	routeDPath := "/routed"
	routeD := radclient.GatewayRoute{
		Destination: &destinationB,
		Path:        &routeDPath,
	}
	routes = append(routes, &routeA)
	routes = append(routes, &routeB)
	routes = append(routes, &routeC)
	routes = append(routes, &routeD)
	properties := radclient.GatewayProperties{
		Routes: routes,
	}
	resource := makeResource(t, properties)
	dependencies := map[string]renderers.RendererDependency{}
	additionalProperties := GetRuntimeOptions()
	expectedHostname := fmt.Sprintf("%s.%s.%s.nip.io", resourceName, applicationName, publicIP)
	expectedURL := "http://" + expectedHostname

	output, err := r.Render(context.Background(), renderers.RenderOptions{Resource: resource, Dependencies: dependencies, Runtime: additionalProperties})
	require.NoError(t, err)
	require.Len(t, output.Resources, 3)
	require.Empty(t, output.SecretValues)
	require.Equal(t, expectedURL, output.ComputedValues["url"].Value)

	expectedIncludes := []contourv1.Include{
		{
			Name: kubernetes.MakeResourceName(applicationName, routeAName),
			Conditions: []contourv1.MatchCondition{
				{
					Prefix: routeAPath,
				},
			},
		},
		{
			Name: kubernetes.MakeResourceName(applicationName, routeBName),
			Conditions: []contourv1.MatchCondition{
				{
					Prefix: routeBPath,
				},
			},
		},
		{
			Name: kubernetes.MakeResourceName(applicationName, routeBName),
			Conditions: []contourv1.MatchCondition{
				{
					Prefix: routeCPath,
				},
			},
		},
		{
			Name: kubernetes.MakeResourceName(applicationName, routeBName),
			Conditions: []contourv1.MatchCondition{
				{
					Prefix: routeDPath,
				},
			},
		},
	}
	validateGateway(t, output.Resources, expectedHostname, expectedIncludes)

	expectedPathRewritePolicy := &contourv1.PathRewritePolicy{
		ReplacePrefix: []contourv1.ReplacePrefix{
			{
				Prefix:      routeBPath,
				Replacement: routeBRewrite,
			},
			{
				Prefix:      routeCPath,
				Replacement: routeCRewrite,
			},
		},
	}
	validateHttpRoute(t, output.Resources, routeAName, 80, nil)
	validateHttpRoute(t, output.Resources, routeBName, 80, expectedPathRewritePolicy)
}

func Test_Render_WithDependencies(t *testing.T) {
	r := &Renderer{}

	var httpRoutePort int32 = 81
	httpRoute := renderHttpRoute(t, httpRoutePort)

	var routes []*radclient.GatewayRoute
	routeName := "routename"
	routeDestination := makeRouteResourceID(routeName)
	routePath := "/routea"
	route := radclient.GatewayRoute{
		Destination: &routeDestination,
		Path:        &routePath,
	}
	routes = append(routes, &route)
	properties := radclient.GatewayProperties{
		Routes: routes,
	}
	resource := makeResource(t, properties)
	dependencies := map[string]renderers.RendererDependency{
		(makeResourceID(t, routeDestination).ID): {
			ResourceID: makeResourceID(t, routeDestination),
			Definition: map[string]interface{}{},
			ComputedValues: map[string]interface{}{
				"port": (httpRoute.ComputedValues["port"].Value).(int32),
			},
		},
	}

	additionalProperties := GetRuntimeOptions()
	expectedHostname := fmt.Sprintf("%s.%s.%s.nip.io", resourceName, applicationName, publicIP)
	expectedURL := "http://" + expectedHostname

	output, err := r.Render(context.Background(), renderers.RenderOptions{Resource: resource, Dependencies: dependencies, Runtime: additionalProperties})
	require.NoError(t, err)
	require.Len(t, output.Resources, 2)
	require.Empty(t, output.SecretValues)
	require.Equal(t, expectedURL, output.ComputedValues["url"].Value)

	expectedIncludes := []contourv1.Include{
		{
			Name: kubernetes.MakeResourceName(applicationName, routeName),
			Conditions: []contourv1.MatchCondition{
				{
					Prefix: routePath,
				},
			},
		},
	}

	validateGateway(t, output.Resources, expectedHostname, expectedIncludes)
	validateHttpRoute(t, output.Resources, routeName, httpRoutePort, nil)
}

func renderHttpRoute(t *testing.T, port int32) renderers.RendererOutput {
	r := &httproutev1alpha3.Renderer{}

	dependencies := map[string]renderers.RendererDependency{}
	properties := radclient.HTTPRouteProperties{
		Port: &port,
	}
	resource := makeResource(t, properties)

	output, err := r.Render(context.Background(), renderers.RenderOptions{Resource: resource, Dependencies: dependencies})
	require.NoError(t, err)

	return output
}

func validateGateway(t *testing.T, outputResources []outputresource.OutputResource, expectedHostname string, expectedIncludes []contourv1.Include) {
	gateway, gatewayOutputResource := kubernetes.FindGateway(outputResources)

	expectedGatewayOutputResource := outputresource.NewKubernetesOutputResource(resourcekinds.Gateway, outputresource.LocalIDGateway, gateway, gateway.ObjectMeta)
	require.Equal(t, expectedGatewayOutputResource, gatewayOutputResource)
	require.Equal(t, kubernetes.MakeResourceName(applicationName, resourceName), gateway.Name)
	require.Equal(t, applicationName, gateway.Namespace)
	require.Equal(t, kubernetes.MakeDescriptiveLabels(applicationName, resourceName), gateway.Labels)

	var expectedVirtualHost *contourv1.VirtualHost = nil
	var expectedGatewaySpec contourv1.HTTPProxySpec
	if expectedHostname != "" {
		expectedVirtualHost = &contourv1.VirtualHost{
			Fqdn: expectedHostname,
		}
		expectedGatewaySpec = contourv1.HTTPProxySpec{
			VirtualHost: expectedVirtualHost,
			Includes:    expectedIncludes,
		}
	} else {
		expectedGatewaySpec = contourv1.HTTPProxySpec{
			Includes: expectedIncludes,
		}
	}

	require.Equal(t, expectedVirtualHost, gateway.Spec.VirtualHost)
	require.Equal(t, expectedGatewaySpec, gateway.Spec)
}

func validateHttpRoute(t *testing.T, outputResources []outputresource.OutputResource, expectedRouteName string, expectedPort int32, expectedRewrite *contourv1.PathRewritePolicy) {
	expectedLocalID := fmt.Sprintf("%s-%s", outputresource.LocalIDHttpRoute, expectedRouteName)
	httpRoute, httpRouteOutputResource := kubernetes.FindHttpRouteByLocalID(outputResources, expectedLocalID)
	expectedHttpRouteOutputResource := outputresource.NewKubernetesOutputResource(resourcekinds.KubernetesHTTPRoute, expectedLocalID, httpRoute, httpRoute.ObjectMeta)
	require.Equal(t, expectedHttpRouteOutputResource, httpRouteOutputResource)

	require.Equal(t, kubernetes.MakeResourceName(applicationName, expectedRouteName), httpRoute.Name)
	require.Equal(t, applicationName, httpRoute.Namespace)
	require.Equal(t, kubernetes.MakeDescriptiveLabels(applicationName, expectedRouteName), httpRoute.Labels)

	require.Nil(t, httpRoute.Spec.VirtualHost)

	expectedServiceName := kubernetes.MakeResourceName(applicationName, expectedRouteName)

	expectedHttpRouteSpec := contourv1.HTTPProxySpec{
		Routes: []contourv1.Route{
			{
				Services: []contourv1.Service{
					{
						Name: expectedServiceName,
						Port: int(expectedPort),
					},
				},
				PathRewritePolicy: expectedRewrite,
			},
		},
	}

	require.Equal(t, expectedHttpRouteSpec, httpRoute.Spec)
}

func makeRouteResourceID(routeName string) string {
	return azresources.MakeID(
		subscriptionID,
		resourceGroup,
		azresources.ResourceType{
			Type: "Microsoft.CustomProviders",
			Name: "resourceProviders/radiusv3",
		},
		azresources.ResourceType{
			Type: "Application",
			Name: applicationName,
		},
		azresources.ResourceType{
			Type: "HttpRoute",
			Name: routeName,
		},
	)
}

func makeResource(t *testing.T, T any) renderers.RendererResource {
	b, err := json.Marshal(&T)
	require.NoError(t, err)

	definition := map[string]interface{}{}
	err = json.Unmarshal(b, &definition)
	require.NoError(t, err)

	return renderers.RendererResource{
		ApplicationName: applicationName,
		ResourceName:    resourceName,
		ResourceType:    ResourceType,
		Definition:      definition,
	}
}

func makeResourceID(t *testing.T, resourceID string) azresources.ResourceID {
	id, err := azresources.Parse(resourceID)
	require.NoError(t, err)

	return id
}

func makeTestGateway(config radclient.GatewayProperties) (radclient.GatewayProperties, []contourv1.Include) {
	routeName := "routeName"
	routeDestination := makeRouteResourceID("routeName")
	routePath := "/"
	defaultRoute := radclient.GatewayRoute{
		Destination: &routeDestination,
		Path:        &routePath,
	}

	includes := []contourv1.Include{
		{
			Name: kubernetes.MakeResourceName(applicationName, routeName),
			Conditions: []contourv1.MatchCondition{
				{
					Prefix: routePath,
				},
			},
		},
	}

	properties := radclient.GatewayProperties{
		Hostname: config.Hostname,
		Routes: []*radclient.GatewayRoute{
			&defaultRoute,
		},
	}

	return properties, includes
}

func GetRuntimeOptions() renderers.RuntimeOptions {
	additionalProperties := renderers.RuntimeOptions{
		Gateway: renderers.GatewayOptions{
			PublicIP: publicIP,
		},
	}
	return additionalProperties
}