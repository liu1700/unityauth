# SignUpAnonymouslyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nonce** | Pointer to **string** | String value used to associate a client session with an Id Token, and to mitigate replay attacks. If this field is provided, the nonce claim in response id token has a matching value. | [optional] 

## Methods

### NewSignUpAnonymouslyRequest

`func NewSignUpAnonymouslyRequest() *SignUpAnonymouslyRequest`

NewSignUpAnonymouslyRequest instantiates a new SignUpAnonymouslyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignUpAnonymouslyRequestWithDefaults

`func NewSignUpAnonymouslyRequestWithDefaults() *SignUpAnonymouslyRequest`

NewSignUpAnonymouslyRequestWithDefaults instantiates a new SignUpAnonymouslyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNonce

`func (o *SignUpAnonymouslyRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *SignUpAnonymouslyRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *SignUpAnonymouslyRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *SignUpAnonymouslyRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


