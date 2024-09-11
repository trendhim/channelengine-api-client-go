/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"encoding/json"
	"fmt"
)

// ReturnLineIdentifier the model 'ReturnLineIdentifier'
type ReturnLineIdentifier string

// List of ReturnLineIdentifier
const (
	RETURNLINEIDENTIFIER_RETURN_LINE_ID ReturnLineIdentifier = "RETURN_LINE_ID"
)

// All allowed values of ReturnLineIdentifier enum
var AllowedReturnLineIdentifierEnumValues = []ReturnLineIdentifier{
	"RETURN_LINE_ID",
}

func (v *ReturnLineIdentifier) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReturnLineIdentifier(value)
	for _, existing := range AllowedReturnLineIdentifierEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReturnLineIdentifier", value)
}

// NewReturnLineIdentifierFromValue returns a pointer to a valid ReturnLineIdentifier
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReturnLineIdentifierFromValue(v string) (*ReturnLineIdentifier, error) {
	ev := ReturnLineIdentifier(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReturnLineIdentifier: valid values are %v", v, AllowedReturnLineIdentifierEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReturnLineIdentifier) IsValid() bool {
	for _, existing := range AllowedReturnLineIdentifierEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReturnLineIdentifier value
func (v ReturnLineIdentifier) Ptr() *ReturnLineIdentifier {
	return &v
}

type NullableReturnLineIdentifier struct {
	value *ReturnLineIdentifier
	isSet bool
}

func (v NullableReturnLineIdentifier) Get() *ReturnLineIdentifier {
	return v.value
}

func (v *NullableReturnLineIdentifier) Set(val *ReturnLineIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnLineIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnLineIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnLineIdentifier(val *ReturnLineIdentifier) *NullableReturnLineIdentifier {
	return &NullableReturnLineIdentifier{value: val, isSet: true}
}

func (v NullableReturnLineIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnLineIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

