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
	upvestdevsamplesdk "github.com/speakeasy-sdks/upvest-dev-sample-sdk"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
	"context"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"log"
)

func main() {
    s := upvestdevsamplesdk.New()


    operationSecurity := operations.RetrieveTaxResidenciesSecurity{
            OauthClientCredentials: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
        }

    ctx := context.Background()
    res, err := s.Taxes.RetrieveTaxResidencies(ctx, operations.RetrieveTaxResidenciesRequest{
        Signature: "<value>",
        SignatureInput: "<value>",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "75c9448a-b642-4df1-885c-88acb31d183e",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TaxResidencyRecord != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.RetrieveTaxResidenciesRequest](../../pkg/models/operations/retrievetaxresidenciesrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.RetrieveTaxResidenciesSecurity](../../pkg/models/operations/retrievetaxresidenciessecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


### Response

**[*operations.RetrieveTaxResidenciesResponse](../../pkg/models/operations/retrievetaxresidenciesresponse.md), error**
| Error Object                               | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.RetrieveTaxResidenciesError      | 401,403,404,406,429,500,503,504            | application/problem+json                   |
| sdkerrors.RetrieveTaxResidenciesTaxesError | 405                                        | application/problem+json                   |
| sdkerrors.SDKError                         | 4xx-5xx                                    | */*                                        |

## SetTaxResidencies

Update tax residencies

### Example Usage

```go
package main

import(
	upvestdevsamplesdk "github.com/speakeasy-sdks/upvest-dev-sample-sdk"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
	"context"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"log"
)

func main() {
    s := upvestdevsamplesdk.New()


    operationSecurity := operations.SetTaxResidenciesSecurity{
            OauthClientCredentials: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
        }

    ctx := context.Background()
    res, err := s.Taxes.SetTaxResidencies(ctx, operations.SetTaxResidenciesRequest{
        IdempotencyKey: "ccb07f42-4104-44ad-8e1f-c660bb7b269c",
        Signature: "<value>",
        SignatureInput: "<value>",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "2bbfa5f7-3a67-4795-aa9c-c38a53af8e19",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TaxResidencyRecord != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.SetTaxResidenciesRequest](../../pkg/models/operations/settaxresidenciesrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.SetTaxResidenciesSecurity](../../pkg/models/operations/settaxresidenciessecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.SetTaxResidenciesResponse](../../pkg/models/operations/settaxresidenciesresponse.md), error**
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.SetTaxResidenciesError    | 400,401,403,404,406,429,500,503,504 | application/problem+json            |
| sdkerrors.SDKError                  | 4xx-5xx                             | */*                                 |
