/*
 * OSCAL REST
 *
 * A draft proposal from [Easy Dynamics](https://www.easydynamics.com) of a REST API specification for interacting   with [OSCAL](https://pages.nist.gov/OSCAL/) models.  Standardized data models like OSCAL lay the groundwork for interoperability of systems, and an ecosystem of meaningful integrations can be brought to life through a standardized REST API. That interface needs to define simple CRUD operations, but should also describe how to manipulate relationships and make partial changes.  Such an API will likely see the most success across various vendors and projects when maintained by a standards body or community, and we're looking to get that conversation started with this effort.  ## Identifier Scope Note that all object identifiers are **scoped to the system** implementing the REST API   and are expected to be unique.  <details>   <summary>Example</summary>   To associate a Party with a known Component Definition with a UUID you might:   1. Search for the Party:       ```       GET /parties/search?query=bob       ```       ```       [         {           \"uuid\": \"d834ed5e-9652-4b78-87e7-a9f8686f4e60\",           \"type\": \"person\",           \"name\": \"Bob Johnson\"         }       ]       ```   2. Use the found Party's ID to associate with the component definition:       ```       PUT /component-definitions/599b6fa5-3e18-4580-bd8f-8a181776c6e8/parties/d834ed5e-9652-4b78-87e7-a9f8686f4e60       ``` </details>  ## Partial Payloads in PATCH Requests All updates to 'root' OSCAL objects via `PATCH` requests should accept a partial payload containing only the changed data and the entire updated object should be returned.  The UUID of the root object must not change in the payload and implementations should throw an error if such a request is made. <details>   <summary>Example</summary>   Example request:   ```   PATCH /system-security-plans/cff8385f-108e-40a5-8f7a-82f3dc0eaba8   {     \"system-security-plan\": {       \"uuid\": \"cff8385f-108e-40a5-8f7a-82f3dc0eaba8\",       \"metadata\": {         \"title\": \"Some New Title\"       }     }   }   ``` </details> 
 *
 * API version: 0.1.0
 * Contact: info@easydynamics.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name        string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method      string
	// Pattern is the pattern of the URI.
	Pattern     string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		http.MethodGet,
		"/oscal/v1/",
		Index,
	},

	{
		"AddCatalog",
		http.MethodPost,
		"/oscal/v1/catalogs",
		AddCatalog,
	},

	{
		"DeleteCatalog",
		http.MethodDelete,
		"/oscal/v1/catalogs/:catalogId",
		DeleteCatalog,
	},

	{
		"FindCatalogsByName",
		http.MethodGet,
		"/oscal/v1/catalogs/search",
		FindCatalogsByName,
	},

	{
		"GetCatalogById",
		http.MethodGet,
		"/oscal/v1/catalogs/:catalogId",
		GetCatalogById,
	},

	{
		"GetCatalogs",
		http.MethodGet,
		"/oscal/v1/catalogs",
		GetCatalogs,
	},

	{
		"ReplaceCatalog",
		http.MethodPut,
		"/oscal/v1/catalogs/:catalogId",
		ReplaceCatalog,
	},

	{
		"UpdateCatalog",
		http.MethodPatch,
		"/oscal/v1/catalogs/:catalogId",
		UpdateCatalog,
	},

	{
		"AddComponentControlImplementationToComponentDefinition",
		http.MethodPost,
		"/oscal/v1/component-definitions/:componentDefinitionId/components/:componentId/control-implementations",
		AddComponentControlImplementationToComponentDefinition,
	},

	{
		"AddComponentDefinition",
		http.MethodPost,
		"/oscal/v1/component-definitions",
		AddComponentDefinition,
	},

	{
		"AddComponentToComponentDefinition",
		http.MethodPost,
		"/oscal/v1/component-definitions/:componentDefinitionId/components",
		AddComponentToComponentDefinition,
	},

	{
		"AddPartyToComponentDefinition",
		http.MethodPut,
		"/oscal/v1/component-definitions/:componentDefinitionId/parties/:partyId",
		AddPartyToComponentDefinition,
	},

	{
		"AddPartyToComponentDefinitionRole",
		http.MethodPost,
		"/oscal/v1/component-definitions/:componentDefinitionId/responsible-parties/:roleId/parties/:partyId",
		AddPartyToComponentDefinitionRole,
	},

	{
		"AddRoleToComponentDefinition",
		http.MethodPost,
		"/oscal/v1/component-definitions/:componentDefinitionId/roles",
		AddRoleToComponentDefinition,
	},

	{
		"DeleteComponentDefinition",
		http.MethodDelete,
		"/oscal/v1/component-definitions/:componentDefinitionId",
		DeleteComponentDefinition,
	},

	{
		"FindComponentDefinitionsByName",
		http.MethodGet,
		"/oscal/v1/component-definitions/search",
		FindComponentDefinitionsByName,
	},

	{
		"GetComponentDefinitionById",
		http.MethodGet,
		"/oscal/v1/component-definitions/:componentDefinitionId",
		GetComponentDefinitionById,
	},

	{
		"GetComponentDefinitions",
		http.MethodGet,
		"/oscal/v1/component-definitions",
		GetComponentDefinitions,
	},

	{
		"RemoveComponentControlImplementationFromComponentDefinition",
		http.MethodDelete,
		"/oscal/v1/component-definitions/:componentDefinitionId/components/:componentId/control-implementations/:componentControlImplementationId",
		RemoveComponentControlImplementationFromComponentDefinition,
	},

	{
		"RemoveComponentFromComponentDefinition",
		http.MethodDelete,
		"/oscal/v1/component-definitions/:componentDefinitionId/components/:componentId",
		RemoveComponentFromComponentDefinition,
	},

	{
		"RemovePartyFromComponentDefinition",
		http.MethodDelete,
		"/oscal/v1/component-definitions/:componentDefinitionId/parties/:partyId",
		RemovePartyFromComponentDefinition,
	},

	{
		"RemovePartyFromComponentDefinitionRole",
		http.MethodDelete,
		"/oscal/v1/component-definitions/:componentDefinitionId/responsible-parties/:roleId/parties/:partyId",
		RemovePartyFromComponentDefinitionRole,
	},

	{
		"RemoveRoleFromComponentDefinition",
		http.MethodDelete,
		"/oscal/v1/component-definitions/:componentDefinitionId/roles/:roleId",
		RemoveRoleFromComponentDefinition,
	},

	{
		"ReplaceComponentDefinition",
		http.MethodPut,
		"/oscal/v1/component-definitions/:componentDefinitionId",
		ReplaceComponentDefinition,
	},

	{
		"UpdateComponentControlImplementationInComponentDefinition",
		http.MethodPatch,
		"/oscal/v1/component-definitions/:componentDefinitionId/components/:componentId/control-implementations/:componentControlImplementationId",
		UpdateComponentControlImplementationInComponentDefinition,
	},

	{
		"UpdateComponentDefinition",
		http.MethodPatch,
		"/oscal/v1/component-definitions/:componentDefinitionId",
		UpdateComponentDefinition,
	},

	{
		"UpdateComponentInComponentDefinition",
		http.MethodPatch,
		"/oscal/v1/component-definitions/:componentDefinitionId/components/:componentId",
		UpdateComponentInComponentDefinition,
	},

	{
		"AddParty",
		http.MethodPost,
		"/oscal/v1/parties",
		AddParty,
	},

	{
		"DeleteParty",
		http.MethodDelete,
		"/oscal/v1/parties/:partyId",
		DeleteParty,
	},

	{
		"FindPartiesByName",
		http.MethodGet,
		"/oscal/v1/parties/search",
		FindPartiesByName,
	},

	{
		"GetParties",
		http.MethodGet,
		"/oscal/v1/parties",
		GetParties,
	},

	{
		"GetPartyById",
		http.MethodGet,
		"/oscal/v1/parties/:partyId",
		GetPartyById,
	},

	{
		"UpdateParty",
		http.MethodPatch,
		"/oscal/v1/parties/:partyId",
		UpdateParty,
	},

	{
		"AddCatalogToProfile",
		http.MethodPut,
		"/oscal/v1/profiles/:profileId/imported-catalogs/:catalogId",
		AddCatalogToProfile,
	},

	{
		"AddPartyToProfile",
		http.MethodPut,
		"/oscal/v1/profiles/:profileId/parties/:partyId",
		AddPartyToProfile,
	},

	{
		"AddPartyToProfileRole",
		http.MethodPost,
		"/oscal/v1/profiles/:profileId/responsible-parties/:roleId/parties/:partyId",
		AddPartyToProfileRole,
	},

	{
		"AddProfile",
		http.MethodPost,
		"/oscal/v1/profiles",
		AddProfile,
	},

	{
		"AddProfileToProfile",
		http.MethodPut,
		"/oscal/v1/profiles/:profileId/imported-profiles/:importProfileId",
		AddProfileToProfile,
	},

	{
		"AddRoleToProfile",
		http.MethodPost,
		"/oscal/v1/profiles/:profileId/roles",
		AddRoleToProfile,
	},

	{
		"DeleteProfile",
		http.MethodDelete,
		"/oscal/v1/profiles/:profileId",
		DeleteProfile,
	},

	{
		"FindProfilesByName",
		http.MethodGet,
		"/oscal/v1/profiles/search",
		FindProfilesByName,
	},

	{
		"GetProfileById",
		http.MethodGet,
		"/oscal/v1/profiles/:profileId",
		GetProfileById,
	},

	{
		"GetProfiles",
		http.MethodGet,
		"/oscal/v1/profiles",
		GetProfiles,
	},

	{
		"RemoveCatalogfromProfile",
		http.MethodDelete,
		"/oscal/v1/profiles/:profileId/imported-catalogs/:catalogId",
		RemoveCatalogfromProfile,
	},

	{
		"RemovePartyfromProfile",
		http.MethodDelete,
		"/oscal/v1/profiles/:profileId/parties/:partyId",
		RemovePartyfromProfile,
	},

	{
		"RemovePartyfromProfileRole",
		http.MethodDelete,
		"/oscal/v1/profiles/:profileId/responsible-parties/:roleId/parties/:partyId",
		RemovePartyfromProfileRole,
	},

	{
		"RemoveProfilefromProfile",
		http.MethodDelete,
		"/oscal/v1/profiles/:profileId/imported-profiles/:importProfileId",
		RemoveProfilefromProfile,
	},

	{
		"RemoveRolefromProfile",
		http.MethodDelete,
		"/oscal/v1/profiles/:profileId/roles/:roleId",
		RemoveRolefromProfile,
	},

	{
		"ReplaceProfile",
		http.MethodPut,
		"/oscal/v1/profiles/:profileId",
		ReplaceProfile,
	},

	{
		"UpdateProfile",
		http.MethodPatch,
		"/oscal/v1/profiles",
		UpdateProfile,
	},

	{
		"AddComponentToSsp",
		http.MethodPut,
		"/oscal/v1/system-security-plans/:sspId/components/:componentId",
		AddComponentToSsp,
	},

	{
		"AddImplementedRequirementToSsp",
		http.MethodPost,
		"/oscal/v1/system-security-plans/:sspId/control-implementation/implemented-requirements",
		AddImplementedRequirementToSsp,
	},

	{
		"AddPartyToSsp",
		http.MethodPut,
		"/oscal/v1/system-security-plans/:sspId/parties/:partyId",
		AddPartyToSsp,
	},

	{
		"AddSsp",
		http.MethodPost,
		"/oscal/v1/system-security-plans",
		AddSsp,
	},

	{
		"DeleteSsp",
		http.MethodDelete,
		"/oscal/v1/system-security-plans/:sspId",
		DeleteSsp,
	},

	{
		"FindSspsByName",
		http.MethodGet,
		"/oscal/v1/system-security-plans/search",
		FindSspsByName,
	},

	{
		"GetSspById",
		http.MethodGet,
		"/oscal/v1/system-security-plans/:sspId",
		GetSspById,
	},

	{
		"GetSsps",
		http.MethodGet,
		"/oscal/v1/system-security-plans",
		GetSsps,
	},

	{
		"RemoveComponentfromSsp",
		http.MethodDelete,
		"/oscal/v1/system-security-plans/:sspId/components/:componentId",
		RemoveComponentfromSsp,
	},

	{
		"RemovePartyfromSsp",
		http.MethodDelete,
		"/oscal/v1/system-security-plans/:sspId/parties/:partyId",
		RemovePartyfromSsp,
	},

	{
		"ReplaceSsp",
		http.MethodPut,
		"/oscal/v1/system-security-plans/:sspId",
		ReplaceSsp,
	},

	{
		"UpdateImplementedRequirementOfSsp",
		http.MethodPut,
		"/oscal/v1/system-security-plans/:sspId/control-implementation/implemented-requirements/:implementedRequirementId",
		UpdateImplementedRequirementOfSsp,
	},

	{
		"UpdateSsp",
		http.MethodPatch,
		"/oscal/v1/system-security-plans/:sspId",
		UpdateSsp,
	},
}
