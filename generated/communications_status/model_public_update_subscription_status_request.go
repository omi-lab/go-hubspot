/*
Subscriptions

Subscriptions allow contacts to control what forms of communications they receive. Contacts can decide whether they want to receive communication pertaining to a specific topic, brand, or an entire HubSpot account.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package communications_status

import (
	"encoding/json"
)

// PublicUpdateSubscriptionStatusRequest A request to change the status of a contact's subscription.
type PublicUpdateSubscriptionStatusRequest struct {
	// Contact's email address.
	EmailAddress string `json:"emailAddress"`
	// ID of the subscription the contact is being resubscribed to.
	SubscriptionId string `json:"subscriptionId"`
	// Legal basis for resubscribing the contact (required for GDPR enabled portals).
	LegalBasis *string `json:"legalBasis,omitempty"`
	// A more detailed explanation to go with the legal basis (required for GDPR enabled portals).
	LegalBasisExplanation *string `json:"legalBasisExplanation,omitempty"`
}

// NewPublicUpdateSubscriptionStatusRequest instantiates a new PublicUpdateSubscriptionStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicUpdateSubscriptionStatusRequest(emailAddress string, subscriptionId string) *PublicUpdateSubscriptionStatusRequest {
	this := PublicUpdateSubscriptionStatusRequest{}
	this.EmailAddress = emailAddress
	this.SubscriptionId = subscriptionId
	return &this
}

// NewPublicUpdateSubscriptionStatusRequestWithDefaults instantiates a new PublicUpdateSubscriptionStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicUpdateSubscriptionStatusRequestWithDefaults() *PublicUpdateSubscriptionStatusRequest {
	this := PublicUpdateSubscriptionStatusRequest{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value
func (o *PublicUpdateSubscriptionStatusRequest) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *PublicUpdateSubscriptionStatusRequest) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *PublicUpdateSubscriptionStatusRequest) SetEmailAddress(v string) {
	o.EmailAddress = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *PublicUpdateSubscriptionStatusRequest) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *PublicUpdateSubscriptionStatusRequest) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *PublicUpdateSubscriptionStatusRequest) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetLegalBasis returns the LegalBasis field value if set, zero value otherwise.
func (o *PublicUpdateSubscriptionStatusRequest) GetLegalBasis() string {
	if o == nil || o.LegalBasis == nil {
		var ret string
		return ret
	}
	return *o.LegalBasis
}

// GetLegalBasisOk returns a tuple with the LegalBasis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicUpdateSubscriptionStatusRequest) GetLegalBasisOk() (*string, bool) {
	if o == nil || o.LegalBasis == nil {
		return nil, false
	}
	return o.LegalBasis, true
}

// HasLegalBasis returns a boolean if a field has been set.
func (o *PublicUpdateSubscriptionStatusRequest) HasLegalBasis() bool {
	if o != nil && o.LegalBasis != nil {
		return true
	}

	return false
}

// SetLegalBasis gets a reference to the given string and assigns it to the LegalBasis field.
func (o *PublicUpdateSubscriptionStatusRequest) SetLegalBasis(v string) {
	o.LegalBasis = &v
}

// GetLegalBasisExplanation returns the LegalBasisExplanation field value if set, zero value otherwise.
func (o *PublicUpdateSubscriptionStatusRequest) GetLegalBasisExplanation() string {
	if o == nil || o.LegalBasisExplanation == nil {
		var ret string
		return ret
	}
	return *o.LegalBasisExplanation
}

// GetLegalBasisExplanationOk returns a tuple with the LegalBasisExplanation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicUpdateSubscriptionStatusRequest) GetLegalBasisExplanationOk() (*string, bool) {
	if o == nil || o.LegalBasisExplanation == nil {
		return nil, false
	}
	return o.LegalBasisExplanation, true
}

// HasLegalBasisExplanation returns a boolean if a field has been set.
func (o *PublicUpdateSubscriptionStatusRequest) HasLegalBasisExplanation() bool {
	if o != nil && o.LegalBasisExplanation != nil {
		return true
	}

	return false
}

// SetLegalBasisExplanation gets a reference to the given string and assigns it to the LegalBasisExplanation field.
func (o *PublicUpdateSubscriptionStatusRequest) SetLegalBasisExplanation(v string) {
	o.LegalBasisExplanation = &v
}

func (o PublicUpdateSubscriptionStatusRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	if true {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if o.LegalBasis != nil {
		toSerialize["legalBasis"] = o.LegalBasis
	}
	if o.LegalBasisExplanation != nil {
		toSerialize["legalBasisExplanation"] = o.LegalBasisExplanation
	}
	return json.Marshal(toSerialize)
}

type NullablePublicUpdateSubscriptionStatusRequest struct {
	value *PublicUpdateSubscriptionStatusRequest
	isSet bool
}

func (v NullablePublicUpdateSubscriptionStatusRequest) Get() *PublicUpdateSubscriptionStatusRequest {
	return v.value
}

func (v *NullablePublicUpdateSubscriptionStatusRequest) Set(val *PublicUpdateSubscriptionStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicUpdateSubscriptionStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicUpdateSubscriptionStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicUpdateSubscriptionStatusRequest(val *PublicUpdateSubscriptionStatusRequest) *NullablePublicUpdateSubscriptionStatusRequest {
	return &NullablePublicUpdateSubscriptionStatusRequest{value: val, isSet: true}
}

func (v NullablePublicUpdateSubscriptionStatusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicUpdateSubscriptionStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
