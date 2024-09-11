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
)

// checks if the MerchantReturnResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantReturnResponse{}

// MerchantReturnResponse struct for MerchantReturnResponse
type MerchantReturnResponse struct {
	// The unique order reference used by the Merchant.
	MerchantOrderNo NullableString `json:"MerchantOrderNo,omitempty"`
	// The unique order reference used by the Channel.
	ChannelOrderNo NullableString `json:"ChannelOrderNo,omitempty"`
	// The id of the channel.
	ChannelId NullableInt32 `json:"ChannelId,omitempty"`
	// The id of the Global Channel.
	GlobalChannelId NullableInt32 `json:"GlobalChannelId,omitempty"`
	// The name of the Global Channel.
	GlobalChannelName NullableString `json:"GlobalChannelName,omitempty"`
	Lines []MerchantReturnLineResponse `json:"Lines,omitempty"`
	// The date at which the return was created in ChannelEngine.
	CreatedAt *time.Time `json:"CreatedAt,omitempty"`
	// The date at which the return was last modified in ChannelEngine.
	UpdatedAt *time.Time `json:"UpdatedAt,omitempty"`
	// The unique return reference used by the Merchant, will be empty in case of a channel or unacknowledged return.
	MerchantReturnNo NullableString `json:"MerchantReturnNo,omitempty"`
	// The unique return reference used by the Channel, will be empty in case of a merchant return.
	ChannelReturnNo NullableString `json:"ChannelReturnNo,omitempty"`
	Status *ReturnStatus `json:"Status,omitempty"`
	// Date of acknowledgement
	AcknowledgedDate NullableTime `json:"AcknowledgedDate,omitempty"`
	// The unique return reference used by ChannelEngine.
	Id *int32 `json:"Id,omitempty"`
	Reason *ReturnReason `json:"Reason,omitempty"`
	// Optional. Comment of customer on the (reason of) the return.
	CustomerComment NullableString `json:"CustomerComment,omitempty"`
	// Optional. Comment of merchant on the return.
	MerchantComment NullableString `json:"MerchantComment,omitempty"`
	// Refund amount incl. VAT.
	RefundInclVat *float32 `json:"RefundInclVat,omitempty"`
	// Refund amount excl. VAT.
	RefundExclVat *float32 `json:"RefundExclVat,omitempty"`
	// The date at which the return was originally created in the source system (if available).
	ReturnDate NullableTime `json:"ReturnDate,omitempty"`
	// Extra data on the return. Each item must have an unqiue key
	ExtraData map[string]string `json:"ExtraData,omitempty"`
}

// NewMerchantReturnResponse instantiates a new MerchantReturnResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantReturnResponse() *MerchantReturnResponse {
	this := MerchantReturnResponse{}
	return &this
}

// NewMerchantReturnResponseWithDefaults instantiates a new MerchantReturnResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantReturnResponseWithDefaults() *MerchantReturnResponse {
	this := MerchantReturnResponse{}
	return &this
}

// GetMerchantOrderNo returns the MerchantOrderNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantReturnResponse) GetMerchantOrderNo() string {
	if o == nil || IsNil(o.MerchantOrderNo.Get()) {
		var ret string
		return ret
	}
	return *o.MerchantOrderNo.Get()
}

// GetMerchantOrderNoOk returns a tuple with the MerchantOrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantReturnResponse) GetMerchantOrderNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MerchantOrderNo.Get(), o.MerchantOrderNo.IsSet()
}

// HasMerchantOrderNo returns a boolean if a field has been set.
func (o *MerchantReturnResponse) HasMerchantOrderNo() bool {
	if o != nil && o.MerchantOrderNo.IsSet() {
		return true
	}

	return false
}

// SetMerchantOrderNo gets a reference to the given NullableString and assigns it to the MerchantOrderNo field.
func (o *MerchantReturnResponse) SetMerchantOrderNo(v string) {
	o.MerchantOrderNo.Set(&v)
}
// SetMerchantOrderNoNil sets the value for MerchantOrderNo to be an explicit nil
func (o *MerchantReturnResponse) SetMerchantOrderNoNil() {
	o.MerchantOrderNo.Set(nil)
}

// UnsetMerchantOrderNo ensures that no value is present for MerchantOrderNo, not even an explicit nil
func (o *MerchantReturnResponse) UnsetMerchantOrderNo() {
	o.MerchantOrderNo.Unset()
}

// GetChannelOrderNo returns the ChannelOrderNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantReturnResponse) GetChannelOrderNo() string {
	if o == nil || IsNil(o.ChannelOrderNo.Get()) {
		var ret string
		return ret
	}
	return *o.ChannelOrderNo.Get()
}

// GetChannelOrderNoOk returns a tuple with the ChannelOrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantReturnResponse) GetChannelOrderNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelOrderNo.Get(), o.ChannelOrderNo.IsSet()
}

// HasChannelOrderNo returns a boolean if a field has been set.
func (o *MerchantReturnResponse) HasChannelOrderNo() bool {
	if o != nil && o.ChannelOrderNo.IsSet() {
		return true
	}

	return false
}

// SetChannelOrderNo gets a reference to the given NullableString and assigns it to the ChannelOrderNo field.
func (o *MerchantReturnResponse) SetChannelOrderNo(v string) {
	o.ChannelOrderNo.Set(&v)
}
// SetChannelOrderNoNil sets the value for ChannelOrderNo to be an explicit nil
func (o *MerchantReturnResponse) SetChannelOrderNoNil() {
	o.ChannelOrderNo.Set(nil)
}

// UnsetChannelOrderNo ensures that no value is present for ChannelOrderNo, not even an explicit nil
func (o *MerchantReturnResponse) UnsetChannelOrderNo() {
	o.ChannelOrderNo.Unset()
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantReturnResponse) GetChannelId() int32 {
	if o == nil || IsNil(o.ChannelId.Get()) {
		var ret int32
		return ret
	}
	return *o.ChannelId.Get()
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantReturnResponse) GetChannelIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelId.Get(), o.ChannelId.IsSet()
}

// HasChannelId returns a boolean if a field has been set.
func (o *MerchantReturnResponse) HasChannelId() bool {
	if o != nil && o.ChannelId.IsSet() {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given NullableInt32 and assigns it to the ChannelId field.
func (o *MerchantReturnResponse) SetChannelId(v int32) {
	o.ChannelId.Set(&v)
}
// SetChannelIdNil sets the value for ChannelId to be an explicit nil
func (o *MerchantReturnResponse) SetChannelIdNil() {
	o.ChannelId.Set(nil)
}

// UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
func (o *MerchantReturnResponse) UnsetChannelId() {
	o.ChannelId.Unset()
}

// GetGlobalChannelId returns the GlobalChannelId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantReturnResponse) GetGlobalChannelId() int32 {
	if o == nil || IsNil(o.GlobalChannelId.Get()) {
		var ret int32
		return ret
	}
	return *o.GlobalChannelId.Get()
}

// GetGlobalChannelIdOk returns a tuple with the GlobalChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantReturnResponse) GetGlobalChannelIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.GlobalChannelId.Get(), o.GlobalChannelId.IsSet()
}

// HasGlobalChannelId returns a boolean if a field has been set.
func (o *MerchantReturnResponse) HasGlobalChannelId() bool {
	if o != nil && o.GlobalChannelId.IsSet() {
		return true
	}

	return false
}

// SetGlobalChannelId gets a reference to the given NullableInt32 and assigns it to the GlobalChannelId field.
func (o *MerchantReturnResponse) SetGlobalChannelId(v int32) {
	o.GlobalChannelId.Set(&v)
}
// SetGlobalChannelIdNil sets the value for GlobalChannelId to be an explicit nil
func (o *MerchantReturnResponse) SetGlobalChannelIdNil() {
	o.GlobalChannelId.Set(nil)
}

// UnsetGlobalChannelId ensures that no value is present for GlobalChannelId, not even an explicit nil
func (o *MerchantReturnResponse) UnsetGlobalChannelId() {
	o.GlobalChannelId.Unset()
}

// GetGlobalChannelName returns the GlobalChannelName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantReturnResponse) GetGlobalChannelName() string {
	if o == nil || IsNil(o.GlobalChannelName.Get()) {
		var ret string
		return ret
	}
	return *o.GlobalChannelName.Get()
}

// GetGlobalChannelNameOk returns a tuple with the GlobalChannelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantReturnResponse) GetGlobalChannelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GlobalChannelName.Get(), o.GlobalChannelName.IsSet()
}

// HasGlobalChannelName returns a boolean if a field has been set.
func (o *MerchantReturnResponse) HasGlobalChannelName() bool {
	if o != nil && o.GlobalChannelName.IsSet() {
		return true
	}

	return false
}

// SetGlobalChannelName gets a reference to the given NullableString and assigns it to the GlobalChannelName field.
func (o *MerchantReturnResponse) SetGlobalChannelName(v string) {
	o.GlobalChannelName.Set(&v)
}
// SetGlobalChannelNameNil sets the value for GlobalChannelName to be an explicit nil
func (o *MerchantReturnResponse) SetGlobalChannelNameNil() {
	o.GlobalChannelName.Set(nil)
}

// UnsetGlobalChannelName ensures that no value is present for GlobalChannelName, not even an explicit nil
func (o *MerchantReturnResponse) UnsetGlobalChannelName() {
	o.GlobalChannelName.Unset()
}

// GetLines returns the Lines field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantReturnResponse) GetLines() []MerchantReturnLineResponse {
	if o == nil {
		var ret []MerchantReturnLineResponse
		return ret
	}
	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantReturnResponse) GetLinesOk() ([]MerchantReturnLineResponse, bool) {
	if o == nil || IsNil(o.Lines) {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *MerchantReturnResponse) HasLines() bool {
	if o != nil && !IsNil(o.Lines) {
		return true
	}

	return false
}

// SetLines gets a reference to the given []MerchantReturnLineResponse and assigns it to the Lines field.
func (o *MerchantReturnResponse) SetLines(v []MerchantReturnLineResponse) {
	o.Lines = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *MerchantReturnResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantReturnResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *MerchantReturnResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *MerchantReturnResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *MerchantReturnResponse) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantReturnResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *MerchantReturnResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *MerchantReturnResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetMerchantReturnNo returns the MerchantReturnNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantReturnResponse) GetMerchantReturnNo() string {
	if o == nil || IsNil(o.MerchantReturnNo.Get()) {
		var ret string
		return ret
	}
	return *o.MerchantReturnNo.Get()
}

// GetMerchantReturnNoOk returns a tuple with the MerchantReturnNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantReturnResponse) GetMerchantReturnNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MerchantReturnNo.Get(), o.MerchantReturnNo.IsSet()
}

// HasMerchantReturnNo returns a boolean if a field has been set.
func (o *MerchantReturnResponse) HasMerchantReturnNo() bool {
	if o != nil && o.MerchantReturnNo.IsSet() {
		return true
	}

	return false
}

// SetMerchantReturnNo gets a reference to the given NullableString and assigns it to the MerchantReturnNo field.
func (o *MerchantReturnResponse) SetMerchantReturnNo(v string) {
	o.MerchantReturnNo.Set(&v)
}
// SetMerchantReturnNoNil sets the value for MerchantReturnNo to be an explicit nil
func (o *MerchantReturnResponse) SetMerchantReturnNoNil() {
	o.MerchantReturnNo.Set(nil)
}

// UnsetMerchantReturnNo ensures that no value is present for MerchantReturnNo, not even an explicit nil
func (o *MerchantReturnResponse) UnsetMerchantReturnNo() {
	o.MerchantReturnNo.Unset()
}

// GetChannelReturnNo returns the ChannelReturnNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantReturnResponse) GetChannelReturnNo() string {
	if o == nil || IsNil(o.ChannelReturnNo.Get()) {
		var ret string
		return ret
	}
	return *o.ChannelReturnNo.Get()
}

// GetChannelReturnNoOk returns a tuple with the ChannelReturnNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantReturnResponse) GetChannelReturnNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelReturnNo.Get(), o.ChannelReturnNo.IsSet()
}

// HasChannelReturnNo returns a boolean if a field has been set.
func (o *MerchantReturnResponse) HasChannelReturnNo() bool {
	if o != nil && o.ChannelReturnNo.IsSet() {
		return true
	}

	return false
}

// SetChannelReturnNo gets a reference to the given NullableString and assigns it to the ChannelReturnNo field.
func (o *MerchantReturnResponse) SetChannelReturnNo(v string) {
	o.ChannelReturnNo.Set(&v)
}
// SetChannelReturnNoNil sets the value for ChannelReturnNo to be an explicit nil
func (o *MerchantReturnResponse) SetChannelReturnNoNil() {
	o.ChannelReturnNo.Set(nil)
}

// UnsetChannelReturnNo ensures that no value is present for ChannelReturnNo, not even an explicit nil
func (o *MerchantReturnResponse) UnsetChannelReturnNo() {
	o.ChannelReturnNo.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MerchantReturnResponse) GetStatus() ReturnStatus {
	if o == nil || IsNil(o.Status) {
		var ret ReturnStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantReturnResponse) GetStatusOk() (*ReturnStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MerchantReturnResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ReturnStatus and assigns it to the Status field.
func (o *MerchantReturnResponse) SetStatus(v ReturnStatus) {
	o.Status = &v
}

// GetAcknowledgedDate returns the AcknowledgedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantReturnResponse) GetAcknowledgedDate() time.Time {
	if o == nil || IsNil(o.AcknowledgedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.AcknowledgedDate.Get()
}

// GetAcknowledgedDateOk returns a tuple with the AcknowledgedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantReturnResponse) GetAcknowledgedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcknowledgedDate.Get(), o.AcknowledgedDate.IsSet()
}

// HasAcknowledgedDate returns a boolean if a field has been set.
func (o *MerchantReturnResponse) HasAcknowledgedDate() bool {
	if o != nil && o.AcknowledgedDate.IsSet() {
		return true
	}

	return false
}

// SetAcknowledgedDate gets a reference to the given NullableTime and assigns it to the AcknowledgedDate field.
func (o *MerchantReturnResponse) SetAcknowledgedDate(v time.Time) {
	o.AcknowledgedDate.Set(&v)
}
// SetAcknowledgedDateNil sets the value for AcknowledgedDate to be an explicit nil
func (o *MerchantReturnResponse) SetAcknowledgedDateNil() {
	o.AcknowledgedDate.Set(nil)
}

// UnsetAcknowledgedDate ensures that no value is present for AcknowledgedDate, not even an explicit nil
func (o *MerchantReturnResponse) UnsetAcknowledgedDate() {
	o.AcknowledgedDate.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MerchantReturnResponse) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantReturnResponse) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MerchantReturnResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MerchantReturnResponse) SetId(v int32) {
	o.Id = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *MerchantReturnResponse) GetReason() ReturnReason {
	if o == nil || IsNil(o.Reason) {
		var ret ReturnReason
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantReturnResponse) GetReasonOk() (*ReturnReason, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *MerchantReturnResponse) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given ReturnReason and assigns it to the Reason field.
func (o *MerchantReturnResponse) SetReason(v ReturnReason) {
	o.Reason = &v
}

// GetCustomerComment returns the CustomerComment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantReturnResponse) GetCustomerComment() string {
	if o == nil || IsNil(o.CustomerComment.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerComment.Get()
}

// GetCustomerCommentOk returns a tuple with the CustomerComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantReturnResponse) GetCustomerCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerComment.Get(), o.CustomerComment.IsSet()
}

// HasCustomerComment returns a boolean if a field has been set.
func (o *MerchantReturnResponse) HasCustomerComment() bool {
	if o != nil && o.CustomerComment.IsSet() {
		return true
	}

	return false
}

// SetCustomerComment gets a reference to the given NullableString and assigns it to the CustomerComment field.
func (o *MerchantReturnResponse) SetCustomerComment(v string) {
	o.CustomerComment.Set(&v)
}
// SetCustomerCommentNil sets the value for CustomerComment to be an explicit nil
func (o *MerchantReturnResponse) SetCustomerCommentNil() {
	o.CustomerComment.Set(nil)
}

// UnsetCustomerComment ensures that no value is present for CustomerComment, not even an explicit nil
func (o *MerchantReturnResponse) UnsetCustomerComment() {
	o.CustomerComment.Unset()
}

// GetMerchantComment returns the MerchantComment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantReturnResponse) GetMerchantComment() string {
	if o == nil || IsNil(o.MerchantComment.Get()) {
		var ret string
		return ret
	}
	return *o.MerchantComment.Get()
}

// GetMerchantCommentOk returns a tuple with the MerchantComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantReturnResponse) GetMerchantCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MerchantComment.Get(), o.MerchantComment.IsSet()
}

// HasMerchantComment returns a boolean if a field has been set.
func (o *MerchantReturnResponse) HasMerchantComment() bool {
	if o != nil && o.MerchantComment.IsSet() {
		return true
	}

	return false
}

// SetMerchantComment gets a reference to the given NullableString and assigns it to the MerchantComment field.
func (o *MerchantReturnResponse) SetMerchantComment(v string) {
	o.MerchantComment.Set(&v)
}
// SetMerchantCommentNil sets the value for MerchantComment to be an explicit nil
func (o *MerchantReturnResponse) SetMerchantCommentNil() {
	o.MerchantComment.Set(nil)
}

// UnsetMerchantComment ensures that no value is present for MerchantComment, not even an explicit nil
func (o *MerchantReturnResponse) UnsetMerchantComment() {
	o.MerchantComment.Unset()
}

// GetRefundInclVat returns the RefundInclVat field value if set, zero value otherwise.
func (o *MerchantReturnResponse) GetRefundInclVat() float32 {
	if o == nil || IsNil(o.RefundInclVat) {
		var ret float32
		return ret
	}
	return *o.RefundInclVat
}

// GetRefundInclVatOk returns a tuple with the RefundInclVat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantReturnResponse) GetRefundInclVatOk() (*float32, bool) {
	if o == nil || IsNil(o.RefundInclVat) {
		return nil, false
	}
	return o.RefundInclVat, true
}

// HasRefundInclVat returns a boolean if a field has been set.
func (o *MerchantReturnResponse) HasRefundInclVat() bool {
	if o != nil && !IsNil(o.RefundInclVat) {
		return true
	}

	return false
}

// SetRefundInclVat gets a reference to the given float32 and assigns it to the RefundInclVat field.
func (o *MerchantReturnResponse) SetRefundInclVat(v float32) {
	o.RefundInclVat = &v
}

// GetRefundExclVat returns the RefundExclVat field value if set, zero value otherwise.
func (o *MerchantReturnResponse) GetRefundExclVat() float32 {
	if o == nil || IsNil(o.RefundExclVat) {
		var ret float32
		return ret
	}
	return *o.RefundExclVat
}

// GetRefundExclVatOk returns a tuple with the RefundExclVat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantReturnResponse) GetRefundExclVatOk() (*float32, bool) {
	if o == nil || IsNil(o.RefundExclVat) {
		return nil, false
	}
	return o.RefundExclVat, true
}

// HasRefundExclVat returns a boolean if a field has been set.
func (o *MerchantReturnResponse) HasRefundExclVat() bool {
	if o != nil && !IsNil(o.RefundExclVat) {
		return true
	}

	return false
}

// SetRefundExclVat gets a reference to the given float32 and assigns it to the RefundExclVat field.
func (o *MerchantReturnResponse) SetRefundExclVat(v float32) {
	o.RefundExclVat = &v
}

// GetReturnDate returns the ReturnDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantReturnResponse) GetReturnDate() time.Time {
	if o == nil || IsNil(o.ReturnDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ReturnDate.Get()
}

// GetReturnDateOk returns a tuple with the ReturnDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantReturnResponse) GetReturnDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReturnDate.Get(), o.ReturnDate.IsSet()
}

// HasReturnDate returns a boolean if a field has been set.
func (o *MerchantReturnResponse) HasReturnDate() bool {
	if o != nil && o.ReturnDate.IsSet() {
		return true
	}

	return false
}

// SetReturnDate gets a reference to the given NullableTime and assigns it to the ReturnDate field.
func (o *MerchantReturnResponse) SetReturnDate(v time.Time) {
	o.ReturnDate.Set(&v)
}
// SetReturnDateNil sets the value for ReturnDate to be an explicit nil
func (o *MerchantReturnResponse) SetReturnDateNil() {
	o.ReturnDate.Set(nil)
}

// UnsetReturnDate ensures that no value is present for ReturnDate, not even an explicit nil
func (o *MerchantReturnResponse) UnsetReturnDate() {
	o.ReturnDate.Unset()
}

// GetExtraData returns the ExtraData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantReturnResponse) GetExtraData() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.ExtraData
}

// GetExtraDataOk returns a tuple with the ExtraData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantReturnResponse) GetExtraDataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ExtraData) {
		return nil, false
	}
	return &o.ExtraData, true
}

// HasExtraData returns a boolean if a field has been set.
func (o *MerchantReturnResponse) HasExtraData() bool {
	if o != nil && !IsNil(o.ExtraData) {
		return true
	}

	return false
}

// SetExtraData gets a reference to the given map[string]string and assigns it to the ExtraData field.
func (o *MerchantReturnResponse) SetExtraData(v map[string]string) {
	o.ExtraData = v
}

func (o MerchantReturnResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantReturnResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MerchantOrderNo.IsSet() {
		toSerialize["MerchantOrderNo"] = o.MerchantOrderNo.Get()
	}
	if o.ChannelOrderNo.IsSet() {
		toSerialize["ChannelOrderNo"] = o.ChannelOrderNo.Get()
	}
	if o.ChannelId.IsSet() {
		toSerialize["ChannelId"] = o.ChannelId.Get()
	}
	if o.GlobalChannelId.IsSet() {
		toSerialize["GlobalChannelId"] = o.GlobalChannelId.Get()
	}
	if o.GlobalChannelName.IsSet() {
		toSerialize["GlobalChannelName"] = o.GlobalChannelName.Get()
	}
	if o.Lines != nil {
		toSerialize["Lines"] = o.Lines
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["CreatedAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["UpdatedAt"] = o.UpdatedAt
	}
	if o.MerchantReturnNo.IsSet() {
		toSerialize["MerchantReturnNo"] = o.MerchantReturnNo.Get()
	}
	if o.ChannelReturnNo.IsSet() {
		toSerialize["ChannelReturnNo"] = o.ChannelReturnNo.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if o.AcknowledgedDate.IsSet() {
		toSerialize["AcknowledgedDate"] = o.AcknowledgedDate.Get()
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.Reason) {
		toSerialize["Reason"] = o.Reason
	}
	if o.CustomerComment.IsSet() {
		toSerialize["CustomerComment"] = o.CustomerComment.Get()
	}
	if o.MerchantComment.IsSet() {
		toSerialize["MerchantComment"] = o.MerchantComment.Get()
	}
	if !IsNil(o.RefundInclVat) {
		toSerialize["RefundInclVat"] = o.RefundInclVat
	}
	if !IsNil(o.RefundExclVat) {
		toSerialize["RefundExclVat"] = o.RefundExclVat
	}
	if o.ReturnDate.IsSet() {
		toSerialize["ReturnDate"] = o.ReturnDate.Get()
	}
	if o.ExtraData != nil {
		toSerialize["ExtraData"] = o.ExtraData
	}
	return toSerialize, nil
}

type NullableMerchantReturnResponse struct {
	value *MerchantReturnResponse
	isSet bool
}

func (v NullableMerchantReturnResponse) Get() *MerchantReturnResponse {
	return v.value
}

func (v *NullableMerchantReturnResponse) Set(val *MerchantReturnResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantReturnResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantReturnResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantReturnResponse(val *MerchantReturnResponse) *NullableMerchantReturnResponse {
	return &NullableMerchantReturnResponse{value: val, isSet: true}
}

func (v NullableMerchantReturnResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantReturnResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


