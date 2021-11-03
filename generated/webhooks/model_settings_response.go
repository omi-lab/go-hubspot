/*
Webhooks API

Provides a way for apps to subscribe to certain change events in HubSpot. Once configured, apps will receive event payloads containing details about the changes at a specified target URL. There can only be one target URL for receiving event notifications per app.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package webhooks

import (
	"encoding/json"
	"time"
)

// SettingsResponse Webhook settings for an app.
type SettingsResponse struct {
	// A publicly available URL for Hubspot to call where event payloads will be delivered. See [link-so-some-doc](#) for details about the format of these event payloads.
	TargetUrl  string             `json:"targetUrl"`
	Throttling ThrottlingSettings `json:"throttling"`
	// When this subscription was created. Formatted as milliseconds from the [Unix epoch](#).
	CreatedAt time.Time `json:"createdAt"`
	// When this subscription was last updated. Formatted as milliseconds from the [Unix epoch](#).
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewSettingsResponse instantiates a new SettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsResponse(targetUrl string, throttling ThrottlingSettings, createdAt time.Time) *SettingsResponse {
	this := SettingsResponse{}
	this.TargetUrl = targetUrl
	this.Throttling = throttling
	this.CreatedAt = createdAt
	return &this
}

// NewSettingsResponseWithDefaults instantiates a new SettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsResponseWithDefaults() *SettingsResponse {
	this := SettingsResponse{}
	return &this
}

// GetTargetUrl returns the TargetUrl field value
func (o *SettingsResponse) GetTargetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetUrl
}

// GetTargetUrlOk returns a tuple with the TargetUrl field value
// and a boolean to check if the value has been set.
func (o *SettingsResponse) GetTargetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetUrl, true
}

// SetTargetUrl sets field value
func (o *SettingsResponse) SetTargetUrl(v string) {
	o.TargetUrl = v
}

// GetThrottling returns the Throttling field value
func (o *SettingsResponse) GetThrottling() ThrottlingSettings {
	if o == nil {
		var ret ThrottlingSettings
		return ret
	}

	return o.Throttling
}

// GetThrottlingOk returns a tuple with the Throttling field value
// and a boolean to check if the value has been set.
func (o *SettingsResponse) GetThrottlingOk() (*ThrottlingSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Throttling, true
}

// SetThrottling sets field value
func (o *SettingsResponse) SetThrottling(v ThrottlingSettings) {
	o.Throttling = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SettingsResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SettingsResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SettingsResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SettingsResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SettingsResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *SettingsResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o SettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["targetUrl"] = o.TargetUrl
	}
	if true {
		toSerialize["throttling"] = o.Throttling
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableSettingsResponse struct {
	value *SettingsResponse
	isSet bool
}

func (v NullableSettingsResponse) Get() *SettingsResponse {
	return v.value
}

func (v *NullableSettingsResponse) Set(val *SettingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsResponse(val *SettingsResponse) *NullableSettingsResponse {
	return &NullableSettingsResponse{value: val, isSet: true}
}

func (v NullableSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
