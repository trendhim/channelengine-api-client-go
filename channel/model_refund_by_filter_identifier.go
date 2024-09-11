/*
ChannelEngine Channel API

ChannelEngine API for channels

API version: 2.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package channel

import (
	"encoding/json"
	"fmt"
)

// RefundByFilterIdentifier the model 'RefundByFilterIdentifier'
type RefundByFilterIdentifier string

// List of RefundByFilterIdentifier
const (
	REFUNDBYFILTERIDENTIFIER_REFUND_ID RefundByFilterIdentifier = "REFUND_ID"
	REFUNDBYFILTERIDENTIFIER_CHANNEL_REFUND_NO RefundByFilterIdentifier = "CHANNEL_REFUND_NO"
	REFUNDBYFILTERIDENTIFIER_MERCHANT_REFUND_NO RefundByFilterIdentifier = "MERCHANT_REFUND_NO"
	REFUNDBYFILTERIDENTIFIER_ORDER_ID RefundByFilterIdentifier = "ORDER_ID"
	REFUNDBYFILTERIDENTIFIER_CHANNEL_ORDER_NO RefundByFilterIdentifier = "CHANNEL_ORDER_NO"
	REFUNDBYFILTERIDENTIFIER_MERCHANT_ORDER_NO RefundByFilterIdentifier = "MERCHANT_ORDER_NO"
	REFUNDBYFILTERIDENTIFIER_RETURN_ID RefundByFilterIdentifier = "RETURN_ID"
	REFUNDBYFILTERIDENTIFIER_CHANNEL_RETURN_NO RefundByFilterIdentifier = "CHANNEL_RETURN_NO"
	REFUNDBYFILTERIDENTIFIER_MERCHANT_RETURN_NO RefundByFilterIdentifier = "MERCHANT_RETURN_NO"
)

// All allowed values of RefundByFilterIdentifier enum
var AllowedRefundByFilterIdentifierEnumValues = []RefundByFilterIdentifier{
	"REFUND_ID",
	"CHANNEL_REFUND_NO",
	"MERCHANT_REFUND_NO",
	"ORDER_ID",
	"CHANNEL_ORDER_NO",
	"MERCHANT_ORDER_NO",
	"RETURN_ID",
	"CHANNEL_RETURN_NO",
	"MERCHANT_RETURN_NO",
}

func (v *RefundByFilterIdentifier) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RefundByFilterIdentifier(value)
	for _, existing := range AllowedRefundByFilterIdentifierEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RefundByFilterIdentifier", value)
}

// NewRefundByFilterIdentifierFromValue returns a pointer to a valid RefundByFilterIdentifier
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRefundByFilterIdentifierFromValue(v string) (*RefundByFilterIdentifier, error) {
	ev := RefundByFilterIdentifier(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RefundByFilterIdentifier: valid values are %v", v, AllowedRefundByFilterIdentifierEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RefundByFilterIdentifier) IsValid() bool {
	for _, existing := range AllowedRefundByFilterIdentifierEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RefundByFilterIdentifier value
func (v RefundByFilterIdentifier) Ptr() *RefundByFilterIdentifier {
	return &v
}

type NullableRefundByFilterIdentifier struct {
	value *RefundByFilterIdentifier
	isSet bool
}

func (v NullableRefundByFilterIdentifier) Get() *RefundByFilterIdentifier {
	return v.value
}

func (v *NullableRefundByFilterIdentifier) Set(val *RefundByFilterIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableRefundByFilterIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableRefundByFilterIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefundByFilterIdentifier(val *RefundByFilterIdentifier) *NullableRefundByFilterIdentifier {
	return &NullableRefundByFilterIdentifier{value: val, isSet: true}
}

func (v NullableRefundByFilterIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefundByFilterIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

