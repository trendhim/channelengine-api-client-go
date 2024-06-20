/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.14.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ProductAttributeGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductAttributeGroupRequest{}

// ProductAttributeGroupRequest struct for ProductAttributeGroupRequest
type ProductAttributeGroupRequest struct {
	GroupName string `json:"GroupName"`
	ProductExtraDataKeys []string `json:"ProductExtraDataKeys"`
}

type _ProductAttributeGroupRequest ProductAttributeGroupRequest

// NewProductAttributeGroupRequest instantiates a new ProductAttributeGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAttributeGroupRequest(groupName string, productExtraDataKeys []string) *ProductAttributeGroupRequest {
	this := ProductAttributeGroupRequest{}
	this.GroupName = groupName
	this.ProductExtraDataKeys = productExtraDataKeys
	return &this
}

// NewProductAttributeGroupRequestWithDefaults instantiates a new ProductAttributeGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAttributeGroupRequestWithDefaults() *ProductAttributeGroupRequest {
	this := ProductAttributeGroupRequest{}
	return &this
}

// GetGroupName returns the GroupName field value
func (o *ProductAttributeGroupRequest) GetGroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value
// and a boolean to check if the value has been set.
func (o *ProductAttributeGroupRequest) GetGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupName, true
}

// SetGroupName sets field value
func (o *ProductAttributeGroupRequest) SetGroupName(v string) {
	o.GroupName = v
}

// GetProductExtraDataKeys returns the ProductExtraDataKeys field value
func (o *ProductAttributeGroupRequest) GetProductExtraDataKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ProductExtraDataKeys
}

// GetProductExtraDataKeysOk returns a tuple with the ProductExtraDataKeys field value
// and a boolean to check if the value has been set.
func (o *ProductAttributeGroupRequest) GetProductExtraDataKeysOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductExtraDataKeys, true
}

// SetProductExtraDataKeys sets field value
func (o *ProductAttributeGroupRequest) SetProductExtraDataKeys(v []string) {
	o.ProductExtraDataKeys = v
}

func (o ProductAttributeGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductAttributeGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["GroupName"] = o.GroupName
	toSerialize["ProductExtraDataKeys"] = o.ProductExtraDataKeys
	return toSerialize, nil
}

func (o *ProductAttributeGroupRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"GroupName",
		"ProductExtraDataKeys",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varProductAttributeGroupRequest := _ProductAttributeGroupRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProductAttributeGroupRequest)

	if err != nil {
		return err
	}

	*o = ProductAttributeGroupRequest(varProductAttributeGroupRequest)

	return err
}

type NullableProductAttributeGroupRequest struct {
	value *ProductAttributeGroupRequest
	isSet bool
}

func (v NullableProductAttributeGroupRequest) Get() *ProductAttributeGroupRequest {
	return v.value
}

func (v *NullableProductAttributeGroupRequest) Set(val *ProductAttributeGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAttributeGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAttributeGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAttributeGroupRequest(val *ProductAttributeGroupRequest) *NullableProductAttributeGroupRequest {
	return &NullableProductAttributeGroupRequest{value: val, isSet: true}
}

func (v NullableProductAttributeGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductAttributeGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


