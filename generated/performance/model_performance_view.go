/*
CMS Performance API

Use these endpoints to get a time series view of your website's performance.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package performance

import (
	"encoding/json"
)

// PerformanceView An individual time series performance data point.
type PerformanceView struct {
	// The timestamp in milliseconds of the start of this interval.
	StartTimestamp int64 `json:"startTimestamp"`
	// The timestamp in milliseconds of the end of this interval.
	EndTimestamp  int64  `json:"endTimestamp"`
	StartDatetime string `json:"startDatetime"`
	EndDatetime   string `json:"endDatetime"`
	// The total number of requests received in this period.
	TotalRequests int32 `json:"totalRequests"`
	// The total number of requests that were served cached responses.
	CacheHits int32 `json:"cacheHits"`
	// The percentage of requests that were served cached responses.
	CacheHitRate     float32 `json:"cacheHitRate"`
	TotalRequestTime int32   `json:"totalRequestTime"`
	// The average response time in milliseconds from the origin to the edge.
	AvgOriginResponseTime int32 `json:"avgOriginResponseTime"`
	// The average response time in milliseconds.
	ResponseTimeMs int32 `json:"responseTimeMs"`
	// The number of responses that had an http status code between 1000-1999.
	Var100X int32 `json:"100X"`
	// The number of responses that had an http status code between 200-299.
	Var20X int32 `json:"20X"`
	// The number of responses that had an http status code between 300-399.
	Var30X int32 `json:"30X"`
	// The number of responses that had an http status code between 400-499.
	Var40X int32 `json:"40X"`
	// The number of responses that had an http status code between 500-599.
	Var50X int32 `json:"50X"`
	// The number of responses that had an http status code of 403.
	Var403 int32 `json:"403"`
	// The number of responses that had an http status code of 404.
	Var404 int32 `json:"404"`
	// The number of responses that had an http status code of 500.
	Var500 int32 `json:"500"`
	// The number of responses that had an http status code of 504.
	Var504 int32 `json:"504"`
	// The 50th percentile response time.
	Var50th int32 `json:"50th"`
	// The 95th percentile response time.
	Var95th int32 `json:"95th"`
	// The 99th percentile response time.
	Var99th int32 `json:"99th"`
}

// NewPerformanceView instantiates a new PerformanceView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerformanceView(startTimestamp int64, endTimestamp int64, startDatetime string, endDatetime string, totalRequests int32, cacheHits int32, cacheHitRate float32, totalRequestTime int32, avgOriginResponseTime int32, responseTimeMs int32, var100X int32, var20X int32, var30X int32, var40X int32, var50X int32, var403 int32, var404 int32, var500 int32, var504 int32, var50th int32, var95th int32, var99th int32) *PerformanceView {
	this := PerformanceView{}
	this.StartTimestamp = startTimestamp
	this.EndTimestamp = endTimestamp
	this.StartDatetime = startDatetime
	this.EndDatetime = endDatetime
	this.TotalRequests = totalRequests
	this.CacheHits = cacheHits
	this.CacheHitRate = cacheHitRate
	this.TotalRequestTime = totalRequestTime
	this.AvgOriginResponseTime = avgOriginResponseTime
	this.ResponseTimeMs = responseTimeMs
	this.Var100X = var100X
	this.Var20X = var20X
	this.Var30X = var30X
	this.Var40X = var40X
	this.Var50X = var50X
	this.Var403 = var403
	this.Var404 = var404
	this.Var500 = var500
	this.Var504 = var504
	this.Var50th = var50th
	this.Var95th = var95th
	this.Var99th = var99th
	return &this
}

// NewPerformanceViewWithDefaults instantiates a new PerformanceView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerformanceViewWithDefaults() *PerformanceView {
	this := PerformanceView{}
	return &this
}

// GetStartTimestamp returns the StartTimestamp field value
func (o *PerformanceView) GetStartTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StartTimestamp
}

// GetStartTimestampOk returns a tuple with the StartTimestamp field value
// and a boolean to check if the value has been set.
func (o *PerformanceView) GetStartTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTimestamp, true
}

// SetStartTimestamp sets field value
func (o *PerformanceView) SetStartTimestamp(v int64) {
	o.StartTimestamp = v
}

// GetEndTimestamp returns the EndTimestamp field value
func (o *PerformanceView) GetEndTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EndTimestamp
}

// GetEndTimestampOk returns a tuple with the EndTimestamp field value
// and a boolean to check if the value has been set.
func (o *PerformanceView) GetEndTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTimestamp, true
}

// SetEndTimestamp sets field value
func (o *PerformanceView) SetEndTimestamp(v int64) {
	o.EndTimestamp = v
}

// GetStartDatetime returns the StartDatetime field value
func (o *PerformanceView) GetStartDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDatetime
}

// GetStartDatetimeOk returns a tuple with the StartDatetime field value
// and a boolean to check if the value has been set.
func (o *PerformanceView) GetStartDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDatetime, true
}

// SetStartDatetime sets field value
func (o *PerformanceView) SetStartDatetime(v string) {
	o.StartDatetime = v
}

// GetEndDatetime returns the EndDatetime field value
func (o *PerformanceView) GetEndDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndDatetime
}

// GetEndDatetimeOk returns a tuple with the EndDatetime field value
// and a boolean to check if the value has been set.
func (o *PerformanceView) GetEndDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDatetime, true
}

// SetEndDatetime sets field value
func (o *PerformanceView) SetEndDatetime(v string) {
	o.EndDatetime = v
}

// GetTotalRequests returns the TotalRequests field value
func (o *PerformanceView) GetTotalRequests() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalRequests
}

// GetTotalRequestsOk returns a tuple with the TotalRequests field value
// and a boolean to check if the value has been set.
func (o *PerformanceView) GetTotalRequestsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalRequests, true
}

// SetTotalRequests sets field value
func (o *PerformanceView) SetTotalRequests(v int32) {
	o.TotalRequests = v
}

// GetCacheHits returns the CacheHits field value
func (o *PerformanceView) GetCacheHits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CacheHits
}

// GetCacheHitsOk returns a tuple with the CacheHits field value
// and a boolean to check if the value has been set.
func (o *PerformanceView) GetCacheHitsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CacheHits, true
}

// SetCacheHits sets field value
func (o *PerformanceView) SetCacheHits(v int32) {
	o.CacheHits = v
}

// GetCacheHitRate returns the CacheHitRate field value
func (o *PerformanceView) GetCacheHitRate() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CacheHitRate
}

// GetCacheHitRateOk returns a tuple with the CacheHitRate field value
// and a boolean to check if the value has been set.
func (o *PerformanceView) GetCacheHitRateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CacheHitRate, true
}

// SetCacheHitRate sets field value
func (o *PerformanceView) SetCacheHitRate(v float32) {
	o.CacheHitRate = v
}

// GetTotalRequestTime returns the TotalRequestTime field value
func (o *PerformanceView) GetTotalRequestTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalRequestTime
}

// GetTotalRequestTimeOk returns a tuple with the TotalRequestTime field value
// and a boolean to check if the value has been set.
func (o *PerformanceView) GetTotalRequestTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalRequestTime, true
}

// SetTotalRequestTime sets field value
func (o *PerformanceView) SetTotalRequestTime(v int32) {
	o.TotalRequestTime = v
}

// GetAvgOriginResponseTime returns the AvgOriginResponseTime field value
func (o *PerformanceView) GetAvgOriginResponseTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AvgOriginResponseTime
}

// GetAvgOriginResponseTimeOk returns a tuple with the AvgOriginResponseTime field value
// and a boolean to check if the value has been set.
func (o *PerformanceView) GetAvgOriginResponseTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgOriginResponseTime, true
}

// SetAvgOriginResponseTime sets field value
func (o *PerformanceView) SetAvgOriginResponseTime(v int32) {
	o.AvgOriginResponseTime = v
}

// GetResponseTimeMs returns the ResponseTimeMs field value
func (o *PerformanceView) GetResponseTimeMs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ResponseTimeMs
}

// GetResponseTimeMsOk returns a tuple with the ResponseTimeMs field value
// and a boolean to check if the value has been set.
func (o *PerformanceView) GetResponseTimeMsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseTimeMs, true
}

// SetResponseTimeMs sets field value
func (o *PerformanceView) SetResponseTimeMs(v int32) {
	o.ResponseTimeMs = v
}

// GetVar100X returns the Var100X field value
func (o *PerformanceView) GetVar100X() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Var100X
}

// GetVar100XOk returns a tuple with the Var100X field value
// and a boolean to check if the value has been set.
func (o *PerformanceView) GetVar100XOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Var100X, true
}

// SetVar100X sets field value
func (o *PerformanceView) SetVar100X(v int32) {
	o.Var100X = v
}

// GetVar20X returns the Var20X field value
func (o *PerformanceView) GetVar20X() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Var20X
}

// GetVar20XOk returns a tuple with the Var20X field value
// and a boolean to check if the value has been set.
func (o *PerformanceView) GetVar20XOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Var20X, true
}

// SetVar20X sets field value
func (o *PerformanceView) SetVar20X(v int32) {
	o.Var20X = v
}

// GetVar30X returns the Var30X field value
func (o *PerformanceView) GetVar30X() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Var30X
}

// GetVar30XOk returns a tuple with the Var30X field value
// and a boolean to check if the value has been set.
func (o *PerformanceView) GetVar30XOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Var30X, true
}

// SetVar30X sets field value
func (o *PerformanceView) SetVar30X(v int32) {
	o.Var30X = v
}

// GetVar40X returns the Var40X field value
func (o *PerformanceView) GetVar40X() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Var40X
}

// GetVar40XOk returns a tuple with the Var40X field value
// and a boolean to check if the value has been set.
func (o *PerformanceView) GetVar40XOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Var40X, true
}

// SetVar40X sets field value
func (o *PerformanceView) SetVar40X(v int32) {
	o.Var40X = v
}

// GetVar50X returns the Var50X field value
func (o *PerformanceView) GetVar50X() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Var50X
}

// GetVar50XOk returns a tuple with the Var50X field value
// and a boolean to check if the value has been set.
func (o *PerformanceView) GetVar50XOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Var50X, true
}

// SetVar50X sets field value
func (o *PerformanceView) SetVar50X(v int32) {
	o.Var50X = v
}

// GetVar403 returns the Var403 field value
func (o *PerformanceView) GetVar403() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Var403
}

// GetVar403Ok returns a tuple with the Var403 field value
// and a boolean to check if the value has been set.
func (o *PerformanceView) GetVar403Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Var403, true
}

// SetVar403 sets field value
func (o *PerformanceView) SetVar403(v int32) {
	o.Var403 = v
}

// GetVar404 returns the Var404 field value
func (o *PerformanceView) GetVar404() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Var404
}

// GetVar404Ok returns a tuple with the Var404 field value
// and a boolean to check if the value has been set.
func (o *PerformanceView) GetVar404Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Var404, true
}

// SetVar404 sets field value
func (o *PerformanceView) SetVar404(v int32) {
	o.Var404 = v
}

// GetVar500 returns the Var500 field value
func (o *PerformanceView) GetVar500() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Var500
}

// GetVar500Ok returns a tuple with the Var500 field value
// and a boolean to check if the value has been set.
func (o *PerformanceView) GetVar500Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Var500, true
}

// SetVar500 sets field value
func (o *PerformanceView) SetVar500(v int32) {
	o.Var500 = v
}

// GetVar504 returns the Var504 field value
func (o *PerformanceView) GetVar504() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Var504
}

// GetVar504Ok returns a tuple with the Var504 field value
// and a boolean to check if the value has been set.
func (o *PerformanceView) GetVar504Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Var504, true
}

// SetVar504 sets field value
func (o *PerformanceView) SetVar504(v int32) {
	o.Var504 = v
}

// GetVar50th returns the Var50th field value
func (o *PerformanceView) GetVar50th() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Var50th
}

// GetVar50thOk returns a tuple with the Var50th field value
// and a boolean to check if the value has been set.
func (o *PerformanceView) GetVar50thOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Var50th, true
}

// SetVar50th sets field value
func (o *PerformanceView) SetVar50th(v int32) {
	o.Var50th = v
}

// GetVar95th returns the Var95th field value
func (o *PerformanceView) GetVar95th() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Var95th
}

// GetVar95thOk returns a tuple with the Var95th field value
// and a boolean to check if the value has been set.
func (o *PerformanceView) GetVar95thOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Var95th, true
}

// SetVar95th sets field value
func (o *PerformanceView) SetVar95th(v int32) {
	o.Var95th = v
}

// GetVar99th returns the Var99th field value
func (o *PerformanceView) GetVar99th() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Var99th
}

// GetVar99thOk returns a tuple with the Var99th field value
// and a boolean to check if the value has been set.
func (o *PerformanceView) GetVar99thOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Var99th, true
}

// SetVar99th sets field value
func (o *PerformanceView) SetVar99th(v int32) {
	o.Var99th = v
}

func (o PerformanceView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["startTimestamp"] = o.StartTimestamp
	}
	if true {
		toSerialize["endTimestamp"] = o.EndTimestamp
	}
	if true {
		toSerialize["startDatetime"] = o.StartDatetime
	}
	if true {
		toSerialize["endDatetime"] = o.EndDatetime
	}
	if true {
		toSerialize["totalRequests"] = o.TotalRequests
	}
	if true {
		toSerialize["cacheHits"] = o.CacheHits
	}
	if true {
		toSerialize["cacheHitRate"] = o.CacheHitRate
	}
	if true {
		toSerialize["totalRequestTime"] = o.TotalRequestTime
	}
	if true {
		toSerialize["avgOriginResponseTime"] = o.AvgOriginResponseTime
	}
	if true {
		toSerialize["responseTimeMs"] = o.ResponseTimeMs
	}
	if true {
		toSerialize["100X"] = o.Var100X
	}
	if true {
		toSerialize["20X"] = o.Var20X
	}
	if true {
		toSerialize["30X"] = o.Var30X
	}
	if true {
		toSerialize["40X"] = o.Var40X
	}
	if true {
		toSerialize["50X"] = o.Var50X
	}
	if true {
		toSerialize["403"] = o.Var403
	}
	if true {
		toSerialize["404"] = o.Var404
	}
	if true {
		toSerialize["500"] = o.Var500
	}
	if true {
		toSerialize["504"] = o.Var504
	}
	if true {
		toSerialize["50th"] = o.Var50th
	}
	if true {
		toSerialize["95th"] = o.Var95th
	}
	if true {
		toSerialize["99th"] = o.Var99th
	}
	return json.Marshal(toSerialize)
}

type NullablePerformanceView struct {
	value *PerformanceView
	isSet bool
}

func (v NullablePerformanceView) Get() *PerformanceView {
	return v.value
}

func (v *NullablePerformanceView) Set(val *PerformanceView) {
	v.value = val
	v.isSet = true
}

func (v NullablePerformanceView) IsSet() bool {
	return v.isSet
}

func (v *NullablePerformanceView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerformanceView(val *PerformanceView) *NullablePerformanceView {
	return &NullablePerformanceView{value: val, isSet: true}
}

func (v NullablePerformanceView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerformanceView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
