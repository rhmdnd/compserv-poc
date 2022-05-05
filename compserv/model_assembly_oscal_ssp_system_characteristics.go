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

// AssemblyOscalSspSystemCharacteristics - Contains the characteristics of the system, such as its name, purpose, and security impact level.
type AssemblyOscalSspSystemCharacteristics struct {

	SystemIds []FieldOscalImplementationCommonSystemId `json:"system-ids"`

	// The full name of the system.
	SystemName string `json:"system-name"`

	// A short name for the system, such as an acronym, that is suitable for display in a data table or summary list.
	SystemNameShort string `json:"system-name-short,omitempty"`

	// A summary of the system.
	Description string `json:"description"`

	Props []AssemblyOscalMetadataProperty3 `json:"props,omitempty"`

	Links []AssemblyOscalMetadataLink `json:"links,omitempty"`

	// The date the system received its authorization.
	DateAuthorized string `json:"date-authorized,omitempty"`

	// The overall information system sensitivity categorization, such as defined by FIPS-199.
	SecuritySensitivityLevel string `json:"security-sensitivity-level"`

	SystemInformation AssemblyOscalSspSystemInformation `json:"system-information"`

	SecurityImpactLevel AssemblyOscalSspSecurityImpactLevel `json:"security-impact-level"`

	Status AssemblyOscalSspStatus `json:"status"`

	AuthorizationBoundary AssemblyOscalSspAuthorizationBoundary `json:"authorization-boundary"`

	NetworkArchitecture AssemblyOscalSspNetworkArchitecture `json:"network-architecture,omitempty"`

	DataFlow AssemblyOscalSspDataFlow `json:"data-flow,omitempty"`

	ResponsibleParties []AssemblyOscalMetadataResponsibleParty3 `json:"responsible-parties,omitempty"`

	// Additional commentary on the containing object.
	Remarks string `json:"remarks,omitempty"`
}

// AssertAssemblyOscalSspSystemCharacteristicsRequired checks if the required fields are not zero-ed
func AssertAssemblyOscalSspSystemCharacteristicsRequired(obj AssemblyOscalSspSystemCharacteristics) error {
	elements := map[string]interface{}{
		"system-ids": obj.SystemIds,
		"system-name": obj.SystemName,
		"description": obj.Description,
		"security-sensitivity-level": obj.SecuritySensitivityLevel,
		"system-information": obj.SystemInformation,
		"security-impact-level": obj.SecurityImpactLevel,
		"status": obj.Status,
		"authorization-boundary": obj.AuthorizationBoundary,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.SystemIds {
		if err := AssertFieldOscalImplementationCommonSystemIdRequired(el); err != nil {
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
	if err := AssertAssemblyOscalSspSystemInformationRequired(obj.SystemInformation); err != nil {
		return err
	}
	if err := AssertAssemblyOscalSspSecurityImpactLevelRequired(obj.SecurityImpactLevel); err != nil {
		return err
	}
	if err := AssertAssemblyOscalSspStatusRequired(obj.Status); err != nil {
		return err
	}
	if err := AssertAssemblyOscalSspAuthorizationBoundaryRequired(obj.AuthorizationBoundary); err != nil {
		return err
	}
	if err := AssertAssemblyOscalSspNetworkArchitectureRequired(obj.NetworkArchitecture); err != nil {
		return err
	}
	if err := AssertAssemblyOscalSspDataFlowRequired(obj.DataFlow); err != nil {
		return err
	}
	for _, el := range obj.ResponsibleParties {
		if err := AssertAssemblyOscalMetadataResponsibleParty3Required(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseAssemblyOscalSspSystemCharacteristicsRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of AssemblyOscalSspSystemCharacteristics (e.g. [][]AssemblyOscalSspSystemCharacteristics), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseAssemblyOscalSspSystemCharacteristicsRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aAssemblyOscalSspSystemCharacteristics, ok := obj.(AssemblyOscalSspSystemCharacteristics)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertAssemblyOscalSspSystemCharacteristicsRequired(aAssemblyOscalSspSystemCharacteristics)
	})
}