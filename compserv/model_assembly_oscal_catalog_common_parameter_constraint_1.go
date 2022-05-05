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

// AssemblyOscalCatalogCommonParameterConstraint1 - A formal or informal expression of a constraint or test
type AssemblyOscalCatalogCommonParameterConstraint1 struct {

	// A textual summary of the constraint to be applied.
	Description string `json:"description,omitempty"`

	Tests []ConstraintTest `json:"tests,omitempty"`
}

// AssertAssemblyOscalCatalogCommonParameterConstraint1Required checks if the required fields are not zero-ed
func AssertAssemblyOscalCatalogCommonParameterConstraint1Required(obj AssemblyOscalCatalogCommonParameterConstraint1) error {
	for _, el := range obj.Tests {
		if err := AssertConstraintTestRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseAssemblyOscalCatalogCommonParameterConstraint1Required recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of AssemblyOscalCatalogCommonParameterConstraint1 (e.g. [][]AssemblyOscalCatalogCommonParameterConstraint1), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseAssemblyOscalCatalogCommonParameterConstraint1Required(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aAssemblyOscalCatalogCommonParameterConstraint1, ok := obj.(AssemblyOscalCatalogCommonParameterConstraint1)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertAssemblyOscalCatalogCommonParameterConstraint1Required(aAssemblyOscalCatalogCommonParameterConstraint1)
	})
}
