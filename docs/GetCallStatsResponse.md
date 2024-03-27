# GetCallStatsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageJitter** | **float32** |  | 
**AverageLatency** | **float32** |  | 
**CallDurationSeconds** | **int64** |  | 
**CallTimeline** | Pointer to [**CallTimeline**](CallTimeline.md) |  | [optional] 
**Duration** | **string** | Duration of the request in human-readable format | 
**MaxParticipants** | **int32** |  | 
**ParticipantReport** | [**map[string]UserStats**](UserStats.md) |  | 
**PublishingParticipants** | **int32** |  | 
**QualityScore** | **int32** |  | 
**SfuCount** | **int32** |  | 
**Sfus** | [**[]SFULocationResponse**](SFULocationResponse.md) |  | 
**TotalFreezesDuration** | **float32** |  | 
**TotalQualityLimitationDuration** | **float32** |  | 

## Methods

### NewGetCallStatsResponse

`func NewGetCallStatsResponse(averageJitter float32, averageLatency float32, callDurationSeconds int64, duration string, maxParticipants int32, participantReport map[string]UserStats, publishingParticipants int32, qualityScore int32, sfuCount int32, sfus []SFULocationResponse, totalFreezesDuration float32, totalQualityLimitationDuration float32, ) *GetCallStatsResponse`

NewGetCallStatsResponse instantiates a new GetCallStatsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCallStatsResponseWithDefaults

`func NewGetCallStatsResponseWithDefaults() *GetCallStatsResponse`

NewGetCallStatsResponseWithDefaults instantiates a new GetCallStatsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageJitter

`func (o *GetCallStatsResponse) GetAverageJitter() float32`

GetAverageJitter returns the AverageJitter field if non-nil, zero value otherwise.

### GetAverageJitterOk

`func (o *GetCallStatsResponse) GetAverageJitterOk() (*float32, bool)`

GetAverageJitterOk returns a tuple with the AverageJitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageJitter

`func (o *GetCallStatsResponse) SetAverageJitter(v float32)`

SetAverageJitter sets AverageJitter field to given value.


### GetAverageLatency

`func (o *GetCallStatsResponse) GetAverageLatency() float32`

GetAverageLatency returns the AverageLatency field if non-nil, zero value otherwise.

### GetAverageLatencyOk

`func (o *GetCallStatsResponse) GetAverageLatencyOk() (*float32, bool)`

GetAverageLatencyOk returns a tuple with the AverageLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageLatency

`func (o *GetCallStatsResponse) SetAverageLatency(v float32)`

SetAverageLatency sets AverageLatency field to given value.


### GetCallDurationSeconds

`func (o *GetCallStatsResponse) GetCallDurationSeconds() int64`

GetCallDurationSeconds returns the CallDurationSeconds field if non-nil, zero value otherwise.

### GetCallDurationSecondsOk

`func (o *GetCallStatsResponse) GetCallDurationSecondsOk() (*int64, bool)`

GetCallDurationSecondsOk returns a tuple with the CallDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallDurationSeconds

`func (o *GetCallStatsResponse) SetCallDurationSeconds(v int64)`

SetCallDurationSeconds sets CallDurationSeconds field to given value.


### GetCallTimeline

`func (o *GetCallStatsResponse) GetCallTimeline() CallTimeline`

GetCallTimeline returns the CallTimeline field if non-nil, zero value otherwise.

### GetCallTimelineOk

`func (o *GetCallStatsResponse) GetCallTimelineOk() (*CallTimeline, bool)`

GetCallTimelineOk returns a tuple with the CallTimeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallTimeline

`func (o *GetCallStatsResponse) SetCallTimeline(v CallTimeline)`

SetCallTimeline sets CallTimeline field to given value.

### HasCallTimeline

`func (o *GetCallStatsResponse) HasCallTimeline() bool`

HasCallTimeline returns a boolean if a field has been set.

### GetDuration

`func (o *GetCallStatsResponse) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetCallStatsResponse) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetCallStatsResponse) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetMaxParticipants

`func (o *GetCallStatsResponse) GetMaxParticipants() int32`

GetMaxParticipants returns the MaxParticipants field if non-nil, zero value otherwise.

### GetMaxParticipantsOk

`func (o *GetCallStatsResponse) GetMaxParticipantsOk() (*int32, bool)`

GetMaxParticipantsOk returns a tuple with the MaxParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParticipants

`func (o *GetCallStatsResponse) SetMaxParticipants(v int32)`

SetMaxParticipants sets MaxParticipants field to given value.


### GetParticipantReport

`func (o *GetCallStatsResponse) GetParticipantReport() map[string]UserStats`

GetParticipantReport returns the ParticipantReport field if non-nil, zero value otherwise.

### GetParticipantReportOk

`func (o *GetCallStatsResponse) GetParticipantReportOk() (*map[string]UserStats, bool)`

GetParticipantReportOk returns a tuple with the ParticipantReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantReport

`func (o *GetCallStatsResponse) SetParticipantReport(v map[string]UserStats)`

SetParticipantReport sets ParticipantReport field to given value.


### GetPublishingParticipants

`func (o *GetCallStatsResponse) GetPublishingParticipants() int32`

GetPublishingParticipants returns the PublishingParticipants field if non-nil, zero value otherwise.

### GetPublishingParticipantsOk

`func (o *GetCallStatsResponse) GetPublishingParticipantsOk() (*int32, bool)`

GetPublishingParticipantsOk returns a tuple with the PublishingParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishingParticipants

`func (o *GetCallStatsResponse) SetPublishingParticipants(v int32)`

SetPublishingParticipants sets PublishingParticipants field to given value.


### GetQualityScore

`func (o *GetCallStatsResponse) GetQualityScore() int32`

GetQualityScore returns the QualityScore field if non-nil, zero value otherwise.

### GetQualityScoreOk

`func (o *GetCallStatsResponse) GetQualityScoreOk() (*int32, bool)`

GetQualityScoreOk returns a tuple with the QualityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityScore

`func (o *GetCallStatsResponse) SetQualityScore(v int32)`

SetQualityScore sets QualityScore field to given value.


### GetSfuCount

`func (o *GetCallStatsResponse) GetSfuCount() int32`

GetSfuCount returns the SfuCount field if non-nil, zero value otherwise.

### GetSfuCountOk

`func (o *GetCallStatsResponse) GetSfuCountOk() (*int32, bool)`

GetSfuCountOk returns a tuple with the SfuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfuCount

`func (o *GetCallStatsResponse) SetSfuCount(v int32)`

SetSfuCount sets SfuCount field to given value.


### GetSfus

`func (o *GetCallStatsResponse) GetSfus() []SFULocationResponse`

GetSfus returns the Sfus field if non-nil, zero value otherwise.

### GetSfusOk

`func (o *GetCallStatsResponse) GetSfusOk() (*[]SFULocationResponse, bool)`

GetSfusOk returns a tuple with the Sfus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfus

`func (o *GetCallStatsResponse) SetSfus(v []SFULocationResponse)`

SetSfus sets Sfus field to given value.


### GetTotalFreezesDuration

`func (o *GetCallStatsResponse) GetTotalFreezesDuration() float32`

GetTotalFreezesDuration returns the TotalFreezesDuration field if non-nil, zero value otherwise.

### GetTotalFreezesDurationOk

`func (o *GetCallStatsResponse) GetTotalFreezesDurationOk() (*float32, bool)`

GetTotalFreezesDurationOk returns a tuple with the TotalFreezesDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFreezesDuration

`func (o *GetCallStatsResponse) SetTotalFreezesDuration(v float32)`

SetTotalFreezesDuration sets TotalFreezesDuration field to given value.


### GetTotalQualityLimitationDuration

`func (o *GetCallStatsResponse) GetTotalQualityLimitationDuration() float32`

GetTotalQualityLimitationDuration returns the TotalQualityLimitationDuration field if non-nil, zero value otherwise.

### GetTotalQualityLimitationDurationOk

`func (o *GetCallStatsResponse) GetTotalQualityLimitationDurationOk() (*float32, bool)`

GetTotalQualityLimitationDurationOk returns a tuple with the TotalQualityLimitationDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalQualityLimitationDuration

`func (o *GetCallStatsResponse) SetTotalQualityLimitationDuration(v float32)`

SetTotalQualityLimitationDuration sets TotalQualityLimitationDuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


