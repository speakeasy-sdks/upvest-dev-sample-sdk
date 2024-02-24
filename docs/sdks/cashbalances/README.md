# CashBalances
(*CashBalances*)

## Overview

All cash balance related paths

### Available Operations

* [RetrieveCashBalance](#retrievecashbalance) - Retrieve an account group's cash balance

## RetrieveCashBalance

Retrieve an account group's cash balance

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


    operationSecurity := operations.RetrieveCashBalanceSecurity{
            OauthClientCredentials: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
        }

    ctx := context.Background()
    res, err := s.CashBalances.RetrieveCashBalance(ctx, operations.RetrieveCashBalanceRequest{
        AccountGroupID: "17c70942-e241-4210-962c-ee16861111ec",
        Signature: "<value>",
        SignatureInput: "<value>",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.RetrieveCashBalanceRequest](../../pkg/models/operations/retrievecashbalancerequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.RetrieveCashBalanceSecurity](../../pkg/models/operations/retrievecashbalancesecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


### Response

**[*operations.RetrieveCashBalanceResponse](../../pkg/models/operations/retrievecashbalanceresponse.md), error**
| Error Object                                   | Status Code                                    | Content Type                                   |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| sdkerrors.RetrieveCashBalanceError             | 401,403,404,429,500,503,504                    | application/problem+json                       |
| sdkerrors.RetrieveCashBalanceCashBalancesError | 406                                            | application/problem+json                       |
| sdkerrors.SDKError                             | 4xx-5xx                                        | */*                                            |
