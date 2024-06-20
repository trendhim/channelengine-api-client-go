/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.14.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"encoding/json"
	"fmt"
)

// ModulesTaxRegistrationType the model 'ModulesTaxRegistrationType'
type ModulesTaxRegistrationType string

// List of ModulesTaxRegistrationType
const (
	MODULESTAXREGISTRATIONTYPE_VAT ModulesTaxRegistrationType = "VAT"
	MODULESTAXREGISTRATIONTYPE_GST ModulesTaxRegistrationType = "GST"
)

// All allowed values of ModulesTaxRegistrationType enum
var AllowedModulesTaxRegistrationTypeEnumValues = []ModulesTaxRegistrationType{
	"VAT",
	"GST",
}

func (v *ModulesTaxRegistrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModulesTaxRegistrationType(value)
	for _, existing := range AllowedModulesTaxRegistrationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModulesTaxRegistrationType", value)
}

// NewModulesTaxRegistrationTypeFromValue returns a pointer to a valid ModulesTaxRegistrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModulesTaxRegistrationTypeFromValue(v string) (*ModulesTaxRegistrationType, error) {
	ev := ModulesTaxRegistrationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModulesTaxRegistrationType: valid values are %v", v, AllowedModulesTaxRegistrationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModulesTaxRegistrationType) IsValid() bool {
	for _, existing := range AllowedModulesTaxRegistrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ModulesTaxRegistrationType value
func (v ModulesTaxRegistrationType) Ptr() *ModulesTaxRegistrationType {
	return &v
}

type NullableModulesTaxRegistrationType struct {
	value *ModulesTaxRegistrationType
	isSet bool
}

func (v NullableModulesTaxRegistrationType) Get() *ModulesTaxRegistrationType {
	return v.value
}

func (v *NullableModulesTaxRegistrationType) Set(val *ModulesTaxRegistrationType) {
	v.value = val
	v.isSet = true
}

func (v NullableModulesTaxRegistrationType) IsSet() bool {
	return v.isSet
}

func (v *NullableModulesTaxRegistrationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModulesTaxRegistrationType(val *ModulesTaxRegistrationType) *NullableModulesTaxRegistrationType {
	return &NullableModulesTaxRegistrationType{value: val, isSet: true}
}

func (v NullableModulesTaxRegistrationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModulesTaxRegistrationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

