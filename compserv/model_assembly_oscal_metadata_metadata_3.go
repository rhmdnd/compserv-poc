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
	"time"
)

// AssemblyOscalMetadataMetadata3 - Provides information about the publication and availability of the containing document.
type AssemblyOscalMetadataMetadata3 struct {

	// A name given to the document, which may be used by a tool for display and navigation.
	Title string `json:"title"`

	// The date and time the document was published. The date-time value must be formatted according to RFC 3339 with full time and time zone included.
	Published time.Time `json:"published,omitempty"`

	// The date and time the document was last modified. The date-time value must be formatted according to RFC 3339 with full time and time zone included.
	LastModified time.Time `json:"last-modified"`

	// A string used to distinguish the current version of the document from other previous (and future) versions.
	Version string `json:"version"`

	// The OSCAL model version the document was authored against.
	OscalVersion string `json:"oscal-version"`

	Revisions []AssemblyOscalMetadataRevision3 `json:"revisions,omitempty"`

	DocumentIds []FieldOscalMetadataDocumentId `json:"document-ids,omitempty"`

	Props []AssemblyOscalMetadataProperty3 `json:"props,omitempty"`

	Links []AssemblyOscalMetadataLink `json:"links,omitempty"`

	Roles []AssemblyOscalMetadataRole3 `json:"roles,omitempty"`

	Locations []AssemblyOscalMetadataLocation3 `json:"locations,omitempty"`

	Parties []AssemblyOscalMetadataParty3 `json:"parties,omitempty"`

	ResponsibleParties []AssemblyOscalMetadataResponsibleParty3 `json:"responsible-parties,omitempty"`

	// Additional commentary on the containing object.
	Remarks string `json:"remarks,omitempty"`
}

// AssertAssemblyOscalMetadataMetadata3Required checks if the required fields are not zero-ed
func AssertAssemblyOscalMetadataMetadata3Required(obj AssemblyOscalMetadataMetadata3) error {
	elements := map[string]interface{}{
		"title": obj.Title,
		"last-modified": obj.LastModified,
		"version": obj.Version,
		"oscal-version": obj.OscalVersion,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Revisions {
		if err := AssertAssemblyOscalMetadataRevision3Required(el); err != nil {
			return err
		}
	}
	for _, el := range obj.DocumentIds {
		if err := AssertFieldOscalMetadataDocumentIdRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Props {
		if err := AssertAssemblyOscalMetadataProperty3Required(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Links {
		if err := AssertAssemblyOscalMetadataLinkRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Roles {
		if err := AssertAssemblyOscalMetadataRole3Required(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Locations {
		if err := AssertAssemblyOscalMetadataLocation3Required(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Parties {
		if err := AssertAssemblyOscalMetadataParty3Required(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ResponsibleParties {
		if err := AssertAssemblyOscalMetadataResponsibleParty3Required(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseAssemblyOscalMetadataMetadata3Required recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of AssemblyOscalMetadataMetadata3 (e.g. [][]AssemblyOscalMetadataMetadata3), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseAssemblyOscalMetadataMetadata3Required(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aAssemblyOscalMetadataMetadata3, ok := obj.(AssemblyOscalMetadataMetadata3)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertAssemblyOscalMetadataMetadata3Required(aAssemblyOscalMetadataMetadata3)
	})
}
