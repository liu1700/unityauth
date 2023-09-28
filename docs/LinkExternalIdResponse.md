# LinkExternalIdResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresIn** | Pointer to **int32** | The number of seconds in which the Id token expires. By default it&#39;s 3600. | [optional] 
**IdToken** | Pointer to **string** | The Id token of the authenticated player. | [optional] 
**SessionToken** | Pointer to **string** | The session token of the authenticated player. This token can be used to sign-in the player again. | [optional] 
**User** | Pointer to [**Player**](Player.md) |  | [optional] 
**UserId** | Pointer to **string** | The Id of the authenticated player. If a project is specified in the request, this field represents the Id of the project scoped player. | [optional] 

## Methods

### NewLinkExternalIdResponse

`func NewLinkExternalIdResponse() *LinkExternalIdResponse`

NewLinkExternalIdResponse instantiates a new LinkExternalIdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkExternalIdResponseWithDefaults

`func NewLinkExternalIdResponseWithDefaults() *LinkExternalIdResponse`

NewLinkExternalIdResponseWithDefaults instantiates a new LinkExternalIdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresIn

`func (o *LinkExternalIdResponse) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *LinkExternalIdResponse) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *LinkExternalIdResponse) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *LinkExternalIdResponse) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetIdToken

`func (o *LinkExternalIdResponse) GetIdToken() string`

GetIdToken returns the IdToken field if non-nil, zero value otherwise.

### GetIdTokenOk

`func (o *LinkExternalIdResponse) GetIdTokenOk() (*string, bool)`

GetIdTokenOk returns a tuple with the IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdToken

`func (o *LinkExternalIdResponse) SetIdToken(v string)`

SetIdToken sets IdToken field to given value.

### HasIdToken

`func (o *LinkExternalIdResponse) HasIdToken() bool`

HasIdToken returns a boolean if a field has been set.

### GetSessionToken

`func (o *LinkExternalIdResponse) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *LinkExternalIdResponse) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *LinkExternalIdResponse) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.

### HasSessionToken

`func (o *LinkExternalIdResponse) HasSessionToken() bool`

HasSessionToken returns a boolean if a field has been set.

### GetUser

`func (o *LinkExternalIdResponse) GetUser() Player`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LinkExternalIdResponse) GetUserOk() (*Player, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LinkExternalIdResponse) SetUser(v Player)`

SetUser sets User field to given value.

### HasUser

`func (o *LinkExternalIdResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserId

`func (o *LinkExternalIdResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *LinkExternalIdResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *LinkExternalIdResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *LinkExternalIdResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


