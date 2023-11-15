# Transactions
(*Transactions*)

## Overview

All transactions related paths.

### Available Operations

* [ListCashTransactions](#listcashtransactions) - List cash transactions
* [ListSecuritiesTransactions](#listsecuritiestransactions) - List securities transactions

## ListCashTransactions

List cash transactions

### Example Usage

```go
package main

import(
	"context"
	"log"
	upvestdevsamplesdk "github.com/speakeasy-sdks/upvest-dev-sample-sdk"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/types"
)

func main() {
    s := upvestdevsamplesdk.New()

    ctx := context.Background()
    res, err := s.Transactions.ListCashTransactions(ctx, operations.ListCashTransactionsRequest{
        EndDate: types.MustDateFromString("2023-01-11"),
        Signature: "string",
        SignatureInput: "string",
        StartDate: types.MustDateFromString("2023-01-03"),
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONCashTransactionListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListCashTransactionsRequest](../../pkg/models/operations/listcashtransactionsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.ListCashTransactionsResponse](../../pkg/models/operations/listcashtransactionsresponse.md), error**
| Error Object                                               | Status Code                                                | Content Type                                               |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| sdkerrors.ListCashTransactionsError                        | 400                                                        | application/problem+json                                   |
| sdkerrors.ListCashTransactionsTransactionsError            | 401                                                        | application/problem+json                                   |
| sdkerrors.ListCashTransactionsTransactionsResponseError    | 403                                                        | application/problem+json                                   |
| sdkerrors.ListCashTransactionsTransactionsResponse404Error | 404                                                        | application/problem+json                                   |
| sdkerrors.ListCashTransactionsTransactionsResponse405Error | 405                                                        | application/problem+json                                   |
| sdkerrors.ListCashTransactionsTransactionsResponse406Error | 406                                                        | application/problem+json                                   |
| sdkerrors.ListCashTransactionsTransactionsResponse429Error | 429                                                        | application/problem+json                                   |
| sdkerrors.ListCashTransactionsTransactionsResponse500Error | 500                                                        | application/problem+json                                   |
| sdkerrors.ListCashTransactionsTransactionsResponse503Error | 503                                                        | application/problem+json                                   |
| sdkerrors.ListCashTransactionsTransactionsResponse504Error | 504                                                        | application/problem+json                                   |
| sdkerrors.SDKError                                         | 400-600                                                    | */*                                                        |

## ListSecuritiesTransactions

List securities transactions

### Example Usage

```go
package main

import(
	"context"
	"log"
	upvestdevsamplesdk "github.com/speakeasy-sdks/upvest-dev-sample-sdk"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/types"
)

func main() {
    s := upvestdevsamplesdk.New()

    ctx := context.Background()
    res, err := s.Transactions.ListSecuritiesTransactions(ctx, operations.ListSecuritiesTransactionsRequest{
        EndDate: types.MustDateFromString("2023-01-11"),
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONSecurityTransactionListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.ListSecuritiesTransactionsRequest](../../pkg/models/operations/listsecuritiestransactionsrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.ListSecuritiesTransactionsResponse](../../pkg/models/operations/listsecuritiestransactionsresponse.md), error**
| Error Object                                                     | Status Code                                                      | Content Type                                                     |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| sdkerrors.ListSecuritiesTransactionsError                        | 400                                                              | application/problem+json                                         |
| sdkerrors.ListSecuritiesTransactionsTransactionsError            | 401                                                              | application/problem+json                                         |
| sdkerrors.ListSecuritiesTransactionsTransactionsResponseError    | 403                                                              | application/problem+json                                         |
| sdkerrors.ListSecuritiesTransactionsTransactionsResponse404Error | 404                                                              | application/problem+json                                         |
| sdkerrors.ListSecuritiesTransactionsTransactionsResponse405Error | 405                                                              | application/problem+json                                         |
| sdkerrors.ListSecuritiesTransactionsTransactionsResponse406Error | 406                                                              | application/problem+json                                         |
| sdkerrors.ListSecuritiesTransactionsTransactionsResponse429Error | 429                                                              | application/problem+json                                         |
| sdkerrors.ListSecuritiesTransactionsTransactionsResponse500Error | 500                                                              | application/problem+json                                         |
| sdkerrors.ListSecuritiesTransactionsTransactionsResponse503Error | 503                                                              | application/problem+json                                         |
| sdkerrors.ListSecuritiesTransactionsTransactionsResponse504Error | 504                                                              | application/problem+json                                         |
| sdkerrors.SDKError                                               | 400-600                                                          | */*                                                              |
