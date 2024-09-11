/*
ChannelEngine Channel API

ChannelEngine API for channels

API version: 2.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package channel

import (
	"encoding/json"
	"fmt"
)

// CreatedByType the model 'CreatedByType'
type CreatedByType string

// List of CreatedByType
const (
	CREATEDBYTYPE_MERCHANT CreatedByType = "MERCHANT"
	CREATEDBYTYPE_CHANNEL CreatedByType = "CHANNEL"
)

// All allowed values of CreatedByType enum
var AllowedCreatedByTypeEnumValues = []CreatedByType{
	"MERCHANT",
	"CHANNEL",
}

func (v *CreatedByType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CreatedByType(value)
	for _, existing := range AllowedCreatedByTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CreatedByType", value)
}

// NewCreatedByTypeFromValue returns a pointer to a valid CreatedByType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreatedByTypeFromValue(v string) (*CreatedByType, error) {
	ev := CreatedByType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreatedByType: valid values are %v", v, AllowedCreatedByTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreatedByType) IsValid() bool {
	for _, existing := range AllowedCreatedByTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreatedByType value
func (v CreatedByType) Ptr() *CreatedByType {
	return &v
}

type NullableCreatedByType struct {
	value *CreatedByType
	isSet bool
}

func (v NullableCreatedByType) Get() *CreatedByType {
	return v.value
}

func (v *NullableCreatedByType) Set(val *CreatedByType) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatedByType) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatedByType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatedByType(val *CreatedByType) *NullableCreatedByType {
	return &NullableCreatedByType{value: val, isSet: true}
}

func (v NullableCreatedByType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatedByType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

