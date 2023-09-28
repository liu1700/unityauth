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

// checks if the CodeLinkInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CodeLinkInfoResponse{}

// CodeLinkInfoResponse Information on the requested code
type CodeLinkInfoResponse struct {
	// Human-readable string to identify the requester device.
	Identifier *string `json:"identifier,omitempty"`
	// The timestamp for when the code is no longer valid in unix time since epoch.
	Expiration *int32 `json:"expiration,omitempty"`
}

// NewCodeLinkInfoResponse instantiates a new CodeLinkInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodeLinkInfoResponse() *CodeLinkInfoResponse {
	this := CodeLinkInfoResponse{}
	return &this
}

// NewCodeLinkInfoResponseWithDefaults instantiates a new CodeLinkInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodeLinkInfoResponseWithDefaults() *CodeLinkInfoResponse {
	this := CodeLinkInfoResponse{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *CodeLinkInfoResponse) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodeLinkInfoResponse) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *CodeLinkInfoResponse) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *CodeLinkInfoResponse) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *CodeLinkInfoResponse) GetExpiration() int32 {
	if o == nil || IsNil(o.Expiration) {
		var ret int32
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodeLinkInfoResponse) GetExpirationOk() (*int32, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *CodeLinkInfoResponse) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given int32 and assigns it to the Expiration field.
func (o *CodeLinkInfoResponse) SetExpiration(v int32) {
	o.Expiration = &v
}

func (o CodeLinkInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CodeLinkInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}
	return toSerialize, nil
}

type NullableCodeLinkInfoResponse struct {
	value *CodeLinkInfoResponse
	isSet bool
}

func (v NullableCodeLinkInfoResponse) Get() *CodeLinkInfoResponse {
	return v.value
}

func (v *NullableCodeLinkInfoResponse) Set(val *CodeLinkInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCodeLinkInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCodeLinkInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodeLinkInfoResponse(val *CodeLinkInfoResponse) *NullableCodeLinkInfoResponse {
	return &NullableCodeLinkInfoResponse{value: val, isSet: true}
}

func (v NullableCodeLinkInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodeLinkInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

