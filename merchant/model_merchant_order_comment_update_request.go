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

// checks if the MerchantOrderCommentUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantOrderCommentUpdateRequest{}

// MerchantOrderCommentUpdateRequest struct for MerchantOrderCommentUpdateRequest
type MerchantOrderCommentUpdateRequest struct {
	// Your own order reference for the order you would like to update the comment for.  Either this field or OrderId is required
	MerchantOrderNo NullableString `json:"MerchantOrderNo,omitempty"`
	// The ChannelEngine order ID of the order you would like to update the comment for.  Either this field or MerchantOrderNo is required
	OrderId NullableInt32 `json:"OrderId,omitempty"`
	// The merchant comment you would like add / update for the order.
	MerchantComment string `json:"MerchantComment"`
}

type _MerchantOrderCommentUpdateRequest MerchantOrderCommentUpdateRequest

// NewMerchantOrderCommentUpdateRequest instantiates a new MerchantOrderCommentUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantOrderCommentUpdateRequest(merchantComment string) *MerchantOrderCommentUpdateRequest {
	this := MerchantOrderCommentUpdateRequest{}
	this.MerchantComment = merchantComment
	return &this
}

// NewMerchantOrderCommentUpdateRequestWithDefaults instantiates a new MerchantOrderCommentUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantOrderCommentUpdateRequestWithDefaults() *MerchantOrderCommentUpdateRequest {
	this := MerchantOrderCommentUpdateRequest{}
	return &this
}

// GetMerchantOrderNo returns the MerchantOrderNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantOrderCommentUpdateRequest) GetMerchantOrderNo() string {
	if o == nil || IsNil(o.MerchantOrderNo.Get()) {
		var ret string
		return ret
	}
	return *o.MerchantOrderNo.Get()
}

// GetMerchantOrderNoOk returns a tuple with the MerchantOrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantOrderCommentUpdateRequest) GetMerchantOrderNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MerchantOrderNo.Get(), o.MerchantOrderNo.IsSet()
}

// HasMerchantOrderNo returns a boolean if a field has been set.
func (o *MerchantOrderCommentUpdateRequest) HasMerchantOrderNo() bool {
	if o != nil && o.MerchantOrderNo.IsSet() {
		return true
	}

	return false
}

// SetMerchantOrderNo gets a reference to the given NullableString and assigns it to the MerchantOrderNo field.
func (o *MerchantOrderCommentUpdateRequest) SetMerchantOrderNo(v string) {
	o.MerchantOrderNo.Set(&v)
}
// SetMerchantOrderNoNil sets the value for MerchantOrderNo to be an explicit nil
func (o *MerchantOrderCommentUpdateRequest) SetMerchantOrderNoNil() {
	o.MerchantOrderNo.Set(nil)
}

// UnsetMerchantOrderNo ensures that no value is present for MerchantOrderNo, not even an explicit nil
func (o *MerchantOrderCommentUpdateRequest) UnsetMerchantOrderNo() {
	o.MerchantOrderNo.Unset()
}

// GetOrderId returns the OrderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantOrderCommentUpdateRequest) GetOrderId() int32 {
	if o == nil || IsNil(o.OrderId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrderId.Get()
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantOrderCommentUpdateRequest) GetOrderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrderId.Get(), o.OrderId.IsSet()
}

// HasOrderId returns a boolean if a field has been set.
func (o *MerchantOrderCommentUpdateRequest) HasOrderId() bool {
	if o != nil && o.OrderId.IsSet() {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given NullableInt32 and assigns it to the OrderId field.
func (o *MerchantOrderCommentUpdateRequest) SetOrderId(v int32) {
	o.OrderId.Set(&v)
}
// SetOrderIdNil sets the value for OrderId to be an explicit nil
func (o *MerchantOrderCommentUpdateRequest) SetOrderIdNil() {
	o.OrderId.Set(nil)
}

// UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
func (o *MerchantOrderCommentUpdateRequest) UnsetOrderId() {
	o.OrderId.Unset()
}

// GetMerchantComment returns the MerchantComment field value
func (o *MerchantOrderCommentUpdateRequest) GetMerchantComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantComment
}

// GetMerchantCommentOk returns a tuple with the MerchantComment field value
// and a boolean to check if the value has been set.
func (o *MerchantOrderCommentUpdateRequest) GetMerchantCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantComment, true
}

// SetMerchantComment sets field value
func (o *MerchantOrderCommentUpdateRequest) SetMerchantComment(v string) {
	o.MerchantComment = v
}

func (o MerchantOrderCommentUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantOrderCommentUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MerchantOrderNo.IsSet() {
		toSerialize["MerchantOrderNo"] = o.MerchantOrderNo.Get()
	}
	if o.OrderId.IsSet() {
		toSerialize["OrderId"] = o.OrderId.Get()
	}
	toSerialize["MerchantComment"] = o.MerchantComment
	return toSerialize, nil
}

func (o *MerchantOrderCommentUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"MerchantComment",
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

	varMerchantOrderCommentUpdateRequest := _MerchantOrderCommentUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMerchantOrderCommentUpdateRequest)

	if err != nil {
		return err
	}

	*o = MerchantOrderCommentUpdateRequest(varMerchantOrderCommentUpdateRequest)

	return err
}

type NullableMerchantOrderCommentUpdateRequest struct {
	value *MerchantOrderCommentUpdateRequest
	isSet bool
}

func (v NullableMerchantOrderCommentUpdateRequest) Get() *MerchantOrderCommentUpdateRequest {
	return v.value
}

func (v *NullableMerchantOrderCommentUpdateRequest) Set(val *MerchantOrderCommentUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantOrderCommentUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantOrderCommentUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantOrderCommentUpdateRequest(val *MerchantOrderCommentUpdateRequest) *NullableMerchantOrderCommentUpdateRequest {
	return &NullableMerchantOrderCommentUpdateRequest{value: val, isSet: true}
}

func (v NullableMerchantOrderCommentUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantOrderCommentUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


