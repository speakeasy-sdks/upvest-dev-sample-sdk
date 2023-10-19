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
        upvestdevsamplesdk.WithSecurity("YOUR_TOKEN"),
    )

    ctx := context.Background()
    res, err := s.Payments.CancelCashWithdrawal(ctx, operations.CancelCashWithdrawalRequest{
        Signature: "Berkshire",
        SignatureInput: "Louisiana",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        WithdrawalID: "65faef4e-74a7-4700-b0c8-4773396ae72f",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CancelCashWithdrawalRequest](../../models/operations/cancelcashwithdrawalrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.CancelCashWithdrawalResponse](../../models/operations/cancelcashwithdrawalresponse.md), error**


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
        upvestdevsamplesdk.WithSecurity("YOUR_TOKEN"),
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
        Signature: "copying",
        SignatureInput: "toward",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentsWithdrawal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateCashWithdrawalRequest](../../models/operations/createcashwithdrawalrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.CreateCashWithdrawalResponse](../../models/operations/createcashwithdrawalresponse.md), error**


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
        upvestdevsamplesdk.WithSecurity("YOUR_TOKEN"),
    )

    ctx := context.Background()
    res, err := s.Payments.CreateDirectDebit(ctx, operations.CreateDirectDebitRequest{
        RequestBody: &operations.CreateDirectDebitPaymentsDirectDebitCreateRequest{
            AccountGroupID: "2c3ba409-0daa-4766-9bd5-09bb9ce0d60c",
            CashAmount: "East",
            MandateID: "d0c6e1a2-0984-4b8d-9b8b-e22446448d8d",
            UserID: "796d7f3a-6651-47f9-b7f7-2202891b8f25",
        },
        IdempotencyKey: "ccb07f42-4104-44ad-8e1f-c660bb7b269c",
        Signature: "Soap",
        SignatureInput: "Southwest",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateDirectDebit200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateDirectDebitRequest](../../models/operations/createdirectdebitrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.CreateDirectDebitResponse](../../models/operations/createdirectdebitresponse.md), error**


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
        upvestdevsamplesdk.WithSecurity("YOUR_TOKEN"),
    )

    ctx := context.Background()
    res, err := s.Payments.ListCashWithdrawals(ctx, operations.ListCashWithdrawalsRequest{
        AccountGroupID: "d8643156-ad77-42f3-879e-e10e652fca89",
        Signature: "Avon",
        SignatureInput: "Bulgaria",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WithdrawalsListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListCashWithdrawalsRequest](../../models/operations/listcashwithdrawalsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.ListCashWithdrawalsResponse](../../models/operations/listcashwithdrawalsresponse.md), error**


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
        upvestdevsamplesdk.WithSecurity("YOUR_TOKEN"),
    )

    ctx := context.Background()
    res, err := s.Payments.ListDirectDebits(ctx, operations.ListDirectDebitsRequest{
        AccountGroupID: "88b424a5-5c79-48a0-940c-d0211ad65dce",
        Signature: "payment",
        SignatureInput: "compressing",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentsDirectDebitsListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListDirectDebitsRequest](../../models/operations/listdirectdebitsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.ListDirectDebitsResponse](../../models/operations/listdirectdebitsresponse.md), error**


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
        upvestdevsamplesdk.WithSecurity("YOUR_TOKEN"),
    )

    ctx := context.Background()
    res, err := s.Payments.RetrieveCashWithdrawal(ctx, operations.RetrieveCashWithdrawalRequest{
        Signature: "weber",
        SignatureInput: "Koch",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
        WithdrawalID: "f83b155d-8f56-499b-8599-f90ce0450334",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentsWithdrawal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.RetrieveCashWithdrawalRequest](../../models/operations/retrievecashwithdrawalrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.RetrieveCashWithdrawalResponse](../../models/operations/retrievecashwithdrawalresponse.md), error**


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
        upvestdevsamplesdk.WithSecurity("YOUR_TOKEN"),
    )

    ctx := context.Background()
    res, err := s.Payments.RetrieveDirectDebit(ctx, operations.RetrieveDirectDebitRequest{
        DirectDebitID: "b0b3f52f-e687-4685-bdab-d93ae664b9eb",
        Signature: "withdrawal",
        SignatureInput: "Blues",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RetrieveDirectDebit200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.RetrieveDirectDebitRequest](../../models/operations/retrievedirectdebitrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.RetrieveDirectDebitResponse](../../models/operations/retrievedirectdebitresponse.md), error**

