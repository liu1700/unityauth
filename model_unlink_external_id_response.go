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

// checks if the UnlinkExternalIdResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnlinkExternalIdResponse{}

// UnlinkExternalIdResponse Response for an unlink request.
type UnlinkExternalIdResponse struct {
	// The number of seconds in which the Id token expires. By default it's 3600.
	ExpiresIn *int32 `json:"expiresIn,omitempty"`
	// The Id token of the authenticated player.
	IdToken *string `json:"idToken,omitempty"`
	// The session token of the authenticated player. This token can be used to sign-in the player again.
	SessionToken *string `json:"sessionToken,omitempty"`
	User *Player `json:"user,omitempty"`
	// The Id of the authenticated player. If a project is specified in the request, this field represents the Id of the project scoped player.
	UserId *string `json:"userId,omitempty"`
}

// NewUnlinkExternalIdResponse instantiates a new UnlinkExternalIdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnlinkExternalIdResponse() *UnlinkExternalIdResponse {
	this := UnlinkExternalIdResponse{}
	return &this
}

// NewUnlinkExternalIdResponseWithDefaults instantiates a new UnlinkExternalIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnlinkExternalIdResponseWithDefaults() *UnlinkExternalIdResponse {
	this := UnlinkExternalIdResponse{}
	return &this
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *UnlinkExternalIdResponse) GetExpiresIn() int32 {
	if o == nil || IsNil(o.ExpiresIn) {
		var ret int32
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnlinkExternalIdResponse) GetExpiresInOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpiresIn) {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *UnlinkExternalIdResponse) HasExpiresIn() bool {
	if o != nil && !IsNil(o.ExpiresIn) {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given int32 and assigns it to the ExpiresIn field.
func (o *UnlinkExternalIdResponse) SetExpiresIn(v int32) {
	o.ExpiresIn = &v
}

// GetIdToken returns the IdToken field value if set, zero value otherwise.
func (o *UnlinkExternalIdResponse) GetIdToken() string {
	if o == nil || IsNil(o.IdToken) {
		var ret string
		return ret
	}
	return *o.IdToken
}

// GetIdTokenOk returns a tuple with the IdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnlinkExternalIdResponse) GetIdTokenOk() (*string, bool) {
	if o == nil || IsNil(o.IdToken) {
		return nil, false
	}
	return o.IdToken, true
}

// HasIdToken returns a boolean if a field has been set.
func (o *UnlinkExternalIdResponse) HasIdToken() bool {
	if o != nil && !IsNil(o.IdToken) {
		return true
	}

	return false
}

// SetIdToken gets a reference to the given string and assigns it to the IdToken field.
func (o *UnlinkExternalIdResponse) SetIdToken(v string) {
	o.IdToken = &v
}

// GetSessionToken returns the SessionToken field value if set, zero value otherwise.
func (o *UnlinkExternalIdResponse) GetSessionToken() string {
	if o == nil || IsNil(o.SessionToken) {
		var ret string
		return ret
	}
	return *o.SessionToken
}

// GetSessionTokenOk returns a tuple with the SessionToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnlinkExternalIdResponse) GetSessionTokenOk() (*string, bool) {
	if o == nil || IsNil(o.SessionToken) {
		return nil, false
	}
	return o.SessionToken, true
}

// HasSessionToken returns a boolean if a field has been set.
func (o *UnlinkExternalIdResponse) HasSessionToken() bool {
	if o != nil && !IsNil(o.SessionToken) {
		return true
	}

	return false
}

// SetSessionToken gets a reference to the given string and assigns it to the SessionToken field.
func (o *UnlinkExternalIdResponse) SetSessionToken(v string) {
	o.SessionToken = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UnlinkExternalIdResponse) GetUser() Player {
	if o == nil || IsNil(o.User) {
		var ret Player
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnlinkExternalIdResponse) GetUserOk() (*Player, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UnlinkExternalIdResponse) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given Player and assigns it to the User field.
func (o *UnlinkExternalIdResponse) SetUser(v Player) {
	o.User = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnlinkExternalIdResponse) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnlinkExternalIdResponse) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnlinkExternalIdResponse) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *UnlinkExternalIdResponse) SetUserId(v string) {
	o.UserId = &v
}

func (o UnlinkExternalIdResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnlinkExternalIdResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpiresIn) {
		toSerialize["expiresIn"] = o.ExpiresIn
	}
	if !IsNil(o.IdToken) {
		toSerialize["idToken"] = o.IdToken
	}
	if !IsNil(o.SessionToken) {
		toSerialize["sessionToken"] = o.SessionToken
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUnlinkExternalIdResponse struct {
	value *UnlinkExternalIdResponse
	isSet bool
}

func (v NullableUnlinkExternalIdResponse) Get() *UnlinkExternalIdResponse {
	return v.value
}

func (v *NullableUnlinkExternalIdResponse) Set(val *UnlinkExternalIdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUnlinkExternalIdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUnlinkExternalIdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnlinkExternalIdResponse(val *UnlinkExternalIdResponse) *NullableUnlinkExternalIdResponse {
	return &NullableUnlinkExternalIdResponse{value: val, isSet: true}
}

func (v NullableUnlinkExternalIdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnlinkExternalIdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


