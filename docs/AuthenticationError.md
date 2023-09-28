# AuthenticationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Detail** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewAuthenticationError

`func NewAuthenticationError() *AuthenticationError`

NewAuthenticationError instantiates a new AuthenticationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationErrorWithDefaults

`func NewAuthenticationErrorWithDefaults() *AuthenticationError`

NewAuthenticationErrorWithDefaults instantiates a new AuthenticationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AuthenticationError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AuthenticationError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AuthenticationError) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AuthenticationError) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *AuthenticationError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthenticationError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthenticationError) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthenticationError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDetail

`func (o *AuthenticationError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *AuthenticationError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *AuthenticationError) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *AuthenticationError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetDetails

`func (o *AuthenticationError) GetDetails() []map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AuthenticationError) GetDetailsOk() (*[]map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AuthenticationError) SetDetails(v []map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *AuthenticationError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

