# CallSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audio** | [**AudioSettings**](AudioSettings.md) |  | 
**Backstage** | [**BackstageSettings**](BackstageSettings.md) |  | 
**Broadcasting** | [**BroadcastSettingsResponse**](BroadcastSettingsResponse.md) |  | 
**Geofencing** | [**GeofenceSettings**](GeofenceSettings.md) |  | 
**Recording** | [**RecordSettingsResponse**](RecordSettingsResponse.md) |  | 
**Ring** | [**RingSettings**](RingSettings.md) |  | 
**Screensharing** | [**ScreensharingSettings**](ScreensharingSettings.md) |  | 
**Thumbnails** | [**ThumbnailsSettings**](ThumbnailsSettings.md) |  | 
**Transcription** | [**TranscriptionSettings**](TranscriptionSettings.md) |  | 
**Video** | [**VideoSettings**](VideoSettings.md) |  | 

## Methods

### NewCallSettingsResponse

`func NewCallSettingsResponse(audio AudioSettings, backstage BackstageSettings, broadcasting BroadcastSettingsResponse, geofencing GeofenceSettings, recording RecordSettingsResponse, ring RingSettings, screensharing ScreensharingSettings, thumbnails ThumbnailsSettings, transcription TranscriptionSettings, video VideoSettings, ) *CallSettingsResponse`

NewCallSettingsResponse instantiates a new CallSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallSettingsResponseWithDefaults

`func NewCallSettingsResponseWithDefaults() *CallSettingsResponse`

NewCallSettingsResponseWithDefaults instantiates a new CallSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudio

`func (o *CallSettingsResponse) GetAudio() AudioSettings`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *CallSettingsResponse) GetAudioOk() (*AudioSettings, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *CallSettingsResponse) SetAudio(v AudioSettings)`

SetAudio sets Audio field to given value.


### GetBackstage

`func (o *CallSettingsResponse) GetBackstage() BackstageSettings`

GetBackstage returns the Backstage field if non-nil, zero value otherwise.

### GetBackstageOk

`func (o *CallSettingsResponse) GetBackstageOk() (*BackstageSettings, bool)`

GetBackstageOk returns a tuple with the Backstage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackstage

`func (o *CallSettingsResponse) SetBackstage(v BackstageSettings)`

SetBackstage sets Backstage field to given value.


### GetBroadcasting

`func (o *CallSettingsResponse) GetBroadcasting() BroadcastSettingsResponse`

GetBroadcasting returns the Broadcasting field if non-nil, zero value otherwise.

### GetBroadcastingOk

`func (o *CallSettingsResponse) GetBroadcastingOk() (*BroadcastSettingsResponse, bool)`

GetBroadcastingOk returns a tuple with the Broadcasting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcasting

`func (o *CallSettingsResponse) SetBroadcasting(v BroadcastSettingsResponse)`

SetBroadcasting sets Broadcasting field to given value.


### GetGeofencing

`func (o *CallSettingsResponse) GetGeofencing() GeofenceSettings`

GetGeofencing returns the Geofencing field if non-nil, zero value otherwise.

### GetGeofencingOk

`func (o *CallSettingsResponse) GetGeofencingOk() (*GeofenceSettings, bool)`

GetGeofencingOk returns a tuple with the Geofencing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofencing

`func (o *CallSettingsResponse) SetGeofencing(v GeofenceSettings)`

SetGeofencing sets Geofencing field to given value.


### GetRecording

`func (o *CallSettingsResponse) GetRecording() RecordSettingsResponse`

GetRecording returns the Recording field if non-nil, zero value otherwise.

### GetRecordingOk

`func (o *CallSettingsResponse) GetRecordingOk() (*RecordSettingsResponse, bool)`

GetRecordingOk returns a tuple with the Recording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecording

`func (o *CallSettingsResponse) SetRecording(v RecordSettingsResponse)`

SetRecording sets Recording field to given value.


### GetRing

`func (o *CallSettingsResponse) GetRing() RingSettings`

GetRing returns the Ring field if non-nil, zero value otherwise.

### GetRingOk

`func (o *CallSettingsResponse) GetRingOk() (*RingSettings, bool)`

GetRingOk returns a tuple with the Ring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRing

`func (o *CallSettingsResponse) SetRing(v RingSettings)`

SetRing sets Ring field to given value.


### GetScreensharing

`func (o *CallSettingsResponse) GetScreensharing() ScreensharingSettings`

GetScreensharing returns the Screensharing field if non-nil, zero value otherwise.

### GetScreensharingOk

`func (o *CallSettingsResponse) GetScreensharingOk() (*ScreensharingSettings, bool)`

GetScreensharingOk returns a tuple with the Screensharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreensharing

`func (o *CallSettingsResponse) SetScreensharing(v ScreensharingSettings)`

SetScreensharing sets Screensharing field to given value.


### GetThumbnails

`func (o *CallSettingsResponse) GetThumbnails() ThumbnailsSettings`

GetThumbnails returns the Thumbnails field if non-nil, zero value otherwise.

### GetThumbnailsOk

`func (o *CallSettingsResponse) GetThumbnailsOk() (*ThumbnailsSettings, bool)`

GetThumbnailsOk returns a tuple with the Thumbnails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnails

`func (o *CallSettingsResponse) SetThumbnails(v ThumbnailsSettings)`

SetThumbnails sets Thumbnails field to given value.


### GetTranscription

`func (o *CallSettingsResponse) GetTranscription() TranscriptionSettings`

GetTranscription returns the Transcription field if non-nil, zero value otherwise.

### GetTranscriptionOk

`func (o *CallSettingsResponse) GetTranscriptionOk() (*TranscriptionSettings, bool)`

GetTranscriptionOk returns a tuple with the Transcription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscription

`func (o *CallSettingsResponse) SetTranscription(v TranscriptionSettings)`

SetTranscription sets Transcription field to given value.


### GetVideo

`func (o *CallSettingsResponse) GetVideo() VideoSettings`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *CallSettingsResponse) GetVideoOk() (*VideoSettings, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *CallSettingsResponse) SetVideo(v VideoSettings)`

SetVideo sets Video field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


