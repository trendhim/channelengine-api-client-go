/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"encoding/json"
	"fmt"
)

// PurchaseOrderAcknowledgementCode the model 'PurchaseOrderAcknowledgementCode'
type PurchaseOrderAcknowledgementCode string

// List of PurchaseOrderAcknowledgementCode
const (
	PURCHASEORDERACKNOWLEDGEMENTCODE_REJECTED PurchaseOrderAcknowledgementCode = "REJECTED"
	PURCHASEORDERACKNOWLEDGEMENTCODE_ACCEPTED PurchaseOrderAcknowledgementCode = "ACCEPTED"
	PURCHASEORDERACKNOWLEDGEMENTCODE_BACKORDERED PurchaseOrderAcknowledgementCode = "BACKORDERED"
)

// All allowed values of PurchaseOrderAcknowledgementCode enum
var AllowedPurchaseOrderAcknowledgementCodeEnumValues = []PurchaseOrderAcknowledgementCode{
	"REJECTED",
	"ACCEPTED",
	"BACKORDERED",
}

func (v *PurchaseOrderAcknowledgementCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PurchaseOrderAcknowledgementCode(value)
	for _, existing := range AllowedPurchaseOrderAcknowledgementCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PurchaseOrderAcknowledgementCode", value)
}

// NewPurchaseOrderAcknowledgementCodeFromValue returns a pointer to a valid PurchaseOrderAcknowledgementCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPurchaseOrderAcknowledgementCodeFromValue(v string) (*PurchaseOrderAcknowledgementCode, error) {
	ev := PurchaseOrderAcknowledgementCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PurchaseOrderAcknowledgementCode: valid values are %v", v, AllowedPurchaseOrderAcknowledgementCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PurchaseOrderAcknowledgementCode) IsValid() bool {
	for _, existing := range AllowedPurchaseOrderAcknowledgementCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PurchaseOrderAcknowledgementCode value
func (v PurchaseOrderAcknowledgementCode) Ptr() *PurchaseOrderAcknowledgementCode {
	return &v
}

type NullablePurchaseOrderAcknowledgementCode struct {
	value *PurchaseOrderAcknowledgementCode
	isSet bool
}

func (v NullablePurchaseOrderAcknowledgementCode) Get() *PurchaseOrderAcknowledgementCode {
	return v.value
}

func (v *NullablePurchaseOrderAcknowledgementCode) Set(val *PurchaseOrderAcknowledgementCode) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseOrderAcknowledgementCode) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseOrderAcknowledgementCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseOrderAcknowledgementCode(val *PurchaseOrderAcknowledgementCode) *NullablePurchaseOrderAcknowledgementCode {
	return &NullablePurchaseOrderAcknowledgementCode{value: val, isSet: true}
}

func (v NullablePurchaseOrderAcknowledgementCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurchaseOrderAcknowledgementCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

