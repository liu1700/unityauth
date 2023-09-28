# SignInWithCodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodeVerifier** | Pointer to **string** | Verifier for PKCE validation. | [optional] 

## Methods

### NewSignInWithCodeRequest

`func NewSignInWithCodeRequest() *SignInWithCodeRequest`

NewSignInWithCodeRequest instantiates a new SignInWithCodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignInWithCodeRequestWithDefaults

`func NewSignInWithCodeRequestWithDefaults() *SignInWithCodeRequest`

NewSignInWithCodeRequestWithDefaults instantiates a new SignInWithCodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodeVerifier

`func (o *SignInWithCodeRequest) GetCodeVerifier() string`

GetCodeVerifier returns the CodeVerifier field if non-nil, zero value otherwise.

### GetCodeVerifierOk

`func (o *SignInWithCodeRequest) GetCodeVerifierOk() (*string, bool)`

GetCodeVerifierOk returns a tuple with the CodeVerifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeVerifier

`func (o *SignInWithCodeRequest) SetCodeVerifier(v string)`

SetCodeVerifier sets CodeVerifier field to given value.

### HasCodeVerifier

`func (o *SignInWithCodeRequest) HasCodeVerifier() bool`

HasCodeVerifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


