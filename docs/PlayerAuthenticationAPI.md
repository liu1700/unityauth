# \PlayerAuthenticationAPI

All URIs are relative to *https://player-auth.services.api.unity.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CodeConfirmation**](PlayerAuthenticationAPI.md#CodeConfirmation) | **Post** /v1/authentication/code-link/confirm | Code Confirmation
[**DeletePlayer**](PlayerAuthenticationAPI.md#DeletePlayer) | **Delete** /v1/users/{PlayerId} | Delete Player
[**GenerateCode**](PlayerAuthenticationAPI.md#GenerateCode) | **Post** /v1/authentication/code-link/generate | Generate Code
[**GetCodeInfo**](PlayerAuthenticationAPI.md#GetCodeInfo) | **Post** /v1/authentication/code-link/info | Get Code Info
[**GetJSONWebKeySet**](PlayerAuthenticationAPI.md#GetJSONWebKeySet) | **Get** /.well-known/jwks.json | Get JWKS
[**GetPlayer**](PlayerAuthenticationAPI.md#GetPlayer) | **Get** /v1/users/{PlayerId} | Get Player
[**LinkExternalId**](PlayerAuthenticationAPI.md#LinkExternalId) | **Post** /v1/authentication/link/{idProvider} | Link External Id
[**SignInWithCode**](PlayerAuthenticationAPI.md#SignInWithCode) | **Post** /v1/authentication/code-link/sign-in/{CodeLinkSessionId} | Sign In With Code
[**SignInWithExternalToken**](PlayerAuthenticationAPI.md#SignInWithExternalToken) | **Post** /v1/authentication/external-token/{idProvider} | External Token Sign In
[**SignInWithSessionToken**](PlayerAuthenticationAPI.md#SignInWithSessionToken) | **Post** /v1/authentication/session-token | Session Token Sign In
[**SignInWithUsernamePassword**](PlayerAuthenticationAPI.md#SignInWithUsernamePassword) | **Post** /v1/authentication/usernamepassword/sign-in | Username Password Sign In
[**SignUpAnonymously**](PlayerAuthenticationAPI.md#SignUpAnonymously) | **Post** /v1/authentication/anonymous | Anonymous Sign Up
[**SignUpWithUsernamePassword**](PlayerAuthenticationAPI.md#SignUpWithUsernamePassword) | **Post** /v1/authentication/usernamepassword/sign-up | Username Password Sign Up
[**UnlinkExternalId**](PlayerAuthenticationAPI.md#UnlinkExternalId) | **Post** /v1/authentication/unlink/{idProvider} | Unlink External Id
[**UsernamePasswordUpdatePassword**](PlayerAuthenticationAPI.md#UsernamePasswordUpdatePassword) | **Post** /v1/authentication/usernamepassword/update-password | Username Password Update Password



## CodeConfirmation

> CodeConfirmation(ctx).ProjectId(projectId).BearerAuth(bearerAuth).UnityEnvironment(unityEnvironment).CodeLinkConfirmationRequest(codeLinkConfirmationRequest).Execute()

Code Confirmation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/liu1700/unityauth"
)

func main() {
    projectId := "8bdacc33-6eef-4577-beb0-633c86259f5b" // string | This is the Unity Project Id. It is a uuid format. 
    bearerAuth := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c" // string | This is the bearer token for the user authorized to call this API. 
    unityEnvironment := "production" // string | This is the Environment you want to authorize a player to access. It is the name of the Environment. If this header is not specified, then the default Environment is used. An invalid environment name is not an acceptable input.  (optional)
    codeLinkConfirmationRequest := *openapiclient.NewCodeLinkConfirmationRequest() // CodeLinkConfirmationRequest | Code confirmation request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PlayerAuthenticationAPI.CodeConfirmation(context.Background()).ProjectId(projectId).BearerAuth(bearerAuth).UnityEnvironment(unityEnvironment).CodeLinkConfirmationRequest(codeLinkConfirmationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerAuthenticationAPI.CodeConfirmation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCodeConfirmationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** | This is the Unity Project Id. It is a uuid format.  | 
 **bearerAuth** | **string** | This is the bearer token for the user authorized to call this API.  | 
 **unityEnvironment** | **string** | This is the Environment you want to authorize a player to access. It is the name of the Environment. If this header is not specified, then the default Environment is used. An invalid environment name is not an acceptable input.  | 
 **codeLinkConfirmationRequest** | [**CodeLinkConfirmationRequest**](CodeLinkConfirmationRequest.md) | Code confirmation request. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePlayer

> DeletePlayer(ctx, playerId).ProjectId(projectId).Execute()

Delete Player



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/liu1700/unityauth"
)

func main() {
    playerId := "99i9ju8juh" // string | This is the player id. 
    projectId := "8bdacc33-6eef-4577-beb0-633c86259f5b" // string | This is the Unity Project Id. It is a uuid format. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PlayerAuthenticationAPI.DeletePlayer(context.Background(), playerId).ProjectId(projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerAuthenticationAPI.DeletePlayer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playerId** | **string** | This is the player id.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlayerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectId** | **string** | This is the Unity Project Id. It is a uuid format.  | 

### Return type

 (empty response body)

### Authorization

[Client](../README.md#Client)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateCode

> GenerateCodeResponse GenerateCode(ctx).ProjectId(projectId).UnityEnvironment(unityEnvironment).GenerateCodeRequest(generateCodeRequest).Execute()

Generate Code



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/liu1700/unityauth"
)

func main() {
    projectId := "8bdacc33-6eef-4577-beb0-633c86259f5b" // string | This is the Unity Project Id. It is a uuid format. 
    unityEnvironment := "production" // string | This is the Environment you want to authorize a player to access. It is the name of the Environment. If this header is not specified, then the default Environment is used. An invalid environment name is not an acceptable input.  (optional)
    generateCodeRequest := *openapiclient.NewGenerateCodeRequest() // GenerateCodeRequest | Generate Code Request Body (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerAuthenticationAPI.GenerateCode(context.Background()).ProjectId(projectId).UnityEnvironment(unityEnvironment).GenerateCodeRequest(generateCodeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerAuthenticationAPI.GenerateCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateCode`: GenerateCodeResponse
    fmt.Fprintf(os.Stdout, "Response from `PlayerAuthenticationAPI.GenerateCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** | This is the Unity Project Id. It is a uuid format.  | 
 **unityEnvironment** | **string** | This is the Environment you want to authorize a player to access. It is the name of the Environment. If this header is not specified, then the default Environment is used. An invalid environment name is not an acceptable input.  | 
 **generateCodeRequest** | [**GenerateCodeRequest**](GenerateCodeRequest.md) | Generate Code Request Body | 

### Return type

[**GenerateCodeResponse**](GenerateCodeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCodeInfo

> CodeLinkInfoResponse GetCodeInfo(ctx).ProjectId(projectId).CodeLinkInfoRequest(codeLinkInfoRequest).UnityEnvironment(unityEnvironment).Execute()

Get Code Info



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/liu1700/unityauth"
)

func main() {
    projectId := "8bdacc33-6eef-4577-beb0-633c86259f5b" // string | This is the Unity Project Id. It is a uuid format. 
    codeLinkInfoRequest := *openapiclient.NewCodeLinkInfoRequest() // CodeLinkInfoRequest | Code Link Request for Sign In
    unityEnvironment := "production" // string | This is the Environment you want to authorize a player to access. It is the name of the Environment. If this header is not specified, then the default Environment is used. An invalid environment name is not an acceptable input.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerAuthenticationAPI.GetCodeInfo(context.Background()).ProjectId(projectId).CodeLinkInfoRequest(codeLinkInfoRequest).UnityEnvironment(unityEnvironment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerAuthenticationAPI.GetCodeInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCodeInfo`: CodeLinkInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `PlayerAuthenticationAPI.GetCodeInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCodeInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** | This is the Unity Project Id. It is a uuid format.  | 
 **codeLinkInfoRequest** | [**CodeLinkInfoRequest**](CodeLinkInfoRequest.md) | Code Link Request for Sign In | 
 **unityEnvironment** | **string** | This is the Environment you want to authorize a player to access. It is the name of the Environment. If this header is not specified, then the default Environment is used. An invalid environment name is not an acceptable input.  | 

### Return type

[**CodeLinkInfoResponse**](CodeLinkInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJSONWebKeySet

> JSONWebKeySet GetJSONWebKeySet(ctx).Execute()

Get JWKS



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/liu1700/unityauth"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerAuthenticationAPI.GetJSONWebKeySet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerAuthenticationAPI.GetJSONWebKeySet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJSONWebKeySet`: JSONWebKeySet
    fmt.Fprintf(os.Stdout, "Response from `PlayerAuthenticationAPI.GetJSONWebKeySet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetJSONWebKeySetRequest struct via the builder pattern


### Return type

[**JSONWebKeySet**](JSONWebKeySet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayer

> PlayerResponse GetPlayer(ctx, playerId).ProjectId(projectId).Execute()

Get Player



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/liu1700/unityauth"
)

func main() {
    playerId := "99i9ju8juh" // string | This is the player id. 
    projectId := "8bdacc33-6eef-4577-beb0-633c86259f5b" // string | This is the Unity Project Id. It is a uuid format. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerAuthenticationAPI.GetPlayer(context.Background(), playerId).ProjectId(projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerAuthenticationAPI.GetPlayer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlayer`: PlayerResponse
    fmt.Fprintf(os.Stdout, "Response from `PlayerAuthenticationAPI.GetPlayer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playerId** | **string** | This is the player id.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectId** | **string** | This is the Unity Project Id. It is a uuid format.  | 

### Return type

[**PlayerResponse**](PlayerResponse.md)

### Authorization

[Client](../README.md#Client)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkExternalId

> LinkExternalIdResponse LinkExternalId(ctx, idProvider).ProjectId(projectId).LinkExternalIdRequest(linkExternalIdRequest).Execute()

Link External Id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/liu1700/unityauth"
)

func main() {
    idProvider := "identity-provider-name" // string | This is the id provider type. 
    projectId := "8bdacc33-6eef-4577-beb0-633c86259f5b" // string | This is the Unity Project Id. It is a uuid format. 
    linkExternalIdRequest := *openapiclient.NewLinkExternalIdRequest() // LinkExternalIdRequest | Link External Id request body

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerAuthenticationAPI.LinkExternalId(context.Background(), idProvider).ProjectId(projectId).LinkExternalIdRequest(linkExternalIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerAuthenticationAPI.LinkExternalId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LinkExternalId`: LinkExternalIdResponse
    fmt.Fprintf(os.Stdout, "Response from `PlayerAuthenticationAPI.LinkExternalId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idProvider** | **string** | This is the id provider type.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLinkExternalIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectId** | **string** | This is the Unity Project Id. It is a uuid format.  | 
 **linkExternalIdRequest** | [**LinkExternalIdRequest**](LinkExternalIdRequest.md) | Link External Id request body | 

### Return type

[**LinkExternalIdResponse**](LinkExternalIdResponse.md)

### Authorization

[Client](../README.md#Client)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignInWithCode

> AuthenticationResponse SignInWithCode(ctx, codeLinkSessionId).ProjectId(projectId).SignInWithCodeRequest(signInWithCodeRequest).UnityEnvironment(unityEnvironment).Execute()

Sign In With Code



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/liu1700/unityauth"
)

func main() {
    projectId := "8bdacc33-6eef-4577-beb0-633c86259f5b" // string | This is the Unity Project Id. It is a uuid format. 
    signInWithCodeRequest := *openapiclient.NewSignInWithCodeRequest() // SignInWithCodeRequest | Code Link Request for Sign In
    unityEnvironment := "production" // string | This is the Environment you want to authorize a player to access. It is the name of the Environment. If this header is not specified, then the default Environment is used. An invalid environment name is not an acceptable input.  (optional)
    codeLinkSessionId := "codeLinkSessionId_example" // string | An identifier for the device requesting the code sign in. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerAuthenticationAPI.SignInWithCode(context.Background(), codeLinkSessionId).ProjectId(projectId).SignInWithCodeRequest(signInWithCodeRequest).UnityEnvironment(unityEnvironment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerAuthenticationAPI.SignInWithCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SignInWithCode`: AuthenticationResponse
    fmt.Fprintf(os.Stdout, "Response from `PlayerAuthenticationAPI.SignInWithCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**codeLinkSessionId** | **string** | An identifier for the device requesting the code sign in. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSignInWithCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** | This is the Unity Project Id. It is a uuid format.  | 
 **signInWithCodeRequest** | [**SignInWithCodeRequest**](SignInWithCodeRequest.md) | Code Link Request for Sign In | 
 **unityEnvironment** | **string** | This is the Environment you want to authorize a player to access. It is the name of the Environment. If this header is not specified, then the default Environment is used. An invalid environment name is not an acceptable input.  | 


### Return type

[**AuthenticationResponse**](AuthenticationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignInWithExternalToken

> AuthenticationResponse SignInWithExternalToken(ctx, idProvider).ProjectId(projectId).SignInWithExternalTokenRequest(signInWithExternalTokenRequest).UnityEnvironment(unityEnvironment).Execute()

External Token Sign In



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/liu1700/unityauth"
)

func main() {
    idProvider := "identity-provider-name" // string | This is the id provider type. 
    projectId := "8bdacc33-6eef-4577-beb0-633c86259f5b" // string | This is the Unity Project Id. It is a uuid format. 
    signInWithExternalTokenRequest := *openapiclient.NewSignInWithExternalTokenRequest() // SignInWithExternalTokenRequest | External Token Authentication request body
    unityEnvironment := "production" // string | This is the Environment you want to authorize a player to access. It is the name of the Environment. If this header is not specified, then the default Environment is used. An invalid environment name is not an acceptable input.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerAuthenticationAPI.SignInWithExternalToken(context.Background(), idProvider).ProjectId(projectId).SignInWithExternalTokenRequest(signInWithExternalTokenRequest).UnityEnvironment(unityEnvironment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerAuthenticationAPI.SignInWithExternalToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SignInWithExternalToken`: AuthenticationResponse
    fmt.Fprintf(os.Stdout, "Response from `PlayerAuthenticationAPI.SignInWithExternalToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idProvider** | **string** | This is the id provider type.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSignInWithExternalTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectId** | **string** | This is the Unity Project Id. It is a uuid format.  | 
 **signInWithExternalTokenRequest** | [**SignInWithExternalTokenRequest**](SignInWithExternalTokenRequest.md) | External Token Authentication request body | 
 **unityEnvironment** | **string** | This is the Environment you want to authorize a player to access. It is the name of the Environment. If this header is not specified, then the default Environment is used. An invalid environment name is not an acceptable input.  | 

### Return type

[**AuthenticationResponse**](AuthenticationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignInWithSessionToken

> AuthenticationResponse SignInWithSessionToken(ctx).ProjectId(projectId).SignInWithSessionTokenRequest(signInWithSessionTokenRequest).UnityEnvironment(unityEnvironment).Execute()

Session Token Sign In



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/liu1700/unityauth"
)

func main() {
    projectId := "8bdacc33-6eef-4577-beb0-633c86259f5b" // string | This is the Unity Project Id. It is a uuid format. 
    signInWithSessionTokenRequest := *openapiclient.NewSignInWithSessionTokenRequest() // SignInWithSessionTokenRequest | Session Token Authentication request body
    unityEnvironment := "production" // string | This is the Environment you want to authorize a player to access. It is the name of the Environment. If this header is not specified, then the default Environment is used. An invalid environment name is not an acceptable input.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerAuthenticationAPI.SignInWithSessionToken(context.Background()).ProjectId(projectId).SignInWithSessionTokenRequest(signInWithSessionTokenRequest).UnityEnvironment(unityEnvironment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerAuthenticationAPI.SignInWithSessionToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SignInWithSessionToken`: AuthenticationResponse
    fmt.Fprintf(os.Stdout, "Response from `PlayerAuthenticationAPI.SignInWithSessionToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignInWithSessionTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** | This is the Unity Project Id. It is a uuid format.  | 
 **signInWithSessionTokenRequest** | [**SignInWithSessionTokenRequest**](SignInWithSessionTokenRequest.md) | Session Token Authentication request body | 
 **unityEnvironment** | **string** | This is the Environment you want to authorize a player to access. It is the name of the Environment. If this header is not specified, then the default Environment is used. An invalid environment name is not an acceptable input.  | 

### Return type

[**AuthenticationResponse**](AuthenticationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignInWithUsernamePassword

> AuthenticationResponse SignInWithUsernamePassword(ctx).ProjectId(projectId).UsernamePasswordRequest(usernamePasswordRequest).Execute()

Username Password Sign In



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/liu1700/unityauth"
)

func main() {
    projectId := "8bdacc33-6eef-4577-beb0-633c86259f5b" // string | This is the Unity Project Id. It is a uuid format. 
    usernamePasswordRequest := *openapiclient.NewUsernamePasswordRequest() // UsernamePasswordRequest | The Username and Password of the player.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerAuthenticationAPI.SignInWithUsernamePassword(context.Background()).ProjectId(projectId).UsernamePasswordRequest(usernamePasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerAuthenticationAPI.SignInWithUsernamePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SignInWithUsernamePassword`: AuthenticationResponse
    fmt.Fprintf(os.Stdout, "Response from `PlayerAuthenticationAPI.SignInWithUsernamePassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignInWithUsernamePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** | This is the Unity Project Id. It is a uuid format.  | 
 **usernamePasswordRequest** | [**UsernamePasswordRequest**](UsernamePasswordRequest.md) | The Username and Password of the player. | 

### Return type

[**AuthenticationResponse**](AuthenticationResponse.md)

### Authorization

[Client](../README.md#Client)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignUpAnonymously

> AuthenticationResponse SignUpAnonymously(ctx).ProjectId(projectId).SignUpAnonymouslyRequest(signUpAnonymouslyRequest).UnityEnvironment(unityEnvironment).Execute()

Anonymous Sign Up



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/liu1700/unityauth"
)

func main() {
    projectId := "8bdacc33-6eef-4577-beb0-633c86259f5b" // string | This is the Unity Project Id. It is a uuid format. 
    signUpAnonymouslyRequest := *openapiclient.NewSignUpAnonymouslyRequest() // SignUpAnonymouslyRequest | Anonymous Sign Up request body
    unityEnvironment := "production" // string | This is the Environment you want to authorize a player to access. It is the name of the Environment. If this header is not specified, then the default Environment is used. An invalid environment name is not an acceptable input.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerAuthenticationAPI.SignUpAnonymously(context.Background()).ProjectId(projectId).SignUpAnonymouslyRequest(signUpAnonymouslyRequest).UnityEnvironment(unityEnvironment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerAuthenticationAPI.SignUpAnonymously``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SignUpAnonymously`: AuthenticationResponse
    fmt.Fprintf(os.Stdout, "Response from `PlayerAuthenticationAPI.SignUpAnonymously`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignUpAnonymouslyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** | This is the Unity Project Id. It is a uuid format.  | 
 **signUpAnonymouslyRequest** | [**SignUpAnonymouslyRequest**](SignUpAnonymouslyRequest.md) | Anonymous Sign Up request body | 
 **unityEnvironment** | **string** | This is the Environment you want to authorize a player to access. It is the name of the Environment. If this header is not specified, then the default Environment is used. An invalid environment name is not an acceptable input.  | 

### Return type

[**AuthenticationResponse**](AuthenticationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignUpWithUsernamePassword

> AuthenticationResponse SignUpWithUsernamePassword(ctx).ProjectId(projectId).UsernamePasswordRequest(usernamePasswordRequest).BearerAuth(bearerAuth).Execute()

Username Password Sign Up



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/liu1700/unityauth"
)

func main() {
    projectId := "8bdacc33-6eef-4577-beb0-633c86259f5b" // string | This is the Unity Project Id. It is a uuid format. 
    usernamePasswordRequest := *openapiclient.NewUsernamePasswordRequest() // UsernamePasswordRequest | The Username and Password of the player.
    bearerAuth := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c" // string | This is the bearer token for the user authorized to call this API. When this is provided, a user associated with the bearer token will be used instead of creating a new user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerAuthenticationAPI.SignUpWithUsernamePassword(context.Background()).ProjectId(projectId).UsernamePasswordRequest(usernamePasswordRequest).BearerAuth(bearerAuth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerAuthenticationAPI.SignUpWithUsernamePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SignUpWithUsernamePassword`: AuthenticationResponse
    fmt.Fprintf(os.Stdout, "Response from `PlayerAuthenticationAPI.SignUpWithUsernamePassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignUpWithUsernamePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** | This is the Unity Project Id. It is a uuid format.  | 
 **usernamePasswordRequest** | [**UsernamePasswordRequest**](UsernamePasswordRequest.md) | The Username and Password of the player. | 
 **bearerAuth** | **string** | This is the bearer token for the user authorized to call this API. When this is provided, a user associated with the bearer token will be used instead of creating a new user. | 

### Return type

[**AuthenticationResponse**](AuthenticationResponse.md)

### Authorization

[Client](../README.md#Client)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkExternalId

> UnlinkExternalIdResponse UnlinkExternalId(ctx, idProvider).ProjectId(projectId).UnlinkExternalIdRequest(unlinkExternalIdRequest).Execute()

Unlink External Id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/liu1700/unityauth"
)

func main() {
    idProvider := "identity-provider-name" // string | This is the id provider type. 
    projectId := "8bdacc33-6eef-4577-beb0-633c86259f5b" // string | This is the Unity Project Id. It is a uuid format. 
    unlinkExternalIdRequest := *openapiclient.NewUnlinkExternalIdRequest() // UnlinkExternalIdRequest | Unlink external Id request body

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerAuthenticationAPI.UnlinkExternalId(context.Background(), idProvider).ProjectId(projectId).UnlinkExternalIdRequest(unlinkExternalIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerAuthenticationAPI.UnlinkExternalId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnlinkExternalId`: UnlinkExternalIdResponse
    fmt.Fprintf(os.Stdout, "Response from `PlayerAuthenticationAPI.UnlinkExternalId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idProvider** | **string** | This is the id provider type.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkExternalIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectId** | **string** | This is the Unity Project Id. It is a uuid format.  | 
 **unlinkExternalIdRequest** | [**UnlinkExternalIdRequest**](UnlinkExternalIdRequest.md) | Unlink external Id request body | 

### Return type

[**UnlinkExternalIdResponse**](UnlinkExternalIdResponse.md)

### Authorization

[Client](../README.md#Client)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsernamePasswordUpdatePassword

> AuthenticationResponse UsernamePasswordUpdatePassword(ctx).ProjectId(projectId).BearerAuth(bearerAuth).UsernamePasswordPasswordUpdateRequest(usernamePasswordPasswordUpdateRequest).Execute()

Username Password Update Password



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/liu1700/unityauth"
)

func main() {
    projectId := "8bdacc33-6eef-4577-beb0-633c86259f5b" // string | This is the Unity Project Id. It is a uuid format. 
    bearerAuth := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c" // string | This is the bearer token for the user authorized to call this API. 
    usernamePasswordPasswordUpdateRequest := *openapiclient.NewUsernamePasswordPasswordUpdateRequest() // UsernamePasswordPasswordUpdateRequest | The Username, current Password and new Password of the player.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerAuthenticationAPI.UsernamePasswordUpdatePassword(context.Background()).ProjectId(projectId).BearerAuth(bearerAuth).UsernamePasswordPasswordUpdateRequest(usernamePasswordPasswordUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerAuthenticationAPI.UsernamePasswordUpdatePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsernamePasswordUpdatePassword`: AuthenticationResponse
    fmt.Fprintf(os.Stdout, "Response from `PlayerAuthenticationAPI.UsernamePasswordUpdatePassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsernamePasswordUpdatePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** | This is the Unity Project Id. It is a uuid format.  | 
 **bearerAuth** | **string** | This is the bearer token for the user authorized to call this API.  | 
 **usernamePasswordPasswordUpdateRequest** | [**UsernamePasswordPasswordUpdateRequest**](UsernamePasswordPasswordUpdateRequest.md) | The Username, current Password and new Password of the player. | 

### Return type

[**AuthenticationResponse**](AuthenticationResponse.md)

### Authorization

[Client](../README.md#Client)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

