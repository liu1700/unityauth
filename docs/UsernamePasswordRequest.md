# UsernamePasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | The username. Case insensitive. Length must be between 3-20 with the allowed characters a-z, 0-9 and the symbols [.][-][@][_]. | [optional] 
**Password** | Pointer to **string** | The password. Length must be between 8-30 and contain at least one uppercase letter, at least one lowercase letter, at least one number and at least one symbol. | [optional] 

## Methods

### NewUsernamePasswordRequest

`func NewUsernamePasswordRequest() *UsernamePasswordRequest`

NewUsernamePasswordRequest instantiates a new UsernamePasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsernamePasswordRequestWithDefaults

`func NewUsernamePasswordRequestWithDefaults() *UsernamePasswordRequest`

NewUsernamePasswordRequestWithDefaults instantiates a new UsernamePasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UsernamePasswordRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UsernamePasswordRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UsernamePasswordRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UsernamePasswordRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *UsernamePasswordRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UsernamePasswordRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UsernamePasswordRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UsernamePasswordRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


