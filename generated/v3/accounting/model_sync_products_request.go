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

// SyncProductsRequest A request to import external accounting product properties in HubSpot
type SyncProductsRequest struct {
	// The ID of the account in the external accounting system. This is the value that will be passed as `accountId` for all outbound calls for the user from HubSpot to the external accounting system.
	AccountId string `json:"accountId"`
	// A list of products to be imported.
	Products []UpdatedProduct `json:"products"`
}

// NewSyncProductsRequest instantiates a new SyncProductsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncProductsRequest(accountId string, products []UpdatedProduct) *SyncProductsRequest {
	this := SyncProductsRequest{}
	this.AccountId = accountId
	this.Products = products
	return &this
}

// NewSyncProductsRequestWithDefaults instantiates a new SyncProductsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncProductsRequestWithDefaults() *SyncProductsRequest {
	this := SyncProductsRequest{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *SyncProductsRequest) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *SyncProductsRequest) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *SyncProductsRequest) SetAccountId(v string) {
	o.AccountId = v
}

// GetProducts returns the Products field value
func (o *SyncProductsRequest) GetProducts() []UpdatedProduct {
	if o == nil {
		var ret []UpdatedProduct
		return ret
	}

	return o.Products
}

// GetProductsOk returns a tuple with the Products field value
// and a boolean to check if the value has been set.
func (o *SyncProductsRequest) GetProductsOk() ([]UpdatedProduct, bool) {
	if o == nil {
		return nil, false
	}
	return o.Products, true
}

// SetProducts sets field value
func (o *SyncProductsRequest) SetProducts(v []UpdatedProduct) {
	o.Products = v
}

func (o SyncProductsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["products"] = o.Products
	}
	return json.Marshal(toSerialize)
}

type NullableSyncProductsRequest struct {
	value *SyncProductsRequest
	isSet bool
}

func (v NullableSyncProductsRequest) Get() *SyncProductsRequest {
	return v.value
}

func (v *NullableSyncProductsRequest) Set(val *SyncProductsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncProductsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncProductsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncProductsRequest(val *SyncProductsRequest) *NullableSyncProductsRequest {
	return &NullableSyncProductsRequest{value: val, isSet: true}
}

func (v NullableSyncProductsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncProductsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


