# CodeLinkInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | Human-readable string to identify the requester device. | [optional] 
**Expiration** | Pointer to **int32** | The timestamp for when the code is no longer valid in unix time since epoch. | [optional] 

## Methods

### NewCodeLinkInfoResponse

`func NewCodeLinkInfoResponse() *CodeLinkInfoResponse`

NewCodeLinkInfoResponse instantiates a new CodeLinkInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeLinkInfoResponseWithDefaults

`func NewCodeLinkInfoResponseWithDefaults() *CodeLinkInfoResponse`

NewCodeLinkInfoResponseWithDefaults instantiates a new CodeLinkInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *CodeLinkInfoResponse) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CodeLinkInfoResponse) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CodeLinkInfoResponse) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *CodeLinkInfoResponse) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetExpiration

`func (o *CodeLinkInfoResponse) GetExpiration() int32`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *CodeLinkInfoResponse) GetExpirationOk() (*int32, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *CodeLinkInfoResponse) SetExpiration(v int32)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *CodeLinkInfoResponse) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


