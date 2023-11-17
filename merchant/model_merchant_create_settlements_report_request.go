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

// checks if the MerchantCreateSettlementsReportRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantCreateSettlementsReportRequest{}

// MerchantCreateSettlementsReportRequest struct for MerchantCreateSettlementsReportRequest
type MerchantCreateSettlementsReportRequest struct {
	SettlementIds []int32 `json:"SettlementIds"`
	Type ReportType `json:"Type"`
}

type _MerchantCreateSettlementsReportRequest MerchantCreateSettlementsReportRequest

// NewMerchantCreateSettlementsReportRequest instantiates a new MerchantCreateSettlementsReportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantCreateSettlementsReportRequest(settlementIds []int32, type_ ReportType) *MerchantCreateSettlementsReportRequest {
	this := MerchantCreateSettlementsReportRequest{}
	this.SettlementIds = settlementIds
	this.Type = type_
	return &this
}

// NewMerchantCreateSettlementsReportRequestWithDefaults instantiates a new MerchantCreateSettlementsReportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantCreateSettlementsReportRequestWithDefaults() *MerchantCreateSettlementsReportRequest {
	this := MerchantCreateSettlementsReportRequest{}
	return &this
}

// GetSettlementIds returns the SettlementIds field value
func (o *MerchantCreateSettlementsReportRequest) GetSettlementIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.SettlementIds
}

// GetSettlementIdsOk returns a tuple with the SettlementIds field value
// and a boolean to check if the value has been set.
func (o *MerchantCreateSettlementsReportRequest) GetSettlementIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SettlementIds, true
}

// SetSettlementIds sets field value
func (o *MerchantCreateSettlementsReportRequest) SetSettlementIds(v []int32) {
	o.SettlementIds = v
}

// GetType returns the Type field value
func (o *MerchantCreateSettlementsReportRequest) GetType() ReportType {
	if o == nil {
		var ret ReportType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MerchantCreateSettlementsReportRequest) GetTypeOk() (*ReportType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MerchantCreateSettlementsReportRequest) SetType(v ReportType) {
	o.Type = v
}

func (o MerchantCreateSettlementsReportRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantCreateSettlementsReportRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["SettlementIds"] = o.SettlementIds
	toSerialize["Type"] = o.Type
	return toSerialize, nil
}

func (o *MerchantCreateSettlementsReportRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"SettlementIds",
		"Type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMerchantCreateSettlementsReportRequest := _MerchantCreateSettlementsReportRequest{}

	err = json.Unmarshal(bytes, &varMerchantCreateSettlementsReportRequest)

	if err != nil {
		return err
	}

	*o = MerchantCreateSettlementsReportRequest(varMerchantCreateSettlementsReportRequest)

	return err
}

type NullableMerchantCreateSettlementsReportRequest struct {
	value *MerchantCreateSettlementsReportRequest
	isSet bool
}

func (v NullableMerchantCreateSettlementsReportRequest) Get() *MerchantCreateSettlementsReportRequest {
	return v.value
}

func (v *NullableMerchantCreateSettlementsReportRequest) Set(val *MerchantCreateSettlementsReportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantCreateSettlementsReportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantCreateSettlementsReportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantCreateSettlementsReportRequest(val *MerchantCreateSettlementsReportRequest) *NullableMerchantCreateSettlementsReportRequest {
	return &NullableMerchantCreateSettlementsReportRequest{value: val, isSet: true}
}

func (v NullableMerchantCreateSettlementsReportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantCreateSettlementsReportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


