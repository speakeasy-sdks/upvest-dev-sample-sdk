# Accounts
(*Accounts*)

## Overview

All accounts related paths

### Available Operations

* [AccountClosure](#accountclosure) - Close a user account by ID
* [AccountGroupClosure](#accountgroupclosure) - Close an account group by ID
* [CreateAccount](#createaccount) - Create an account
* [CreateAccountGroup](#createaccountgroup) - Create an account group
* [ListAccountGroups](#listaccountgroups) - Get account groups
* [ListAccounts](#listaccounts) - Get accounts
* [RetrieveAccount](#retrieveaccount) - Get a user account by ID
* [RetrieveAccountGroup](#retrieveaccountgroup) - Get an account group by ID
* [UpdateAccount](#updateaccount) - Update user account

## AccountClosure

Initiates the closure request for an account specified by its ID.

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
    res, err := s.Accounts.AccountClosure(ctx, operations.AccountClosureRequest{
        AccountID: "87f46f4c-298e-4960-b531-5043c3be9e8d",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.AccountClosureRequest](../../pkg/models/operations/accountclosurerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.AccountClosureResponse](../../pkg/models/operations/accountclosureresponse.md), error**
| Error Object                                     | Status Code                                      | Content Type                                     |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| sdkerrors.AccountClosureError                    | 401                                              | application/problem+json                         |
| sdkerrors.AccountClosureAccountsError            | 403                                              | application/problem+json                         |
| sdkerrors.AccountClosureAccountsResponseError    | 404                                              | application/problem+json                         |
| sdkerrors.AccountClosureAccountsResponse406Error | 406                                              | application/problem+json                         |
| sdkerrors.AccountClosureAccountsResponse409Error | 409                                              | application/problem+json                         |
| sdkerrors.AccountClosureAccountsResponse429Error | 429                                              | application/problem+json                         |
| sdkerrors.AccountClosureAccountsResponse500Error | 500                                              | application/problem+json                         |
| sdkerrors.AccountClosureAccountsResponse503Error | 503                                              | application/problem+json                         |
| sdkerrors.AccountClosureAccountsResponse504Error | 504                                              | application/problem+json                         |
| sdkerrors.SDKError                               | 400-600                                          | */*                                              |

## AccountGroupClosure

Initiates the closure request for an account group specified by its ID.

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
    res, err := s.Accounts.AccountGroupClosure(ctx, operations.AccountGroupClosureRequest{
        AccountGroupID: "b19f4004-9a18-414a-bb76-73957c03485e",
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.AccountGroupClosureRequest](../../pkg/models/operations/accountgroupclosurerequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.AccountGroupClosureResponse](../../pkg/models/operations/accountgroupclosureresponse.md), error**
| Error Object                                          | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| sdkerrors.AccountGroupClosureError                    | 401                                                   | application/problem+json                              |
| sdkerrors.AccountGroupClosureAccountsError            | 403                                                   | application/problem+json                              |
| sdkerrors.AccountGroupClosureAccountsResponseError    | 404                                                   | application/problem+json                              |
| sdkerrors.AccountGroupClosureAccountsResponse406Error | 406                                                   | application/problem+json                              |
| sdkerrors.AccountGroupClosureAccountsResponse409Error | 409                                                   | application/problem+json                              |
| sdkerrors.AccountGroupClosureAccountsResponse429Error | 429                                                   | application/problem+json                              |
| sdkerrors.AccountGroupClosureAccountsResponse500Error | 500                                                   | application/problem+json                              |
| sdkerrors.AccountGroupClosureAccountsResponse503Error | 503                                                   | application/problem+json                              |
| sdkerrors.AccountGroupClosureAccountsResponse504Error | 504                                                   | application/problem+json                              |
| sdkerrors.SDKError                                    | 400-600                                               | */*                                                   |

## CreateAccount

Creates an account.

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
    res, err := s.Accounts.CreateAccount(ctx, operations.CreateAccountRequest{
        RequestBody: &operations.CreateAccountRequestBody{
            AccountGroupID: "e9562292-f304-4c6a-8db0-ea541f32fba9",
            Type: operations.TypeTrading,
            UserID: "d04cd2d5-ae02-4bb1-9118-75a95a0f2373",
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

    if res.TwoHundredApplicationJSONAccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateAccountRequest](../../pkg/models/operations/createaccountrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.CreateAccountResponse](../../pkg/models/operations/createaccountresponse.md), error**
| Error Object                                    | Status Code                                     | Content Type                                    |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| sdkerrors.CreateAccountError                    | 400                                             | application/problem+json                        |
| sdkerrors.CreateAccountAccountsError            | 401                                             | application/problem+json                        |
| sdkerrors.CreateAccountAccountsResponseError    | 403                                             | application/problem+json                        |
| sdkerrors.CreateAccountAccountsResponse404Error | 404                                             | application/problem+json                        |
| sdkerrors.CreateAccountAccountsResponse406Error | 406                                             | application/problem+json                        |
| sdkerrors.CreateAccountAccountsResponse429Error | 429                                             | application/problem+json                        |
| sdkerrors.CreateAccountAccountsResponse500Error | 500                                             | application/problem+json                        |
| sdkerrors.CreateAccountAccountsResponse503Error | 503                                             | application/problem+json                        |
| sdkerrors.CreateAccountAccountsResponse504Error | 504                                             | application/problem+json                        |
| sdkerrors.SDKError                              | 400-600                                         | */*                                             |

## CreateAccountGroup

Creates an account group.

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
    res, err := s.Accounts.CreateAccountGroup(ctx, operations.CreateAccountGroupRequest{
        RequestBody: &operations.CreateAccountGroupAccountGroupCreateRequest{
            Type: operations.CreateAccountGroupTypeLegalEntity,
            UserID: "9172e12f-f215-477c-9ccc-f257f38b8e8a",
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

    if res.TwoHundredApplicationJSONAccountGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateAccountGroupRequest](../../pkg/models/operations/createaccountgrouprequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.CreateAccountGroupResponse](../../pkg/models/operations/createaccountgroupresponse.md), error**
| Error Object                                         | Status Code                                          | Content Type                                         |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| sdkerrors.CreateAccountGroupError                    | 400                                                  | application/problem+json                             |
| sdkerrors.CreateAccountGroupAccountsError            | 401                                                  | application/problem+json                             |
| sdkerrors.CreateAccountGroupAccountsResponseError    | 403                                                  | application/problem+json                             |
| sdkerrors.CreateAccountGroupAccountsResponse404Error | 404                                                  | application/problem+json                             |
| sdkerrors.CreateAccountGroupAccountsResponse406Error | 406                                                  | application/problem+json                             |
| sdkerrors.CreateAccountGroupAccountsResponse429Error | 429                                                  | application/problem+json                             |
| sdkerrors.CreateAccountGroupAccountsResponse500Error | 500                                                  | application/problem+json                             |
| sdkerrors.CreateAccountGroupAccountsResponse503Error | 503                                                  | application/problem+json                             |
| sdkerrors.CreateAccountGroupAccountsResponse504Error | 504                                                  | application/problem+json                             |
| sdkerrors.SDKError                                   | 400-600                                              | */*                                                  |

## ListAccountGroups

Returns a list of all account groups.

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
    res, err := s.Accounts.ListAccountGroups(ctx, operations.ListAccountGroupsRequest{
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONAccountGroupsListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListAccountGroupsRequest](../../pkg/models/operations/listaccountgroupsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.ListAccountGroupsResponse](../../pkg/models/operations/listaccountgroupsresponse.md), error**
| Error Object                                        | Status Code                                         | Content Type                                        |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| sdkerrors.ListAccountGroupsError                    | 400                                                 | application/problem+json                            |
| sdkerrors.ListAccountGroupsAccountsError            | 401                                                 | application/problem+json                            |
| sdkerrors.ListAccountGroupsAccountsResponseError    | 403                                                 | application/problem+json                            |
| sdkerrors.ListAccountGroupsAccountsResponse406Error | 406                                                 | application/problem+json                            |
| sdkerrors.ListAccountGroupsAccountsResponse429Error | 429                                                 | application/problem+json                            |
| sdkerrors.ListAccountGroupsAccountsResponse500Error | 500                                                 | application/problem+json                            |
| sdkerrors.ListAccountGroupsAccountsResponse503Error | 503                                                 | application/problem+json                            |
| sdkerrors.ListAccountGroupsAccountsResponse504Error | 504                                                 | application/problem+json                            |
| sdkerrors.SDKError                                  | 400-600                                             | */*                                                 |

## ListAccounts

Returns a list of all accounts.

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
    res, err := s.Accounts.ListAccounts(ctx, operations.ListAccountsRequest{
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONAccountsListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListAccountsRequest](../../pkg/models/operations/listaccountsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.ListAccountsResponse](../../pkg/models/operations/listaccountsresponse.md), error**
| Error Object                                   | Status Code                                    | Content Type                                   |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| sdkerrors.ListAccountsError                    | 400                                            | application/problem+json                       |
| sdkerrors.ListAccountsAccountsError            | 401                                            | application/problem+json                       |
| sdkerrors.ListAccountsAccountsResponseError    | 403                                            | application/problem+json                       |
| sdkerrors.ListAccountsAccountsResponse406Error | 406                                            | application/problem+json                       |
| sdkerrors.ListAccountsAccountsResponse429Error | 429                                            | application/problem+json                       |
| sdkerrors.ListAccountsAccountsResponse500Error | 500                                            | application/problem+json                       |
| sdkerrors.ListAccountsAccountsResponse503Error | 503                                            | application/problem+json                       |
| sdkerrors.ListAccountsAccountsResponse504Error | 504                                            | application/problem+json                       |
| sdkerrors.SDKError                             | 400-600                                        | */*                                            |

## RetrieveAccount

Returns the user account specified by its ID.

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
    res, err := s.Accounts.RetrieveAccount(ctx, operations.RetrieveAccountRequest{
        AccountID: "4e13c906-e457-4f4f-a3a8-bf3db1cfc77e",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONAccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.RetrieveAccountRequest](../../pkg/models/operations/retrieveaccountrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.RetrieveAccountResponse](../../pkg/models/operations/retrieveaccountresponse.md), error**
| Error Object                                      | Status Code                                       | Content Type                                      |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| sdkerrors.RetrieveAccountError                    | 401                                               | application/problem+json                          |
| sdkerrors.RetrieveAccountAccountsError            | 403                                               | application/problem+json                          |
| sdkerrors.RetrieveAccountAccountsResponseError    | 404                                               | application/problem+json                          |
| sdkerrors.RetrieveAccountAccountsResponse406Error | 406                                               | application/problem+json                          |
| sdkerrors.RetrieveAccountAccountsResponse429Error | 429                                               | application/problem+json                          |
| sdkerrors.RetrieveAccountAccountsResponse500Error | 500                                               | application/problem+json                          |
| sdkerrors.RetrieveAccountAccountsResponse503Error | 503                                               | application/problem+json                          |
| sdkerrors.RetrieveAccountAccountsResponse504Error | 504                                               | application/problem+json                          |
| sdkerrors.SDKError                                | 400-600                                           | */*                                               |

## RetrieveAccountGroup

Returns the account group specified bu its ID.

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
    res, err := s.Accounts.RetrieveAccountGroup(ctx, operations.RetrieveAccountGroupRequest{
        AccountGroupID: "e5a4d5a1-aa6d-4a78-8b31-db3645b07f04",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONAccountGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.RetrieveAccountGroupRequest](../../pkg/models/operations/retrieveaccountgrouprequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.RetrieveAccountGroupResponse](../../pkg/models/operations/retrieveaccountgroupresponse.md), error**
| Error Object                                           | Status Code                                            | Content Type                                           |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| sdkerrors.RetrieveAccountGroupError                    | 401                                                    | application/problem+json                               |
| sdkerrors.RetrieveAccountGroupAccountsError            | 403                                                    | application/problem+json                               |
| sdkerrors.RetrieveAccountGroupAccountsResponseError    | 404                                                    | application/problem+json                               |
| sdkerrors.RetrieveAccountGroupAccountsResponse406Error | 406                                                    | application/problem+json                               |
| sdkerrors.RetrieveAccountGroupAccountsResponse429Error | 429                                                    | application/problem+json                               |
| sdkerrors.RetrieveAccountGroupAccountsResponse500Error | 500                                                    | application/problem+json                               |
| sdkerrors.RetrieveAccountGroupAccountsResponse503Error | 503                                                    | application/problem+json                               |
| sdkerrors.RetrieveAccountGroupAccountsResponse504Error | 504                                                    | application/problem+json                               |
| sdkerrors.SDKError                                     | 400-600                                                | */*                                                    |

## UpdateAccount

Updates the user account specified by its ID.

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
    res, err := s.Accounts.UpdateAccount(ctx, operations.UpdateAccountRequest{
        RequestBody: &operations.UpdateAccountAccountUpdateRequest{},
        AccountID: "9adf3964-ab4d-47a0-bf84-a54e6f410e38",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONAccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateAccountRequest](../../pkg/models/operations/updateaccountrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.UpdateAccountResponse](../../pkg/models/operations/updateaccountresponse.md), error**
| Error Object                                    | Status Code                                     | Content Type                                    |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| sdkerrors.UpdateAccountError                    | 400                                             | application/problem+json                        |
| sdkerrors.UpdateAccountAccountsError            | 401                                             | application/problem+json                        |
| sdkerrors.UpdateAccountAccountsResponseError    | 403                                             | application/problem+json                        |
| sdkerrors.UpdateAccountAccountsResponse404Error | 404                                             | application/problem+json                        |
| sdkerrors.UpdateAccountAccountsResponse406Error | 406                                             | application/problem+json                        |
| sdkerrors.UpdateAccountAccountsResponse429Error | 429                                             | application/problem+json                        |
| sdkerrors.UpdateAccountAccountsResponse500Error | 500                                             | application/problem+json                        |
| sdkerrors.UpdateAccountAccountsResponse503Error | 503                                             | application/problem+json                        |
| sdkerrors.UpdateAccountAccountsResponse504Error | 504                                             | application/problem+json                        |
| sdkerrors.SDKError                              | 400-600                                         | */*                                             |
