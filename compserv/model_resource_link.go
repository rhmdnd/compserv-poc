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

// ResourceLink - A pointer to an external resource with an optional hash for verification and change detection.
type ResourceLink struct {

	// A resolvable URI reference to a resource.
	Href string `json:"href"`

	// Specifies a media type as defined by the Internet Assigned Numbers Authority (IANA) Media Types Registry.
	MediaType string `json:"media-type,omitempty"`

	Hashes []FieldOscalMetadataHash `json:"hashes,omitempty"`
}

// AssertResourceLinkRequired checks if the required fields are not zero-ed
func AssertResourceLinkRequired(obj ResourceLink) error {
	elements := map[string]interface{}{
		"href": obj.Href,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Hashes {
		if err := AssertFieldOscalMetadataHashRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseResourceLinkRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ResourceLink (e.g. [][]ResourceLink), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseResourceLinkRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aResourceLink, ok := obj.(ResourceLink)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertResourceLinkRequired(aResourceLink)
	})
}
