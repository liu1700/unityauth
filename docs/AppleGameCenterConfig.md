# AppleGameCenterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TeamPlayerId** | Pointer to **string** | String value of the Apple Game Center player&#39;s team player Id. | [optional] 
**Timestamp** | Pointer to **int32** | Integer value of the timestamp. | [optional] 
**PublicKeyUrl** | Pointer to **string** | String value of the Apple Game Center public key url. | [optional] 
**Salt** | Pointer to **string** | String value of the base64 encoded salt. | [optional] 

## Methods

### NewAppleGameCenterConfig

`func NewAppleGameCenterConfig() *AppleGameCenterConfig`

NewAppleGameCenterConfig instantiates a new AppleGameCenterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppleGameCenterConfigWithDefaults

`func NewAppleGameCenterConfigWithDefaults() *AppleGameCenterConfig`

NewAppleGameCenterConfigWithDefaults instantiates a new AppleGameCenterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeamPlayerId

`func (o *AppleGameCenterConfig) GetTeamPlayerId() string`

GetTeamPlayerId returns the TeamPlayerId field if non-nil, zero value otherwise.

### GetTeamPlayerIdOk

`func (o *AppleGameCenterConfig) GetTeamPlayerIdOk() (*string, bool)`

GetTeamPlayerIdOk returns a tuple with the TeamPlayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamPlayerId

`func (o *AppleGameCenterConfig) SetTeamPlayerId(v string)`

SetTeamPlayerId sets TeamPlayerId field to given value.

### HasTeamPlayerId

`func (o *AppleGameCenterConfig) HasTeamPlayerId() bool`

HasTeamPlayerId returns a boolean if a field has been set.

### GetTimestamp

`func (o *AppleGameCenterConfig) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AppleGameCenterConfig) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AppleGameCenterConfig) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AppleGameCenterConfig) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetPublicKeyUrl

`func (o *AppleGameCenterConfig) GetPublicKeyUrl() string`

GetPublicKeyUrl returns the PublicKeyUrl field if non-nil, zero value otherwise.

### GetPublicKeyUrlOk

`func (o *AppleGameCenterConfig) GetPublicKeyUrlOk() (*string, bool)`

GetPublicKeyUrlOk returns a tuple with the PublicKeyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyUrl

`func (o *AppleGameCenterConfig) SetPublicKeyUrl(v string)`

SetPublicKeyUrl sets PublicKeyUrl field to given value.

### HasPublicKeyUrl

`func (o *AppleGameCenterConfig) HasPublicKeyUrl() bool`

HasPublicKeyUrl returns a boolean if a field has been set.

### GetSalt

`func (o *AppleGameCenterConfig) GetSalt() string`

GetSalt returns the Salt field if non-nil, zero value otherwise.

### GetSaltOk

`func (o *AppleGameCenterConfig) GetSaltOk() (*string, bool)`

GetSaltOk returns a tuple with the Salt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalt

`func (o *AppleGameCenterConfig) SetSalt(v string)`

SetSalt sets Salt field to given value.

### HasSalt

`func (o *AppleGameCenterConfig) HasSalt() bool`

HasSalt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


