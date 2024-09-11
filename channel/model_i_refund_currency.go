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

// checks if the IRefundCurrency type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IRefundCurrency{}

// IRefundCurrency struct for IRefundCurrency
type IRefundCurrency struct {
	TenantCurrencyCode NullableString `json:"TenantCurrencyCode,omitempty"`
	OriginalCurrencyCode NullableString `json:"OriginalCurrencyCode,omitempty"`
	ExchangeRate *float32 `json:"ExchangeRate,omitempty"`
}

// NewIRefundCurrency instantiates a new IRefundCurrency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIRefundCurrency() *IRefundCurrency {
	this := IRefundCurrency{}
	return &this
}

// NewIRefundCurrencyWithDefaults instantiates a new IRefundCurrency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIRefundCurrencyWithDefaults() *IRefundCurrency {
	this := IRefundCurrency{}
	return &this
}

// GetTenantCurrencyCode returns the TenantCurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IRefundCurrency) GetTenantCurrencyCode() string {
	if o == nil || IsNil(o.TenantCurrencyCode.Get()) {
		var ret string
		return ret
	}
	return *o.TenantCurrencyCode.Get()
}

// GetTenantCurrencyCodeOk returns a tuple with the TenantCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IRefundCurrency) GetTenantCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantCurrencyCode.Get(), o.TenantCurrencyCode.IsSet()
}

// HasTenantCurrencyCode returns a boolean if a field has been set.
func (o *IRefundCurrency) HasTenantCurrencyCode() bool {
	if o != nil && o.TenantCurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetTenantCurrencyCode gets a reference to the given NullableString and assigns it to the TenantCurrencyCode field.
func (o *IRefundCurrency) SetTenantCurrencyCode(v string) {
	o.TenantCurrencyCode.Set(&v)
}
// SetTenantCurrencyCodeNil sets the value for TenantCurrencyCode to be an explicit nil
func (o *IRefundCurrency) SetTenantCurrencyCodeNil() {
	o.TenantCurrencyCode.Set(nil)
}

// UnsetTenantCurrencyCode ensures that no value is present for TenantCurrencyCode, not even an explicit nil
func (o *IRefundCurrency) UnsetTenantCurrencyCode() {
	o.TenantCurrencyCode.Unset()
}

// GetOriginalCurrencyCode returns the OriginalCurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IRefundCurrency) GetOriginalCurrencyCode() string {
	if o == nil || IsNil(o.OriginalCurrencyCode.Get()) {
		var ret string
		return ret
	}
	return *o.OriginalCurrencyCode.Get()
}

// GetOriginalCurrencyCodeOk returns a tuple with the OriginalCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IRefundCurrency) GetOriginalCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginalCurrencyCode.Get(), o.OriginalCurrencyCode.IsSet()
}

// HasOriginalCurrencyCode returns a boolean if a field has been set.
func (o *IRefundCurrency) HasOriginalCurrencyCode() bool {
	if o != nil && o.OriginalCurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetOriginalCurrencyCode gets a reference to the given NullableString and assigns it to the OriginalCurrencyCode field.
func (o *IRefundCurrency) SetOriginalCurrencyCode(v string) {
	o.OriginalCurrencyCode.Set(&v)
}
// SetOriginalCurrencyCodeNil sets the value for OriginalCurrencyCode to be an explicit nil
func (o *IRefundCurrency) SetOriginalCurrencyCodeNil() {
	o.OriginalCurrencyCode.Set(nil)
}

// UnsetOriginalCurrencyCode ensures that no value is present for OriginalCurrencyCode, not even an explicit nil
func (o *IRefundCurrency) UnsetOriginalCurrencyCode() {
	o.OriginalCurrencyCode.Unset()
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise.
func (o *IRefundCurrency) GetExchangeRate() float32 {
	if o == nil || IsNil(o.ExchangeRate) {
		var ret float32
		return ret
	}
	return *o.ExchangeRate
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IRefundCurrency) GetExchangeRateOk() (*float32, bool) {
	if o == nil || IsNil(o.ExchangeRate) {
		return nil, false
	}
	return o.ExchangeRate, true
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *IRefundCurrency) HasExchangeRate() bool {
	if o != nil && !IsNil(o.ExchangeRate) {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given float32 and assigns it to the ExchangeRate field.
func (o *IRefundCurrency) SetExchangeRate(v float32) {
	o.ExchangeRate = &v
}

func (o IRefundCurrency) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IRefundCurrency) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TenantCurrencyCode.IsSet() {
		toSerialize["TenantCurrencyCode"] = o.TenantCurrencyCode.Get()
	}
	if o.OriginalCurrencyCode.IsSet() {
		toSerialize["OriginalCurrencyCode"] = o.OriginalCurrencyCode.Get()
	}
	if !IsNil(o.ExchangeRate) {
		toSerialize["ExchangeRate"] = o.ExchangeRate
	}
	return toSerialize, nil
}

type NullableIRefundCurrency struct {
	value *IRefundCurrency
	isSet bool
}

func (v NullableIRefundCurrency) Get() *IRefundCurrency {
	return v.value
}

func (v *NullableIRefundCurrency) Set(val *IRefundCurrency) {
	v.value = val
	v.isSet = true
}

func (v NullableIRefundCurrency) IsSet() bool {
	return v.isSet
}

func (v *NullableIRefundCurrency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIRefundCurrency(val *IRefundCurrency) *NullableIRefundCurrency {
	return &NullableIRefundCurrency{value: val, isSet: true}
}

func (v NullableIRefundCurrency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIRefundCurrency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

