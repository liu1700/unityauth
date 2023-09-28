# CodeLinkInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignInCode** | Pointer to **string** | The code from which to get the info. | [optional] 

## Methods

### NewCodeLinkInfoRequest

`func NewCodeLinkInfoRequest() *CodeLinkInfoRequest`

NewCodeLinkInfoRequest instantiates a new CodeLinkInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeLinkInfoRequestWithDefaults

`func NewCodeLinkInfoRequestWithDefaults() *CodeLinkInfoRequest`

NewCodeLinkInfoRequestWithDefaults instantiates a new CodeLinkInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignInCode

`func (o *CodeLinkInfoRequest) GetSignInCode() string`

GetSignInCode returns the SignInCode field if non-nil, zero value otherwise.

### GetSignInCodeOk

`func (o *CodeLinkInfoRequest) GetSignInCodeOk() (*string, bool)`

GetSignInCodeOk returns a tuple with the SignInCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInCode

`func (o *CodeLinkInfoRequest) SetSignInCode(v string)`

SetSignInCode sets SignInCode field to given value.

### HasSignInCode

`func (o *CodeLinkInfoRequest) HasSignInCode() bool`

HasSignInCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


