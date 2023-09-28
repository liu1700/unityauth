# SignInWithSessionTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nonce** | Pointer to **string** | String value used to associate a Client session with an Id Token, and to mitigate replay attacks. If this field is provided, the nonce claim in response Id token has a matching value. | [optional] 
**SessionToken** | Pointer to **string** | The session token of the player. | [optional] 

## Methods

### NewSignInWithSessionTokenRequest

`func NewSignInWithSessionTokenRequest() *SignInWithSessionTokenRequest`

NewSignInWithSessionTokenRequest instantiates a new SignInWithSessionTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignInWithSessionTokenRequestWithDefaults

`func NewSignInWithSessionTokenRequestWithDefaults() *SignInWithSessionTokenRequest`

NewSignInWithSessionTokenRequestWithDefaults instantiates a new SignInWithSessionTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNonce

`func (o *SignInWithSessionTokenRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *SignInWithSessionTokenRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *SignInWithSessionTokenRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *SignInWithSessionTokenRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetSessionToken

`func (o *SignInWithSessionTokenRequest) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *SignInWithSessionTokenRequest) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *SignInWithSessionTokenRequest) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.

### HasSessionToken

`func (o *SignInWithSessionTokenRequest) HasSessionToken() bool`

HasSessionToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


