# LinkExternalIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceLink** | Pointer to **bool** | Force a link between the player specified in the UAS access token and the external Id. If a different UAS player is already linked to the external id, unlink that player from the external id before linking the request&#39;s player. | [optional] 
**Token** | Pointer to **string** | External token that can be verified to represent a player from the Id provider. This may be an Id token or an access token. | [optional] 
**OculusConfig** | Pointer to [**OculusConfig**](OculusConfig.md) |  | [optional] 
**AppleGameCenterConfig** | Pointer to [**AppleGameCenterConfig**](AppleGameCenterConfig.md) |  | [optional] 
**SteamConfig** | Pointer to [**SteamConfig**](SteamConfig.md) |  | [optional] 

## Methods

### NewLinkExternalIdRequest

`func NewLinkExternalIdRequest() *LinkExternalIdRequest`

NewLinkExternalIdRequest instantiates a new LinkExternalIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkExternalIdRequestWithDefaults

`func NewLinkExternalIdRequestWithDefaults() *LinkExternalIdRequest`

NewLinkExternalIdRequestWithDefaults instantiates a new LinkExternalIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceLink

`func (o *LinkExternalIdRequest) GetForceLink() bool`

GetForceLink returns the ForceLink field if non-nil, zero value otherwise.

### GetForceLinkOk

`func (o *LinkExternalIdRequest) GetForceLinkOk() (*bool, bool)`

GetForceLinkOk returns a tuple with the ForceLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceLink

`func (o *LinkExternalIdRequest) SetForceLink(v bool)`

SetForceLink sets ForceLink field to given value.

### HasForceLink

`func (o *LinkExternalIdRequest) HasForceLink() bool`

HasForceLink returns a boolean if a field has been set.

### GetToken

`func (o *LinkExternalIdRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LinkExternalIdRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LinkExternalIdRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LinkExternalIdRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetOculusConfig

`func (o *LinkExternalIdRequest) GetOculusConfig() OculusConfig`

GetOculusConfig returns the OculusConfig field if non-nil, zero value otherwise.

### GetOculusConfigOk

`func (o *LinkExternalIdRequest) GetOculusConfigOk() (*OculusConfig, bool)`

GetOculusConfigOk returns a tuple with the OculusConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOculusConfig

`func (o *LinkExternalIdRequest) SetOculusConfig(v OculusConfig)`

SetOculusConfig sets OculusConfig field to given value.

### HasOculusConfig

`func (o *LinkExternalIdRequest) HasOculusConfig() bool`

HasOculusConfig returns a boolean if a field has been set.

### GetAppleGameCenterConfig

`func (o *LinkExternalIdRequest) GetAppleGameCenterConfig() AppleGameCenterConfig`

GetAppleGameCenterConfig returns the AppleGameCenterConfig field if non-nil, zero value otherwise.

### GetAppleGameCenterConfigOk

`func (o *LinkExternalIdRequest) GetAppleGameCenterConfigOk() (*AppleGameCenterConfig, bool)`

GetAppleGameCenterConfigOk returns a tuple with the AppleGameCenterConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleGameCenterConfig

`func (o *LinkExternalIdRequest) SetAppleGameCenterConfig(v AppleGameCenterConfig)`

SetAppleGameCenterConfig sets AppleGameCenterConfig field to given value.

### HasAppleGameCenterConfig

`func (o *LinkExternalIdRequest) HasAppleGameCenterConfig() bool`

HasAppleGameCenterConfig returns a boolean if a field has been set.

### GetSteamConfig

`func (o *LinkExternalIdRequest) GetSteamConfig() SteamConfig`

GetSteamConfig returns the SteamConfig field if non-nil, zero value otherwise.

### GetSteamConfigOk

`func (o *LinkExternalIdRequest) GetSteamConfigOk() (*SteamConfig, bool)`

GetSteamConfigOk returns a tuple with the SteamConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamConfig

`func (o *LinkExternalIdRequest) SetSteamConfig(v SteamConfig)`

SetSteamConfig sets SteamConfig field to given value.

### HasSteamConfig

`func (o *LinkExternalIdRequest) HasSteamConfig() bool`

HasSteamConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


