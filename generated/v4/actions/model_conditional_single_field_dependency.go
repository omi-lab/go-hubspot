/*
Custom Workflow Actions

Create custom workflow actions

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package actions

import (
	"encoding/json"
)

// ConditionalSingleFieldDependency struct for ConditionalSingleFieldDependency
type ConditionalSingleFieldDependency struct {
	DependencyType string `json:"dependencyType"`
	DependentFieldNames []string `json:"dependentFieldNames"`
	ControllingFieldName string `json:"controllingFieldName"`
	ControllingFieldValue string `json:"controllingFieldValue"`
}

// NewConditionalSingleFieldDependency instantiates a new ConditionalSingleFieldDependency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionalSingleFieldDependency(dependencyType string, dependentFieldNames []string, controllingFieldName string, controllingFieldValue string) *ConditionalSingleFieldDependency {
	this := ConditionalSingleFieldDependency{}
	this.DependencyType = dependencyType
	this.DependentFieldNames = dependentFieldNames
	this.ControllingFieldName = controllingFieldName
	this.ControllingFieldValue = controllingFieldValue
	return &this
}

// NewConditionalSingleFieldDependencyWithDefaults instantiates a new ConditionalSingleFieldDependency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionalSingleFieldDependencyWithDefaults() *ConditionalSingleFieldDependency {
	this := ConditionalSingleFieldDependency{}
	var dependencyType string = "CONDITIONAL_SINGLE_FIELD"
	this.DependencyType = dependencyType
	return &this
}

// GetDependencyType returns the DependencyType field value
func (o *ConditionalSingleFieldDependency) GetDependencyType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DependencyType
}

// GetDependencyTypeOk returns a tuple with the DependencyType field value
// and a boolean to check if the value has been set.
func (o *ConditionalSingleFieldDependency) GetDependencyTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DependencyType, true
}

// SetDependencyType sets field value
func (o *ConditionalSingleFieldDependency) SetDependencyType(v string) {
	o.DependencyType = v
}

// GetDependentFieldNames returns the DependentFieldNames field value
func (o *ConditionalSingleFieldDependency) GetDependentFieldNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DependentFieldNames
}

// GetDependentFieldNamesOk returns a tuple with the DependentFieldNames field value
// and a boolean to check if the value has been set.
func (o *ConditionalSingleFieldDependency) GetDependentFieldNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DependentFieldNames, true
}

// SetDependentFieldNames sets field value
func (o *ConditionalSingleFieldDependency) SetDependentFieldNames(v []string) {
	o.DependentFieldNames = v
}

// GetControllingFieldName returns the ControllingFieldName field value
func (o *ConditionalSingleFieldDependency) GetControllingFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ControllingFieldName
}

// GetControllingFieldNameOk returns a tuple with the ControllingFieldName field value
// and a boolean to check if the value has been set.
func (o *ConditionalSingleFieldDependency) GetControllingFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ControllingFieldName, true
}

// SetControllingFieldName sets field value
func (o *ConditionalSingleFieldDependency) SetControllingFieldName(v string) {
	o.ControllingFieldName = v
}

// GetControllingFieldValue returns the ControllingFieldValue field value
func (o *ConditionalSingleFieldDependency) GetControllingFieldValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ControllingFieldValue
}

// GetControllingFieldValueOk returns a tuple with the ControllingFieldValue field value
// and a boolean to check if the value has been set.
func (o *ConditionalSingleFieldDependency) GetControllingFieldValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ControllingFieldValue, true
}

// SetControllingFieldValue sets field value
func (o *ConditionalSingleFieldDependency) SetControllingFieldValue(v string) {
	o.ControllingFieldValue = v
}

func (o ConditionalSingleFieldDependency) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dependencyType"] = o.DependencyType
	}
	if true {
		toSerialize["dependentFieldNames"] = o.DependentFieldNames
	}
	if true {
		toSerialize["controllingFieldName"] = o.ControllingFieldName
	}
	if true {
		toSerialize["controllingFieldValue"] = o.ControllingFieldValue
	}
	return json.Marshal(toSerialize)
}

type NullableConditionalSingleFieldDependency struct {
	value *ConditionalSingleFieldDependency
	isSet bool
}

func (v NullableConditionalSingleFieldDependency) Get() *ConditionalSingleFieldDependency {
	return v.value
}

func (v *NullableConditionalSingleFieldDependency) Set(val *ConditionalSingleFieldDependency) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionalSingleFieldDependency) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionalSingleFieldDependency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionalSingleFieldDependency(val *ConditionalSingleFieldDependency) *NullableConditionalSingleFieldDependency {
	return &NullableConditionalSingleFieldDependency{value: val, isSet: true}
}

func (v NullableConditionalSingleFieldDependency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionalSingleFieldDependency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


