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

// checks if the MerchantCreateRefundLine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantCreateRefundLine{}

// MerchantCreateRefundLine struct for MerchantCreateRefundLine
type MerchantCreateRefundLine struct {
	OrderLineIdentifier NullableString `json:"OrderLineIdentifier,omitempty"`
	LineAmountInclTax *float32 `json:"LineAmountInclTax,omitempty"`
	MerchantRefundLineNo NullableString `json:"MerchantRefundLineNo,omitempty"`
	ExtraData map[string]string `json:"ExtraData,omitempty"`
}

// NewMerchantCreateRefundLine instantiates a new MerchantCreateRefundLine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantCreateRefundLine() *MerchantCreateRefundLine {
	this := MerchantCreateRefundLine{}
	return &this
}

// NewMerchantCreateRefundLineWithDefaults instantiates a new MerchantCreateRefundLine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantCreateRefundLineWithDefaults() *MerchantCreateRefundLine {
	this := MerchantCreateRefundLine{}
	return &this
}

// GetOrderLineIdentifier returns the OrderLineIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantCreateRefundLine) GetOrderLineIdentifier() string {
	if o == nil || IsNil(o.OrderLineIdentifier.Get()) {
		var ret string
		return ret
	}
	return *o.OrderLineIdentifier.Get()
}

// GetOrderLineIdentifierOk returns a tuple with the OrderLineIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantCreateRefundLine) GetOrderLineIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrderLineIdentifier.Get(), o.OrderLineIdentifier.IsSet()
}

// HasOrderLineIdentifier returns a boolean if a field has been set.
func (o *MerchantCreateRefundLine) HasOrderLineIdentifier() bool {
	if o != nil && o.OrderLineIdentifier.IsSet() {
		return true
	}

	return false
}

// SetOrderLineIdentifier gets a reference to the given NullableString and assigns it to the OrderLineIdentifier field.
func (o *MerchantCreateRefundLine) SetOrderLineIdentifier(v string) {
	o.OrderLineIdentifier.Set(&v)
}
// SetOrderLineIdentifierNil sets the value for OrderLineIdentifier to be an explicit nil
func (o *MerchantCreateRefundLine) SetOrderLineIdentifierNil() {
	o.OrderLineIdentifier.Set(nil)
}

// UnsetOrderLineIdentifier ensures that no value is present for OrderLineIdentifier, not even an explicit nil
func (o *MerchantCreateRefundLine) UnsetOrderLineIdentifier() {
	o.OrderLineIdentifier.Unset()
}

// GetLineAmountInclTax returns the LineAmountInclTax field value if set, zero value otherwise.
func (o *MerchantCreateRefundLine) GetLineAmountInclTax() float32 {
	if o == nil || IsNil(o.LineAmountInclTax) {
		var ret float32
		return ret
	}
	return *o.LineAmountInclTax
}

// GetLineAmountInclTaxOk returns a tuple with the LineAmountInclTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCreateRefundLine) GetLineAmountInclTaxOk() (*float32, bool) {
	if o == nil || IsNil(o.LineAmountInclTax) {
		return nil, false
	}
	return o.LineAmountInclTax, true
}

// HasLineAmountInclTax returns a boolean if a field has been set.
func (o *MerchantCreateRefundLine) HasLineAmountInclTax() bool {
	if o != nil && !IsNil(o.LineAmountInclTax) {
		return true
	}

	return false
}

// SetLineAmountInclTax gets a reference to the given float32 and assigns it to the LineAmountInclTax field.
func (o *MerchantCreateRefundLine) SetLineAmountInclTax(v float32) {
	o.LineAmountInclTax = &v
}

// GetMerchantRefundLineNo returns the MerchantRefundLineNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantCreateRefundLine) GetMerchantRefundLineNo() string {
	if o == nil || IsNil(o.MerchantRefundLineNo.Get()) {
		var ret string
		return ret
	}
	return *o.MerchantRefundLineNo.Get()
}

// GetMerchantRefundLineNoOk returns a tuple with the MerchantRefundLineNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantCreateRefundLine) GetMerchantRefundLineNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MerchantRefundLineNo.Get(), o.MerchantRefundLineNo.IsSet()
}

// HasMerchantRefundLineNo returns a boolean if a field has been set.
func (o *MerchantCreateRefundLine) HasMerchantRefundLineNo() bool {
	if o != nil && o.MerchantRefundLineNo.IsSet() {
		return true
	}

	return false
}

// SetMerchantRefundLineNo gets a reference to the given NullableString and assigns it to the MerchantRefundLineNo field.
func (o *MerchantCreateRefundLine) SetMerchantRefundLineNo(v string) {
	o.MerchantRefundLineNo.Set(&v)
}
// SetMerchantRefundLineNoNil sets the value for MerchantRefundLineNo to be an explicit nil
func (o *MerchantCreateRefundLine) SetMerchantRefundLineNoNil() {
	o.MerchantRefundLineNo.Set(nil)
}

// UnsetMerchantRefundLineNo ensures that no value is present for MerchantRefundLineNo, not even an explicit nil
func (o *MerchantCreateRefundLine) UnsetMerchantRefundLineNo() {
	o.MerchantRefundLineNo.Unset()
}

// GetExtraData returns the ExtraData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantCreateRefundLine) GetExtraData() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.ExtraData
}

// GetExtraDataOk returns a tuple with the ExtraData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantCreateRefundLine) GetExtraDataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ExtraData) {
		return nil, false
	}
	return &o.ExtraData, true
}

// HasExtraData returns a boolean if a field has been set.
func (o *MerchantCreateRefundLine) HasExtraData() bool {
	if o != nil && !IsNil(o.ExtraData) {
		return true
	}

	return false
}

// SetExtraData gets a reference to the given map[string]string and assigns it to the ExtraData field.
func (o *MerchantCreateRefundLine) SetExtraData(v map[string]string) {
	o.ExtraData = v
}

func (o MerchantCreateRefundLine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantCreateRefundLine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OrderLineIdentifier.IsSet() {
		toSerialize["OrderLineIdentifier"] = o.OrderLineIdentifier.Get()
	}
	if !IsNil(o.LineAmountInclTax) {
		toSerialize["LineAmountInclTax"] = o.LineAmountInclTax
	}
	if o.MerchantRefundLineNo.IsSet() {
		toSerialize["MerchantRefundLineNo"] = o.MerchantRefundLineNo.Get()
	}
	if o.ExtraData != nil {
		toSerialize["ExtraData"] = o.ExtraData
	}
	return toSerialize, nil
}

type NullableMerchantCreateRefundLine struct {
	value *MerchantCreateRefundLine
	isSet bool
}

func (v NullableMerchantCreateRefundLine) Get() *MerchantCreateRefundLine {
	return v.value
}

func (v *NullableMerchantCreateRefundLine) Set(val *MerchantCreateRefundLine) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantCreateRefundLine) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantCreateRefundLine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantCreateRefundLine(val *MerchantCreateRefundLine) *NullableMerchantCreateRefundLine {
	return &NullableMerchantCreateRefundLine{value: val, isSet: true}
}

func (v NullableMerchantCreateRefundLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantCreateRefundLine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


