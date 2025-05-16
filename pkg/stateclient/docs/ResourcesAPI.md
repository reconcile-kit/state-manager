# \ResourcesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResource**](ResourcesAPI.md#CreateResource) | **Post** /api/v1/groups/{resource_group}/namespaces/{namespace}/kinds/{kind}/resources | Create a new resource
[**DeleteResource**](ResourcesAPI.md#DeleteResource) | **Delete** /api/v1/groups/{resource_group}/namespaces/{namespace}/kinds/{kind}/resources/{name} | Delete a resource
[**GetResource**](ResourcesAPI.md#GetResource) | **Get** /api/v1/groups/{resource_group}/namespaces/{namespace}/kinds/{kind}/resources/{name} | Get a resource by key
[**ListResources**](ResourcesAPI.md#ListResources) | **Get** /api/v1/resources | List pending resources
[**UpdateResource**](ResourcesAPI.md#UpdateResource) | **Put** /api/v1/groups/{resource_group}/namespaces/{namespace}/kinds/{kind}/resources/{name} | Update a resource
[**UpdateResourceStatus**](ResourcesAPI.md#UpdateResourceStatus) | **Put** /api/v1/groups/{resource_group}/namespaces/{namespace}/kinds/{kind}/resources/{name}/status | Update a resource



## CreateResource

> CreateResource200Response CreateResource(ctx, resourceGroup, kind, namespace).Resource(resource).Execute()

Create a new resource



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	resourceGroup := "resourceGroup_example" // string | Resource Group
	kind := "kind_example" // string | Kind
	namespace := "namespace_example" // string | Namespace
	resource := *openapiclient.NewCreateResourceRequest("Name_example", "ShardId_example") // CreateResourceRequest | Resource details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.CreateResource(context.Background(), resourceGroup, kind, namespace).Resource(resource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.CreateResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateResource`: CreateResource200Response
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.CreateResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceGroup** | **string** | Resource Group | 
**kind** | **string** | Kind | 
**namespace** | **string** | Namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **resource** | [**CreateResourceRequest**](CreateResourceRequest.md) | Resource details | 

### Return type

[**CreateResource200Response**](CreateResource200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResource

> DeleteResource(ctx, resourceGroup, kind, namespace, name).Execute()

Delete a resource



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	resourceGroup := "resourceGroup_example" // string | Resource Group
	kind := "kind_example" // string | Kind
	namespace := "namespace_example" // string | Namespace
	name := "name_example" // string | Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourcesAPI.DeleteResource(context.Background(), resourceGroup, kind, namespace, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.DeleteResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceGroup** | **string** | Resource Group | 
**kind** | **string** | Kind | 
**namespace** | **string** | Namespace | 
**name** | **string** | Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResource

> CreateResource200Response GetResource(ctx, resourceGroup, kind, namespace, name).Execute()

Get a resource by key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	resourceGroup := "resourceGroup_example" // string | Resource Group
	kind := "kind_example" // string | Kind
	namespace := "namespace_example" // string | Namespace
	name := "name_example" // string | Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.GetResource(context.Background(), resourceGroup, kind, namespace, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.GetResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResource`: CreateResource200Response
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.GetResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceGroup** | **string** | Resource Group | 
**kind** | **string** | Kind | 
**namespace** | **string** | Namespace | 
**name** | **string** | Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**CreateResource200Response**](CreateResource200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResources

> CreateResource200Response ListResources(ctx).Kind(kind).Limit(limit).Name(name).Namespace(namespace).Offset(offset).Pending(pending).ResourceGroup(resourceGroup).ShardId(shardId).Execute()

List pending resources



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	kind := "kind_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	name := "name_example" // string |  (optional)
	namespace := "namespace_example" // string |  (optional)
	offset := int32(56) // int32 |  (optional)
	pending := true // bool |  (optional)
	resourceGroup := "resourceGroup_example" // string |  (optional)
	shardId := "shardId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.ListResources(context.Background()).Kind(kind).Limit(limit).Name(name).Namespace(namespace).Offset(offset).Pending(pending).ResourceGroup(resourceGroup).ShardId(shardId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.ListResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListResources`: CreateResource200Response
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.ListResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **string** |  | 
 **limit** | **int32** |  | 
 **name** | **string** |  | 
 **namespace** | **string** |  | 
 **offset** | **int32** |  | 
 **pending** | **bool** |  | 
 **resourceGroup** | **string** |  | 
 **shardId** | **string** |  | 

### Return type

[**CreateResource200Response**](CreateResource200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResource

> CreateResource200Response UpdateResource(ctx, resourceGroup, kind, namespace, name).Resource(resource).Execute()

Update a resource



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	resourceGroup := "resourceGroup_example" // string | Resource Group
	kind := "kind_example" // string | Kind
	namespace := "namespace_example" // string | Namespace
	name := "name_example" // string | Name
	resource := *openapiclient.NewUpdateResourceRequest("ShardId_example") // UpdateResourceRequest | Resource details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.UpdateResource(context.Background(), resourceGroup, kind, namespace, name).Resource(resource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.UpdateResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateResource`: CreateResource200Response
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.UpdateResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceGroup** | **string** | Resource Group | 
**kind** | **string** | Kind | 
**namespace** | **string** | Namespace | 
**name** | **string** | Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **resource** | [**UpdateResourceRequest**](UpdateResourceRequest.md) | Resource details | 

### Return type

[**CreateResource200Response**](CreateResource200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResourceStatus

> CreateResource200Response UpdateResourceStatus(ctx, resourceGroup, kind, namespace, name).Resource(resource).Execute()

Update a resource



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	resourceGroup := "resourceGroup_example" // string | Resource Group
	kind := "kind_example" // string | Kind
	namespace := "namespace_example" // string | Namespace
	name := "name_example" // string | Name
	resource := *openapiclient.NewUpdateResourceStatusRequest("ShardId_example") // UpdateResourceStatusRequest | Resource details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.UpdateResourceStatus(context.Background(), resourceGroup, kind, namespace, name).Resource(resource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.UpdateResourceStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateResourceStatus`: CreateResource200Response
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.UpdateResourceStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceGroup** | **string** | Resource Group | 
**kind** | **string** | Kind | 
**namespace** | **string** | Namespace | 
**name** | **string** | Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResourceStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **resource** | [**UpdateResourceStatusRequest**](UpdateResourceStatusRequest.md) | Resource details | 

### Return type

[**CreateResource200Response**](CreateResource200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

