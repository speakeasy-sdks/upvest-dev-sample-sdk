# Taxes
(*Taxes*)

## Overview

All taxes related paths.

### Available Operations

* [RetrieveTaxResidencies](#retrievetaxresidencies) - Retrieve tax residencies
* [SetTaxResidencies](#settaxresidencies) - Update tax residencies

## RetrieveTaxResidencies

Retrieve tax residencies

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
        upvestdevsamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Taxes.RetrieveTaxResidencies(ctx, operations.RetrieveTaxResidenciesRequest{
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "75c9448a-b642-4df1-885c-88acb31d183e",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TaxResidencyRecord != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.RetrieveTaxResidenciesRequest](../../models/operations/retrievetaxresidenciesrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.RetrieveTaxResidenciesResponse](../../models/operations/retrievetaxresidenciesresponse.md), error**


## SetTaxResidencies

Update tax residencies

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
        upvestdevsamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Taxes.SetTaxResidencies(ctx, operations.SetTaxResidenciesRequest{
        RequestBody: &operations.SetTaxResidenciesTaxResidenciesSetRequest{
            TaxResidencies: []operations.SetTaxResidenciesTaxResidenciesSetRequestTaxResidencyForCreateRequest{
                operations.CreateSetTaxResidenciesTaxResidenciesSetRequestTaxResidencyForCreateRequestSetTaxResidenciesTaxResidenciesSetRequestTaxResidencyForCreateRequestWithTaxIdentifierNumber(
                    operations.SetTaxResidenciesTaxResidenciesSetRequestTaxResidencyForCreateRequestWithTaxIdentifierNumber{
                        Country: "Portugal",
                        TaxIdentifierNumber: "string",
                    },
                ),
            },
        },
        IdempotencyKey: "ccb07f42-4104-44ad-8e1f-c660bb7b269c",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "bfa5f73a-6779-45aa-9cc3-8a53af8e1995",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TaxResidencyRecord != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.SetTaxResidenciesRequest](../../models/operations/settaxresidenciesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.SetTaxResidenciesResponse](../../models/operations/settaxresidenciesresponse.md), error**

