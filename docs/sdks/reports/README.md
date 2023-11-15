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
	"context"
	"log"
	upvestdevsamplesdk "github.com/speakeasy-sdks/upvest-dev-sample-sdk"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
)

func main() {
    s := upvestdevsamplesdk.New()

    ctx := context.Background()
    res, err := s.Reports.ListReports(ctx, operations.ListReportsRequest{
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "7c866e4d-4e33-41a7-889e-0c03d98591db",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONReportsListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListReportsRequest](../../pkg/models/operations/listreportsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.ListReportsResponse](../../pkg/models/operations/listreportsresponse.md), error**
| Error Object                                 | Status Code                                  | Content Type                                 |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| sdkerrors.ListReportsError                   | 400                                          | application/problem+json                     |
| sdkerrors.ListReportsReportsError            | 401                                          | application/problem+json                     |
| sdkerrors.ListReportsReportsResponseError    | 403                                          | application/problem+json                     |
| sdkerrors.ListReportsReportsResponse404Error | 404                                          | application/problem+json                     |
| sdkerrors.ListReportsReportsResponse406Error | 406                                          | application/problem+json                     |
| sdkerrors.ListReportsReportsResponse429Error | 429                                          | application/problem+json                     |
| sdkerrors.ListReportsReportsResponse500Error | 500                                          | application/problem+json                     |
| sdkerrors.ListReportsReportsResponse503Error | 503                                          | application/problem+json                     |
| sdkerrors.ListReportsReportsResponse504Error | 504                                          | application/problem+json                     |
| sdkerrors.SDKError                           | 400-600                                      | */*                                          |

## RetrieveReport

Retrieve a user report

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
    res, err := s.Reports.RetrieveReport(ctx, operations.RetrieveReportRequest{
        ReportID: "008164dd-ebf0-4325-9e09-780b39dc8ee3",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONReport != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.RetrieveReportRequest](../../pkg/models/operations/retrievereportrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.RetrieveReportResponse](../../pkg/models/operations/retrievereportresponse.md), error**
| Error Object                                    | Status Code                                     | Content Type                                    |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| sdkerrors.RetrieveReportError                   | 401                                             | application/problem+json                        |
| sdkerrors.RetrieveReportReportsError            | 403                                             | application/problem+json                        |
| sdkerrors.RetrieveReportReportsResponseError    | 404                                             | application/problem+json                        |
| sdkerrors.RetrieveReportReportsResponse406Error | 406                                             | application/problem+json                        |
| sdkerrors.RetrieveReportReportsResponse429Error | 429                                             | application/problem+json                        |
| sdkerrors.RetrieveReportReportsResponse500Error | 500                                             | application/problem+json                        |
| sdkerrors.RetrieveReportReportsResponse503Error | 503                                             | application/problem+json                        |
| sdkerrors.RetrieveReportReportsResponse504Error | 504                                             | application/problem+json                        |
| sdkerrors.SDKError                              | 400-600                                         | */*                                             |
