# \CallbacksApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallbackComplete**](CallbacksApi.md#CallbackComplete) | **Post** /automation/v4/actions/callbacks/{callbackId}/complete | Complete a callback
[**CallbackCompleteBatch**](CallbacksApi.md#CallbackCompleteBatch) | **Post** /automation/v4/actions/callbacks/complete | Complete a batch of callbacks



## CallbackComplete

> CallbackComplete(ctx, callbackId).CallbackCompletionRequest(callbackCompletionRequest).Execute()

Complete a callback



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    callbackId := "callbackId_example" // string | The ID of the target app.
    callbackCompletionRequest := *openapiclient.NewCallbackCompletionRequest(map[string]string{"key": "Inner_example"}) // CallbackCompletionRequest | The result of the completed action.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CallbacksApi.CallbackComplete(context.Background(), callbackId).CallbackCompletionRequest(callbackCompletionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbacksApi.CallbackComplete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callbackId** | **string** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCallbackCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **callbackCompletionRequest** | [**CallbackCompletionRequest**](CallbackCompletionRequest.md) | The result of the completed action. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallbackCompleteBatch

> CallbackCompleteBatch(ctx).BatchInputCallbackCompletionBatchRequest(batchInputCallbackCompletionBatchRequest).Execute()

Complete a batch of callbacks



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    batchInputCallbackCompletionBatchRequest := *openapiclient.NewBatchInputCallbackCompletionBatchRequest([]openapiclient.CallbackCompletionBatchRequest{*openapiclient.NewCallbackCompletionBatchRequest("CallbackId_example", map[string]string{"key": "Inner_example"})}) // BatchInputCallbackCompletionBatchRequest | The result of the completed action.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CallbacksApi.CallbackCompleteBatch(context.Background()).BatchInputCallbackCompletionBatchRequest(batchInputCallbackCompletionBatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbacksApi.CallbackCompleteBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallbackCompleteBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputCallbackCompletionBatchRequest** | [**BatchInputCallbackCompletionBatchRequest**](BatchInputCallbackCompletionBatchRequest.md) | The result of the completed action. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

