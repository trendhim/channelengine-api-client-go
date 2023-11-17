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

// checks if the MerchantReturnAcknowledgeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantReturnAcknowledgeRequest{}

// MerchantReturnAcknowledgeRequest struct for MerchantReturnAcknowledgeRequest
type MerchantReturnAcknowledgeRequest struct {
	ReturnId *int32 `json:"ReturnId,omitempty"`
	MerchantReturnNo string `json:"MerchantReturnNo"`
}

type _MerchantReturnAcknowledgeRequest MerchantReturnAcknowledgeRequest

// NewMerchantReturnAcknowledgeRequest instantiates a new MerchantReturnAcknowledgeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantReturnAcknowledgeRequest(merchantReturnNo string) *MerchantReturnAcknowledgeRequest {
	this := MerchantReturnAcknowledgeRequest{}
	this.MerchantReturnNo = merchantReturnNo
	return &this
}

// NewMerchantReturnAcknowledgeRequestWithDefaults instantiates a new MerchantReturnAcknowledgeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantReturnAcknowledgeRequestWithDefaults() *MerchantReturnAcknowledgeRequest {
	this := MerchantReturnAcknowledgeRequest{}
	return &this
}

// GetReturnId returns the ReturnId field value if set, zero value otherwise.
func (o *MerchantReturnAcknowledgeRequest) GetReturnId() int32 {
	if o == nil || IsNil(o.ReturnId) {
		var ret int32
		return ret
	}
	return *o.ReturnId
}

// GetReturnIdOk returns a tuple with the ReturnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantReturnAcknowledgeRequest) GetReturnIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ReturnId) {
		return nil, false
	}
	return o.ReturnId, true
}

// HasReturnId returns a boolean if a field has been set.
func (o *MerchantReturnAcknowledgeRequest) HasReturnId() bool {
	if o != nil && !IsNil(o.ReturnId) {
		return true
	}

	return false
}

// SetReturnId gets a reference to the given int32 and assigns it to the ReturnId field.
func (o *MerchantReturnAcknowledgeRequest) SetReturnId(v int32) {
	o.ReturnId = &v
}

// GetMerchantReturnNo returns the MerchantReturnNo field value
func (o *MerchantReturnAcknowledgeRequest) GetMerchantReturnNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantReturnNo
}

// GetMerchantReturnNoOk returns a tuple with the MerchantReturnNo field value
// and a boolean to check if the value has been set.
func (o *MerchantReturnAcknowledgeRequest) GetMerchantReturnNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantReturnNo, true
}

// SetMerchantReturnNo sets field value
func (o *MerchantReturnAcknowledgeRequest) SetMerchantReturnNo(v string) {
	o.MerchantReturnNo = v
}

func (o MerchantReturnAcknowledgeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantReturnAcknowledgeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReturnId) {
		toSerialize["ReturnId"] = o.ReturnId
	}
	toSerialize["MerchantReturnNo"] = o.MerchantReturnNo
	return toSerialize, nil
}

func (o *MerchantReturnAcknowledgeRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"MerchantReturnNo",
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

	varMerchantReturnAcknowledgeRequest := _MerchantReturnAcknowledgeRequest{}

	err = json.Unmarshal(bytes, &varMerchantReturnAcknowledgeRequest)

	if err != nil {
		return err
	}

	*o = MerchantReturnAcknowledgeRequest(varMerchantReturnAcknowledgeRequest)

	return err
}

type NullableMerchantReturnAcknowledgeRequest struct {
	value *MerchantReturnAcknowledgeRequest
	isSet bool
}

func (v NullableMerchantReturnAcknowledgeRequest) Get() *MerchantReturnAcknowledgeRequest {
	return v.value
}

func (v *NullableMerchantReturnAcknowledgeRequest) Set(val *MerchantReturnAcknowledgeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantReturnAcknowledgeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantReturnAcknowledgeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantReturnAcknowledgeRequest(val *MerchantReturnAcknowledgeRequest) *NullableMerchantReturnAcknowledgeRequest {
	return &NullableMerchantReturnAcknowledgeRequest{value: val, isSet: true}
}

func (v NullableMerchantReturnAcknowledgeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantReturnAcknowledgeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

