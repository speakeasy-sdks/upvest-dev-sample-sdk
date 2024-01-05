# Valuations
(*Valuations*)

## Overview

All valuations related paths.

### Available Operations

* [GetAccountValuation](#getaccountvaluation) - Get current valuation for an account
* [ListAccountValuationHistory](#listaccountvaluationhistory) - List valuation history for an account

## GetAccountValuation

Get current valuation for an account

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Valuations.GetAccountValuation(ctx, operations.GetAccountValuationRequest{
        AccountID: "4767938f-4376-4441-a2b6-0a26456b8e4f",
        PriceQuality: operations.PriceQualityEod,
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountValuation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetAccountValuationRequest](../../pkg/models/operations/getaccountvaluationrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.GetAccountValuationResponse](../../pkg/models/operations/getaccountvaluationresponse.md), error**
| Error Object                                            | Status Code                                             | Content Type                                            |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| sdkerrors.GetAccountValuationError                      | 400                                                     | application/problem+json                                |
| sdkerrors.GetAccountValuationValuationsError            | 401                                                     | application/problem+json                                |
| sdkerrors.GetAccountValuationValuationsResponseError    | 403                                                     | application/problem+json                                |
| sdkerrors.GetAccountValuationValuationsResponse404Error | 404                                                     | application/problem+json                                |
| sdkerrors.GetAccountValuationValuationsResponse406Error | 406                                                     | application/problem+json                                |
| sdkerrors.GetAccountValuationValuationsResponse429Error | 429                                                     | application/problem+json                                |
| sdkerrors.GetAccountValuationValuationsResponse500Error | 500                                                     | application/problem+json                                |
| sdkerrors.GetAccountValuationValuationsResponse503Error | 503                                                     | application/problem+json                                |
| sdkerrors.GetAccountValuationValuationsResponse504Error | 504                                                     | application/problem+json                                |
| sdkerrors.SDKError                                      | 4xx-5xx                                                 | */*                                                     |

## ListAccountValuationHistory

List valuation history for an account

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Valuations.ListAccountValuationHistory(ctx, operations.ListAccountValuationHistoryRequest{
        AccountID: "cbc89a08-4aea-491e-b3e5-762fdd9e8431",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountValuationListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.ListAccountValuationHistoryRequest](../../pkg/models/operations/listaccountvaluationhistoryrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.ListAccountValuationHistoryResponse](../../pkg/models/operations/listaccountvaluationhistoryresponse.md), error**
| Error Object                                                    | Status Code                                                     | Content Type                                                    |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| sdkerrors.ListAccountValuationHistoryError                      | 400                                                             | application/problem+json                                        |
| sdkerrors.ListAccountValuationHistoryValuationsError            | 401                                                             | application/problem+json                                        |
| sdkerrors.ListAccountValuationHistoryValuationsResponseError    | 403                                                             | application/problem+json                                        |
| sdkerrors.ListAccountValuationHistoryValuationsResponse404Error | 404                                                             | application/problem+json                                        |
| sdkerrors.ListAccountValuationHistoryValuationsResponse405Error | 405                                                             | application/problem+json                                        |
| sdkerrors.ListAccountValuationHistoryValuationsResponse406Error | 406                                                             | application/problem+json                                        |
| sdkerrors.ListAccountValuationHistoryValuationsResponse429Error | 429                                                             | application/problem+json                                        |
| sdkerrors.ListAccountValuationHistoryValuationsResponse500Error | 500                                                             | application/problem+json                                        |
| sdkerrors.ListAccountValuationHistoryValuationsResponse503Error | 503                                                             | application/problem+json                                        |
| sdkerrors.ListAccountValuationHistoryValuationsResponse504Error | 504                                                             | application/problem+json                                        |
| sdkerrors.SDKError                                              | 4xx-5xx                                                         | */*                                                             |
