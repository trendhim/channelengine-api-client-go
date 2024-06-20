/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.14.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"encoding/json"
	"time"
)

// checks if the MerchantAcknowledgePurchaseOrderLine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantAcknowledgePurchaseOrderLine{}

// MerchantAcknowledgePurchaseOrderLine struct for MerchantAcknowledgePurchaseOrderLine
type MerchantAcknowledgePurchaseOrderLine struct {
	OrderLineIdentifier NullableString `json:"OrderLineIdentifier,omitempty"`
	AcknowledgementCode *PurchaseOrderAcknowledgementCode `json:"AcknowledgementCode,omitempty"`
	AcknowledgedQuantity *int32 `json:"AcknowledgedQuantity,omitempty"`
	RejectionReason *PurchaseOrderRejectionReason `json:"RejectionReason,omitempty"`
	ScheduledShipDate NullableTime `json:"ScheduledShipDate,omitempty"`
	ScheduledDeliveryDate NullableTime `json:"ScheduledDeliveryDate,omitempty"`
}

// NewMerchantAcknowledgePurchaseOrderLine instantiates a new MerchantAcknowledgePurchaseOrderLine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantAcknowledgePurchaseOrderLine() *MerchantAcknowledgePurchaseOrderLine {
	this := MerchantAcknowledgePurchaseOrderLine{}
	return &this
}

// NewMerchantAcknowledgePurchaseOrderLineWithDefaults instantiates a new MerchantAcknowledgePurchaseOrderLine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantAcknowledgePurchaseOrderLineWithDefaults() *MerchantAcknowledgePurchaseOrderLine {
	this := MerchantAcknowledgePurchaseOrderLine{}
	return &this
}

// GetOrderLineIdentifier returns the OrderLineIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAcknowledgePurchaseOrderLine) GetOrderLineIdentifier() string {
	if o == nil || IsNil(o.OrderLineIdentifier.Get()) {
		var ret string
		return ret
	}
	return *o.OrderLineIdentifier.Get()
}

// GetOrderLineIdentifierOk returns a tuple with the OrderLineIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAcknowledgePurchaseOrderLine) GetOrderLineIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrderLineIdentifier.Get(), o.OrderLineIdentifier.IsSet()
}

// HasOrderLineIdentifier returns a boolean if a field has been set.
func (o *MerchantAcknowledgePurchaseOrderLine) HasOrderLineIdentifier() bool {
	if o != nil && o.OrderLineIdentifier.IsSet() {
		return true
	}

	return false
}

// SetOrderLineIdentifier gets a reference to the given NullableString and assigns it to the OrderLineIdentifier field.
func (o *MerchantAcknowledgePurchaseOrderLine) SetOrderLineIdentifier(v string) {
	o.OrderLineIdentifier.Set(&v)
}
// SetOrderLineIdentifierNil sets the value for OrderLineIdentifier to be an explicit nil
func (o *MerchantAcknowledgePurchaseOrderLine) SetOrderLineIdentifierNil() {
	o.OrderLineIdentifier.Set(nil)
}

// UnsetOrderLineIdentifier ensures that no value is present for OrderLineIdentifier, not even an explicit nil
func (o *MerchantAcknowledgePurchaseOrderLine) UnsetOrderLineIdentifier() {
	o.OrderLineIdentifier.Unset()
}

// GetAcknowledgementCode returns the AcknowledgementCode field value if set, zero value otherwise.
func (o *MerchantAcknowledgePurchaseOrderLine) GetAcknowledgementCode() PurchaseOrderAcknowledgementCode {
	if o == nil || IsNil(o.AcknowledgementCode) {
		var ret PurchaseOrderAcknowledgementCode
		return ret
	}
	return *o.AcknowledgementCode
}

// GetAcknowledgementCodeOk returns a tuple with the AcknowledgementCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantAcknowledgePurchaseOrderLine) GetAcknowledgementCodeOk() (*PurchaseOrderAcknowledgementCode, bool) {
	if o == nil || IsNil(o.AcknowledgementCode) {
		return nil, false
	}
	return o.AcknowledgementCode, true
}

// HasAcknowledgementCode returns a boolean if a field has been set.
func (o *MerchantAcknowledgePurchaseOrderLine) HasAcknowledgementCode() bool {
	if o != nil && !IsNil(o.AcknowledgementCode) {
		return true
	}

	return false
}

// SetAcknowledgementCode gets a reference to the given PurchaseOrderAcknowledgementCode and assigns it to the AcknowledgementCode field.
func (o *MerchantAcknowledgePurchaseOrderLine) SetAcknowledgementCode(v PurchaseOrderAcknowledgementCode) {
	o.AcknowledgementCode = &v
}

// GetAcknowledgedQuantity returns the AcknowledgedQuantity field value if set, zero value otherwise.
func (o *MerchantAcknowledgePurchaseOrderLine) GetAcknowledgedQuantity() int32 {
	if o == nil || IsNil(o.AcknowledgedQuantity) {
		var ret int32
		return ret
	}
	return *o.AcknowledgedQuantity
}

// GetAcknowledgedQuantityOk returns a tuple with the AcknowledgedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantAcknowledgePurchaseOrderLine) GetAcknowledgedQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.AcknowledgedQuantity) {
		return nil, false
	}
	return o.AcknowledgedQuantity, true
}

// HasAcknowledgedQuantity returns a boolean if a field has been set.
func (o *MerchantAcknowledgePurchaseOrderLine) HasAcknowledgedQuantity() bool {
	if o != nil && !IsNil(o.AcknowledgedQuantity) {
		return true
	}

	return false
}

// SetAcknowledgedQuantity gets a reference to the given int32 and assigns it to the AcknowledgedQuantity field.
func (o *MerchantAcknowledgePurchaseOrderLine) SetAcknowledgedQuantity(v int32) {
	o.AcknowledgedQuantity = &v
}

// GetRejectionReason returns the RejectionReason field value if set, zero value otherwise.
func (o *MerchantAcknowledgePurchaseOrderLine) GetRejectionReason() PurchaseOrderRejectionReason {
	if o == nil || IsNil(o.RejectionReason) {
		var ret PurchaseOrderRejectionReason
		return ret
	}
	return *o.RejectionReason
}

// GetRejectionReasonOk returns a tuple with the RejectionReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantAcknowledgePurchaseOrderLine) GetRejectionReasonOk() (*PurchaseOrderRejectionReason, bool) {
	if o == nil || IsNil(o.RejectionReason) {
		return nil, false
	}
	return o.RejectionReason, true
}

// HasRejectionReason returns a boolean if a field has been set.
func (o *MerchantAcknowledgePurchaseOrderLine) HasRejectionReason() bool {
	if o != nil && !IsNil(o.RejectionReason) {
		return true
	}

	return false
}

// SetRejectionReason gets a reference to the given PurchaseOrderRejectionReason and assigns it to the RejectionReason field.
func (o *MerchantAcknowledgePurchaseOrderLine) SetRejectionReason(v PurchaseOrderRejectionReason) {
	o.RejectionReason = &v
}

// GetScheduledShipDate returns the ScheduledShipDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAcknowledgePurchaseOrderLine) GetScheduledShipDate() time.Time {
	if o == nil || IsNil(o.ScheduledShipDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledShipDate.Get()
}

// GetScheduledShipDateOk returns a tuple with the ScheduledShipDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAcknowledgePurchaseOrderLine) GetScheduledShipDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduledShipDate.Get(), o.ScheduledShipDate.IsSet()
}

// HasScheduledShipDate returns a boolean if a field has been set.
func (o *MerchantAcknowledgePurchaseOrderLine) HasScheduledShipDate() bool {
	if o != nil && o.ScheduledShipDate.IsSet() {
		return true
	}

	return false
}

// SetScheduledShipDate gets a reference to the given NullableTime and assigns it to the ScheduledShipDate field.
func (o *MerchantAcknowledgePurchaseOrderLine) SetScheduledShipDate(v time.Time) {
	o.ScheduledShipDate.Set(&v)
}
// SetScheduledShipDateNil sets the value for ScheduledShipDate to be an explicit nil
func (o *MerchantAcknowledgePurchaseOrderLine) SetScheduledShipDateNil() {
	o.ScheduledShipDate.Set(nil)
}

// UnsetScheduledShipDate ensures that no value is present for ScheduledShipDate, not even an explicit nil
func (o *MerchantAcknowledgePurchaseOrderLine) UnsetScheduledShipDate() {
	o.ScheduledShipDate.Unset()
}

// GetScheduledDeliveryDate returns the ScheduledDeliveryDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAcknowledgePurchaseOrderLine) GetScheduledDeliveryDate() time.Time {
	if o == nil || IsNil(o.ScheduledDeliveryDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledDeliveryDate.Get()
}

// GetScheduledDeliveryDateOk returns a tuple with the ScheduledDeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAcknowledgePurchaseOrderLine) GetScheduledDeliveryDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduledDeliveryDate.Get(), o.ScheduledDeliveryDate.IsSet()
}

// HasScheduledDeliveryDate returns a boolean if a field has been set.
func (o *MerchantAcknowledgePurchaseOrderLine) HasScheduledDeliveryDate() bool {
	if o != nil && o.ScheduledDeliveryDate.IsSet() {
		return true
	}

	return false
}

// SetScheduledDeliveryDate gets a reference to the given NullableTime and assigns it to the ScheduledDeliveryDate field.
func (o *MerchantAcknowledgePurchaseOrderLine) SetScheduledDeliveryDate(v time.Time) {
	o.ScheduledDeliveryDate.Set(&v)
}
// SetScheduledDeliveryDateNil sets the value for ScheduledDeliveryDate to be an explicit nil
func (o *MerchantAcknowledgePurchaseOrderLine) SetScheduledDeliveryDateNil() {
	o.ScheduledDeliveryDate.Set(nil)
}

// UnsetScheduledDeliveryDate ensures that no value is present for ScheduledDeliveryDate, not even an explicit nil
func (o *MerchantAcknowledgePurchaseOrderLine) UnsetScheduledDeliveryDate() {
	o.ScheduledDeliveryDate.Unset()
}

func (o MerchantAcknowledgePurchaseOrderLine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantAcknowledgePurchaseOrderLine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OrderLineIdentifier.IsSet() {
		toSerialize["OrderLineIdentifier"] = o.OrderLineIdentifier.Get()
	}
	if !IsNil(o.AcknowledgementCode) {
		toSerialize["AcknowledgementCode"] = o.AcknowledgementCode
	}
	if !IsNil(o.AcknowledgedQuantity) {
		toSerialize["AcknowledgedQuantity"] = o.AcknowledgedQuantity
	}
	if !IsNil(o.RejectionReason) {
		toSerialize["RejectionReason"] = o.RejectionReason
	}
	if o.ScheduledShipDate.IsSet() {
		toSerialize["ScheduledShipDate"] = o.ScheduledShipDate.Get()
	}
	if o.ScheduledDeliveryDate.IsSet() {
		toSerialize["ScheduledDeliveryDate"] = o.ScheduledDeliveryDate.Get()
	}
	return toSerialize, nil
}

type NullableMerchantAcknowledgePurchaseOrderLine struct {
	value *MerchantAcknowledgePurchaseOrderLine
	isSet bool
}

func (v NullableMerchantAcknowledgePurchaseOrderLine) Get() *MerchantAcknowledgePurchaseOrderLine {
	return v.value
}

func (v *NullableMerchantAcknowledgePurchaseOrderLine) Set(val *MerchantAcknowledgePurchaseOrderLine) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantAcknowledgePurchaseOrderLine) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantAcknowledgePurchaseOrderLine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantAcknowledgePurchaseOrderLine(val *MerchantAcknowledgePurchaseOrderLine) *NullableMerchantAcknowledgePurchaseOrderLine {
	return &NullableMerchantAcknowledgePurchaseOrderLine{value: val, isSet: true}
}

func (v NullableMerchantAcknowledgePurchaseOrderLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantAcknowledgePurchaseOrderLine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


