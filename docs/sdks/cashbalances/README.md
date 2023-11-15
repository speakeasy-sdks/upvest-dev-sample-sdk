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
	"context"
	"log"
	upvestdevsamplesdk "github.com/speakeasy-sdks/upvest-dev-sample-sdk"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
)

func main() {
    s := upvestdevsamplesdk.New()

    ctx := context.Background()
    res, err := s.CashBalances.RetrieveCashBalance(ctx, operations.RetrieveCashBalanceRequest{
        AccountGroupID: "17c70942-e241-4210-962c-ee16861111ec",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.RetrieveCashBalanceRequest](../../pkg/models/operations/retrievecashbalancerequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.RetrieveCashBalanceResponse](../../pkg/models/operations/retrievecashbalanceresponse.md), error**
| Error Object                                              | Status Code                                               | Content Type                                              |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| sdkerrors.RetrieveCashBalanceError                        | 401                                                       | application/problem+json                                  |
| sdkerrors.RetrieveCashBalanceCashBalancesError            | 403                                                       | application/problem+json                                  |
| sdkerrors.RetrieveCashBalanceCashBalancesResponseError    | 404                                                       | application/problem+json                                  |
| sdkerrors.RetrieveCashBalanceCashBalancesResponse406Error | 406                                                       | application/problem+json                                  |
| sdkerrors.RetrieveCashBalanceCashBalancesResponse429Error | 429                                                       | application/problem+json                                  |
| sdkerrors.RetrieveCashBalanceCashBalancesResponse500Error | 500                                                       | application/problem+json                                  |
| sdkerrors.RetrieveCashBalanceCashBalancesResponse503Error | 503                                                       | application/problem+json                                  |
| sdkerrors.RetrieveCashBalanceCashBalancesResponse504Error | 504                                                       | application/problem+json                                  |
| sdkerrors.SDKError                                        | 400-600                                                   | */*                                                       |
