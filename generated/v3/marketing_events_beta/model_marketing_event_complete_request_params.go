/*
Marketing Events Extension

These APIs allow you to interact with HubSpot's Marketing Events Extension. It allows you to: * Create, Read or update Marketing Event information in HubSpot * Specify whether a HubSpot contact has registered, attended or cancelled a registration to a Marketing Event. * Specify a URL that can be called to get the details of a Marketing Event. 

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marketing_events_beta

import (
	"encoding/json"
	"time"
)

// MarketingEventCompleteRequestParams struct for MarketingEventCompleteRequestParams
type MarketingEventCompleteRequestParams struct {
	StartDateTime time.Time `json:"startDateTime"`
	EndDateTime time.Time `json:"endDateTime"`
}

// NewMarketingEventCompleteRequestParams instantiates a new MarketingEventCompleteRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketingEventCompleteRequestParams(startDateTime time.Time, endDateTime time.Time) *MarketingEventCompleteRequestParams {
	this := MarketingEventCompleteRequestParams{}
	this.StartDateTime = startDateTime
	this.EndDateTime = endDateTime
	return &this
}

// NewMarketingEventCompleteRequestParamsWithDefaults instantiates a new MarketingEventCompleteRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketingEventCompleteRequestParamsWithDefaults() *MarketingEventCompleteRequestParams {
	this := MarketingEventCompleteRequestParams{}
	return &this
}

// GetStartDateTime returns the StartDateTime field value
func (o *MarketingEventCompleteRequestParams) GetStartDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value
// and a boolean to check if the value has been set.
func (o *MarketingEventCompleteRequestParams) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDateTime, true
}

// SetStartDateTime sets field value
func (o *MarketingEventCompleteRequestParams) SetStartDateTime(v time.Time) {
	o.StartDateTime = v
}

// GetEndDateTime returns the EndDateTime field value
func (o *MarketingEventCompleteRequestParams) GetEndDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndDateTime
}

// GetEndDateTimeOk returns a tuple with the EndDateTime field value
// and a boolean to check if the value has been set.
func (o *MarketingEventCompleteRequestParams) GetEndDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDateTime, true
}

// SetEndDateTime sets field value
func (o *MarketingEventCompleteRequestParams) SetEndDateTime(v time.Time) {
	o.EndDateTime = v
}

func (o MarketingEventCompleteRequestParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["startDateTime"] = o.StartDateTime
	}
	if true {
		toSerialize["endDateTime"] = o.EndDateTime
	}
	return json.Marshal(toSerialize)
}

type NullableMarketingEventCompleteRequestParams struct {
	value *MarketingEventCompleteRequestParams
	isSet bool
}

func (v NullableMarketingEventCompleteRequestParams) Get() *MarketingEventCompleteRequestParams {
	return v.value
}

func (v *NullableMarketingEventCompleteRequestParams) Set(val *MarketingEventCompleteRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketingEventCompleteRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketingEventCompleteRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketingEventCompleteRequestParams(val *MarketingEventCompleteRequestParams) *NullableMarketingEventCompleteRequestParams {
	return &NullableMarketingEventCompleteRequestParams{value: val, isSet: true}
}

func (v NullableMarketingEventCompleteRequestParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketingEventCompleteRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


