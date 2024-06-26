/*
Stream API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v100.14.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package getstream_video

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the CallTypeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallTypeResponse{}

// CallTypeResponse struct for CallTypeResponse
type CallTypeResponse struct {
	CreatedAt time.Time `json:"created_at"`
	ExternalStorage *string `json:"external_storage,omitempty"`
	Grants map[string][]string `json:"grants"`
	Name string `json:"name"`
	NotificationSettings NotificationSettings `json:"notification_settings"`
	Settings CallSettingsResponse `json:"settings"`
	UpdatedAt time.Time `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _CallTypeResponse CallTypeResponse

// NewCallTypeResponse instantiates a new CallTypeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallTypeResponse(createdAt time.Time, grants map[string][]string, name string, notificationSettings NotificationSettings, settings CallSettingsResponse, updatedAt time.Time) *CallTypeResponse {
	this := CallTypeResponse{}
	this.CreatedAt = createdAt
	this.Grants = grants
	this.Name = name
	this.NotificationSettings = notificationSettings
	this.Settings = settings
	this.UpdatedAt = updatedAt
	return &this
}

// NewCallTypeResponseWithDefaults instantiates a new CallTypeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallTypeResponseWithDefaults() *CallTypeResponse {
	this := CallTypeResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *CallTypeResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CallTypeResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CallTypeResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetExternalStorage returns the ExternalStorage field value if set, zero value otherwise.
func (o *CallTypeResponse) GetExternalStorage() string {
	if o == nil || IsNil(o.ExternalStorage) {
		var ret string
		return ret
	}
	return *o.ExternalStorage
}

// GetExternalStorageOk returns a tuple with the ExternalStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTypeResponse) GetExternalStorageOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalStorage) {
		return nil, false
	}
	return o.ExternalStorage, true
}

// HasExternalStorage returns a boolean if a field has been set.
func (o *CallTypeResponse) HasExternalStorage() bool {
	if o != nil && !IsNil(o.ExternalStorage) {
		return true
	}

	return false
}

// SetExternalStorage gets a reference to the given string and assigns it to the ExternalStorage field.
func (o *CallTypeResponse) SetExternalStorage(v string) {
	o.ExternalStorage = &v
}

// GetGrants returns the Grants field value
func (o *CallTypeResponse) GetGrants() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}

	return o.Grants
}

// GetGrantsOk returns a tuple with the Grants field value
// and a boolean to check if the value has been set.
func (o *CallTypeResponse) GetGrantsOk() (*map[string][]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Grants, true
}

// SetGrants sets field value
func (o *CallTypeResponse) SetGrants(v map[string][]string) {
	o.Grants = v
}

// GetName returns the Name field value
func (o *CallTypeResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CallTypeResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CallTypeResponse) SetName(v string) {
	o.Name = v
}

// GetNotificationSettings returns the NotificationSettings field value
func (o *CallTypeResponse) GetNotificationSettings() NotificationSettings {
	if o == nil {
		var ret NotificationSettings
		return ret
	}

	return o.NotificationSettings
}

// GetNotificationSettingsOk returns a tuple with the NotificationSettings field value
// and a boolean to check if the value has been set.
func (o *CallTypeResponse) GetNotificationSettingsOk() (*NotificationSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationSettings, true
}

// SetNotificationSettings sets field value
func (o *CallTypeResponse) SetNotificationSettings(v NotificationSettings) {
	o.NotificationSettings = v
}

// GetSettings returns the Settings field value
func (o *CallTypeResponse) GetSettings() CallSettingsResponse {
	if o == nil {
		var ret CallSettingsResponse
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *CallTypeResponse) GetSettingsOk() (*CallSettingsResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *CallTypeResponse) SetSettings(v CallSettingsResponse) {
	o.Settings = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CallTypeResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CallTypeResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CallTypeResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o CallTypeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallTypeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.ExternalStorage) {
		toSerialize["external_storage"] = o.ExternalStorage
	}
	toSerialize["grants"] = o.Grants
	toSerialize["name"] = o.Name
	toSerialize["notification_settings"] = o.NotificationSettings
	toSerialize["settings"] = o.Settings
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CallTypeResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created_at",
		"grants",
		"name",
		"notification_settings",
		"settings",
		"updated_at",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCallTypeResponse := _CallTypeResponse{}

	err = json.Unmarshal(data, &varCallTypeResponse)

	if err != nil {
		return err
	}

	*o = CallTypeResponse(varCallTypeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "external_storage")
		delete(additionalProperties, "grants")
		delete(additionalProperties, "name")
		delete(additionalProperties, "notification_settings")
		delete(additionalProperties, "settings")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCallTypeResponse struct {
	value *CallTypeResponse
	isSet bool
}

func (v NullableCallTypeResponse) Get() *CallTypeResponse {
	return v.value
}

func (v *NullableCallTypeResponse) Set(val *CallTypeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCallTypeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCallTypeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallTypeResponse(val *CallTypeResponse) *NullableCallTypeResponse {
	return &NullableCallTypeResponse{value: val, isSet: true}
}

func (v NullableCallTypeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallTypeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


