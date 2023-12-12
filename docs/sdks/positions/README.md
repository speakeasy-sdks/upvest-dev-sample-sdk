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
    res, err := s.Positions.ListPositions(ctx, operations.ListPositionsRequest{
        AccountID: "ff0795ec-4572-4135-9b42-4822adc99bc7",
        Signature: "string",
        SignatureInput: "string",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListPositionsRequest](../../pkg/models/operations/listpositionsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.ListPositionsResponse](../../pkg/models/operations/listpositionsresponse.md), error**
| Error Object                                     | Status Code                                      | Content Type                                     |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| sdkerrors.ListPositionsError                     | 400                                              | application/problem+json                         |
| sdkerrors.ListPositionsPositionsError            | 401                                              | application/problem+json                         |
| sdkerrors.ListPositionsPositionsResponseError    | 403                                              | application/problem+json                         |
| sdkerrors.ListPositionsPositionsResponse404Error | 404                                              | application/problem+json                         |
| sdkerrors.ListPositionsPositionsResponse406Error | 406                                              | application/problem+json                         |
| sdkerrors.ListPositionsPositionsResponse429Error | 429                                              | application/problem+json                         |
| sdkerrors.ListPositionsPositionsResponse500Error | 500                                              | application/problem+json                         |
| sdkerrors.ListPositionsPositionsResponse503Error | 503                                              | application/problem+json                         |
| sdkerrors.ListPositionsPositionsResponse504Error | 504                                              | application/problem+json                         |
| sdkerrors.SDKError                               | 400-600                                          | */*                                              |

## RetrievePosition

Retrieve an account position

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
    res, err := s.Positions.RetrievePosition(ctx, operations.RetrievePositionRequest{
        AccountID: "a4c2b2f9-c5c3-4199-ab2d-18703af419c4",
        InstrumentID: "string",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.RetrievePositionRequest](../../pkg/models/operations/retrievepositionrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.RetrievePositionResponse](../../pkg/models/operations/retrievepositionresponse.md), error**
| Error Object                                        | Status Code                                         | Content Type                                        |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| sdkerrors.RetrievePositionError                     | 401                                                 | application/problem+json                            |
| sdkerrors.RetrievePositionPositionsError            | 403                                                 | application/problem+json                            |
| sdkerrors.RetrievePositionPositionsResponseError    | 404                                                 | application/problem+json                            |
| sdkerrors.RetrievePositionPositionsResponse406Error | 406                                                 | application/problem+json                            |
| sdkerrors.RetrievePositionPositionsResponse429Error | 429                                                 | application/problem+json                            |
| sdkerrors.RetrievePositionPositionsResponse500Error | 500                                                 | application/problem+json                            |
| sdkerrors.RetrievePositionPositionsResponse503Error | 503                                                 | application/problem+json                            |
| sdkerrors.RetrievePositionPositionsResponse504Error | 504                                                 | application/problem+json                            |
| sdkerrors.SDKError                                  | 400-600                                             | */*                                                 |
