# Liquidations
(*Liquidations*)

## Overview

All accounts liquidations related paths.

### Available Operations

* [CancelAccountLiquidation](#cancelaccountliquidation) - Cancel account liquidation
* [CreateAccountLiquidation](#createaccountliquidation) - Create account liquidation request
* [ListAccountsLiquidations](#listaccountsliquidations) - List accounts liquidations
* [RetrieveAccountLiquidation](#retrieveaccountliquidation) - Retrieve account liquidation

## CancelAccountLiquidation

Cancel account liquidation

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	upvestdevsamplesdk "github.com/speakeasy-sdks/upvest-dev-sample-sdk"
	"context"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := upvestdevsamplesdk.New(
        upvestdevsamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Liquidations.CancelAccountLiquidation(ctx, operations.CancelAccountLiquidationRequest{
        AccountID: "cd706758-20cd-4939-8f8e-40e83550ab21",
        AccountLiquidationID: "dd0efae8-3d16-49cb-8eda-a5ad5143e798",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.CancelAccountLiquidationRequest](../../pkg/models/operations/cancelaccountliquidationrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.CancelAccountLiquidationResponse](../../pkg/models/operations/cancelaccountliquidationresponse.md), error**
| Error Object                                                   | Status Code                                                    | Content Type                                                   |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| sdkerrors.CancelAccountLiquidationError                        | 401                                                            | application/problem+json                                       |
| sdkerrors.CancelAccountLiquidationLiquidationsError            | 403                                                            | application/problem+json                                       |
| sdkerrors.CancelAccountLiquidationLiquidationsResponseError    | 404                                                            | application/problem+json                                       |
| sdkerrors.CancelAccountLiquidationLiquidationsResponse406Error | 406                                                            | application/problem+json                                       |
| sdkerrors.CancelAccountLiquidationLiquidationsResponse429Error | 429                                                            | application/problem+json                                       |
| sdkerrors.CancelAccountLiquidationLiquidationsResponse500Error | 500                                                            | application/problem+json                                       |
| sdkerrors.CancelAccountLiquidationLiquidationsResponse503Error | 503                                                            | application/problem+json                                       |
| sdkerrors.CancelAccountLiquidationLiquidationsResponse504Error | 504                                                            | application/problem+json                                       |
| sdkerrors.SDKError                                             | 400-600                                                        | */*                                                            |

## CreateAccountLiquidation

Create account liquidation request

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	upvestdevsamplesdk "github.com/speakeasy-sdks/upvest-dev-sample-sdk"
	"context"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := upvestdevsamplesdk.New(
        upvestdevsamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Liquidations.CreateAccountLiquidation(ctx, operations.CreateAccountLiquidationRequest{
        RequestBody: &operations.CreateAccountLiquidationAccountLiquidationRequest{
            UserID: "b90ecfdb-b70d-41ab-a334-98a3dd3e0dc0",
        },
        AccountID: "99bc53ab-4f38-4564-a664-63304471a03f",
        IdempotencyKey: "ccb07f42-4104-44ad-8e1f-c660bb7b269c",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountLiquidation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.CreateAccountLiquidationRequest](../../pkg/models/operations/createaccountliquidationrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.CreateAccountLiquidationResponse](../../pkg/models/operations/createaccountliquidationresponse.md), error**
| Error Object                                                   | Status Code                                                    | Content Type                                                   |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| sdkerrors.CreateAccountLiquidationError                        | 400                                                            | application/problem+json                                       |
| sdkerrors.CreateAccountLiquidationLiquidationsError            | 401                                                            | application/problem+json                                       |
| sdkerrors.CreateAccountLiquidationLiquidationsResponseError    | 403                                                            | application/problem+json                                       |
| sdkerrors.CreateAccountLiquidationLiquidationsResponse404Error | 404                                                            | application/problem+json                                       |
| sdkerrors.CreateAccountLiquidationLiquidationsResponse406Error | 406                                                            | application/problem+json                                       |
| sdkerrors.CreateAccountLiquidationLiquidationsResponse429Error | 429                                                            | application/problem+json                                       |
| sdkerrors.CreateAccountLiquidationLiquidationsResponse500Error | 500                                                            | application/problem+json                                       |
| sdkerrors.CreateAccountLiquidationLiquidationsResponse503Error | 503                                                            | application/problem+json                                       |
| sdkerrors.CreateAccountLiquidationLiquidationsResponse504Error | 504                                                            | application/problem+json                                       |
| sdkerrors.SDKError                                             | 400-600                                                        | */*                                                            |

## ListAccountsLiquidations

List accounts liquidations

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	upvestdevsamplesdk "github.com/speakeasy-sdks/upvest-dev-sample-sdk"
	"context"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := upvestdevsamplesdk.New(
        upvestdevsamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Liquidations.ListAccountsLiquidations(ctx, operations.ListAccountsLiquidationsRequest{
        AccountID: "d85dad7e-cc61-429e-9b09-5ebf96fc84b7",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PortfoliosOrdersListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListAccountsLiquidationsRequest](../../pkg/models/operations/listaccountsliquidationsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.ListAccountsLiquidationsResponse](../../pkg/models/operations/listaccountsliquidationsresponse.md), error**
| Error Object                                                   | Status Code                                                    | Content Type                                                   |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| sdkerrors.ListAccountsLiquidationsError                        | 400                                                            | application/problem+json                                       |
| sdkerrors.ListAccountsLiquidationsLiquidationsError            | 401                                                            | application/problem+json                                       |
| sdkerrors.ListAccountsLiquidationsLiquidationsResponseError    | 403                                                            | application/problem+json                                       |
| sdkerrors.ListAccountsLiquidationsLiquidationsResponse404Error | 404                                                            | application/problem+json                                       |
| sdkerrors.ListAccountsLiquidationsLiquidationsResponse405Error | 405                                                            | application/problem+json                                       |
| sdkerrors.ListAccountsLiquidationsLiquidationsResponse406Error | 406                                                            | application/problem+json                                       |
| sdkerrors.ListAccountsLiquidationsLiquidationsResponse429Error | 429                                                            | application/problem+json                                       |
| sdkerrors.ListAccountsLiquidationsLiquidationsResponse500Error | 500                                                            | application/problem+json                                       |
| sdkerrors.ListAccountsLiquidationsLiquidationsResponse503Error | 503                                                            | application/problem+json                                       |
| sdkerrors.ListAccountsLiquidationsLiquidationsResponse504Error | 504                                                            | application/problem+json                                       |
| sdkerrors.SDKError                                             | 400-600                                                        | */*                                                            |

## RetrieveAccountLiquidation

Retrieve account liquidation

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	upvestdevsamplesdk "github.com/speakeasy-sdks/upvest-dev-sample-sdk"
	"context"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := upvestdevsamplesdk.New(
        upvestdevsamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Liquidations.RetrieveAccountLiquidation(ctx, operations.RetrieveAccountLiquidationRequest{
        AccountID: "9647ae8a-5755-45e7-b8d8-97c31d5fbeb6",
        AccountLiquidationID: "8b0e85a4-b0ad-48ab-a1e2-30d4bd9e6a9d",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountLiquidation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.RetrieveAccountLiquidationRequest](../../pkg/models/operations/retrieveaccountliquidationrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.RetrieveAccountLiquidationResponse](../../pkg/models/operations/retrieveaccountliquidationresponse.md), error**
| Error Object                                                     | Status Code                                                      | Content Type                                                     |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| sdkerrors.RetrieveAccountLiquidationError                        | 401                                                              | application/problem+json                                         |
| sdkerrors.RetrieveAccountLiquidationLiquidationsError            | 403                                                              | application/problem+json                                         |
| sdkerrors.RetrieveAccountLiquidationLiquidationsResponseError    | 404                                                              | application/problem+json                                         |
| sdkerrors.RetrieveAccountLiquidationLiquidationsResponse406Error | 406                                                              | application/problem+json                                         |
| sdkerrors.RetrieveAccountLiquidationLiquidationsResponse429Error | 429                                                              | application/problem+json                                         |
| sdkerrors.RetrieveAccountLiquidationLiquidationsResponse500Error | 500                                                              | application/problem+json                                         |
| sdkerrors.RetrieveAccountLiquidationLiquidationsResponse503Error | 503                                                              | application/problem+json                                         |
| sdkerrors.RetrieveAccountLiquidationLiquidationsResponse504Error | 504                                                              | application/problem+json                                         |
| sdkerrors.SDKError                                               | 400-600                                                          | */*                                                              |
