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

// MarketingEventDefaultResponse struct for MarketingEventDefaultResponse
type MarketingEventDefaultResponse struct {
	// The name of the marketing event.
	EventName string `json:"eventName"`
	// The type of the marketing event.
	EventType *string `json:"eventType,omitempty"`
	// The start date and time of the marketing event.
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// The end date and time of the marketing event.
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// The name of the organizer of the marketing event.
	EventOrganizer string `json:"eventOrganizer"`
	// The description of the marketing event.
	EventDescription *string `json:"eventDescription,omitempty"`
	// The URL in the external event application where the marketing event can be managed.
	EventUrl *string `json:"eventUrl,omitempty"`
	// Indicates if the marketing event has been cancelled.
	EventCancelled *bool `json:"eventCancelled,omitempty"`
	// A list of PropertyValues. These can be whatever kind of property names and values you want. However, they must already exist on the HubSpot account's definition of the MarketingEvent Object. If they don't they will be filtered out and not set. In order to do this you'll need to create a new PropertyGroup on the HubSpot account's MarketingEvent object for your specific app and create the Custom Property you want to track on that HubSpot account. Do not create any new default properties on the MarketingEvent object as that will apply to all HubSpot accounts.
	CustomProperties *[]PropertyValue `json:"customProperties,omitempty"`
}

// NewMarketingEventDefaultResponse instantiates a new MarketingEventDefaultResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketingEventDefaultResponse(eventName string, eventOrganizer string) *MarketingEventDefaultResponse {
	this := MarketingEventDefaultResponse{}
	this.EventName = eventName
	this.EventOrganizer = eventOrganizer
	return &this
}

// NewMarketingEventDefaultResponseWithDefaults instantiates a new MarketingEventDefaultResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketingEventDefaultResponseWithDefaults() *MarketingEventDefaultResponse {
	this := MarketingEventDefaultResponse{}
	return &this
}

// GetEventName returns the EventName field value
func (o *MarketingEventDefaultResponse) GetEventName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value
// and a boolean to check if the value has been set.
func (o *MarketingEventDefaultResponse) GetEventNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventName, true
}

// SetEventName sets field value
func (o *MarketingEventDefaultResponse) SetEventName(v string) {
	o.EventName = v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *MarketingEventDefaultResponse) GetEventType() string {
	if o == nil || o.EventType == nil {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventDefaultResponse) GetEventTypeOk() (*string, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *MarketingEventDefaultResponse) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *MarketingEventDefaultResponse) SetEventType(v string) {
	o.EventType = &v
}

// GetStartDateTime returns the StartDateTime field value if set, zero value otherwise.
func (o *MarketingEventDefaultResponse) GetStartDateTime() time.Time {
	if o == nil || o.StartDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventDefaultResponse) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil || o.StartDateTime == nil {
		return nil, false
	}
	return o.StartDateTime, true
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *MarketingEventDefaultResponse) HasStartDateTime() bool {
	if o != nil && o.StartDateTime != nil {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given time.Time and assigns it to the StartDateTime field.
func (o *MarketingEventDefaultResponse) SetStartDateTime(v time.Time) {
	o.StartDateTime = &v
}

// GetEndDateTime returns the EndDateTime field value if set, zero value otherwise.
func (o *MarketingEventDefaultResponse) GetEndDateTime() time.Time {
	if o == nil || o.EndDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDateTime
}

// GetEndDateTimeOk returns a tuple with the EndDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventDefaultResponse) GetEndDateTimeOk() (*time.Time, bool) {
	if o == nil || o.EndDateTime == nil {
		return nil, false
	}
	return o.EndDateTime, true
}

// HasEndDateTime returns a boolean if a field has been set.
func (o *MarketingEventDefaultResponse) HasEndDateTime() bool {
	if o != nil && o.EndDateTime != nil {
		return true
	}

	return false
}

// SetEndDateTime gets a reference to the given time.Time and assigns it to the EndDateTime field.
func (o *MarketingEventDefaultResponse) SetEndDateTime(v time.Time) {
	o.EndDateTime = &v
}

// GetEventOrganizer returns the EventOrganizer field value
func (o *MarketingEventDefaultResponse) GetEventOrganizer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventOrganizer
}

// GetEventOrganizerOk returns a tuple with the EventOrganizer field value
// and a boolean to check if the value has been set.
func (o *MarketingEventDefaultResponse) GetEventOrganizerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventOrganizer, true
}

// SetEventOrganizer sets field value
func (o *MarketingEventDefaultResponse) SetEventOrganizer(v string) {
	o.EventOrganizer = v
}

// GetEventDescription returns the EventDescription field value if set, zero value otherwise.
func (o *MarketingEventDefaultResponse) GetEventDescription() string {
	if o == nil || o.EventDescription == nil {
		var ret string
		return ret
	}
	return *o.EventDescription
}

// GetEventDescriptionOk returns a tuple with the EventDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventDefaultResponse) GetEventDescriptionOk() (*string, bool) {
	if o == nil || o.EventDescription == nil {
		return nil, false
	}
	return o.EventDescription, true
}

// HasEventDescription returns a boolean if a field has been set.
func (o *MarketingEventDefaultResponse) HasEventDescription() bool {
	if o != nil && o.EventDescription != nil {
		return true
	}

	return false
}

// SetEventDescription gets a reference to the given string and assigns it to the EventDescription field.
func (o *MarketingEventDefaultResponse) SetEventDescription(v string) {
	o.EventDescription = &v
}

// GetEventUrl returns the EventUrl field value if set, zero value otherwise.
func (o *MarketingEventDefaultResponse) GetEventUrl() string {
	if o == nil || o.EventUrl == nil {
		var ret string
		return ret
	}
	return *o.EventUrl
}

// GetEventUrlOk returns a tuple with the EventUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventDefaultResponse) GetEventUrlOk() (*string, bool) {
	if o == nil || o.EventUrl == nil {
		return nil, false
	}
	return o.EventUrl, true
}

// HasEventUrl returns a boolean if a field has been set.
func (o *MarketingEventDefaultResponse) HasEventUrl() bool {
	if o != nil && o.EventUrl != nil {
		return true
	}

	return false
}

// SetEventUrl gets a reference to the given string and assigns it to the EventUrl field.
func (o *MarketingEventDefaultResponse) SetEventUrl(v string) {
	o.EventUrl = &v
}

// GetEventCancelled returns the EventCancelled field value if set, zero value otherwise.
func (o *MarketingEventDefaultResponse) GetEventCancelled() bool {
	if o == nil || o.EventCancelled == nil {
		var ret bool
		return ret
	}
	return *o.EventCancelled
}

// GetEventCancelledOk returns a tuple with the EventCancelled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventDefaultResponse) GetEventCancelledOk() (*bool, bool) {
	if o == nil || o.EventCancelled == nil {
		return nil, false
	}
	return o.EventCancelled, true
}

// HasEventCancelled returns a boolean if a field has been set.
func (o *MarketingEventDefaultResponse) HasEventCancelled() bool {
	if o != nil && o.EventCancelled != nil {
		return true
	}

	return false
}

// SetEventCancelled gets a reference to the given bool and assigns it to the EventCancelled field.
func (o *MarketingEventDefaultResponse) SetEventCancelled(v bool) {
	o.EventCancelled = &v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *MarketingEventDefaultResponse) GetCustomProperties() []PropertyValue {
	if o == nil || o.CustomProperties == nil {
		var ret []PropertyValue
		return ret
	}
	return *o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventDefaultResponse) GetCustomPropertiesOk() (*[]PropertyValue, bool) {
	if o == nil || o.CustomProperties == nil {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *MarketingEventDefaultResponse) HasCustomProperties() bool {
	if o != nil && o.CustomProperties != nil {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given []PropertyValue and assigns it to the CustomProperties field.
func (o *MarketingEventDefaultResponse) SetCustomProperties(v []PropertyValue) {
	o.CustomProperties = &v
}

func (o MarketingEventDefaultResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["eventName"] = o.EventName
	}
	if o.EventType != nil {
		toSerialize["eventType"] = o.EventType
	}
	if o.StartDateTime != nil {
		toSerialize["startDateTime"] = o.StartDateTime
	}
	if o.EndDateTime != nil {
		toSerialize["endDateTime"] = o.EndDateTime
	}
	if true {
		toSerialize["eventOrganizer"] = o.EventOrganizer
	}
	if o.EventDescription != nil {
		toSerialize["eventDescription"] = o.EventDescription
	}
	if o.EventUrl != nil {
		toSerialize["eventUrl"] = o.EventUrl
	}
	if o.EventCancelled != nil {
		toSerialize["eventCancelled"] = o.EventCancelled
	}
	if o.CustomProperties != nil {
		toSerialize["customProperties"] = o.CustomProperties
	}
	return json.Marshal(toSerialize)
}

type NullableMarketingEventDefaultResponse struct {
	value *MarketingEventDefaultResponse
	isSet bool
}

func (v NullableMarketingEventDefaultResponse) Get() *MarketingEventDefaultResponse {
	return v.value
}

func (v *NullableMarketingEventDefaultResponse) Set(val *MarketingEventDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketingEventDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketingEventDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketingEventDefaultResponse(val *MarketingEventDefaultResponse) *NullableMarketingEventDefaultResponse {
	return &NullableMarketingEventDefaultResponse{value: val, isSet: true}
}

func (v NullableMarketingEventDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketingEventDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
