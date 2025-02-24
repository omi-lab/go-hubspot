/*
HubDB endpoints

HubDB is a relational data store that presents data as rows, columns, and cells in a table, much like a spreadsheet. HubDB tables can be added or modified [in the HubSpot CMS](https://knowledge.hubspot.com/cos-general/how-to-edit-hubdb-tables), but you can also use the API endpoints documented here. For more information on HubDB tables and using their data on a HubSpot site, see the [CMS developers site](https://designers.hubspot.com/docs/tools/hubdb). You can also see the [documentation for dynamic pages](https://designers.hubspot.com/docs/tutorials/how-to-build-dynamic-pages-with-hubdb) for more details about the `useForPages` field.  HubDB tables support `draft` and `published` versions. This allows you to update data in the table, either for testing or to allow for a manual approval process, without affecting any live pages using the existing data. Draft data can be reviewed, and published by a user working in HubSpot or published via the API. Draft data can also be discarded, allowing users to go back to the published version of the data without disrupting it. If a table is set to be `allowed for public access`, you can access the published version of the table and rows without any authentication by specifying the portal id via the query parameter `portalId`.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hubdb

import (
	"encoding/json"
)

// BatchInputHubDbTableRowV3Request struct for BatchInputHubDbTableRowV3Request
type BatchInputHubDbTableRowV3Request struct {
	Inputs []HubDbTableRowV3Request `json:"inputs"`
}

// NewBatchInputHubDbTableRowV3Request instantiates a new BatchInputHubDbTableRowV3Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchInputHubDbTableRowV3Request(inputs []HubDbTableRowV3Request) *BatchInputHubDbTableRowV3Request {
	this := BatchInputHubDbTableRowV3Request{}
	this.Inputs = inputs
	return &this
}

// NewBatchInputHubDbTableRowV3RequestWithDefaults instantiates a new BatchInputHubDbTableRowV3Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchInputHubDbTableRowV3RequestWithDefaults() *BatchInputHubDbTableRowV3Request {
	this := BatchInputHubDbTableRowV3Request{}
	return &this
}

// GetInputs returns the Inputs field value
func (o *BatchInputHubDbTableRowV3Request) GetInputs() []HubDbTableRowV3Request {
	if o == nil {
		var ret []HubDbTableRowV3Request
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchInputHubDbTableRowV3Request) GetInputsOk() ([]HubDbTableRowV3Request, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchInputHubDbTableRowV3Request) SetInputs(v []HubDbTableRowV3Request) {
	o.Inputs = v
}

func (o BatchInputHubDbTableRowV3Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["inputs"] = o.Inputs
	}
	return json.Marshal(toSerialize)
}

type NullableBatchInputHubDbTableRowV3Request struct {
	value *BatchInputHubDbTableRowV3Request
	isSet bool
}

func (v NullableBatchInputHubDbTableRowV3Request) Get() *BatchInputHubDbTableRowV3Request {
	return v.value
}

func (v *NullableBatchInputHubDbTableRowV3Request) Set(val *BatchInputHubDbTableRowV3Request) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchInputHubDbTableRowV3Request) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchInputHubDbTableRowV3Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchInputHubDbTableRowV3Request(val *BatchInputHubDbTableRowV3Request) *NullableBatchInputHubDbTableRowV3Request {
	return &NullableBatchInputHubDbTableRowV3Request{value: val, isSet: true}
}

func (v NullableBatchInputHubDbTableRowV3Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchInputHubDbTableRowV3Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


