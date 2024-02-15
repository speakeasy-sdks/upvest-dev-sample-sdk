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
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	upvestdevsamplesdk "github.com/speakeasy-sdks/upvest-dev-sample-sdk"
	"context"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/types"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := upvestdevsamplesdk.New(
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Transactions.ListCashTransactions(ctx, operations.ListCashTransactionsRequest{
        EndDate: types.MustNewDateFromString("2023-01-11T00:00:00Z"),
        Signature: "<value>",
        SignatureInput: "<value>",
        StartDate: types.MustNewDateFromString("2023-01-03T00:00:00Z"),
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CashTransactionListResponse != nil {
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
| Error Object                                    | Status Code                                     | Content Type                                    |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| sdkerrors.ListCashTransactionsError             | 400,401,403,404,406,429,500,503,504             | application/problem+json                        |
| sdkerrors.ListCashTransactionsTransactionsError | 405                                             | application/problem+json                        |
| sdkerrors.SDKError                              | 4xx-5xx                                         | */*                                             |

## ListSecuritiesTransactions

List securities transactions

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	upvestdevsamplesdk "github.com/speakeasy-sdks/upvest-dev-sample-sdk"
	"context"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/types"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := upvestdevsamplesdk.New(
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Transactions.ListSecuritiesTransactions(ctx, operations.ListSecuritiesTransactionsRequest{
        EndDate: types.MustNewDateFromString("2023-01-11T00:00:00Z"),
        Signature: "<value>",
        SignatureInput: "<value>",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SecurityTransactionListResponse != nil {
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
| Error Object                                          | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| sdkerrors.ListSecuritiesTransactionsError             | 400,401,403,404,406,429,500,503,504                   | application/problem+json                              |
| sdkerrors.ListSecuritiesTransactionsTransactionsError | 405                                                   | application/problem+json                              |
| sdkerrors.SDKError                                    | 4xx-5xx                                               | */*                                                   |
