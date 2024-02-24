# Returns
(*Returns*)

## Overview

All accounts returns related paths.

### Available Operations

* [ListAccountReturns](#listaccountreturns) - List account returns

## ListAccountReturns

List account returns

### Example Usage

```go
package main

import(
	upvestdevsamplesdk "github.com/speakeasy-sdks/upvest-dev-sample-sdk"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
	"context"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"log"
)

func main() {
    s := upvestdevsamplesdk.New()


    operationSecurity := operations.ListAccountReturnsSecurity{
            OauthClientCredentials: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
        }

    ctx := context.Background()
    res, err := s.Returns.ListAccountReturns(ctx, operations.ListAccountReturnsRequest{
        AccountID: "5b95c1fa-d06f-4a63-8a4b-4d67218cd5f1",
        Signature: "<value>",
        SignatureInput: "<value>",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountReturnListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListAccountReturnsRequest](../../pkg/models/operations/listaccountreturnsrequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.ListAccountReturnsSecurity](../../pkg/models/operations/listaccountreturnssecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.ListAccountReturnsResponse](../../pkg/models/operations/listaccountreturnsresponse.md), error**
| Error Object                             | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| sdkerrors.ListAccountReturnsError        | 400,401,403,404,406,429,500,503,504      | application/problem+json                 |
| sdkerrors.ListAccountReturnsReturnsError | 405                                      | application/problem+json                 |
| sdkerrors.SDKError                       | 4xx-5xx                                  | */*                                      |
