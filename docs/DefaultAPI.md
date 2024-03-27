# \DefaultAPI

All URIs are relative to *https://stream-io-api.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlockUser**](DefaultAPI.md#BlockUser) | **Post** /video/call/{type}/{id}/block | Block user on a call
[**CreateDevice**](DefaultAPI.md#CreateDevice) | **Post** /video/devices | Create device
[**CreateGuest**](DefaultAPI.md#CreateGuest) | **Post** /video/guest | Create Guest
[**DeleteDevice**](DefaultAPI.md#DeleteDevice) | **Delete** /video/devices | Delete device
[**EndCall**](DefaultAPI.md#EndCall) | **Post** /video/call/{type}/{id}/mark_ended | End call
[**GetCall**](DefaultAPI.md#GetCall) | **Get** /video/call/{type}/{id} | Get Call
[**GetCallStats**](DefaultAPI.md#GetCallStats) | **Get** /video/call/{type}/{id}/stats/{session} | Get Call Stats
[**GetEdges**](DefaultAPI.md#GetEdges) | **Get** /video/edges | Get Edges
[**GetOrCreateCall**](DefaultAPI.md#GetOrCreateCall) | **Post** /video/call/{type}/{id} | Get or create a call
[**GoLive**](DefaultAPI.md#GoLive) | **Post** /video/call/{type}/{id}/go_live | Set call as live
[**ListDevices**](DefaultAPI.md#ListDevices) | **Get** /video/devices | List devices
[**ListRecordings**](DefaultAPI.md#ListRecordings) | **Get** /video/call/{type}/{id}/recordings | List recordings
[**ListTranscriptions**](DefaultAPI.md#ListTranscriptions) | **Get** /video/call/{type}/{id}/transcriptions | List transcriptions
[**MuteUsers**](DefaultAPI.md#MuteUsers) | **Post** /video/call/{type}/{id}/mute_users | Mute users
[**QueryCalls**](DefaultAPI.md#QueryCalls) | **Post** /video/calls | Query call
[**QueryMembers**](DefaultAPI.md#QueryMembers) | **Post** /video/call/members | Query call members
[**SendEvent**](DefaultAPI.md#SendEvent) | **Post** /video/call/{type}/{id}/event | Send custom event
[**StartHLSBroadcasting**](DefaultAPI.md#StartHLSBroadcasting) | **Post** /video/call/{type}/{id}/start_broadcasting | Start HLS broadcasting
[**StartRecording**](DefaultAPI.md#StartRecording) | **Post** /video/call/{type}/{id}/start_recording | Start recording
[**StartTranscription**](DefaultAPI.md#StartTranscription) | **Post** /video/call/{type}/{id}/start_transcription | Start transcription
[**StopHLSBroadcasting**](DefaultAPI.md#StopHLSBroadcasting) | **Post** /video/call/{type}/{id}/stop_broadcasting | Stop HLS broadcasting
[**StopLive**](DefaultAPI.md#StopLive) | **Post** /video/call/{type}/{id}/stop_live | Set call as not live
[**StopRecording**](DefaultAPI.md#StopRecording) | **Post** /video/call/{type}/{id}/stop_recording | Stop recording
[**StopTranscription**](DefaultAPI.md#StopTranscription) | **Post** /video/call/{type}/{id}/stop_transcription | Stop transcription
[**UnblockUser**](DefaultAPI.md#UnblockUser) | **Post** /video/call/{type}/{id}/unblock | Unblocks user on a call
[**UpdateCall**](DefaultAPI.md#UpdateCall) | **Patch** /video/call/{type}/{id} | Update Call
[**UpdateCallMembers**](DefaultAPI.md#UpdateCallMembers) | **Post** /video/call/{type}/{id}/members | Update Call Member
[**UpdateUserPermissions**](DefaultAPI.md#UpdateUserPermissions) | **Post** /video/call/{type}/{id}/user_permissions | Update user permissions
[**VideoConnect**](DefaultAPI.md#VideoConnect) | **Get** /video/connect | Video Connect (WebSocket)
[**VideoPin**](DefaultAPI.md#VideoPin) | **Post** /video/call/{type}/{id}/pin | Pin
[**VideoUnpin**](DefaultAPI.md#VideoUnpin) | **Post** /video/call/{type}/{id}/unpin | Unpin



## BlockUser

> BlockUserResponse BlockUser(ctx, type_, id).BlockUserRequest(blockUserRequest).Execute()

Block user on a call



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
	type_ := "type__example" // string | 
	id := "id_example" // string | 
	blockUserRequest := *openapiclient.NewBlockUserRequest("UserId_example") // BlockUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.BlockUser(context.Background(), type_, id).BlockUserRequest(blockUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BlockUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlockUser`: BlockUserResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BlockUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlockUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **blockUserRequest** | [**BlockUserRequest**](BlockUserRequest.md) |  | 

### Return type

[**BlockUserResponse**](BlockUserResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDevice

> Response CreateDevice(ctx).CreateDeviceRequest(createDeviceRequest).Execute()

Create device



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
	createDeviceRequest := *openapiclient.NewCreateDeviceRequest() // CreateDeviceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateDevice(context.Background()).CreateDeviceRequest(createDeviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDevice`: Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateDevice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDeviceRequest** | [**CreateDeviceRequest**](CreateDeviceRequest.md) |  | 

### Return type

[**Response**](Response.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGuest

> CreateGuestResponse CreateGuest(ctx).CreateGuestRequest(createGuestRequest).Execute()

Create Guest



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
	createGuestRequest := *openapiclient.NewCreateGuestRequest(*openapiclient.NewUserRequest("Id_example")) // CreateGuestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateGuest(context.Background()).CreateGuestRequest(createGuestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateGuest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGuest`: CreateGuestResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateGuest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGuestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createGuestRequest** | [**CreateGuestRequest**](CreateGuestRequest.md) |  | 

### Return type

[**CreateGuestResponse**](CreateGuestResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDevice

> Response DeleteDevice(ctx).Id(id).UserId(userId).Execute()

Delete device



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
	id := "id_example" // string |  (optional)
	userId := "userId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteDevice(context.Background()).Id(id).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDevice`: Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteDevice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **userId** | **string** |  | 

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


## EndCall

> EndCallResponse EndCall(ctx, type_, id).Execute()

End call



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
	type_ := "type__example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.EndCall(context.Background(), type_, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EndCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndCall`: EndCallResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.EndCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EndCallResponse**](EndCallResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCall

> GetCallResponse GetCall(ctx, type_, id).ConnectionId(connectionId).MembersLimit(membersLimit).Ring(ring).Notify(notify).Execute()

Get Call



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
	type_ := "type__example" // string | 
	id := "id_example" // string | 
	connectionId := "connectionId_example" // string |  (optional)
	membersLimit := int32(56) // int32 |  (optional)
	ring := true // bool |  (optional)
	notify := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetCall(context.Background(), type_, id).ConnectionId(connectionId).MembersLimit(membersLimit).Ring(ring).Notify(notify).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCall`: GetCallResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **connectionId** | **string** |  | 
 **membersLimit** | **int32** |  | 
 **ring** | **bool** |  | 
 **notify** | **bool** |  | 

### Return type

[**GetCallResponse**](GetCallResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCallStats

> GetCallStatsResponse GetCallStats(ctx, type_, id, session).Execute()

Get Call Stats



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
	type_ := "type__example" // string | 
	id := "id_example" // string | 
	session := "session_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetCallStats(context.Background(), type_, id, session).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetCallStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCallStats`: GetCallStatsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetCallStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**id** | **string** |  | 
**session** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCallStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetCallStatsResponse**](GetCallStatsResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEdges

> GetEdgesResponse GetEdges(ctx).Execute()

Get Edges



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
	resp, r, err := apiClient.DefaultAPI.GetEdges(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetEdges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEdges`: GetEdgesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetEdges`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEdgesRequest struct via the builder pattern


### Return type

[**GetEdgesResponse**](GetEdgesResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrCreateCall

> GetOrCreateCallResponse GetOrCreateCall(ctx, type_, id).GetOrCreateCallRequest(getOrCreateCallRequest).ConnectionId(connectionId).Execute()

Get or create a call



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
	type_ := "type__example" // string | 
	id := "id_example" // string | 
	getOrCreateCallRequest := *openapiclient.NewGetOrCreateCallRequest() // GetOrCreateCallRequest | 
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetOrCreateCall(context.Background(), type_, id).GetOrCreateCallRequest(getOrCreateCallRequest).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetOrCreateCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrCreateCall`: GetOrCreateCallResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetOrCreateCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrCreateCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **getOrCreateCallRequest** | [**GetOrCreateCallRequest**](GetOrCreateCallRequest.md) |  | 
 **connectionId** | **string** |  | 

### Return type

[**GetOrCreateCallResponse**](GetOrCreateCallResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GoLive

> GoLiveResponse GoLive(ctx, type_, id).GoLiveRequest(goLiveRequest).Execute()

Set call as live



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
	type_ := "type__example" // string | 
	id := "id_example" // string | 
	goLiveRequest := *openapiclient.NewGoLiveRequest() // GoLiveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GoLive(context.Background(), type_, id).GoLiveRequest(goLiveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GoLive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GoLive`: GoLiveResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GoLive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGoLiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **goLiveRequest** | [**GoLiveRequest**](GoLiveRequest.md) |  | 

### Return type

[**GoLiveResponse**](GoLiveResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDevices

> ListDevicesResponse ListDevices(ctx).UserId(userId).Execute()

List devices



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
	userId := "userId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListDevices(context.Background()).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDevices`: ListDevicesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** |  | 

### Return type

[**ListDevicesResponse**](ListDevicesResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRecordings

> ListRecordingsResponse ListRecordings(ctx, type_, id).Execute()

List recordings



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
	type_ := "type__example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListRecordings(context.Background(), type_, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListRecordings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRecordings`: ListRecordingsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListRecordings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRecordingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListRecordingsResponse**](ListRecordingsResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTranscriptions

> ListTranscriptionsResponse ListTranscriptions(ctx, type_, id).Execute()

List transcriptions



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
	type_ := "type__example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListTranscriptions(context.Background(), type_, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListTranscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTranscriptions`: ListTranscriptionsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListTranscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTranscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListTranscriptionsResponse**](ListTranscriptionsResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MuteUsers

> MuteUsersResponse MuteUsers(ctx, type_, id).MuteUsersRequest(muteUsersRequest).Execute()

Mute users



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
	type_ := "type__example" // string | 
	id := "id_example" // string | 
	muteUsersRequest := *openapiclient.NewMuteUsersRequest() // MuteUsersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.MuteUsers(context.Background(), type_, id).MuteUsersRequest(muteUsersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MuteUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MuteUsers`: MuteUsersResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.MuteUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMuteUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **muteUsersRequest** | [**MuteUsersRequest**](MuteUsersRequest.md) |  | 

### Return type

[**MuteUsersResponse**](MuteUsersResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryCalls

> QueryCallsResponse QueryCalls(ctx).QueryCallsRequest(queryCallsRequest).ConnectionId(connectionId).Execute()

Query call



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
	queryCallsRequest := *openapiclient.NewQueryCallsRequest() // QueryCallsRequest | 
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QueryCalls(context.Background()).QueryCallsRequest(queryCallsRequest).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QueryCalls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryCalls`: QueryCallsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QueryCalls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryCallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryCallsRequest** | [**QueryCallsRequest**](QueryCallsRequest.md) |  | 
 **connectionId** | **string** |  | 

### Return type

[**QueryCallsResponse**](QueryCallsResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryMembers

> QueryMembersResponse QueryMembers(ctx).QueryMembersRequest(queryMembersRequest).Execute()

Query call members



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
	queryMembersRequest := *openapiclient.NewQueryMembersRequest("Id_example", "Type_example") // QueryMembersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QueryMembers(context.Background()).QueryMembersRequest(queryMembersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QueryMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryMembers`: QueryMembersResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QueryMembers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryMembersRequest** | [**QueryMembersRequest**](QueryMembersRequest.md) |  | 

### Return type

[**QueryMembersResponse**](QueryMembersResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendEvent

> SendEventResponse SendEvent(ctx, type_, id).SendEventRequest(sendEventRequest).Execute()

Send custom event



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
	type_ := "type__example" // string | 
	id := "id_example" // string | 
	sendEventRequest := *openapiclient.NewSendEventRequest() // SendEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.SendEvent(context.Background(), type_, id).SendEventRequest(sendEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SendEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendEvent`: SendEventResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SendEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sendEventRequest** | [**SendEventRequest**](SendEventRequest.md) |  | 

### Return type

[**SendEventResponse**](SendEventResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartHLSBroadcasting

> StartHLSBroadcastingResponse StartHLSBroadcasting(ctx, type_, id).Execute()

Start HLS broadcasting



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
	type_ := "type__example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.StartHLSBroadcasting(context.Background(), type_, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.StartHLSBroadcasting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartHLSBroadcasting`: StartHLSBroadcastingResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.StartHLSBroadcasting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartHLSBroadcastingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StartHLSBroadcastingResponse**](StartHLSBroadcastingResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartRecording

> StartRecordingResponse StartRecording(ctx, type_, id).StartRecordingRequest(startRecordingRequest).Execute()

Start recording



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
	type_ := "type__example" // string | 
	id := "id_example" // string | 
	startRecordingRequest := *openapiclient.NewStartRecordingRequest() // StartRecordingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.StartRecording(context.Background(), type_, id).StartRecordingRequest(startRecordingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.StartRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartRecording`: StartRecordingResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.StartRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startRecordingRequest** | [**StartRecordingRequest**](StartRecordingRequest.md) |  | 

### Return type

[**StartRecordingResponse**](StartRecordingResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartTranscription

> StartTranscriptionResponse StartTranscription(ctx, type_, id).StartTranscriptionRequest(startTranscriptionRequest).Execute()

Start transcription



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
	type_ := "type__example" // string | 
	id := "id_example" // string | 
	startTranscriptionRequest := *openapiclient.NewStartTranscriptionRequest() // StartTranscriptionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.StartTranscription(context.Background(), type_, id).StartTranscriptionRequest(startTranscriptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.StartTranscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartTranscription`: StartTranscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.StartTranscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartTranscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startTranscriptionRequest** | [**StartTranscriptionRequest**](StartTranscriptionRequest.md) |  | 

### Return type

[**StartTranscriptionResponse**](StartTranscriptionResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopHLSBroadcasting

> StopHLSBroadcastingResponse StopHLSBroadcasting(ctx, type_, id).Execute()

Stop HLS broadcasting



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
	type_ := "type__example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.StopHLSBroadcasting(context.Background(), type_, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.StopHLSBroadcasting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopHLSBroadcasting`: StopHLSBroadcastingResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.StopHLSBroadcasting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopHLSBroadcastingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StopHLSBroadcastingResponse**](StopHLSBroadcastingResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopLive

> StopLiveResponse StopLive(ctx, type_, id).Execute()

Set call as not live



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
	type_ := "type__example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.StopLive(context.Background(), type_, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.StopLive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopLive`: StopLiveResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.StopLive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopLiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StopLiveResponse**](StopLiveResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopRecording

> StopRecordingResponse StopRecording(ctx, type_, id).Execute()

Stop recording



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
	type_ := "type__example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.StopRecording(context.Background(), type_, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.StopRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopRecording`: StopRecordingResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.StopRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StopRecordingResponse**](StopRecordingResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopTranscription

> StopTranscriptionResponse StopTranscription(ctx, type_, id).Execute()

Stop transcription



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
	type_ := "type__example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.StopTranscription(context.Background(), type_, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.StopTranscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopTranscription`: StopTranscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.StopTranscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopTranscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StopTranscriptionResponse**](StopTranscriptionResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnblockUser

> UnblockUserResponse UnblockUser(ctx, type_, id).UnblockUserRequest(unblockUserRequest).Execute()

Unblocks user on a call



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
	type_ := "type__example" // string | 
	id := "id_example" // string | 
	unblockUserRequest := *openapiclient.NewUnblockUserRequest("UserId_example") // UnblockUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UnblockUser(context.Background(), type_, id).UnblockUserRequest(unblockUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UnblockUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnblockUser`: UnblockUserResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UnblockUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnblockUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **unblockUserRequest** | [**UnblockUserRequest**](UnblockUserRequest.md) |  | 

### Return type

[**UnblockUserResponse**](UnblockUserResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCall

> UpdateCallResponse UpdateCall(ctx, type_, id).UpdateCallRequest(updateCallRequest).Execute()

Update Call



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
	type_ := "type__example" // string | 
	id := "id_example" // string | 
	updateCallRequest := *openapiclient.NewUpdateCallRequest() // UpdateCallRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateCall(context.Background(), type_, id).UpdateCallRequest(updateCallRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCall`: UpdateCallResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateCallRequest** | [**UpdateCallRequest**](UpdateCallRequest.md) |  | 

### Return type

[**UpdateCallResponse**](UpdateCallResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCallMembers

> UpdateCallMembersResponse UpdateCallMembers(ctx, type_, id).UpdateCallMembersRequest(updateCallMembersRequest).Execute()

Update Call Member



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
	type_ := "type__example" // string | 
	id := "id_example" // string | 
	updateCallMembersRequest := *openapiclient.NewUpdateCallMembersRequest() // UpdateCallMembersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateCallMembers(context.Background(), type_, id).UpdateCallMembersRequest(updateCallMembersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateCallMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCallMembers`: UpdateCallMembersResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateCallMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCallMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateCallMembersRequest** | [**UpdateCallMembersRequest**](UpdateCallMembersRequest.md) |  | 

### Return type

[**UpdateCallMembersResponse**](UpdateCallMembersResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserPermissions

> UpdateUserPermissionsResponse UpdateUserPermissions(ctx, type_, id).UpdateUserPermissionsRequest(updateUserPermissionsRequest).Execute()

Update user permissions



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
	type_ := "type__example" // string | 
	id := "id_example" // string | 
	updateUserPermissionsRequest := *openapiclient.NewUpdateUserPermissionsRequest("UserId_example") // UpdateUserPermissionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateUserPermissions(context.Background(), type_, id).UpdateUserPermissionsRequest(updateUserPermissionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateUserPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserPermissions`: UpdateUserPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateUserPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateUserPermissionsRequest** | [**UpdateUserPermissionsRequest**](UpdateUserPermissionsRequest.md) |  | 

### Return type

[**UpdateUserPermissionsResponse**](UpdateUserPermissionsResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VideoConnect

> VideoConnect(ctx).WSAuthMessageRequest(wSAuthMessageRequest).Execute()

Video Connect (WebSocket)



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
	wSAuthMessageRequest := *openapiclient.NewWSAuthMessageRequest("Token_example", *openapiclient.NewConnectUserDetailsRequest("Id_example")) // WSAuthMessageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.VideoConnect(context.Background()).WSAuthMessageRequest(wSAuthMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.VideoConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVideoConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wSAuthMessageRequest** | [**WSAuthMessageRequest**](WSAuthMessageRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VideoPin

> PinResponse VideoPin(ctx, type_, id).PinRequest(pinRequest).Execute()

Pin



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
	type_ := "type__example" // string | 
	id := "id_example" // string | 
	pinRequest := *openapiclient.NewPinRequest("SessionId_example", "UserId_example") // PinRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.VideoPin(context.Background(), type_, id).PinRequest(pinRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.VideoPin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VideoPin`: PinResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.VideoPin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVideoPinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pinRequest** | [**PinRequest**](PinRequest.md) |  | 

### Return type

[**PinResponse**](PinResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VideoUnpin

> UnpinResponse VideoUnpin(ctx, type_, id).UnpinRequest(unpinRequest).Execute()

Unpin



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
	type_ := "type__example" // string | 
	id := "id_example" // string | 
	unpinRequest := *openapiclient.NewUnpinRequest("SessionId_example", "UserId_example") // UnpinRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.VideoUnpin(context.Background(), type_, id).UnpinRequest(unpinRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.VideoUnpin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VideoUnpin`: UnpinResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.VideoUnpin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVideoUnpinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **unpinRequest** | [**UnpinRequest**](UnpinRequest.md) |  | 

### Return type

[**UnpinResponse**](UnpinResponse.md)

### Authorization

[stream-auth-type](../README.md#stream-auth-type), [api_key](../README.md#api_key), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

