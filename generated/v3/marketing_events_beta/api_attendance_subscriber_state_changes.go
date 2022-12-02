/*
Marketing Events Extension

These APIs allow you to interact with HubSpot's Marketing Events Extension. It allows you to: * Create, Read or update Marketing Event information in HubSpot * Specify whether a HubSpot contact has registered, attended or cancelled a registration to a Marketing Event. * Specify a URL that can be called to get the details of a Marketing Event.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marketing_events_beta

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"github.com/omi-lab/go-hubspot"
	"net/url"
	"strings"
)

// AttendanceSubscriberStateChangesApiService AttendanceSubscriberStateChangesApi service
type AttendanceSubscriberStateChangesApiService service

type ApiAttendanceCreateRequest struct {
	ctx                                context.Context
	ApiService                         *AttendanceSubscriberStateChangesApiService
	externalEventId                    string
	subscriberState                    string
	batchInputMarketingEventSubscriber *BatchInputMarketingEventSubscriber
	externalAccountId                  *string
}

// The details of the contacts to subscribe to the event. Parameters of join and left time if state is Attended.
func (r ApiAttendanceCreateRequest) BatchInputMarketingEventSubscriber(batchInputMarketingEventSubscriber BatchInputMarketingEventSubscriber) ApiAttendanceCreateRequest {
	r.batchInputMarketingEventSubscriber = &batchInputMarketingEventSubscriber
	return r
}

// The account id associated with the marketing event
func (r ApiAttendanceCreateRequest) ExternalAccountId(externalAccountId string) ApiAttendanceCreateRequest {
	r.externalAccountId = &externalAccountId
	return r
}

func (r ApiAttendanceCreateRequest) Execute() (*BatchResponseSubscriberVidResponse, *http.Response, error) {
	return r.ApiService.AttendanceCreateExecute(r)
}

/*
AttendanceCreate Record

Record a subscription state between multiple HubSpot contacts and a marketing event, using HubSpot contact ids.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalEventId The id of the marketing event
 @param subscriberState The new subscriber state for the HubSpot contacts and the specified marketing event
 @return ApiAttendanceCreateRequest
*/
func (a *AttendanceSubscriberStateChangesApiService) AttendanceCreate(ctx context.Context, externalEventId string, subscriberState string) ApiAttendanceCreateRequest {
	return ApiAttendanceCreateRequest{
		ApiService:      a,
		ctx:             ctx,
		externalEventId: externalEventId,
		subscriberState: subscriberState,
	}
}

// Execute executes the request
//  @return BatchResponseSubscriberVidResponse
func (a *AttendanceSubscriberStateChangesApiService) AttendanceCreateExecute(r ApiAttendanceCreateRequest) (*BatchResponseSubscriberVidResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BatchResponseSubscriberVidResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AttendanceSubscriberStateChangesApiService.AttendanceCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/marketing/v3/marketing-events/attendance/{externalEventId}/{subscriberState}/create"
	localVarPath = strings.Replace(localVarPath, "{"+"externalEventId"+"}", url.PathEscape(parameterToString(r.externalEventId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscriberState"+"}", url.PathEscape(parameterToString(r.subscriberState, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputMarketingEventSubscriber == nil {
		return localVarReturnValue, nil, reportError("batchInputMarketingEventSubscriber is required and must be specified")
	}

	if r.externalAccountId != nil {
		localVarQueryParams.Add("externalAccountId", parameterToString(*r.externalAccountId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.batchInputMarketingEventSubscriber
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiAttendanceEmailCreateRequest struct {
	ctx                                     context.Context
	ApiService                              *AttendanceSubscriberStateChangesApiService
	externalEventId                         string
	subscriberState                         string
	batchInputMarketingEventEmailSubscriber *BatchInputMarketingEventEmailSubscriber
	externalAccountId                       *string
}

// The details of the contacts to subscribe to the event. Parameters of join and left time if state is Attended.
func (r ApiAttendanceEmailCreateRequest) BatchInputMarketingEventEmailSubscriber(batchInputMarketingEventEmailSubscriber BatchInputMarketingEventEmailSubscriber) ApiAttendanceEmailCreateRequest {
	r.batchInputMarketingEventEmailSubscriber = &batchInputMarketingEventEmailSubscriber
	return r
}

// The account id associated with the marketing event
func (r ApiAttendanceEmailCreateRequest) ExternalAccountId(externalAccountId string) ApiAttendanceEmailCreateRequest {
	r.externalAccountId = &externalAccountId
	return r
}

func (r ApiAttendanceEmailCreateRequest) Execute() (*BatchResponseSubscriberEmailResponse, *http.Response, error) {
	return r.ApiService.AttendanceEmailCreateExecute(r)
}

/*
AttendanceEmailCreate Record

Record a subscription state between multiple HubSpot contacts and a marketing event, using contact email addresses. If contact is not present it will be automatically created. If you set params

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalEventId The id of the marketing event
 @param subscriberState The new subscriber state for the HubSpot contacts and the specified marketing event
 @return ApiAttendanceEmailCreateRequest
*/
func (a *AttendanceSubscriberStateChangesApiService) AttendanceEmailCreate(ctx context.Context, externalEventId string, subscriberState string) ApiAttendanceEmailCreateRequest {
	return ApiAttendanceEmailCreateRequest{
		ApiService:      a,
		ctx:             ctx,
		externalEventId: externalEventId,
		subscriberState: subscriberState,
	}
}

// Execute executes the request
//  @return BatchResponseSubscriberEmailResponse
func (a *AttendanceSubscriberStateChangesApiService) AttendanceEmailCreateExecute(r ApiAttendanceEmailCreateRequest) (*BatchResponseSubscriberEmailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BatchResponseSubscriberEmailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AttendanceSubscriberStateChangesApiService.AttendanceEmailCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/marketing/v3/marketing-events/attendance/{externalEventId}/{subscriberState}/email-create"
	localVarPath = strings.Replace(localVarPath, "{"+"externalEventId"+"}", url.PathEscape(parameterToString(r.externalEventId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscriberState"+"}", url.PathEscape(parameterToString(r.subscriberState, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputMarketingEventEmailSubscriber == nil {
		return localVarReturnValue, nil, reportError("batchInputMarketingEventEmailSubscriber is required and must be specified")
	}

	if r.externalAccountId != nil {
		localVarQueryParams.Add("externalAccountId", parameterToString(*r.externalAccountId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.batchInputMarketingEventEmailSubscriber
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
