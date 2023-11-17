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

// checks if the ChannelProductReferencesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelProductReferencesRequest{}

// ChannelProductReferencesRequest struct for ChannelProductReferencesRequest
type ChannelProductReferencesRequest struct {
	// The unique ChannelEngine product ID.
	Id *int32 `json:"Id,omitempty"`
	// The unique product reference used by the Channel.
	ChannelProductNo string `json:"ChannelProductNo"`
}

type _ChannelProductReferencesRequest ChannelProductReferencesRequest

// NewChannelProductReferencesRequest instantiates a new ChannelProductReferencesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelProductReferencesRequest(channelProductNo string) *ChannelProductReferencesRequest {
	this := ChannelProductReferencesRequest{}
	this.ChannelProductNo = channelProductNo
	return &this
}

// NewChannelProductReferencesRequestWithDefaults instantiates a new ChannelProductReferencesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelProductReferencesRequestWithDefaults() *ChannelProductReferencesRequest {
	this := ChannelProductReferencesRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChannelProductReferencesRequest) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelProductReferencesRequest) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChannelProductReferencesRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ChannelProductReferencesRequest) SetId(v int32) {
	o.Id = &v
}

// GetChannelProductNo returns the ChannelProductNo field value
func (o *ChannelProductReferencesRequest) GetChannelProductNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelProductNo
}

// GetChannelProductNoOk returns a tuple with the ChannelProductNo field value
// and a boolean to check if the value has been set.
func (o *ChannelProductReferencesRequest) GetChannelProductNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelProductNo, true
}

// SetChannelProductNo sets field value
func (o *ChannelProductReferencesRequest) SetChannelProductNo(v string) {
	o.ChannelProductNo = v
}

func (o ChannelProductReferencesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelProductReferencesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	toSerialize["ChannelProductNo"] = o.ChannelProductNo
	return toSerialize, nil
}

func (o *ChannelProductReferencesRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ChannelProductNo",
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

	varChannelProductReferencesRequest := _ChannelProductReferencesRequest{}

	err = json.Unmarshal(bytes, &varChannelProductReferencesRequest)

	if err != nil {
		return err
	}

	*o = ChannelProductReferencesRequest(varChannelProductReferencesRequest)

	return err
}

type NullableChannelProductReferencesRequest struct {
	value *ChannelProductReferencesRequest
	isSet bool
}

func (v NullableChannelProductReferencesRequest) Get() *ChannelProductReferencesRequest {
	return v.value
}

func (v *NullableChannelProductReferencesRequest) Set(val *ChannelProductReferencesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelProductReferencesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelProductReferencesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelProductReferencesRequest(val *ChannelProductReferencesRequest) *NullableChannelProductReferencesRequest {
	return &NullableChannelProductReferencesRequest{value: val, isSet: true}
}

func (v NullableChannelProductReferencesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelProductReferencesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


