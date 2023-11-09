# Payments
(*Payments*)

## Overview

All payments related paths

### Available Operations

* [CancelCashWithdrawal](#cancelcashwithdrawal) - Cancel withdrawal by ID
* [CreateCashWithdrawal](#createcashwithdrawal) - Trigger a withdrawal
* [CreateDirectDebit](#createdirectdebit) - Trigger a direct debit
* [ListCashWithdrawals](#listcashwithdrawals) - List withdrawals
* [ListDirectDebits](#listdirectdebits) - List direct debits
* [RetrieveCashWithdrawal](#retrievecashwithdrawal) - Retrieve withdrawal
* [RetrieveDirectDebit](#retrievedirectdebit) - Retrieve a direct debit

## CancelCashWithdrawal

Cancels a withdrawal specified by its ID. It is only possible to cancel a withdrawal if it has the status `NEW`.

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
    res, err := s.Payments.CancelCashWithdrawal(ctx, operations.CancelCashWithdrawalRequest{
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        WithdrawalID: "171565fa-ef4e-474a-b700-f0c84773396a",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CancelCashWithdrawalRequest](../../pkg/models/operations/cancelcashwithdrawalrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.CancelCashWithdrawalResponse](../../pkg/models/operations/cancelcashwithdrawalresponse.md), error**
| Error Object                                           | Status Code                                            | Content Type                                           |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| sdkerrors.CancelCashWithdrawalError                    | 401                                                    | application/problem+json                               |
| sdkerrors.CancelCashWithdrawalPaymentsError            | 403                                                    | application/problem+json                               |
| sdkerrors.CancelCashWithdrawalPaymentsResponseError    | 404                                                    | application/problem+json                               |
| sdkerrors.CancelCashWithdrawalPaymentsResponse422Error | 422                                                    | application/problem+json                               |
| sdkerrors.CancelCashWithdrawalPaymentsResponse429Error | 429                                                    | application/problem+json                               |
| sdkerrors.CancelCashWithdrawalPaymentsResponse500Error | 500                                                    | application/problem+json                               |
| sdkerrors.CancelCashWithdrawalPaymentsResponse503Error | 503                                                    | application/problem+json                               |
| sdkerrors.CancelCashWithdrawalPaymentsResponse504Error | 504                                                    | application/problem+json                               |
| sdkerrors.SDKError                                     | 400-600                                                | */*                                                    |

## CreateCashWithdrawal

Trigger a withdrawal

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
    res, err := s.Payments.CreateCashWithdrawal(ctx, operations.CreateCashWithdrawalRequest{
        RequestBody: &operations.CreateCashWithdrawalPaymentsWithdrawalCreateRequest{
            AccountGroupID: "f243f592-9a68-4da1-9557-1ab0437bf62f",
            Amount: "12.46",
            ReferenceAccountID: "c49a0e62-4ea0-479d-9020-324bce22f352",
            UserID: "c470db78-0f32-4502-9944-477d265e5fa2",
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

    if res.TwoHundredApplicationJSONPaymentsWithdrawal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateCashWithdrawalRequest](../../pkg/models/operations/createcashwithdrawalrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.CreateCashWithdrawalResponse](../../pkg/models/operations/createcashwithdrawalresponse.md), error**
| Error Object                                           | Status Code                                            | Content Type                                           |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| sdkerrors.CreateCashWithdrawalError                    | 400                                                    | application/problem+json                               |
| sdkerrors.CreateCashWithdrawalPaymentsError            | 401                                                    | application/problem+json                               |
| sdkerrors.CreateCashWithdrawalPaymentsResponseError    | 403                                                    | application/problem+json                               |
| sdkerrors.CreateCashWithdrawalPaymentsResponse404Error | 404                                                    | application/problem+json                               |
| sdkerrors.CreateCashWithdrawalPaymentsResponse406Error | 406                                                    | application/problem+json                               |
| sdkerrors.CreateCashWithdrawalPaymentsResponse429Error | 429                                                    | application/problem+json                               |
| sdkerrors.CreateCashWithdrawalPaymentsResponse500Error | 500                                                    | application/problem+json                               |
| sdkerrors.CreateCashWithdrawalPaymentsResponse503Error | 503                                                    | application/problem+json                               |
| sdkerrors.CreateCashWithdrawalPaymentsResponse504Error | 504                                                    | application/problem+json                               |
| sdkerrors.SDKError                                     | 400-600                                                | */*                                                    |

## CreateDirectDebit

Trigger a direct debit

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
    res, err := s.Payments.CreateDirectDebit(ctx, operations.CreateDirectDebitRequest{
        RequestBody: &operations.CreateDirectDebitPaymentsDirectDebitCreateRequest{
            AccountGroupID: "2c3ba409-0daa-4766-9bd5-09bb9ce0d60c",
            CashAmount: "string",
            MandateID: "07d0c6e1-a209-484b-8ddb-8be22446448d",
            UserID: "8d796d7f-3a66-4517-b9b7-f72202891b8f",
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

    if res.TwoHundredApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateDirectDebitRequest](../../pkg/models/operations/createdirectdebitrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.CreateDirectDebitResponse](../../pkg/models/operations/createdirectdebitresponse.md), error**
| Error Object                                        | Status Code                                         | Content Type                                        |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| sdkerrors.CreateDirectDebitError                    | 400                                                 | application/problem+json                            |
| sdkerrors.CreateDirectDebitPaymentsError            | 401                                                 | application/problem+json                            |
| sdkerrors.CreateDirectDebitPaymentsResponseError    | 403                                                 | application/problem+json                            |
| sdkerrors.CreateDirectDebitPaymentsResponse404Error | 404                                                 | application/problem+json                            |
| sdkerrors.CreateDirectDebitPaymentsResponse406Error | 406                                                 | application/problem+json                            |
| sdkerrors.CreateDirectDebitPaymentsResponse429Error | 429                                                 | application/problem+json                            |
| sdkerrors.CreateDirectDebitPaymentsResponse500Error | 500                                                 | application/problem+json                            |
| sdkerrors.CreateDirectDebitPaymentsResponse503Error | 503                                                 | application/problem+json                            |
| sdkerrors.CreateDirectDebitPaymentsResponse504Error | 504                                                 | application/problem+json                            |
| sdkerrors.SDKError                                  | 400-600                                             | */*                                                 |

## ListCashWithdrawals

List withdrawals

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
    res, err := s.Payments.ListCashWithdrawals(ctx, operations.ListCashWithdrawalsRequest{
        AccountGroupID: "d8643156-ad77-42f3-879e-e10e652fca89",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONWithdrawalsListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListCashWithdrawalsRequest](../../pkg/models/operations/listcashwithdrawalsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.ListCashWithdrawalsResponse](../../pkg/models/operations/listcashwithdrawalsresponse.md), error**
| Error Object                                          | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| sdkerrors.ListCashWithdrawalsError                    | 400                                                   | application/problem+json                              |
| sdkerrors.ListCashWithdrawalsPaymentsError            | 401                                                   | application/problem+json                              |
| sdkerrors.ListCashWithdrawalsPaymentsResponseError    | 403                                                   | application/problem+json                              |
| sdkerrors.ListCashWithdrawalsPaymentsResponse404Error | 404                                                   | application/problem+json                              |
| sdkerrors.ListCashWithdrawalsPaymentsResponse406Error | 406                                                   | application/problem+json                              |
| sdkerrors.ListCashWithdrawalsPaymentsResponse429Error | 429                                                   | application/problem+json                              |
| sdkerrors.ListCashWithdrawalsPaymentsResponse500Error | 500                                                   | application/problem+json                              |
| sdkerrors.ListCashWithdrawalsPaymentsResponse503Error | 503                                                   | application/problem+json                              |
| sdkerrors.ListCashWithdrawalsPaymentsResponse504Error | 504                                                   | application/problem+json                              |
| sdkerrors.SDKError                                    | 400-600                                               | */*                                                   |

## ListDirectDebits

List direct debits

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
    res, err := s.Payments.ListDirectDebits(ctx, operations.ListDirectDebitsRequest{
        AccountGroupID: "88b424a5-5c79-48a0-940c-d0211ad65dce",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONPaymentsDirectDebitsListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListDirectDebitsRequest](../../pkg/models/operations/listdirectdebitsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.ListDirectDebitsResponse](../../pkg/models/operations/listdirectdebitsresponse.md), error**
| Error Object                                       | Status Code                                        | Content Type                                       |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| sdkerrors.ListDirectDebitsError                    | 400                                                | application/problem+json                           |
| sdkerrors.ListDirectDebitsPaymentsError            | 401                                                | application/problem+json                           |
| sdkerrors.ListDirectDebitsPaymentsResponseError    | 403                                                | application/problem+json                           |
| sdkerrors.ListDirectDebitsPaymentsResponse404Error | 404                                                | application/problem+json                           |
| sdkerrors.ListDirectDebitsPaymentsResponse406Error | 406                                                | application/problem+json                           |
| sdkerrors.ListDirectDebitsPaymentsResponse429Error | 429                                                | application/problem+json                           |
| sdkerrors.ListDirectDebitsPaymentsResponse500Error | 500                                                | application/problem+json                           |
| sdkerrors.ListDirectDebitsPaymentsResponse503Error | 503                                                | application/problem+json                           |
| sdkerrors.ListDirectDebitsPaymentsResponse504Error | 504                                                | application/problem+json                           |
| sdkerrors.SDKError                                 | 400-600                                            | */*                                                |

## RetrieveCashWithdrawal

Retrieve withdrawal

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
    res, err := s.Payments.RetrieveCashWithdrawal(ctx, operations.RetrieveCashWithdrawalRequest{
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        WithdrawalID: "cd2b7c3f-83b1-455d-8f56-99b4599f90ce",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONPaymentsWithdrawal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.RetrieveCashWithdrawalRequest](../../pkg/models/operations/retrievecashwithdrawalrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.RetrieveCashWithdrawalResponse](../../pkg/models/operations/retrievecashwithdrawalresponse.md), error**
| Error Object                                             | Status Code                                              | Content Type                                             |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| sdkerrors.RetrieveCashWithdrawalError                    | 401                                                      | application/problem+json                                 |
| sdkerrors.RetrieveCashWithdrawalPaymentsError            | 403                                                      | application/problem+json                                 |
| sdkerrors.RetrieveCashWithdrawalPaymentsResponseError    | 404                                                      | application/problem+json                                 |
| sdkerrors.RetrieveCashWithdrawalPaymentsResponse406Error | 406                                                      | application/problem+json                                 |
| sdkerrors.RetrieveCashWithdrawalPaymentsResponse429Error | 429                                                      | application/problem+json                                 |
| sdkerrors.RetrieveCashWithdrawalPaymentsResponse500Error | 500                                                      | application/problem+json                                 |
| sdkerrors.RetrieveCashWithdrawalPaymentsResponse503Error | 503                                                      | application/problem+json                                 |
| sdkerrors.RetrieveCashWithdrawalPaymentsResponse504Error | 504                                                      | application/problem+json                                 |
| sdkerrors.SDKError                                       | 400-600                                                  | */*                                                      |

## RetrieveDirectDebit

Retrieve a direct debit

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
    res, err := s.Payments.RetrieveDirectDebit(ctx, operations.RetrieveDirectDebitRequest{
        DirectDebitID: "b0b3f52f-e687-4685-bdab-d93ae664b9eb",
        Signature: "string",
        SignatureInput: "string",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.RetrieveDirectDebitRequest](../../pkg/models/operations/retrievedirectdebitrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.RetrieveDirectDebitResponse](../../pkg/models/operations/retrievedirectdebitresponse.md), error**
| Error Object                                          | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| sdkerrors.RetrieveDirectDebitError                    | 401                                                   | application/problem+json                              |
| sdkerrors.RetrieveDirectDebitPaymentsError            | 403                                                   | application/problem+json                              |
| sdkerrors.RetrieveDirectDebitPaymentsResponseError    | 404                                                   | application/problem+json                              |
| sdkerrors.RetrieveDirectDebitPaymentsResponse406Error | 406                                                   | application/problem+json                              |
| sdkerrors.RetrieveDirectDebitPaymentsResponse429Error | 429                                                   | application/problem+json                              |
| sdkerrors.RetrieveDirectDebitPaymentsResponse500Error | 500                                                   | application/problem+json                              |
| sdkerrors.RetrieveDirectDebitPaymentsResponse503Error | 503                                                   | application/problem+json                              |
| sdkerrors.RetrieveDirectDebitPaymentsResponse504Error | 504                                                   | application/problem+json                              |
| sdkerrors.SDKError                                    | 400-600                                               | */*                                                   |
