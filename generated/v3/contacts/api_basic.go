/*
Contacts

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package contacts

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	
	"github.com/omi-lab/go-hubspot"
"net/url"
	"strings"
	"reflect"
)


// BasicApiService BasicApi service
type BasicApiService service

type ApiArchiveRequest struct {
	ctx context.Context
	ApiService *BasicApiService
	contactId string
}

func (r ApiArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.ArchiveExecute(r)
}

/*
Archive Archive

Move an Object identified by `{contactId}` to the recycling bin.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param contactId
 @return ApiArchiveRequest
*/
func (a *BasicApiService) Archive(ctx context.Context, contactId string) ApiArchiveRequest {
	return ApiArchiveRequest{
		ApiService: a,
		ctx: ctx,
		contactId: contactId,
	}
}

// Execute executes the request
func (a *BasicApiService) ArchiveExecute(r ApiArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BasicApiService.Archive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/contacts/{contactId}"
	localVarPath = strings.Replace(localVarPath, "{"+"contactId"+"}", url.PathEscape(parameterToString(r.contactId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiCreateRequest struct {
	ctx context.Context
	ApiService *BasicApiService
	simplePublicObjectInput *SimplePublicObjectInput
}

func (r ApiCreateRequest) SimplePublicObjectInput(simplePublicObjectInput SimplePublicObjectInput) ApiCreateRequest {
	r.simplePublicObjectInput = &simplePublicObjectInput
	return r
}

func (r ApiCreateRequest) Execute() (*SimplePublicObject, *http.Response, error) {
	return r.ApiService.CreateExecute(r)
}

/*
Create Create

Create a contact with the given properties and return a copy of the object, including the ID. Documentation and examples for creating standard contacts is provided.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateRequest
*/
func (a *BasicApiService) Create(ctx context.Context) ApiCreateRequest {
	return ApiCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SimplePublicObject
func (a *BasicApiService) CreateExecute(r ApiCreateRequest) (*SimplePublicObject, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SimplePublicObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BasicApiService.Create")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/contacts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.simplePublicObjectInput == nil {
		return localVarReturnValue, nil, reportError("simplePublicObjectInput is required and must be specified")
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
	localVarPostBody = r.simplePublicObjectInput
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

type ApiGetByIDRequest struct {
	ctx context.Context
	ApiService *BasicApiService
	contactId string
	properties *[]string
	propertiesWithHistory *[]string
	associations *[]string
	archived *bool
	idProperty *string
}

// A comma separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored.
func (r ApiGetByIDRequest) Properties(properties []string) ApiGetByIDRequest {
	r.properties = &properties
	return r
}

// A comma separated list of the properties to be returned along with their history of previous values. If any of the specified properties are not present on the requested object(s), they will be ignored.
func (r ApiGetByIDRequest) PropertiesWithHistory(propertiesWithHistory []string) ApiGetByIDRequest {
	r.propertiesWithHistory = &propertiesWithHistory
	return r
}

// A comma separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored.
func (r ApiGetByIDRequest) Associations(associations []string) ApiGetByIDRequest {
	r.associations = &associations
	return r
}

// Whether to return only results that have been archived.
func (r ApiGetByIDRequest) Archived(archived bool) ApiGetByIDRequest {
	r.archived = &archived
	return r
}

// The name of a property whose values are unique for this object type
func (r ApiGetByIDRequest) IdProperty(idProperty string) ApiGetByIDRequest {
	r.idProperty = &idProperty
	return r
}

func (r ApiGetByIDRequest) Execute() (*SimplePublicObjectWithAssociations, *http.Response, error) {
	return r.ApiService.GetByIDExecute(r)
}

/*
GetByID Read

Read an Object identified by `{contactId}`. `{contactId}` refers to the internal object ID.  Control what is returned via the `properties` query param.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param contactId
 @return ApiGetByIDRequest
*/
func (a *BasicApiService) GetByID(ctx context.Context, contactId string) ApiGetByIDRequest {
	return ApiGetByIDRequest{
		ApiService: a,
		ctx: ctx,
		contactId: contactId,
	}
}

// Execute executes the request
//  @return SimplePublicObjectWithAssociations
func (a *BasicApiService) GetByIDExecute(r ApiGetByIDRequest) (*SimplePublicObjectWithAssociations, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SimplePublicObjectWithAssociations
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BasicApiService.GetByID")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/contacts/{contactId}"
	localVarPath = strings.Replace(localVarPath, "{"+"contactId"+"}", url.PathEscape(parameterToString(r.contactId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.properties != nil {
		t := *r.properties
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("properties", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("properties", parameterToString(t, "multi"))
		}
	}
	if r.propertiesWithHistory != nil {
		t := *r.propertiesWithHistory
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("propertiesWithHistory", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("propertiesWithHistory", parameterToString(t, "multi"))
		}
	}
	if r.associations != nil {
		t := *r.associations
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("associations", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("associations", parameterToString(t, "multi"))
		}
	}
	if r.archived != nil {
		localVarQueryParams.Add("archived", parameterToString(*r.archived, ""))
	}
	if r.idProperty != nil {
		localVarQueryParams.Add("idProperty", parameterToString(*r.idProperty, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiGetPageRequest struct {
	ctx context.Context
	ApiService *BasicApiService
	limit *int32
	after *string
	properties *[]string
	propertiesWithHistory *[]string
	associations *[]string
	archived *bool
}

// The maximum number of results to display per page.
func (r ApiGetPageRequest) Limit(limit int32) ApiGetPageRequest {
	r.limit = &limit
	return r
}

// The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results.
func (r ApiGetPageRequest) After(after string) ApiGetPageRequest {
	r.after = &after
	return r
}

// A comma separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored.
func (r ApiGetPageRequest) Properties(properties []string) ApiGetPageRequest {
	r.properties = &properties
	return r
}

// A comma separated list of the properties to be returned along with their history of previous values. If any of the specified properties are not present on the requested object(s), they will be ignored. Usage of this parameter will reduce the maximum number of objects that can be read by a single request.
func (r ApiGetPageRequest) PropertiesWithHistory(propertiesWithHistory []string) ApiGetPageRequest {
	r.propertiesWithHistory = &propertiesWithHistory
	return r
}

// A comma separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored.
func (r ApiGetPageRequest) Associations(associations []string) ApiGetPageRequest {
	r.associations = &associations
	return r
}

// Whether to return only results that have been archived.
func (r ApiGetPageRequest) Archived(archived bool) ApiGetPageRequest {
	r.archived = &archived
	return r
}

func (r ApiGetPageRequest) Execute() (*CollectionResponseSimplePublicObjectWithAssociationsForwardPaging, *http.Response, error) {
	return r.ApiService.GetPageExecute(r)
}

/*
GetPage List

Read a page of contacts. Control what is returned via the `properties` query param.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPageRequest
*/
func (a *BasicApiService) GetPage(ctx context.Context) ApiGetPageRequest {
	return ApiGetPageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionResponseSimplePublicObjectWithAssociationsForwardPaging
func (a *BasicApiService) GetPageExecute(r ApiGetPageRequest) (*CollectionResponseSimplePublicObjectWithAssociationsForwardPaging, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionResponseSimplePublicObjectWithAssociationsForwardPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BasicApiService.GetPage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/contacts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.properties != nil {
		t := *r.properties
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("properties", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("properties", parameterToString(t, "multi"))
		}
	}
	if r.propertiesWithHistory != nil {
		t := *r.propertiesWithHistory
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("propertiesWithHistory", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("propertiesWithHistory", parameterToString(t, "multi"))
		}
	}
	if r.associations != nil {
		t := *r.associations
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("associations", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("associations", parameterToString(t, "multi"))
		}
	}
	if r.archived != nil {
		localVarQueryParams.Add("archived", parameterToString(*r.archived, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiUpdateRequest struct {
	ctx context.Context
	ApiService *BasicApiService
	contactId string
	simplePublicObjectInput *SimplePublicObjectInput
}

func (r ApiUpdateRequest) SimplePublicObjectInput(simplePublicObjectInput SimplePublicObjectInput) ApiUpdateRequest {
	r.simplePublicObjectInput = &simplePublicObjectInput
	return r
}

func (r ApiUpdateRequest) Execute() (*SimplePublicObject, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update Update

Perform a partial update of an Object identified by `{contactId}`. `{contactId}` refers to the internal object ID. Provided property values will be overwritten. Read-only and non-existent properties will be ignored. Properties values can be cleared by passing an empty string.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param contactId
 @return ApiUpdateRequest
*/
func (a *BasicApiService) Update(ctx context.Context, contactId string) ApiUpdateRequest {
	return ApiUpdateRequest{
		ApiService: a,
		ctx: ctx,
		contactId: contactId,
	}
}

// Execute executes the request
//  @return SimplePublicObject
func (a *BasicApiService) UpdateExecute(r ApiUpdateRequest) (*SimplePublicObject, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SimplePublicObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BasicApiService.Update")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/contacts/{contactId}"
	localVarPath = strings.Replace(localVarPath, "{"+"contactId"+"}", url.PathEscape(parameterToString(r.contactId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.simplePublicObjectInput == nil {
		return localVarReturnValue, nil, reportError("simplePublicObjectInput is required and must be specified")
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
	localVarPostBody = r.simplePublicObjectInput
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
