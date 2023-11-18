# Instruments
(*Instruments*)

## Overview

All instrument related paths.

### Available Operations

* [ListInstruments](#listinstruments) - List instruments
* [RetrieveInstrument](#retrieveinstrument) - Retrieve instrument

## ListInstruments

List instruments

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
    res, err := s.Instruments.ListInstruments(ctx, operations.ListInstrumentsRequest{
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.InstrumentsListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListInstrumentsRequest](../../pkg/models/operations/listinstrumentsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.ListInstrumentsResponse](../../pkg/models/operations/listinstrumentsresponse.md), error**
| Error Object                                         | Status Code                                          | Content Type                                         |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| sdkerrors.ListInstrumentsError                       | 400                                                  | application/problem+json                             |
| sdkerrors.ListInstrumentsInstrumentsError            | 401                                                  | application/problem+json                             |
| sdkerrors.ListInstrumentsInstrumentsResponseError    | 403                                                  | application/problem+json                             |
| sdkerrors.ListInstrumentsInstrumentsResponse406Error | 406                                                  | application/problem+json                             |
| sdkerrors.ListInstrumentsInstrumentsResponse429Error | 429                                                  | application/problem+json                             |
| sdkerrors.ListInstrumentsInstrumentsResponse500Error | 500                                                  | application/problem+json                             |
| sdkerrors.ListInstrumentsInstrumentsResponse503Error | 503                                                  | application/problem+json                             |
| sdkerrors.ListInstrumentsInstrumentsResponse504Error | 504                                                  | application/problem+json                             |
| sdkerrors.SDKError                                   | 400-600                                              | */*                                                  |

## RetrieveInstrument

Retrieve instrument

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
    res, err := s.Instruments.RetrieveInstrument(ctx, operations.RetrieveInstrumentRequest{
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.RetrieveInstrumentRequest](../../pkg/models/operations/retrieveinstrumentrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.RetrieveInstrumentResponse](../../pkg/models/operations/retrieveinstrumentresponse.md), error**
| Error Object                                            | Status Code                                             | Content Type                                            |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| sdkerrors.RetrieveInstrumentError                       | 401                                                     | application/problem+json                                |
| sdkerrors.RetrieveInstrumentInstrumentsError            | 403                                                     | application/problem+json                                |
| sdkerrors.RetrieveInstrumentInstrumentsResponseError    | 404                                                     | application/problem+json                                |
| sdkerrors.RetrieveInstrumentInstrumentsResponse406Error | 406                                                     | application/problem+json                                |
| sdkerrors.RetrieveInstrumentInstrumentsResponse429Error | 429                                                     | application/problem+json                                |
| sdkerrors.RetrieveInstrumentInstrumentsResponse500Error | 500                                                     | application/problem+json                                |
| sdkerrors.RetrieveInstrumentInstrumentsResponse503Error | 503                                                     | application/problem+json                                |
| sdkerrors.RetrieveInstrumentInstrumentsResponse504Error | 504                                                     | application/problem+json                                |
| sdkerrors.SDKError                                      | 400-600                                                 | */*                                                     |
