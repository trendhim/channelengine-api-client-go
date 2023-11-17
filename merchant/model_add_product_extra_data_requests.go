/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"encoding/json"
)

// checks if the AddProductExtraDataRequests type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddProductExtraDataRequests{}

// AddProductExtraDataRequests struct for AddProductExtraDataRequests
type AddProductExtraDataRequests struct {
	ProductExtraDataKeys []string `json:"ProductExtraDataKeys,omitempty"`
}

// NewAddProductExtraDataRequests instantiates a new AddProductExtraDataRequests object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddProductExtraDataRequests() *AddProductExtraDataRequests {
	this := AddProductExtraDataRequests{}
	return &this
}

// NewAddProductExtraDataRequestsWithDefaults instantiates a new AddProductExtraDataRequests object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddProductExtraDataRequestsWithDefaults() *AddProductExtraDataRequests {
	this := AddProductExtraDataRequests{}
	return &this
}

// GetProductExtraDataKeys returns the ProductExtraDataKeys field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddProductExtraDataRequests) GetProductExtraDataKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProductExtraDataKeys
}

// GetProductExtraDataKeysOk returns a tuple with the ProductExtraDataKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddProductExtraDataRequests) GetProductExtraDataKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductExtraDataKeys) {
		return nil, false
	}
	return o.ProductExtraDataKeys, true
}

// HasProductExtraDataKeys returns a boolean if a field has been set.
func (o *AddProductExtraDataRequests) HasProductExtraDataKeys() bool {
	if o != nil && IsNil(o.ProductExtraDataKeys) {
		return true
	}

	return false
}

// SetProductExtraDataKeys gets a reference to the given []string and assigns it to the ProductExtraDataKeys field.
func (o *AddProductExtraDataRequests) SetProductExtraDataKeys(v []string) {
	o.ProductExtraDataKeys = v
}

func (o AddProductExtraDataRequests) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddProductExtraDataRequests) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ProductExtraDataKeys != nil {
		toSerialize["ProductExtraDataKeys"] = o.ProductExtraDataKeys
	}
	return toSerialize, nil
}

type NullableAddProductExtraDataRequests struct {
	value *AddProductExtraDataRequests
	isSet bool
}

func (v NullableAddProductExtraDataRequests) Get() *AddProductExtraDataRequests {
	return v.value
}

func (v *NullableAddProductExtraDataRequests) Set(val *AddProductExtraDataRequests) {
	v.value = val
	v.isSet = true
}

func (v NullableAddProductExtraDataRequests) IsSet() bool {
	return v.isSet
}

func (v *NullableAddProductExtraDataRequests) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddProductExtraDataRequests(val *AddProductExtraDataRequests) *NullableAddProductExtraDataRequests {
	return &NullableAddProductExtraDataRequests{value: val, isSet: true}
}

func (v NullableAddProductExtraDataRequests) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddProductExtraDataRequests) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

