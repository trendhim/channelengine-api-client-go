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

// DataChangesProductType the model 'DataChangesProductType'
type DataChangesProductType string

// List of DataChangesProductType
const (
	DATACHANGESPRODUCTTYPE_SINGLE DataChangesProductType = "SINGLE"
	DATACHANGESPRODUCTTYPE_PARENT DataChangesProductType = "PARENT"
	DATACHANGESPRODUCTTYPE_CHILD DataChangesProductType = "CHILD"
	DATACHANGESPRODUCTTYPE_GRANDPARENT DataChangesProductType = "GRANDPARENT"
)

// All allowed values of DataChangesProductType enum
var AllowedDataChangesProductTypeEnumValues = []DataChangesProductType{
	"SINGLE",
	"PARENT",
	"CHILD",
	"GRANDPARENT",
}

func (v *DataChangesProductType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataChangesProductType(value)
	for _, existing := range AllowedDataChangesProductTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataChangesProductType", value)
}

// NewDataChangesProductTypeFromValue returns a pointer to a valid DataChangesProductType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataChangesProductTypeFromValue(v string) (*DataChangesProductType, error) {
	ev := DataChangesProductType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataChangesProductType: valid values are %v", v, AllowedDataChangesProductTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataChangesProductType) IsValid() bool {
	for _, existing := range AllowedDataChangesProductTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataChangesProductType value
func (v DataChangesProductType) Ptr() *DataChangesProductType {
	return &v
}

type NullableDataChangesProductType struct {
	value *DataChangesProductType
	isSet bool
}

func (v NullableDataChangesProductType) Get() *DataChangesProductType {
	return v.value
}

func (v *NullableDataChangesProductType) Set(val *DataChangesProductType) {
	v.value = val
	v.isSet = true
}

func (v NullableDataChangesProductType) IsSet() bool {
	return v.isSet
}

func (v *NullableDataChangesProductType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataChangesProductType(val *DataChangesProductType) *NullableDataChangesProductType {
	return &NullableDataChangesProductType{value: val, isSet: true}
}

func (v NullableDataChangesProductType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataChangesProductType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

