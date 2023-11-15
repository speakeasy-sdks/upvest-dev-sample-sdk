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
    s := upvestdevsamplesdk.New()

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

    if res.TwoHundredApplicationJSONTaxResidencyRecord != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.RetrieveTaxResidenciesRequest](../../pkg/models/operations/retrievetaxresidenciesrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.RetrieveTaxResidenciesResponse](../../pkg/models/operations/retrievetaxresidenciesresponse.md), error**
| Error Object                                          | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| sdkerrors.RetrieveTaxResidenciesError                 | 401                                                   | application/problem+json                              |
| sdkerrors.RetrieveTaxResidenciesTaxesError            | 403                                                   | application/problem+json                              |
| sdkerrors.RetrieveTaxResidenciesTaxesResponseError    | 404                                                   | application/problem+json                              |
| sdkerrors.RetrieveTaxResidenciesTaxesResponse405Error | 405                                                   | application/problem+json                              |
| sdkerrors.RetrieveTaxResidenciesTaxesResponse406Error | 406                                                   | application/problem+json                              |
| sdkerrors.RetrieveTaxResidenciesTaxesResponse429Error | 429                                                   | application/problem+json                              |
| sdkerrors.RetrieveTaxResidenciesTaxesResponse500Error | 500                                                   | application/problem+json                              |
| sdkerrors.RetrieveTaxResidenciesTaxesResponse503Error | 503                                                   | application/problem+json                              |
| sdkerrors.RetrieveTaxResidenciesTaxesResponse504Error | 504                                                   | application/problem+json                              |
| sdkerrors.SDKError                                    | 400-600                                               | */*                                                   |

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
    s := upvestdevsamplesdk.New()

    ctx := context.Background()
    res, err := s.Taxes.SetTaxResidencies(ctx, operations.SetTaxResidenciesRequest{
        RequestBody: &operations.SetTaxResidenciesTaxResidenciesSetRequest{
            TaxResidencies: []operations.TaxResidencyForCreateRequest{
                operations.CreateTaxResidencyForCreateRequestWithTaxIdentifierNumber(
                    operations.WithTaxIdentifierNumber{
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

    if res.TwoHundredApplicationJSONTaxResidencyRecord != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.SetTaxResidenciesRequest](../../pkg/models/operations/settaxresidenciesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.SetTaxResidenciesResponse](../../pkg/models/operations/settaxresidenciesresponse.md), error**
| Error Object                                     | Status Code                                      | Content Type                                     |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| sdkerrors.SetTaxResidenciesError                 | 400                                              | application/problem+json                         |
| sdkerrors.SetTaxResidenciesTaxesError            | 401                                              | application/problem+json                         |
| sdkerrors.SetTaxResidenciesTaxesResponseError    | 403                                              | application/problem+json                         |
| sdkerrors.SetTaxResidenciesTaxesResponse404Error | 404                                              | application/problem+json                         |
| sdkerrors.SetTaxResidenciesTaxesResponse406Error | 406                                              | application/problem+json                         |
| sdkerrors.SetTaxResidenciesTaxesResponse429Error | 429                                              | application/problem+json                         |
| sdkerrors.SetTaxResidenciesTaxesResponse500Error | 500                                              | application/problem+json                         |
| sdkerrors.SetTaxResidenciesTaxesResponse503Error | 503                                              | application/problem+json                         |
| sdkerrors.SetTaxResidenciesTaxesResponse504Error | 504                                              | application/problem+json                         |
| sdkerrors.SDKError                               | 400-600                                          | */*                                              |
