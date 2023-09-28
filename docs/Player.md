# Player

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** | Whether the player is disabled by admin. | [optional] 
**ExternalIds** | Pointer to [**[]ExternalId**](ExternalId.md) | The list of linked external providers info. The list itself doesn&#39;t limit the number of external accounts linked with this player. In our current version, we allow only one external account of each type to link with the player. For example, you can&#39;t link two different Facebook Ids with the same player. | [optional] 
**Id** | Pointer to **string** | The player Id. The Id is unique within the Id domain. | [optional] 
**Username** | Pointer to **string** | The username used to sign in using the Username Password IdProvider. | [optional] 

## Methods

### NewPlayer

`func NewPlayer() *Player`

NewPlayer instantiates a new Player object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerWithDefaults

`func NewPlayerWithDefaults() *Player`

NewPlayerWithDefaults instantiates a new Player object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *Player) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Player) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Player) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Player) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetExternalIds

`func (o *Player) GetExternalIds() []ExternalId`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *Player) GetExternalIdsOk() (*[]ExternalId, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *Player) SetExternalIds(v []ExternalId)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *Player) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetId

`func (o *Player) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Player) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Player) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Player) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *Player) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Player) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Player) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Player) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


