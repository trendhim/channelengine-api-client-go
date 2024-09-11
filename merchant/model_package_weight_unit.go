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

// PackageWeightUnit the model 'PackageWeightUnit'
type PackageWeightUnit string

// List of PackageWeightUnit
const (
	PACKAGEWEIGHTUNIT_GRAM PackageWeightUnit = "GRAM"
	PACKAGEWEIGHTUNIT_OUNCE PackageWeightUnit = "OUNCE"
)

// All allowed values of PackageWeightUnit enum
var AllowedPackageWeightUnitEnumValues = []PackageWeightUnit{
	"GRAM",
	"OUNCE",
}

func (v *PackageWeightUnit) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PackageWeightUnit(value)
	for _, existing := range AllowedPackageWeightUnitEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PackageWeightUnit", value)
}

// NewPackageWeightUnitFromValue returns a pointer to a valid PackageWeightUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPackageWeightUnitFromValue(v string) (*PackageWeightUnit, error) {
	ev := PackageWeightUnit(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PackageWeightUnit: valid values are %v", v, AllowedPackageWeightUnitEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PackageWeightUnit) IsValid() bool {
	for _, existing := range AllowedPackageWeightUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PackageWeightUnit value
func (v PackageWeightUnit) Ptr() *PackageWeightUnit {
	return &v
}

type NullablePackageWeightUnit struct {
	value *PackageWeightUnit
	isSet bool
}

func (v NullablePackageWeightUnit) Get() *PackageWeightUnit {
	return v.value
}

func (v *NullablePackageWeightUnit) Set(val *PackageWeightUnit) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageWeightUnit) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageWeightUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageWeightUnit(val *PackageWeightUnit) *NullablePackageWeightUnit {
	return &NullablePackageWeightUnit{value: val, isSet: true}
}

func (v NullablePackageWeightUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageWeightUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

