/*
Stream API

Testing DefaultAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package getstream_video

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/cnztech/getstream_video"
)

func Test_getstream_video_DefaultAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DefaultAPIService BlockUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		resp, httpRes, err := apiClient.DefaultAPI.BlockUser(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService CreateDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DefaultAPI.CreateDevice(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService CreateGuest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DefaultAPI.CreateGuest(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService DeleteDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DefaultAPI.DeleteDevice(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService EndCall", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		resp, httpRes, err := apiClient.DefaultAPI.EndCall(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GetCall", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		resp, httpRes, err := apiClient.DefaultAPI.GetCall(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GetCallStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string
		var session string

		resp, httpRes, err := apiClient.DefaultAPI.GetCallStats(context.Background(), type_, id, session).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GetEdges", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DefaultAPI.GetEdges(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GetOrCreateCall", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		resp, httpRes, err := apiClient.DefaultAPI.GetOrCreateCall(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GoLive", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		resp, httpRes, err := apiClient.DefaultAPI.GoLive(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ListDevices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DefaultAPI.ListDevices(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ListRecordings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		resp, httpRes, err := apiClient.DefaultAPI.ListRecordings(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ListTranscriptions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		resp, httpRes, err := apiClient.DefaultAPI.ListTranscriptions(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService MuteUsers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		resp, httpRes, err := apiClient.DefaultAPI.MuteUsers(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService QueryCalls", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DefaultAPI.QueryCalls(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService QueryMembers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DefaultAPI.QueryMembers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService SendEvent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		resp, httpRes, err := apiClient.DefaultAPI.SendEvent(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService StartHLSBroadcasting", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		resp, httpRes, err := apiClient.DefaultAPI.StartHLSBroadcasting(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService StartRecording", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		resp, httpRes, err := apiClient.DefaultAPI.StartRecording(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService StartTranscription", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		resp, httpRes, err := apiClient.DefaultAPI.StartTranscription(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService StopHLSBroadcasting", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		resp, httpRes, err := apiClient.DefaultAPI.StopHLSBroadcasting(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService StopLive", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		resp, httpRes, err := apiClient.DefaultAPI.StopLive(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService StopRecording", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		resp, httpRes, err := apiClient.DefaultAPI.StopRecording(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService StopTranscription", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		resp, httpRes, err := apiClient.DefaultAPI.StopTranscription(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService UnblockUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		resp, httpRes, err := apiClient.DefaultAPI.UnblockUser(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService UpdateCall", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		resp, httpRes, err := apiClient.DefaultAPI.UpdateCall(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService UpdateCallMembers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		resp, httpRes, err := apiClient.DefaultAPI.UpdateCallMembers(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService UpdateUserPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		resp, httpRes, err := apiClient.DefaultAPI.UpdateUserPermissions(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService VideoConnect", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.DefaultAPI.VideoConnect(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService VideoPin", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		resp, httpRes, err := apiClient.DefaultAPI.VideoPin(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService VideoUnpin", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		resp, httpRes, err := apiClient.DefaultAPI.VideoUnpin(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
