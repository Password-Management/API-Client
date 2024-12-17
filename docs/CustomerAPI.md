# \CustomerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerPost**](CustomerAPI.md#CustomerPost) | **Post** /customer | Create a new custmer in the database
[**GetResultsGet**](CustomerAPI.md#GetResultsGet) | **Get** /get-results | API to get counts of votes/ Winner of the election



## CustomerPost

> CustomerPost200Response CustomerPost(ctx).CustomerPostRequest(customerPostRequest).Execute()

Create a new custmer in the database

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
	customerPostRequest := *openapiclient.NewCustomerPostRequest("Email_example", "Name_example", "Plan_example", "Algorithm_example", "Platform_example") // CustomerPostRequest | User login credentials

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CustomerPost(context.Background()).CustomerPostRequest(customerPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.CustomerPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerPost`: CustomerPost200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.CustomerPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerPostRequest** | [**CustomerPostRequest**](CustomerPostRequest.md) | User login credentials | 

### Return type

[**CustomerPost200Response**](CustomerPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResultsGet

> GetResultsGet200Response GetResultsGet(ctx).MasterId(masterId).Execute()

API to get counts of votes/ Winner of the election

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
	masterId := "masterId_example" // string | The search keyword

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.GetResultsGet(context.Background()).MasterId(masterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.GetResultsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResultsGet`: GetResultsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.GetResultsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetResultsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **masterId** | **string** | The search keyword | 

### Return type

[**GetResultsGet200Response**](GetResultsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

