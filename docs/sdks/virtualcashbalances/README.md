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
    res, err := s.VirtualCashBalances.CancelVirtualCashDecrease(ctx, operations.CancelVirtualCashDecreaseRequest{
        Signature: "<value>",
        SignatureInput: "<value>",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        VirtualCashDecreaseID: "29770b4c-40f4-4e9d-8aa1-a63d51b4005a",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.CancelVirtualCashDecreaseRequest](../../pkg/models/operations/cancelvirtualcashdecreaserequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.CancelVirtualCashDecreaseResponse](../../pkg/models/operations/cancelvirtualcashdecreaseresponse.md), error**
| Error Object                                                | Status Code                                                 | Content Type                                                |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| sdkerrors.CancelVirtualCashDecreaseError                    | 401,403,404,429,500,503,504                                 | application/problem+json                                    |
| sdkerrors.CancelVirtualCashDecreaseVirtualCashBalancesError | 422                                                         | application/problem+json                                    |
| sdkerrors.SDKError                                          | 4xx-5xx                                                     | */*                                                         |

## CreateVirtualCashDecrease

Trigger a virtual cash decrease

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
    res, err := s.VirtualCashBalances.CreateVirtualCashDecrease(ctx, operations.CreateVirtualCashDecreaseRequest{
        IdempotencyKey: "ccb07f42-4104-44ad-8e1f-c660bb7b269c",
        Signature: "<value>",
        SignatureInput: "<value>",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.CreateVirtualCashDecreaseRequest](../../pkg/models/operations/createvirtualcashdecreaserequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.CreateVirtualCashDecreaseResponse](../../pkg/models/operations/createvirtualcashdecreaseresponse.md), error**
| Error Object                             | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| sdkerrors.CreateVirtualCashDecreaseError | 400,401,403,404,406,429,500,503,504      | application/problem+json                 |
| sdkerrors.SDKError                       | 4xx-5xx                                  | */*                                      |

## CreateVirtualCashIncrease

Trigger a virtual cash increase

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
    res, err := s.VirtualCashBalances.CreateVirtualCashIncrease(ctx, operations.CreateVirtualCashIncreaseRequest{
        IdempotencyKey: "ccb07f42-4104-44ad-8e1f-c660bb7b269c",
        Signature: "<value>",
        SignatureInput: "<value>",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.CreateVirtualCashIncreaseRequest](../../pkg/models/operations/createvirtualcashincreaserequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.CreateVirtualCashIncreaseResponse](../../pkg/models/operations/createvirtualcashincreaseresponse.md), error**
| Error Object                             | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| sdkerrors.CreateVirtualCashIncreaseError | 400,401,403,404,406,429,500,503,504      | application/problem+json                 |
| sdkerrors.SDKError                       | 4xx-5xx                                  | */*                                      |
