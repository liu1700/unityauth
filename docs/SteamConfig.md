# SteamConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | Pointer to **string** | If this identity string was passed to GetAuthTicketForWebApi then the same value must be provided as the identity property. This value should be the comprised of alphanumeric characters with a length between 5 and 30. | [optional] 

## Methods

### NewSteamConfig

`func NewSteamConfig() *SteamConfig`

NewSteamConfig instantiates a new SteamConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSteamConfigWithDefaults

`func NewSteamConfigWithDefaults() *SteamConfig`

NewSteamConfigWithDefaults instantiates a new SteamConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *SteamConfig) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *SteamConfig) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *SteamConfig) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *SteamConfig) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


