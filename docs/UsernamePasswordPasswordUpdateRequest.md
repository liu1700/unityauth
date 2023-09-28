# UsernamePasswordPasswordUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | The password. Length must be between 8-30 and contain at least one uppercase letter, at least one lowercase letter, at least one number and at least one symbol. | [optional] 
**NewPassword** | Pointer to **string** | The password to be changed. Length must be between 8-30 and contain at least one uppercase letter, at least one lowercase letter, at least one number and at least one symbol. | [optional] 

## Methods

### NewUsernamePasswordPasswordUpdateRequest

`func NewUsernamePasswordPasswordUpdateRequest() *UsernamePasswordPasswordUpdateRequest`

NewUsernamePasswordPasswordUpdateRequest instantiates a new UsernamePasswordPasswordUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsernamePasswordPasswordUpdateRequestWithDefaults

`func NewUsernamePasswordPasswordUpdateRequestWithDefaults() *UsernamePasswordPasswordUpdateRequest`

NewUsernamePasswordPasswordUpdateRequestWithDefaults instantiates a new UsernamePasswordPasswordUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *UsernamePasswordPasswordUpdateRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UsernamePasswordPasswordUpdateRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UsernamePasswordPasswordUpdateRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UsernamePasswordPasswordUpdateRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetNewPassword

`func (o *UsernamePasswordPasswordUpdateRequest) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *UsernamePasswordPasswordUpdateRequest) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *UsernamePasswordPasswordUpdateRequest) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *UsernamePasswordPasswordUpdateRequest) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


