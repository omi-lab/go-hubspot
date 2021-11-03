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

// PropertyModificationMetadata struct for PropertyModificationMetadata
type PropertyModificationMetadata struct {
	Archivable         bool  `json:"archivable"`
	ReadOnlyDefinition bool  `json:"readOnlyDefinition"`
	ReadOnlyOptions    *bool `json:"readOnlyOptions,omitempty"`
	ReadOnlyValue      bool  `json:"readOnlyValue"`
}

// NewPropertyModificationMetadata instantiates a new PropertyModificationMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyModificationMetadata(archivable bool, readOnlyDefinition bool, readOnlyValue bool) *PropertyModificationMetadata {
	this := PropertyModificationMetadata{}
	this.Archivable = archivable
	this.ReadOnlyDefinition = readOnlyDefinition
	this.ReadOnlyValue = readOnlyValue
	return &this
}

// NewPropertyModificationMetadataWithDefaults instantiates a new PropertyModificationMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyModificationMetadataWithDefaults() *PropertyModificationMetadata {
	this := PropertyModificationMetadata{}
	return &this
}

// GetArchivable returns the Archivable field value
func (o *PropertyModificationMetadata) GetArchivable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Archivable
}

// GetArchivableOk returns a tuple with the Archivable field value
// and a boolean to check if the value has been set.
func (o *PropertyModificationMetadata) GetArchivableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archivable, true
}

// SetArchivable sets field value
func (o *PropertyModificationMetadata) SetArchivable(v bool) {
	o.Archivable = v
}

// GetReadOnlyDefinition returns the ReadOnlyDefinition field value
func (o *PropertyModificationMetadata) GetReadOnlyDefinition() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ReadOnlyDefinition
}

// GetReadOnlyDefinitionOk returns a tuple with the ReadOnlyDefinition field value
// and a boolean to check if the value has been set.
func (o *PropertyModificationMetadata) GetReadOnlyDefinitionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadOnlyDefinition, true
}

// SetReadOnlyDefinition sets field value
func (o *PropertyModificationMetadata) SetReadOnlyDefinition(v bool) {
	o.ReadOnlyDefinition = v
}

// GetReadOnlyOptions returns the ReadOnlyOptions field value if set, zero value otherwise.
func (o *PropertyModificationMetadata) GetReadOnlyOptions() bool {
	if o == nil || o.ReadOnlyOptions == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnlyOptions
}

// GetReadOnlyOptionsOk returns a tuple with the ReadOnlyOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyModificationMetadata) GetReadOnlyOptionsOk() (*bool, bool) {
	if o == nil || o.ReadOnlyOptions == nil {
		return nil, false
	}
	return o.ReadOnlyOptions, true
}

// HasReadOnlyOptions returns a boolean if a field has been set.
func (o *PropertyModificationMetadata) HasReadOnlyOptions() bool {
	if o != nil && o.ReadOnlyOptions != nil {
		return true
	}

	return false
}

// SetReadOnlyOptions gets a reference to the given bool and assigns it to the ReadOnlyOptions field.
func (o *PropertyModificationMetadata) SetReadOnlyOptions(v bool) {
	o.ReadOnlyOptions = &v
}

// GetReadOnlyValue returns the ReadOnlyValue field value
func (o *PropertyModificationMetadata) GetReadOnlyValue() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ReadOnlyValue
}

// GetReadOnlyValueOk returns a tuple with the ReadOnlyValue field value
// and a boolean to check if the value has been set.
func (o *PropertyModificationMetadata) GetReadOnlyValueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadOnlyValue, true
}

// SetReadOnlyValue sets field value
func (o *PropertyModificationMetadata) SetReadOnlyValue(v bool) {
	o.ReadOnlyValue = v
}

func (o PropertyModificationMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["archivable"] = o.Archivable
	}
	if true {
		toSerialize["readOnlyDefinition"] = o.ReadOnlyDefinition
	}
	if o.ReadOnlyOptions != nil {
		toSerialize["readOnlyOptions"] = o.ReadOnlyOptions
	}
	if true {
		toSerialize["readOnlyValue"] = o.ReadOnlyValue
	}
	return json.Marshal(toSerialize)
}

type NullablePropertyModificationMetadata struct {
	value *PropertyModificationMetadata
	isSet bool
}

func (v NullablePropertyModificationMetadata) Get() *PropertyModificationMetadata {
	return v.value
}

func (v *NullablePropertyModificationMetadata) Set(val *PropertyModificationMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyModificationMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyModificationMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyModificationMetadata(val *PropertyModificationMetadata) *NullablePropertyModificationMetadata {
	return &NullablePropertyModificationMetadata{value: val, isSet: true}
}

func (v NullablePropertyModificationMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyModificationMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
