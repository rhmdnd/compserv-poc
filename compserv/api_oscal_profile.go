/*
 * OSCAL REST
 *
 * A draft proposal from [Easy Dynamics](https://www.easydynamics.com) of a REST API specification for interacting   with [OSCAL](https://pages.nist.gov/OSCAL/) models.  Standardized data models like OSCAL lay the groundwork for interoperability of systems, and an ecosystem of meaningful integrations can be brought to life through a standardized REST API. That interface needs to define simple CRUD operations, but should also describe how to manipulate relationships and make partial changes.  Such an API will likely see the most success across various vendors and projects when maintained by a standards body or community, and we're looking to get that conversation started with this effort.  ## Identifier Scope Note that all object identifiers are **scoped to the system** implementing the REST API   and are expected to be unique.  <details>   <summary>Example</summary>   To associate a Party with a known Component Definition with a UUID you might:   1. Search for the Party:       ```       GET /parties/search?query=bob       ```       ```       [         {           \"uuid\": \"d834ed5e-9652-4b78-87e7-a9f8686f4e60\",           \"type\": \"person\",           \"name\": \"Bob Johnson\"         }       ]       ```   2. Use the found Party's ID to associate with the component definition:       ```       PUT /component-definitions/599b6fa5-3e18-4580-bd8f-8a181776c6e8/parties/d834ed5e-9652-4b78-87e7-a9f8686f4e60       ``` </details>  ## Partial Payloads in PATCH Requests All updates to 'root' OSCAL objects via `PATCH` requests should accept a partial payload containing only the changed data and the entire updated object should be returned.  The UUID of the root object must not change in the payload and implementations should throw an error if such a request is made. <details>   <summary>Example</summary>   Example request:   ```   PATCH /system-security-plans/cff8385f-108e-40a5-8f7a-82f3dc0eaba8   {     \"system-security-plan\": {       \"uuid\": \"cff8385f-108e-40a5-8f7a-82f3dc0eaba8\",       \"metadata\": {         \"title\": \"Some New Title\"       }     }   }   ``` </details> 
 *
 * API version: 0.1.0
 * Contact: info@easydynamics.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package compserv

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// OSCALProfileApiController binds http requests to an api service and writes the service results to the http response
type OSCALProfileApiController struct {
	service OSCALProfileApiServicer
	errorHandler ErrorHandler
}

// OSCALProfileApiOption for how the controller is set up.
type OSCALProfileApiOption func(*OSCALProfileApiController)

// WithOSCALProfileApiErrorHandler inject ErrorHandler into controller
func WithOSCALProfileApiErrorHandler(h ErrorHandler) OSCALProfileApiOption {
	return func(c *OSCALProfileApiController) {
		c.errorHandler = h
	}
}

// NewOSCALProfileApiController creates a default api controller
func NewOSCALProfileApiController(s OSCALProfileApiServicer, opts ...OSCALProfileApiOption) Router {
	controller := &OSCALProfileApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the OSCALProfileApiController
func (c *OSCALProfileApiController) Routes() Routes {
	return Routes{ 
		{
			"AddCatalogToProfile",
			strings.ToUpper("Put"),
			"/oscal/v1/profiles/{profileId}/imported-catalogs/{catalogId}",
			c.AddCatalogToProfile,
		},
		{
			"AddPartyToProfile",
			strings.ToUpper("Put"),
			"/oscal/v1/profiles/{profileId}/parties/{partyId}",
			c.AddPartyToProfile,
		},
		{
			"AddPartyToProfileRole",
			strings.ToUpper("Post"),
			"/oscal/v1/profiles/{profileId}/responsible-parties/{roleId}/parties/{partyId}",
			c.AddPartyToProfileRole,
		},
		{
			"AddProfile",
			strings.ToUpper("Post"),
			"/oscal/v1/profiles",
			c.AddProfile,
		},
		{
			"AddProfileToProfile",
			strings.ToUpper("Put"),
			"/oscal/v1/profiles/{profileId}/imported-profiles/{importProfileId}",
			c.AddProfileToProfile,
		},
		{
			"AddRoleToProfile",
			strings.ToUpper("Post"),
			"/oscal/v1/profiles/{profileId}/roles",
			c.AddRoleToProfile,
		},
		{
			"DeleteProfile",
			strings.ToUpper("Delete"),
			"/oscal/v1/profiles/{profileId}",
			c.DeleteProfile,
		},
		{
			"FindProfilesByName",
			strings.ToUpper("Get"),
			"/oscal/v1/profiles/search",
			c.FindProfilesByName,
		},
		{
			"GetProfileById",
			strings.ToUpper("Get"),
			"/oscal/v1/profiles/{profileId}",
			c.GetProfileById,
		},
		{
			"GetProfiles",
			strings.ToUpper("Get"),
			"/oscal/v1/profiles",
			c.GetProfiles,
		},
		{
			"RemoveCatalogfromProfile",
			strings.ToUpper("Delete"),
			"/oscal/v1/profiles/{profileId}/imported-catalogs/{catalogId}",
			c.RemoveCatalogfromProfile,
		},
		{
			"RemovePartyfromProfile",
			strings.ToUpper("Delete"),
			"/oscal/v1/profiles/{profileId}/parties/{partyId}",
			c.RemovePartyfromProfile,
		},
		{
			"RemovePartyfromProfileRole",
			strings.ToUpper("Delete"),
			"/oscal/v1/profiles/{profileId}/responsible-parties/{roleId}/parties/{partyId}",
			c.RemovePartyfromProfileRole,
		},
		{
			"RemoveProfilefromProfile",
			strings.ToUpper("Delete"),
			"/oscal/v1/profiles/{profileId}/imported-profiles/{importProfileId}",
			c.RemoveProfilefromProfile,
		},
		{
			"RemoveRolefromProfile",
			strings.ToUpper("Delete"),
			"/oscal/v1/profiles/{profileId}/roles/{roleId}",
			c.RemoveRolefromProfile,
		},
		{
			"ReplaceProfile",
			strings.ToUpper("Put"),
			"/oscal/v1/profiles/{profileId}",
			c.ReplaceProfile,
		},
		{
			"UpdateProfile",
			strings.ToUpper("Patch"),
			"/oscal/v1/profiles",
			c.UpdateProfile,
		},
	}
}

// AddCatalogToProfile - Defines a catalog import in an OSCAL profile
func (c *OSCALProfileApiController) AddCatalogToProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	profileIdParam := params["profileId"]
	
	catalogIdParam := params["catalogId"]
	
	apiKeyParam := r.Header.Get("api_key")
	inlineObjectParam := InlineObject{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&inlineObjectParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertInlineObjectRequired(inlineObjectParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AddCatalogToProfile(r.Context(), profileIdParam, catalogIdParam, apiKeyParam, inlineObjectParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// AddPartyToProfile - Associates a party with an OSCAL profile
func (c *OSCALProfileApiController) AddPartyToProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	profileIdParam := params["profileId"]
	
	partyIdParam := params["partyId"]
	
	apiKeyParam := r.Header.Get("api_key")
	result, err := c.service.AddPartyToProfile(r.Context(), profileIdParam, partyIdParam, apiKeyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// AddPartyToProfileRole - Associates a party with role within an OSCAL profile
func (c *OSCALProfileApiController) AddPartyToProfileRole(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	profileIdParam := params["profileId"]
	
	roleIdParam := params["roleId"]
	
	partyIdParam := params["partyId"]
	
	apiKeyParam := r.Header.Get("api_key")
	result, err := c.service.AddPartyToProfileRole(r.Context(), profileIdParam, roleIdParam, partyIdParam, apiKeyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// AddProfile - Adds a new OSCAL profile
func (c *OSCALProfileApiController) AddProfile(w http.ResponseWriter, r *http.Request) {
	bodyParam := OscalProfile{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertOscalProfileRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AddProfile(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// AddProfileToProfile - Defines the import of another profile in an OSCAL profile
func (c *OSCALProfileApiController) AddProfileToProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	profileIdParam := params["profileId"]
	
	importProfileIdParam := params["importProfileId"]
	
	apiKeyParam := r.Header.Get("api_key")
	inlineObject1Param := InlineObject1{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&inlineObject1Param); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertInlineObject1Required(inlineObject1Param); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AddProfileToProfile(r.Context(), profileIdParam, importProfileIdParam, apiKeyParam, inlineObject1Param)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// AddRoleToProfile - Adds a new role to an OSCAL profile
func (c *OSCALProfileApiController) AddRoleToProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	profileIdParam := params["profileId"]
	
	oscalRoleParam := OscalRole{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&oscalRoleParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertOscalRoleRequired(oscalRoleParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	apiKeyParam := r.Header.Get("api_key")
	result, err := c.service.AddRoleToProfile(r.Context(), profileIdParam, oscalRoleParam, apiKeyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeleteProfile - Deletes an OSCAL profile
func (c *OSCALProfileApiController) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	profileIdParam := params["profileId"]
	
	apiKeyParam := r.Header.Get("api_key")
	result, err := c.service.DeleteProfile(r.Context(), profileIdParam, apiKeyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// FindProfilesByName - Searches for OSCAL profiles by name
func (c *OSCALProfileApiController) FindProfilesByName(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	queryParam := query.Get("query")
	result, err := c.service.FindProfilesByName(r.Context(), queryParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetProfileById - Finds an OSCAL profile by ID
func (c *OSCALProfileApiController) GetProfileById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	profileIdParam := params["profileId"]
	
	result, err := c.service.GetProfileById(r.Context(), profileIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetProfiles - Returns all OSCAL profiles
func (c *OSCALProfileApiController) GetProfiles(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetProfiles(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// RemoveCatalogfromProfile - Removes a catalog from an OSCAL profile imports
func (c *OSCALProfileApiController) RemoveCatalogfromProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	profileIdParam := params["profileId"]
	
	catalogIdParam := params["catalogId"]
	
	apiKeyParam := r.Header.Get("api_key")
	result, err := c.service.RemoveCatalogfromProfile(r.Context(), profileIdParam, catalogIdParam, apiKeyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// RemovePartyfromProfile - Removes a party from an OSCAL profile
func (c *OSCALProfileApiController) RemovePartyfromProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	profileIdParam := params["profileId"]
	
	partyIdParam := params["partyId"]
	
	apiKeyParam := r.Header.Get("api_key")
	result, err := c.service.RemovePartyfromProfile(r.Context(), profileIdParam, partyIdParam, apiKeyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// RemovePartyfromProfileRole - Removes a party from a role within an OSCAL profile
func (c *OSCALProfileApiController) RemovePartyfromProfileRole(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	profileIdParam := params["profileId"]
	
	roleIdParam := params["roleId"]
	
	partyIdParam := params["partyId"]
	
	apiKeyParam := r.Header.Get("api_key")
	result, err := c.service.RemovePartyfromProfileRole(r.Context(), profileIdParam, roleIdParam, partyIdParam, apiKeyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// RemoveProfilefromProfile - Removes a another profile from an OSCAL profile imports
func (c *OSCALProfileApiController) RemoveProfilefromProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	profileIdParam := params["profileId"]
	
	importProfileIdParam := params["importProfileId"]
	
	apiKeyParam := r.Header.Get("api_key")
	result, err := c.service.RemoveProfilefromProfile(r.Context(), profileIdParam, importProfileIdParam, apiKeyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// RemoveRolefromProfile - Removes a role from an OSCAL profile
func (c *OSCALProfileApiController) RemoveRolefromProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	profileIdParam := params["profileId"]
	
	roleIdParam := params["roleId"]
	
	apiKeyParam := r.Header.Get("api_key")
	result, err := c.service.RemoveRolefromProfile(r.Context(), profileIdParam, roleIdParam, apiKeyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ReplaceProfile - Replaces an existing OSCAL profile
func (c *OSCALProfileApiController) ReplaceProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	profileIdParam := params["profileId"]
	
	bodyParam := OscalProfile{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertOscalProfileRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.ReplaceProfile(r.Context(), profileIdParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpdateProfile - Updates an existing OSCAL profile
func (c *OSCALProfileApiController) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	bodyParam := OscalProfileUpdateExample{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertOscalProfileUpdateExampleRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateProfile(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}