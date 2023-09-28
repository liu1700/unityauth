/*
Player Authentication Admin API

# Introduction This is the Admin API specification for the Unity Authentication service that allows player authentication. To use this API, you must first enable it through the Unity Gaming Services dashboard.  For more information about how to set up Service Account Authentication, please read here: https://services.docs.unity.com/docs/service-account-auth  ## Rate Limits The API has rate limiting in place. Request are limited to 10 requests per second, and 500 requests per 30 minute period.  The API responds with a `429` HTTP status code if the rate limit is exceeded.  It will also respond with a `Retry-After` header to be used in conjunction with a client's retry logic. The value is the number of seconds until a request for the given player will be accepted. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unityauth

import (
	"encoding/json"
)

// checks if the PlayerAuthValidationErrorV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlayerAuthValidationErrorV1{}

// PlayerAuthValidationErrorV1 struct for PlayerAuthValidationErrorV1
type PlayerAuthValidationErrorV1 struct {
	Title *string `json:"title,omitempty"`
	Status *int32 `json:"status,omitempty"`
	Detail *string `json:"detail,omitempty"`
	Details []map[string]interface{} `json:"details,omitempty"`
}

// NewPlayerAuthValidationErrorV1 instantiates a new PlayerAuthValidationErrorV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlayerAuthValidationErrorV1() *PlayerAuthValidationErrorV1 {
	this := PlayerAuthValidationErrorV1{}
	return &this
}

// NewPlayerAuthValidationErrorV1WithDefaults instantiates a new PlayerAuthValidationErrorV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlayerAuthValidationErrorV1WithDefaults() *PlayerAuthValidationErrorV1 {
	this := PlayerAuthValidationErrorV1{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PlayerAuthValidationErrorV1) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerAuthValidationErrorV1) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PlayerAuthValidationErrorV1) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PlayerAuthValidationErrorV1) SetTitle(v string) {
	o.Title = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PlayerAuthValidationErrorV1) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerAuthValidationErrorV1) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PlayerAuthValidationErrorV1) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *PlayerAuthValidationErrorV1) SetStatus(v int32) {
	o.Status = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *PlayerAuthValidationErrorV1) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerAuthValidationErrorV1) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *PlayerAuthValidationErrorV1) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *PlayerAuthValidationErrorV1) SetDetail(v string) {
	o.Detail = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *PlayerAuthValidationErrorV1) GetDetails() []map[string]interface{} {
	if o == nil || IsNil(o.Details) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerAuthValidationErrorV1) GetDetailsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *PlayerAuthValidationErrorV1) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []map[string]interface{} and assigns it to the Details field.
func (o *PlayerAuthValidationErrorV1) SetDetails(v []map[string]interface{}) {
	o.Details = v
}

func (o PlayerAuthValidationErrorV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlayerAuthValidationErrorV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullablePlayerAuthValidationErrorV1 struct {
	value *PlayerAuthValidationErrorV1
	isSet bool
}

func (v NullablePlayerAuthValidationErrorV1) Get() *PlayerAuthValidationErrorV1 {
	return v.value
}

func (v *NullablePlayerAuthValidationErrorV1) Set(val *PlayerAuthValidationErrorV1) {
	v.value = val
	v.isSet = true
}

func (v NullablePlayerAuthValidationErrorV1) IsSet() bool {
	return v.isSet
}

func (v *NullablePlayerAuthValidationErrorV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlayerAuthValidationErrorV1(val *PlayerAuthValidationErrorV1) *NullablePlayerAuthValidationErrorV1 {
	return &NullablePlayerAuthValidationErrorV1{value: val, isSet: true}
}

func (v NullablePlayerAuthValidationErrorV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlayerAuthValidationErrorV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


