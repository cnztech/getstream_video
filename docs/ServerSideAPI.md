# \ServerSideAPI

All URIs are relative to *https://stream-io-api.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckExternalStorage**](ServerSideAPI.md#CheckExternalStorage) | **Get** /video/external_storage/{name}/check | Check External Storage
[**CreateCallType**](ServerSideAPI.md#CreateCallType) | **Post** /video/calltypes | Create Call Type
[**DeleteCallType**](ServerSideAPI.md#DeleteCallType) | **Delete** /video/calltypes/{name} | Delete Call Type
[**GetCallType**](ServerSideAPI.md#GetCallType) | **Get** /video/calltypes/{name} | Get Call Type
[**ListCallTypes**](ServerSideAPI.md#ListCallTypes) | **Get** /video/calltypes | List Call Type
[**UpdateCallType**](ServerSideAPI.md#UpdateCallType) | **Put** /video/calltypes/{name} | Update Call Type
[**UpdateExternalStorage**](ServerSideAPI.md#UpdateExternalStorage) | **Put** /video/external_storage/{name} | Update External Storage



## CheckExternalStorage

> CheckExternalStorageResponse CheckExternalStorage(ctx, name).Execute()

Check External Storage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cnztech/getstream_video"
)

func main() {
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerSideAPI.CheckExternalStorage(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerSideAPI.CheckExternalStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckExternalStorage`: CheckExternalStorageResponse
	fmt.Fprintf(os.Stdout, "Response from `ServerSideAPI.CheckExternalStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckExternalStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CheckExternalStorageResponse**](CheckExternalStorageResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCallType

> CreateCallTypeResponse CreateCallType(ctx).CreateCallTypeRequest(createCallTypeRequest).Execute()

Create Call Type



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cnztech/getstream_video"
)

func main() {
	createCallTypeRequest := *openapiclient.NewCreateCallTypeRequest("Name_example") // CreateCallTypeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerSideAPI.CreateCallType(context.Background()).CreateCallTypeRequest(createCallTypeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerSideAPI.CreateCallType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCallType`: CreateCallTypeResponse
	fmt.Fprintf(os.Stdout, "Response from `ServerSideAPI.CreateCallType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCallTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCallTypeRequest** | [**CreateCallTypeRequest**](CreateCallTypeRequest.md) |  | 

### Return type

[**CreateCallTypeResponse**](CreateCallTypeResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCallType

> Response DeleteCallType(ctx, name).Execute()

Delete Call Type



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cnztech/getstream_video"
)

func main() {
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerSideAPI.DeleteCallType(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerSideAPI.DeleteCallType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCallType`: Response
	fmt.Fprintf(os.Stdout, "Response from `ServerSideAPI.DeleteCallType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCallTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Response**](Response.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCallType

> GetCallTypeResponse GetCallType(ctx, name).Execute()

Get Call Type



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cnztech/getstream_video"
)

func main() {
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerSideAPI.GetCallType(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerSideAPI.GetCallType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCallType`: GetCallTypeResponse
	fmt.Fprintf(os.Stdout, "Response from `ServerSideAPI.GetCallType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCallTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCallTypeResponse**](GetCallTypeResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCallTypes

> ListCallTypeResponse ListCallTypes(ctx).Execute()

List Call Type



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cnztech/getstream_video"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerSideAPI.ListCallTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerSideAPI.ListCallTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCallTypes`: ListCallTypeResponse
	fmt.Fprintf(os.Stdout, "Response from `ServerSideAPI.ListCallTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCallTypesRequest struct via the builder pattern


### Return type

[**ListCallTypeResponse**](ListCallTypeResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCallType

> UpdateCallTypeResponse UpdateCallType(ctx, name).UpdateCallTypeRequest(updateCallTypeRequest).Execute()

Update Call Type



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cnztech/getstream_video"
)

func main() {
	name := "name_example" // string | 
	updateCallTypeRequest := *openapiclient.NewUpdateCallTypeRequest() // UpdateCallTypeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerSideAPI.UpdateCallType(context.Background(), name).UpdateCallTypeRequest(updateCallTypeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerSideAPI.UpdateCallType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCallType`: UpdateCallTypeResponse
	fmt.Fprintf(os.Stdout, "Response from `ServerSideAPI.UpdateCallType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCallTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCallTypeRequest** | [**UpdateCallTypeRequest**](UpdateCallTypeRequest.md) |  | 

### Return type

[**UpdateCallTypeResponse**](UpdateCallTypeResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExternalStorage

> UpdateExternalStorageResponse UpdateExternalStorage(ctx, name).UpdateExternalStorageRequest(updateExternalStorageRequest).Execute()

Update External Storage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cnztech/getstream_video"
)

func main() {
	name := "name_example" // string | 
	updateExternalStorageRequest := *openapiclient.NewUpdateExternalStorageRequest("Bucket_example", "StorageType_example") // UpdateExternalStorageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerSideAPI.UpdateExternalStorage(context.Background(), name).UpdateExternalStorageRequest(updateExternalStorageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerSideAPI.UpdateExternalStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateExternalStorage`: UpdateExternalStorageResponse
	fmt.Fprintf(os.Stdout, "Response from `ServerSideAPI.UpdateExternalStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExternalStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateExternalStorageRequest** | [**UpdateExternalStorageRequest**](UpdateExternalStorageRequest.md) |  | 

### Return type

[**UpdateExternalStorageResponse**](UpdateExternalStorageResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

