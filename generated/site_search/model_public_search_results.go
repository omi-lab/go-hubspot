/*
CMS Site Search

Use these endpoints for searching content on your HubSpot hosted CMS website(s).

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package site_search

import (
	"encoding/json"
)

// PublicSearchResults struct for PublicSearchResults
type PublicSearchResults struct {
	Total      int32                 `json:"total"`
	Offset     int32                 `json:"offset"`
	Limit      int32                 `json:"limit"`
	Results    []ContentSearchResult `json:"results"`
	SearchTerm *string               `json:"searchTerm,omitempty"`
	Page       int32                 `json:"page"`
}

// NewPublicSearchResults instantiates a new PublicSearchResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicSearchResults(total int32, offset int32, limit int32, results []ContentSearchResult, page int32) *PublicSearchResults {
	this := PublicSearchResults{}
	this.Total = total
	this.Offset = offset
	this.Limit = limit
	this.Results = results
	this.Page = page
	return &this
}

// NewPublicSearchResultsWithDefaults instantiates a new PublicSearchResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicSearchResultsWithDefaults() *PublicSearchResults {
	this := PublicSearchResults{}
	return &this
}

// GetTotal returns the Total field value
func (o *PublicSearchResults) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *PublicSearchResults) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *PublicSearchResults) SetTotal(v int32) {
	o.Total = v
}

// GetOffset returns the Offset field value
func (o *PublicSearchResults) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *PublicSearchResults) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *PublicSearchResults) SetOffset(v int32) {
	o.Offset = v
}

// GetLimit returns the Limit field value
func (o *PublicSearchResults) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *PublicSearchResults) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *PublicSearchResults) SetLimit(v int32) {
	o.Limit = v
}

// GetResults returns the Results field value
func (o *PublicSearchResults) GetResults() []ContentSearchResult {
	if o == nil {
		var ret []ContentSearchResult
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PublicSearchResults) GetResultsOk() (*[]ContentSearchResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PublicSearchResults) SetResults(v []ContentSearchResult) {
	o.Results = v
}

// GetSearchTerm returns the SearchTerm field value if set, zero value otherwise.
func (o *PublicSearchResults) GetSearchTerm() string {
	if o == nil || o.SearchTerm == nil {
		var ret string
		return ret
	}
	return *o.SearchTerm
}

// GetSearchTermOk returns a tuple with the SearchTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicSearchResults) GetSearchTermOk() (*string, bool) {
	if o == nil || o.SearchTerm == nil {
		return nil, false
	}
	return o.SearchTerm, true
}

// HasSearchTerm returns a boolean if a field has been set.
func (o *PublicSearchResults) HasSearchTerm() bool {
	if o != nil && o.SearchTerm != nil {
		return true
	}

	return false
}

// SetSearchTerm gets a reference to the given string and assigns it to the SearchTerm field.
func (o *PublicSearchResults) SetSearchTerm(v string) {
	o.SearchTerm = &v
}

// GetPage returns the Page field value
func (o *PublicSearchResults) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *PublicSearchResults) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *PublicSearchResults) SetPage(v int32) {
	o.Page = v
}

func (o PublicSearchResults) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["offset"] = o.Offset
	}
	if true {
		toSerialize["limit"] = o.Limit
	}
	if true {
		toSerialize["results"] = o.Results
	}
	if o.SearchTerm != nil {
		toSerialize["searchTerm"] = o.SearchTerm
	}
	if true {
		toSerialize["page"] = o.Page
	}
	return json.Marshal(toSerialize)
}

type NullablePublicSearchResults struct {
	value *PublicSearchResults
	isSet bool
}

func (v NullablePublicSearchResults) Get() *PublicSearchResults {
	return v.value
}

func (v *NullablePublicSearchResults) Set(val *PublicSearchResults) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicSearchResults) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicSearchResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicSearchResults(val *PublicSearchResults) *NullablePublicSearchResults {
	return &NullablePublicSearchResults{value: val, isSet: true}
}

func (v NullablePublicSearchResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicSearchResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
