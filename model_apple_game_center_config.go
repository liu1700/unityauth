/*
Player Authentication API

# Introduction This is the API specification for the Unity Authentication service that allows player authentication.  ## Rate Limits The API has rate limiting in place. The endpoints are limited to 15 requests per second on a per-IP basis, and 300 requests over 30 minutes. The API responds with a `429` HTTP status code if the rate limit is exceeded. It will also respond with a `Retry-After` header to be used in conjunction with a client's retry logic. The value is the number of seconds until a request for the given player will be accepted. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unityauth

import (
	"encoding/json"
)

// checks if the AppleGameCenterConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppleGameCenterConfig{}

// AppleGameCenterConfig The request body for Apple Game Center authentication. This field is not applicable for any other Id provider.
type AppleGameCenterConfig struct {
	// String value of the Apple Game Center player's team player Id.
	TeamPlayerId *string `json:"teamPlayerId,omitempty"`
	// Integer value of the timestamp.
	Timestamp *int32 `json:"timestamp,omitempty"`
	// String value of the Apple Game Center public key url.
	PublicKeyUrl *string `json:"publicKeyUrl,omitempty"`
	// String value of the base64 encoded salt.
	Salt *string `json:"salt,omitempty"`
}

// NewAppleGameCenterConfig instantiates a new AppleGameCenterConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppleGameCenterConfig() *AppleGameCenterConfig {
	this := AppleGameCenterConfig{}
	return &this
}

// NewAppleGameCenterConfigWithDefaults instantiates a new AppleGameCenterConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppleGameCenterConfigWithDefaults() *AppleGameCenterConfig {
	this := AppleGameCenterConfig{}
	return &this
}

// GetTeamPlayerId returns the TeamPlayerId field value if set, zero value otherwise.
func (o *AppleGameCenterConfig) GetTeamPlayerId() string {
	if o == nil || IsNil(o.TeamPlayerId) {
		var ret string
		return ret
	}
	return *o.TeamPlayerId
}

// GetTeamPlayerIdOk returns a tuple with the TeamPlayerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleGameCenterConfig) GetTeamPlayerIdOk() (*string, bool) {
	if o == nil || IsNil(o.TeamPlayerId) {
		return nil, false
	}
	return o.TeamPlayerId, true
}

// HasTeamPlayerId returns a boolean if a field has been set.
func (o *AppleGameCenterConfig) HasTeamPlayerId() bool {
	if o != nil && !IsNil(o.TeamPlayerId) {
		return true
	}

	return false
}

// SetTeamPlayerId gets a reference to the given string and assigns it to the TeamPlayerId field.
func (o *AppleGameCenterConfig) SetTeamPlayerId(v string) {
	o.TeamPlayerId = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *AppleGameCenterConfig) GetTimestamp() int32 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleGameCenterConfig) GetTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *AppleGameCenterConfig) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int32 and assigns it to the Timestamp field.
func (o *AppleGameCenterConfig) SetTimestamp(v int32) {
	o.Timestamp = &v
}

// GetPublicKeyUrl returns the PublicKeyUrl field value if set, zero value otherwise.
func (o *AppleGameCenterConfig) GetPublicKeyUrl() string {
	if o == nil || IsNil(o.PublicKeyUrl) {
		var ret string
		return ret
	}
	return *o.PublicKeyUrl
}

// GetPublicKeyUrlOk returns a tuple with the PublicKeyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleGameCenterConfig) GetPublicKeyUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKeyUrl) {
		return nil, false
	}
	return o.PublicKeyUrl, true
}

// HasPublicKeyUrl returns a boolean if a field has been set.
func (o *AppleGameCenterConfig) HasPublicKeyUrl() bool {
	if o != nil && !IsNil(o.PublicKeyUrl) {
		return true
	}

	return false
}

// SetPublicKeyUrl gets a reference to the given string and assigns it to the PublicKeyUrl field.
func (o *AppleGameCenterConfig) SetPublicKeyUrl(v string) {
	o.PublicKeyUrl = &v
}

// GetSalt returns the Salt field value if set, zero value otherwise.
func (o *AppleGameCenterConfig) GetSalt() string {
	if o == nil || IsNil(o.Salt) {
		var ret string
		return ret
	}
	return *o.Salt
}

// GetSaltOk returns a tuple with the Salt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleGameCenterConfig) GetSaltOk() (*string, bool) {
	if o == nil || IsNil(o.Salt) {
		return nil, false
	}
	return o.Salt, true
}

// HasSalt returns a boolean if a field has been set.
func (o *AppleGameCenterConfig) HasSalt() bool {
	if o != nil && !IsNil(o.Salt) {
		return true
	}

	return false
}

// SetSalt gets a reference to the given string and assigns it to the Salt field.
func (o *AppleGameCenterConfig) SetSalt(v string) {
	o.Salt = &v
}

func (o AppleGameCenterConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppleGameCenterConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TeamPlayerId) {
		toSerialize["teamPlayerId"] = o.TeamPlayerId
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.PublicKeyUrl) {
		toSerialize["publicKeyUrl"] = o.PublicKeyUrl
	}
	if !IsNil(o.Salt) {
		toSerialize["salt"] = o.Salt
	}
	return toSerialize, nil
}

type NullableAppleGameCenterConfig struct {
	value *AppleGameCenterConfig
	isSet bool
}

func (v NullableAppleGameCenterConfig) Get() *AppleGameCenterConfig {
	return v.value
}

func (v *NullableAppleGameCenterConfig) Set(val *AppleGameCenterConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAppleGameCenterConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAppleGameCenterConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppleGameCenterConfig(val *AppleGameCenterConfig) *NullableAppleGameCenterConfig {
	return &NullableAppleGameCenterConfig{value: val, isSet: true}
}

func (v NullableAppleGameCenterConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppleGameCenterConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


