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
	"context"
	"log"
	upvestdevsamplesdk "github.com/speakeasy-sdks/upvest-dev-sample-sdk"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
)

func main() {
    s := upvestdevsamplesdk.New(
        upvestdevsamplesdk.WithSecurity("YOUR_TOKEN"),
    )

    ctx := context.Background()
    res, err := s.Users.CreateIdentifier(ctx, operations.CreateIdentifierRequest{
        RequestBody: &operations.CreateIdentifierIdentifierCreateRequest{
            Identifier: "string",
            IssuingCountry: "string",
        },
        Signature: "string",
        SignatureInput: "string",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateIdentifierRequest](../../models/operations/createidentifierrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.CreateIdentifierResponse](../../models/operations/createidentifierresponse.md), error**


## CreateUser

Creates a user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	upvestdevsamplesdk "github.com/speakeasy-sdks/upvest-dev-sample-sdk"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/types"
)

func main() {
    s := upvestdevsamplesdk.New(
        upvestdevsamplesdk.WithSecurity("YOUR_TOKEN"),
    )

    ctx := context.Background()
    res, err := s.Users.CreateUser(ctx, operations.CreateUserRequest{
        RequestBody: operations.CreateCreateUserUserCreateRequestCreateUserUserCreateRequestUserBYOLCreateRequest(
                operations.CreateUserUserCreateRequestUserBYOLCreateRequest{
                    Address: operations.CreateUserUserCreateRequestUserBYOLCreateRequestAddress{
                        AddressLine1: "string",
                        City: "Bettiechester",
                        Country: "Saint Vincent and the Grenadines",
                        Postcode: "45093-7750",
                    },
                    BirthDate: types.MustDateFromString("2023-11-27"),
                    FirstName: "Betty",
                    LastName: "Wuckert",
                    Nationalities: []string{
                        "string",
                    },
                    PostalAddress: &operations.CreateUserUserCreateRequestUserBYOLCreateRequestAddress{
                        AddressLine1: "string",
                        City: "Zackeryfurt",
                        Country: "Mozambique",
                        Postcode: "71396-6780",
                    },
                },
        ),
        IdempotencyKey: "ccb07f42-4104-44ad-8e1f-c660bb7b269c",
        Signature: "string",
        SignatureInput: "string",
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.CreateUserRequest](../../models/operations/createuserrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.CreateUserResponse](../../models/operations/createuserresponse.md), error**


## CreateUserCheck

Creates a new check for a user specified by ID.

### Example Usage

```go
package main

import(
	"context"
	"log"
	upvestdevsamplesdk "github.com/speakeasy-sdks/upvest-dev-sample-sdk"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/types"
)

func main() {
    s := upvestdevsamplesdk.New(
        upvestdevsamplesdk.WithSecurity("YOUR_TOKEN"),
    )

    ctx := context.Background()
    res, err := s.Users.CreateUserCheck(ctx, operations.CreateUserCheckRequest{
        RequestBody: operations.CreateCreateUserCheckUserCheckCreateRequestCreateUserCheckUserCheckCreateRequestUserCheckProofOfResidencyCreateRequest(
                operations.CreateUserCheckUserCheckCreateRequestUserCheckProofOfResidencyCreateRequest{
                    CheckConfirmedAt: types.MustTimeFromString("2023-02-24T14:05:33.909Z"),
                    ConfirmedAddress: operations.CreateUserCheckUserCheckCreateRequestUserCheckProofOfResidencyCreateRequestAddress{
                        AddressLine1: "string",
                        City: "New Sheahaven",
                        Country: "Virgin Islands, British",
                        Postcode: "78236-1673",
                    },
                    DataDownloadLink: "http://overcooked-job.org",
                    DocumentType: operations.CreateUserCheckUserCheckCreateRequestUserCheckProofOfResidencyCreateRequestDocumentTypeTelephoneBill,
                    IssuanceDate: types.MustDateFromString("2023-01-16"),
                },
        ),
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "98c63118-bd2e-45df-ae3c-05476d362f77",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateUserCheckRequest](../../models/operations/createusercheckrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.CreateUserCheckResponse](../../models/operations/createusercheckresponse.md), error**


## ListUserAccountGroups

Lists the account groups of a user specified by ID.

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
        upvestdevsamplesdk.WithSecurity("YOUR_TOKEN"),
    )

    ctx := context.Background()
    res, err := s.Users.ListUserAccountGroups(ctx, operations.ListUserAccountGroupsRequest{
        Signature: "string",
        SignatureInput: "string",
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListUserAccountGroupsRequest](../../models/operations/listuseraccountgroupsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.ListUserAccountGroupsResponse](../../models/operations/listuseraccountgroupsresponse.md), error**


## ListUserAccounts

Lists the accounts of a user specified by ID.

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
        upvestdevsamplesdk.WithSecurity("YOUR_TOKEN"),
    )

    ctx := context.Background()
    res, err := s.Users.ListUserAccounts(ctx, operations.ListUserAccountsRequest{
        Signature: "string",
        SignatureInput: "string",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListUserAccountsRequest](../../models/operations/listuseraccountsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.ListUserAccountsResponse](../../models/operations/listuseraccountsresponse.md), error**


## ListUserChecks

Lists all checks for a user specified by ID.

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
        upvestdevsamplesdk.WithSecurity("YOUR_TOKEN"),
    )

    ctx := context.Background()
    res, err := s.Users.ListUserChecks(ctx, operations.ListUserChecksRequest{
        Signature: "string",
        SignatureInput: "string",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListUserChecksRequest](../../models/operations/listuserchecksrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.ListUserChecksResponse](../../models/operations/listuserchecksresponse.md), error**


## ListUserIdentifiers

Lists all existing identifiers of a user used for transaction reporting.

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
        upvestdevsamplesdk.WithSecurity("YOUR_TOKEN"),
    )

    ctx := context.Background()
    res, err := s.Users.ListUserIdentifiers(ctx, operations.ListUserIdentifiersRequest{
        Signature: "string",
        SignatureInput: "string",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListUserIdentifiersRequest](../../models/operations/listuseridentifiersrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.ListUserIdentifiersResponse](../../models/operations/listuseridentifiersresponse.md), error**


## ListUsers

Returns the list of all users.

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
        upvestdevsamplesdk.WithSecurity("YOUR_TOKEN"),
    )

    ctx := context.Background()
    res, err := s.Users.ListUsers(ctx, operations.ListUsersRequest{
        Signature: "string",
        SignatureInput: "string",
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.ListUsersRequest](../../models/operations/listusersrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.ListUsersResponse](../../models/operations/listusersresponse.md), error**


## OffboardUser

Starts the user offboarding process in the background.

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
        upvestdevsamplesdk.WithSecurity("YOUR_TOKEN"),
    )

    ctx := context.Background()
    res, err := s.Users.OffboardUser(ctx, operations.OffboardUserRequest{
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "2ec7a027-b1d2-4089-a65e-dad15450d274",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.OffboardUserRequest](../../models/operations/offboarduserrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.OffboardUserResponse](../../models/operations/offboarduserresponse.md), error**


## RetrieveIdentifier

Returns an existing identifier of a given user used for transaction reporting.

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
        upvestdevsamplesdk.WithSecurity("YOUR_TOKEN"),
    )

    ctx := context.Background()
    res, err := s.Users.RetrieveIdentifier(ctx, operations.RetrieveIdentifierRequest{
        IdentifierID: "7599778d-08e5-4301-b229-50983005b604",
        Signature: "string",
        SignatureInput: "string",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.RetrieveIdentifierRequest](../../models/operations/retrieveidentifierrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.RetrieveIdentifierResponse](../../models/operations/retrieveidentifierresponse.md), error**


## RetrieveUser

Returns the user specified by ID.

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
        upvestdevsamplesdk.WithSecurity("YOUR_TOKEN"),
    )

    ctx := context.Background()
    res, err := s.Users.RetrieveUser(ctx, operations.RetrieveUserRequest{
        Signature: "string",
        SignatureInput: "string",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.RetrieveUserRequest](../../models/operations/retrieveuserrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.RetrieveUserResponse](../../models/operations/retrieveuserresponse.md), error**


## RetrieveUserCheck

Retrieves a check for a user specified by its ID.

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
        upvestdevsamplesdk.WithSecurity("YOUR_TOKEN"),
    )

    ctx := context.Background()
    res, err := s.Users.RetrieveUserCheck(ctx, operations.RetrieveUserCheckRequest{
        CheckID: "7c557256-cf5c-4ad1-a9aa-07e6ee46d87a",
        Signature: "string",
        SignatureInput: "string",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.RetrieveUserCheckRequest](../../models/operations/retrieveusercheckrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.RetrieveUserCheckResponse](../../models/operations/retrieveusercheckresponse.md), error**


## UpdateIdentifier

Updates an existing identifier of a given user used for transaction reporting.

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
        upvestdevsamplesdk.WithSecurity("YOUR_TOKEN"),
    )

    ctx := context.Background()
    res, err := s.Users.UpdateIdentifier(ctx, operations.UpdateIdentifierRequest{
        RequestBody: &operations.UpdateIdentifierIdentifierUpdateRequest{
            Identifier: "string",
        },
        IdentifierID: "f9ef9c4d-678e-40e8-9c2b-f312e052d47d",
        Signature: "string",
        SignatureInput: "string",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateIdentifierRequest](../../models/operations/updateidentifierrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UpdateIdentifierResponse](../../models/operations/updateidentifierresponse.md), error**


## UserDataChange

Requests a data change for a user specified by ID.

### Example Usage

```go
package main

import(
	"context"
	"log"
	upvestdevsamplesdk "github.com/speakeasy-sdks/upvest-dev-sample-sdk"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/types"
)

func main() {
    s := upvestdevsamplesdk.New(
        upvestdevsamplesdk.WithSecurity("YOUR_TOKEN"),
    )

    ctx := context.Background()
    res, err := s.Users.UserDataChange(ctx, operations.UserDataChangeRequest{
        RequestBody: operations.CreateUserDataChangeUserDataChangeRequestUserDataChangeUserDataChangeRequestUserBYOLDataChangeRequest(
                operations.UserDataChangeUserDataChangeRequestUserBYOLDataChangeRequest{
                    Address: &operations.UserDataChangeUserDataChangeRequestUserBYOLDataChangeRequestAddress{
                        AddressLine1: "string",
                        City: "Schinnerfurt",
                        Country: "Syrian Arab Republic",
                        Postcode: "46592-2751",
                    },
                    Nationalities: []string{
                        "string",
                    },
                    PostalAddress: &operations.UserDataChangeUserDataChangeRequestUserBYOLDataChangeRequestAddress{
                        AddressLine1: "string",
                        City: "Fort Johathan",
                        Country: "Wallis and Futuna",
                        Postcode: "04090",
                    },
                },
        ),
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "6d84a24c-2d1c-4b75-ac2d-e28b7e8f8d39",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UserDataChangeRequest](../../models/operations/userdatachangerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.UserDataChangeResponse](../../models/operations/userdatachangeresponse.md), error**
