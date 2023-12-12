# Orders
(*Orders*)

## Overview

All order related paths.

### Available Operations

* [CancelOrder](#cancelorder) - Cancel an order by ID
* [ListAccountOrders](#listaccountorders) - Get orders for an account by ID
* [PlaceOrder](#placeorder) - Place an order
* [RetrieveOrder](#retrieveorder) - Get an order by ID
* [RetrieveOrderExecution](#retrieveorderexecution) - Get an order execution by ID

## CancelOrder

Cancels an order specified by its ID. It is possible to cancel an order of the `NEW` status. Once a cancellation has been accepted, the further processing steps take place asynchronously and depending on the respective order status.

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
    res, err := s.Orders.CancelOrder(ctx, operations.CancelOrderRequest{
        OrderID: "8d2637b8-60e3-49d2-9ef3-bc3a907bc7fc",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OrderCancelResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CancelOrderRequest](../../pkg/models/operations/cancelorderrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.CancelOrderResponse](../../pkg/models/operations/cancelorderresponse.md), error**
| Error Object                                | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| sdkerrors.CancelOrderError                  | 401                                         | application/problem+json                    |
| sdkerrors.CancelOrderOrdersError            | 403                                         | application/problem+json                    |
| sdkerrors.CancelOrderOrdersResponseError    | 404                                         | application/problem+json                    |
| sdkerrors.CancelOrderOrdersResponse406Error | 406                                         | application/problem+json                    |
| sdkerrors.CancelOrderOrdersResponse429Error | 429                                         | application/problem+json                    |
| sdkerrors.CancelOrderOrdersResponse500Error | 500                                         | application/problem+json                    |
| sdkerrors.CancelOrderOrdersResponse503Error | 503                                         | application/problem+json                    |
| sdkerrors.CancelOrderOrdersResponse504Error | 504                                         | application/problem+json                    |
| sdkerrors.SDKError                          | 400-600                                     | */*                                         |

## ListAccountOrders

Returns a list of all orders for the account specified by its ID.

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
    res, err := s.Orders.ListAccountOrders(ctx, operations.ListAccountOrdersRequest{
        AccountID: "ad45b27f-c0e7-4cfb-b48c-a83670dbdfbd",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OrdersListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListAccountOrdersRequest](../../pkg/models/operations/listaccountordersrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.ListAccountOrdersResponse](../../pkg/models/operations/listaccountordersresponse.md), error**
| Error Object                                      | Status Code                                       | Content Type                                      |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| sdkerrors.ListAccountOrdersError                  | 400                                               | application/problem+json                          |
| sdkerrors.ListAccountOrdersOrdersError            | 401                                               | application/problem+json                          |
| sdkerrors.ListAccountOrdersOrdersResponseError    | 403                                               | application/problem+json                          |
| sdkerrors.ListAccountOrdersOrdersResponse404Error | 404                                               | application/problem+json                          |
| sdkerrors.ListAccountOrdersOrdersResponse406Error | 406                                               | application/problem+json                          |
| sdkerrors.ListAccountOrdersOrdersResponse429Error | 429                                               | application/problem+json                          |
| sdkerrors.ListAccountOrdersOrdersResponse500Error | 500                                               | application/problem+json                          |
| sdkerrors.ListAccountOrdersOrdersResponse503Error | 503                                               | application/problem+json                          |
| sdkerrors.ListAccountOrdersOrdersResponse504Error | 504                                               | application/problem+json                          |
| sdkerrors.SDKError                                | 400-600                                           | */*                                               |

## PlaceOrder

Places a new order. After the creation request for the order is accepted, further processing takes place asynchronously.

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
    res, err := s.Orders.PlaceOrder(ctx, operations.PlaceOrderRequest{
        RequestBody: &operations.PlaceOrderOrderPlaceRequest{
            AccountID: "b95bd99c-2bac-4393-b09f-a0cc26cb9bc2",
            InstrumentID: "string",
            Side: operations.SideBuy,
            UserID: "69a1c9be-906c-4213-b0cb-fa43f8d73fd7",
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

    if res.Order != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.PlaceOrderRequest](../../pkg/models/operations/placeorderrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.PlaceOrderResponse](../../pkg/models/operations/placeorderresponse.md), error**
| Error Object                               | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.PlaceOrderError                  | 400                                        | application/problem+json                   |
| sdkerrors.PlaceOrderOrdersError            | 401                                        | application/problem+json                   |
| sdkerrors.PlaceOrderOrdersResponseError    | 403                                        | application/problem+json                   |
| sdkerrors.PlaceOrderOrdersResponse406Error | 406                                        | application/problem+json                   |
| sdkerrors.PlaceOrderOrdersResponse422Error | 422                                        | application/problem+json                   |
| sdkerrors.PlaceOrderOrdersResponse429Error | 429                                        | application/problem+json                   |
| sdkerrors.PlaceOrderOrdersResponse500Error | 500                                        | application/problem+json                   |
| sdkerrors.PlaceOrderOrdersResponse503Error | 503                                        | application/problem+json                   |
| sdkerrors.PlaceOrderOrdersResponse504Error | 504                                        | application/problem+json                   |
| sdkerrors.SDKError                         | 400-600                                    | */*                                        |

## RetrieveOrder

Returns the order specified by its ID.

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
    res, err := s.Orders.RetrieveOrder(ctx, operations.RetrieveOrderRequest{
        OrderID: "6b222a05-ca10-4e5a-8de6-a4b10ac87ee1",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Order != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.RetrieveOrderRequest](../../pkg/models/operations/retrieveorderrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.RetrieveOrderResponse](../../pkg/models/operations/retrieveorderresponse.md), error**
| Error Object                                  | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| sdkerrors.RetrieveOrderError                  | 401                                           | application/problem+json                      |
| sdkerrors.RetrieveOrderOrdersError            | 403                                           | application/problem+json                      |
| sdkerrors.RetrieveOrderOrdersResponseError    | 404                                           | application/problem+json                      |
| sdkerrors.RetrieveOrderOrdersResponse406Error | 406                                           | application/problem+json                      |
| sdkerrors.RetrieveOrderOrdersResponse429Error | 429                                           | application/problem+json                      |
| sdkerrors.RetrieveOrderOrdersResponse500Error | 500                                           | application/problem+json                      |
| sdkerrors.RetrieveOrderOrdersResponse503Error | 503                                           | application/problem+json                      |
| sdkerrors.RetrieveOrderOrdersResponse504Error | 504                                           | application/problem+json                      |
| sdkerrors.SDKError                            | 400-600                                       | */*                                           |

## RetrieveOrderExecution

Returns the order execution specified by its ID.

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
    res, err := s.Orders.RetrieveOrderExecution(ctx, operations.RetrieveOrderExecutionRequest{
        ExecutionID: "469610b9-c442-4475-8f5b-e30a097443e4",
        OrderID: "9798733b-65ea-45d3-9976-592362ab8751",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OrderExecution != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.RetrieveOrderExecutionRequest](../../pkg/models/operations/retrieveorderexecutionrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.RetrieveOrderExecutionResponse](../../pkg/models/operations/retrieveorderexecutionresponse.md), error**
| Error Object                                           | Status Code                                            | Content Type                                           |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| sdkerrors.RetrieveOrderExecutionError                  | 401                                                    | application/problem+json                               |
| sdkerrors.RetrieveOrderExecutionOrdersError            | 403                                                    | application/problem+json                               |
| sdkerrors.RetrieveOrderExecutionOrdersResponseError    | 404                                                    | application/problem+json                               |
| sdkerrors.RetrieveOrderExecutionOrdersResponse406Error | 406                                                    | application/problem+json                               |
| sdkerrors.RetrieveOrderExecutionOrdersResponse429Error | 429                                                    | application/problem+json                               |
| sdkerrors.RetrieveOrderExecutionOrdersResponse500Error | 500                                                    | application/problem+json                               |
| sdkerrors.RetrieveOrderExecutionOrdersResponse503Error | 503                                                    | application/problem+json                               |
| sdkerrors.RetrieveOrderExecutionOrdersResponse504Error | 504                                                    | application/problem+json                               |
| sdkerrors.SDKError                                     | 400-600                                                | */*                                                    |
