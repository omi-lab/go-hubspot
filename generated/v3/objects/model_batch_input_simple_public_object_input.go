/*
CRM Objects

CRM objects such as companies, contacts, deals, line items, products, tickets, and quotes are standard objects in HubSpot’s CRM. These core building blocks support custom properties, store critical information, and play a central role in the HubSpot application.  ## Supported Object Types  This API provides access to collections of CRM objects, which return a map of property names to values. Each object type has its own set of default properties, which can be found by exploring the [CRM Object Properties API](https://developers.hubspot.com/docs/methods/crm-properties/crm-properties-overview).  |Object Type |Properties returned by default | |--|--| | `companies` | `name`, `domain` | | `contacts` | `firstname`, `lastname`, `email` | | `deals` | `dealname`, `amount`, `closedate`, `pipeline`, `dealstage` | | `products` | `name`, `description`, `price` | | `tickets` | `content`, `hs_pipeline`, `hs_pipeline_stage`, `hs_ticket_category`, `hs_ticket_priority`, `subject` |  Find a list of all properties for an object type using the [CRM Object Properties](https://developers.hubspot.com/docs/methods/crm-properties/get-properties) API. e.g. `GET https://api.hubapi.com/properties/v2/companies/properties`. Change the properties returned in the response using the `properties` array in the request body.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objects

import (
	"encoding/json"
)

// BatchInputSimplePublicObjectInput struct for BatchInputSimplePublicObjectInput
type BatchInputSimplePublicObjectInput struct {
	Inputs []SimplePublicObjectInput `json:"inputs"`
}

// NewBatchInputSimplePublicObjectInput instantiates a new BatchInputSimplePublicObjectInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchInputSimplePublicObjectInput(inputs []SimplePublicObjectInput) *BatchInputSimplePublicObjectInput {
	this := BatchInputSimplePublicObjectInput{}
	this.Inputs = inputs
	return &this
}

// NewBatchInputSimplePublicObjectInputWithDefaults instantiates a new BatchInputSimplePublicObjectInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchInputSimplePublicObjectInputWithDefaults() *BatchInputSimplePublicObjectInput {
	this := BatchInputSimplePublicObjectInput{}
	return &this
}

// GetInputs returns the Inputs field value
func (o *BatchInputSimplePublicObjectInput) GetInputs() []SimplePublicObjectInput {
	if o == nil {
		var ret []SimplePublicObjectInput
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchInputSimplePublicObjectInput) GetInputsOk() ([]SimplePublicObjectInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchInputSimplePublicObjectInput) SetInputs(v []SimplePublicObjectInput) {
	o.Inputs = v
}

func (o BatchInputSimplePublicObjectInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["inputs"] = o.Inputs
	}
	return json.Marshal(toSerialize)
}

type NullableBatchInputSimplePublicObjectInput struct {
	value *BatchInputSimplePublicObjectInput
	isSet bool
}

func (v NullableBatchInputSimplePublicObjectInput) Get() *BatchInputSimplePublicObjectInput {
	return v.value
}

func (v *NullableBatchInputSimplePublicObjectInput) Set(val *BatchInputSimplePublicObjectInput) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchInputSimplePublicObjectInput) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchInputSimplePublicObjectInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchInputSimplePublicObjectInput(val *BatchInputSimplePublicObjectInput) *NullableBatchInputSimplePublicObjectInput {
	return &NullableBatchInputSimplePublicObjectInput{value: val, isSet: true}
}

func (v NullableBatchInputSimplePublicObjectInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchInputSimplePublicObjectInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


