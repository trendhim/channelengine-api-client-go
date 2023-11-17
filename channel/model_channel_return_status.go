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

// ChannelReturnStatus the model 'ChannelReturnStatus'
type ChannelReturnStatus string

// List of ChannelReturnStatus
const (
	CHANNELRETURNSTATUS_IN_PROGRESS ChannelReturnStatus = "IN_PROGRESS"
	CHANNELRETURNSTATUS_RECEIVED ChannelReturnStatus = "RECEIVED"
)

// All allowed values of ChannelReturnStatus enum
var AllowedChannelReturnStatusEnumValues = []ChannelReturnStatus{
	"IN_PROGRESS",
	"RECEIVED",
}

func (v *ChannelReturnStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ChannelReturnStatus(value)
	for _, existing := range AllowedChannelReturnStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ChannelReturnStatus", value)
}

// NewChannelReturnStatusFromValue returns a pointer to a valid ChannelReturnStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChannelReturnStatusFromValue(v string) (*ChannelReturnStatus, error) {
	ev := ChannelReturnStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChannelReturnStatus: valid values are %v", v, AllowedChannelReturnStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChannelReturnStatus) IsValid() bool {
	for _, existing := range AllowedChannelReturnStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChannelReturnStatus value
func (v ChannelReturnStatus) Ptr() *ChannelReturnStatus {
	return &v
}

type NullableChannelReturnStatus struct {
	value *ChannelReturnStatus
	isSet bool
}

func (v NullableChannelReturnStatus) Get() *ChannelReturnStatus {
	return v.value
}

func (v *NullableChannelReturnStatus) Set(val *ChannelReturnStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelReturnStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelReturnStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelReturnStatus(val *ChannelReturnStatus) *NullableChannelReturnStatus {
	return &NullableChannelReturnStatus{value: val, isSet: true}
}

func (v NullableChannelReturnStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelReturnStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

