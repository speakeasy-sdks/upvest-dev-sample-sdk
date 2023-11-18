# Mandates
(*Mandates*)

## Overview

All direct debit mandates related paths

### Available Operations

* [CreateMandate](#createmandate) - Create a mandate
* [DeleteMandate](#deletemandate) - Delete mandate
* [ListMandates](#listmandates) - List mandates
* [RetrieveMandate](#retrievemandate) - Retrieve a direct debit mandate

## CreateMandate

Create a mandate

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	upvestdevsamplesdk "github.com/speakeasy-sdks/upvest-dev-sample-sdk"
	"context"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/types"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := upvestdevsamplesdk.New(
        upvestdevsamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Mandates.CreateMandate(ctx, operations.CreateMandateRequest{
        RequestBody: &operations.CreateMandateMandateCreateRequest{
            Bic: "string",
            ConfirmedAt: types.MustTimeFromString("2023-09-15T12:57:03.635Z"),
            Iban: "GT09333126I454209K17J3608U37",
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

    if res.DirectDebitMandate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateMandateRequest](../../pkg/models/operations/createmandaterequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.CreateMandateResponse](../../pkg/models/operations/createmandateresponse.md), error**
| Error Object                                    | Status Code                                     | Content Type                                    |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| sdkerrors.CreateMandateError                    | 400                                             | application/problem+json                        |
| sdkerrors.CreateMandateMandatesError            | 401                                             | application/problem+json                        |
| sdkerrors.CreateMandateMandatesResponseError    | 403                                             | application/problem+json                        |
| sdkerrors.CreateMandateMandatesResponse404Error | 404                                             | application/problem+json                        |
| sdkerrors.CreateMandateMandatesResponse406Error | 406                                             | application/problem+json                        |
| sdkerrors.CreateMandateMandatesResponse429Error | 429                                             | application/problem+json                        |
| sdkerrors.CreateMandateMandatesResponse500Error | 500                                             | application/problem+json                        |
| sdkerrors.CreateMandateMandatesResponse503Error | 503                                             | application/problem+json                        |
| sdkerrors.CreateMandateMandatesResponse504Error | 504                                             | application/problem+json                        |
| sdkerrors.SDKError                              | 400-600                                         | */*                                             |

## DeleteMandate

Delete mandate

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
    res, err := s.Mandates.DeleteMandate(ctx, operations.DeleteMandateRequest{
        MandateID: "6c46d0ca-673a-42ce-b95f-6db83c7daae3",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.DeleteMandateRequest](../../pkg/models/operations/deletemandaterequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.DeleteMandateResponse](../../pkg/models/operations/deletemandateresponse.md), error**
| Error Object                                    | Status Code                                     | Content Type                                    |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| sdkerrors.DeleteMandateError                    | 401                                             | application/problem+json                        |
| sdkerrors.DeleteMandateMandatesError            | 403                                             | application/problem+json                        |
| sdkerrors.DeleteMandateMandatesResponseError    | 404                                             | application/problem+json                        |
| sdkerrors.DeleteMandateMandatesResponse429Error | 429                                             | application/problem+json                        |
| sdkerrors.DeleteMandateMandatesResponse500Error | 500                                             | application/problem+json                        |
| sdkerrors.DeleteMandateMandatesResponse503Error | 503                                             | application/problem+json                        |
| sdkerrors.DeleteMandateMandatesResponse504Error | 504                                             | application/problem+json                        |
| sdkerrors.SDKError                              | 400-600                                         | */*                                             |

## ListMandates

List mandates

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
    res, err := s.Mandates.ListMandates(ctx, operations.ListMandatesRequest{
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "af42547b-fef5-4eb5-b512-a618cd73b2e3",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MandatesListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListMandatesRequest](../../pkg/models/operations/listmandatesrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.ListMandatesResponse](../../pkg/models/operations/listmandatesresponse.md), error**
| Error Object                                   | Status Code                                    | Content Type                                   |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| sdkerrors.ListMandatesError                    | 400                                            | application/problem+json                       |
| sdkerrors.ListMandatesMandatesError            | 401                                            | application/problem+json                       |
| sdkerrors.ListMandatesMandatesResponseError    | 403                                            | application/problem+json                       |
| sdkerrors.ListMandatesMandatesResponse404Error | 404                                            | application/problem+json                       |
| sdkerrors.ListMandatesMandatesResponse406Error | 406                                            | application/problem+json                       |
| sdkerrors.ListMandatesMandatesResponse429Error | 429                                            | application/problem+json                       |
| sdkerrors.ListMandatesMandatesResponse500Error | 500                                            | application/problem+json                       |
| sdkerrors.ListMandatesMandatesResponse503Error | 503                                            | application/problem+json                       |
| sdkerrors.ListMandatesMandatesResponse504Error | 504                                            | application/problem+json                       |
| sdkerrors.SDKError                             | 400-600                                        | */*                                            |

## RetrieveMandate

Retrieve a direct debit mandate

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
    res, err := s.Mandates.RetrieveMandate(ctx, operations.RetrieveMandateRequest{
        MandateID: "7fb0ccc3-5941-4c86-893a-aea9214922ce",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectDebitMandate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.RetrieveMandateRequest](../../pkg/models/operations/retrievemandaterequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.RetrieveMandateResponse](../../pkg/models/operations/retrievemandateresponse.md), error**
| Error Object                                      | Status Code                                       | Content Type                                      |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| sdkerrors.RetrieveMandateError                    | 401                                               | application/problem+json                          |
| sdkerrors.RetrieveMandateMandatesError            | 403                                               | application/problem+json                          |
| sdkerrors.RetrieveMandateMandatesResponseError    | 404                                               | application/problem+json                          |
| sdkerrors.RetrieveMandateMandatesResponse406Error | 406                                               | application/problem+json                          |
| sdkerrors.RetrieveMandateMandatesResponse429Error | 429                                               | application/problem+json                          |
| sdkerrors.RetrieveMandateMandatesResponse500Error | 500                                               | application/problem+json                          |
| sdkerrors.RetrieveMandateMandatesResponse503Error | 503                                               | application/problem+json                          |
| sdkerrors.RetrieveMandateMandatesResponse504Error | 504                                               | application/problem+json                          |
| sdkerrors.SDKError                                | 400-600                                           | */*                                               |
