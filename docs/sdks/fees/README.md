# Fees
(*Fees*)

## Overview

All fees related paths.

### Available Operations

* [CreateFeeCollection](#createfeecollection) - Create a fee collection
* [ListFeeCollections](#listfeecollections) - Get fee collections
* [RetrieveFeeCollection](#retrievefeecollection) - Get a fee collection by ID

## CreateFeeCollection

Creates a fee collection for pre-calculated fee amounts.

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
    res, err := s.Fees.CreateFeeCollection(ctx, operations.CreateFeeCollectionRequest{
        RequestBody: &operations.CreateFeeCollectionFeeCollectionCreateRequest{
            AccountID: "f7894af9-5b70-47e1-ab25-69b76e33b134",
            CollectionAmount: "string",
            PeriodEnd: types.MustDateFromString("2022-12-20"),
            PeriodStart: types.MustDateFromString("2022-02-15"),
            Type: operations.CreateFeeCollectionTypeServiceFee,
        },
        IdempotencyKey: "ccb07f42-4104-44ad-8e1f-c660bb7b269c",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONFeeCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateFeeCollectionRequest](../../pkg/models/operations/createfeecollectionrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.CreateFeeCollectionResponse](../../pkg/models/operations/createfeecollectionresponse.md), error**
| Error Object                                      | Status Code                                       | Content Type                                      |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| sdkerrors.CreateFeeCollectionError                | 400                                               | application/problem+json                          |
| sdkerrors.CreateFeeCollectionFeesError            | 401                                               | application/problem+json                          |
| sdkerrors.CreateFeeCollectionFeesResponseError    | 403                                               | application/problem+json                          |
| sdkerrors.CreateFeeCollectionFeesResponse404Error | 404                                               | application/problem+json                          |
| sdkerrors.CreateFeeCollectionFeesResponse406Error | 406                                               | application/problem+json                          |
| sdkerrors.CreateFeeCollectionFeesResponse429Error | 429                                               | application/problem+json                          |
| sdkerrors.CreateFeeCollectionFeesResponse500Error | 500                                               | application/problem+json                          |
| sdkerrors.CreateFeeCollectionFeesResponse503Error | 503                                               | application/problem+json                          |
| sdkerrors.CreateFeeCollectionFeesResponse504Error | 504                                               | application/problem+json                          |
| sdkerrors.SDKError                                | 400-600                                           | */*                                               |

## ListFeeCollections

Returns a list of fee collections.

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
    res, err := s.Fees.ListFeeCollections(ctx, operations.ListFeeCollectionsRequest{
        AccountGroupID: "288586aa-6cdf-41bc-86dc-1989f6f91027",
        AccountID: "67a4f793-f2bf-4b5f-a5f7-059dafaa6425",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONFeeCollectionListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListFeeCollectionsRequest](../../pkg/models/operations/listfeecollectionsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.ListFeeCollectionsResponse](../../pkg/models/operations/listfeecollectionsresponse.md), error**
| Error Object                                     | Status Code                                      | Content Type                                     |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| sdkerrors.ListFeeCollectionsError                | 400                                              | application/problem+json                         |
| sdkerrors.ListFeeCollectionsFeesError            | 401                                              | application/problem+json                         |
| sdkerrors.ListFeeCollectionsFeesResponseError    | 403                                              | application/problem+json                         |
| sdkerrors.ListFeeCollectionsFeesResponse406Error | 406                                              | application/problem+json                         |
| sdkerrors.ListFeeCollectionsFeesResponse429Error | 429                                              | application/problem+json                         |
| sdkerrors.ListFeeCollectionsFeesResponse500Error | 500                                              | application/problem+json                         |
| sdkerrors.ListFeeCollectionsFeesResponse503Error | 503                                              | application/problem+json                         |
| sdkerrors.ListFeeCollectionsFeesResponse504Error | 504                                              | application/problem+json                         |
| sdkerrors.SDKError                               | 400-600                                          | */*                                              |

## RetrieveFeeCollection

Returns the fee collection specified by its ID.

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
    res, err := s.Fees.RetrieveFeeCollection(ctx, operations.RetrieveFeeCollectionRequest{
        FeeCollectionID: "abd1b9d9-1afb-4f4c-91e2-b1b47a2dde72",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONFeeCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.RetrieveFeeCollectionRequest](../../pkg/models/operations/retrievefeecollectionrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.RetrieveFeeCollectionResponse](../../pkg/models/operations/retrievefeecollectionresponse.md), error**
| Error Object                                        | Status Code                                         | Content Type                                        |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| sdkerrors.RetrieveFeeCollectionError                | 401                                                 | application/problem+json                            |
| sdkerrors.RetrieveFeeCollectionFeesError            | 403                                                 | application/problem+json                            |
| sdkerrors.RetrieveFeeCollectionFeesResponseError    | 404                                                 | application/problem+json                            |
| sdkerrors.RetrieveFeeCollectionFeesResponse406Error | 406                                                 | application/problem+json                            |
| sdkerrors.RetrieveFeeCollectionFeesResponse429Error | 429                                                 | application/problem+json                            |
| sdkerrors.RetrieveFeeCollectionFeesResponse500Error | 500                                                 | application/problem+json                            |
| sdkerrors.RetrieveFeeCollectionFeesResponse503Error | 503                                                 | application/problem+json                            |
| sdkerrors.RetrieveFeeCollectionFeesResponse504Error | 504                                                 | application/problem+json                            |
| sdkerrors.SDKError                                  | 400-600                                             | */*                                                 |
