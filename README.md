# Upvest Sample Go SDK

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://github.com/speakeasy-sdks/upvest-dev-sample-sdk.git/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/upvest-dev-sample-sdk/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
    
</div>

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/upvest-dev-sample-sdk
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
### Example

```go
package main

import (
	"context"
	upvestdevsamplesdk "github.com/speakeasy-sdks/upvest-dev-sample-sdk"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
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
			Type:           operations.TypeTrading,
			UserID:         "d04cd2d5-ae02-4bb1-9118-75a95a0f2373",
		},
		IdempotencyKey:   "ccb07f42-4104-44ad-8e1f-c660bb7b269c",
		Signature:        "string",
		SignatureInput:   "string",
		UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
		UpvestClientID:   "ebabcf4d-61c3-4942-875c-e265a7c2d062",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Account != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Accounts](docs/sdks/accounts/README.md)

* [AccountClosure](docs/sdks/accounts/README.md#accountclosure) - Close a user account by ID
* [AccountGroupClosure](docs/sdks/accounts/README.md#accountgroupclosure) - Close an account group by ID
* [CreateAccount](docs/sdks/accounts/README.md#createaccount) - Create an account
* [CreateAccountGroup](docs/sdks/accounts/README.md#createaccountgroup) - Create an account group
* [ListAccountGroups](docs/sdks/accounts/README.md#listaccountgroups) - Get account groups
* [ListAccounts](docs/sdks/accounts/README.md#listaccounts) - Get accounts
* [RetrieveAccount](docs/sdks/accounts/README.md#retrieveaccount) - Get a user account by ID
* [RetrieveAccountGroup](docs/sdks/accounts/README.md#retrieveaccountgroup) - Get an account group by ID
* [UpdateAccount](docs/sdks/accounts/README.md#updateaccount) - Update user account

### [CashBalances](docs/sdks/cashbalances/README.md)

* [RetrieveCashBalance](docs/sdks/cashbalances/README.md#retrievecashbalance) - Retrieve an account group's cash balance

### [Payments](docs/sdks/payments/README.md)

* [CancelCashWithdrawal](docs/sdks/payments/README.md#cancelcashwithdrawal) - Cancel withdrawal by ID
* [CreateCashWithdrawal](docs/sdks/payments/README.md#createcashwithdrawal) - Trigger a withdrawal
* [CreateDirectDebit](docs/sdks/payments/README.md#createdirectdebit) - Trigger a direct debit
* [ListCashWithdrawals](docs/sdks/payments/README.md#listcashwithdrawals) - List withdrawals
* [ListDirectDebits](docs/sdks/payments/README.md#listdirectdebits) - List direct debits
* [RetrieveCashWithdrawal](docs/sdks/payments/README.md#retrievecashwithdrawal) - Retrieve withdrawal
* [RetrieveDirectDebit](docs/sdks/payments/README.md#retrievedirectdebit) - Retrieve a direct debit

### [Liquidations](docs/sdks/liquidations/README.md)

* [CancelAccountLiquidation](docs/sdks/liquidations/README.md#cancelaccountliquidation) - Cancel account liquidation
* [CreateAccountLiquidation](docs/sdks/liquidations/README.md#createaccountliquidation) - Create account liquidation request
* [ListAccountsLiquidations](docs/sdks/liquidations/README.md#listaccountsliquidations) - List accounts liquidations
* [RetrieveAccountLiquidation](docs/sdks/liquidations/README.md#retrieveaccountliquidation) - Retrieve account liquidation

### [Orders](docs/sdks/orders/README.md)

* [CancelOrder](docs/sdks/orders/README.md#cancelorder) - Cancel an order by ID
* [ListAccountOrders](docs/sdks/orders/README.md#listaccountorders) - Get orders for an account by ID
* [PlaceOrder](docs/sdks/orders/README.md#placeorder) - Place an order
* [RetrieveOrder](docs/sdks/orders/README.md#retrieveorder) - Get an order by ID
* [RetrieveOrderExecution](docs/sdks/orders/README.md#retrieveorderexecution) - Get an order execution by ID

### [Positions](docs/sdks/positions/README.md)

* [ListPositions](docs/sdks/positions/README.md#listpositions) - List positions
* [RetrievePosition](docs/sdks/positions/README.md#retrieveposition) - Retrieve position

### [Returns](docs/sdks/returns/README.md)

* [ListAccountReturns](docs/sdks/returns/README.md#listaccountreturns) - List account returns

### [Valuations](docs/sdks/valuations/README.md)

* [GetAccountValuation](docs/sdks/valuations/README.md#getaccountvaluation) - Get current valuation for an account
* [ListAccountValuationHistory](docs/sdks/valuations/README.md#listaccountvaluationhistory) - List valuation history for an account

### [AccessTokens](docs/sdks/accesstokens/README.md)

* [IssueToken](docs/sdks/accesstokens/README.md#issuetoken) - Get an access token for requested scopes

### [Webhooks](docs/sdks/webhooks/README.md)

* [CreateWebhook](docs/sdks/webhooks/README.md#createwebhook) - Create a webhook subscription
* [DeleteWebhook](docs/sdks/webhooks/README.md#deletewebhook) - Delete a webhook subscription
* [GetJwks](docs/sdks/webhooks/README.md#getjwks) - Get signing keys
* [ListWebhooks](docs/sdks/webhooks/README.md#listwebhooks) - List all webhooks
* [RetrieveWebhook](docs/sdks/webhooks/README.md#retrievewebhook) - Retrieve a webhook subscription
* [TestWebhook](docs/sdks/webhooks/README.md#testwebhook) - Test a webhook
* [UpdateWebhook](docs/sdks/webhooks/README.md#updatewebhook) - Update a webhook subscription

### [Fees](docs/sdks/fees/README.md)

* [CreateFeeCollection](docs/sdks/fees/README.md#createfeecollection) - Create a fee collection
* [ListFeeCollections](docs/sdks/fees/README.md#listfeecollections) - Get fee collections
* [RetrieveFeeCollection](docs/sdks/fees/README.md#retrievefeecollection) - Get a fee collection by ID

### [Instruments](docs/sdks/instruments/README.md)

* [ListInstruments](docs/sdks/instruments/README.md#listinstruments) - List instruments
* [RetrieveInstrument](docs/sdks/instruments/README.md#retrieveinstrument) - Retrieve instrument

### [Mandates](docs/sdks/mandates/README.md)

* [CreateMandate](docs/sdks/mandates/README.md#createmandate) - Create a mandate
* [DeleteMandate](docs/sdks/mandates/README.md#deletemandate) - Delete mandate
* [ListMandates](docs/sdks/mandates/README.md#listmandates) - List mandates
* [RetrieveMandate](docs/sdks/mandates/README.md#retrievemandate) - Retrieve a direct debit mandate

### [ReferenceAccounts](docs/sdks/referenceaccounts/README.md)

* [CreateReferenceAccount](docs/sdks/referenceaccounts/README.md#createreferenceaccount) - Create a reference account
* [DeleteReferenceAccount](docs/sdks/referenceaccounts/README.md#deletereferenceaccount) - Delete a reference account by ID
* [ListReferenceAccounts](docs/sdks/referenceaccounts/README.md#listreferenceaccounts) - Get reference accounts of a user
* [RetrieveReferenceAccount](docs/sdks/referenceaccounts/README.md#retrievereferenceaccount) - Get a reference account by ID

### [Portfolios](docs/sdks/portfolios/README.md)

* [CancelPortfoliosOrder](docs/sdks/portfolios/README.md#cancelportfoliosorder) - Cancel portfolios order
* [CreatePortfoliosAllocation](docs/sdks/portfolios/README.md#createportfoliosallocation) - Create portfolios allocation
* [CreatePortfoliosConfiguration](docs/sdks/portfolios/README.md#createportfoliosconfiguration) - Create portfolios configuration
* [CreatePortfoliosOrder](docs/sdks/portfolios/README.md#createportfoliosorder) - Create portfolios order
* [CreatePortfoliosRebalancingStrategy](docs/sdks/portfolios/README.md#createportfoliosrebalancingstrategy) - Create portfolios rebalancing strategy
* [ListPortfolioRebalancingExecutionOrders](docs/sdks/portfolios/README.md#listportfoliorebalancingexecutionorders) - List portfolio rebalancing execution orders
* [ListPortfoliosAllocationAccounts](docs/sdks/portfolios/README.md#listportfoliosallocationaccounts) - List portfolios allocation accounts
* [ListPortfoliosAllocations](docs/sdks/portfolios/README.md#listportfoliosallocations) - List portfolios allocations
* [ListPortfoliosConfigurations](docs/sdks/portfolios/README.md#listportfoliosconfigurations) - List portfolios configurations
* [ListPortfoliosOrders](docs/sdks/portfolios/README.md#listportfoliosorders) - List portfolios orders
* [ListPortfoliosRebalancingStrategies](docs/sdks/portfolios/README.md#listportfoliosrebalancingstrategies) - List portfolios rebalancing strategies
* [RetrievePortfoliosAllocation](docs/sdks/portfolios/README.md#retrieveportfoliosallocation) - Retrieve portfolios allocation
* [RetrievePortfoliosConfiguration](docs/sdks/portfolios/README.md#retrieveportfoliosconfiguration) - Retrieve portfolios configuration
* [RetrievePortfoliosOrder](docs/sdks/portfolios/README.md#retrieveportfoliosorder) - Retrieve portfolios order
* [RetrievePortfoliosRebalancingExecution](docs/sdks/portfolios/README.md#retrieveportfoliosrebalancingexecution) - Retrieve portfolios rebalancing execution
* [RetrievePortfoliosRebalancingStrategy](docs/sdks/portfolios/README.md#retrieveportfoliosrebalancingstrategy) - Retrieve portfolios rebalancing strategy
* [TriggerPortfolioRebalancing](docs/sdks/portfolios/README.md#triggerportfoliorebalancing) - Trigger portfolio rebalancing
* [UpdatePortfoliosAllocation](docs/sdks/portfolios/README.md#updateportfoliosallocation) - Update portfolios allocation
* [UpdatePortfoliosConfiguration](docs/sdks/portfolios/README.md#updateportfoliosconfiguration) - Update portfolios configuration

### [Reports](docs/sdks/reports/README.md)

* [ListReports](docs/sdks/reports/README.md#listreports) - List user reports
* [RetrieveReport](docs/sdks/reports/README.md#retrievereport) - Retrieve a user report

### [Transactions](docs/sdks/transactions/README.md)

* [ListCashTransactions](docs/sdks/transactions/README.md#listcashtransactions) - List cash transactions
* [ListSecuritiesTransactions](docs/sdks/transactions/README.md#listsecuritiestransactions) - List securities transactions

### [Users](docs/sdks/users/README.md)

* [CreateIdentifier](docs/sdks/users/README.md#createidentifier) - Create a user identifier
* [CreateUser](docs/sdks/users/README.md#createuser) - Create a user
* [CreateUserCheck](docs/sdks/users/README.md#createusercheck) - Create a new check for a user
* [ListUserAccountGroups](docs/sdks/users/README.md#listuseraccountgroups) - Get user account groups
* [ListUserAccounts](docs/sdks/users/README.md#listuseraccounts) - Get user accounts
* [ListUserChecks](docs/sdks/users/README.md#listuserchecks) - Get checks for a user
* [ListUserIdentifiers](docs/sdks/users/README.md#listuseridentifiers) - Get user identifiers
* [ListUsers](docs/sdks/users/README.md#listusers) - Get all users
* [OffboardUser](docs/sdks/users/README.md#offboarduser) - Offboard a user
* [RetrieveIdentifier](docs/sdks/users/README.md#retrieveidentifier) - Get a user identifier by ID
* [RetrieveUser](docs/sdks/users/README.md#retrieveuser) - Get a user by ID
* [RetrieveUserCheck](docs/sdks/users/README.md#retrieveusercheck) - Get a user check by ID
* [UpdateIdentifier](docs/sdks/users/README.md#updateidentifier) - Update a user identifier by ID
* [UserDataChange](docs/sdks/users/README.md#userdatachange) - Change user data

### [Taxes](docs/sdks/taxes/README.md)

* [RetrieveTaxResidencies](docs/sdks/taxes/README.md#retrievetaxresidencies) - Retrieve tax residencies
* [SetTaxResidencies](docs/sdks/taxes/README.md#settaxresidencies) - Update tax residencies

### [VirtualCashBalances](docs/sdks/virtualcashbalances/README.md)

* [CancelVirtualCashDecrease](docs/sdks/virtualcashbalances/README.md#cancelvirtualcashdecrease) - Cancel virtual cash decrease by ID
* [CreateVirtualCashDecrease](docs/sdks/virtualcashbalances/README.md#createvirtualcashdecrease) - Trigger a virtual cash decrease
* [CreateVirtualCashIncrease](docs/sdks/virtualcashbalances/README.md#createvirtualcashincrease) - Trigger a virtual cash increase
<!-- End SDK Available Operations -->

<!-- Start Dev Containers -->

<!-- End Dev Containers -->

<!-- Start Go Types -->
# Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

## Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
<!-- End Go Types -->



<!-- Start Error Handling -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

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

### Example

```go
package main

import (
	"context"
	"errors"
	upvestdevsamplesdk "github.com/speakeasy-sdks/upvest-dev-sample-sdk"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"log"
)

func main() {
	s := upvestdevsamplesdk.New(
		upvestdevsamplesdk.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Accounts.AccountClosure(ctx, operations.AccountClosureRequest{
		AccountID:        "87f46f4c-298e-4960-b531-5043c3be9e8d",
		Signature:        "string",
		SignatureInput:   "string",
		UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
		UpvestClientID:   "ebabcf4d-61c3-4942-875c-e265a7c2d062",
	})
	if err != nil {

		var e *sdkerrors.AccountClosureError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.AccountClosureAccountsError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.AccountClosureAccountsResponseError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.AccountClosureAccountsResponse406Error
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.AccountClosureAccountsResponse409Error
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.AccountClosureAccountsResponse429Error
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.AccountClosureAccountsResponse500Error
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.AccountClosureAccountsResponse503Error
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.AccountClosureAccountsResponse504Error
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling -->



<!-- Start Server Selection -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://sandbox.upvest.co` | None |
| 1 | `https://api.upvest.co` | None |

#### Example

```go
package main

import (
	"context"
	upvestdevsamplesdk "github.com/speakeasy-sdks/upvest-dev-sample-sdk"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"log"
	"net/http"
)

func main() {
	s := upvestdevsamplesdk.New(
		upvestdevsamplesdk.WithServerIndex(1),
		upvestdevsamplesdk.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Accounts.AccountClosure(ctx, operations.AccountClosureRequest{
		AccountID:        "87f46f4c-298e-4960-b531-5043c3be9e8d",
		Signature:        "string",
		SignatureInput:   "string",
		UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
		UpvestClientID:   "ebabcf4d-61c3-4942-875c-e265a7c2d062",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	upvestdevsamplesdk "github.com/speakeasy-sdks/upvest-dev-sample-sdk"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"log"
	"net/http"
)

func main() {
	s := upvestdevsamplesdk.New(
		upvestdevsamplesdk.WithServerURL("https://sandbox.upvest.co"),
		upvestdevsamplesdk.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Accounts.AccountClosure(ctx, operations.AccountClosureRequest{
		AccountID:        "87f46f4c-298e-4960-b531-5043c3be9e8d",
		Signature:        "string",
		SignatureInput:   "string",
		UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
		UpvestClientID:   "ebabcf4d-61c3-4942-875c-e265a7c2d062",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		// handle response
	}
}

```
<!-- End Server Selection -->



<!-- Start Custom HTTP Client -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client -->



<!-- Start Authentication -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name                     | Type                     | Scheme                   |
| ------------------------ | ------------------------ | ------------------------ |
| `OauthClientCredentials` | oauth2                   | OAuth2 token             |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	upvestdevsamplesdk "github.com/speakeasy-sdks/upvest-dev-sample-sdk"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"log"
	"net/http"
)

func main() {
	s := upvestdevsamplesdk.New(
		upvestdevsamplesdk.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Accounts.AccountClosure(ctx, operations.AccountClosureRequest{
		AccountID:        "87f46f4c-298e-4960-b531-5043c3be9e8d",
		Signature:        "string",
		SignatureInput:   "string",
		UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
		UpvestClientID:   "ebabcf4d-61c3-4942-875c-e265a7c2d062",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		// handle response
	}
}

```
<!-- End Authentication -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
