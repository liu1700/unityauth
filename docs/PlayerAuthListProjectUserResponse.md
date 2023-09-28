# PlayerAuthListProjectUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **string** | The next page token if available. Needed for the next batch GET request of players. | [optional] 
**Results** | Pointer to [**[]PlayerAuthListProjectUserResponseUser**](PlayerAuthListProjectUserResponseUser.md) | Resulting collection of Players. | [optional] 

## Methods

### NewPlayerAuthListProjectUserResponse

`func NewPlayerAuthListProjectUserResponse() *PlayerAuthListProjectUserResponse`

NewPlayerAuthListProjectUserResponse instantiates a new PlayerAuthListProjectUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerAuthListProjectUserResponseWithDefaults

`func NewPlayerAuthListProjectUserResponseWithDefaults() *PlayerAuthListProjectUserResponse`

NewPlayerAuthListProjectUserResponseWithDefaults instantiates a new PlayerAuthListProjectUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *PlayerAuthListProjectUserResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PlayerAuthListProjectUserResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PlayerAuthListProjectUserResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PlayerAuthListProjectUserResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetResults

`func (o *PlayerAuthListProjectUserResponse) GetResults() []PlayerAuthListProjectUserResponseUser`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PlayerAuthListProjectUserResponse) GetResultsOk() (*[]PlayerAuthListProjectUserResponseUser, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PlayerAuthListProjectUserResponse) SetResults(v []PlayerAuthListProjectUserResponseUser)`

SetResults sets Results field to given value.

### HasResults

`func (o *PlayerAuthListProjectUserResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


