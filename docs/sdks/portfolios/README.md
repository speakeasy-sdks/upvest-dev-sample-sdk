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
)

func main() {
    s := upvestdevsamplesdk.New(
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Portfolios.CancelPortfoliosOrder(ctx, operations.CancelPortfoliosOrderRequest{
        PortfolioOrderID: "42863460-c076-471d-9b66-566e4a20d8de",
        Signature: "<value>",
        SignatureInput: "<value>",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CancelPortfoliosOrderRequest](../../pkg/models/operations/cancelportfoliosorderrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.CancelPortfoliosOrderResponse](../../pkg/models/operations/cancelportfoliosorderresponse.md), error**
| Error Object                         | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| sdkerrors.CancelPortfoliosOrderError | 401,403,404,406,429,500,503,504      | application/problem+json             |
| sdkerrors.SDKError                   | 4xx-5xx                              | */*                                  |

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Portfolios.CreatePortfoliosAllocation(ctx, operations.CreatePortfoliosAllocationRequest{
        IdempotencyKey: "ccb07f42-4104-44ad-8e1f-c660bb7b269c",
        Signature: "<value>",
        SignatureInput: "<value>",
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
| `request`                                                                                                        | [operations.CreatePortfoliosAllocationRequest](../../pkg/models/operations/createportfoliosallocationrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.CreatePortfoliosAllocationResponse](../../pkg/models/operations/createportfoliosallocationresponse.md), error**
| Error Object                              | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| sdkerrors.CreatePortfoliosAllocationError | 400,401,403,404,406,429,500,503,504       | application/problem+json                  |
| sdkerrors.SDKError                        | 4xx-5xx                                   | */*                                       |

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Portfolios.CreatePortfoliosConfiguration(ctx, operations.CreatePortfoliosConfigurationRequest{
        IdempotencyKey: "ccb07f42-4104-44ad-8e1f-c660bb7b269c",
        Signature: "<value>",
        SignatureInput: "<value>",
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
| `request`                                                                                                              | [operations.CreatePortfoliosConfigurationRequest](../../pkg/models/operations/createportfoliosconfigurationrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.CreatePortfoliosConfigurationResponse](../../pkg/models/operations/createportfoliosconfigurationresponse.md), error**
| Error Object                                 | Status Code                                  | Content Type                                 |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| sdkerrors.CreatePortfoliosConfigurationError | 400,401,403,404,406,429,500,503,504          | application/problem+json                     |
| sdkerrors.SDKError                           | 4xx-5xx                                      | */*                                          |

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Portfolios.CreatePortfoliosOrder(ctx, operations.CreatePortfoliosOrderRequest{
        IdempotencyKey: "ccb07f42-4104-44ad-8e1f-c660bb7b269c",
        Signature: "<value>",
        SignatureInput: "<value>",
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
| `request`                                                                                              | [operations.CreatePortfoliosOrderRequest](../../pkg/models/operations/createportfoliosorderrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.CreatePortfoliosOrderResponse](../../pkg/models/operations/createportfoliosorderresponse.md), error**
| Error Object                         | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| sdkerrors.CreatePortfoliosOrderError | 400,401,403,404,406,429,500,503,504  | application/problem+json             |
| sdkerrors.SDKError                   | 4xx-5xx                              | */*                                  |

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Portfolios.CreatePortfoliosRebalancingStrategy(ctx, operations.CreatePortfoliosRebalancingStrategyRequest{
        IdempotencyKey: "ccb07f42-4104-44ad-8e1f-c660bb7b269c",
        Signature: "<value>",
        SignatureInput: "<value>",
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
| `request`                                                                                                                          | [operations.CreatePortfoliosRebalancingStrategyRequest](../../pkg/models/operations/createportfoliosrebalancingstrategyrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |


### Response

**[*operations.CreatePortfoliosRebalancingStrategyResponse](../../pkg/models/operations/createportfoliosrebalancingstrategyresponse.md), error**
| Error Object                                       | Status Code                                        | Content Type                                       |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| sdkerrors.CreatePortfoliosRebalancingStrategyError | 400,401,403,404,406,429,500,503,504                | application/problem+json                           |
| sdkerrors.SDKError                                 | 4xx-5xx                                            | */*                                                |

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Portfolios.ListPortfolioRebalancingExecutionOrders(ctx, operations.ListPortfolioRebalancingExecutionOrdersRequest{
        ExecutionID: "faccb01b-05e2-49df-ba2f-94e57ee7faa0",
        Signature: "<value>",
        SignatureInput: "<value>",
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

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.ListPortfolioRebalancingExecutionOrdersRequest](../../pkg/models/operations/listportfoliorebalancingexecutionordersrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |


### Response

**[*operations.ListPortfolioRebalancingExecutionOrdersResponse](../../pkg/models/operations/listportfoliorebalancingexecutionordersresponse.md), error**
| Error Object                                                     | Status Code                                                      | Content Type                                                     |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| sdkerrors.ListPortfolioRebalancingExecutionOrdersError           | 400,401,403,404,406,429,500,503,504                              | application/problem+json                                         |
| sdkerrors.ListPortfolioRebalancingExecutionOrdersPortfoliosError | 405                                                              | application/problem+json                                         |
| sdkerrors.SDKError                                               | 4xx-5xx                                                          | */*                                                              |

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Portfolios.ListPortfoliosAllocationAccounts(ctx, operations.ListPortfoliosAllocationAccountsRequest{
        AllocationID: "9260907b-4e94-45c8-a31a-10d01d471a66",
        Signature: "<value>",
        SignatureInput: "<value>",
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

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.ListPortfoliosAllocationAccountsRequest](../../pkg/models/operations/listportfoliosallocationaccountsrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.ListPortfoliosAllocationAccountsResponse](../../pkg/models/operations/listportfoliosallocationaccountsresponse.md), error**
| Error Object                                              | Status Code                                               | Content Type                                              |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| sdkerrors.ListPortfoliosAllocationAccountsError           | 400,401,403,404,406,429,500,503,504                       | application/problem+json                                  |
| sdkerrors.ListPortfoliosAllocationAccountsPortfoliosError | 405                                                       | application/problem+json                                  |
| sdkerrors.SDKError                                        | 4xx-5xx                                                   | */*                                                       |

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Portfolios.ListPortfoliosAllocations(ctx, operations.ListPortfoliosAllocationsRequest{
        Signature: "<value>",
        SignatureInput: "<value>",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.ListPortfoliosAllocationsRequest](../../pkg/models/operations/listportfoliosallocationsrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.ListPortfoliosAllocationsResponse](../../pkg/models/operations/listportfoliosallocationsresponse.md), error**
| Error Object                                       | Status Code                                        | Content Type                                       |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| sdkerrors.ListPortfoliosAllocationsError           | 400,401,403,404,406,429,500,503,504                | application/problem+json                           |
| sdkerrors.ListPortfoliosAllocationsPortfoliosError | 405                                                | application/problem+json                           |
| sdkerrors.SDKError                                 | 4xx-5xx                                            | */*                                                |

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Portfolios.ListPortfoliosConfigurations(ctx, operations.ListPortfoliosConfigurationsRequest{
        Signature: "<value>",
        SignatureInput: "<value>",
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

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.ListPortfoliosConfigurationsRequest](../../pkg/models/operations/listportfoliosconfigurationsrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.ListPortfoliosConfigurationsResponse](../../pkg/models/operations/listportfoliosconfigurationsresponse.md), error**
| Error Object                                          | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| sdkerrors.ListPortfoliosConfigurationsError           | 400,401,403,404,406,429,500,503,504                   | application/problem+json                              |
| sdkerrors.ListPortfoliosConfigurationsPortfoliosError | 405                                                   | application/problem+json                              |
| sdkerrors.SDKError                                    | 4xx-5xx                                               | */*                                                   |

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Portfolios.ListPortfoliosOrders(ctx, operations.ListPortfoliosOrdersRequest{
        Signature: "<value>",
        SignatureInput: "<value>",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListPortfoliosOrdersRequest](../../pkg/models/operations/listportfoliosordersrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.ListPortfoliosOrdersResponse](../../pkg/models/operations/listportfoliosordersresponse.md), error**
| Error Object                                  | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| sdkerrors.ListPortfoliosOrdersError           | 400,401,403,404,406,429,500,503,504           | application/problem+json                      |
| sdkerrors.ListPortfoliosOrdersPortfoliosError | 405                                           | application/problem+json                      |
| sdkerrors.SDKError                            | 4xx-5xx                                       | */*                                           |

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Portfolios.ListPortfoliosRebalancingStrategies(ctx, operations.ListPortfoliosRebalancingStrategiesRequest{
        Signature: "<value>",
        SignatureInput: "<value>",
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

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.ListPortfoliosRebalancingStrategiesRequest](../../pkg/models/operations/listportfoliosrebalancingstrategiesrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |


### Response

**[*operations.ListPortfoliosRebalancingStrategiesResponse](../../pkg/models/operations/listportfoliosrebalancingstrategiesresponse.md), error**
| Error Object                                                 | Status Code                                                  | Content Type                                                 |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| sdkerrors.ListPortfoliosRebalancingStrategiesError           | 400,401,403,404,406,429,500,503,504                          | application/problem+json                                     |
| sdkerrors.ListPortfoliosRebalancingStrategiesPortfoliosError | 405                                                          | application/problem+json                                     |
| sdkerrors.SDKError                                           | 4xx-5xx                                                      | */*                                                          |

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Portfolios.RetrievePortfoliosAllocation(ctx, operations.RetrievePortfoliosAllocationRequest{
        AllocationID: "8902e6a1-152d-44ee-8ba2-1f95e37ca0c1",
        Signature: "<value>",
        SignatureInput: "<value>",
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

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.RetrievePortfoliosAllocationRequest](../../pkg/models/operations/retrieveportfoliosallocationrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.RetrievePortfoliosAllocationResponse](../../pkg/models/operations/retrieveportfoliosallocationresponse.md), error**
| Error Object                                | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| sdkerrors.RetrievePortfoliosAllocationError | 401,403,404,406,429,500,503,504             | application/problem+json                    |
| sdkerrors.SDKError                          | 4xx-5xx                                     | */*                                         |

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Portfolios.RetrievePortfoliosConfiguration(ctx, operations.RetrievePortfoliosConfigurationRequest{
        AccountID: "823e6a8e-d630-4a74-9292-5b3ae05f7ca5",
        Signature: "<value>",
        SignatureInput: "<value>",
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.RetrievePortfoliosConfigurationRequest](../../pkg/models/operations/retrieveportfoliosconfigurationrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.RetrievePortfoliosConfigurationResponse](../../pkg/models/operations/retrieveportfoliosconfigurationresponse.md), error**
| Error Object                                             | Status Code                                              | Content Type                                             |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| sdkerrors.RetrievePortfoliosConfigurationError           | 401,403,404,406,429,500,503,504                          | application/problem+json                                 |
| sdkerrors.RetrievePortfoliosConfigurationPortfoliosError | 405                                                      | application/problem+json                                 |
| sdkerrors.SDKError                                       | 4xx-5xx                                                  | */*                                                      |

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Portfolios.RetrievePortfoliosOrder(ctx, operations.RetrievePortfoliosOrderRequest{
        PortfolioOrderID: "e8307df3-b75a-46de-a05a-9fe332308857",
        Signature: "<value>",
        SignatureInput: "<value>",
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.RetrievePortfoliosOrderRequest](../../pkg/models/operations/retrieveportfoliosorderrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.RetrievePortfoliosOrderResponse](../../pkg/models/operations/retrieveportfoliosorderresponse.md), error**
| Error Object                           | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.RetrievePortfoliosOrderError | 401,403,404,406,429,500,503,504        | application/problem+json               |
| sdkerrors.SDKError                     | 4xx-5xx                                | */*                                    |

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Portfolios.RetrievePortfoliosRebalancingExecution(ctx, operations.RetrievePortfoliosRebalancingExecutionRequest{
        ExecutionID: "99606d30-438e-43e3-948b-6cd364997761",
        Signature: "<value>",
        SignatureInput: "<value>",
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

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.RetrievePortfoliosRebalancingExecutionRequest](../../pkg/models/operations/retrieveportfoliosrebalancingexecutionrequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |


### Response

**[*operations.RetrievePortfoliosRebalancingExecutionResponse](../../pkg/models/operations/retrieveportfoliosrebalancingexecutionresponse.md), error**
| Error Object                                                    | Status Code                                                     | Content Type                                                    |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| sdkerrors.RetrievePortfoliosRebalancingExecutionError           | 401,403,404,406,429,500,503,504                                 | application/problem+json                                        |
| sdkerrors.RetrievePortfoliosRebalancingExecutionPortfoliosError | 405                                                             | application/problem+json                                        |
| sdkerrors.SDKError                                              | 4xx-5xx                                                         | */*                                                             |

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Portfolios.RetrievePortfoliosRebalancingStrategy(ctx, operations.RetrievePortfoliosRebalancingStrategyRequest{
        Signature: "<value>",
        SignatureInput: "<value>",
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

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.RetrievePortfoliosRebalancingStrategyRequest](../../pkg/models/operations/retrieveportfoliosrebalancingstrategyrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |


### Response

**[*operations.RetrievePortfoliosRebalancingStrategyResponse](../../pkg/models/operations/retrieveportfoliosrebalancingstrategyresponse.md), error**
| Error Object                                                   | Status Code                                                    | Content Type                                                   |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| sdkerrors.RetrievePortfoliosRebalancingStrategyError           | 401,403,404,406,429,500,503,504                                | application/problem+json                                       |
| sdkerrors.RetrievePortfoliosRebalancingStrategyPortfoliosError | 405                                                            | application/problem+json                                       |
| sdkerrors.SDKError                                             | 4xx-5xx                                                        | */*                                                            |

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Portfolios.TriggerPortfolioRebalancing(ctx, operations.TriggerPortfolioRebalancingRequest{
        IdempotencyKey: "ccb07f42-4104-44ad-8e1f-c660bb7b269c",
        Signature: "<value>",
        SignatureInput: "<value>",
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.TriggerPortfolioRebalancingRequest](../../pkg/models/operations/triggerportfoliorebalancingrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.TriggerPortfolioRebalancingResponse](../../pkg/models/operations/triggerportfoliorebalancingresponse.md), error**
| Error Object                               | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.TriggerPortfolioRebalancingError | 400,401,403,404,406,429,500,503,504        | application/problem+json                   |
| sdkerrors.SDKError                         | 4xx-5xx                                    | */*                                        |

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Portfolios.UpdatePortfoliosAllocation(ctx, operations.UpdatePortfoliosAllocationRequest{
        AllocationID: "c7ddcc6d-56bb-41c2-a344-b57fb4000627",
        Signature: "<value>",
        SignatureInput: "<value>",
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
| `request`                                                                                                        | [operations.UpdatePortfoliosAllocationRequest](../../pkg/models/operations/updateportfoliosallocationrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.UpdatePortfoliosAllocationResponse](../../pkg/models/operations/updateportfoliosallocationresponse.md), error**
| Error Object                              | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| sdkerrors.UpdatePortfoliosAllocationError | 400,401,403,404,406,429,500,503,504       | application/problem+json                  |
| sdkerrors.SDKError                        | 4xx-5xx                                   | */*                                       |

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Portfolios.UpdatePortfoliosConfiguration(ctx, operations.UpdatePortfoliosConfigurationRequest{
        AccountID: "be034715-bf00-452a-b6f6-ba5a9dfa7ad7",
        IdempotencyKey: "ccb07f42-4104-44ad-8e1f-c660bb7b269c",
        Signature: "<value>",
        SignatureInput: "<value>",
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
| `request`                                                                                                              | [operations.UpdatePortfoliosConfigurationRequest](../../pkg/models/operations/updateportfoliosconfigurationrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.UpdatePortfoliosConfigurationResponse](../../pkg/models/operations/updateportfoliosconfigurationresponse.md), error**
| Error Object                                 | Status Code                                  | Content Type                                 |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| sdkerrors.UpdatePortfoliosConfigurationError | 400,401,403,404,406,429,500,503,504          | application/problem+json                     |
| sdkerrors.SDKError                           | 4xx-5xx                                      | */*                                          |
