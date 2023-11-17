/*
ChannelEngine Channel API

ChannelEngine API for channels

API version: 2.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package channel

import (
	"encoding/json"
)

// checks if the SingleChannelCreateCancellationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SingleChannelCreateCancellationRequest{}

// SingleChannelCreateCancellationRequest struct for SingleChannelCreateCancellationRequest
type SingleChannelCreateCancellationRequest struct {
	LineIdentifierType *ChannelCreateCancellationsOrderLineIdentifierType `json:"LineIdentifierType,omitempty"`
	IdentifierType *ChannelCreateCancellationsOrderIdentifierType `json:"IdentifierType,omitempty"`
	Model *ChannelCancellation `json:"Model,omitempty"`
}

// NewSingleChannelCreateCancellationRequest instantiates a new SingleChannelCreateCancellationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleChannelCreateCancellationRequest() *SingleChannelCreateCancellationRequest {
	this := SingleChannelCreateCancellationRequest{}
	return &this
}

// NewSingleChannelCreateCancellationRequestWithDefaults instantiates a new SingleChannelCreateCancellationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleChannelCreateCancellationRequestWithDefaults() *SingleChannelCreateCancellationRequest {
	this := SingleChannelCreateCancellationRequest{}
	return &this
}

// GetLineIdentifierType returns the LineIdentifierType field value if set, zero value otherwise.
func (o *SingleChannelCreateCancellationRequest) GetLineIdentifierType() ChannelCreateCancellationsOrderLineIdentifierType {
	if o == nil || IsNil(o.LineIdentifierType) {
		var ret ChannelCreateCancellationsOrderLineIdentifierType
		return ret
	}
	return *o.LineIdentifierType
}

// GetLineIdentifierTypeOk returns a tuple with the LineIdentifierType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleChannelCreateCancellationRequest) GetLineIdentifierTypeOk() (*ChannelCreateCancellationsOrderLineIdentifierType, bool) {
	if o == nil || IsNil(o.LineIdentifierType) {
		return nil, false
	}
	return o.LineIdentifierType, true
}

// HasLineIdentifierType returns a boolean if a field has been set.
func (o *SingleChannelCreateCancellationRequest) HasLineIdentifierType() bool {
	if o != nil && !IsNil(o.LineIdentifierType) {
		return true
	}

	return false
}

// SetLineIdentifierType gets a reference to the given ChannelCreateCancellationsOrderLineIdentifierType and assigns it to the LineIdentifierType field.
func (o *SingleChannelCreateCancellationRequest) SetLineIdentifierType(v ChannelCreateCancellationsOrderLineIdentifierType) {
	o.LineIdentifierType = &v
}

// GetIdentifierType returns the IdentifierType field value if set, zero value otherwise.
func (o *SingleChannelCreateCancellationRequest) GetIdentifierType() ChannelCreateCancellationsOrderIdentifierType {
	if o == nil || IsNil(o.IdentifierType) {
		var ret ChannelCreateCancellationsOrderIdentifierType
		return ret
	}
	return *o.IdentifierType
}

// GetIdentifierTypeOk returns a tuple with the IdentifierType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleChannelCreateCancellationRequest) GetIdentifierTypeOk() (*ChannelCreateCancellationsOrderIdentifierType, bool) {
	if o == nil || IsNil(o.IdentifierType) {
		return nil, false
	}
	return o.IdentifierType, true
}

// HasIdentifierType returns a boolean if a field has been set.
func (o *SingleChannelCreateCancellationRequest) HasIdentifierType() bool {
	if o != nil && !IsNil(o.IdentifierType) {
		return true
	}

	return false
}

// SetIdentifierType gets a reference to the given ChannelCreateCancellationsOrderIdentifierType and assigns it to the IdentifierType field.
func (o *SingleChannelCreateCancellationRequest) SetIdentifierType(v ChannelCreateCancellationsOrderIdentifierType) {
	o.IdentifierType = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *SingleChannelCreateCancellationRequest) GetModel() ChannelCancellation {
	if o == nil || IsNil(o.Model) {
		var ret ChannelCancellation
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleChannelCreateCancellationRequest) GetModelOk() (*ChannelCancellation, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *SingleChannelCreateCancellationRequest) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given ChannelCancellation and assigns it to the Model field.
func (o *SingleChannelCreateCancellationRequest) SetModel(v ChannelCancellation) {
	o.Model = &v
}

func (o SingleChannelCreateCancellationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SingleChannelCreateCancellationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LineIdentifierType) {
		toSerialize["LineIdentifierType"] = o.LineIdentifierType
	}
	if !IsNil(o.IdentifierType) {
		toSerialize["IdentifierType"] = o.IdentifierType
	}
	if !IsNil(o.Model) {
		toSerialize["Model"] = o.Model
	}
	return toSerialize, nil
}

type NullableSingleChannelCreateCancellationRequest struct {
	value *SingleChannelCreateCancellationRequest
	isSet bool
}

func (v NullableSingleChannelCreateCancellationRequest) Get() *SingleChannelCreateCancellationRequest {
	return v.value
}

func (v *NullableSingleChannelCreateCancellationRequest) Set(val *SingleChannelCreateCancellationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleChannelCreateCancellationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleChannelCreateCancellationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleChannelCreateCancellationRequest(val *SingleChannelCreateCancellationRequest) *NullableSingleChannelCreateCancellationRequest {
	return &NullableSingleChannelCreateCancellationRequest{value: val, isSet: true}
}

func (v NullableSingleChannelCreateCancellationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleChannelCreateCancellationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

