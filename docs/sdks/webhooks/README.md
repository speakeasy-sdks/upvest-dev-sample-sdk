# Webhooks
(*Webhooks*)

## Overview

All webhook related paths.

### Available Operations

* [CreateWebhook](#createwebhook) - Create a webhook subscription
* [DeleteWebhook](#deletewebhook) - Delete a webhook subscription
* [GetJwks](#getjwks) - Get signing keys
* [ListWebhooks](#listwebhooks) - List all webhooks
* [RetrieveWebhook](#retrievewebhook) - Retrieve a webhook subscription
* [TestWebhook](#testwebhook) - Test a webhook
* [UpdateWebhook](#updatewebhook) - Update a webhook subscription

## CreateWebhook

Create a webhook subscription

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
    res, err := s.Webhooks.CreateWebhook(ctx, operations.CreateWebhookRequest{
        RequestBody: &operations.CreateWebhookWebhookCreateRequest{
            Config: &operations.Config{},
            Title: "string",
            Type: []operations.CreateWebhookType{
                operations.CreateWebhookTypeUser,
            },
            URL: "https://pointless-banner.org",
        },
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredAndOneApplicationJSONWebhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateWebhookRequest](../../pkg/models/operations/createwebhookrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.CreateWebhookResponse](../../pkg/models/operations/createwebhookresponse.md), error**
| Error Object                                    | Status Code                                     | Content Type                                    |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| sdkerrors.CreateWebhookError                    | 400                                             | application/problem+json                        |
| sdkerrors.CreateWebhookWebhooksError            | 401                                             | application/problem+json                        |
| sdkerrors.CreateWebhookWebhooksResponseError    | 403                                             | application/problem+json                        |
| sdkerrors.CreateWebhookWebhooksResponse406Error | 406                                             | application/problem+json                        |
| sdkerrors.CreateWebhookWebhooksResponse429Error | 429                                             | application/problem+json                        |
| sdkerrors.CreateWebhookWebhooksResponse500Error | 500                                             | application/problem+json                        |
| sdkerrors.CreateWebhookWebhooksResponse503Error | 503                                             | application/problem+json                        |
| sdkerrors.CreateWebhookWebhooksResponse504Error | 504                                             | application/problem+json                        |
| sdkerrors.SDKError                              | 400-600                                         | */*                                             |

## DeleteWebhook

Delete a webhook subscription

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
    res, err := s.Webhooks.DeleteWebhook(ctx, operations.DeleteWebhookRequest{
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        WebhookID: "5aa6d655-26dd-43c7-a4d9-800d9681eebf",
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
| `request`                                                                              | [operations.DeleteWebhookRequest](../../pkg/models/operations/deletewebhookrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.DeleteWebhookResponse](../../pkg/models/operations/deletewebhookresponse.md), error**
| Error Object                                    | Status Code                                     | Content Type                                    |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| sdkerrors.DeleteWebhookError                    | 401                                             | application/problem+json                        |
| sdkerrors.DeleteWebhookWebhooksError            | 403                                             | application/problem+json                        |
| sdkerrors.DeleteWebhookWebhooksResponseError    | 404                                             | application/problem+json                        |
| sdkerrors.DeleteWebhookWebhooksResponse429Error | 429                                             | application/problem+json                        |
| sdkerrors.DeleteWebhookWebhooksResponse500Error | 500                                             | application/problem+json                        |
| sdkerrors.DeleteWebhookWebhooksResponse503Error | 503                                             | application/problem+json                        |
| sdkerrors.DeleteWebhookWebhooksResponse504Error | 504                                             | application/problem+json                        |
| sdkerrors.SDKError                              | 400-600                                         | */*                                             |

## GetJwks

Get list of signing keys used to verify webhooks

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
    res, err := s.Webhooks.GetJwks(ctx, operations.GetJwksRequest{
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONAuthVerificationKeys != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetJwksRequest](../../pkg/models/operations/getjwksrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetJwksResponse](../../pkg/models/operations/getjwksresponse.md), error**
| Error Object                              | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| sdkerrors.GetJwksError                    | 401                                       | application/problem+json                  |
| sdkerrors.GetJwksWebhooksError            | 403                                       | application/problem+json                  |
| sdkerrors.GetJwksWebhooksResponseError    | 406                                       | application/problem+json                  |
| sdkerrors.GetJwksWebhooksResponse429Error | 429                                       | application/problem+json                  |
| sdkerrors.GetJwksWebhooksResponse500Error | 500                                       | application/problem+json                  |
| sdkerrors.GetJwksWebhooksResponse503Error | 503                                       | application/problem+json                  |
| sdkerrors.GetJwksWebhooksResponse504Error | 504                                       | application/problem+json                  |
| sdkerrors.SDKError                        | 400-600                                   | */*                                       |

## ListWebhooks

List all webhooks

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
    res, err := s.Webhooks.ListWebhooks(ctx, operations.ListWebhooksRequest{
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONWebhooksListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListWebhooksRequest](../../pkg/models/operations/listwebhooksrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.ListWebhooksResponse](../../pkg/models/operations/listwebhooksresponse.md), error**
| Error Object                                   | Status Code                                    | Content Type                                   |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| sdkerrors.ListWebhooksError                    | 400                                            | application/problem+json                       |
| sdkerrors.ListWebhooksWebhooksError            | 401                                            | application/problem+json                       |
| sdkerrors.ListWebhooksWebhooksResponseError    | 403                                            | application/problem+json                       |
| sdkerrors.ListWebhooksWebhooksResponse406Error | 406                                            | application/problem+json                       |
| sdkerrors.ListWebhooksWebhooksResponse429Error | 429                                            | application/problem+json                       |
| sdkerrors.ListWebhooksWebhooksResponse500Error | 500                                            | application/problem+json                       |
| sdkerrors.ListWebhooksWebhooksResponse503Error | 503                                            | application/problem+json                       |
| sdkerrors.ListWebhooksWebhooksResponse504Error | 504                                            | application/problem+json                       |
| sdkerrors.SDKError                             | 400-600                                        | */*                                            |

## RetrieveWebhook

Retrieve a webhook subscription

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
    res, err := s.Webhooks.RetrieveWebhook(ctx, operations.RetrieveWebhookRequest{
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        WebhookID: "1b96f803-b0a9-4178-8c24-38b86638f043",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONWebhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.RetrieveWebhookRequest](../../pkg/models/operations/retrievewebhookrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.RetrieveWebhookResponse](../../pkg/models/operations/retrievewebhookresponse.md), error**
| Error Object                                      | Status Code                                       | Content Type                                      |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| sdkerrors.RetrieveWebhookError                    | 401                                               | application/problem+json                          |
| sdkerrors.RetrieveWebhookWebhooksError            | 403                                               | application/problem+json                          |
| sdkerrors.RetrieveWebhookWebhooksResponseError    | 404                                               | application/problem+json                          |
| sdkerrors.RetrieveWebhookWebhooksResponse406Error | 406                                               | application/problem+json                          |
| sdkerrors.RetrieveWebhookWebhooksResponse429Error | 429                                               | application/problem+json                          |
| sdkerrors.RetrieveWebhookWebhooksResponse500Error | 500                                               | application/problem+json                          |
| sdkerrors.RetrieveWebhookWebhooksResponse503Error | 503                                               | application/problem+json                          |
| sdkerrors.RetrieveWebhookWebhooksResponse504Error | 504                                               | application/problem+json                          |
| sdkerrors.SDKError                                | 400-600                                           | */*                                               |

## TestWebhook

Test a webhook

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
    res, err := s.Webhooks.TestWebhook(ctx, operations.TestWebhookRequest{
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        WebhookID: "6b48b079-7b31-4425-968d-e3994a64058e",
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
| `request`                                                                          | [operations.TestWebhookRequest](../../pkg/models/operations/testwebhookrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.TestWebhookResponse](../../pkg/models/operations/testwebhookresponse.md), error**
| Error Object                                  | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| sdkerrors.TestWebhookError                    | 401                                           | application/problem+json                      |
| sdkerrors.TestWebhookWebhooksError            | 403                                           | application/problem+json                      |
| sdkerrors.TestWebhookWebhooksResponseError    | 404                                           | application/problem+json                      |
| sdkerrors.TestWebhookWebhooksResponse406Error | 406                                           | application/problem+json                      |
| sdkerrors.TestWebhookWebhooksResponse429Error | 429                                           | application/problem+json                      |
| sdkerrors.TestWebhookWebhooksResponse500Error | 500                                           | application/problem+json                      |
| sdkerrors.TestWebhookWebhooksResponse503Error | 503                                           | application/problem+json                      |
| sdkerrors.TestWebhookWebhooksResponse504Error | 504                                           | application/problem+json                      |
| sdkerrors.SDKError                            | 400-600                                       | */*                                           |

## UpdateWebhook

Update a webhook subscription

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
    res, err := s.Webhooks.UpdateWebhook(ctx, operations.UpdateWebhookRequest{
        RequestBody: &operations.UpdateWebhookWebhookUpdateRequest{
            Config: &operations.UpdateWebhookConfig{},
            Type: []operations.UpdateWebhookType{
                operations.UpdateWebhookTypeIntradayAccountValuation,
            },
        },
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        WebhookID: "f1f2b873-8bbd-4246-95a9-1c147ec420b7",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONWebhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateWebhookRequest](../../pkg/models/operations/updatewebhookrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.UpdateWebhookResponse](../../pkg/models/operations/updatewebhookresponse.md), error**
| Error Object                                    | Status Code                                     | Content Type                                    |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| sdkerrors.UpdateWebhookError                    | 400                                             | application/problem+json                        |
| sdkerrors.UpdateWebhookWebhooksError            | 401                                             | application/problem+json                        |
| sdkerrors.UpdateWebhookWebhooksResponseError    | 403                                             | application/problem+json                        |
| sdkerrors.UpdateWebhookWebhooksResponse404Error | 404                                             | application/problem+json                        |
| sdkerrors.UpdateWebhookWebhooksResponse406Error | 406                                             | application/problem+json                        |
| sdkerrors.UpdateWebhookWebhooksResponse429Error | 429                                             | application/problem+json                        |
| sdkerrors.UpdateWebhookWebhooksResponse500Error | 500                                             | application/problem+json                        |
| sdkerrors.UpdateWebhookWebhooksResponse503Error | 503                                             | application/problem+json                        |
| sdkerrors.UpdateWebhookWebhooksResponse504Error | 504                                             | application/problem+json                        |
| sdkerrors.SDKError                              | 400-600                                         | */*                                             |
