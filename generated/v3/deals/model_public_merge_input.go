/*
Deals

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package deals

import (
	"encoding/json"
)

// PublicMergeInput struct for PublicMergeInput
type PublicMergeInput struct {
	PrimaryObjectId string `json:"primaryObjectId"`
	ObjectIdToMerge string `json:"objectIdToMerge"`
}

// NewPublicMergeInput instantiates a new PublicMergeInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicMergeInput(primaryObjectId string, objectIdToMerge string) *PublicMergeInput {
	this := PublicMergeInput{}
	this.PrimaryObjectId = primaryObjectId
	this.ObjectIdToMerge = objectIdToMerge
	return &this
}

// NewPublicMergeInputWithDefaults instantiates a new PublicMergeInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicMergeInputWithDefaults() *PublicMergeInput {
	this := PublicMergeInput{}
	return &this
}

// GetPrimaryObjectId returns the PrimaryObjectId field value
func (o *PublicMergeInput) GetPrimaryObjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrimaryObjectId
}

// GetPrimaryObjectIdOk returns a tuple with the PrimaryObjectId field value
// and a boolean to check if the value has been set.
func (o *PublicMergeInput) GetPrimaryObjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryObjectId, true
}

// SetPrimaryObjectId sets field value
func (o *PublicMergeInput) SetPrimaryObjectId(v string) {
	o.PrimaryObjectId = v
}

// GetObjectIdToMerge returns the ObjectIdToMerge field value
func (o *PublicMergeInput) GetObjectIdToMerge() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectIdToMerge
}

// GetObjectIdToMergeOk returns a tuple with the ObjectIdToMerge field value
// and a boolean to check if the value has been set.
func (o *PublicMergeInput) GetObjectIdToMergeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectIdToMerge, true
}

// SetObjectIdToMerge sets field value
func (o *PublicMergeInput) SetObjectIdToMerge(v string) {
	o.ObjectIdToMerge = v
}

func (o PublicMergeInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["primaryObjectId"] = o.PrimaryObjectId
	}
	if true {
		toSerialize["objectIdToMerge"] = o.ObjectIdToMerge
	}
	return json.Marshal(toSerialize)
}

type NullablePublicMergeInput struct {
	value *PublicMergeInput
	isSet bool
}

func (v NullablePublicMergeInput) Get() *PublicMergeInput {
	return v.value
}

func (v *NullablePublicMergeInput) Set(val *PublicMergeInput) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicMergeInput) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicMergeInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicMergeInput(val *PublicMergeInput) *NullablePublicMergeInput {
	return &NullablePublicMergeInput{value: val, isSet: true}
}

func (v NullablePublicMergeInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicMergeInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


