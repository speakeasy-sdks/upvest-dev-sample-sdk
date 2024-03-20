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
    res, err := s.Fees.CreateFeeCollection(ctx, operations.CreateFeeCollectionRequest{
        IdempotencyKey: "ccb07f42-4104-44ad-8e1f-c660bb7b269c",
        Signature: "<value>",
        SignatureInput: "<value>",
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
| `request`                                                                                          | [operations.CreateFeeCollectionRequest](../../pkg/models/operations/createfeecollectionrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.CreateFeeCollectionResponse](../../pkg/models/operations/createfeecollectionresponse.md), error**
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.CreateFeeCollectionError  | 400,401,403,404,406,429,500,503,504 | application/problem+json            |
| sdkerrors.SDKError                  | 4xx-5xx                             | */*                                 |

## ListFeeCollections

Returns a list of fee collections.

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
    res, err := s.Fees.ListFeeCollections(ctx, operations.ListFeeCollectionsRequest{
        AccountGroupID: "288586aa-6cdf-41bc-86dc-1989f6f91027",
        AccountID: "67a4f793-f2bf-4b5f-a5f7-059dafaa6425",
        Signature: "<value>",
        SignatureInput: "<value>",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListFeeCollectionsRequest](../../pkg/models/operations/listfeecollectionsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.ListFeeCollectionsResponse](../../pkg/models/operations/listfeecollectionsresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.ListFeeCollectionsError | 400,401,403,406,429,500,503,504   | application/problem+json          |
| sdkerrors.SDKError                | 4xx-5xx                           | */*                               |

## RetrieveFeeCollection

Returns the fee collection specified by its ID.

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
    res, err := s.Fees.RetrieveFeeCollection(ctx, operations.RetrieveFeeCollectionRequest{
        FeeCollectionID: "abd1b9d9-1afb-4f4c-91e2-b1b47a2dde72",
        Signature: "<value>",
        SignatureInput: "<value>",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.RetrieveFeeCollectionRequest](../../pkg/models/operations/retrievefeecollectionrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.RetrieveFeeCollectionResponse](../../pkg/models/operations/retrievefeecollectionresponse.md), error**
| Error Object                         | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| sdkerrors.RetrieveFeeCollectionError | 401,403,404,406,429,500,503,504      | application/problem+json             |
| sdkerrors.SDKError                   | 4xx-5xx                              | */*                                  |
