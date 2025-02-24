/*
Blog Post endpoints

Use these endpoints for interacting with Blog Posts, Blog Authors, and Blog Tags

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authors

import (
	"encoding/json"
)

// BatchInputBlogAuthor Wrapper for providing an array of blog authors as inputs.
type BatchInputBlogAuthor struct {
	// Blog authors to input.
	Inputs []BlogAuthor `json:"inputs"`
}

// NewBatchInputBlogAuthor instantiates a new BatchInputBlogAuthor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchInputBlogAuthor(inputs []BlogAuthor) *BatchInputBlogAuthor {
	this := BatchInputBlogAuthor{}
	this.Inputs = inputs
	return &this
}

// NewBatchInputBlogAuthorWithDefaults instantiates a new BatchInputBlogAuthor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchInputBlogAuthorWithDefaults() *BatchInputBlogAuthor {
	this := BatchInputBlogAuthor{}
	return &this
}

// GetInputs returns the Inputs field value
func (o *BatchInputBlogAuthor) GetInputs() []BlogAuthor {
	if o == nil {
		var ret []BlogAuthor
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchInputBlogAuthor) GetInputsOk() ([]BlogAuthor, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchInputBlogAuthor) SetInputs(v []BlogAuthor) {
	o.Inputs = v
}

func (o BatchInputBlogAuthor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["inputs"] = o.Inputs
	}
	return json.Marshal(toSerialize)
}

type NullableBatchInputBlogAuthor struct {
	value *BatchInputBlogAuthor
	isSet bool
}

func (v NullableBatchInputBlogAuthor) Get() *BatchInputBlogAuthor {
	return v.value
}

func (v *NullableBatchInputBlogAuthor) Set(val *BatchInputBlogAuthor) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchInputBlogAuthor) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchInputBlogAuthor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchInputBlogAuthor(val *BatchInputBlogAuthor) *NullableBatchInputBlogAuthor {
	return &NullableBatchInputBlogAuthor{value: val, isSet: true}
}

func (v NullableBatchInputBlogAuthor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchInputBlogAuthor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


