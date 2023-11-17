/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"encoding/json"
	"fmt"
)

// MancoReason the model 'MancoReason'
type MancoReason string

// List of MancoReason
const (
	MANCOREASON_NOT_IN_STOCK MancoReason = "NOT_IN_STOCK"
	MANCOREASON_DAMAGED MancoReason = "DAMAGED"
	MANCOREASON_INCOMPLETE MancoReason = "INCOMPLETE"
	MANCOREASON_CLIENT_CANCELLED MancoReason = "CLIENT_CANCELLED"
	MANCOREASON_INVALID_ADDRESS MancoReason = "INVALID_ADDRESS"
	MANCOREASON_OTHER MancoReason = "OTHER"
)

// All allowed values of MancoReason enum
var AllowedMancoReasonEnumValues = []MancoReason{
	"NOT_IN_STOCK",
	"DAMAGED",
	"INCOMPLETE",
	"CLIENT_CANCELLED",
	"INVALID_ADDRESS",
	"OTHER",
}

func (v *MancoReason) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MancoReason(value)
	for _, existing := range AllowedMancoReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MancoReason", value)
}

// NewMancoReasonFromValue returns a pointer to a valid MancoReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMancoReasonFromValue(v string) (*MancoReason, error) {
	ev := MancoReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MancoReason: valid values are %v", v, AllowedMancoReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MancoReason) IsValid() bool {
	for _, existing := range AllowedMancoReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MancoReason value
func (v MancoReason) Ptr() *MancoReason {
	return &v
}

type NullableMancoReason struct {
	value *MancoReason
	isSet bool
}

func (v NullableMancoReason) Get() *MancoReason {
	return v.value
}

func (v *NullableMancoReason) Set(val *MancoReason) {
	v.value = val
	v.isSet = true
}

func (v NullableMancoReason) IsSet() bool {
	return v.isSet
}

func (v *NullableMancoReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMancoReason(val *MancoReason) *NullableMancoReason {
	return &NullableMancoReason{value: val, isSet: true}
}

func (v NullableMancoReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMancoReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

