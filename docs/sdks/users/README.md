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
        upvestdevsamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Users.CreateIdentifier(ctx, operations.CreateIdentifierRequest{
        RequestBody: &operations.CreateIdentifierIdentifierCreateRequest{
            Identifier: "improvement",
            IssuingCountry: "Director",
        },
        Signature: "Intranet",
        SignatureInput: "Liaison",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "06eb515e-27f0-43f3-a460-35141216a7dd",
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
        upvestdevsamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Users.CreateUser(ctx, operations.CreateUserRequest{
        RequestBody: operations.CreateCreateUserUserCreateRequestCreateUserUserCreateRequestUserBYOLCreateRequest(
                operations.CreateUserUserCreateRequestUserBYOLCreateRequest{
                    Address: operations.CreateUserUserCreateRequestUserBYOLCreateRequestAddress{
                        AddressLine1: "dicta",
                        City: "Russelcester",
                        Country: "Democratic People's Republic of Korea",
                        Postcode: "09377-5091",
                    },
                    BirthDate: types.MustDateFromString("2023-12-11"),
                    FirstName: "Kaylah",
                    LastName: "Zemlak",
                    Nationalities: []string{
                        "Communications",
                    },
                    PostalAddress: &operations.CreateUserUserCreateRequestUserBYOLCreateRequestAddress{
                        AddressLine1: "Accounts",
                        City: "West Taylorchester",
                        Country: "Peru",
                        Postcode: "80229-1482",
                    },
                },
        ),
        IdempotencyKey: "ccb07f42-4104-44ad-8e1f-c660bb7b269c",
        Signature: "collaboration",
        SignatureInput: "neural",
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
        upvestdevsamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Users.CreateUserCheck(ctx, operations.CreateUserCheckRequest{
        RequestBody: operations.CreateCreateUserCheckUserCheckCreateRequestCreateUserCheckUserCheckCreateRequestUserCheckProofOfResidencyCreateRequest(
                operations.CreateUserCheckUserCheckCreateRequestUserCheckProofOfResidencyCreateRequest{
                    CheckConfirmedAt: types.MustTimeFromString("2023-02-24T14:05:33.909Z"),
                    ConfirmedAddress: operations.CreateUserCheckUserCheckCreateRequestUserCheckProofOfResidencyCreateRequestAddress{
                        AddressLine1: "Newark",
                        City: "Spokane",
                        Country: "Virgin Islands, British",
                        Postcode: "78236-1673",
                    },
                    DataDownloadLink: "http://overcooked-job.org",
                    DocumentType: operations.CreateUserCheckUserCheckCreateRequestUserCheckProofOfResidencyCreateRequestDocumentTypeTelephoneBill,
                    IssuanceDate: types.MustDateFromString("2023-01-16"),
                },
        ),
        Signature: "trans",
        SignatureInput: "Algerian",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "18bd2e5d-fae3-4c05-876d-362f7741e6fa",
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
        upvestdevsamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Users.ListUserAccountGroups(ctx, operations.ListUserAccountGroupsRequest{
        Signature: "capability",
        SignatureInput: "scale",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "566f6bdc-0397-4183-845d-956c1118fc37",
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
        upvestdevsamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Users.ListUserAccounts(ctx, operations.ListUserAccountsRequest{
        Signature: "Ways",
        SignatureInput: "tempora",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "282c4190-a6e0-4d67-96f2-13432d40a499",
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
        upvestdevsamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Users.ListUserChecks(ctx, operations.ListUserChecksRequest{
        Signature: "Legacy",
        SignatureInput: "pascal",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "5d140a30-46db-4475-bdbf-fd92ae246c7b",
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
        upvestdevsamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Users.ListUserIdentifiers(ctx, operations.ListUserIdentifiersRequest{
        Signature: "Latin",
        SignatureInput: "annual",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "a72af34b-c90b-498c-86cb-d026f279cc7c",
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
        upvestdevsamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Users.ListUsers(ctx, operations.ListUsersRequest{
        Signature: "Representative",
        SignatureInput: "Blues",
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
        upvestdevsamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Users.OffboardUser(ctx, operations.OffboardUserRequest{
        Signature: "Lakes",
        SignatureInput: "African",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "b1d20892-65ed-4ad1-9450-d2749d0e752f",
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
        upvestdevsamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Users.RetrieveIdentifier(ctx, operations.RetrieveIdentifierRequest{
        IdentifierID: "7599778d-08e5-4301-b229-50983005b604",
        Signature: "withdrawal",
        SignatureInput: "Guaynabo",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "d1553807-bfe8-4fea-99b8-f7cdd39f42ac",
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
        upvestdevsamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Users.RetrieveUser(ctx, operations.RetrieveUserRequest{
        Signature: "male",
        SignatureInput: "Bronze",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "80860346-9e3d-422a-9d72-cd6eb21664af",
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
        upvestdevsamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Users.RetrieveUserCheck(ctx, operations.RetrieveUserCheckRequest{
        CheckID: "7c557256-cf5c-4ad1-a9aa-07e6ee46d87a",
        Signature: "Sleek",
        SignatureInput: "overriding",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "8df6df78-5464-4a07-9b80-ad51fe682a3f",
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
        upvestdevsamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Users.UpdateIdentifier(ctx, operations.UpdateIdentifierRequest{
        RequestBody: &operations.UpdateIdentifierIdentifierUpdateRequest{
            Identifier: "notwithstanding",
        },
        IdentifierID: "ef9c4d67-8e0e-489c-abf3-12e052d47df1",
        Signature: "regarding",
        SignatureInput: "Garden",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "7ee7b82e-a337-48cf-b11b-13e2bc3a0263",
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
        upvestdevsamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Users.UserDataChange(ctx, operations.UserDataChangeRequest{
        RequestBody: operations.CreateUserDataChangeUserDataChangeRequestUserDataChangeUserDataChangeRequestUserBYOLDataChangeRequest(
                operations.UserDataChangeUserDataChangeRequestUserBYOLDataChangeRequest{
                    Address: &operations.UserDataChangeUserDataChangeRequestUserBYOLDataChangeRequestAddress{
                        AddressLine1: "Data",
                        City: "Streichcester",
                        Country: "Kenya",
                        Postcode: "59227-5128",
                    },
                    Nationalities: []string{
                        "parsing",
                    },
                    PostalAddress: &operations.UserDataChangeUserDataChangeRequestUserBYOLDataChangeRequestAddress{
                        AddressLine1: "approach",
                        City: "Adriannaside",
                        Country: "Belgium",
                        Postcode: "85261",
                    },
                },
        ),
        Signature: "Soft",
        SignatureInput: "Texas",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "1cb756c2-de28-4b7e-8f8d-39737e16e9c4",
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

