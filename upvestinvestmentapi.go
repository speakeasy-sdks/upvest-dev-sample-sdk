// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package upvestdevsamplesdk

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/internal/hooks"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// Sandbox environment
	"https://sandbox.upvest.co",
	// Live environment
	"https://api.upvest.co",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
	Hooks             *hooks.Hooks
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// UpvestInvestmentAPI - Upvest Investment API: Upvest Investment API.
type UpvestInvestmentAPI struct {
	// All accounts related paths
	Accounts *Accounts
	// All cash balance related paths
	CashBalances *CashBalances
	// All payments related paths
	Payments *Payments
	// All accounts liquidations related paths.
	Liquidations *Liquidations
	// All order related paths.
	Orders *Orders
	// All positions related paths.
	Positions *Positions
	// All accounts returns related paths.
	Returns *Returns
	// All valuations related paths.
	Valuations *Valuations
	// All authentication related paths.
	AccessTokens *AccessTokens
	// All webhook related paths.
	Webhooks *Webhooks
	// All fees related paths.
	Fees *Fees
	// All instrument related paths.
	Instruments *Instruments
	// All direct debit mandates related paths
	Mandates *Mandates
	// All reference account related paths
	ReferenceAccounts *ReferenceAccounts
	// All portfolios related paths.
	Portfolios *Portfolios
	// All reports related paths.
	Reports *Reports
	// All transactions related paths.
	Transactions *Transactions
	// All user related paths.
	Users *Users
	// All taxes related paths.
	Taxes *Taxes
	// All virtual cash balances related paths
	VirtualCashBalances *VirtualCashBalances

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*UpvestInvestmentAPI)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *UpvestInvestmentAPI) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *UpvestInvestmentAPI) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *UpvestInvestmentAPI) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *UpvestInvestmentAPI) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

func withSecurity(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return &security, nil
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(oauthClientCredentials string) SDKOption {
	return func(sdk *UpvestInvestmentAPI) {
		security := shared.Security{OauthClientCredentials: &oauthClientCredentials}
		sdk.sdkConfiguration.Security = withSecurity(&security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *UpvestInvestmentAPI) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *UpvestInvestmentAPI) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *UpvestInvestmentAPI {
	sdk := &UpvestInvestmentAPI{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.9.0",
			SDKVersion:        "0.9.1",
			GenVersion:        "2.263.3",
			UserAgent:         "speakeasy-sdk/go 0.9.1 2.263.3 1.9.0 github.com/speakeasy-sdks/upvest-dev-sample-sdk",
			Hooks:             hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	sdk.sdkConfiguration.DefaultClient = sdk.sdkConfiguration.Hooks.ClientInit(sdk.sdkConfiguration.DefaultClient)

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.Accounts = newAccounts(sdk.sdkConfiguration)

	sdk.CashBalances = newCashBalances(sdk.sdkConfiguration)

	sdk.Payments = newPayments(sdk.sdkConfiguration)

	sdk.Liquidations = newLiquidations(sdk.sdkConfiguration)

	sdk.Orders = newOrders(sdk.sdkConfiguration)

	sdk.Positions = newPositions(sdk.sdkConfiguration)

	sdk.Returns = newReturns(sdk.sdkConfiguration)

	sdk.Valuations = newValuations(sdk.sdkConfiguration)

	sdk.AccessTokens = newAccessTokens(sdk.sdkConfiguration)

	sdk.Webhooks = newWebhooks(sdk.sdkConfiguration)

	sdk.Fees = newFees(sdk.sdkConfiguration)

	sdk.Instruments = newInstruments(sdk.sdkConfiguration)

	sdk.Mandates = newMandates(sdk.sdkConfiguration)

	sdk.ReferenceAccounts = newReferenceAccounts(sdk.sdkConfiguration)

	sdk.Portfolios = newPortfolios(sdk.sdkConfiguration)

	sdk.Reports = newReports(sdk.sdkConfiguration)

	sdk.Transactions = newTransactions(sdk.sdkConfiguration)

	sdk.Users = newUsers(sdk.sdkConfiguration)

	sdk.Taxes = newTaxes(sdk.sdkConfiguration)

	sdk.VirtualCashBalances = newVirtualCashBalances(sdk.sdkConfiguration)

	return sdk
}
