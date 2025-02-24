/*
Marketing Events Extension

These APIs allow you to interact with HubSpot's Marketing Events Extension. It allows you to: * Create, Read or update Marketing Event information in HubSpot * Specify whether a HubSpot contact has registered, attended or cancelled a registration to a Marketing Event. * Specify a URL that can be called to get the details of a Marketing Event. 

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marketing_events_beta

import (
	"encoding/json"
)

// SubscriberEmailResponse struct for SubscriberEmailResponse
type SubscriberEmailResponse struct {
	Vid int32 `json:"vid"`
	Email string `json:"email"`
}

// NewSubscriberEmailResponse instantiates a new SubscriberEmailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriberEmailResponse(vid int32, email string) *SubscriberEmailResponse {
	this := SubscriberEmailResponse{}
	this.Vid = vid
	this.Email = email
	return &this
}

// NewSubscriberEmailResponseWithDefaults instantiates a new SubscriberEmailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriberEmailResponseWithDefaults() *SubscriberEmailResponse {
	this := SubscriberEmailResponse{}
	return &this
}

// GetVid returns the Vid field value
func (o *SubscriberEmailResponse) GetVid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vid
}

// GetVidOk returns a tuple with the Vid field value
// and a boolean to check if the value has been set.
func (o *SubscriberEmailResponse) GetVidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vid, true
}

// SetVid sets field value
func (o *SubscriberEmailResponse) SetVid(v int32) {
	o.Vid = v
}

// GetEmail returns the Email field value
func (o *SubscriberEmailResponse) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SubscriberEmailResponse) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *SubscriberEmailResponse) SetEmail(v string) {
	o.Email = v
}

func (o SubscriberEmailResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["vid"] = o.Vid
	}
	if true {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriberEmailResponse struct {
	value *SubscriberEmailResponse
	isSet bool
}

func (v NullableSubscriberEmailResponse) Get() *SubscriberEmailResponse {
	return v.value
}

func (v *NullableSubscriberEmailResponse) Set(val *SubscriberEmailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriberEmailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriberEmailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriberEmailResponse(val *SubscriberEmailResponse) *NullableSubscriberEmailResponse {
	return &NullableSubscriberEmailResponse{value: val, isSet: true}
}

func (v NullableSubscriberEmailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriberEmailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


