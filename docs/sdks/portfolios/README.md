# Portfolios
(*Portfolios*)

## Overview

All portfolios related paths.

### Available Operations

* [CancelPortfoliosOrder](#cancelportfoliosorder) - Cancel portfolios order
* [CreatePortfoliosAllocation](#createportfoliosallocation) - Create portfolios allocation
* [CreatePortfoliosConfiguration](#createportfoliosconfiguration) - Create portfolios configuration
* [CreatePortfoliosOrder](#createportfoliosorder) - Create portfolios order
* [CreatePortfoliosRebalancingStrategy](#createportfoliosrebalancingstrategy) - Create portfolios rebalancing strategy
* [ListPortfolioRebalancingExecutionOrders](#listportfoliorebalancingexecutionorders) - List portfolio rebalancing execution orders
* [ListPortfoliosAllocationAccounts](#listportfoliosallocationaccounts) - List portfolios allocation accounts
* [ListPortfoliosAllocations](#listportfoliosallocations) - List portfolios allocations
* [ListPortfoliosConfigurations](#listportfoliosconfigurations) - List portfolios configurations
* [ListPortfoliosOrders](#listportfoliosorders) - List portfolios orders
* [ListPortfoliosRebalancingStrategies](#listportfoliosrebalancingstrategies) - List portfolios rebalancing strategies
* [RetrievePortfoliosAllocation](#retrieveportfoliosallocation) - Retrieve portfolios allocation
* [RetrievePortfoliosConfiguration](#retrieveportfoliosconfiguration) - Retrieve portfolios configuration
* [RetrievePortfoliosOrder](#retrieveportfoliosorder) - Retrieve portfolios order
* [RetrievePortfoliosRebalancingExecution](#retrieveportfoliosrebalancingexecution) - Retrieve portfolios rebalancing execution
* [RetrievePortfoliosRebalancingStrategy](#retrieveportfoliosrebalancingstrategy) - Retrieve portfolios rebalancing strategy
* [TriggerPortfolioRebalancing](#triggerportfoliorebalancing) - Trigger portfolio rebalancing
* [UpdatePortfoliosAllocation](#updateportfoliosallocation) - Update portfolios allocation
* [UpdatePortfoliosConfiguration](#updateportfoliosconfiguration) - Update portfolios configuration

## CancelPortfoliosOrder

Cancel portfolios order

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
    res, err := s.Portfolios.CancelPortfoliosOrder(ctx, operations.CancelPortfoliosOrderRequest{
        PortfolioOrderID: "42863460-c076-471d-9b66-566e4a20d8de",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CancelPortfoliosOrderRequest](../../pkg/models/operations/cancelportfoliosorderrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.CancelPortfoliosOrderResponse](../../pkg/models/operations/cancelportfoliosorderresponse.md), error**
| Error Object                                              | Status Code                                               | Content Type                                              |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| sdkerrors.CancelPortfoliosOrderError                      | 401                                                       | application/problem+json                                  |
| sdkerrors.CancelPortfoliosOrderPortfoliosError            | 403                                                       | application/problem+json                                  |
| sdkerrors.CancelPortfoliosOrderPortfoliosResponseError    | 404                                                       | application/problem+json                                  |
| sdkerrors.CancelPortfoliosOrderPortfoliosResponse406Error | 406                                                       | application/problem+json                                  |
| sdkerrors.CancelPortfoliosOrderPortfoliosResponse429Error | 429                                                       | application/problem+json                                  |
| sdkerrors.CancelPortfoliosOrderPortfoliosResponse500Error | 500                                                       | application/problem+json                                  |
| sdkerrors.CancelPortfoliosOrderPortfoliosResponse503Error | 503                                                       | application/problem+json                                  |
| sdkerrors.CancelPortfoliosOrderPortfoliosResponse504Error | 504                                                       | application/problem+json                                  |
| sdkerrors.SDKError                                        | 400-600                                                   | */*                                                       |

## CreatePortfoliosAllocation

Create portfolios allocation

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
    res, err := s.Portfolios.CreatePortfoliosAllocation(ctx, operations.CreatePortfoliosAllocationRequest{
        RequestBody: &operations.CreatePortfoliosAllocationPortfoliosAllocationCreateRequest{
            Allocation: []operations.Allocation{
                operations.Allocation{
                    InstrumentID: "string",
                    Weight: "string",
                },
            },
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

    if res.TwoHundredApplicationJSONPortfoliosAllocation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.CreatePortfoliosAllocationRequest](../../pkg/models/operations/createportfoliosallocationrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.CreatePortfoliosAllocationResponse](../../pkg/models/operations/createportfoliosallocationresponse.md), error**
| Error Object                                                   | Status Code                                                    | Content Type                                                   |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| sdkerrors.CreatePortfoliosAllocationError                      | 400                                                            | application/problem+json                                       |
| sdkerrors.CreatePortfoliosAllocationPortfoliosError            | 401                                                            | application/problem+json                                       |
| sdkerrors.CreatePortfoliosAllocationPortfoliosResponseError    | 403                                                            | application/problem+json                                       |
| sdkerrors.CreatePortfoliosAllocationPortfoliosResponse404Error | 404                                                            | application/problem+json                                       |
| sdkerrors.CreatePortfoliosAllocationPortfoliosResponse406Error | 406                                                            | application/problem+json                                       |
| sdkerrors.CreatePortfoliosAllocationPortfoliosResponse429Error | 429                                                            | application/problem+json                                       |
| sdkerrors.CreatePortfoliosAllocationPortfoliosResponse500Error | 500                                                            | application/problem+json                                       |
| sdkerrors.CreatePortfoliosAllocationPortfoliosResponse503Error | 503                                                            | application/problem+json                                       |
| sdkerrors.CreatePortfoliosAllocationPortfoliosResponse504Error | 504                                                            | application/problem+json                                       |
| sdkerrors.SDKError                                             | 400-600                                                        | */*                                                            |

## CreatePortfoliosConfiguration

Create portfolios configuration

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
    res, err := s.Portfolios.CreatePortfoliosConfiguration(ctx, operations.CreatePortfoliosConfigurationRequest{
        RequestBody: &operations.CreatePortfoliosConfigurationPortfoliosConfigurationCreateRequest{
            AccountID: "b71530a0-f18c-4c45-b08c-4985b08dec50",
            AllocationID: "ef099f22-e4c3-48b2-a7ad-03c136e81b79",
            RebalancingStrategyIds: []string{
                "529a411e-a19e-4841-bc64-1d029019161f",
            },
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

    if res.TwoHundredApplicationJSONPortfoliosConfiguration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.CreatePortfoliosConfigurationRequest](../../pkg/models/operations/createportfoliosconfigurationrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.CreatePortfoliosConfigurationResponse](../../pkg/models/operations/createportfoliosconfigurationresponse.md), error**
| Error Object                                                      | Status Code                                                       | Content Type                                                      |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| sdkerrors.CreatePortfoliosConfigurationError                      | 400                                                               | application/problem+json                                          |
| sdkerrors.CreatePortfoliosConfigurationPortfoliosError            | 401                                                               | application/problem+json                                          |
| sdkerrors.CreatePortfoliosConfigurationPortfoliosResponseError    | 403                                                               | application/problem+json                                          |
| sdkerrors.CreatePortfoliosConfigurationPortfoliosResponse404Error | 404                                                               | application/problem+json                                          |
| sdkerrors.CreatePortfoliosConfigurationPortfoliosResponse406Error | 406                                                               | application/problem+json                                          |
| sdkerrors.CreatePortfoliosConfigurationPortfoliosResponse429Error | 429                                                               | application/problem+json                                          |
| sdkerrors.CreatePortfoliosConfigurationPortfoliosResponse500Error | 500                                                               | application/problem+json                                          |
| sdkerrors.CreatePortfoliosConfigurationPortfoliosResponse503Error | 503                                                               | application/problem+json                                          |
| sdkerrors.CreatePortfoliosConfigurationPortfoliosResponse504Error | 504                                                               | application/problem+json                                          |
| sdkerrors.SDKError                                                | 400-600                                                           | */*                                                               |

## CreatePortfoliosOrder

Create portfolios order

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
    res, err := s.Portfolios.CreatePortfoliosOrder(ctx, operations.CreatePortfoliosOrderRequest{
        RequestBody: &operations.CreatePortfoliosOrderPortfoliosOrderPlaceRequest{
            AccountID: "09386917-edc7-47c9-8e4c-97774c688b9c",
            CashAmount: "string",
            Side: operations.CreatePortfoliosOrderSideBuy,
            UserID: "72dd2124-1a6c-40cd-8d75-ad07e0d9c0bc",
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

    if res.TwoHundredApplicationJSONPortfoliosOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreatePortfoliosOrderRequest](../../pkg/models/operations/createportfoliosorderrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.CreatePortfoliosOrderResponse](../../pkg/models/operations/createportfoliosorderresponse.md), error**
| Error Object                                              | Status Code                                               | Content Type                                              |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| sdkerrors.CreatePortfoliosOrderError                      | 400                                                       | application/problem+json                                  |
| sdkerrors.CreatePortfoliosOrderPortfoliosError            | 401                                                       | application/problem+json                                  |
| sdkerrors.CreatePortfoliosOrderPortfoliosResponseError    | 403                                                       | application/problem+json                                  |
| sdkerrors.CreatePortfoliosOrderPortfoliosResponse404Error | 404                                                       | application/problem+json                                  |
| sdkerrors.CreatePortfoliosOrderPortfoliosResponse406Error | 406                                                       | application/problem+json                                  |
| sdkerrors.CreatePortfoliosOrderPortfoliosResponse429Error | 429                                                       | application/problem+json                                  |
| sdkerrors.CreatePortfoliosOrderPortfoliosResponse500Error | 500                                                       | application/problem+json                                  |
| sdkerrors.CreatePortfoliosOrderPortfoliosResponse503Error | 503                                                       | application/problem+json                                  |
| sdkerrors.CreatePortfoliosOrderPortfoliosResponse504Error | 504                                                       | application/problem+json                                  |
| sdkerrors.SDKError                                        | 400-600                                                   | */*                                                       |

## CreatePortfoliosRebalancingStrategy

Create portfolios rebalancing strategy

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
    res, err := s.Portfolios.CreatePortfoliosRebalancingStrategy(ctx, operations.CreatePortfoliosRebalancingStrategyRequest{
        RequestBody: &operations.CreatePortfoliosRebalancingStrategyPortfoliosRebalancingStrategyRequest{
            Conditions: []operations.Conditions{
                operations.Conditions{
                    AdditionalProperties: map[string]interface{}{
                        "key": "string",
                    },
                    Name: "string",
                    Type: operations.CreatePortfoliosRebalancingStrategyTypeScheduled,
                },
            },
            Name: "string",
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

    if res.TwoHundredApplicationJSONPortfoliosRebalancingStrategy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.CreatePortfoliosRebalancingStrategyRequest](../../pkg/models/operations/createportfoliosrebalancingstrategyrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |


### Response

**[*operations.CreatePortfoliosRebalancingStrategyResponse](../../pkg/models/operations/createportfoliosrebalancingstrategyresponse.md), error**
| Error Object                                                            | Status Code                                                             | Content Type                                                            |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| sdkerrors.CreatePortfoliosRebalancingStrategyError                      | 400                                                                     | application/problem+json                                                |
| sdkerrors.CreatePortfoliosRebalancingStrategyPortfoliosError            | 401                                                                     | application/problem+json                                                |
| sdkerrors.CreatePortfoliosRebalancingStrategyPortfoliosResponseError    | 403                                                                     | application/problem+json                                                |
| sdkerrors.CreatePortfoliosRebalancingStrategyPortfoliosResponse404Error | 404                                                                     | application/problem+json                                                |
| sdkerrors.CreatePortfoliosRebalancingStrategyPortfoliosResponse406Error | 406                                                                     | application/problem+json                                                |
| sdkerrors.CreatePortfoliosRebalancingStrategyPortfoliosResponse429Error | 429                                                                     | application/problem+json                                                |
| sdkerrors.CreatePortfoliosRebalancingStrategyPortfoliosResponse500Error | 500                                                                     | application/problem+json                                                |
| sdkerrors.CreatePortfoliosRebalancingStrategyPortfoliosResponse503Error | 503                                                                     | application/problem+json                                                |
| sdkerrors.CreatePortfoliosRebalancingStrategyPortfoliosResponse504Error | 504                                                                     | application/problem+json                                                |
| sdkerrors.SDKError                                                      | 400-600                                                                 | */*                                                                     |

## ListPortfolioRebalancingExecutionOrders

List portfolio rebalancing execution orders

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
    res, err := s.Portfolios.ListPortfolioRebalancingExecutionOrders(ctx, operations.ListPortfolioRebalancingExecutionOrdersRequest{
        ExecutionID: "faccb01b-05e2-49df-ba2f-94e57ee7faa0",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONPortfoliosRebalancingExecutionOrderListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.ListPortfolioRebalancingExecutionOrdersRequest](../../pkg/models/operations/listportfoliorebalancingexecutionordersrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |


### Response

**[*operations.ListPortfolioRebalancingExecutionOrdersResponse](../../pkg/models/operations/listportfoliorebalancingexecutionordersresponse.md), error**
| Error Object                                                                | Status Code                                                                 | Content Type                                                                |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| sdkerrors.ListPortfolioRebalancingExecutionOrdersError                      | 400                                                                         | application/problem+json                                                    |
| sdkerrors.ListPortfolioRebalancingExecutionOrdersPortfoliosError            | 401                                                                         | application/problem+json                                                    |
| sdkerrors.ListPortfolioRebalancingExecutionOrdersPortfoliosResponseError    | 403                                                                         | application/problem+json                                                    |
| sdkerrors.ListPortfolioRebalancingExecutionOrdersPortfoliosResponse404Error | 404                                                                         | application/problem+json                                                    |
| sdkerrors.ListPortfolioRebalancingExecutionOrdersPortfoliosResponse405Error | 405                                                                         | application/problem+json                                                    |
| sdkerrors.ListPortfolioRebalancingExecutionOrdersPortfoliosResponse406Error | 406                                                                         | application/problem+json                                                    |
| sdkerrors.ListPortfolioRebalancingExecutionOrdersPortfoliosResponse429Error | 429                                                                         | application/problem+json                                                    |
| sdkerrors.ListPortfolioRebalancingExecutionOrdersPortfoliosResponse500Error | 500                                                                         | application/problem+json                                                    |
| sdkerrors.ListPortfolioRebalancingExecutionOrdersPortfoliosResponse503Error | 503                                                                         | application/problem+json                                                    |
| sdkerrors.ListPortfolioRebalancingExecutionOrdersPortfoliosResponse504Error | 504                                                                         | application/problem+json                                                    |
| sdkerrors.SDKError                                                          | 400-600                                                                     | */*                                                                         |

## ListPortfoliosAllocationAccounts

List portfolios allocation accounts

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
    res, err := s.Portfolios.ListPortfoliosAllocationAccounts(ctx, operations.ListPortfoliosAllocationAccountsRequest{
        AllocationID: "9260907b-4e94-45c8-a31a-10d01d471a66",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONPortfoliosAllocationAccountsListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.ListPortfoliosAllocationAccountsRequest](../../pkg/models/operations/listportfoliosallocationaccountsrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.ListPortfoliosAllocationAccountsResponse](../../pkg/models/operations/listportfoliosallocationaccountsresponse.md), error**
| Error Object                                                         | Status Code                                                          | Content Type                                                         |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| sdkerrors.ListPortfoliosAllocationAccountsError                      | 400                                                                  | application/problem+json                                             |
| sdkerrors.ListPortfoliosAllocationAccountsPortfoliosError            | 401                                                                  | application/problem+json                                             |
| sdkerrors.ListPortfoliosAllocationAccountsPortfoliosResponseError    | 403                                                                  | application/problem+json                                             |
| sdkerrors.ListPortfoliosAllocationAccountsPortfoliosResponse404Error | 404                                                                  | application/problem+json                                             |
| sdkerrors.ListPortfoliosAllocationAccountsPortfoliosResponse405Error | 405                                                                  | application/problem+json                                             |
| sdkerrors.ListPortfoliosAllocationAccountsPortfoliosResponse406Error | 406                                                                  | application/problem+json                                             |
| sdkerrors.ListPortfoliosAllocationAccountsPortfoliosResponse429Error | 429                                                                  | application/problem+json                                             |
| sdkerrors.ListPortfoliosAllocationAccountsPortfoliosResponse500Error | 500                                                                  | application/problem+json                                             |
| sdkerrors.ListPortfoliosAllocationAccountsPortfoliosResponse503Error | 503                                                                  | application/problem+json                                             |
| sdkerrors.ListPortfoliosAllocationAccountsPortfoliosResponse504Error | 504                                                                  | application/problem+json                                             |
| sdkerrors.SDKError                                                   | 400-600                                                              | */*                                                                  |

## ListPortfoliosAllocations

List portfolios allocations

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
    res, err := s.Portfolios.ListPortfoliosAllocations(ctx, operations.ListPortfoliosAllocationsRequest{
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONPortfoliosAllocationsListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.ListPortfoliosAllocationsRequest](../../pkg/models/operations/listportfoliosallocationsrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.ListPortfoliosAllocationsResponse](../../pkg/models/operations/listportfoliosallocationsresponse.md), error**
| Error Object                                                  | Status Code                                                   | Content Type                                                  |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| sdkerrors.ListPortfoliosAllocationsError                      | 400                                                           | application/problem+json                                      |
| sdkerrors.ListPortfoliosAllocationsPortfoliosError            | 401                                                           | application/problem+json                                      |
| sdkerrors.ListPortfoliosAllocationsPortfoliosResponseError    | 403                                                           | application/problem+json                                      |
| sdkerrors.ListPortfoliosAllocationsPortfoliosResponse404Error | 404                                                           | application/problem+json                                      |
| sdkerrors.ListPortfoliosAllocationsPortfoliosResponse405Error | 405                                                           | application/problem+json                                      |
| sdkerrors.ListPortfoliosAllocationsPortfoliosResponse406Error | 406                                                           | application/problem+json                                      |
| sdkerrors.ListPortfoliosAllocationsPortfoliosResponse429Error | 429                                                           | application/problem+json                                      |
| sdkerrors.ListPortfoliosAllocationsPortfoliosResponse500Error | 500                                                           | application/problem+json                                      |
| sdkerrors.ListPortfoliosAllocationsPortfoliosResponse503Error | 503                                                           | application/problem+json                                      |
| sdkerrors.ListPortfoliosAllocationsPortfoliosResponse504Error | 504                                                           | application/problem+json                                      |
| sdkerrors.SDKError                                            | 400-600                                                       | */*                                                           |

## ListPortfoliosConfigurations

List portfolios configurations

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
    res, err := s.Portfolios.ListPortfoliosConfigurations(ctx, operations.ListPortfoliosConfigurationsRequest{
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONPortfoliosConfigurationsListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.ListPortfoliosConfigurationsRequest](../../pkg/models/operations/listportfoliosconfigurationsrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.ListPortfoliosConfigurationsResponse](../../pkg/models/operations/listportfoliosconfigurationsresponse.md), error**
| Error Object                                                     | Status Code                                                      | Content Type                                                     |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| sdkerrors.ListPortfoliosConfigurationsError                      | 400                                                              | application/problem+json                                         |
| sdkerrors.ListPortfoliosConfigurationsPortfoliosError            | 401                                                              | application/problem+json                                         |
| sdkerrors.ListPortfoliosConfigurationsPortfoliosResponseError    | 403                                                              | application/problem+json                                         |
| sdkerrors.ListPortfoliosConfigurationsPortfoliosResponse404Error | 404                                                              | application/problem+json                                         |
| sdkerrors.ListPortfoliosConfigurationsPortfoliosResponse405Error | 405                                                              | application/problem+json                                         |
| sdkerrors.ListPortfoliosConfigurationsPortfoliosResponse406Error | 406                                                              | application/problem+json                                         |
| sdkerrors.ListPortfoliosConfigurationsPortfoliosResponse429Error | 429                                                              | application/problem+json                                         |
| sdkerrors.ListPortfoliosConfigurationsPortfoliosResponse500Error | 500                                                              | application/problem+json                                         |
| sdkerrors.ListPortfoliosConfigurationsPortfoliosResponse503Error | 503                                                              | application/problem+json                                         |
| sdkerrors.ListPortfoliosConfigurationsPortfoliosResponse504Error | 504                                                              | application/problem+json                                         |
| sdkerrors.SDKError                                               | 400-600                                                          | */*                                                              |

## ListPortfoliosOrders

List portfolios orders

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
    res, err := s.Portfolios.ListPortfoliosOrders(ctx, operations.ListPortfoliosOrdersRequest{
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONPortfoliosOrdersListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListPortfoliosOrdersRequest](../../pkg/models/operations/listportfoliosordersrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.ListPortfoliosOrdersResponse](../../pkg/models/operations/listportfoliosordersresponse.md), error**
| Error Object                                             | Status Code                                              | Content Type                                             |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| sdkerrors.ListPortfoliosOrdersError                      | 400                                                      | application/problem+json                                 |
| sdkerrors.ListPortfoliosOrdersPortfoliosError            | 401                                                      | application/problem+json                                 |
| sdkerrors.ListPortfoliosOrdersPortfoliosResponseError    | 403                                                      | application/problem+json                                 |
| sdkerrors.ListPortfoliosOrdersPortfoliosResponse404Error | 404                                                      | application/problem+json                                 |
| sdkerrors.ListPortfoliosOrdersPortfoliosResponse405Error | 405                                                      | application/problem+json                                 |
| sdkerrors.ListPortfoliosOrdersPortfoliosResponse406Error | 406                                                      | application/problem+json                                 |
| sdkerrors.ListPortfoliosOrdersPortfoliosResponse429Error | 429                                                      | application/problem+json                                 |
| sdkerrors.ListPortfoliosOrdersPortfoliosResponse500Error | 500                                                      | application/problem+json                                 |
| sdkerrors.ListPortfoliosOrdersPortfoliosResponse503Error | 503                                                      | application/problem+json                                 |
| sdkerrors.ListPortfoliosOrdersPortfoliosResponse504Error | 504                                                      | application/problem+json                                 |
| sdkerrors.SDKError                                       | 400-600                                                  | */*                                                      |

## ListPortfoliosRebalancingStrategies

List portfolios rebalancing strategies

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
    res, err := s.Portfolios.ListPortfoliosRebalancingStrategies(ctx, operations.ListPortfoliosRebalancingStrategiesRequest{
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONPortfoliosRebalancingStrategyListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.ListPortfoliosRebalancingStrategiesRequest](../../pkg/models/operations/listportfoliosrebalancingstrategiesrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |


### Response

**[*operations.ListPortfoliosRebalancingStrategiesResponse](../../pkg/models/operations/listportfoliosrebalancingstrategiesresponse.md), error**
| Error Object                                                            | Status Code                                                             | Content Type                                                            |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| sdkerrors.ListPortfoliosRebalancingStrategiesError                      | 400                                                                     | application/problem+json                                                |
| sdkerrors.ListPortfoliosRebalancingStrategiesPortfoliosError            | 401                                                                     | application/problem+json                                                |
| sdkerrors.ListPortfoliosRebalancingStrategiesPortfoliosResponseError    | 403                                                                     | application/problem+json                                                |
| sdkerrors.ListPortfoliosRebalancingStrategiesPortfoliosResponse404Error | 404                                                                     | application/problem+json                                                |
| sdkerrors.ListPortfoliosRebalancingStrategiesPortfoliosResponse405Error | 405                                                                     | application/problem+json                                                |
| sdkerrors.ListPortfoliosRebalancingStrategiesPortfoliosResponse406Error | 406                                                                     | application/problem+json                                                |
| sdkerrors.ListPortfoliosRebalancingStrategiesPortfoliosResponse429Error | 429                                                                     | application/problem+json                                                |
| sdkerrors.ListPortfoliosRebalancingStrategiesPortfoliosResponse500Error | 500                                                                     | application/problem+json                                                |
| sdkerrors.ListPortfoliosRebalancingStrategiesPortfoliosResponse503Error | 503                                                                     | application/problem+json                                                |
| sdkerrors.ListPortfoliosRebalancingStrategiesPortfoliosResponse504Error | 504                                                                     | application/problem+json                                                |
| sdkerrors.SDKError                                                      | 400-600                                                                 | */*                                                                     |

## RetrievePortfoliosAllocation

Retrieve portfolios allocation

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
    res, err := s.Portfolios.RetrievePortfoliosAllocation(ctx, operations.RetrievePortfoliosAllocationRequest{
        AllocationID: "8902e6a1-152d-44ee-8ba2-1f95e37ca0c1",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONPortfoliosAllocation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.RetrievePortfoliosAllocationRequest](../../pkg/models/operations/retrieveportfoliosallocationrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.RetrievePortfoliosAllocationResponse](../../pkg/models/operations/retrieveportfoliosallocationresponse.md), error**
| Error Object                                                     | Status Code                                                      | Content Type                                                     |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| sdkerrors.RetrievePortfoliosAllocationError                      | 401                                                              | application/problem+json                                         |
| sdkerrors.RetrievePortfoliosAllocationPortfoliosError            | 403                                                              | application/problem+json                                         |
| sdkerrors.RetrievePortfoliosAllocationPortfoliosResponseError    | 404                                                              | application/problem+json                                         |
| sdkerrors.RetrievePortfoliosAllocationPortfoliosResponse406Error | 406                                                              | application/problem+json                                         |
| sdkerrors.RetrievePortfoliosAllocationPortfoliosResponse429Error | 429                                                              | application/problem+json                                         |
| sdkerrors.RetrievePortfoliosAllocationPortfoliosResponse500Error | 500                                                              | application/problem+json                                         |
| sdkerrors.RetrievePortfoliosAllocationPortfoliosResponse503Error | 503                                                              | application/problem+json                                         |
| sdkerrors.RetrievePortfoliosAllocationPortfoliosResponse504Error | 504                                                              | application/problem+json                                         |
| sdkerrors.SDKError                                               | 400-600                                                          | */*                                                              |

## RetrievePortfoliosConfiguration

Retrieve portfolios configuration

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
    res, err := s.Portfolios.RetrievePortfoliosConfiguration(ctx, operations.RetrievePortfoliosConfigurationRequest{
        AccountID: "823e6a8e-d630-4a74-9292-5b3ae05f7ca5",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONPortfoliosConfiguration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.RetrievePortfoliosConfigurationRequest](../../pkg/models/operations/retrieveportfoliosconfigurationrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.RetrievePortfoliosConfigurationResponse](../../pkg/models/operations/retrieveportfoliosconfigurationresponse.md), error**
| Error Object                                                        | Status Code                                                         | Content Type                                                        |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| sdkerrors.RetrievePortfoliosConfigurationError                      | 401                                                                 | application/problem+json                                            |
| sdkerrors.RetrievePortfoliosConfigurationPortfoliosError            | 403                                                                 | application/problem+json                                            |
| sdkerrors.RetrievePortfoliosConfigurationPortfoliosResponseError    | 404                                                                 | application/problem+json                                            |
| sdkerrors.RetrievePortfoliosConfigurationPortfoliosResponse405Error | 405                                                                 | application/problem+json                                            |
| sdkerrors.RetrievePortfoliosConfigurationPortfoliosResponse406Error | 406                                                                 | application/problem+json                                            |
| sdkerrors.RetrievePortfoliosConfigurationPortfoliosResponse429Error | 429                                                                 | application/problem+json                                            |
| sdkerrors.RetrievePortfoliosConfigurationPortfoliosResponse500Error | 500                                                                 | application/problem+json                                            |
| sdkerrors.RetrievePortfoliosConfigurationPortfoliosResponse503Error | 503                                                                 | application/problem+json                                            |
| sdkerrors.RetrievePortfoliosConfigurationPortfoliosResponse504Error | 504                                                                 | application/problem+json                                            |
| sdkerrors.SDKError                                                  | 400-600                                                             | */*                                                                 |

## RetrievePortfoliosOrder

Retrieve portfolios order

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
    res, err := s.Portfolios.RetrievePortfoliosOrder(ctx, operations.RetrievePortfoliosOrderRequest{
        PortfolioOrderID: "e8307df3-b75a-46de-a05a-9fe332308857",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONPortfoliosOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.RetrievePortfoliosOrderRequest](../../pkg/models/operations/retrieveportfoliosorderrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.RetrievePortfoliosOrderResponse](../../pkg/models/operations/retrieveportfoliosorderresponse.md), error**
| Error Object                                                | Status Code                                                 | Content Type                                                |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| sdkerrors.RetrievePortfoliosOrderError                      | 401                                                         | application/problem+json                                    |
| sdkerrors.RetrievePortfoliosOrderPortfoliosError            | 403                                                         | application/problem+json                                    |
| sdkerrors.RetrievePortfoliosOrderPortfoliosResponseError    | 404                                                         | application/problem+json                                    |
| sdkerrors.RetrievePortfoliosOrderPortfoliosResponse406Error | 406                                                         | application/problem+json                                    |
| sdkerrors.RetrievePortfoliosOrderPortfoliosResponse429Error | 429                                                         | application/problem+json                                    |
| sdkerrors.RetrievePortfoliosOrderPortfoliosResponse500Error | 500                                                         | application/problem+json                                    |
| sdkerrors.RetrievePortfoliosOrderPortfoliosResponse503Error | 503                                                         | application/problem+json                                    |
| sdkerrors.RetrievePortfoliosOrderPortfoliosResponse504Error | 504                                                         | application/problem+json                                    |
| sdkerrors.SDKError                                          | 400-600                                                     | */*                                                         |

## RetrievePortfoliosRebalancingExecution

Retrieve portfolios rebalancing execution

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
    res, err := s.Portfolios.RetrievePortfoliosRebalancingExecution(ctx, operations.RetrievePortfoliosRebalancingExecutionRequest{
        ExecutionID: "99606d30-438e-43e3-948b-6cd364997761",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONPortfoliosRebalancingExecution != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.RetrievePortfoliosRebalancingExecutionRequest](../../pkg/models/operations/retrieveportfoliosrebalancingexecutionrequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |


### Response

**[*operations.RetrievePortfoliosRebalancingExecutionResponse](../../pkg/models/operations/retrieveportfoliosrebalancingexecutionresponse.md), error**
| Error Object                                                               | Status Code                                                                | Content Type                                                               |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| sdkerrors.RetrievePortfoliosRebalancingExecutionError                      | 401                                                                        | application/problem+json                                                   |
| sdkerrors.RetrievePortfoliosRebalancingExecutionPortfoliosError            | 403                                                                        | application/problem+json                                                   |
| sdkerrors.RetrievePortfoliosRebalancingExecutionPortfoliosResponseError    | 404                                                                        | application/problem+json                                                   |
| sdkerrors.RetrievePortfoliosRebalancingExecutionPortfoliosResponse405Error | 405                                                                        | application/problem+json                                                   |
| sdkerrors.RetrievePortfoliosRebalancingExecutionPortfoliosResponse406Error | 406                                                                        | application/problem+json                                                   |
| sdkerrors.RetrievePortfoliosRebalancingExecutionPortfoliosResponse429Error | 429                                                                        | application/problem+json                                                   |
| sdkerrors.RetrievePortfoliosRebalancingExecutionPortfoliosResponse500Error | 500                                                                        | application/problem+json                                                   |
| sdkerrors.RetrievePortfoliosRebalancingExecutionPortfoliosResponse503Error | 503                                                                        | application/problem+json                                                   |
| sdkerrors.RetrievePortfoliosRebalancingExecutionPortfoliosResponse504Error | 504                                                                        | application/problem+json                                                   |
| sdkerrors.SDKError                                                         | 400-600                                                                    | */*                                                                        |

## RetrievePortfoliosRebalancingStrategy

Retrieve portfolios rebalancing strategy

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
    res, err := s.Portfolios.RetrievePortfoliosRebalancingStrategy(ctx, operations.RetrievePortfoliosRebalancingStrategyRequest{
        Signature: "string",
        SignatureInput: "string",
        StrategyID: "0c1478b8-c314-4033-ba97-4a21434d1d07",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONPortfoliosRebalancingStrategy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.RetrievePortfoliosRebalancingStrategyRequest](../../pkg/models/operations/retrieveportfoliosrebalancingstrategyrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |


### Response

**[*operations.RetrievePortfoliosRebalancingStrategyResponse](../../pkg/models/operations/retrieveportfoliosrebalancingstrategyresponse.md), error**
| Error Object                                                              | Status Code                                                               | Content Type                                                              |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| sdkerrors.RetrievePortfoliosRebalancingStrategyError                      | 401                                                                       | application/problem+json                                                  |
| sdkerrors.RetrievePortfoliosRebalancingStrategyPortfoliosError            | 403                                                                       | application/problem+json                                                  |
| sdkerrors.RetrievePortfoliosRebalancingStrategyPortfoliosResponseError    | 404                                                                       | application/problem+json                                                  |
| sdkerrors.RetrievePortfoliosRebalancingStrategyPortfoliosResponse405Error | 405                                                                       | application/problem+json                                                  |
| sdkerrors.RetrievePortfoliosRebalancingStrategyPortfoliosResponse406Error | 406                                                                       | application/problem+json                                                  |
| sdkerrors.RetrievePortfoliosRebalancingStrategyPortfoliosResponse429Error | 429                                                                       | application/problem+json                                                  |
| sdkerrors.RetrievePortfoliosRebalancingStrategyPortfoliosResponse500Error | 500                                                                       | application/problem+json                                                  |
| sdkerrors.RetrievePortfoliosRebalancingStrategyPortfoliosResponse503Error | 503                                                                       | application/problem+json                                                  |
| sdkerrors.RetrievePortfoliosRebalancingStrategyPortfoliosResponse504Error | 504                                                                       | application/problem+json                                                  |
| sdkerrors.SDKError                                                        | 400-600                                                                   | */*                                                                       |

## TriggerPortfolioRebalancing

Trigger portfolio rebalancing

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
    res, err := s.Portfolios.TriggerPortfolioRebalancing(ctx, operations.TriggerPortfolioRebalancingRequest{
        RequestBody: operations.CreateTriggerPortfolioRebalancingTriggerPortfolioRebalancingRequestAccounts(
                operations.Accounts{
                    Accounts: []string{
                        "04eef0b9-9d66-414b-85ad-3037ab127a63",
                    },
                },
        ),
        IdempotencyKey: "ccb07f42-4104-44ad-8e1f-c660bb7b269c",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONTriggerPortfolioRebalancingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.TriggerPortfolioRebalancingRequest](../../pkg/models/operations/triggerportfoliorebalancingrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.TriggerPortfolioRebalancingResponse](../../pkg/models/operations/triggerportfoliorebalancingresponse.md), error**
| Error Object                                                    | Status Code                                                     | Content Type                                                    |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| sdkerrors.TriggerPortfolioRebalancingError                      | 400                                                             | application/problem+json                                        |
| sdkerrors.TriggerPortfolioRebalancingPortfoliosError            | 401                                                             | application/problem+json                                        |
| sdkerrors.TriggerPortfolioRebalancingPortfoliosResponseError    | 403                                                             | application/problem+json                                        |
| sdkerrors.TriggerPortfolioRebalancingPortfoliosResponse404Error | 404                                                             | application/problem+json                                        |
| sdkerrors.TriggerPortfolioRebalancingPortfoliosResponse406Error | 406                                                             | application/problem+json                                        |
| sdkerrors.TriggerPortfolioRebalancingPortfoliosResponse429Error | 429                                                             | application/problem+json                                        |
| sdkerrors.TriggerPortfolioRebalancingPortfoliosResponse500Error | 500                                                             | application/problem+json                                        |
| sdkerrors.TriggerPortfolioRebalancingPortfoliosResponse503Error | 503                                                             | application/problem+json                                        |
| sdkerrors.TriggerPortfolioRebalancingPortfoliosResponse504Error | 504                                                             | application/problem+json                                        |
| sdkerrors.SDKError                                              | 400-600                                                         | */*                                                             |

## UpdatePortfoliosAllocation

Update portfolios allocation

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
    res, err := s.Portfolios.UpdatePortfoliosAllocation(ctx, operations.UpdatePortfoliosAllocationRequest{
        RequestBody: &operations.UpdatePortfoliosAllocationPortfoliosAllocationUpdateRequest{
            Allocation: []operations.UpdatePortfoliosAllocationAllocation{
                operations.UpdatePortfoliosAllocationAllocation{
                    InstrumentID: "string",
                    Weight: "string",
                },
            },
        },
        AllocationID: "c7ddcc6d-56bb-41c2-a344-b57fb4000627",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONPortfoliosAllocation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.UpdatePortfoliosAllocationRequest](../../pkg/models/operations/updateportfoliosallocationrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.UpdatePortfoliosAllocationResponse](../../pkg/models/operations/updateportfoliosallocationresponse.md), error**
| Error Object                                                   | Status Code                                                    | Content Type                                                   |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| sdkerrors.UpdatePortfoliosAllocationError                      | 400                                                            | application/problem+json                                       |
| sdkerrors.UpdatePortfoliosAllocationPortfoliosError            | 401                                                            | application/problem+json                                       |
| sdkerrors.UpdatePortfoliosAllocationPortfoliosResponseError    | 403                                                            | application/problem+json                                       |
| sdkerrors.UpdatePortfoliosAllocationPortfoliosResponse404Error | 404                                                            | application/problem+json                                       |
| sdkerrors.UpdatePortfoliosAllocationPortfoliosResponse406Error | 406                                                            | application/problem+json                                       |
| sdkerrors.UpdatePortfoliosAllocationPortfoliosResponse429Error | 429                                                            | application/problem+json                                       |
| sdkerrors.UpdatePortfoliosAllocationPortfoliosResponse500Error | 500                                                            | application/problem+json                                       |
| sdkerrors.UpdatePortfoliosAllocationPortfoliosResponse503Error | 503                                                            | application/problem+json                                       |
| sdkerrors.UpdatePortfoliosAllocationPortfoliosResponse504Error | 504                                                            | application/problem+json                                       |
| sdkerrors.SDKError                                             | 400-600                                                        | */*                                                            |

## UpdatePortfoliosConfiguration

Update portfolios configuration

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
    res, err := s.Portfolios.UpdatePortfoliosConfiguration(ctx, operations.UpdatePortfoliosConfigurationRequest{
        RequestBody: &operations.UpdatePortfoliosConfigurationPortfoliosConfigurationUpdateRequest{
            RebalancingStrategyIds: []string{
                "be034715-bf00-452a-b6f6-ba5a9dfa7ad7",
            },
        },
        AccountID: "4243e3c9-f31f-4939-b731-77f5e67f038c",
        IdempotencyKey: "ccb07f42-4104-44ad-8e1f-c660bb7b269c",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONPortfoliosConfiguration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.UpdatePortfoliosConfigurationRequest](../../pkg/models/operations/updateportfoliosconfigurationrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.UpdatePortfoliosConfigurationResponse](../../pkg/models/operations/updateportfoliosconfigurationresponse.md), error**
| Error Object                                                      | Status Code                                                       | Content Type                                                      |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| sdkerrors.UpdatePortfoliosConfigurationError                      | 400                                                               | application/problem+json                                          |
| sdkerrors.UpdatePortfoliosConfigurationPortfoliosError            | 401                                                               | application/problem+json                                          |
| sdkerrors.UpdatePortfoliosConfigurationPortfoliosResponseError    | 403                                                               | application/problem+json                                          |
| sdkerrors.UpdatePortfoliosConfigurationPortfoliosResponse404Error | 404                                                               | application/problem+json                                          |
| sdkerrors.UpdatePortfoliosConfigurationPortfoliosResponse406Error | 406                                                               | application/problem+json                                          |
| sdkerrors.UpdatePortfoliosConfigurationPortfoliosResponse429Error | 429                                                               | application/problem+json                                          |
| sdkerrors.UpdatePortfoliosConfigurationPortfoliosResponse500Error | 500                                                               | application/problem+json                                          |
| sdkerrors.UpdatePortfoliosConfigurationPortfoliosResponse503Error | 503                                                               | application/problem+json                                          |
| sdkerrors.UpdatePortfoliosConfigurationPortfoliosResponse504Error | 504                                                               | application/problem+json                                          |
| sdkerrors.SDKError                                                | 400-600                                                           | */*                                                               |
