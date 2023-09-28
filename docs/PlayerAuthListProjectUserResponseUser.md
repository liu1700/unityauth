# PlayerAuthListProjectUserResponseUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The player Id. The Id is unique within the Id domain. | [optional] 
**Disabled** | Pointer to **bool** | Whether the player is enabled or disabled by admin. | [optional] 
**ExternalIds** | Pointer to [**[]PlayerAuthListProjectUserResponseExternalId**](PlayerAuthListProjectUserResponseExternalId.md) | List of externalIds that have been linked to the player. | [optional] 
**CreatedAt** | Pointer to **string** | When the player was created at. It is a unix timestamp. | [optional] 
**LastLoginAt** | Pointer to **string** | When the player last logged in. It is a unix timestamp. | [optional] 

## Methods

### NewPlayerAuthListProjectUserResponseUser

`func NewPlayerAuthListProjectUserResponseUser() *PlayerAuthListProjectUserResponseUser`

NewPlayerAuthListProjectUserResponseUser instantiates a new PlayerAuthListProjectUserResponseUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerAuthListProjectUserResponseUserWithDefaults

`func NewPlayerAuthListProjectUserResponseUserWithDefaults() *PlayerAuthListProjectUserResponseUser`

NewPlayerAuthListProjectUserResponseUserWithDefaults instantiates a new PlayerAuthListProjectUserResponseUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlayerAuthListProjectUserResponseUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlayerAuthListProjectUserResponseUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlayerAuthListProjectUserResponseUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlayerAuthListProjectUserResponseUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisabled

`func (o *PlayerAuthListProjectUserResponseUser) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *PlayerAuthListProjectUserResponseUser) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *PlayerAuthListProjectUserResponseUser) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *PlayerAuthListProjectUserResponseUser) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetExternalIds

`func (o *PlayerAuthListProjectUserResponseUser) GetExternalIds() []PlayerAuthListProjectUserResponseExternalId`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *PlayerAuthListProjectUserResponseUser) GetExternalIdsOk() (*[]PlayerAuthListProjectUserResponseExternalId, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *PlayerAuthListProjectUserResponseUser) SetExternalIds(v []PlayerAuthListProjectUserResponseExternalId)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *PlayerAuthListProjectUserResponseUser) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PlayerAuthListProjectUserResponseUser) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PlayerAuthListProjectUserResponseUser) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PlayerAuthListProjectUserResponseUser) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PlayerAuthListProjectUserResponseUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastLoginAt

`func (o *PlayerAuthListProjectUserResponseUser) GetLastLoginAt() string`

GetLastLoginAt returns the LastLoginAt field if non-nil, zero value otherwise.

### GetLastLoginAtOk

`func (o *PlayerAuthListProjectUserResponseUser) GetLastLoginAtOk() (*string, bool)`

GetLastLoginAtOk returns a tuple with the LastLoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginAt

`func (o *PlayerAuthListProjectUserResponseUser) SetLastLoginAt(v string)`

SetLastLoginAt sets LastLoginAt field to given value.

### HasLastLoginAt

`func (o *PlayerAuthListProjectUserResponseUser) HasLastLoginAt() bool`

HasLastLoginAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


