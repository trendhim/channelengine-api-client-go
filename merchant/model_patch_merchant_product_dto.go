/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.14.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"encoding/json"
)

// checks if the PatchMerchantProductDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchMerchantProductDto{}

// PatchMerchantProductDto struct for PatchMerchantProductDto
type PatchMerchantProductDto struct {
	// Fields to update
	PropertiesToUpdate []string `json:"PropertiesToUpdate,omitempty"`
	// Products to be updated
	MerchantProductRequestModels []MerchantProductRequest `json:"MerchantProductRequestModels,omitempty"`
}

// NewPatchMerchantProductDto instantiates a new PatchMerchantProductDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchMerchantProductDto() *PatchMerchantProductDto {
	this := PatchMerchantProductDto{}
	return &this
}

// NewPatchMerchantProductDtoWithDefaults instantiates a new PatchMerchantProductDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchMerchantProductDtoWithDefaults() *PatchMerchantProductDto {
	this := PatchMerchantProductDto{}
	return &this
}

// GetPropertiesToUpdate returns the PropertiesToUpdate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchMerchantProductDto) GetPropertiesToUpdate() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PropertiesToUpdate
}

// GetPropertiesToUpdateOk returns a tuple with the PropertiesToUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchMerchantProductDto) GetPropertiesToUpdateOk() ([]string, bool) {
	if o == nil || IsNil(o.PropertiesToUpdate) {
		return nil, false
	}
	return o.PropertiesToUpdate, true
}

// HasPropertiesToUpdate returns a boolean if a field has been set.
func (o *PatchMerchantProductDto) HasPropertiesToUpdate() bool {
	if o != nil && !IsNil(o.PropertiesToUpdate) {
		return true
	}

	return false
}

// SetPropertiesToUpdate gets a reference to the given []string and assigns it to the PropertiesToUpdate field.
func (o *PatchMerchantProductDto) SetPropertiesToUpdate(v []string) {
	o.PropertiesToUpdate = v
}

// GetMerchantProductRequestModels returns the MerchantProductRequestModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchMerchantProductDto) GetMerchantProductRequestModels() []MerchantProductRequest {
	if o == nil {
		var ret []MerchantProductRequest
		return ret
	}
	return o.MerchantProductRequestModels
}

// GetMerchantProductRequestModelsOk returns a tuple with the MerchantProductRequestModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchMerchantProductDto) GetMerchantProductRequestModelsOk() ([]MerchantProductRequest, bool) {
	if o == nil || IsNil(o.MerchantProductRequestModels) {
		return nil, false
	}
	return o.MerchantProductRequestModels, true
}

// HasMerchantProductRequestModels returns a boolean if a field has been set.
func (o *PatchMerchantProductDto) HasMerchantProductRequestModels() bool {
	if o != nil && !IsNil(o.MerchantProductRequestModels) {
		return true
	}

	return false
}

// SetMerchantProductRequestModels gets a reference to the given []MerchantProductRequest and assigns it to the MerchantProductRequestModels field.
func (o *PatchMerchantProductDto) SetMerchantProductRequestModels(v []MerchantProductRequest) {
	o.MerchantProductRequestModels = v
}

func (o PatchMerchantProductDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchMerchantProductDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PropertiesToUpdate != nil {
		toSerialize["PropertiesToUpdate"] = o.PropertiesToUpdate
	}
	if o.MerchantProductRequestModels != nil {
		toSerialize["MerchantProductRequestModels"] = o.MerchantProductRequestModels
	}
	return toSerialize, nil
}

type NullablePatchMerchantProductDto struct {
	value *PatchMerchantProductDto
	isSet bool
}

func (v NullablePatchMerchantProductDto) Get() *PatchMerchantProductDto {
	return v.value
}

func (v *NullablePatchMerchantProductDto) Set(val *PatchMerchantProductDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchMerchantProductDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchMerchantProductDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchMerchantProductDto(val *PatchMerchantProductDto) *NullablePatchMerchantProductDto {
	return &NullablePatchMerchantProductDto{value: val, isSet: true}
}

func (v NullablePatchMerchantProductDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchMerchantProductDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


