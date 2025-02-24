/*
Accounting Extension

These APIs allow you to interact with HubSpot's Accounting Extension. It allows you to: * Specify the URLs that HubSpot will use when making webhook requests to your external accounting system. * Respond to webhook calls made to your external accounting system by HubSpot 

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accounting

import (
	"encoding/json"
)

// SyncContactsRequest A request to import external accounting contact properties in HubSpot
type SyncContactsRequest struct {
	// The ID of the account in the external accounting system. This is the value that will be passed as `accountId` for all outbound calls for the user from HubSpot to the external accounting system.
	AccountId string `json:"accountId"`
	// A list of contacts to be imported.
	Contacts []UpdatedContact `json:"contacts"`
}

// NewSyncContactsRequest instantiates a new SyncContactsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncContactsRequest(accountId string, contacts []UpdatedContact) *SyncContactsRequest {
	this := SyncContactsRequest{}
	this.AccountId = accountId
	this.Contacts = contacts
	return &this
}

// NewSyncContactsRequestWithDefaults instantiates a new SyncContactsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncContactsRequestWithDefaults() *SyncContactsRequest {
	this := SyncContactsRequest{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *SyncContactsRequest) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *SyncContactsRequest) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *SyncContactsRequest) SetAccountId(v string) {
	o.AccountId = v
}

// GetContacts returns the Contacts field value
func (o *SyncContactsRequest) GetContacts() []UpdatedContact {
	if o == nil {
		var ret []UpdatedContact
		return ret
	}

	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value
// and a boolean to check if the value has been set.
func (o *SyncContactsRequest) GetContactsOk() ([]UpdatedContact, bool) {
	if o == nil {
		return nil, false
	}
	return o.Contacts, true
}

// SetContacts sets field value
func (o *SyncContactsRequest) SetContacts(v []UpdatedContact) {
	o.Contacts = v
}

func (o SyncContactsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["contacts"] = o.Contacts
	}
	return json.Marshal(toSerialize)
}

type NullableSyncContactsRequest struct {
	value *SyncContactsRequest
	isSet bool
}

func (v NullableSyncContactsRequest) Get() *SyncContactsRequest {
	return v.value
}

func (v *NullableSyncContactsRequest) Set(val *SyncContactsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncContactsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncContactsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncContactsRequest(val *SyncContactsRequest) *NullableSyncContactsRequest {
	return &NullableSyncContactsRequest{value: val, isSet: true}
}

func (v NullableSyncContactsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncContactsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


