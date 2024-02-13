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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
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
| Error Object                          | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.AccountClosureError         | 401,403,404,406,429,500,503,504       | application/problem+json              |
| sdkerrors.AccountClosureAccountsError | 409                                   | application/problem+json              |
| sdkerrors.SDKError                    | 4xx-5xx                               | */*                                   |

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
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
| Error Object                               | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.AccountGroupClosureError         | 401,403,404,406,429,500,503,504            | application/problem+json                   |
| sdkerrors.AccountGroupClosureAccountsError | 409                                        | application/problem+json                   |
| sdkerrors.SDKError                         | 4xx-5xx                                    | */*                                        |

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Accounts.CreateAccount(ctx, operations.CreateAccountRequest{
        IdempotencyKey: "ccb07f42-4104-44ad-8e1f-c660bb7b269c",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Account != nil {
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
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.CreateAccountError        | 400,401,403,404,406,429,500,503,504 | application/problem+json            |
| sdkerrors.SDKError                  | 4xx-5xx                             | */*                                 |

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Accounts.CreateAccountGroup(ctx, operations.CreateAccountGroupRequest{
        IdempotencyKey: "ccb07f42-4104-44ad-8e1f-c660bb7b269c",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountGroup != nil {
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
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.CreateAccountGroupError   | 400,401,403,404,406,429,500,503,504 | application/problem+json            |
| sdkerrors.SDKError                  | 4xx-5xx                             | */*                                 |

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
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

    if res.AccountGroupsListResponse != nil {
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
| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.ListAccountGroupsError | 400,401,403,406,429,500,503,504  | application/problem+json         |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
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

    if res.AccountsListResponse != nil {
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
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ListAccountsError     | 400,401,403,406,429,500,503,504 | application/problem+json        |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
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

    if res.Account != nil {
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
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.RetrieveAccountError  | 401,403,404,406,429,500,503,504 | application/problem+json        |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
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

    if res.AccountGroup != nil {
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
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.RetrieveAccountGroupError | 401,403,404,406,429,500,503,504     | application/problem+json            |
| sdkerrors.SDKError                  | 4xx-5xx                             | */*                                 |

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
        upvestdevsamplesdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Accounts.UpdateAccount(ctx, operations.UpdateAccountRequest{
        AccountID: "9adf3964-ab4d-47a0-bf84-a54e6f410e38",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Account != nil {
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
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.UpdateAccountError        | 400,401,403,404,406,429,500,503,504 | application/problem+json            |
| sdkerrors.SDKError                  | 4xx-5xx                             | */*                                 |
