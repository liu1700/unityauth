# UsernamePasswordResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | The username. Case insensitive. Length must be between 3-20 with the allowed characters a-z, 0-9 and the symbols [.][-][@][_]. | [optional] 
**CreatedAt** | Pointer to **string** | When the username/password account was created. It is a unix timestamp. | [optional] 
**LastLoginAt** | Pointer to **string** | When the username/password account last logged in. It is a unix timestamp. | [optional] 
**PasswordUpdatedAt** | Pointer to **string** | When the username/password account password was last updated. It is a unix timestamp. | [optional] 

## Methods

### NewUsernamePasswordResponse

`func NewUsernamePasswordResponse() *UsernamePasswordResponse`

NewUsernamePasswordResponse instantiates a new UsernamePasswordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsernamePasswordResponseWithDefaults

`func NewUsernamePasswordResponseWithDefaults() *UsernamePasswordResponse`

NewUsernamePasswordResponseWithDefaults instantiates a new UsernamePasswordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UsernamePasswordResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UsernamePasswordResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UsernamePasswordResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UsernamePasswordResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UsernamePasswordResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UsernamePasswordResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UsernamePasswordResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UsernamePasswordResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastLoginAt

`func (o *UsernamePasswordResponse) GetLastLoginAt() string`

GetLastLoginAt returns the LastLoginAt field if non-nil, zero value otherwise.

### GetLastLoginAtOk

`func (o *UsernamePasswordResponse) GetLastLoginAtOk() (*string, bool)`

GetLastLoginAtOk returns a tuple with the LastLoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginAt

`func (o *UsernamePasswordResponse) SetLastLoginAt(v string)`

SetLastLoginAt sets LastLoginAt field to given value.

### HasLastLoginAt

`func (o *UsernamePasswordResponse) HasLastLoginAt() bool`

HasLastLoginAt returns a boolean if a field has been set.

### GetPasswordUpdatedAt

`func (o *UsernamePasswordResponse) GetPasswordUpdatedAt() string`

GetPasswordUpdatedAt returns the PasswordUpdatedAt field if non-nil, zero value otherwise.

### GetPasswordUpdatedAtOk

`func (o *UsernamePasswordResponse) GetPasswordUpdatedAtOk() (*string, bool)`

GetPasswordUpdatedAtOk returns a tuple with the PasswordUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordUpdatedAt

`func (o *UsernamePasswordResponse) SetPasswordUpdatedAt(v string)`

SetPasswordUpdatedAt sets PasswordUpdatedAt field to given value.

### HasPasswordUpdatedAt

`func (o *UsernamePasswordResponse) HasPasswordUpdatedAt() bool`

HasPasswordUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


