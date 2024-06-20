/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.14.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"encoding/json"
)

// checks if the SettingsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingsResponse{}

// SettingsResponse struct for SettingsResponse
type SettingsResponse struct {
	Shipment *ShipmentSettingsResponse `json:"Shipment,omitempty"`
	Advanced *AdvanceSettingsResponse `json:"Advanced,omitempty"`
}

// NewSettingsResponse instantiates a new SettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsResponse() *SettingsResponse {
	this := SettingsResponse{}
	return &this
}

// NewSettingsResponseWithDefaults instantiates a new SettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsResponseWithDefaults() *SettingsResponse {
	this := SettingsResponse{}
	return &this
}

// GetShipment returns the Shipment field value if set, zero value otherwise.
func (o *SettingsResponse) GetShipment() ShipmentSettingsResponse {
	if o == nil || IsNil(o.Shipment) {
		var ret ShipmentSettingsResponse
		return ret
	}
	return *o.Shipment
}

// GetShipmentOk returns a tuple with the Shipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsResponse) GetShipmentOk() (*ShipmentSettingsResponse, bool) {
	if o == nil || IsNil(o.Shipment) {
		return nil, false
	}
	return o.Shipment, true
}

// HasShipment returns a boolean if a field has been set.
func (o *SettingsResponse) HasShipment() bool {
	if o != nil && !IsNil(o.Shipment) {
		return true
	}

	return false
}

// SetShipment gets a reference to the given ShipmentSettingsResponse and assigns it to the Shipment field.
func (o *SettingsResponse) SetShipment(v ShipmentSettingsResponse) {
	o.Shipment = &v
}

// GetAdvanced returns the Advanced field value if set, zero value otherwise.
func (o *SettingsResponse) GetAdvanced() AdvanceSettingsResponse {
	if o == nil || IsNil(o.Advanced) {
		var ret AdvanceSettingsResponse
		return ret
	}
	return *o.Advanced
}

// GetAdvancedOk returns a tuple with the Advanced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsResponse) GetAdvancedOk() (*AdvanceSettingsResponse, bool) {
	if o == nil || IsNil(o.Advanced) {
		return nil, false
	}
	return o.Advanced, true
}

// HasAdvanced returns a boolean if a field has been set.
func (o *SettingsResponse) HasAdvanced() bool {
	if o != nil && !IsNil(o.Advanced) {
		return true
	}

	return false
}

// SetAdvanced gets a reference to the given AdvanceSettingsResponse and assigns it to the Advanced field.
func (o *SettingsResponse) SetAdvanced(v AdvanceSettingsResponse) {
	o.Advanced = &v
}

func (o SettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Shipment) {
		toSerialize["Shipment"] = o.Shipment
	}
	if !IsNil(o.Advanced) {
		toSerialize["Advanced"] = o.Advanced
	}
	return toSerialize, nil
}

type NullableSettingsResponse struct {
	value *SettingsResponse
	isSet bool
}

func (v NullableSettingsResponse) Get() *SettingsResponse {
	return v.value
}

func (v *NullableSettingsResponse) Set(val *SettingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsResponse(val *SettingsResponse) *NullableSettingsResponse {
	return &NullableSettingsResponse{value: val, isSet: true}
}

func (v NullableSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


