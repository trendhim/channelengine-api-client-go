/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"encoding/json"
)

// checks if the DeleteTargetView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteTargetView{}

// DeleteTargetView struct for DeleteTargetView
type DeleteTargetView struct {
	Month *int32 `json:"Month,omitempty"`
	Year *int32 `json:"Year,omitempty"`
}

// NewDeleteTargetView instantiates a new DeleteTargetView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteTargetView() *DeleteTargetView {
	this := DeleteTargetView{}
	return &this
}

// NewDeleteTargetViewWithDefaults instantiates a new DeleteTargetView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteTargetViewWithDefaults() *DeleteTargetView {
	this := DeleteTargetView{}
	return &this
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *DeleteTargetView) GetMonth() int32 {
	if o == nil || IsNil(o.Month) {
		var ret int32
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteTargetView) GetMonthOk() (*int32, bool) {
	if o == nil || IsNil(o.Month) {
		return nil, false
	}
	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *DeleteTargetView) HasMonth() bool {
	if o != nil && !IsNil(o.Month) {
		return true
	}

	return false
}

// SetMonth gets a reference to the given int32 and assigns it to the Month field.
func (o *DeleteTargetView) SetMonth(v int32) {
	o.Month = &v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *DeleteTargetView) GetYear() int32 {
	if o == nil || IsNil(o.Year) {
		var ret int32
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteTargetView) GetYearOk() (*int32, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *DeleteTargetView) HasYear() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given int32 and assigns it to the Year field.
func (o *DeleteTargetView) SetYear(v int32) {
	o.Year = &v
}

func (o DeleteTargetView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteTargetView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Month) {
		toSerialize["Month"] = o.Month
	}
	if !IsNil(o.Year) {
		toSerialize["Year"] = o.Year
	}
	return toSerialize, nil
}

type NullableDeleteTargetView struct {
	value *DeleteTargetView
	isSet bool
}

func (v NullableDeleteTargetView) Get() *DeleteTargetView {
	return v.value
}

func (v *NullableDeleteTargetView) Set(val *DeleteTargetView) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteTargetView) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteTargetView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteTargetView(val *DeleteTargetView) *NullableDeleteTargetView {
	return &NullableDeleteTargetView{value: val, isSet: true}
}

func (v NullableDeleteTargetView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteTargetView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


