# Reports
(*Reports*)

## Overview

All reports related paths.

### Available Operations

* [ListReports](#listreports) - List user reports
* [RetrieveReport](#retrievereport) - Retrieve a user report

## ListReports

List user reports

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


    operationSecurity := operations.ListReportsSecurity{
            OauthClientCredentials: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
        }

    ctx := context.Background()
    res, err := s.Reports.ListReports(ctx, operations.ListReportsRequest{
        Signature: "<value>",
        SignatureInput: "<value>",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "7c866e4d-4e33-41a7-889e-0c03d98591db",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }
    if res.ReportsListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListReportsRequest](../../pkg/models/operations/listreportsrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.ListReportsSecurity](../../pkg/models/operations/listreportssecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.ListReportsResponse](../../pkg/models/operations/listreportsresponse.md), error**
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.ListReportsError          | 400,401,403,404,406,429,500,503,504 | application/problem+json            |
| sdkerrors.SDKError                  | 4xx-5xx                             | */*                                 |

## RetrieveReport

Retrieve a user report

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


    operationSecurity := operations.RetrieveReportSecurity{
            OauthClientCredentials: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
        }

    ctx := context.Background()
    res, err := s.Reports.RetrieveReport(ctx, operations.RetrieveReportRequest{
        ReportID: "008164dd-ebf0-4325-9e09-780b39dc8ee3",
        Signature: "<value>",
        SignatureInput: "<value>",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }
    if res.TwoHundredApplicationJSONReport != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.RetrieveReportRequest](../../pkg/models/operations/retrievereportrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.RetrieveReportSecurity](../../pkg/models/operations/retrievereportsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.RetrieveReportResponse](../../pkg/models/operations/retrievereportresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.RetrieveReportError   | 401,403,404,406,429,500,503,504 | application/problem+json        |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |
