/*
Properties

All HubSpot objects store data in default and custom properties. These endpoints provide access to read and modify object properties in HubSpot.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package properties

import (
	"encoding/json"
)

// BatchInputPropertyName struct for BatchInputPropertyName
type BatchInputPropertyName struct {
	Inputs []PropertyName `json:"inputs"`
}

// NewBatchInputPropertyName instantiates a new BatchInputPropertyName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchInputPropertyName(inputs []PropertyName) *BatchInputPropertyName {
	this := BatchInputPropertyName{}
	this.Inputs = inputs
	return &this
}

// NewBatchInputPropertyNameWithDefaults instantiates a new BatchInputPropertyName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchInputPropertyNameWithDefaults() *BatchInputPropertyName {
	this := BatchInputPropertyName{}
	return &this
}

// GetInputs returns the Inputs field value
func (o *BatchInputPropertyName) GetInputs() []PropertyName {
	if o == nil {
		var ret []PropertyName
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchInputPropertyName) GetInputsOk() (*[]PropertyName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value
func (o *BatchInputPropertyName) SetInputs(v []PropertyName) {
	o.Inputs = v
}

func (o BatchInputPropertyName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["inputs"] = o.Inputs
	}
	return json.Marshal(toSerialize)
}

type NullableBatchInputPropertyName struct {
	value *BatchInputPropertyName
	isSet bool
}

func (v NullableBatchInputPropertyName) Get() *BatchInputPropertyName {
	return v.value
}

func (v *NullableBatchInputPropertyName) Set(val *BatchInputPropertyName) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchInputPropertyName) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchInputPropertyName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchInputPropertyName(val *BatchInputPropertyName) *NullableBatchInputPropertyName {
	return &NullableBatchInputPropertyName{value: val, isSet: true}
}

func (v NullableBatchInputPropertyName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchInputPropertyName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
