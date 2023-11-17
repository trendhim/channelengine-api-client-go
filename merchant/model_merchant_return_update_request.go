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

// checks if the MerchantReturnUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantReturnUpdateRequest{}

// MerchantReturnUpdateRequest struct for MerchantReturnUpdateRequest
type MerchantReturnUpdateRequest struct {
	// The ChannelEngine return ID of the return you would like to update.
	ReturnId int32 `json:"ReturnId"`
	Lines []MerchantReturnLineUpdateRequest `json:"Lines"`
}

type _MerchantReturnUpdateRequest MerchantReturnUpdateRequest

// NewMerchantReturnUpdateRequest instantiates a new MerchantReturnUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantReturnUpdateRequest(returnId int32, lines []MerchantReturnLineUpdateRequest) *MerchantReturnUpdateRequest {
	this := MerchantReturnUpdateRequest{}
	this.ReturnId = returnId
	this.Lines = lines
	return &this
}

// NewMerchantReturnUpdateRequestWithDefaults instantiates a new MerchantReturnUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantReturnUpdateRequestWithDefaults() *MerchantReturnUpdateRequest {
	this := MerchantReturnUpdateRequest{}
	return &this
}

// GetReturnId returns the ReturnId field value
func (o *MerchantReturnUpdateRequest) GetReturnId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReturnId
}

// GetReturnIdOk returns a tuple with the ReturnId field value
// and a boolean to check if the value has been set.
func (o *MerchantReturnUpdateRequest) GetReturnIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReturnId, true
}

// SetReturnId sets field value
func (o *MerchantReturnUpdateRequest) SetReturnId(v int32) {
	o.ReturnId = v
}

// GetLines returns the Lines field value
func (o *MerchantReturnUpdateRequest) GetLines() []MerchantReturnLineUpdateRequest {
	if o == nil {
		var ret []MerchantReturnLineUpdateRequest
		return ret
	}

	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value
// and a boolean to check if the value has been set.
func (o *MerchantReturnUpdateRequest) GetLinesOk() ([]MerchantReturnLineUpdateRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lines, true
}

// SetLines sets field value
func (o *MerchantReturnUpdateRequest) SetLines(v []MerchantReturnLineUpdateRequest) {
	o.Lines = v
}

func (o MerchantReturnUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantReturnUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ReturnId"] = o.ReturnId
	toSerialize["Lines"] = o.Lines
	return toSerialize, nil
}

func (o *MerchantReturnUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ReturnId",
		"Lines",
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

	varMerchantReturnUpdateRequest := _MerchantReturnUpdateRequest{}

	err = json.Unmarshal(bytes, &varMerchantReturnUpdateRequest)

	if err != nil {
		return err
	}

	*o = MerchantReturnUpdateRequest(varMerchantReturnUpdateRequest)

	return err
}

type NullableMerchantReturnUpdateRequest struct {
	value *MerchantReturnUpdateRequest
	isSet bool
}

func (v NullableMerchantReturnUpdateRequest) Get() *MerchantReturnUpdateRequest {
	return v.value
}

func (v *NullableMerchantReturnUpdateRequest) Set(val *MerchantReturnUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantReturnUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantReturnUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantReturnUpdateRequest(val *MerchantReturnUpdateRequest) *NullableMerchantReturnUpdateRequest {
	return &NullableMerchantReturnUpdateRequest{value: val, isSet: true}
}

func (v NullableMerchantReturnUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantReturnUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


