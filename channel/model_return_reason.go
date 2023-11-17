/*
ChannelEngine Channel API

ChannelEngine API for channels

API version: 2.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package channel

import (
	"encoding/json"
	"fmt"
)

// ReturnReason the model 'ReturnReason'
type ReturnReason string

// List of ReturnReason
const (
	RETURNREASON_PRODUCT_DEFECT ReturnReason = "PRODUCT_DEFECT"
	RETURNREASON_PRODUCT_UNSATISFACTORY ReturnReason = "PRODUCT_UNSATISFACTORY"
	RETURNREASON_WRONG_PRODUCT ReturnReason = "WRONG_PRODUCT"
	RETURNREASON_TOO_MANY_PRODUCTS ReturnReason = "TOO_MANY_PRODUCTS"
	RETURNREASON_REFUSED ReturnReason = "REFUSED"
	RETURNREASON_REFUSED_DAMAGED ReturnReason = "REFUSED_DAMAGED"
	RETURNREASON_WRONG_ADDRESS ReturnReason = "WRONG_ADDRESS"
	RETURNREASON_NOT_COLLECTED ReturnReason = "NOT_COLLECTED"
	RETURNREASON_WRONG_SIZE ReturnReason = "WRONG_SIZE"
	RETURNREASON_OTHER ReturnReason = "OTHER"
)

// All allowed values of ReturnReason enum
var AllowedReturnReasonEnumValues = []ReturnReason{
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

func (v *ReturnReason) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReturnReason(value)
	for _, existing := range AllowedReturnReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReturnReason", value)
}

// NewReturnReasonFromValue returns a pointer to a valid ReturnReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReturnReasonFromValue(v string) (*ReturnReason, error) {
	ev := ReturnReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReturnReason: valid values are %v", v, AllowedReturnReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReturnReason) IsValid() bool {
	for _, existing := range AllowedReturnReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReturnReason value
func (v ReturnReason) Ptr() *ReturnReason {
	return &v
}

type NullableReturnReason struct {
	value *ReturnReason
	isSet bool
}

func (v NullableReturnReason) Get() *ReturnReason {
	return v.value
}

func (v *NullableReturnReason) Set(val *ReturnReason) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnReason) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnReason(val *ReturnReason) *NullableReturnReason {
	return &NullableReturnReason{value: val, isSet: true}
}

func (v NullableReturnReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
