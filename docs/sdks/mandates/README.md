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
	"context"
	"log"
	upvestdevsamplesdk "github.com/speakeasy-sdks/upvest-dev-sample-sdk"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/types"
)

func main() {
    s := upvestdevsamplesdk.New(
        upvestdevsamplesdk.WithSecurity("YOUR_TOKEN"),
    )

    ctx := context.Background()
    res, err := s.Mandates.CreateMandate(ctx, operations.CreateMandateRequest{
        RequestBody: &operations.CreateMandateMandateCreateRequest{
            Bic: "instead",
            ConfirmedAt: types.MustTimeFromString("2021-06-03T23:13:21.087Z"),
            Iban: "GL7533112009623524",
        },
        IdempotencyKey: "ccb07f42-4104-44ad-8e1f-c660bb7b269c",
        Signature: "Brookline",
        SignatureInput: "while",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreateMandateRequest](../../models/operations/createmandaterequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.CreateMandateResponse](../../models/operations/createmandateresponse.md), error**


## DeleteMandate

Delete mandate

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
    res, err := s.Mandates.DeleteMandate(ctx, operations.DeleteMandateRequest{
        MandateID: "6c46d0ca-673a-42ce-b95f-6db83c7daae3",
        Signature: "Tasty",
        SignatureInput: "hybrid",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.DeleteMandateRequest](../../models/operations/deletemandaterequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.DeleteMandateResponse](../../models/operations/deletemandateresponse.md), error**


## ListMandates

List mandates

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
    res, err := s.Mandates.ListMandates(ctx, operations.ListMandatesRequest{
        Signature: "Functionality",
        SignatureInput: "Sausages",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "f5eb5b51-2a61-48cd-b3b2-e39eb375a342",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListMandatesRequest](../../models/operations/listmandatesrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.ListMandatesResponse](../../models/operations/listmandatesresponse.md), error**


## RetrieveMandate

Retrieve a direct debit mandate

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
    res, err := s.Mandates.RetrieveMandate(ctx, operations.RetrieveMandateRequest{
        MandateID: "7fb0ccc3-5941-4c86-893a-aea9214922ce",
        Signature: "North",
        SignatureInput: "Loan",
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
| `request`                                                                              | [operations.RetrieveMandateRequest](../../models/operations/retrievemandaterequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.RetrieveMandateResponse](../../models/operations/retrievemandateresponse.md), error**

