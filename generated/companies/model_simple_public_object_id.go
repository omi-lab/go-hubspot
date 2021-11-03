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

// SimplePublicObjectId struct for SimplePublicObjectId
type SimplePublicObjectId struct {
	Id string `json:"id"`
}

// NewSimplePublicObjectId instantiates a new SimplePublicObjectId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimplePublicObjectId(id string) *SimplePublicObjectId {
	this := SimplePublicObjectId{}
	this.Id = id
	return &this
}

// NewSimplePublicObjectIdWithDefaults instantiates a new SimplePublicObjectId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimplePublicObjectIdWithDefaults() *SimplePublicObjectId {
	this := SimplePublicObjectId{}
	return &this
}

// GetId returns the Id field value
func (o *SimplePublicObjectId) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectId) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SimplePublicObjectId) SetId(v string) {
	o.Id = v
}

func (o SimplePublicObjectId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableSimplePublicObjectId struct {
	value *SimplePublicObjectId
	isSet bool
}

func (v NullableSimplePublicObjectId) Get() *SimplePublicObjectId {
	return v.value
}

func (v *NullableSimplePublicObjectId) Set(val *SimplePublicObjectId) {
	v.value = val
	v.isSet = true
}

func (v NullableSimplePublicObjectId) IsSet() bool {
	return v.isSet
}

func (v *NullableSimplePublicObjectId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimplePublicObjectId(val *SimplePublicObjectId) *NullableSimplePublicObjectId {
	return &NullableSimplePublicObjectId{value: val, isSet: true}
}

func (v NullableSimplePublicObjectId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimplePublicObjectId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
