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

// ReportStatus the model 'ReportStatus'
type ReportStatus string

// List of ReportStatus
const (
	REPORTSTATUS_IN_PROGRESS ReportStatus = "IN_PROGRESS"
	REPORTSTATUS_DONE ReportStatus = "DONE"
	REPORTSTATUS_FAILED ReportStatus = "FAILED"
	REPORTSTATUS_NOT_FOUND ReportStatus = "NOT_FOUND"
)

// All allowed values of ReportStatus enum
var AllowedReportStatusEnumValues = []ReportStatus{
	"IN_PROGRESS",
	"DONE",
	"FAILED",
	"NOT_FOUND",
}

func (v *ReportStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReportStatus(value)
	for _, existing := range AllowedReportStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReportStatus", value)
}

// NewReportStatusFromValue returns a pointer to a valid ReportStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportStatusFromValue(v string) (*ReportStatus, error) {
	ev := ReportStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportStatus: valid values are %v", v, AllowedReportStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportStatus) IsValid() bool {
	for _, existing := range AllowedReportStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReportStatus value
func (v ReportStatus) Ptr() *ReportStatus {
	return &v
}

type NullableReportStatus struct {
	value *ReportStatus
	isSet bool
}

func (v NullableReportStatus) Get() *ReportStatus {
	return v.value
}

func (v *NullableReportStatus) Set(val *ReportStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableReportStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableReportStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportStatus(val *ReportStatus) *NullableReportStatus {
	return &NullableReportStatus{value: val, isSet: true}
}

func (v NullableReportStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

