/*
Feedback Submissions

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package feedback_submissions

import (
	"encoding/json"
)

// CollectionResponseAssociatedId struct for CollectionResponseAssociatedId
type CollectionResponseAssociatedId struct {
	Results []AssociatedId `json:"results"`
	Paging  *Paging        `json:"paging,omitempty"`
}

// NewCollectionResponseAssociatedId instantiates a new CollectionResponseAssociatedId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResponseAssociatedId(results []AssociatedId) *CollectionResponseAssociatedId {
	this := CollectionResponseAssociatedId{}
	this.Results = results
	return &this
}

// NewCollectionResponseAssociatedIdWithDefaults instantiates a new CollectionResponseAssociatedId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResponseAssociatedIdWithDefaults() *CollectionResponseAssociatedId {
	this := CollectionResponseAssociatedId{}
	return &this
}

// GetResults returns the Results field value
func (o *CollectionResponseAssociatedId) GetResults() []AssociatedId {
	if o == nil {
		var ret []AssociatedId
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *CollectionResponseAssociatedId) GetResultsOk() (*[]AssociatedId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *CollectionResponseAssociatedId) SetResults(v []AssociatedId) {
	o.Results = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *CollectionResponseAssociatedId) GetPaging() Paging {
	if o == nil || o.Paging == nil {
		var ret Paging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResponseAssociatedId) GetPagingOk() (*Paging, bool) {
	if o == nil || o.Paging == nil {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *CollectionResponseAssociatedId) HasPaging() bool {
	if o != nil && o.Paging != nil {
		return true
	}

	return false
}

// SetPaging gets a reference to the given Paging and assigns it to the Paging field.
func (o *CollectionResponseAssociatedId) SetPaging(v Paging) {
	o.Paging = &v
}

func (o CollectionResponseAssociatedId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["results"] = o.Results
	}
	if o.Paging != nil {
		toSerialize["paging"] = o.Paging
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionResponseAssociatedId struct {
	value *CollectionResponseAssociatedId
	isSet bool
}

func (v NullableCollectionResponseAssociatedId) Get() *CollectionResponseAssociatedId {
	return v.value
}

func (v *NullableCollectionResponseAssociatedId) Set(val *CollectionResponseAssociatedId) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResponseAssociatedId) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResponseAssociatedId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResponseAssociatedId(val *CollectionResponseAssociatedId) *NullableCollectionResponseAssociatedId {
	return &NullableCollectionResponseAssociatedId{value: val, isSet: true}
}

func (v NullableCollectionResponseAssociatedId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResponseAssociatedId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
