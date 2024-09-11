/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the MerchantShipmentLineRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantShipmentLineRequest{}

// MerchantShipmentLineRequest struct for MerchantShipmentLineRequest
type MerchantShipmentLineRequest struct {
	// The unique product reference used by the Merchant (sku).
	MerchantProductNo string `json:"MerchantProductNo"`
	// Extra data on the order. Each item must have an unqiue key
	ExtraData map[string]string `json:"ExtraData,omitempty"`
	// Number of items of the product in the shipment.
	Quantity int32 `json:"Quantity"`
}

type _MerchantShipmentLineRequest MerchantShipmentLineRequest

// NewMerchantShipmentLineRequest instantiates a new MerchantShipmentLineRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantShipmentLineRequest(merchantProductNo string, quantity int32) *MerchantShipmentLineRequest {
	this := MerchantShipmentLineRequest{}
	this.MerchantProductNo = merchantProductNo
	this.Quantity = quantity
	return &this
}

// NewMerchantShipmentLineRequestWithDefaults instantiates a new MerchantShipmentLineRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantShipmentLineRequestWithDefaults() *MerchantShipmentLineRequest {
	this := MerchantShipmentLineRequest{}
	return &this
}

// GetMerchantProductNo returns the MerchantProductNo field value
func (o *MerchantShipmentLineRequest) GetMerchantProductNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantProductNo
}

// GetMerchantProductNoOk returns a tuple with the MerchantProductNo field value
// and a boolean to check if the value has been set.
func (o *MerchantShipmentLineRequest) GetMerchantProductNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantProductNo, true
}

// SetMerchantProductNo sets field value
func (o *MerchantShipmentLineRequest) SetMerchantProductNo(v string) {
	o.MerchantProductNo = v
}

// GetExtraData returns the ExtraData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantShipmentLineRequest) GetExtraData() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.ExtraData
}

// GetExtraDataOk returns a tuple with the ExtraData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantShipmentLineRequest) GetExtraDataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ExtraData) {
		return nil, false
	}
	return &o.ExtraData, true
}

// HasExtraData returns a boolean if a field has been set.
func (o *MerchantShipmentLineRequest) HasExtraData() bool {
	if o != nil && !IsNil(o.ExtraData) {
		return true
	}

	return false
}

// SetExtraData gets a reference to the given map[string]string and assigns it to the ExtraData field.
func (o *MerchantShipmentLineRequest) SetExtraData(v map[string]string) {
	o.ExtraData = v
}

// GetQuantity returns the Quantity field value
func (o *MerchantShipmentLineRequest) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *MerchantShipmentLineRequest) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *MerchantShipmentLineRequest) SetQuantity(v int32) {
	o.Quantity = v
}

func (o MerchantShipmentLineRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantShipmentLineRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["MerchantProductNo"] = o.MerchantProductNo
	if o.ExtraData != nil {
		toSerialize["ExtraData"] = o.ExtraData
	}
	toSerialize["Quantity"] = o.Quantity
	return toSerialize, nil
}

func (o *MerchantShipmentLineRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"MerchantProductNo",
		"Quantity",
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

	varMerchantShipmentLineRequest := _MerchantShipmentLineRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMerchantShipmentLineRequest)

	if err != nil {
		return err
	}

	*o = MerchantShipmentLineRequest(varMerchantShipmentLineRequest)

	return err
}

type NullableMerchantShipmentLineRequest struct {
	value *MerchantShipmentLineRequest
	isSet bool
}

func (v NullableMerchantShipmentLineRequest) Get() *MerchantShipmentLineRequest {
	return v.value
}

func (v *NullableMerchantShipmentLineRequest) Set(val *MerchantShipmentLineRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantShipmentLineRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantShipmentLineRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantShipmentLineRequest(val *MerchantShipmentLineRequest) *NullableMerchantShipmentLineRequest {
	return &NullableMerchantShipmentLineRequest{value: val, isSet: true}
}

func (v NullableMerchantShipmentLineRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantShipmentLineRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


