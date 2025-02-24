/*
Timeline events

This feature allows an app to create and configure custom events that can show up in the timelines of certain CRM objects like contacts, companies, tickets, or deals. You'll find multiple use cases for this API in the sections below.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package timeline

import (
	"encoding/json"
)

// TimelineEventTemplateTokenOption struct for TimelineEventTemplateTokenOption
type TimelineEventTemplateTokenOption struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

// NewTimelineEventTemplateTokenOption instantiates a new TimelineEventTemplateTokenOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimelineEventTemplateTokenOption(value string, label string) *TimelineEventTemplateTokenOption {
	this := TimelineEventTemplateTokenOption{}
	this.Value = value
	this.Label = label
	return &this
}

// NewTimelineEventTemplateTokenOptionWithDefaults instantiates a new TimelineEventTemplateTokenOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimelineEventTemplateTokenOptionWithDefaults() *TimelineEventTemplateTokenOption {
	this := TimelineEventTemplateTokenOption{}
	return &this
}

// GetValue returns the Value field value
func (o *TimelineEventTemplateTokenOption) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TimelineEventTemplateTokenOption) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TimelineEventTemplateTokenOption) SetValue(v string) {
	o.Value = v
}

// GetLabel returns the Label field value
func (o *TimelineEventTemplateTokenOption) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *TimelineEventTemplateTokenOption) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *TimelineEventTemplateTokenOption) SetLabel(v string) {
	o.Label = v
}

func (o TimelineEventTemplateTokenOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["label"] = o.Label
	}
	return json.Marshal(toSerialize)
}

type NullableTimelineEventTemplateTokenOption struct {
	value *TimelineEventTemplateTokenOption
	isSet bool
}

func (v NullableTimelineEventTemplateTokenOption) Get() *TimelineEventTemplateTokenOption {
	return v.value
}

func (v *NullableTimelineEventTemplateTokenOption) Set(val *TimelineEventTemplateTokenOption) {
	v.value = val
	v.isSet = true
}

func (v NullableTimelineEventTemplateTokenOption) IsSet() bool {
	return v.isSet
}

func (v *NullableTimelineEventTemplateTokenOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimelineEventTemplateTokenOption(val *TimelineEventTemplateTokenOption) *NullableTimelineEventTemplateTokenOption {
	return &NullableTimelineEventTemplateTokenOption{value: val, isSet: true}
}

func (v NullableTimelineEventTemplateTokenOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimelineEventTemplateTokenOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


