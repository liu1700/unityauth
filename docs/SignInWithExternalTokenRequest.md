# SignInWithExternalTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nonce** | Pointer to **string** | String value used to associate a client session with an Id token, and to mitigate replay attacks. If this field is provided, the nonce claim in response Id token will have a matching value. | [optional] 
**SignInOnly** | Pointer to **bool** | Whether the API should only attempt to sign-in and do not create a new player if the player does not exist. | [optional] 
**Token** | Pointer to **string** | External token that can be verified to represent a player from the id provider. This may be an id token or an access token. | [optional] 
**OculusConfig** | Pointer to [**OculusConfig**](OculusConfig.md) |  | [optional] 
**AppleGameCenterConfig** | Pointer to [**AppleGameCenterConfig**](AppleGameCenterConfig.md) |  | [optional] 
**SteamConfig** | Pointer to [**SteamConfig**](SteamConfig.md) |  | [optional] 

## Methods

### NewSignInWithExternalTokenRequest

`func NewSignInWithExternalTokenRequest() *SignInWithExternalTokenRequest`

NewSignInWithExternalTokenRequest instantiates a new SignInWithExternalTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignInWithExternalTokenRequestWithDefaults

`func NewSignInWithExternalTokenRequestWithDefaults() *SignInWithExternalTokenRequest`

NewSignInWithExternalTokenRequestWithDefaults instantiates a new SignInWithExternalTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNonce

`func (o *SignInWithExternalTokenRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *SignInWithExternalTokenRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *SignInWithExternalTokenRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *SignInWithExternalTokenRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetSignInOnly

`func (o *SignInWithExternalTokenRequest) GetSignInOnly() bool`

GetSignInOnly returns the SignInOnly field if non-nil, zero value otherwise.

### GetSignInOnlyOk

`func (o *SignInWithExternalTokenRequest) GetSignInOnlyOk() (*bool, bool)`

GetSignInOnlyOk returns a tuple with the SignInOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInOnly

`func (o *SignInWithExternalTokenRequest) SetSignInOnly(v bool)`

SetSignInOnly sets SignInOnly field to given value.

### HasSignInOnly

`func (o *SignInWithExternalTokenRequest) HasSignInOnly() bool`

HasSignInOnly returns a boolean if a field has been set.

### GetToken

`func (o *SignInWithExternalTokenRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SignInWithExternalTokenRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SignInWithExternalTokenRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SignInWithExternalTokenRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetOculusConfig

`func (o *SignInWithExternalTokenRequest) GetOculusConfig() OculusConfig`

GetOculusConfig returns the OculusConfig field if non-nil, zero value otherwise.

### GetOculusConfigOk

`func (o *SignInWithExternalTokenRequest) GetOculusConfigOk() (*OculusConfig, bool)`

GetOculusConfigOk returns a tuple with the OculusConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOculusConfig

`func (o *SignInWithExternalTokenRequest) SetOculusConfig(v OculusConfig)`

SetOculusConfig sets OculusConfig field to given value.

### HasOculusConfig

`func (o *SignInWithExternalTokenRequest) HasOculusConfig() bool`

HasOculusConfig returns a boolean if a field has been set.

### GetAppleGameCenterConfig

`func (o *SignInWithExternalTokenRequest) GetAppleGameCenterConfig() AppleGameCenterConfig`

GetAppleGameCenterConfig returns the AppleGameCenterConfig field if non-nil, zero value otherwise.

### GetAppleGameCenterConfigOk

`func (o *SignInWithExternalTokenRequest) GetAppleGameCenterConfigOk() (*AppleGameCenterConfig, bool)`

GetAppleGameCenterConfigOk returns a tuple with the AppleGameCenterConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleGameCenterConfig

`func (o *SignInWithExternalTokenRequest) SetAppleGameCenterConfig(v AppleGameCenterConfig)`

SetAppleGameCenterConfig sets AppleGameCenterConfig field to given value.

### HasAppleGameCenterConfig

`func (o *SignInWithExternalTokenRequest) HasAppleGameCenterConfig() bool`

HasAppleGameCenterConfig returns a boolean if a field has been set.

### GetSteamConfig

`func (o *SignInWithExternalTokenRequest) GetSteamConfig() SteamConfig`

GetSteamConfig returns the SteamConfig field if non-nil, zero value otherwise.

### GetSteamConfigOk

`func (o *SignInWithExternalTokenRequest) GetSteamConfigOk() (*SteamConfig, bool)`

GetSteamConfigOk returns a tuple with the SteamConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamConfig

`func (o *SignInWithExternalTokenRequest) SetSteamConfig(v SteamConfig)`

SetSteamConfig sets SteamConfig field to given value.

### HasSteamConfig

`func (o *SignInWithExternalTokenRequest) HasSteamConfig() bool`

HasSteamConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


