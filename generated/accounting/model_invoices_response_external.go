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

// InvoicesResponseExternal A response to a request for invoices.
type InvoicesResponseExternal struct {
	// Designates if the response is a success ('OK') or failure ('ERR').
	Result *string `json:"@result,omitempty"`
	// The list of invoices that were found for the request.
	Invoices []AccountingExtensionInvoice `json:"invoices"`
}

// NewInvoicesResponseExternal instantiates a new InvoicesResponseExternal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoicesResponseExternal(invoices []AccountingExtensionInvoice) *InvoicesResponseExternal {
	this := InvoicesResponseExternal{}
	this.Invoices = invoices
	return &this
}

// NewInvoicesResponseExternalWithDefaults instantiates a new InvoicesResponseExternal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoicesResponseExternalWithDefaults() *InvoicesResponseExternal {
	this := InvoicesResponseExternal{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *InvoicesResponseExternal) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicesResponseExternal) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *InvoicesResponseExternal) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *InvoicesResponseExternal) SetResult(v string) {
	o.Result = &v
}

// GetInvoices returns the Invoices field value
func (o *InvoicesResponseExternal) GetInvoices() []AccountingExtensionInvoice {
	if o == nil {
		var ret []AccountingExtensionInvoice
		return ret
	}

	return o.Invoices
}

// GetInvoicesOk returns a tuple with the Invoices field value
// and a boolean to check if the value has been set.
func (o *InvoicesResponseExternal) GetInvoicesOk() (*[]AccountingExtensionInvoice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Invoices, true
}

// SetInvoices sets field value
func (o *InvoicesResponseExternal) SetInvoices(v []AccountingExtensionInvoice) {
	o.Invoices = v
}

func (o InvoicesResponseExternal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Result != nil {
		toSerialize["@result"] = o.Result
	}
	if true {
		toSerialize["invoices"] = o.Invoices
	}
	return json.Marshal(toSerialize)
}

type NullableInvoicesResponseExternal struct {
	value *InvoicesResponseExternal
	isSet bool
}

func (v NullableInvoicesResponseExternal) Get() *InvoicesResponseExternal {
	return v.value
}

func (v *NullableInvoicesResponseExternal) Set(val *InvoicesResponseExternal) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoicesResponseExternal) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoicesResponseExternal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoicesResponseExternal(val *InvoicesResponseExternal) *NullableInvoicesResponseExternal {
	return &NullableInvoicesResponseExternal{value: val, isSet: true}
}

func (v NullableInvoicesResponseExternal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoicesResponseExternal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
