# VirtualCashBalances
(*VirtualCashBalances*)

## Overview

All virtual cash balances related paths

### Available Operations

* [CancelVirtualCashDecrease](#cancelvirtualcashdecrease) - Cancel virtual cash decrease by ID
* [CreateVirtualCashDecrease](#createvirtualcashdecrease) - Trigger a virtual cash decrease
* [CreateVirtualCashIncrease](#createvirtualcashincrease) - Trigger a virtual cash increase

## CancelVirtualCashDecrease

Cancels a virtual cash decrease specified by its ID. It is only possible to cancel a virtual cash decrease if it has the status `ISSUED` or `QUEUED`.

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
        upvestdevsamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.VirtualCashBalances.CancelVirtualCashDecrease(ctx, operations.CancelVirtualCashDecreaseRequest{
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        VirtualCashDecreaseID: "29770b4c-40f4-4e9d-8aa1-a63d51b4005a",
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.CancelVirtualCashDecreaseRequest](../../models/operations/cancelvirtualcashdecreaserequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.CancelVirtualCashDecreaseResponse](../../models/operations/cancelvirtualcashdecreaseresponse.md), error**


## CreateVirtualCashDecrease

Trigger a virtual cash decrease

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
        upvestdevsamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.VirtualCashBalances.CreateVirtualCashDecrease(ctx, operations.CreateVirtualCashDecreaseRequest{
        RequestBody: &operations.CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCreateRequest{
            AccountGroupID: "2e6f782f-c6dc-4f9d-91c1-313da4f8ef04",
            Amount: "436.77",
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

    if res.VirtualCashBalanceVirtualCashDecrease != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.CreateVirtualCashDecreaseRequest](../../models/operations/createvirtualcashdecreaserequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.CreateVirtualCashDecreaseResponse](../../models/operations/createvirtualcashdecreaseresponse.md), error**


## CreateVirtualCashIncrease

Trigger a virtual cash increase

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
        upvestdevsamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.VirtualCashBalances.CreateVirtualCashIncrease(ctx, operations.CreateVirtualCashIncreaseRequest{
        RequestBody: &operations.CreateVirtualCashIncreaseVirtualCashBalanceVirtualCashIncreaseCreateRequest{
            AccountGroupID: "ecd9a7c7-c554-4819-87b7-feb3aef3e971",
            Amount: "856.18",
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

    if res.VirtualCashBalanceVirtualCashIncrease != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.CreateVirtualCashIncreaseRequest](../../models/operations/createvirtualcashincreaserequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.CreateVirtualCashIncreaseResponse](../../models/operations/createvirtualcashincreaseresponse.md), error**

