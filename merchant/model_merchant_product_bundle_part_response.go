/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"encoding/json"
)

// checks if the MerchantProductBundlePartResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantProductBundlePartResponse{}

// MerchantProductBundlePartResponse struct for MerchantProductBundlePartResponse
type MerchantProductBundlePartResponse struct {
	MerchantProductNo NullableString `json:"MerchantProductNo,omitempty"`
	Ean NullableString `json:"Ean,omitempty"`
	Name NullableString `json:"Name,omitempty"`
	Quantity *int32 `json:"Quantity,omitempty"`
	Price *float32 `json:"Price,omitempty"`
}

// NewMerchantProductBundlePartResponse instantiates a new MerchantProductBundlePartResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantProductBundlePartResponse() *MerchantProductBundlePartResponse {
	this := MerchantProductBundlePartResponse{}
	return &this
}

// NewMerchantProductBundlePartResponseWithDefaults instantiates a new MerchantProductBundlePartResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantProductBundlePartResponseWithDefaults() *MerchantProductBundlePartResponse {
	this := MerchantProductBundlePartResponse{}
	return &this
}

// GetMerchantProductNo returns the MerchantProductNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantProductBundlePartResponse) GetMerchantProductNo() string {
	if o == nil || IsNil(o.MerchantProductNo.Get()) {
		var ret string
		return ret
	}
	return *o.MerchantProductNo.Get()
}

// GetMerchantProductNoOk returns a tuple with the MerchantProductNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantProductBundlePartResponse) GetMerchantProductNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MerchantProductNo.Get(), o.MerchantProductNo.IsSet()
}

// HasMerchantProductNo returns a boolean if a field has been set.
func (o *MerchantProductBundlePartResponse) HasMerchantProductNo() bool {
	if o != nil && o.MerchantProductNo.IsSet() {
		return true
	}

	return false
}

// SetMerchantProductNo gets a reference to the given NullableString and assigns it to the MerchantProductNo field.
func (o *MerchantProductBundlePartResponse) SetMerchantProductNo(v string) {
	o.MerchantProductNo.Set(&v)
}
// SetMerchantProductNoNil sets the value for MerchantProductNo to be an explicit nil
func (o *MerchantProductBundlePartResponse) SetMerchantProductNoNil() {
	o.MerchantProductNo.Set(nil)
}

// UnsetMerchantProductNo ensures that no value is present for MerchantProductNo, not even an explicit nil
func (o *MerchantProductBundlePartResponse) UnsetMerchantProductNo() {
	o.MerchantProductNo.Unset()
}

// GetEan returns the Ean field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantProductBundlePartResponse) GetEan() string {
	if o == nil || IsNil(o.Ean.Get()) {
		var ret string
		return ret
	}
	return *o.Ean.Get()
}

// GetEanOk returns a tuple with the Ean field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantProductBundlePartResponse) GetEanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ean.Get(), o.Ean.IsSet()
}

// HasEan returns a boolean if a field has been set.
func (o *MerchantProductBundlePartResponse) HasEan() bool {
	if o != nil && o.Ean.IsSet() {
		return true
	}

	return false
}

// SetEan gets a reference to the given NullableString and assigns it to the Ean field.
func (o *MerchantProductBundlePartResponse) SetEan(v string) {
	o.Ean.Set(&v)
}
// SetEanNil sets the value for Ean to be an explicit nil
func (o *MerchantProductBundlePartResponse) SetEanNil() {
	o.Ean.Set(nil)
}

// UnsetEan ensures that no value is present for Ean, not even an explicit nil
func (o *MerchantProductBundlePartResponse) UnsetEan() {
	o.Ean.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantProductBundlePartResponse) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantProductBundlePartResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MerchantProductBundlePartResponse) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MerchantProductBundlePartResponse) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MerchantProductBundlePartResponse) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MerchantProductBundlePartResponse) UnsetName() {
	o.Name.Unset()
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *MerchantProductBundlePartResponse) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantProductBundlePartResponse) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *MerchantProductBundlePartResponse) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *MerchantProductBundlePartResponse) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *MerchantProductBundlePartResponse) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantProductBundlePartResponse) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *MerchantProductBundlePartResponse) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *MerchantProductBundlePartResponse) SetPrice(v float32) {
	o.Price = &v
}

func (o MerchantProductBundlePartResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantProductBundlePartResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MerchantProductNo.IsSet() {
		toSerialize["MerchantProductNo"] = o.MerchantProductNo.Get()
	}
	if o.Ean.IsSet() {
		toSerialize["Ean"] = o.Ean.Get()
	}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if !IsNil(o.Quantity) {
		toSerialize["Quantity"] = o.Quantity
	}
	if !IsNil(o.Price) {
		toSerialize["Price"] = o.Price
	}
	return toSerialize, nil
}

type NullableMerchantProductBundlePartResponse struct {
	value *MerchantProductBundlePartResponse
	isSet bool
}

func (v NullableMerchantProductBundlePartResponse) Get() *MerchantProductBundlePartResponse {
	return v.value
}

func (v *NullableMerchantProductBundlePartResponse) Set(val *MerchantProductBundlePartResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantProductBundlePartResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantProductBundlePartResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantProductBundlePartResponse(val *MerchantProductBundlePartResponse) *NullableMerchantProductBundlePartResponse {
	return &NullableMerchantProductBundlePartResponse{value: val, isSet: true}
}

func (v NullableMerchantProductBundlePartResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantProductBundlePartResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


