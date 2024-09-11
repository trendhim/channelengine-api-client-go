/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"encoding/json"
)

// checks if the SingleMerchantCreateRefundRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SingleMerchantCreateRefundRequest{}

// SingleMerchantCreateRefundRequest struct for SingleMerchantCreateRefundRequest
type SingleMerchantCreateRefundRequest struct {
	ChannelId NullableInt32 `json:"ChannelId,omitempty"`
	LineIdentifierType *OrderLineIdentifier `json:"LineIdentifierType,omitempty"`
	IdentifierType *OrderIdentifier `json:"IdentifierType,omitempty"`
	Model *MerchantCreateRefund `json:"Model,omitempty"`
}

// NewSingleMerchantCreateRefundRequest instantiates a new SingleMerchantCreateRefundRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleMerchantCreateRefundRequest() *SingleMerchantCreateRefundRequest {
	this := SingleMerchantCreateRefundRequest{}
	return &this
}

// NewSingleMerchantCreateRefundRequestWithDefaults instantiates a new SingleMerchantCreateRefundRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleMerchantCreateRefundRequestWithDefaults() *SingleMerchantCreateRefundRequest {
	this := SingleMerchantCreateRefundRequest{}
	return &this
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SingleMerchantCreateRefundRequest) GetChannelId() int32 {
	if o == nil || IsNil(o.ChannelId.Get()) {
		var ret int32
		return ret
	}
	return *o.ChannelId.Get()
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SingleMerchantCreateRefundRequest) GetChannelIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelId.Get(), o.ChannelId.IsSet()
}

// HasChannelId returns a boolean if a field has been set.
func (o *SingleMerchantCreateRefundRequest) HasChannelId() bool {
	if o != nil && o.ChannelId.IsSet() {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given NullableInt32 and assigns it to the ChannelId field.
func (o *SingleMerchantCreateRefundRequest) SetChannelId(v int32) {
	o.ChannelId.Set(&v)
}
// SetChannelIdNil sets the value for ChannelId to be an explicit nil
func (o *SingleMerchantCreateRefundRequest) SetChannelIdNil() {
	o.ChannelId.Set(nil)
}

// UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
func (o *SingleMerchantCreateRefundRequest) UnsetChannelId() {
	o.ChannelId.Unset()
}

// GetLineIdentifierType returns the LineIdentifierType field value if set, zero value otherwise.
func (o *SingleMerchantCreateRefundRequest) GetLineIdentifierType() OrderLineIdentifier {
	if o == nil || IsNil(o.LineIdentifierType) {
		var ret OrderLineIdentifier
		return ret
	}
	return *o.LineIdentifierType
}

// GetLineIdentifierTypeOk returns a tuple with the LineIdentifierType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleMerchantCreateRefundRequest) GetLineIdentifierTypeOk() (*OrderLineIdentifier, bool) {
	if o == nil || IsNil(o.LineIdentifierType) {
		return nil, false
	}
	return o.LineIdentifierType, true
}

// HasLineIdentifierType returns a boolean if a field has been set.
func (o *SingleMerchantCreateRefundRequest) HasLineIdentifierType() bool {
	if o != nil && !IsNil(o.LineIdentifierType) {
		return true
	}

	return false
}

// SetLineIdentifierType gets a reference to the given OrderLineIdentifier and assigns it to the LineIdentifierType field.
func (o *SingleMerchantCreateRefundRequest) SetLineIdentifierType(v OrderLineIdentifier) {
	o.LineIdentifierType = &v
}

// GetIdentifierType returns the IdentifierType field value if set, zero value otherwise.
func (o *SingleMerchantCreateRefundRequest) GetIdentifierType() OrderIdentifier {
	if o == nil || IsNil(o.IdentifierType) {
		var ret OrderIdentifier
		return ret
	}
	return *o.IdentifierType
}

// GetIdentifierTypeOk returns a tuple with the IdentifierType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleMerchantCreateRefundRequest) GetIdentifierTypeOk() (*OrderIdentifier, bool) {
	if o == nil || IsNil(o.IdentifierType) {
		return nil, false
	}
	return o.IdentifierType, true
}

// HasIdentifierType returns a boolean if a field has been set.
func (o *SingleMerchantCreateRefundRequest) HasIdentifierType() bool {
	if o != nil && !IsNil(o.IdentifierType) {
		return true
	}

	return false
}

// SetIdentifierType gets a reference to the given OrderIdentifier and assigns it to the IdentifierType field.
func (o *SingleMerchantCreateRefundRequest) SetIdentifierType(v OrderIdentifier) {
	o.IdentifierType = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *SingleMerchantCreateRefundRequest) GetModel() MerchantCreateRefund {
	if o == nil || IsNil(o.Model) {
		var ret MerchantCreateRefund
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleMerchantCreateRefundRequest) GetModelOk() (*MerchantCreateRefund, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *SingleMerchantCreateRefundRequest) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given MerchantCreateRefund and assigns it to the Model field.
func (o *SingleMerchantCreateRefundRequest) SetModel(v MerchantCreateRefund) {
	o.Model = &v
}

func (o SingleMerchantCreateRefundRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SingleMerchantCreateRefundRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ChannelId.IsSet() {
		toSerialize["ChannelId"] = o.ChannelId.Get()
	}
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

type NullableSingleMerchantCreateRefundRequest struct {
	value *SingleMerchantCreateRefundRequest
	isSet bool
}

func (v NullableSingleMerchantCreateRefundRequest) Get() *SingleMerchantCreateRefundRequest {
	return v.value
}

func (v *NullableSingleMerchantCreateRefundRequest) Set(val *SingleMerchantCreateRefundRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleMerchantCreateRefundRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleMerchantCreateRefundRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleMerchantCreateRefundRequest(val *SingleMerchantCreateRefundRequest) *NullableSingleMerchantCreateRefundRequest {
	return &NullableSingleMerchantCreateRefundRequest{value: val, isSet: true}
}

func (v NullableSingleMerchantCreateRefundRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleMerchantCreateRefundRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

