# AuthenticationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresIn** | Pointer to **int32** | The number of seconds in which the Id token expires. By default it&#39;s 3600. | [optional] 
**IdToken** | Pointer to **string** | The Id token of the authenticated player. | [optional] 
**SessionToken** | Pointer to **string** | The session token of the authenticated player. This token can be used to sign-in the player again. | [optional] 
**User** | Pointer to [**Player**](Player.md) |  | [optional] 
**UserId** | Pointer to **string** | The Id of the authenticated player. If a project is specified in the request, this field represents the Id of the project scoped player. | [optional] 

## Methods

### NewAuthenticationResponse

`func NewAuthenticationResponse() *AuthenticationResponse`

NewAuthenticationResponse instantiates a new AuthenticationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationResponseWithDefaults

`func NewAuthenticationResponseWithDefaults() *AuthenticationResponse`

NewAuthenticationResponseWithDefaults instantiates a new AuthenticationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresIn

`func (o *AuthenticationResponse) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *AuthenticationResponse) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *AuthenticationResponse) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *AuthenticationResponse) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetIdToken

`func (o *AuthenticationResponse) GetIdToken() string`

GetIdToken returns the IdToken field if non-nil, zero value otherwise.

### GetIdTokenOk

`func (o *AuthenticationResponse) GetIdTokenOk() (*string, bool)`

GetIdTokenOk returns a tuple with the IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdToken

`func (o *AuthenticationResponse) SetIdToken(v string)`

SetIdToken sets IdToken field to given value.

### HasIdToken

`func (o *AuthenticationResponse) HasIdToken() bool`

HasIdToken returns a boolean if a field has been set.

### GetSessionToken

`func (o *AuthenticationResponse) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *AuthenticationResponse) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *AuthenticationResponse) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.

### HasSessionToken

`func (o *AuthenticationResponse) HasSessionToken() bool`

HasSessionToken returns a boolean if a field has been set.

### GetUser

`func (o *AuthenticationResponse) GetUser() Player`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AuthenticationResponse) GetUserOk() (*Player, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AuthenticationResponse) SetUser(v Player)`

SetUser sets User field to given value.

### HasUser

`func (o *AuthenticationResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserId

`func (o *AuthenticationResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AuthenticationResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AuthenticationResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AuthenticationResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


