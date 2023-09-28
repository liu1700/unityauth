# \PlayerAuthenticationAdminAPI

All URIs are relative to *https://services.api.unity.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePlayer**](PlayerAuthenticationAdminAPI.md#DeletePlayer) | **Delete** /player-identity/v1/projects/{projectId}/users/{playerId} | Delete Player
[**DisablePlayer**](PlayerAuthenticationAdminAPI.md#DisablePlayer) | **Post** /player-identity/v1/projects/{projectId}/users/{playerId}/disable | Disable Player
[**EnablePlayer**](PlayerAuthenticationAdminAPI.md#EnablePlayer) | **Post** /player-identity/v1/projects/{projectId}/users/{playerId}/enable | Enable Player
[**GetPlayer**](PlayerAuthenticationAdminAPI.md#GetPlayer) | **Get** /player-identity/v1/projects/{projectId}/users/{playerId} | Get Player
[**ListPlayers**](PlayerAuthenticationAdminAPI.md#ListPlayers) | **Get** /player-identity/v1/projects/{projectId}/users | List Players



## DeletePlayer

> map[string]interface{} DeletePlayer(ctx, playerId, projectId).Execute()

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
    playerId := "99i9ju8juh" // string | This is the player Id. 
    projectId := "8bdacc33-6eef-4577-beb0-633c86259f5b" // string | Id of the project.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerAuthenticationAdminAPI.DeletePlayer(context.Background(), playerId, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerAuthenticationAdminAPI.DeletePlayer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePlayer`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PlayerAuthenticationAdminAPI.DeletePlayer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playerId** | **string** | This is the player Id.  | 
**projectId** | **string** | Id of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlayerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[Admin](../README.md#Admin)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisablePlayer

> PlayerAuthPlayerProjectResponse DisablePlayer(ctx, playerId, projectId).Execute()

Disable Player



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
    playerId := "99i9ju8juh" // string | This is the player Id. 
    projectId := "8bdacc33-6eef-4577-beb0-633c86259f5b" // string | Id of the project.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerAuthenticationAdminAPI.DisablePlayer(context.Background(), playerId, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerAuthenticationAdminAPI.DisablePlayer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisablePlayer`: PlayerAuthPlayerProjectResponse
    fmt.Fprintf(os.Stdout, "Response from `PlayerAuthenticationAdminAPI.DisablePlayer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playerId** | **string** | This is the player Id.  | 
**projectId** | **string** | Id of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisablePlayerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PlayerAuthPlayerProjectResponse**](PlayerAuthPlayerProjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnablePlayer

> PlayerAuthPlayerProjectResponse EnablePlayer(ctx, playerId, projectId).Execute()

Enable Player



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
    playerId := "99i9ju8juh" // string | This is the player Id. 
    projectId := "8bdacc33-6eef-4577-beb0-633c86259f5b" // string | Id of the project.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerAuthenticationAdminAPI.EnablePlayer(context.Background(), playerId, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerAuthenticationAdminAPI.EnablePlayer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnablePlayer`: PlayerAuthPlayerProjectResponse
    fmt.Fprintf(os.Stdout, "Response from `PlayerAuthenticationAdminAPI.EnablePlayer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playerId** | **string** | This is the player Id.  | 
**projectId** | **string** | Id of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnablePlayerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PlayerAuthPlayerProjectResponse**](PlayerAuthPlayerProjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayer

> PlayerAuthPlayerProjectResponse GetPlayer(ctx, playerId, projectId).Execute()

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
    playerId := "99i9ju8juh" // string | This is the player Id. 
    projectId := "8bdacc33-6eef-4577-beb0-633c86259f5b" // string | Id of the project.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerAuthenticationAdminAPI.GetPlayer(context.Background(), playerId, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerAuthenticationAdminAPI.GetPlayer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlayer`: PlayerAuthPlayerProjectResponse
    fmt.Fprintf(os.Stdout, "Response from `PlayerAuthenticationAdminAPI.GetPlayer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playerId** | **string** | This is the player Id.  | 
**projectId** | **string** | Id of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PlayerAuthPlayerProjectResponse**](PlayerAuthPlayerProjectResponse.md)

### Authorization

[Admin](../README.md#Admin)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPlayers

> PlayerAuthListProjectUserResponse ListPlayers(ctx, projectId).Limit(limit).Page(page).Execute()

List Players



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
    projectId := "8bdacc33-6eef-4577-beb0-633c86259f5b" // string | Id of the project.
    limit := int32(56) // int32 | The number of players to return. (optional)
    page := "page_example" // string | (Optional) The next page token. If not set, returns players without any offset. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerAuthenticationAdminAPI.ListPlayers(context.Background(), projectId).Limit(limit).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerAuthenticationAdminAPI.ListPlayers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPlayers`: PlayerAuthListProjectUserResponse
    fmt.Fprintf(os.Stdout, "Response from `PlayerAuthenticationAdminAPI.ListPlayers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Id of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPlayersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The number of players to return. | 
 **page** | **string** | (Optional) The next page token. If not set, returns players without any offset. | 

### Return type

[**PlayerAuthListProjectUserResponse**](PlayerAuthListProjectUserResponse.md)

### Authorization

[Admin](../README.md#Admin)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

