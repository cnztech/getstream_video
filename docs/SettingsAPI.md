# \SettingsAPI

All URIs are relative to *https://stream-io-api.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExternalStorage**](SettingsAPI.md#CreateExternalStorage) | **Post** /video/external_storage | Create external storage
[**DeleteExternalStorage**](SettingsAPI.md#DeleteExternalStorage) | **Delete** /video/external_storage/{name} | Delete external storage
[**ListExternalStorage**](SettingsAPI.md#ListExternalStorage) | **Get** /video/external_storage | List external storage



## CreateExternalStorage

> CreateExternalStorageResponse CreateExternalStorage(ctx).CreateExternalStorageRequest(createExternalStorageRequest).Execute()

Create external storage



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
	createExternalStorageRequest := *openapiclient.NewCreateExternalStorageRequest("Bucket_example", "Name_example", "StorageType_example") // CreateExternalStorageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.CreateExternalStorage(context.Background()).CreateExternalStorageRequest(createExternalStorageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.CreateExternalStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExternalStorage`: CreateExternalStorageResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.CreateExternalStorage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExternalStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createExternalStorageRequest** | [**CreateExternalStorageRequest**](CreateExternalStorageRequest.md) |  | 

### Return type

[**CreateExternalStorageResponse**](CreateExternalStorageResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExternalStorage

> DeleteExternalStorageResponse DeleteExternalStorage(ctx, name).Execute()

Delete external storage



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
	resp, r, err := apiClient.SettingsAPI.DeleteExternalStorage(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.DeleteExternalStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteExternalStorage`: DeleteExternalStorageResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.DeleteExternalStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExternalStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteExternalStorageResponse**](DeleteExternalStorageResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExternalStorage

> ListExternalStorageResponse ListExternalStorage(ctx).Execute()

List external storage



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
	resp, r, err := apiClient.SettingsAPI.ListExternalStorage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.ListExternalStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListExternalStorage`: ListExternalStorageResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.ListExternalStorage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListExternalStorageRequest struct via the builder pattern


### Return type

[**ListExternalStorageResponse**](ListExternalStorageResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

