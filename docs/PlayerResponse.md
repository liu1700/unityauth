# PlayerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** | Whether the player is disabled by admin. | [optional] 
**ExternalIds** | Pointer to [**[]ExternalId**](ExternalId.md) | This is the response returned when get player is called. | [optional] 
**Id** | Pointer to **string** | The player Id. The Id is unique within the Id domain. | [optional] 
**CreatedAt** | Pointer to **string** | When the player was created. It is a unix timestamp. | [optional] 
**LastLoginAt** | Pointer to **string** | When the player last logged in. It is a unix timestamp. | [optional] 
**Usernamepassword** | Pointer to [**UsernamePasswordResponse**](UsernamePasswordResponse.md) |  | [optional] 

## Methods

### NewPlayerResponse

`func NewPlayerResponse() *PlayerResponse`

NewPlayerResponse instantiates a new PlayerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerResponseWithDefaults

`func NewPlayerResponseWithDefaults() *PlayerResponse`

NewPlayerResponseWithDefaults instantiates a new PlayerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *PlayerResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *PlayerResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *PlayerResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *PlayerResponse) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetExternalIds

`func (o *PlayerResponse) GetExternalIds() []ExternalId`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *PlayerResponse) GetExternalIdsOk() (*[]ExternalId, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *PlayerResponse) SetExternalIds(v []ExternalId)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *PlayerResponse) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetId

`func (o *PlayerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlayerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlayerResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlayerResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PlayerResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PlayerResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PlayerResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PlayerResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastLoginAt

`func (o *PlayerResponse) GetLastLoginAt() string`

GetLastLoginAt returns the LastLoginAt field if non-nil, zero value otherwise.

### GetLastLoginAtOk

`func (o *PlayerResponse) GetLastLoginAtOk() (*string, bool)`

GetLastLoginAtOk returns a tuple with the LastLoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginAt

`func (o *PlayerResponse) SetLastLoginAt(v string)`

SetLastLoginAt sets LastLoginAt field to given value.

### HasLastLoginAt

`func (o *PlayerResponse) HasLastLoginAt() bool`

HasLastLoginAt returns a boolean if a field has been set.

### GetUsernamepassword

`func (o *PlayerResponse) GetUsernamepassword() UsernamePasswordResponse`

GetUsernamepassword returns the Usernamepassword field if non-nil, zero value otherwise.

### GetUsernamepasswordOk

`func (o *PlayerResponse) GetUsernamepasswordOk() (*UsernamePasswordResponse, bool)`

GetUsernamepasswordOk returns a tuple with the Usernamepassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernamepassword

`func (o *PlayerResponse) SetUsernamepassword(v UsernamePasswordResponse)`

SetUsernamepassword sets Usernamepassword field to given value.

### HasUsernamepassword

`func (o *PlayerResponse) HasUsernamepassword() bool`

HasUsernamepassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


