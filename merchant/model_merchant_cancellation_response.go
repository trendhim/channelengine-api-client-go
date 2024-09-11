/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the MerchantCancellationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantCancellationResponse{}

// MerchantCancellationResponse struct for MerchantCancellationResponse
type MerchantCancellationResponse struct {
	// The unique cancellation identifier used by ChannelEngine.
	Id *int32 `json:"Id,omitempty"`
	// The unique cancellation reference used by the Merchant (sku).
	MerchantCancellationNo string `json:"MerchantCancellationNo"`
	// The unique order reference used by the Merchant.
	MerchantOrderNo string `json:"MerchantOrderNo"`
	// The unique order reference used by the Channel.
	ChannelOrderNo NullableString `json:"ChannelOrderNo,omitempty"`
	Lines []MerchantCancellationLineResponse `json:"Lines"`
	// The date at which the cancellation was created in ChannelEngine.
	CreatedAt *time.Time `json:"CreatedAt,omitempty"`
	// Reason for cancellation (text).
	Reason NullableString `json:"Reason,omitempty"`
	ReasonCode *MancoReason `json:"ReasonCode,omitempty"`
}

type _MerchantCancellationResponse MerchantCancellationResponse

// NewMerchantCancellationResponse instantiates a new MerchantCancellationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantCancellationResponse(merchantCancellationNo string, merchantOrderNo string, lines []MerchantCancellationLineResponse) *MerchantCancellationResponse {
	this := MerchantCancellationResponse{}
	this.MerchantCancellationNo = merchantCancellationNo
	this.MerchantOrderNo = merchantOrderNo
	this.Lines = lines
	return &this
}

// NewMerchantCancellationResponseWithDefaults instantiates a new MerchantCancellationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantCancellationResponseWithDefaults() *MerchantCancellationResponse {
	this := MerchantCancellationResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MerchantCancellationResponse) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCancellationResponse) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MerchantCancellationResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MerchantCancellationResponse) SetId(v int32) {
	o.Id = &v
}

// GetMerchantCancellationNo returns the MerchantCancellationNo field value
func (o *MerchantCancellationResponse) GetMerchantCancellationNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantCancellationNo
}

// GetMerchantCancellationNoOk returns a tuple with the MerchantCancellationNo field value
// and a boolean to check if the value has been set.
func (o *MerchantCancellationResponse) GetMerchantCancellationNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantCancellationNo, true
}

// SetMerchantCancellationNo sets field value
func (o *MerchantCancellationResponse) SetMerchantCancellationNo(v string) {
	o.MerchantCancellationNo = v
}

// GetMerchantOrderNo returns the MerchantOrderNo field value
func (o *MerchantCancellationResponse) GetMerchantOrderNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantOrderNo
}

// GetMerchantOrderNoOk returns a tuple with the MerchantOrderNo field value
// and a boolean to check if the value has been set.
func (o *MerchantCancellationResponse) GetMerchantOrderNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantOrderNo, true
}

// SetMerchantOrderNo sets field value
func (o *MerchantCancellationResponse) SetMerchantOrderNo(v string) {
	o.MerchantOrderNo = v
}

// GetChannelOrderNo returns the ChannelOrderNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantCancellationResponse) GetChannelOrderNo() string {
	if o == nil || IsNil(o.ChannelOrderNo.Get()) {
		var ret string
		return ret
	}
	return *o.ChannelOrderNo.Get()
}

// GetChannelOrderNoOk returns a tuple with the ChannelOrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantCancellationResponse) GetChannelOrderNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelOrderNo.Get(), o.ChannelOrderNo.IsSet()
}

// HasChannelOrderNo returns a boolean if a field has been set.
func (o *MerchantCancellationResponse) HasChannelOrderNo() bool {
	if o != nil && o.ChannelOrderNo.IsSet() {
		return true
	}

	return false
}

// SetChannelOrderNo gets a reference to the given NullableString and assigns it to the ChannelOrderNo field.
func (o *MerchantCancellationResponse) SetChannelOrderNo(v string) {
	o.ChannelOrderNo.Set(&v)
}
// SetChannelOrderNoNil sets the value for ChannelOrderNo to be an explicit nil
func (o *MerchantCancellationResponse) SetChannelOrderNoNil() {
	o.ChannelOrderNo.Set(nil)
}

// UnsetChannelOrderNo ensures that no value is present for ChannelOrderNo, not even an explicit nil
func (o *MerchantCancellationResponse) UnsetChannelOrderNo() {
	o.ChannelOrderNo.Unset()
}

// GetLines returns the Lines field value
func (o *MerchantCancellationResponse) GetLines() []MerchantCancellationLineResponse {
	if o == nil {
		var ret []MerchantCancellationLineResponse
		return ret
	}

	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value
// and a boolean to check if the value has been set.
func (o *MerchantCancellationResponse) GetLinesOk() ([]MerchantCancellationLineResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lines, true
}

// SetLines sets field value
func (o *MerchantCancellationResponse) SetLines(v []MerchantCancellationLineResponse) {
	o.Lines = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *MerchantCancellationResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCancellationResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *MerchantCancellationResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *MerchantCancellationResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantCancellationResponse) GetReason() string {
	if o == nil || IsNil(o.Reason.Get()) {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantCancellationResponse) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *MerchantCancellationResponse) HasReason() bool {
	if o != nil && o.Reason.IsSet() {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableString and assigns it to the Reason field.
func (o *MerchantCancellationResponse) SetReason(v string) {
	o.Reason.Set(&v)
}
// SetReasonNil sets the value for Reason to be an explicit nil
func (o *MerchantCancellationResponse) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil
func (o *MerchantCancellationResponse) UnsetReason() {
	o.Reason.Unset()
}

// GetReasonCode returns the ReasonCode field value if set, zero value otherwise.
func (o *MerchantCancellationResponse) GetReasonCode() MancoReason {
	if o == nil || IsNil(o.ReasonCode) {
		var ret MancoReason
		return ret
	}
	return *o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCancellationResponse) GetReasonCodeOk() (*MancoReason, bool) {
	if o == nil || IsNil(o.ReasonCode) {
		return nil, false
	}
	return o.ReasonCode, true
}

// HasReasonCode returns a boolean if a field has been set.
func (o *MerchantCancellationResponse) HasReasonCode() bool {
	if o != nil && !IsNil(o.ReasonCode) {
		return true
	}

	return false
}

// SetReasonCode gets a reference to the given MancoReason and assigns it to the ReasonCode field.
func (o *MerchantCancellationResponse) SetReasonCode(v MancoReason) {
	o.ReasonCode = &v
}

func (o MerchantCancellationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantCancellationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	toSerialize["MerchantCancellationNo"] = o.MerchantCancellationNo
	toSerialize["MerchantOrderNo"] = o.MerchantOrderNo
	if o.ChannelOrderNo.IsSet() {
		toSerialize["ChannelOrderNo"] = o.ChannelOrderNo.Get()
	}
	toSerialize["Lines"] = o.Lines
	if !IsNil(o.CreatedAt) {
		toSerialize["CreatedAt"] = o.CreatedAt
	}
	if o.Reason.IsSet() {
		toSerialize["Reason"] = o.Reason.Get()
	}
	if !IsNil(o.ReasonCode) {
		toSerialize["ReasonCode"] = o.ReasonCode
	}
	return toSerialize, nil
}

func (o *MerchantCancellationResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"MerchantCancellationNo",
		"MerchantOrderNo",
		"Lines",
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

	varMerchantCancellationResponse := _MerchantCancellationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMerchantCancellationResponse)

	if err != nil {
		return err
	}

	*o = MerchantCancellationResponse(varMerchantCancellationResponse)

	return err
}

type NullableMerchantCancellationResponse struct {
	value *MerchantCancellationResponse
	isSet bool
}

func (v NullableMerchantCancellationResponse) Get() *MerchantCancellationResponse {
	return v.value
}

func (v *NullableMerchantCancellationResponse) Set(val *MerchantCancellationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantCancellationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantCancellationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantCancellationResponse(val *MerchantCancellationResponse) *NullableMerchantCancellationResponse {
	return &NullableMerchantCancellationResponse{value: val, isSet: true}
}

func (v NullableMerchantCancellationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantCancellationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


