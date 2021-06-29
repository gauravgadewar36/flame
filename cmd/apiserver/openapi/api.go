/*
 * Fledge REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
	"os"

	objects2 "wwwin-github.cisco.com/fledge/fledge/pkg/objects"
)

// DesignApiRouter defines the required methods for binding the api requests to a responses for the DesignApi
// The DesignApiRouter implementation should parse necessary information from the http request,
// pass the data to a DesignApiServicer to perform the required actions, then write the service results to the http response.
type DesignApiRouter interface {
	CreateDesign(http.ResponseWriter, *http.Request)
	GetDesign(http.ResponseWriter, *http.Request)
}

// DesignCodeApiRouter defines the required methods for binding the api requests to a responses for the DesignCodeApi
// The DesignCodeApiRouter implementation should parse necessary information from the http request,
// pass the data to a DesignCodeApiServicer to perform the required actions, then write the service results to the http response.
type DesignCodeApiRouter interface {
	GetDesignCode(http.ResponseWriter, *http.Request)
	UpdateDesignCode(http.ResponseWriter, *http.Request)
}

// DesignSchemaApiRouter defines the required methods for binding the api requests to a responses for the DesignSchemaApi
// The DesignSchemaApiRouter implementation should parse necessary information from the http request,
// pass the data to a DesignSchemaApiServicer to perform the required actions, then write the service results to the http response.
type DesignSchemaApiRouter interface {
	GetDesignSchema(http.ResponseWriter, *http.Request)
	UpdateDesignSchema(http.ResponseWriter, *http.Request)
}

// DesignsApiRouter defines the required methods for binding the api requests to a responses for the DesignsApi
// The DesignsApiRouter implementation should parse necessary information from the http request,
// pass the data to a DesignsApiServicer to perform the required actions, then write the service results to the http response.
type DesignsApiRouter interface {
	GetDesigns(http.ResponseWriter, *http.Request)
}

// DesignApiServicer defines the api actions for the DesignApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DesignApiServicer interface {
	CreateDesign(context.Context, string, objects2.DesignInfo) (ImplResponse, error)
	GetDesign(context.Context, string, string) (ImplResponse, error)
}

// DesignCodeApiServicer defines the api actions for the DesignCodeApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DesignCodeApiServicer interface {
	GetDesignCode(context.Context, string, string) (ImplResponse, error)
	UpdateDesignCode(context.Context, string, string, *os.File) (ImplResponse, error)
}

// DesignSchemaApiServicer defines the api actions for the DesignSchemaApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DesignSchemaApiServicer interface {
	GetDesignSchema(context.Context, string, string, string, string) (ImplResponse, error)
	UpdateDesignSchema(context.Context, string, string, objects2.DesignSchema) (ImplResponse, error)
}

// DesignsApiServicer defines the api actions for the DesignsApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DesignsApiServicer interface {
	GetDesigns(context.Context, string, int32) (ImplResponse, error)
}