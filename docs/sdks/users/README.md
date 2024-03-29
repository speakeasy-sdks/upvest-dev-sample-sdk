# Users
(*Users*)

## Overview

All user related paths.

### Available Operations

* [CreateIdentifier](#createidentifier) - Create a user identifier
* [CreateUser](#createuser) - Create a user
* [CreateUserCheck](#createusercheck) - Create a new check for a user
* [ListUserAccountGroups](#listuseraccountgroups) - Get user account groups
* [ListUserAccounts](#listuseraccounts) - Get user accounts
* [ListUserChecks](#listuserchecks) - Get checks for a user
* [ListUserIdentifiers](#listuseridentifiers) - Get user identifiers
* [ListUsers](#listusers) - Get all users
* [OffboardUser](#offboarduser) - Offboard a user
* [RetrieveIdentifier](#retrieveidentifier) - Get a user identifier by ID
* [RetrieveUser](#retrieveuser) - Get a user by ID
* [RetrieveUserCheck](#retrieveusercheck) - Get a user check by ID
* [UpdateIdentifier](#updateidentifier) - Update a user identifier by ID
* [UserDataChange](#userdatachange) - Change user data

## CreateIdentifier

Creates a new identifier for a user that will be used for transaction reporting obligations. This identifier is required for user activation if the user's nationalities do not allow reporting using the CONCAT format. More information can be found in the [guides](https://docs.upvest.co/guides/users_and_accounts/identifiers_low_prio).

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
    res, err := s.Users.CreateIdentifier(ctx, operations.CreateIdentifierRequest{
        Signature: "<value>",
        SignatureInput: "<value>",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "6cdaa45b-a19b-46b2-86eb-515e27f03f3a",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Identifier != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateIdentifierRequest](../../pkg/models/operations/createidentifierrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.CreateIdentifierResponse](../../pkg/models/operations/createidentifierresponse.md), error**
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.CreateIdentifierError     | 400,401,403,404,406,429,500,503,504 | application/problem+json            |
| sdkerrors.SDKError                  | 4xx-5xx                             | */*                                 |

## CreateUser

Creates a user.

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
    res, err := s.Users.CreateUser(ctx, operations.CreateUserRequest{
        IdempotencyKey: "ccb07f42-4104-44ad-8e1f-c660bb7b269c",
        Signature: "<value>",
        SignatureInput: "<value>",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserCreateRequest != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.CreateUserRequest](../../pkg/models/operations/createuserrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.CreateUserResponse](../../pkg/models/operations/createuserresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.CreateUserError       | 400,401,403,406,429,500,503,504 | application/problem+json        |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |

## CreateUserCheck

Creates a new check for a user specified by ID.

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
    res, err := s.Users.CreateUserCheck(ctx, operations.CreateUserCheckRequest{
        Signature: "<value>",
        SignatureInput: "<value>",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "9b09edfa-be45-4a19-8409-7f3a98c63118",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserCheckCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateUserCheckRequest](../../pkg/models/operations/createusercheckrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.CreateUserCheckResponse](../../pkg/models/operations/createusercheckresponse.md), error**
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.CreateUserCheckError      | 400,401,403,404,406,429,500,503,504 | application/problem+json            |
| sdkerrors.SDKError                  | 4xx-5xx                             | */*                                 |

## ListUserAccountGroups

Lists the account groups of a user specified by ID.

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
    res, err := s.Users.ListUserAccountGroups(ctx, operations.ListUserAccountGroupsRequest{
        Signature: "<value>",
        SignatureInput: "<value>",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "624b566f-6bdc-4039-b183-845d956c1118",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListUserAccountGroupsRequest](../../pkg/models/operations/listuseraccountgroupsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.ListUserAccountGroupsResponse](../../pkg/models/operations/listuseraccountgroupsresponse.md), error**
| Error Object                         | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| sdkerrors.ListUserAccountGroupsError | 400,401,403,404,406,429,500,503,504  | application/problem+json             |
| sdkerrors.SDKError                   | 4xx-5xx                              | */*                                  |

## ListUserAccounts

Lists the accounts of a user specified by ID.

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
    res, err := s.Users.ListUserAccounts(ctx, operations.ListUserAccountsRequest{
        Signature: "<value>",
        SignatureInput: "<value>",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "2f4fc942-82c4-4190-a6e0-d67d6f213432",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListUserAccountsRequest](../../pkg/models/operations/listuseraccountsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.ListUserAccountsResponse](../../pkg/models/operations/listuseraccountsresponse.md), error**
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.ListUserAccountsError     | 400,401,403,404,406,429,500,503,504 | application/problem+json            |
| sdkerrors.SDKError                  | 4xx-5xx                             | */*                                 |

## ListUserChecks

Lists all checks for a user specified by ID.

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
    res, err := s.Users.ListUserChecks(ctx, operations.ListUserChecksRequest{
        Signature: "<value>",
        SignatureInput: "<value>",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "ab810b55-d140-4a30-86db-4757dbffd92a",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserCheckListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListUserChecksRequest](../../pkg/models/operations/listuserchecksrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.ListUserChecksResponse](../../pkg/models/operations/listuserchecksresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ListUserChecksError   | 401,403,404,406,429,500,503,504 | application/problem+json        |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |

## ListUserIdentifiers

Lists all existing identifiers of a user used for transaction reporting.

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
    res, err := s.Users.ListUserIdentifiers(ctx, operations.ListUserIdentifiersRequest{
        Signature: "<value>",
        SignatureInput: "<value>",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "9ad0a72a-f34b-4c90-b98c-c6cbd026f279",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IdentifiersListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListUserIdentifiersRequest](../../pkg/models/operations/listuseridentifiersrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.ListUserIdentifiersResponse](../../pkg/models/operations/listuseridentifiersresponse.md), error**
| Error Object                       | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.ListUserIdentifiersError | 401,403,404,406,429,500,503,504    | application/problem+json           |
| sdkerrors.SDKError                 | 4xx-5xx                            | */*                                |

## ListUsers

Returns the list of all users.

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
    res, err := s.Users.ListUsers(ctx, operations.ListUsersRequest{
        Signature: "<value>",
        SignatureInput: "<value>",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListUsersRequest](../../pkg/models/operations/listusersrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.ListUsersResponse](../../pkg/models/operations/listusersresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ListUsersError        | 400,401,403,406,429,500,503,504 | application/problem+json        |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |

## OffboardUser

Starts the user offboarding process in the background.

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
    res, err := s.Users.OffboardUser(ctx, operations.OffboardUserRequest{
        Signature: "<value>",
        SignatureInput: "<value>",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "2ec7a027-b1d2-4089-a65e-dad15450d274",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.OffboardUserRequest](../../pkg/models/operations/offboarduserrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.OffboardUserResponse](../../pkg/models/operations/offboarduserresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.OffboardUserError     | 401,403,404,406,429,500,503,504 | application/problem+json        |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |

## RetrieveIdentifier

Returns an existing identifier of a given user used for transaction reporting.

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
    res, err := s.Users.RetrieveIdentifier(ctx, operations.RetrieveIdentifierRequest{
        IdentifierID: "7599778d-08e5-4301-b229-50983005b604",
        Signature: "<value>",
        SignatureInput: "<value>",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "7505d155-3807-4bfe-8fea-19b8f7cdd39f",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Identifier != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.RetrieveIdentifierRequest](../../pkg/models/operations/retrieveidentifierrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.RetrieveIdentifierResponse](../../pkg/models/operations/retrieveidentifierresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.RetrieveIdentifierError | 401,403,404,406,429,500,503,504   | application/problem+json          |
| sdkerrors.SDKError                | 4xx-5xx                           | */*                               |

## RetrieveUser

Returns the user specified by ID.

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
    res, err := s.Users.RetrieveUser(ctx, operations.RetrieveUserRequest{
        Signature: "<value>",
        SignatureInput: "<value>",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "bc418086-0346-49e3-922a-9d72cd6eb216",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.RetrieveUserRequest](../../pkg/models/operations/retrieveuserrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.RetrieveUserResponse](../../pkg/models/operations/retrieveuserresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.RetrieveUserError     | 401,403,404,406,429,500,503,504 | application/problem+json        |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |

## RetrieveUserCheck

Retrieves a check for a user specified by its ID.

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
    res, err := s.Users.RetrieveUserCheck(ctx, operations.RetrieveUserCheckRequest{
        CheckID: "7c557256-cf5c-4ad1-a9aa-07e6ee46d87a",
        Signature: "<value>",
        SignatureInput: "<value>",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "38838df6-df78-4546-8a07-db80ad51fe68",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserCheck != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.RetrieveUserCheckRequest](../../pkg/models/operations/retrieveusercheckrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.RetrieveUserCheckResponse](../../pkg/models/operations/retrieveusercheckresponse.md), error**
| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.RetrieveUserCheckError | 401,403,404,406,429,500,503,504  | application/problem+json         |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |

## UpdateIdentifier

Updates an existing identifier of a given user used for transaction reporting.

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
    res, err := s.Users.UpdateIdentifier(ctx, operations.UpdateIdentifierRequest{
        IdentifierID: "f9ef9c4d-678e-40e8-9c2b-f312e052d47d",
        Signature: "<value>",
        SignatureInput: "<value>",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "f1fb357e-e7b8-42ea-b378-cff11b13e2bc",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Identifier != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateIdentifierRequest](../../pkg/models/operations/updateidentifierrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.UpdateIdentifierResponse](../../pkg/models/operations/updateidentifierresponse.md), error**
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.UpdateIdentifierError     | 400,401,403,404,406,429,500,503,504 | application/problem+json            |
| sdkerrors.SDKError                  | 4xx-5xx                             | */*                                 |

## UserDataChange

Requests a data change for a user specified by ID.

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
    res, err := s.Users.UserDataChange(ctx, operations.UserDataChangeRequest{
        Signature: "<value>",
        SignatureInput: "<value>",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "1ac9db79-8f33-4b82-8e8f-6170e16d84a2",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UserDataChangeRequest](../../pkg/models/operations/userdatachangerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UserDataChangeResponse](../../pkg/models/operations/userdatachangeresponse.md), error**
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.UserDataChangeError       | 400,401,403,404,406,429,500,503,504 | application/problem+json            |
| sdkerrors.SDKError                  | 4xx-5xx                             | */*                                 |
