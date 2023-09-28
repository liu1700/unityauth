# GenerateCodeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodeLinkSessionId** | Pointer to **string** | UUID identifying the code linking session. | [optional] 
**SignInCode** | Pointer to **string** | The code required to perform the code confirmation. | [optional] 
**Expiration** | Pointer to **int32** | The timestamp for when the code is no longer valid in unix time since epoch. | [optional] 

## Methods

### NewGenerateCodeResponse

`func NewGenerateCodeResponse() *GenerateCodeResponse`

NewGenerateCodeResponse instantiates a new GenerateCodeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateCodeResponseWithDefaults

`func NewGenerateCodeResponseWithDefaults() *GenerateCodeResponse`

NewGenerateCodeResponseWithDefaults instantiates a new GenerateCodeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodeLinkSessionId

`func (o *GenerateCodeResponse) GetCodeLinkSessionId() string`

GetCodeLinkSessionId returns the CodeLinkSessionId field if non-nil, zero value otherwise.

### GetCodeLinkSessionIdOk

`func (o *GenerateCodeResponse) GetCodeLinkSessionIdOk() (*string, bool)`

GetCodeLinkSessionIdOk returns a tuple with the CodeLinkSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeLinkSessionId

`func (o *GenerateCodeResponse) SetCodeLinkSessionId(v string)`

SetCodeLinkSessionId sets CodeLinkSessionId field to given value.

### HasCodeLinkSessionId

`func (o *GenerateCodeResponse) HasCodeLinkSessionId() bool`

HasCodeLinkSessionId returns a boolean if a field has been set.

### GetSignInCode

`func (o *GenerateCodeResponse) GetSignInCode() string`

GetSignInCode returns the SignInCode field if non-nil, zero value otherwise.

### GetSignInCodeOk

`func (o *GenerateCodeResponse) GetSignInCodeOk() (*string, bool)`

GetSignInCodeOk returns a tuple with the SignInCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInCode

`func (o *GenerateCodeResponse) SetSignInCode(v string)`

SetSignInCode sets SignInCode field to given value.

### HasSignInCode

`func (o *GenerateCodeResponse) HasSignInCode() bool`

HasSignInCode returns a boolean if a field has been set.

### GetExpiration

`func (o *GenerateCodeResponse) GetExpiration() int32`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *GenerateCodeResponse) GetExpirationOk() (*int32, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *GenerateCodeResponse) SetExpiration(v int32)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *GenerateCodeResponse) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


