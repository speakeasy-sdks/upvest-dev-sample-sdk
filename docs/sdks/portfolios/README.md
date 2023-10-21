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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CancelPortfoliosOrderRequest](../../models/operations/cancelportfoliosorderrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.CancelPortfoliosOrderResponse](../../models/operations/cancelportfoliosorderresponse.md), error**


## CreatePortfoliosAllocation

Create portfolios allocation

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
    res, err := s.Portfolios.CreatePortfoliosAllocation(ctx, operations.CreatePortfoliosAllocationRequest{
        RequestBody: &operations.CreatePortfoliosAllocationPortfoliosAllocationCreateRequest{
            Allocation: []operations.CreatePortfoliosAllocationPortfoliosAllocationCreateRequestAllocation{
                operations.CreatePortfoliosAllocationPortfoliosAllocationCreateRequestAllocation{
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

    if res.PortfoliosAllocation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.CreatePortfoliosAllocationRequest](../../models/operations/createportfoliosallocationrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.CreatePortfoliosAllocationResponse](../../models/operations/createportfoliosallocationresponse.md), error**


## CreatePortfoliosConfiguration

Create portfolios configuration

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

    if res.PortfoliosConfiguration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.CreatePortfoliosConfigurationRequest](../../models/operations/createportfoliosconfigurationrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.CreatePortfoliosConfigurationResponse](../../models/operations/createportfoliosconfigurationresponse.md), error**


## CreatePortfoliosOrder

Create portfolios order

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
    res, err := s.Portfolios.CreatePortfoliosOrder(ctx, operations.CreatePortfoliosOrderRequest{
        RequestBody: &operations.CreatePortfoliosOrderPortfoliosOrderPlaceRequest{
            AccountID: "09386917-edc7-47c9-8e4c-97774c688b9c",
            CashAmount: "string",
            Side: operations.CreatePortfoliosOrderPortfoliosOrderPlaceRequestSideBuy,
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

    if res.PortfoliosOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreatePortfoliosOrderRequest](../../models/operations/createportfoliosorderrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.CreatePortfoliosOrderResponse](../../models/operations/createportfoliosorderresponse.md), error**


## CreatePortfoliosRebalancingStrategy

Create portfolios rebalancing strategy

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
    res, err := s.Portfolios.CreatePortfoliosRebalancingStrategy(ctx, operations.CreatePortfoliosRebalancingStrategyRequest{
        RequestBody: &operations.CreatePortfoliosRebalancingStrategyPortfoliosRebalancingStrategyRequest{
            Conditions: []operations.CreatePortfoliosRebalancingStrategyPortfoliosRebalancingStrategyRequestConditions{
                operations.CreatePortfoliosRebalancingStrategyPortfoliosRebalancingStrategyRequestConditions{
                    AdditionalProperties: map[string]interface{}{
                        "key": "string",
                    },
                    Name: "string",
                    Type: operations.CreatePortfoliosRebalancingStrategyPortfoliosRebalancingStrategyRequestConditionsTypeScheduled,
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

    if res.PortfoliosRebalancingStrategy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.CreatePortfoliosRebalancingStrategyRequest](../../models/operations/createportfoliosrebalancingstrategyrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.CreatePortfoliosRebalancingStrategyResponse](../../models/operations/createportfoliosrebalancingstrategyresponse.md), error**


## ListPortfolioRebalancingExecutionOrders

List portfolio rebalancing execution orders

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

    if res.PortfoliosRebalancingExecutionOrderListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.ListPortfolioRebalancingExecutionOrdersRequest](../../models/operations/listportfoliorebalancingexecutionordersrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |


### Response

**[*operations.ListPortfolioRebalancingExecutionOrdersResponse](../../models/operations/listportfoliorebalancingexecutionordersresponse.md), error**


## ListPortfoliosAllocationAccounts

List portfolios allocation accounts

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

    if res.PortfoliosAllocationAccountsListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.ListPortfoliosAllocationAccountsRequest](../../models/operations/listportfoliosallocationaccountsrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


### Response

**[*operations.ListPortfoliosAllocationAccountsResponse](../../models/operations/listportfoliosallocationaccountsresponse.md), error**


## ListPortfoliosAllocations

List portfolios allocations

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
    res, err := s.Portfolios.ListPortfoliosAllocations(ctx, operations.ListPortfoliosAllocationsRequest{
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PortfoliosAllocationsListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ListPortfoliosAllocationsRequest](../../models/operations/listportfoliosallocationsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.ListPortfoliosAllocationsResponse](../../models/operations/listportfoliosallocationsresponse.md), error**


## ListPortfoliosConfigurations

List portfolios configurations

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
    res, err := s.Portfolios.ListPortfoliosConfigurations(ctx, operations.ListPortfoliosConfigurationsRequest{
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PortfoliosConfigurationsListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.ListPortfoliosConfigurationsRequest](../../models/operations/listportfoliosconfigurationsrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.ListPortfoliosConfigurationsResponse](../../models/operations/listportfoliosconfigurationsresponse.md), error**


## ListPortfoliosOrders

List portfolios orders

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
    res, err := s.Portfolios.ListPortfoliosOrders(ctx, operations.ListPortfoliosOrdersRequest{
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListPortfoliosOrdersRequest](../../models/operations/listportfoliosordersrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.ListPortfoliosOrdersResponse](../../models/operations/listportfoliosordersresponse.md), error**


## ListPortfoliosRebalancingStrategies

List portfolios rebalancing strategies

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
    res, err := s.Portfolios.ListPortfoliosRebalancingStrategies(ctx, operations.ListPortfoliosRebalancingStrategiesRequest{
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PortfoliosRebalancingStrategyListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.ListPortfoliosRebalancingStrategiesRequest](../../models/operations/listportfoliosrebalancingstrategiesrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.ListPortfoliosRebalancingStrategiesResponse](../../models/operations/listportfoliosrebalancingstrategiesresponse.md), error**


## RetrievePortfoliosAllocation

Retrieve portfolios allocation

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

    if res.PortfoliosAllocation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.RetrievePortfoliosAllocationRequest](../../models/operations/retrieveportfoliosallocationrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.RetrievePortfoliosAllocationResponse](../../models/operations/retrieveportfoliosallocationresponse.md), error**


## RetrievePortfoliosConfiguration

Retrieve portfolios configuration

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

    if res.PortfoliosConfiguration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.RetrievePortfoliosConfigurationRequest](../../models/operations/retrieveportfoliosconfigurationrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.RetrievePortfoliosConfigurationResponse](../../models/operations/retrieveportfoliosconfigurationresponse.md), error**


## RetrievePortfoliosOrder

Retrieve portfolios order

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

    if res.PortfoliosOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.RetrievePortfoliosOrderRequest](../../models/operations/retrieveportfoliosorderrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.RetrievePortfoliosOrderResponse](../../models/operations/retrieveportfoliosorderresponse.md), error**


## RetrievePortfoliosRebalancingExecution

Retrieve portfolios rebalancing execution

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

    if res.PortfoliosRebalancingExecution != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.RetrievePortfoliosRebalancingExecutionRequest](../../models/operations/retrieveportfoliosrebalancingexecutionrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |


### Response

**[*operations.RetrievePortfoliosRebalancingExecutionResponse](../../models/operations/retrieveportfoliosrebalancingexecutionresponse.md), error**


## RetrievePortfoliosRebalancingStrategy

Retrieve portfolios rebalancing strategy

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

    if res.PortfoliosRebalancingStrategy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.RetrievePortfoliosRebalancingStrategyRequest](../../models/operations/retrieveportfoliosrebalancingstrategyrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |


### Response

**[*operations.RetrievePortfoliosRebalancingStrategyResponse](../../models/operations/retrieveportfoliosrebalancingstrategyresponse.md), error**


## TriggerPortfolioRebalancing

Trigger portfolio rebalancing

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
    res, err := s.Portfolios.TriggerPortfolioRebalancing(ctx, operations.TriggerPortfolioRebalancingRequest{
        RequestBody: operations.CreateTriggerPortfolioRebalancingTriggerPortfolioRebalancingRequestTriggerPortfolioRebalancingTriggerPortfolioRebalancingRequestAccounts(
                operations.TriggerPortfolioRebalancingTriggerPortfolioRebalancingRequestAccounts{
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

    if res.TriggerPortfolioRebalancingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.TriggerPortfolioRebalancingRequest](../../models/operations/triggerportfoliorebalancingrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.TriggerPortfolioRebalancingResponse](../../models/operations/triggerportfoliorebalancingresponse.md), error**


## UpdatePortfoliosAllocation

Update portfolios allocation

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
    res, err := s.Portfolios.UpdatePortfoliosAllocation(ctx, operations.UpdatePortfoliosAllocationRequest{
        RequestBody: &operations.UpdatePortfoliosAllocationPortfoliosAllocationUpdateRequest{
            Allocation: []operations.UpdatePortfoliosAllocationPortfoliosAllocationUpdateRequestAllocation{
                operations.UpdatePortfoliosAllocationPortfoliosAllocationUpdateRequestAllocation{
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

    if res.PortfoliosAllocation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.UpdatePortfoliosAllocationRequest](../../models/operations/updateportfoliosallocationrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.UpdatePortfoliosAllocationResponse](../../models/operations/updateportfoliosallocationresponse.md), error**


## UpdatePortfoliosConfiguration

Update portfolios configuration

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

    if res.PortfoliosConfiguration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.UpdatePortfoliosConfigurationRequest](../../models/operations/updateportfoliosconfigurationrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.UpdatePortfoliosConfigurationResponse](../../models/operations/updateportfoliosconfigurationresponse.md), error**

