# CallSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audio** | Pointer to [**AudioSettingsRequest**](AudioSettingsRequest.md) |  | [optional] 
**Backstage** | Pointer to [**BackstageSettingsRequest**](BackstageSettingsRequest.md) |  | [optional] 
**Broadcasting** | Pointer to [**BroadcastSettingsRequest**](BroadcastSettingsRequest.md) |  | [optional] 
**Geofencing** | Pointer to [**GeofenceSettingsRequest**](GeofenceSettingsRequest.md) |  | [optional] 
**Recording** | Pointer to [**RecordSettingsRequest**](RecordSettingsRequest.md) |  | [optional] 
**Ring** | Pointer to [**RingSettingsRequest**](RingSettingsRequest.md) |  | [optional] 
**Screensharing** | Pointer to [**ScreensharingSettingsRequest**](ScreensharingSettingsRequest.md) |  | [optional] 
**Thumbnails** | Pointer to [**ThumbnailsSettingsRequest**](ThumbnailsSettingsRequest.md) |  | [optional] 
**Transcription** | Pointer to [**TranscriptionSettingsRequest**](TranscriptionSettingsRequest.md) |  | [optional] 
**Video** | Pointer to [**VideoSettingsRequest**](VideoSettingsRequest.md) |  | [optional] 

## Methods

### NewCallSettingsRequest

`func NewCallSettingsRequest() *CallSettingsRequest`

NewCallSettingsRequest instantiates a new CallSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallSettingsRequestWithDefaults

`func NewCallSettingsRequestWithDefaults() *CallSettingsRequest`

NewCallSettingsRequestWithDefaults instantiates a new CallSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudio

`func (o *CallSettingsRequest) GetAudio() AudioSettingsRequest`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *CallSettingsRequest) GetAudioOk() (*AudioSettingsRequest, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *CallSettingsRequest) SetAudio(v AudioSettingsRequest)`

SetAudio sets Audio field to given value.

### HasAudio

`func (o *CallSettingsRequest) HasAudio() bool`

HasAudio returns a boolean if a field has been set.

### GetBackstage

`func (o *CallSettingsRequest) GetBackstage() BackstageSettingsRequest`

GetBackstage returns the Backstage field if non-nil, zero value otherwise.

### GetBackstageOk

`func (o *CallSettingsRequest) GetBackstageOk() (*BackstageSettingsRequest, bool)`

GetBackstageOk returns a tuple with the Backstage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackstage

`func (o *CallSettingsRequest) SetBackstage(v BackstageSettingsRequest)`

SetBackstage sets Backstage field to given value.

### HasBackstage

`func (o *CallSettingsRequest) HasBackstage() bool`

HasBackstage returns a boolean if a field has been set.

### GetBroadcasting

`func (o *CallSettingsRequest) GetBroadcasting() BroadcastSettingsRequest`

GetBroadcasting returns the Broadcasting field if non-nil, zero value otherwise.

### GetBroadcastingOk

`func (o *CallSettingsRequest) GetBroadcastingOk() (*BroadcastSettingsRequest, bool)`

GetBroadcastingOk returns a tuple with the Broadcasting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcasting

`func (o *CallSettingsRequest) SetBroadcasting(v BroadcastSettingsRequest)`

SetBroadcasting sets Broadcasting field to given value.

### HasBroadcasting

`func (o *CallSettingsRequest) HasBroadcasting() bool`

HasBroadcasting returns a boolean if a field has been set.

### GetGeofencing

`func (o *CallSettingsRequest) GetGeofencing() GeofenceSettingsRequest`

GetGeofencing returns the Geofencing field if non-nil, zero value otherwise.

### GetGeofencingOk

`func (o *CallSettingsRequest) GetGeofencingOk() (*GeofenceSettingsRequest, bool)`

GetGeofencingOk returns a tuple with the Geofencing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofencing

`func (o *CallSettingsRequest) SetGeofencing(v GeofenceSettingsRequest)`

SetGeofencing sets Geofencing field to given value.

### HasGeofencing

`func (o *CallSettingsRequest) HasGeofencing() bool`

HasGeofencing returns a boolean if a field has been set.

### GetRecording

`func (o *CallSettingsRequest) GetRecording() RecordSettingsRequest`

GetRecording returns the Recording field if non-nil, zero value otherwise.

### GetRecordingOk

`func (o *CallSettingsRequest) GetRecordingOk() (*RecordSettingsRequest, bool)`

GetRecordingOk returns a tuple with the Recording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecording

`func (o *CallSettingsRequest) SetRecording(v RecordSettingsRequest)`

SetRecording sets Recording field to given value.

### HasRecording

`func (o *CallSettingsRequest) HasRecording() bool`

HasRecording returns a boolean if a field has been set.

### GetRing

`func (o *CallSettingsRequest) GetRing() RingSettingsRequest`

GetRing returns the Ring field if non-nil, zero value otherwise.

### GetRingOk

`func (o *CallSettingsRequest) GetRingOk() (*RingSettingsRequest, bool)`

GetRingOk returns a tuple with the Ring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRing

`func (o *CallSettingsRequest) SetRing(v RingSettingsRequest)`

SetRing sets Ring field to given value.

### HasRing

`func (o *CallSettingsRequest) HasRing() bool`

HasRing returns a boolean if a field has been set.

### GetScreensharing

`func (o *CallSettingsRequest) GetScreensharing() ScreensharingSettingsRequest`

GetScreensharing returns the Screensharing field if non-nil, zero value otherwise.

### GetScreensharingOk

`func (o *CallSettingsRequest) GetScreensharingOk() (*ScreensharingSettingsRequest, bool)`

GetScreensharingOk returns a tuple with the Screensharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreensharing

`func (o *CallSettingsRequest) SetScreensharing(v ScreensharingSettingsRequest)`

SetScreensharing sets Screensharing field to given value.

### HasScreensharing

`func (o *CallSettingsRequest) HasScreensharing() bool`

HasScreensharing returns a boolean if a field has been set.

### GetThumbnails

`func (o *CallSettingsRequest) GetThumbnails() ThumbnailsSettingsRequest`

GetThumbnails returns the Thumbnails field if non-nil, zero value otherwise.

### GetThumbnailsOk

`func (o *CallSettingsRequest) GetThumbnailsOk() (*ThumbnailsSettingsRequest, bool)`

GetThumbnailsOk returns a tuple with the Thumbnails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnails

`func (o *CallSettingsRequest) SetThumbnails(v ThumbnailsSettingsRequest)`

SetThumbnails sets Thumbnails field to given value.

### HasThumbnails

`func (o *CallSettingsRequest) HasThumbnails() bool`

HasThumbnails returns a boolean if a field has been set.

### GetTranscription

`func (o *CallSettingsRequest) GetTranscription() TranscriptionSettingsRequest`

GetTranscription returns the Transcription field if non-nil, zero value otherwise.

### GetTranscriptionOk

`func (o *CallSettingsRequest) GetTranscriptionOk() (*TranscriptionSettingsRequest, bool)`

GetTranscriptionOk returns a tuple with the Transcription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscription

`func (o *CallSettingsRequest) SetTranscription(v TranscriptionSettingsRequest)`

SetTranscription sets Transcription field to given value.

### HasTranscription

`func (o *CallSettingsRequest) HasTranscription() bool`

HasTranscription returns a boolean if a field has been set.

### GetVideo

`func (o *CallSettingsRequest) GetVideo() VideoSettingsRequest`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *CallSettingsRequest) GetVideoOk() (*VideoSettingsRequest, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *CallSettingsRequest) SetVideo(v VideoSettingsRequest)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *CallSettingsRequest) HasVideo() bool`

HasVideo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


