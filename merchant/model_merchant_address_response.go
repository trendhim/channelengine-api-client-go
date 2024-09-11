/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"encoding/json"
)

// checks if the MerchantAddressResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantAddressResponse{}

// MerchantAddressResponse struct for MerchantAddressResponse
type MerchantAddressResponse struct {
	// The first address line, use this field if address validation is disabled in ChannelEngine.
	Line1 NullableString `json:"Line1,omitempty"`
	// The second address line, use this field if address validation is disabled in ChannelEngine.
	Line2 NullableString `json:"Line2,omitempty"`
	// The third address line, use this field if address validation is disabled in ChannelEngine.
	Line3 NullableString `json:"Line3,omitempty"`
	Gender *Gender `json:"Gender,omitempty"`
	// Optional. Company addressed too.
	CompanyName NullableString `json:"CompanyName,omitempty"`
	// The first name of the customer.
	FirstName NullableString `json:"FirstName,omitempty"`
	// The last name of the customer (includes the surname prefix [tussenvoegsel] like 'de', 'van', 'du').
	LastName NullableString `json:"LastName,omitempty"`
	// The name of the street (without house number information)  This field might be empty if address validation is disabled in ChannelEngine.
	StreetName NullableString `json:"StreetName,omitempty"`
	// The house number  This field might be empty if address validation is disabled in ChannelEngine.
	HouseNr NullableString `json:"HouseNr,omitempty"`
	// Optional. Addition to the house number  If the address is: Groenhazengracht 2c, the address will be:  StreetName: Groenhazengracht  HouseNo: 2  HouseNrAddition: c  This field might be empty if address validation is disabled in ChannelEngine.
	HouseNrAddition NullableString `json:"HouseNrAddition,omitempty"`
	// The zip or postal code.
	ZipCode NullableString `json:"ZipCode,omitempty"`
	// The name of the city.
	City NullableString `json:"City,omitempty"`
	// Optional. State/province/region.
	Region NullableString `json:"Region,omitempty"`
	// For example: NL, BE, FR.
	CountryIso NullableString `json:"CountryIso,omitempty"`
	// Optional. The address as a single string: use in case the address lines are entered  as single lines and later parsed into street, house number and house number addition.
	// Deprecated
	Original NullableString `json:"Original,omitempty"`
}

// NewMerchantAddressResponse instantiates a new MerchantAddressResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantAddressResponse() *MerchantAddressResponse {
	this := MerchantAddressResponse{}
	return &this
}

// NewMerchantAddressResponseWithDefaults instantiates a new MerchantAddressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantAddressResponseWithDefaults() *MerchantAddressResponse {
	this := MerchantAddressResponse{}
	return &this
}

// GetLine1 returns the Line1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAddressResponse) GetLine1() string {
	if o == nil || IsNil(o.Line1.Get()) {
		var ret string
		return ret
	}
	return *o.Line1.Get()
}

// GetLine1Ok returns a tuple with the Line1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAddressResponse) GetLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Line1.Get(), o.Line1.IsSet()
}

// HasLine1 returns a boolean if a field has been set.
func (o *MerchantAddressResponse) HasLine1() bool {
	if o != nil && o.Line1.IsSet() {
		return true
	}

	return false
}

// SetLine1 gets a reference to the given NullableString and assigns it to the Line1 field.
func (o *MerchantAddressResponse) SetLine1(v string) {
	o.Line1.Set(&v)
}
// SetLine1Nil sets the value for Line1 to be an explicit nil
func (o *MerchantAddressResponse) SetLine1Nil() {
	o.Line1.Set(nil)
}

// UnsetLine1 ensures that no value is present for Line1, not even an explicit nil
func (o *MerchantAddressResponse) UnsetLine1() {
	o.Line1.Unset()
}

// GetLine2 returns the Line2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAddressResponse) GetLine2() string {
	if o == nil || IsNil(o.Line2.Get()) {
		var ret string
		return ret
	}
	return *o.Line2.Get()
}

// GetLine2Ok returns a tuple with the Line2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAddressResponse) GetLine2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Line2.Get(), o.Line2.IsSet()
}

// HasLine2 returns a boolean if a field has been set.
func (o *MerchantAddressResponse) HasLine2() bool {
	if o != nil && o.Line2.IsSet() {
		return true
	}

	return false
}

// SetLine2 gets a reference to the given NullableString and assigns it to the Line2 field.
func (o *MerchantAddressResponse) SetLine2(v string) {
	o.Line2.Set(&v)
}
// SetLine2Nil sets the value for Line2 to be an explicit nil
func (o *MerchantAddressResponse) SetLine2Nil() {
	o.Line2.Set(nil)
}

// UnsetLine2 ensures that no value is present for Line2, not even an explicit nil
func (o *MerchantAddressResponse) UnsetLine2() {
	o.Line2.Unset()
}

// GetLine3 returns the Line3 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAddressResponse) GetLine3() string {
	if o == nil || IsNil(o.Line3.Get()) {
		var ret string
		return ret
	}
	return *o.Line3.Get()
}

// GetLine3Ok returns a tuple with the Line3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAddressResponse) GetLine3Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Line3.Get(), o.Line3.IsSet()
}

// HasLine3 returns a boolean if a field has been set.
func (o *MerchantAddressResponse) HasLine3() bool {
	if o != nil && o.Line3.IsSet() {
		return true
	}

	return false
}

// SetLine3 gets a reference to the given NullableString and assigns it to the Line3 field.
func (o *MerchantAddressResponse) SetLine3(v string) {
	o.Line3.Set(&v)
}
// SetLine3Nil sets the value for Line3 to be an explicit nil
func (o *MerchantAddressResponse) SetLine3Nil() {
	o.Line3.Set(nil)
}

// UnsetLine3 ensures that no value is present for Line3, not even an explicit nil
func (o *MerchantAddressResponse) UnsetLine3() {
	o.Line3.Unset()
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *MerchantAddressResponse) GetGender() Gender {
	if o == nil || IsNil(o.Gender) {
		var ret Gender
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantAddressResponse) GetGenderOk() (*Gender, bool) {
	if o == nil || IsNil(o.Gender) {
		return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *MerchantAddressResponse) HasGender() bool {
	if o != nil && !IsNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given Gender and assigns it to the Gender field.
func (o *MerchantAddressResponse) SetGender(v Gender) {
	o.Gender = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAddressResponse) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName.Get()) {
		var ret string
		return ret
	}
	return *o.CompanyName.Get()
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAddressResponse) GetCompanyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompanyName.Get(), o.CompanyName.IsSet()
}

// HasCompanyName returns a boolean if a field has been set.
func (o *MerchantAddressResponse) HasCompanyName() bool {
	if o != nil && o.CompanyName.IsSet() {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given NullableString and assigns it to the CompanyName field.
func (o *MerchantAddressResponse) SetCompanyName(v string) {
	o.CompanyName.Set(&v)
}
// SetCompanyNameNil sets the value for CompanyName to be an explicit nil
func (o *MerchantAddressResponse) SetCompanyNameNil() {
	o.CompanyName.Set(nil)
}

// UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
func (o *MerchantAddressResponse) UnsetCompanyName() {
	o.CompanyName.Unset()
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAddressResponse) GetFirstName() string {
	if o == nil || IsNil(o.FirstName.Get()) {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAddressResponse) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *MerchantAddressResponse) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *MerchantAddressResponse) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *MerchantAddressResponse) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *MerchantAddressResponse) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAddressResponse) GetLastName() string {
	if o == nil || IsNil(o.LastName.Get()) {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAddressResponse) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *MerchantAddressResponse) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *MerchantAddressResponse) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *MerchantAddressResponse) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *MerchantAddressResponse) UnsetLastName() {
	o.LastName.Unset()
}

// GetStreetName returns the StreetName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAddressResponse) GetStreetName() string {
	if o == nil || IsNil(o.StreetName.Get()) {
		var ret string
		return ret
	}
	return *o.StreetName.Get()
}

// GetStreetNameOk returns a tuple with the StreetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAddressResponse) GetStreetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StreetName.Get(), o.StreetName.IsSet()
}

// HasStreetName returns a boolean if a field has been set.
func (o *MerchantAddressResponse) HasStreetName() bool {
	if o != nil && o.StreetName.IsSet() {
		return true
	}

	return false
}

// SetStreetName gets a reference to the given NullableString and assigns it to the StreetName field.
func (o *MerchantAddressResponse) SetStreetName(v string) {
	o.StreetName.Set(&v)
}
// SetStreetNameNil sets the value for StreetName to be an explicit nil
func (o *MerchantAddressResponse) SetStreetNameNil() {
	o.StreetName.Set(nil)
}

// UnsetStreetName ensures that no value is present for StreetName, not even an explicit nil
func (o *MerchantAddressResponse) UnsetStreetName() {
	o.StreetName.Unset()
}

// GetHouseNr returns the HouseNr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAddressResponse) GetHouseNr() string {
	if o == nil || IsNil(o.HouseNr.Get()) {
		var ret string
		return ret
	}
	return *o.HouseNr.Get()
}

// GetHouseNrOk returns a tuple with the HouseNr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAddressResponse) GetHouseNrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HouseNr.Get(), o.HouseNr.IsSet()
}

// HasHouseNr returns a boolean if a field has been set.
func (o *MerchantAddressResponse) HasHouseNr() bool {
	if o != nil && o.HouseNr.IsSet() {
		return true
	}

	return false
}

// SetHouseNr gets a reference to the given NullableString and assigns it to the HouseNr field.
func (o *MerchantAddressResponse) SetHouseNr(v string) {
	o.HouseNr.Set(&v)
}
// SetHouseNrNil sets the value for HouseNr to be an explicit nil
func (o *MerchantAddressResponse) SetHouseNrNil() {
	o.HouseNr.Set(nil)
}

// UnsetHouseNr ensures that no value is present for HouseNr, not even an explicit nil
func (o *MerchantAddressResponse) UnsetHouseNr() {
	o.HouseNr.Unset()
}

// GetHouseNrAddition returns the HouseNrAddition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAddressResponse) GetHouseNrAddition() string {
	if o == nil || IsNil(o.HouseNrAddition.Get()) {
		var ret string
		return ret
	}
	return *o.HouseNrAddition.Get()
}

// GetHouseNrAdditionOk returns a tuple with the HouseNrAddition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAddressResponse) GetHouseNrAdditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HouseNrAddition.Get(), o.HouseNrAddition.IsSet()
}

// HasHouseNrAddition returns a boolean if a field has been set.
func (o *MerchantAddressResponse) HasHouseNrAddition() bool {
	if o != nil && o.HouseNrAddition.IsSet() {
		return true
	}

	return false
}

// SetHouseNrAddition gets a reference to the given NullableString and assigns it to the HouseNrAddition field.
func (o *MerchantAddressResponse) SetHouseNrAddition(v string) {
	o.HouseNrAddition.Set(&v)
}
// SetHouseNrAdditionNil sets the value for HouseNrAddition to be an explicit nil
func (o *MerchantAddressResponse) SetHouseNrAdditionNil() {
	o.HouseNrAddition.Set(nil)
}

// UnsetHouseNrAddition ensures that no value is present for HouseNrAddition, not even an explicit nil
func (o *MerchantAddressResponse) UnsetHouseNrAddition() {
	o.HouseNrAddition.Unset()
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAddressResponse) GetZipCode() string {
	if o == nil || IsNil(o.ZipCode.Get()) {
		var ret string
		return ret
	}
	return *o.ZipCode.Get()
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAddressResponse) GetZipCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ZipCode.Get(), o.ZipCode.IsSet()
}

// HasZipCode returns a boolean if a field has been set.
func (o *MerchantAddressResponse) HasZipCode() bool {
	if o != nil && o.ZipCode.IsSet() {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given NullableString and assigns it to the ZipCode field.
func (o *MerchantAddressResponse) SetZipCode(v string) {
	o.ZipCode.Set(&v)
}
// SetZipCodeNil sets the value for ZipCode to be an explicit nil
func (o *MerchantAddressResponse) SetZipCodeNil() {
	o.ZipCode.Set(nil)
}

// UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil
func (o *MerchantAddressResponse) UnsetZipCode() {
	o.ZipCode.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAddressResponse) GetCity() string {
	if o == nil || IsNil(o.City.Get()) {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAddressResponse) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *MerchantAddressResponse) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *MerchantAddressResponse) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *MerchantAddressResponse) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *MerchantAddressResponse) UnsetCity() {
	o.City.Unset()
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAddressResponse) GetRegion() string {
	if o == nil || IsNil(o.Region.Get()) {
		var ret string
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAddressResponse) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *MerchantAddressResponse) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableString and assigns it to the Region field.
func (o *MerchantAddressResponse) SetRegion(v string) {
	o.Region.Set(&v)
}
// SetRegionNil sets the value for Region to be an explicit nil
func (o *MerchantAddressResponse) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *MerchantAddressResponse) UnsetRegion() {
	o.Region.Unset()
}

// GetCountryIso returns the CountryIso field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAddressResponse) GetCountryIso() string {
	if o == nil || IsNil(o.CountryIso.Get()) {
		var ret string
		return ret
	}
	return *o.CountryIso.Get()
}

// GetCountryIsoOk returns a tuple with the CountryIso field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAddressResponse) GetCountryIsoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountryIso.Get(), o.CountryIso.IsSet()
}

// HasCountryIso returns a boolean if a field has been set.
func (o *MerchantAddressResponse) HasCountryIso() bool {
	if o != nil && o.CountryIso.IsSet() {
		return true
	}

	return false
}

// SetCountryIso gets a reference to the given NullableString and assigns it to the CountryIso field.
func (o *MerchantAddressResponse) SetCountryIso(v string) {
	o.CountryIso.Set(&v)
}
// SetCountryIsoNil sets the value for CountryIso to be an explicit nil
func (o *MerchantAddressResponse) SetCountryIsoNil() {
	o.CountryIso.Set(nil)
}

// UnsetCountryIso ensures that no value is present for CountryIso, not even an explicit nil
func (o *MerchantAddressResponse) UnsetCountryIso() {
	o.CountryIso.Unset()
}

// GetOriginal returns the Original field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *MerchantAddressResponse) GetOriginal() string {
	if o == nil || IsNil(o.Original.Get()) {
		var ret string
		return ret
	}
	return *o.Original.Get()
}

// GetOriginalOk returns a tuple with the Original field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *MerchantAddressResponse) GetOriginalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Original.Get(), o.Original.IsSet()
}

// HasOriginal returns a boolean if a field has been set.
func (o *MerchantAddressResponse) HasOriginal() bool {
	if o != nil && o.Original.IsSet() {
		return true
	}

	return false
}

// SetOriginal gets a reference to the given NullableString and assigns it to the Original field.
// Deprecated
func (o *MerchantAddressResponse) SetOriginal(v string) {
	o.Original.Set(&v)
}
// SetOriginalNil sets the value for Original to be an explicit nil
func (o *MerchantAddressResponse) SetOriginalNil() {
	o.Original.Set(nil)
}

// UnsetOriginal ensures that no value is present for Original, not even an explicit nil
func (o *MerchantAddressResponse) UnsetOriginal() {
	o.Original.Unset()
}

func (o MerchantAddressResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantAddressResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Line1.IsSet() {
		toSerialize["Line1"] = o.Line1.Get()
	}
	if o.Line2.IsSet() {
		toSerialize["Line2"] = o.Line2.Get()
	}
	if o.Line3.IsSet() {
		toSerialize["Line3"] = o.Line3.Get()
	}
	if !IsNil(o.Gender) {
		toSerialize["Gender"] = o.Gender
	}
	if o.CompanyName.IsSet() {
		toSerialize["CompanyName"] = o.CompanyName.Get()
	}
	if o.FirstName.IsSet() {
		toSerialize["FirstName"] = o.FirstName.Get()
	}
	if o.LastName.IsSet() {
		toSerialize["LastName"] = o.LastName.Get()
	}
	if o.StreetName.IsSet() {
		toSerialize["StreetName"] = o.StreetName.Get()
	}
	if o.HouseNr.IsSet() {
		toSerialize["HouseNr"] = o.HouseNr.Get()
	}
	if o.HouseNrAddition.IsSet() {
		toSerialize["HouseNrAddition"] = o.HouseNrAddition.Get()
	}
	if o.ZipCode.IsSet() {
		toSerialize["ZipCode"] = o.ZipCode.Get()
	}
	if o.City.IsSet() {
		toSerialize["City"] = o.City.Get()
	}
	if o.Region.IsSet() {
		toSerialize["Region"] = o.Region.Get()
	}
	if o.CountryIso.IsSet() {
		toSerialize["CountryIso"] = o.CountryIso.Get()
	}
	if o.Original.IsSet() {
		toSerialize["Original"] = o.Original.Get()
	}
	return toSerialize, nil
}

type NullableMerchantAddressResponse struct {
	value *MerchantAddressResponse
	isSet bool
}

func (v NullableMerchantAddressResponse) Get() *MerchantAddressResponse {
	return v.value
}

func (v *NullableMerchantAddressResponse) Set(val *MerchantAddressResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantAddressResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantAddressResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantAddressResponse(val *MerchantAddressResponse) *NullableMerchantAddressResponse {
	return &NullableMerchantAddressResponse{value: val, isSet: true}
}

func (v NullableMerchantAddressResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantAddressResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


