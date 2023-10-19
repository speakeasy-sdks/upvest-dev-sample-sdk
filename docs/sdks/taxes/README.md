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
        Signature: "HDD",
        SignatureInput: "orchid",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "448ab642-df1c-485c-88ac-b31d183e94aa",
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
                        TaxIdentifierNumber: "katal",
                    },
                ),
            },
        },
        IdempotencyKey: "ccb07f42-4104-44ad-8e1f-c660bb7b269c",
        Signature: "National",
        SignatureInput: "jibe",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "3a67795a-a9cc-438a-93af-8e1995010f78",
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

