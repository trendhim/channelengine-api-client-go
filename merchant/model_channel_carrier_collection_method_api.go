/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"encoding/json"
	"fmt"
)

// ChannelCarrierCollectionMethodApi the model 'ChannelCarrierCollectionMethodApi'
type ChannelCarrierCollectionMethodApi string

// List of ChannelCarrierCollectionMethodApi
const (
	CHANNELCARRIERCOLLECTIONMETHODAPI_DROP_OFF ChannelCarrierCollectionMethodApi = "DROP_OFF"
	CHANNELCARRIERCOLLECTIONMETHODAPI_PICK_UP ChannelCarrierCollectionMethodApi = "PICK_UP"
)

// All allowed values of ChannelCarrierCollectionMethodApi enum
var AllowedChannelCarrierCollectionMethodApiEnumValues = []ChannelCarrierCollectionMethodApi{
	"DROP_OFF",
	"PICK_UP",
}

func (v *ChannelCarrierCollectionMethodApi) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ChannelCarrierCollectionMethodApi(value)
	for _, existing := range AllowedChannelCarrierCollectionMethodApiEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ChannelCarrierCollectionMethodApi", value)
}

// NewChannelCarrierCollectionMethodApiFromValue returns a pointer to a valid ChannelCarrierCollectionMethodApi
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChannelCarrierCollectionMethodApiFromValue(v string) (*ChannelCarrierCollectionMethodApi, error) {
	ev := ChannelCarrierCollectionMethodApi(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChannelCarrierCollectionMethodApi: valid values are %v", v, AllowedChannelCarrierCollectionMethodApiEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChannelCarrierCollectionMethodApi) IsValid() bool {
	for _, existing := range AllowedChannelCarrierCollectionMethodApiEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChannelCarrierCollectionMethodApi value
func (v ChannelCarrierCollectionMethodApi) Ptr() *ChannelCarrierCollectionMethodApi {
	return &v
}

type NullableChannelCarrierCollectionMethodApi struct {
	value *ChannelCarrierCollectionMethodApi
	isSet bool
}

func (v NullableChannelCarrierCollectionMethodApi) Get() *ChannelCarrierCollectionMethodApi {
	return v.value
}

func (v *NullableChannelCarrierCollectionMethodApi) Set(val *ChannelCarrierCollectionMethodApi) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelCarrierCollectionMethodApi) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelCarrierCollectionMethodApi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelCarrierCollectionMethodApi(val *ChannelCarrierCollectionMethodApi) *NullableChannelCarrierCollectionMethodApi {
	return &NullableChannelCarrierCollectionMethodApi{value: val, isSet: true}
}

func (v NullableChannelCarrierCollectionMethodApi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelCarrierCollectionMethodApi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
