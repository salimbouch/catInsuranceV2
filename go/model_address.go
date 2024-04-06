/*
 * Cat Insurance API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi


import (
	"errors"
)



type Address struct {

	Street string `json:"street"`

	HouseNumber string `json:"houseNumber"`

	ZipCode float32 `json:"zipCode"`

	City string `json:"city"`

	Id string `json:"id"`
}

// AssertAddressRequired checks if the required fields are not zero-ed
func AssertAddressRequired(obj Address) error {
	elements := map[string]interface{}{
		"street": obj.Street,
		"houseNumber": obj.HouseNumber,
		"zipCode": obj.ZipCode,
		"city": obj.City,
		"id": obj.Id,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertAddressConstraints checks if the values respects the defined constraints
func AssertAddressConstraints(obj Address) error {
	if obj.ZipCode < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.ZipCode > 99999 {
		return &ParsingError{Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
