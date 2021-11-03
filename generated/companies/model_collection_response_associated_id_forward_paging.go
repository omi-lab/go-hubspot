/*
Companies

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package companies

import (
	"encoding/json"
)

// CollectionResponseAssociatedIdForwardPaging struct for CollectionResponseAssociatedIdForwardPaging
type CollectionResponseAssociatedIdForwardPaging struct {
	Results []AssociatedId `json:"results"`
	Paging  *ForwardPaging `json:"paging,omitempty"`
}

// NewCollectionResponseAssociatedIdForwardPaging instantiates a new CollectionResponseAssociatedIdForwardPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResponseAssociatedIdForwardPaging(results []AssociatedId) *CollectionResponseAssociatedIdForwardPaging {
	this := CollectionResponseAssociatedIdForwardPaging{}
	this.Results = results
	return &this
}

// NewCollectionResponseAssociatedIdForwardPagingWithDefaults instantiates a new CollectionResponseAssociatedIdForwardPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResponseAssociatedIdForwardPagingWithDefaults() *CollectionResponseAssociatedIdForwardPaging {
	this := CollectionResponseAssociatedIdForwardPaging{}
	return &this
}

// GetResults returns the Results field value
func (o *CollectionResponseAssociatedIdForwardPaging) GetResults() []AssociatedId {
	if o == nil {
		var ret []AssociatedId
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *CollectionResponseAssociatedIdForwardPaging) GetResultsOk() (*[]AssociatedId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *CollectionResponseAssociatedIdForwardPaging) SetResults(v []AssociatedId) {
	o.Results = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *CollectionResponseAssociatedIdForwardPaging) GetPaging() ForwardPaging {
	if o == nil || o.Paging == nil {
		var ret ForwardPaging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResponseAssociatedIdForwardPaging) GetPagingOk() (*ForwardPaging, bool) {
	if o == nil || o.Paging == nil {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *CollectionResponseAssociatedIdForwardPaging) HasPaging() bool {
	if o != nil && o.Paging != nil {
		return true
	}

	return false
}

// SetPaging gets a reference to the given ForwardPaging and assigns it to the Paging field.
func (o *CollectionResponseAssociatedIdForwardPaging) SetPaging(v ForwardPaging) {
	o.Paging = &v
}

func (o CollectionResponseAssociatedIdForwardPaging) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["results"] = o.Results
	}
	if o.Paging != nil {
		toSerialize["paging"] = o.Paging
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionResponseAssociatedIdForwardPaging struct {
	value *CollectionResponseAssociatedIdForwardPaging
	isSet bool
}

func (v NullableCollectionResponseAssociatedIdForwardPaging) Get() *CollectionResponseAssociatedIdForwardPaging {
	return v.value
}

func (v *NullableCollectionResponseAssociatedIdForwardPaging) Set(val *CollectionResponseAssociatedIdForwardPaging) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResponseAssociatedIdForwardPaging) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResponseAssociatedIdForwardPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResponseAssociatedIdForwardPaging(val *CollectionResponseAssociatedIdForwardPaging) *NullableCollectionResponseAssociatedIdForwardPaging {
	return &NullableCollectionResponseAssociatedIdForwardPaging{value: val, isSet: true}
}

func (v NullableCollectionResponseAssociatedIdForwardPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResponseAssociatedIdForwardPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
