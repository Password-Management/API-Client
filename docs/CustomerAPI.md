# \CustomerAPI

All URIs are relative to *http://localhost:8001*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerGet**](CustomerAPI.md#CustomerGet) | **Get** /customer | Get customer details using the MasterId
[**CustomerPost**](CustomerAPI.md#CustomerPost) | **Post** /customer | Create a new customer in the database



## CustomerGet

> CustomerGet200Response CustomerGet(ctx).MasterId(masterId).Execute()

Get customer details using the MasterId

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Password-Management/API-Client"
)

func main() {
	masterId := "masterId_example" // string | The Master ID of the customer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CustomerGet(context.Background()).MasterId(masterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.CustomerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerGet`: CustomerGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.CustomerGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **masterId** | **string** | The Master ID of the customer | 

### Return type

[**CustomerGet200Response**](CustomerGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerPost

> CustomerPost200Response CustomerPost(ctx).CustomerPostRequest(customerPostRequest).Execute()

Create a new customer in the database

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Password-Management/API-Client"
)

func main() {
	customerPostRequest := *openapiclient.NewCustomerPostRequest("Email_example", "Name_example", "Plan_example", "Algorithm_example", "Platform_example") // CustomerPostRequest | Customer details

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
 **customerPostRequest** | [**CustomerPostRequest**](CustomerPostRequest.md) | Customer details | 

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

