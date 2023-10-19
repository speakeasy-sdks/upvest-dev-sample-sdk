# Positions
(*Positions*)

## Overview

All positions related paths.

### Available Operations

* [ListPositions](#listpositions) - List positions
* [RetrievePosition](#retrieveposition) - Retrieve position

## ListPositions

List of account positions

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
    res, err := s.Positions.ListPositions(ctx, operations.ListPositionsRequest{
        AccountID: "ff0795ec-4572-4135-9b42-4822adc99bc7",
        Signature: "whiteboard",
        SignatureInput: "aha",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PositionsListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListPositionsRequest](../../models/operations/listpositionsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.ListPositionsResponse](../../models/operations/listpositionsresponse.md), error**


## RetrievePosition

Retrieve an account position

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
    res, err := s.Positions.RetrievePosition(ctx, operations.RetrievePositionRequest{
        AccountID: "a4c2b2f9-c5c3-4199-ab2d-18703af419c4",
        InstrumentID: "Myanmar",
        Signature: "orange",
        SignatureInput: "Savings",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RetrievePosition200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.RetrievePositionRequest](../../models/operations/retrievepositionrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.RetrievePositionResponse](../../models/operations/retrievepositionresponse.md), error**

