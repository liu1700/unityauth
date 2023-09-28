# PlayerAuthExternalId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | The Id of the player from the external provider. | [optional] 
**ProviderId** | Pointer to **string** | The provider Id within the Id domain that provided the link. | [optional] 

## Methods

### NewPlayerAuthExternalId

`func NewPlayerAuthExternalId() *PlayerAuthExternalId`

NewPlayerAuthExternalId instantiates a new PlayerAuthExternalId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerAuthExternalIdWithDefaults

`func NewPlayerAuthExternalIdWithDefaults() *PlayerAuthExternalId`

NewPlayerAuthExternalIdWithDefaults instantiates a new PlayerAuthExternalId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *PlayerAuthExternalId) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *PlayerAuthExternalId) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *PlayerAuthExternalId) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *PlayerAuthExternalId) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProviderId

`func (o *PlayerAuthExternalId) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *PlayerAuthExternalId) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *PlayerAuthExternalId) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *PlayerAuthExternalId) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


