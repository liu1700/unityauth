# CodeLinkConfirmationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignInCode** | Pointer to **string** | The code returned in the GenerateCodeResponse. | [optional] 
**SessionToken** | Pointer to **string** | The authenticated device session token, for added security. | [optional] 
**IdProvider** | Pointer to **string** | This is the id provider type. Only for consoles. | [optional] 
**ExternalToken** | Pointer to **string** | External token to validate the user. Only for consoles. | [optional] 

## Methods

### NewCodeLinkConfirmationRequest

`func NewCodeLinkConfirmationRequest() *CodeLinkConfirmationRequest`

NewCodeLinkConfirmationRequest instantiates a new CodeLinkConfirmationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeLinkConfirmationRequestWithDefaults

`func NewCodeLinkConfirmationRequestWithDefaults() *CodeLinkConfirmationRequest`

NewCodeLinkConfirmationRequestWithDefaults instantiates a new CodeLinkConfirmationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignInCode

`func (o *CodeLinkConfirmationRequest) GetSignInCode() string`

GetSignInCode returns the SignInCode field if non-nil, zero value otherwise.

### GetSignInCodeOk

`func (o *CodeLinkConfirmationRequest) GetSignInCodeOk() (*string, bool)`

GetSignInCodeOk returns a tuple with the SignInCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInCode

`func (o *CodeLinkConfirmationRequest) SetSignInCode(v string)`

SetSignInCode sets SignInCode field to given value.

### HasSignInCode

`func (o *CodeLinkConfirmationRequest) HasSignInCode() bool`

HasSignInCode returns a boolean if a field has been set.

### GetSessionToken

`func (o *CodeLinkConfirmationRequest) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *CodeLinkConfirmationRequest) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *CodeLinkConfirmationRequest) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.

### HasSessionToken

`func (o *CodeLinkConfirmationRequest) HasSessionToken() bool`

HasSessionToken returns a boolean if a field has been set.

### GetIdProvider

`func (o *CodeLinkConfirmationRequest) GetIdProvider() string`

GetIdProvider returns the IdProvider field if non-nil, zero value otherwise.

### GetIdProviderOk

`func (o *CodeLinkConfirmationRequest) GetIdProviderOk() (*string, bool)`

GetIdProviderOk returns a tuple with the IdProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdProvider

`func (o *CodeLinkConfirmationRequest) SetIdProvider(v string)`

SetIdProvider sets IdProvider field to given value.

### HasIdProvider

`func (o *CodeLinkConfirmationRequest) HasIdProvider() bool`

HasIdProvider returns a boolean if a field has been set.

### GetExternalToken

`func (o *CodeLinkConfirmationRequest) GetExternalToken() string`

GetExternalToken returns the ExternalToken field if non-nil, zero value otherwise.

### GetExternalTokenOk

`func (o *CodeLinkConfirmationRequest) GetExternalTokenOk() (*string, bool)`

GetExternalTokenOk returns a tuple with the ExternalToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalToken

`func (o *CodeLinkConfirmationRequest) SetExternalToken(v string)`

SetExternalToken sets ExternalToken field to given value.

### HasExternalToken

`func (o *CodeLinkConfirmationRequest) HasExternalToken() bool`

HasExternalToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


