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

// ListedProductExportStatus the model 'ListedProductExportStatus'
type ListedProductExportStatus string

// List of ListedProductExportStatus
const (
	LISTEDPRODUCTEXPORTSTATUS_CREATED ListedProductExportStatus = "CREATED"
	LISTEDPRODUCTEXPORTSTATUS_UPDATED ListedProductExportStatus = "UPDATED"
	LISTEDPRODUCTEXPORTSTATUS_DELETED ListedProductExportStatus = "DELETED"
	LISTEDPRODUCTEXPORTSTATUS_CREATE_FAILED ListedProductExportStatus = "CREATE_FAILED"
)

// All allowed values of ListedProductExportStatus enum
var AllowedListedProductExportStatusEnumValues = []ListedProductExportStatus{
	"CREATED",
	"UPDATED",
	"DELETED",
	"CREATE_FAILED",
}

func (v *ListedProductExportStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ListedProductExportStatus(value)
	for _, existing := range AllowedListedProductExportStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ListedProductExportStatus", value)
}

// NewListedProductExportStatusFromValue returns a pointer to a valid ListedProductExportStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewListedProductExportStatusFromValue(v string) (*ListedProductExportStatus, error) {
	ev := ListedProductExportStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ListedProductExportStatus: valid values are %v", v, AllowedListedProductExportStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ListedProductExportStatus) IsValid() bool {
	for _, existing := range AllowedListedProductExportStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListedProductExportStatus value
func (v ListedProductExportStatus) Ptr() *ListedProductExportStatus {
	return &v
}

type NullableListedProductExportStatus struct {
	value *ListedProductExportStatus
	isSet bool
}

func (v NullableListedProductExportStatus) Get() *ListedProductExportStatus {
	return v.value
}

func (v *NullableListedProductExportStatus) Set(val *ListedProductExportStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableListedProductExportStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableListedProductExportStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListedProductExportStatus(val *ListedProductExportStatus) *NullableListedProductExportStatus {
	return &NullableListedProductExportStatus{value: val, isSet: true}
}

func (v NullableListedProductExportStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListedProductExportStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

