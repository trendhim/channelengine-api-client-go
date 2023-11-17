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

// checks if the MerchantReturnLineRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantReturnLineRequest{}

// MerchantReturnLineRequest struct for MerchantReturnLineRequest
type MerchantReturnLineRequest struct {
	// The unique product reference used by the Merchant (sku).
	MerchantProductNo string `json:"MerchantProductNo"`
	// Number of items of the product in this return.
	Quantity int32 `json:"Quantity"`
	// Extra data on the returnline. Each item must have an unqiue key
	ExtraData map[string]string `json:"ExtraData,omitempty"`
}

type _MerchantReturnLineRequest MerchantReturnLineRequest

// NewMerchantReturnLineRequest instantiates a new MerchantReturnLineRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantReturnLineRequest(merchantProductNo string, quantity int32) *MerchantReturnLineRequest {
	this := MerchantReturnLineRequest{}
	this.MerchantProductNo = merchantProductNo
	this.Quantity = quantity
	return &this
}

// NewMerchantReturnLineRequestWithDefaults instantiates a new MerchantReturnLineRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantReturnLineRequestWithDefaults() *MerchantReturnLineRequest {
	this := MerchantReturnLineRequest{}
	return &this
}

// GetMerchantProductNo returns the MerchantProductNo field value
func (o *MerchantReturnLineRequest) GetMerchantProductNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantProductNo
}

// GetMerchantProductNoOk returns a tuple with the MerchantProductNo field value
// and a boolean to check if the value has been set.
func (o *MerchantReturnLineRequest) GetMerchantProductNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantProductNo, true
}

// SetMerchantProductNo sets field value
func (o *MerchantReturnLineRequest) SetMerchantProductNo(v string) {
	o.MerchantProductNo = v
}

// GetQuantity returns the Quantity field value
func (o *MerchantReturnLineRequest) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *MerchantReturnLineRequest) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *MerchantReturnLineRequest) SetQuantity(v int32) {
	o.Quantity = v
}

// GetExtraData returns the ExtraData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantReturnLineRequest) GetExtraData() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.ExtraData
}

// GetExtraDataOk returns a tuple with the ExtraData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantReturnLineRequest) GetExtraDataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ExtraData) {
		return nil, false
	}
	return &o.ExtraData, true
}

// HasExtraData returns a boolean if a field has been set.
func (o *MerchantReturnLineRequest) HasExtraData() bool {
	if o != nil && IsNil(o.ExtraData) {
		return true
	}

	return false
}

// SetExtraData gets a reference to the given map[string]string and assigns it to the ExtraData field.
func (o *MerchantReturnLineRequest) SetExtraData(v map[string]string) {
	o.ExtraData = v
}

func (o MerchantReturnLineRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantReturnLineRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["MerchantProductNo"] = o.MerchantProductNo
	toSerialize["Quantity"] = o.Quantity
	if o.ExtraData != nil {
		toSerialize["ExtraData"] = o.ExtraData
	}
	return toSerialize, nil
}

func (o *MerchantReturnLineRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"MerchantProductNo",
		"Quantity",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMerchantReturnLineRequest := _MerchantReturnLineRequest{}

	err = json.Unmarshal(bytes, &varMerchantReturnLineRequest)

	if err != nil {
		return err
	}

	*o = MerchantReturnLineRequest(varMerchantReturnLineRequest)

	return err
}

type NullableMerchantReturnLineRequest struct {
	value *MerchantReturnLineRequest
	isSet bool
}

func (v NullableMerchantReturnLineRequest) Get() *MerchantReturnLineRequest {
	return v.value
}

func (v *NullableMerchantReturnLineRequest) Set(val *MerchantReturnLineRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantReturnLineRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantReturnLineRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantReturnLineRequest(val *MerchantReturnLineRequest) *NullableMerchantReturnLineRequest {
	return &NullableMerchantReturnLineRequest{value: val, isSet: true}
}

func (v NullableMerchantReturnLineRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantReturnLineRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


