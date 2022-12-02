/*
CMS Source Code

Endpoints for interacting with files in the CMS Developer File System. These files include HTML templates, CSS, JS, modules, and other assets which are used to create CMS content.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package source_code

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"github.com/omi-lab/go-hubspot"
	"net/url"
	"strings"
)

// MetadataApiService MetadataApi service
type MetadataApiService service

type ApiMetadataGetRequest struct {
	ctx         context.Context
	ApiService  *MetadataApiService
	environment string
	path        string
}

func (r ApiMetadataGetRequest) Execute() (*AssetFileMetadata, *http.Response, error) {
	return r.ApiService.MetadataGetExecute(r)
}

/*
MetadataGet Get the metadata for a file

Gets the metadata object for the file at the specified path in the specified environment.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environment The environment of the file (\"draft\" or \"published\").
 @param path The file system location of the file.
 @return ApiMetadataGetRequest
*/
func (a *MetadataApiService) MetadataGet(ctx context.Context, environment string, path string) ApiMetadataGetRequest {
	return ApiMetadataGetRequest{
		ApiService:  a,
		ctx:         ctx,
		environment: environment,
		path:        path,
	}
}

// Execute executes the request
//  @return AssetFileMetadata
func (a *MetadataApiService) MetadataGetExecute(r ApiMetadataGetRequest) (*AssetFileMetadata, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AssetFileMetadata
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetadataApiService.MetadataGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cms/v3/source-code/{environment}/metadata/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"environment"+"}", url.PathEscape(parameterToString(r.environment, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", url.PathEscape(parameterToString(r.path, "")), -1)

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
