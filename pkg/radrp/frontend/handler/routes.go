// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/project-radius/radius/pkg/azure/azresources"
	"github.com/project-radius/radius/pkg/radrp/frontend/resourceprovider"
	"github.com/project-radius/radius/pkg/radrp/rest"
	"github.com/project-radius/radius/pkg/radrp/schema"
)

func AddRoutes(rp resourceprovider.ResourceProvider, router *mux.Router, validatorFactory ValidatorFactory, swaggerDocRoute string) {
	// Nothing for now

	h := Handler{RP: rp, ValidatorFactory: validatorFactory, PathPrefix: swaggerDocRoute}
	var subrouter *mux.Router

	var providerPath = fmt.Sprintf(
		"%s/subscriptions/{%s}/resourcegroups/{%s}/providers/Microsoft.CustomProviders/resourceProviders/radiusv3",
		swaggerDocRoute,
		azresources.SubscriptionIDKey,
		azresources.ResourceGroupKey)

	if swaggerDocRoute != "" {
		router.Path(swaggerDocRoute).Methods("GET").HandlerFunc(h.GetSwaggerDoc)
	}

	router.Path(fmt.Sprintf("%s/listSecrets", providerPath)).Methods("POST").HandlerFunc(h.ListSecrets)

	var applicationCollectionPath = fmt.Sprintf("%s/Application", providerPath)
	var applicationItemPath = fmt.Sprintf("%s/{%s}", applicationCollectionPath, azresources.ApplicationNameKey)

	var resourceCollectionPath = fmt.Sprintf("%s/{%s}", applicationItemPath, azresources.ResourceTypeKey)
	var resourceItemPath = fmt.Sprintf("%s/{%s}", resourceCollectionPath, azresources.ResourceNameKey)
	var operationItemPath = fmt.Sprintf("%s/{%s}/{%s}", resourceItemPath, "OperationResults", azresources.OperationIDKey)

	var allResourceCollectionPath = fmt.Sprintf("%s/%s", applicationItemPath, schema.GenericResourceType)
	var allResourceItemPath = fmt.Sprintf("%s/{%s}", allResourceCollectionPath, azresources.ResourceNameKey)

	router.Path(applicationCollectionPath).Methods("GET").HandlerFunc(h.ListApplications)
	subrouter = router.Path(applicationItemPath).Subrouter()
	subrouter.Methods("GET").HandlerFunc(h.GetApplication)
	subrouter.Methods("PUT").HandlerFunc(h.UpdateApplication)
	subrouter.Methods("DELETE").HandlerFunc(h.DeleteApplication)

	router.Path(allResourceCollectionPath).Methods("GET").HandlerFunc(h.ListAllV3ResourcesByApplication)
	router.Path(allResourceItemPath).HandlerFunc(notSupported)

	router.Path(resourceCollectionPath).Methods("GET").HandlerFunc(h.ListResources)
	subrouter = router.Path(resourceItemPath).Subrouter()
	subrouter.Methods("GET").HandlerFunc(h.GetResource)
	subrouter.Methods("PUT").HandlerFunc(h.UpdateResource)
	subrouter.Methods("DELETE").HandlerFunc(h.DeleteResource)

	subrouter = router.Path(operationItemPath).Subrouter()
	subrouter.Methods("GET").HandlerFunc(h.GetOperation)
}

func notSupported(w http.ResponseWriter, req *http.Request) {
	response := rest.NewBadRequestResponse(fmt.Sprintf("Route not suported: %s", req.URL.Path))
	_ = response.Apply(req.Context(), w, req)
}
