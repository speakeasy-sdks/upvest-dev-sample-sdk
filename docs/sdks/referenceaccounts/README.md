# ReferenceAccounts
(*ReferenceAccounts*)

## Overview

All reference account related paths

### Available Operations

* [CreateReferenceAccount](#createreferenceaccount) - Create a reference account
* [DeleteReferenceAccount](#deletereferenceaccount) - Delete a reference account by ID
* [ListReferenceAccounts](#listreferenceaccounts) - Get reference accounts of a user
* [RetrieveReferenceAccount](#retrievereferenceaccount) - Get a reference account by ID

## CreateReferenceAccount

Creates a new reference account for a user specified by ID.

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.ReferenceAccounts.CreateReferenceAccount(ctx, operations.CreateReferenceAccountRequest{
        RequestBody: &operations.CreateReferenceAccountReferenceAccountCreateRequest{
            AccountOwner: "string",
            Bic: "string",
            ConfirmedAt: types.MustTimeFromString("2021-10-21T09:53:29.074Z"),
            Iban: "CH82077325Y83934M284R",
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

    if res.ReferenceAccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.CreateReferenceAccountRequest](../../pkg/models/operations/createreferenceaccountrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.CreateReferenceAccountResponse](../../pkg/models/operations/createreferenceaccountresponse.md), error**
| Error Object                                                      | Status Code                                                       | Content Type                                                      |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| sdkerrors.CreateReferenceAccountError                             | 400                                                               | application/problem+json                                          |
| sdkerrors.CreateReferenceAccountReferenceAccountsError            | 401                                                               | application/problem+json                                          |
| sdkerrors.CreateReferenceAccountReferenceAccountsResponseError    | 403                                                               | application/problem+json                                          |
| sdkerrors.CreateReferenceAccountReferenceAccountsResponse404Error | 404                                                               | application/problem+json                                          |
| sdkerrors.CreateReferenceAccountReferenceAccountsResponse406Error | 406                                                               | application/problem+json                                          |
| sdkerrors.CreateReferenceAccountReferenceAccountsResponse429Error | 429                                                               | application/problem+json                                          |
| sdkerrors.CreateReferenceAccountReferenceAccountsResponse500Error | 500                                                               | application/problem+json                                          |
| sdkerrors.CreateReferenceAccountReferenceAccountsResponse503Error | 503                                                               | application/problem+json                                          |
| sdkerrors.CreateReferenceAccountReferenceAccountsResponse504Error | 504                                                               | application/problem+json                                          |
| sdkerrors.SDKError                                                | 400-600                                                           | */*                                                               |

## DeleteReferenceAccount

Deletes the reference account specified by its ID.

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.ReferenceAccounts.DeleteReferenceAccount(ctx, operations.DeleteReferenceAccountRequest{
        ReferenceAccountID: "b658e83d-389d-47da-aa73-769e9336c8e7",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.DeleteReferenceAccountRequest](../../pkg/models/operations/deletereferenceaccountrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.DeleteReferenceAccountResponse](../../pkg/models/operations/deletereferenceaccountresponse.md), error**
| Error Object                                                      | Status Code                                                       | Content Type                                                      |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| sdkerrors.DeleteReferenceAccountError                             | 401                                                               | application/problem+json                                          |
| sdkerrors.DeleteReferenceAccountReferenceAccountsError            | 403                                                               | application/problem+json                                          |
| sdkerrors.DeleteReferenceAccountReferenceAccountsResponseError    | 404                                                               | application/problem+json                                          |
| sdkerrors.DeleteReferenceAccountReferenceAccountsResponse429Error | 429                                                               | application/problem+json                                          |
| sdkerrors.DeleteReferenceAccountReferenceAccountsResponse500Error | 500                                                               | application/problem+json                                          |
| sdkerrors.DeleteReferenceAccountReferenceAccountsResponse503Error | 503                                                               | application/problem+json                                          |
| sdkerrors.DeleteReferenceAccountReferenceAccountsResponse504Error | 504                                                               | application/problem+json                                          |
| sdkerrors.SDKError                                                | 400-600                                                           | */*                                                               |

## ListReferenceAccounts

Returns the list of reference accounts of a user specified by ID.

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.ReferenceAccounts.ListReferenceAccounts(ctx, operations.ListReferenceAccountsRequest{
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "e50a1010-768d-4d7a-9f1c-4bb86369603d",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UsersListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListReferenceAccountsRequest](../../pkg/models/operations/listreferenceaccountsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.ListReferenceAccountsResponse](../../pkg/models/operations/listreferenceaccountsresponse.md), error**
| Error Object                                                     | Status Code                                                      | Content Type                                                     |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| sdkerrors.ListReferenceAccountsError                             | 400                                                              | application/problem+json                                         |
| sdkerrors.ListReferenceAccountsReferenceAccountsError            | 401                                                              | application/problem+json                                         |
| sdkerrors.ListReferenceAccountsReferenceAccountsResponseError    | 403                                                              | application/problem+json                                         |
| sdkerrors.ListReferenceAccountsReferenceAccountsResponse404Error | 404                                                              | application/problem+json                                         |
| sdkerrors.ListReferenceAccountsReferenceAccountsResponse406Error | 406                                                              | application/problem+json                                         |
| sdkerrors.ListReferenceAccountsReferenceAccountsResponse429Error | 429                                                              | application/problem+json                                         |
| sdkerrors.ListReferenceAccountsReferenceAccountsResponse500Error | 500                                                              | application/problem+json                                         |
| sdkerrors.ListReferenceAccountsReferenceAccountsResponse503Error | 503                                                              | application/problem+json                                         |
| sdkerrors.ListReferenceAccountsReferenceAccountsResponse504Error | 504                                                              | application/problem+json                                         |
| sdkerrors.SDKError                                               | 400-600                                                          | */*                                                              |

## RetrieveReferenceAccount

Retrieves the reference account specified by its ID.

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.ReferenceAccounts.RetrieveReferenceAccount(ctx, operations.RetrieveReferenceAccountRequest{
        ReferenceAccountID: "7e2cbf41-0b63-42ab-90dc-add7f6c2774a",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ReferenceAccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.RetrieveReferenceAccountRequest](../../pkg/models/operations/retrievereferenceaccountrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.RetrieveReferenceAccountResponse](../../pkg/models/operations/retrievereferenceaccountresponse.md), error**
| Error Object                                                        | Status Code                                                         | Content Type                                                        |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| sdkerrors.RetrieveReferenceAccountError                             | 401                                                                 | application/problem+json                                            |
| sdkerrors.RetrieveReferenceAccountReferenceAccountsError            | 403                                                                 | application/problem+json                                            |
| sdkerrors.RetrieveReferenceAccountReferenceAccountsResponseError    | 404                                                                 | application/problem+json                                            |
| sdkerrors.RetrieveReferenceAccountReferenceAccountsResponse406Error | 406                                                                 | application/problem+json                                            |
| sdkerrors.RetrieveReferenceAccountReferenceAccountsResponse429Error | 429                                                                 | application/problem+json                                            |
| sdkerrors.RetrieveReferenceAccountReferenceAccountsResponse500Error | 500                                                                 | application/problem+json                                            |
| sdkerrors.RetrieveReferenceAccountReferenceAccountsResponse503Error | 503                                                                 | application/problem+json                                            |
| sdkerrors.RetrieveReferenceAccountReferenceAccountsResponse504Error | 504                                                                 | application/problem+json                                            |
| sdkerrors.SDKError                                                  | 400-600                                                             | */*                                                                 |
