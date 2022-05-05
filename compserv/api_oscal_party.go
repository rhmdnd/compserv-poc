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

// OSCALPartyApiController binds http requests to an api service and writes the service results to the http response
type OSCALPartyApiController struct {
	service OSCALPartyApiServicer
	errorHandler ErrorHandler
}

// OSCALPartyApiOption for how the controller is set up.
type OSCALPartyApiOption func(*OSCALPartyApiController)

// WithOSCALPartyApiErrorHandler inject ErrorHandler into controller
func WithOSCALPartyApiErrorHandler(h ErrorHandler) OSCALPartyApiOption {
	return func(c *OSCALPartyApiController) {
		c.errorHandler = h
	}
}

// NewOSCALPartyApiController creates a default api controller
func NewOSCALPartyApiController(s OSCALPartyApiServicer, opts ...OSCALPartyApiOption) Router {
	controller := &OSCALPartyApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the OSCALPartyApiController
func (c *OSCALPartyApiController) Routes() Routes {
	return Routes{ 
		{
			"AddParty",
			strings.ToUpper("Post"),
			"/oscal/v1/parties",
			c.AddParty,
		},
		{
			"DeleteParty",
			strings.ToUpper("Delete"),
			"/oscal/v1/parties/{partyId}",
			c.DeleteParty,
		},
		{
			"FindPartiesByName",
			strings.ToUpper("Get"),
			"/oscal/v1/parties/search",
			c.FindPartiesByName,
		},
		{
			"GetParties",
			strings.ToUpper("Get"),
			"/oscal/v1/parties",
			c.GetParties,
		},
		{
			"GetPartyById",
			strings.ToUpper("Get"),
			"/oscal/v1/parties/{partyId}",
			c.GetPartyById,
		},
		{
			"UpdateParty",
			strings.ToUpper("Patch"),
			"/oscal/v1/parties/{partyId}",
			c.UpdateParty,
		},
	}
}

// AddParty - Adds a new OSCAL party
func (c *OSCALPartyApiController) AddParty(w http.ResponseWriter, r *http.Request) {
	bodyParam := OscalParty{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertOscalPartyRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AddParty(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeleteParty - Deletes an existing OSCAL party
func (c *OSCALPartyApiController) DeleteParty(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	partyIdParam := params["partyId"]
	
	apiKeyParam := r.Header.Get("api_key")
	result, err := c.service.DeleteParty(r.Context(), partyIdParam, apiKeyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// FindPartiesByName - Searches for OSCAL party by name
func (c *OSCALPartyApiController) FindPartiesByName(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	queryParam := query.Get("query")
	result, err := c.service.FindPartiesByName(r.Context(), queryParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetParties - Returns all OSCAL parties
func (c *OSCALPartyApiController) GetParties(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetParties(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetPartyById - Finds an OSCAL party by ID
func (c *OSCALPartyApiController) GetPartyById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	partyIdParam := params["partyId"]
	
	result, err := c.service.GetPartyById(r.Context(), partyIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpdateParty - Updates an existing OSCAL party
func (c *OSCALPartyApiController) UpdateParty(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	partyIdParam := params["partyId"]
	
	bodyParam := OscalParty{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertOscalPartyRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateParty(r.Context(), partyIdParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}