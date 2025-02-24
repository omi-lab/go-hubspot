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

// ImportResult The result of import operation
type ImportResult struct {
	// List of errors during import
	Errors []Error `json:"errors"`
	// Specifies number of rows imported
	RowsImported int32 `json:"rowsImported"`
	// Specifies number of duplicate rows
	DuplicateRows int32 `json:"duplicateRows"`
	// Specifies whether row limit exceeded during import
	RowLimitExceeded bool `json:"rowLimitExceeded"`
}

// NewImportResult instantiates a new ImportResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportResult(errors []Error, rowsImported int32, duplicateRows int32, rowLimitExceeded bool) *ImportResult {
	this := ImportResult{}
	this.Errors = errors
	this.RowsImported = rowsImported
	this.DuplicateRows = duplicateRows
	this.RowLimitExceeded = rowLimitExceeded
	return &this
}

// NewImportResultWithDefaults instantiates a new ImportResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportResultWithDefaults() *ImportResult {
	this := ImportResult{}
	return &this
}

// GetErrors returns the Errors field value
func (o *ImportResult) GetErrors() []Error {
	if o == nil {
		var ret []Error
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *ImportResult) GetErrorsOk() ([]Error, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *ImportResult) SetErrors(v []Error) {
	o.Errors = v
}

// GetRowsImported returns the RowsImported field value
func (o *ImportResult) GetRowsImported() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RowsImported
}

// GetRowsImportedOk returns a tuple with the RowsImported field value
// and a boolean to check if the value has been set.
func (o *ImportResult) GetRowsImportedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RowsImported, true
}

// SetRowsImported sets field value
func (o *ImportResult) SetRowsImported(v int32) {
	o.RowsImported = v
}

// GetDuplicateRows returns the DuplicateRows field value
func (o *ImportResult) GetDuplicateRows() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DuplicateRows
}

// GetDuplicateRowsOk returns a tuple with the DuplicateRows field value
// and a boolean to check if the value has been set.
func (o *ImportResult) GetDuplicateRowsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DuplicateRows, true
}

// SetDuplicateRows sets field value
func (o *ImportResult) SetDuplicateRows(v int32) {
	o.DuplicateRows = v
}

// GetRowLimitExceeded returns the RowLimitExceeded field value
func (o *ImportResult) GetRowLimitExceeded() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RowLimitExceeded
}

// GetRowLimitExceededOk returns a tuple with the RowLimitExceeded field value
// and a boolean to check if the value has been set.
func (o *ImportResult) GetRowLimitExceededOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RowLimitExceeded, true
}

// SetRowLimitExceeded sets field value
func (o *ImportResult) SetRowLimitExceeded(v bool) {
	o.RowLimitExceeded = v
}

func (o ImportResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["errors"] = o.Errors
	}
	if true {
		toSerialize["rowsImported"] = o.RowsImported
	}
	if true {
		toSerialize["duplicateRows"] = o.DuplicateRows
	}
	if true {
		toSerialize["rowLimitExceeded"] = o.RowLimitExceeded
	}
	return json.Marshal(toSerialize)
}

type NullableImportResult struct {
	value *ImportResult
	isSet bool
}

func (v NullableImportResult) Get() *ImportResult {
	return v.value
}

func (v *NullableImportResult) Set(val *ImportResult) {
	v.value = val
	v.isSet = true
}

func (v NullableImportResult) IsSet() bool {
	return v.isSet
}

func (v *NullableImportResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportResult(val *ImportResult) *NullableImportResult {
	return &NullableImportResult{value: val, isSet: true}
}

func (v NullableImportResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


