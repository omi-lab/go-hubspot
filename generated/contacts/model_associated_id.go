/*
Contacts

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package contacts

import (
	"encoding/json"
)

// AssociatedId struct for AssociatedId
type AssociatedId struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

// NewAssociatedId instantiates a new AssociatedId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssociatedId(id string, type_ string) *AssociatedId {
	this := AssociatedId{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewAssociatedIdWithDefaults instantiates a new AssociatedId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssociatedIdWithDefaults() *AssociatedId {
	this := AssociatedId{}
	return &this
}

// GetId returns the Id field value
func (o *AssociatedId) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AssociatedId) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AssociatedId) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *AssociatedId) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AssociatedId) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AssociatedId) SetType(v string) {
	o.Type = v
}

func (o AssociatedId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableAssociatedId struct {
	value *AssociatedId
	isSet bool
}

func (v NullableAssociatedId) Get() *AssociatedId {
	return v.value
}

func (v *NullableAssociatedId) Set(val *AssociatedId) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociatedId) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociatedId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociatedId(val *AssociatedId) *NullableAssociatedId {
	return &NullableAssociatedId{value: val, isSet: true}
}

func (v NullableAssociatedId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssociatedId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
