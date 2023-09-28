# GenerateCodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | Human-readable string to identify the requester device. | [optional] 
**CodeChallenge** | Pointer to **string** | SHA-256 string challenge for PKCE validation | [optional] 

## Methods

### NewGenerateCodeRequest

`func NewGenerateCodeRequest() *GenerateCodeRequest`

NewGenerateCodeRequest instantiates a new GenerateCodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateCodeRequestWithDefaults

`func NewGenerateCodeRequestWithDefaults() *GenerateCodeRequest`

NewGenerateCodeRequestWithDefaults instantiates a new GenerateCodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *GenerateCodeRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *GenerateCodeRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *GenerateCodeRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *GenerateCodeRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetCodeChallenge

`func (o *GenerateCodeRequest) GetCodeChallenge() string`

GetCodeChallenge returns the CodeChallenge field if non-nil, zero value otherwise.

### GetCodeChallengeOk

`func (o *GenerateCodeRequest) GetCodeChallengeOk() (*string, bool)`

GetCodeChallengeOk returns a tuple with the CodeChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeChallenge

`func (o *GenerateCodeRequest) SetCodeChallenge(v string)`

SetCodeChallenge sets CodeChallenge field to given value.

### HasCodeChallenge

`func (o *GenerateCodeRequest) HasCodeChallenge() bool`

HasCodeChallenge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


