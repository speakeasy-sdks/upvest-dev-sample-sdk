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

    if res.TwoHundredApplicationJSONIdentifier != nil {
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
| Error Object                                    | Status Code                                     | Content Type                                    |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| sdkerrors.CreateIdentifierError                 | 400                                             | application/problem+json                        |
| sdkerrors.CreateIdentifierUsersError            | 401                                             | application/problem+json                        |
| sdkerrors.CreateIdentifierUsersResponseError    | 403                                             | application/problem+json                        |
| sdkerrors.CreateIdentifierUsersResponse404Error | 404                                             | application/problem+json                        |
| sdkerrors.CreateIdentifierUsersResponse406Error | 406                                             | application/problem+json                        |
| sdkerrors.CreateIdentifierUsersResponse429Error | 429                                             | application/problem+json                        |
| sdkerrors.CreateIdentifierUsersResponse500Error | 500                                             | application/problem+json                        |
| sdkerrors.CreateIdentifierUsersResponse503Error | 503                                             | application/problem+json                        |
| sdkerrors.CreateIdentifierUsersResponse504Error | 504                                             | application/problem+json                        |
| sdkerrors.SDKError                              | 400-600                                         | */*                                             |

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
        RequestBody: operations.CreateCreateUserUserCreateRequestUserBYOLCreateRequest(
                operations.UserBYOLCreateRequest{
                    Address: operations.CreateUserAddress{
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
                    PostalAddress: operations.CreatePostalAddressAddress(
                            operations.Address{
                                AddressLine1: "string",
                                City: "Zackeryfurt",
                                Country: "Mozambique",
                                Postcode: "71396-6780",
                            },
                    ),
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

    if res.TwoHundredApplicationJSONUserCreateRequest != nil {
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
| Error Object                              | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| sdkerrors.CreateUserError                 | 400                                       | application/problem+json                  |
| sdkerrors.CreateUserUsersError            | 401                                       | application/problem+json                  |
| sdkerrors.CreateUserUsersResponseError    | 403                                       | application/problem+json                  |
| sdkerrors.CreateUserUsersResponse406Error | 406                                       | application/problem+json                  |
| sdkerrors.CreateUserUsersResponse429Error | 429                                       | application/problem+json                  |
| sdkerrors.CreateUserUsersResponse500Error | 500                                       | application/problem+json                  |
| sdkerrors.CreateUserUsersResponse503Error | 503                                       | application/problem+json                  |
| sdkerrors.CreateUserUsersResponse504Error | 504                                       | application/problem+json                  |
| sdkerrors.SDKError                        | 400-600                                   | */*                                       |

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
        RequestBody: operations.CreateCreateUserCheckUserCheckCreateRequestUserCheckProofOfResidencyCreateRequest(
                operations.UserCheckProofOfResidencyCreateRequest{
                    CheckConfirmedAt: types.MustTimeFromString("2023-02-24T14:05:33.909Z"),
                    ConfirmedAddress: operations.CreateUserCheckUsersAddress{
                        AddressLine1: "string",
                        City: "New Sheahaven",
                        Country: "Virgin Islands, British",
                        Postcode: "78236-1673",
                    },
                    DataDownloadLink: "http://overcooked-job.org",
                    DocumentType: operations.CreateUserCheckDocumentTypeTelephoneBill,
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

    if res.TwoHundredAndTwoApplicationJSONUserCheckCreateResponse != nil {
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
| Error Object                                   | Status Code                                    | Content Type                                   |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| sdkerrors.CreateUserCheckError                 | 400                                            | application/problem+json                       |
| sdkerrors.CreateUserCheckUsersError            | 401                                            | application/problem+json                       |
| sdkerrors.CreateUserCheckUsersResponseError    | 403                                            | application/problem+json                       |
| sdkerrors.CreateUserCheckUsersResponse404Error | 404                                            | application/problem+json                       |
| sdkerrors.CreateUserCheckUsersResponse406Error | 406                                            | application/problem+json                       |
| sdkerrors.CreateUserCheckUsersResponse429Error | 429                                            | application/problem+json                       |
| sdkerrors.CreateUserCheckUsersResponse500Error | 500                                            | application/problem+json                       |
| sdkerrors.CreateUserCheckUsersResponse503Error | 503                                            | application/problem+json                       |
| sdkerrors.CreateUserCheckUsersResponse504Error | 504                                            | application/problem+json                       |
| sdkerrors.SDKError                             | 400-600                                        | */*                                            |

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
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "624b566f-6bdc-4039-b183-845d956c1118",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListUserAccountGroupsRequest](../../pkg/models/operations/listuseraccountgroupsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.ListUserAccountGroupsResponse](../../pkg/models/operations/listuseraccountgroupsresponse.md), error**
| Error Object                                         | Status Code                                          | Content Type                                         |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| sdkerrors.ListUserAccountGroupsError                 | 400                                                  | application/problem+json                             |
| sdkerrors.ListUserAccountGroupsUsersError            | 401                                                  | application/problem+json                             |
| sdkerrors.ListUserAccountGroupsUsersResponseError    | 403                                                  | application/problem+json                             |
| sdkerrors.ListUserAccountGroupsUsersResponse404Error | 404                                                  | application/problem+json                             |
| sdkerrors.ListUserAccountGroupsUsersResponse406Error | 406                                                  | application/problem+json                             |
| sdkerrors.ListUserAccountGroupsUsersResponse429Error | 429                                                  | application/problem+json                             |
| sdkerrors.ListUserAccountGroupsUsersResponse500Error | 500                                                  | application/problem+json                             |
| sdkerrors.ListUserAccountGroupsUsersResponse503Error | 503                                                  | application/problem+json                             |
| sdkerrors.ListUserAccountGroupsUsersResponse504Error | 504                                                  | application/problem+json                             |
| sdkerrors.SDKError                                   | 400-600                                              | */*                                                  |

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
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "2f4fc942-82c4-4190-a6e0-d67d6f213432",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListUserAccountsRequest](../../pkg/models/operations/listuseraccountsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.ListUserAccountsResponse](../../pkg/models/operations/listuseraccountsresponse.md), error**
| Error Object                                    | Status Code                                     | Content Type                                    |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| sdkerrors.ListUserAccountsError                 | 400                                             | application/problem+json                        |
| sdkerrors.ListUserAccountsUsersError            | 401                                             | application/problem+json                        |
| sdkerrors.ListUserAccountsUsersResponseError    | 403                                             | application/problem+json                        |
| sdkerrors.ListUserAccountsUsersResponse404Error | 404                                             | application/problem+json                        |
| sdkerrors.ListUserAccountsUsersResponse406Error | 406                                             | application/problem+json                        |
| sdkerrors.ListUserAccountsUsersResponse429Error | 429                                             | application/problem+json                        |
| sdkerrors.ListUserAccountsUsersResponse500Error | 500                                             | application/problem+json                        |
| sdkerrors.ListUserAccountsUsersResponse503Error | 503                                             | application/problem+json                        |
| sdkerrors.ListUserAccountsUsersResponse504Error | 504                                             | application/problem+json                        |
| sdkerrors.SDKError                              | 400-600                                         | */*                                             |

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
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "ab810b55-d140-4a30-86db-4757dbffd92a",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONUserCheckListResponse != nil {
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
| Error Object                                  | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| sdkerrors.ListUserChecksError                 | 401                                           | application/problem+json                      |
| sdkerrors.ListUserChecksUsersError            | 403                                           | application/problem+json                      |
| sdkerrors.ListUserChecksUsersResponseError    | 404                                           | application/problem+json                      |
| sdkerrors.ListUserChecksUsersResponse406Error | 406                                           | application/problem+json                      |
| sdkerrors.ListUserChecksUsersResponse429Error | 429                                           | application/problem+json                      |
| sdkerrors.ListUserChecksUsersResponse500Error | 500                                           | application/problem+json                      |
| sdkerrors.ListUserChecksUsersResponse503Error | 503                                           | application/problem+json                      |
| sdkerrors.ListUserChecksUsersResponse504Error | 504                                           | application/problem+json                      |
| sdkerrors.SDKError                            | 400-600                                       | */*                                           |

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
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "9ad0a72a-f34b-4c90-b98c-c6cbd026f279",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONIdentifiersListResponse != nil {
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
| Error Object                                       | Status Code                                        | Content Type                                       |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| sdkerrors.ListUserIdentifiersError                 | 401                                                | application/problem+json                           |
| sdkerrors.ListUserIdentifiersUsersError            | 403                                                | application/problem+json                           |
| sdkerrors.ListUserIdentifiersUsersResponseError    | 404                                                | application/problem+json                           |
| sdkerrors.ListUserIdentifiersUsersResponse406Error | 406                                                | application/problem+json                           |
| sdkerrors.ListUserIdentifiersUsersResponse429Error | 429                                                | application/problem+json                           |
| sdkerrors.ListUserIdentifiersUsersResponse500Error | 500                                                | application/problem+json                           |
| sdkerrors.ListUserIdentifiersUsersResponse503Error | 503                                                | application/problem+json                           |
| sdkerrors.ListUserIdentifiersUsersResponse504Error | 504                                                | application/problem+json                           |
| sdkerrors.SDKError                                 | 400-600                                            | */*                                                |

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
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONUsersListResponse != nil {
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
| Error Object                             | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| sdkerrors.ListUsersError                 | 400                                      | application/problem+json                 |
| sdkerrors.ListUsersUsersError            | 401                                      | application/problem+json                 |
| sdkerrors.ListUsersUsersResponseError    | 403                                      | application/problem+json                 |
| sdkerrors.ListUsersUsersResponse406Error | 406                                      | application/problem+json                 |
| sdkerrors.ListUsersUsersResponse429Error | 429                                      | application/problem+json                 |
| sdkerrors.ListUsersUsersResponse500Error | 500                                      | application/problem+json                 |
| sdkerrors.ListUsersUsersResponse503Error | 503                                      | application/problem+json                 |
| sdkerrors.ListUsersUsersResponse504Error | 504                                      | application/problem+json                 |
| sdkerrors.SDKError                       | 400-600                                  | */*                                      |

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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.OffboardUserRequest](../../pkg/models/operations/offboarduserrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.OffboardUserResponse](../../pkg/models/operations/offboarduserresponse.md), error**
| Error Object                                | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| sdkerrors.OffboardUserError                 | 401                                         | application/problem+json                    |
| sdkerrors.OffboardUserUsersError            | 403                                         | application/problem+json                    |
| sdkerrors.OffboardUserUsersResponseError    | 404                                         | application/problem+json                    |
| sdkerrors.OffboardUserUsersResponse406Error | 406                                         | application/problem+json                    |
| sdkerrors.OffboardUserUsersResponse429Error | 429                                         | application/problem+json                    |
| sdkerrors.OffboardUserUsersResponse500Error | 500                                         | application/problem+json                    |
| sdkerrors.OffboardUserUsersResponse503Error | 503                                         | application/problem+json                    |
| sdkerrors.OffboardUserUsersResponse504Error | 504                                         | application/problem+json                    |
| sdkerrors.SDKError                          | 400-600                                     | */*                                         |

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
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "7505d155-3807-4bfe-8fea-19b8f7cdd39f",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONIdentifier != nil {
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
| Error Object                                      | Status Code                                       | Content Type                                      |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| sdkerrors.RetrieveIdentifierError                 | 401                                               | application/problem+json                          |
| sdkerrors.RetrieveIdentifierUsersError            | 403                                               | application/problem+json                          |
| sdkerrors.RetrieveIdentifierUsersResponseError    | 404                                               | application/problem+json                          |
| sdkerrors.RetrieveIdentifierUsersResponse406Error | 406                                               | application/problem+json                          |
| sdkerrors.RetrieveIdentifierUsersResponse429Error | 429                                               | application/problem+json                          |
| sdkerrors.RetrieveIdentifierUsersResponse500Error | 500                                               | application/problem+json                          |
| sdkerrors.RetrieveIdentifierUsersResponse503Error | 503                                               | application/problem+json                          |
| sdkerrors.RetrieveIdentifierUsersResponse504Error | 504                                               | application/problem+json                          |
| sdkerrors.SDKError                                | 400-600                                           | */*                                               |

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
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "bc418086-0346-49e3-922a-9d72cd6eb216",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONUserGetResponse != nil {
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
| Error Object                                | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| sdkerrors.RetrieveUserError                 | 401                                         | application/problem+json                    |
| sdkerrors.RetrieveUserUsersError            | 403                                         | application/problem+json                    |
| sdkerrors.RetrieveUserUsersResponseError    | 404                                         | application/problem+json                    |
| sdkerrors.RetrieveUserUsersResponse406Error | 406                                         | application/problem+json                    |
| sdkerrors.RetrieveUserUsersResponse429Error | 429                                         | application/problem+json                    |
| sdkerrors.RetrieveUserUsersResponse500Error | 500                                         | application/problem+json                    |
| sdkerrors.RetrieveUserUsersResponse503Error | 503                                         | application/problem+json                    |
| sdkerrors.RetrieveUserUsersResponse504Error | 504                                         | application/problem+json                    |
| sdkerrors.SDKError                          | 400-600                                     | */*                                         |

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
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        UserID: "38838df6-df78-4546-8a07-db80ad51fe68",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONUserCheck != nil {
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
| Error Object                                     | Status Code                                      | Content Type                                     |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| sdkerrors.RetrieveUserCheckError                 | 401                                              | application/problem+json                         |
| sdkerrors.RetrieveUserCheckUsersError            | 403                                              | application/problem+json                         |
| sdkerrors.RetrieveUserCheckUsersResponseError    | 404                                              | application/problem+json                         |
| sdkerrors.RetrieveUserCheckUsersResponse406Error | 406                                              | application/problem+json                         |
| sdkerrors.RetrieveUserCheckUsersResponse429Error | 429                                              | application/problem+json                         |
| sdkerrors.RetrieveUserCheckUsersResponse500Error | 500                                              | application/problem+json                         |
| sdkerrors.RetrieveUserCheckUsersResponse503Error | 503                                              | application/problem+json                         |
| sdkerrors.RetrieveUserCheckUsersResponse504Error | 504                                              | application/problem+json                         |
| sdkerrors.SDKError                               | 400-600                                          | */*                                              |

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

    if res.TwoHundredApplicationJSONIdentifier != nil {
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
| Error Object                                    | Status Code                                     | Content Type                                    |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| sdkerrors.UpdateIdentifierError                 | 400                                             | application/problem+json                        |
| sdkerrors.UpdateIdentifierUsersError            | 401                                             | application/problem+json                        |
| sdkerrors.UpdateIdentifierUsersResponseError    | 403                                             | application/problem+json                        |
| sdkerrors.UpdateIdentifierUsersResponse404Error | 404                                             | application/problem+json                        |
| sdkerrors.UpdateIdentifierUsersResponse406Error | 406                                             | application/problem+json                        |
| sdkerrors.UpdateIdentifierUsersResponse429Error | 429                                             | application/problem+json                        |
| sdkerrors.UpdateIdentifierUsersResponse500Error | 500                                             | application/problem+json                        |
| sdkerrors.UpdateIdentifierUsersResponse503Error | 503                                             | application/problem+json                        |
| sdkerrors.UpdateIdentifierUsersResponse504Error | 504                                             | application/problem+json                        |
| sdkerrors.SDKError                              | 400-600                                         | */*                                             |

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
        RequestBody: operations.CreateUserDataChangeUserDataChangeRequestUserBYOLDataChangeRequest(
                operations.UserBYOLDataChangeRequest{
                    Address: &operations.UserDataChangeUsersAddress{
                        AddressLine1: "string",
                        City: "Schinnerfurt",
                        Country: "Syrian Arab Republic",
                        Postcode: "46592-2751",
                    },
                    Nationalities: []string{
                        "string",
                    },
                    PostalAddress: operations.CreateUserDataChangePostalAddressUserDataChangeAddress(
                            operations.UserDataChangeAddress{
                                AddressLine1: "string",
                                City: "Fort Johathan",
                                Country: "Wallis and Futuna",
                                Postcode: "04090",
                            },
                    ),
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UserDataChangeRequest](../../pkg/models/operations/userdatachangerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UserDataChangeResponse](../../pkg/models/operations/userdatachangeresponse.md), error**
| Error Object                                  | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| sdkerrors.UserDataChangeError                 | 400                                           | application/problem+json                      |
| sdkerrors.UserDataChangeUsersError            | 401                                           | application/problem+json                      |
| sdkerrors.UserDataChangeUsersResponseError    | 403                                           | application/problem+json                      |
| sdkerrors.UserDataChangeUsersResponse404Error | 404                                           | application/problem+json                      |
| sdkerrors.UserDataChangeUsersResponse406Error | 406                                           | application/problem+json                      |
| sdkerrors.UserDataChangeUsersResponse429Error | 429                                           | application/problem+json                      |
| sdkerrors.UserDataChangeUsersResponse500Error | 500                                           | application/problem+json                      |
| sdkerrors.UserDataChangeUsersResponse503Error | 503                                           | application/problem+json                      |
| sdkerrors.UserDataChangeUsersResponse504Error | 504                                           | application/problem+json                      |
| sdkerrors.SDKError                            | 400-600                                       | */*                                           |
