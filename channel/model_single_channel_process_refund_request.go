/*
ChannelEngine Channel API

ChannelEngine API for channels

API version: 2.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package channel

import (
	"encoding/json"
)

// checks if the SingleChannelProcessRefundRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SingleChannelProcessRefundRequest{}

// SingleChannelProcessRefundRequest struct for SingleChannelProcessRefundRequest
type SingleChannelProcessRefundRequest struct {
	IdentifierType *RefundIdentifier `json:"IdentifierType,omitempty"`
	Model *ChannelProcessRefund `json:"Model,omitempty"`
}

// NewSingleChannelProcessRefundRequest instantiates a new SingleChannelProcessRefundRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleChannelProcessRefundRequest() *SingleChannelProcessRefundRequest {
	this := SingleChannelProcessRefundRequest{}
	return &this
}

// NewSingleChannelProcessRefundRequestWithDefaults instantiates a new SingleChannelProcessRefundRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleChannelProcessRefundRequestWithDefaults() *SingleChannelProcessRefundRequest {
	this := SingleChannelProcessRefundRequest{}
	return &this
}

// GetIdentifierType returns the IdentifierType field value if set, zero value otherwise.
func (o *SingleChannelProcessRefundRequest) GetIdentifierType() RefundIdentifier {
	if o == nil || IsNil(o.IdentifierType) {
		var ret RefundIdentifier
		return ret
	}
	return *o.IdentifierType
}

// GetIdentifierTypeOk returns a tuple with the IdentifierType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleChannelProcessRefundRequest) GetIdentifierTypeOk() (*RefundIdentifier, bool) {
	if o == nil || IsNil(o.IdentifierType) {
		return nil, false
	}
	return o.IdentifierType, true
}

// HasIdentifierType returns a boolean if a field has been set.
func (o *SingleChannelProcessRefundRequest) HasIdentifierType() bool {
	if o != nil && !IsNil(o.IdentifierType) {
		return true
	}

	return false
}

// SetIdentifierType gets a reference to the given RefundIdentifier and assigns it to the IdentifierType field.
func (o *SingleChannelProcessRefundRequest) SetIdentifierType(v RefundIdentifier) {
	o.IdentifierType = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *SingleChannelProcessRefundRequest) GetModel() ChannelProcessRefund {
	if o == nil || IsNil(o.Model) {
		var ret ChannelProcessRefund
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleChannelProcessRefundRequest) GetModelOk() (*ChannelProcessRefund, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *SingleChannelProcessRefundRequest) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given ChannelProcessRefund and assigns it to the Model field.
func (o *SingleChannelProcessRefundRequest) SetModel(v ChannelProcessRefund) {
	o.Model = &v
}

func (o SingleChannelProcessRefundRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SingleChannelProcessRefundRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IdentifierType) {
		toSerialize["IdentifierType"] = o.IdentifierType
	}
	if !IsNil(o.Model) {
		toSerialize["Model"] = o.Model
	}
	return toSerialize, nil
}

type NullableSingleChannelProcessRefundRequest struct {
	value *SingleChannelProcessRefundRequest
	isSet bool
}

func (v NullableSingleChannelProcessRefundRequest) Get() *SingleChannelProcessRefundRequest {
	return v.value
}

func (v *NullableSingleChannelProcessRefundRequest) Set(val *SingleChannelProcessRefundRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleChannelProcessRefundRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleChannelProcessRefundRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleChannelProcessRefundRequest(val *SingleChannelProcessRefundRequest) *NullableSingleChannelProcessRefundRequest {
	return &NullableSingleChannelProcessRefundRequest{value: val, isSet: true}
}

func (v NullableSingleChannelProcessRefundRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleChannelProcessRefundRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


