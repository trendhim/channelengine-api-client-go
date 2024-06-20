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

// ShipmentLineStatus the model 'ShipmentLineStatus'
type ShipmentLineStatus string

// List of ShipmentLineStatus
const (
	SHIPMENTLINESTATUS_SHIPPED ShipmentLineStatus = "SHIPPED"
	SHIPMENTLINESTATUS_IN_BACKORDER ShipmentLineStatus = "IN_BACKORDER"
	SHIPMENTLINESTATUS_MANCO ShipmentLineStatus = "MANCO"
)

// All allowed values of ShipmentLineStatus enum
var AllowedShipmentLineStatusEnumValues = []ShipmentLineStatus{
	"SHIPPED",
	"IN_BACKORDER",
	"MANCO",
}

func (v *ShipmentLineStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ShipmentLineStatus(value)
	for _, existing := range AllowedShipmentLineStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ShipmentLineStatus", value)
}

// NewShipmentLineStatusFromValue returns a pointer to a valid ShipmentLineStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewShipmentLineStatusFromValue(v string) (*ShipmentLineStatus, error) {
	ev := ShipmentLineStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ShipmentLineStatus: valid values are %v", v, AllowedShipmentLineStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ShipmentLineStatus) IsValid() bool {
	for _, existing := range AllowedShipmentLineStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ShipmentLineStatus value
func (v ShipmentLineStatus) Ptr() *ShipmentLineStatus {
	return &v
}

type NullableShipmentLineStatus struct {
	value *ShipmentLineStatus
	isSet bool
}

func (v NullableShipmentLineStatus) Get() *ShipmentLineStatus {
	return v.value
}

func (v *NullableShipmentLineStatus) Set(val *ShipmentLineStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentLineStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentLineStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentLineStatus(val *ShipmentLineStatus) *NullableShipmentLineStatus {
	return &NullableShipmentLineStatus{value: val, isSet: true}
}

func (v NullableShipmentLineStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentLineStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

