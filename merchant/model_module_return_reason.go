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

// ModuleReturnReason the model 'ModuleReturnReason'
type ModuleReturnReason string

// List of ModuleReturnReason
const (
	MODULERETURNREASON_PRODUCT_DEFECT ModuleReturnReason = "PRODUCT_DEFECT"
	MODULERETURNREASON_PRODUCT_UNSATISFACTORY ModuleReturnReason = "PRODUCT_UNSATISFACTORY"
	MODULERETURNREASON_WRONG_PRODUCT ModuleReturnReason = "WRONG_PRODUCT"
	MODULERETURNREASON_TOO_MANY_PRODUCTS ModuleReturnReason = "TOO_MANY_PRODUCTS"
	MODULERETURNREASON_REFUSED ModuleReturnReason = "REFUSED"
	MODULERETURNREASON_REFUSED_DAMAGED ModuleReturnReason = "REFUSED_DAMAGED"
	MODULERETURNREASON_WRONG_ADDRESS ModuleReturnReason = "WRONG_ADDRESS"
	MODULERETURNREASON_NOT_COLLECTED ModuleReturnReason = "NOT_COLLECTED"
	MODULERETURNREASON_WRONG_SIZE ModuleReturnReason = "WRONG_SIZE"
	MODULERETURNREASON_OTHER ModuleReturnReason = "OTHER"
)

// All allowed values of ModuleReturnReason enum
var AllowedModuleReturnReasonEnumValues = []ModuleReturnReason{
	"PRODUCT_DEFECT",
	"PRODUCT_UNSATISFACTORY",
	"WRONG_PRODUCT",
	"TOO_MANY_PRODUCTS",
	"REFUSED",
	"REFUSED_DAMAGED",
	"WRONG_ADDRESS",
	"NOT_COLLECTED",
	"WRONG_SIZE",
	"OTHER",
}

func (v *ModuleReturnReason) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModuleReturnReason(value)
	for _, existing := range AllowedModuleReturnReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModuleReturnReason", value)
}

// NewModuleReturnReasonFromValue returns a pointer to a valid ModuleReturnReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModuleReturnReasonFromValue(v string) (*ModuleReturnReason, error) {
	ev := ModuleReturnReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModuleReturnReason: valid values are %v", v, AllowedModuleReturnReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModuleReturnReason) IsValid() bool {
	for _, existing := range AllowedModuleReturnReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ModuleReturnReason value
func (v ModuleReturnReason) Ptr() *ModuleReturnReason {
	return &v
}

type NullableModuleReturnReason struct {
	value *ModuleReturnReason
	isSet bool
}

func (v NullableModuleReturnReason) Get() *ModuleReturnReason {
	return v.value
}

func (v *NullableModuleReturnReason) Set(val *ModuleReturnReason) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleReturnReason) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleReturnReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleReturnReason(val *ModuleReturnReason) *NullableModuleReturnReason {
	return &NullableModuleReturnReason{value: val, isSet: true}
}

func (v NullableModuleReturnReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleReturnReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

