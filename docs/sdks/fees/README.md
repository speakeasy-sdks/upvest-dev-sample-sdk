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
    s := upvestdevsamplesdk.New(
        upvestdevsamplesdk.WithSecurity("YOUR_TOKEN"),
    )

    ctx := context.Background()
    res, err := s.Fees.CreateFeeCollection(ctx, operations.CreateFeeCollectionRequest{
        RequestBody: &operations.CreateFeeCollectionFeeCollectionCreateRequest{
            AccountID: "f7894af9-5b70-47e1-ab25-69b76e33b134",
            CollectionAmount: "National",
            PeriodEnd: types.MustDateFromString("2022-02-19"),
            PeriodStart: types.MustDateFromString("2023-07-28"),
            Type: operations.CreateFeeCollectionFeeCollectionCreateRequestTypeServiceFee,
        },
        IdempotencyKey: "ccb07f42-4104-44ad-8e1f-c660bb7b269c",
        Signature: "Alaska",
        SignatureInput: "USB",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FeeCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateFeeCollectionRequest](../../models/operations/createfeecollectionrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.CreateFeeCollectionResponse](../../models/operations/createfeecollectionresponse.md), error**


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
    s := upvestdevsamplesdk.New(
        upvestdevsamplesdk.WithSecurity("YOUR_TOKEN"),
    )

    ctx := context.Background()
    res, err := s.Fees.ListFeeCollections(ctx, operations.ListFeeCollectionsRequest{
        AccountGroupID: "288586aa-6cdf-41bc-86dc-1989f6f91027",
        AccountID: "67a4f793-f2bf-4b5f-a5f7-059dafaa6425",
        Signature: "Pants",
        SignatureInput: "blockchains",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FeeCollectionListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListFeeCollectionsRequest](../../models/operations/listfeecollectionsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.ListFeeCollectionsResponse](../../models/operations/listfeecollectionsresponse.md), error**


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
    s := upvestdevsamplesdk.New(
        upvestdevsamplesdk.WithSecurity("YOUR_TOKEN"),
    )

    ctx := context.Background()
    res, err := s.Fees.RetrieveFeeCollection(ctx, operations.RetrieveFeeCollectionRequest{
        FeeCollectionID: "abd1b9d9-1afb-4f4c-91e2-b1b47a2dde72",
        Signature: "toward",
        SignatureInput: "Branding",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FeeCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.RetrieveFeeCollectionRequest](../../models/operations/retrievefeecollectionrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.RetrieveFeeCollectionResponse](../../models/operations/retrievefeecollectionresponse.md), error**

