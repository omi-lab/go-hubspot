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

// BatchResponseSubscriptionResponseWithErrors struct for BatchResponseSubscriptionResponseWithErrors
type BatchResponseSubscriptionResponseWithErrors struct {
	Status      string                 `json:"status"`
	Results     []SubscriptionResponse `json:"results"`
	NumErrors   *int32                 `json:"numErrors,omitempty"`
	Errors      *[]StandardError       `json:"errors,omitempty"`
	RequestedAt *time.Time             `json:"requestedAt,omitempty"`
	StartedAt   time.Time              `json:"startedAt"`
	CompletedAt time.Time              `json:"completedAt"`
	Links       *map[string]string     `json:"links,omitempty"`
}

// NewBatchResponseSubscriptionResponseWithErrors instantiates a new BatchResponseSubscriptionResponseWithErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchResponseSubscriptionResponseWithErrors(status string, results []SubscriptionResponse, startedAt time.Time, completedAt time.Time) *BatchResponseSubscriptionResponseWithErrors {
	this := BatchResponseSubscriptionResponseWithErrors{}
	this.Status = status
	this.Results = results
	this.StartedAt = startedAt
	this.CompletedAt = completedAt
	return &this
}

// NewBatchResponseSubscriptionResponseWithErrorsWithDefaults instantiates a new BatchResponseSubscriptionResponseWithErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchResponseSubscriptionResponseWithErrorsWithDefaults() *BatchResponseSubscriptionResponseWithErrors {
	this := BatchResponseSubscriptionResponseWithErrors{}
	return &this
}

// GetStatus returns the Status field value
func (o *BatchResponseSubscriptionResponseWithErrors) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BatchResponseSubscriptionResponseWithErrors) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BatchResponseSubscriptionResponseWithErrors) SetStatus(v string) {
	o.Status = v
}

// GetResults returns the Results field value
func (o *BatchResponseSubscriptionResponseWithErrors) GetResults() []SubscriptionResponse {
	if o == nil {
		var ret []SubscriptionResponse
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *BatchResponseSubscriptionResponseWithErrors) GetResultsOk() (*[]SubscriptionResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *BatchResponseSubscriptionResponseWithErrors) SetResults(v []SubscriptionResponse) {
	o.Results = v
}

// GetNumErrors returns the NumErrors field value if set, zero value otherwise.
func (o *BatchResponseSubscriptionResponseWithErrors) GetNumErrors() int32 {
	if o == nil || o.NumErrors == nil {
		var ret int32
		return ret
	}
	return *o.NumErrors
}

// GetNumErrorsOk returns a tuple with the NumErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseSubscriptionResponseWithErrors) GetNumErrorsOk() (*int32, bool) {
	if o == nil || o.NumErrors == nil {
		return nil, false
	}
	return o.NumErrors, true
}

// HasNumErrors returns a boolean if a field has been set.
func (o *BatchResponseSubscriptionResponseWithErrors) HasNumErrors() bool {
	if o != nil && o.NumErrors != nil {
		return true
	}

	return false
}

// SetNumErrors gets a reference to the given int32 and assigns it to the NumErrors field.
func (o *BatchResponseSubscriptionResponseWithErrors) SetNumErrors(v int32) {
	o.NumErrors = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *BatchResponseSubscriptionResponseWithErrors) GetErrors() []StandardError {
	if o == nil || o.Errors == nil {
		var ret []StandardError
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseSubscriptionResponseWithErrors) GetErrorsOk() (*[]StandardError, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *BatchResponseSubscriptionResponseWithErrors) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []StandardError and assigns it to the Errors field.
func (o *BatchResponseSubscriptionResponseWithErrors) SetErrors(v []StandardError) {
	o.Errors = &v
}

// GetRequestedAt returns the RequestedAt field value if set, zero value otherwise.
func (o *BatchResponseSubscriptionResponseWithErrors) GetRequestedAt() time.Time {
	if o == nil || o.RequestedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.RequestedAt
}

// GetRequestedAtOk returns a tuple with the RequestedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseSubscriptionResponseWithErrors) GetRequestedAtOk() (*time.Time, bool) {
	if o == nil || o.RequestedAt == nil {
		return nil, false
	}
	return o.RequestedAt, true
}

// HasRequestedAt returns a boolean if a field has been set.
func (o *BatchResponseSubscriptionResponseWithErrors) HasRequestedAt() bool {
	if o != nil && o.RequestedAt != nil {
		return true
	}

	return false
}

// SetRequestedAt gets a reference to the given time.Time and assigns it to the RequestedAt field.
func (o *BatchResponseSubscriptionResponseWithErrors) SetRequestedAt(v time.Time) {
	o.RequestedAt = &v
}

// GetStartedAt returns the StartedAt field value
func (o *BatchResponseSubscriptionResponseWithErrors) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *BatchResponseSubscriptionResponseWithErrors) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *BatchResponseSubscriptionResponseWithErrors) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetCompletedAt returns the CompletedAt field value
func (o *BatchResponseSubscriptionResponseWithErrors) GetCompletedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value
// and a boolean to check if the value has been set.
func (o *BatchResponseSubscriptionResponseWithErrors) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletedAt, true
}

// SetCompletedAt sets field value
func (o *BatchResponseSubscriptionResponseWithErrors) SetCompletedAt(v time.Time) {
	o.CompletedAt = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BatchResponseSubscriptionResponseWithErrors) GetLinks() map[string]string {
	if o == nil || o.Links == nil {
		var ret map[string]string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseSubscriptionResponseWithErrors) GetLinksOk() (*map[string]string, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BatchResponseSubscriptionResponseWithErrors) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]string and assigns it to the Links field.
func (o *BatchResponseSubscriptionResponseWithErrors) SetLinks(v map[string]string) {
	o.Links = &v
}

func (o BatchResponseSubscriptionResponseWithErrors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["results"] = o.Results
	}
	if o.NumErrors != nil {
		toSerialize["numErrors"] = o.NumErrors
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.RequestedAt != nil {
		toSerialize["requestedAt"] = o.RequestedAt
	}
	if true {
		toSerialize["startedAt"] = o.StartedAt
	}
	if true {
		toSerialize["completedAt"] = o.CompletedAt
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableBatchResponseSubscriptionResponseWithErrors struct {
	value *BatchResponseSubscriptionResponseWithErrors
	isSet bool
}

func (v NullableBatchResponseSubscriptionResponseWithErrors) Get() *BatchResponseSubscriptionResponseWithErrors {
	return v.value
}

func (v *NullableBatchResponseSubscriptionResponseWithErrors) Set(val *BatchResponseSubscriptionResponseWithErrors) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchResponseSubscriptionResponseWithErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchResponseSubscriptionResponseWithErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchResponseSubscriptionResponseWithErrors(val *BatchResponseSubscriptionResponseWithErrors) *NullableBatchResponseSubscriptionResponseWithErrors {
	return &NullableBatchResponseSubscriptionResponseWithErrors{value: val, isSet: true}
}

func (v NullableBatchResponseSubscriptionResponseWithErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchResponseSubscriptionResponseWithErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
