// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/utils"
	"net/http"
	"time"
)

// Config - Configuration of webhook packages collection.
type Config struct {
	// Maximum time of package collection.
	Delay *string `json:"delay,omitempty"`
	// Maximum package size (bytes)
	MaxPackageSize *int64 `json:"max_package_size,omitempty"`
}

func (o *Config) GetDelay() *string {
	if o == nil {
		return nil
	}
	return o.Delay
}

func (o *Config) GetMaxPackageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxPackageSize
}

type CreateWebhookType string

const (
	CreateWebhookTypeAll                      CreateWebhookType = "ALL"
	CreateWebhookTypeUser                     CreateWebhookType = "USER"
	CreateWebhookTypeUserCheck                CreateWebhookType = "USER_CHECK"
	CreateWebhookTypeOrder                    CreateWebhookType = "ORDER"
	CreateWebhookTypeOrderCancellation        CreateWebhookType = "ORDER_CANCELLATION"
	CreateWebhookTypeExecution                CreateWebhookType = "EXECUTION"
	CreateWebhookTypePosition                 CreateWebhookType = "POSITION"
	CreateWebhookTypeCashBalance              CreateWebhookType = "CASH_BALANCE"
	CreateWebhookTypeAccount                  CreateWebhookType = "ACCOUNT"
	CreateWebhookTypeAccountGroup             CreateWebhookType = "ACCOUNT_GROUP"
	CreateWebhookTypeReport                   CreateWebhookType = "REPORT"
	CreateWebhookTypeTreasuryReport           CreateWebhookType = "TREASURY_REPORT"
	CreateWebhookTypeDirectDebit              CreateWebhookType = "DIRECT_DEBIT"
	CreateWebhookTypeWithdrawal               CreateWebhookType = "WITHDRAWAL"
	CreateWebhookTypePortfolio                CreateWebhookType = "PORTFOLIO"
	CreateWebhookTypePortfolioAllocation      CreateWebhookType = "PORTFOLIO_ALLOCATION"
	CreateWebhookTypePortfolioOrder           CreateWebhookType = "PORTFOLIO_ORDER"
	CreateWebhookTypeCorporateAction          CreateWebhookType = "CORPORATE_ACTION"
	CreateWebhookTypeAccountValuation         CreateWebhookType = "ACCOUNT_VALUATION"
	CreateWebhookTypeIntradayAccountValuation CreateWebhookType = "INTRADAY_ACCOUNT_VALUATION"
	CreateWebhookTypeCashTransaction          CreateWebhookType = "CASH_TRANSACTION"
	CreateWebhookTypeSecurityTransaction      CreateWebhookType = "SECURITY_TRANSACTION"
	CreateWebhookTypeAccountLiquidation       CreateWebhookType = "ACCOUNT_LIQUIDATION"
	CreateWebhookTypeAccountReturns           CreateWebhookType = "ACCOUNT_RETURNS"
	CreateWebhookTypeVirtualCashIncrease      CreateWebhookType = "VIRTUAL_CASH_INCREASE"
	CreateWebhookTypeVirtualCashDecrease      CreateWebhookType = "VIRTUAL_CASH_DECREASE"
	CreateWebhookTypeFeeCollection            CreateWebhookType = "FEE_COLLECTION"
)

func (e CreateWebhookType) ToPointer() *CreateWebhookType {
	return &e
}

func (e *CreateWebhookType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ALL":
		fallthrough
	case "USER":
		fallthrough
	case "USER_CHECK":
		fallthrough
	case "ORDER":
		fallthrough
	case "ORDER_CANCELLATION":
		fallthrough
	case "EXECUTION":
		fallthrough
	case "POSITION":
		fallthrough
	case "CASH_BALANCE":
		fallthrough
	case "ACCOUNT":
		fallthrough
	case "ACCOUNT_GROUP":
		fallthrough
	case "REPORT":
		fallthrough
	case "TREASURY_REPORT":
		fallthrough
	case "DIRECT_DEBIT":
		fallthrough
	case "WITHDRAWAL":
		fallthrough
	case "PORTFOLIO":
		fallthrough
	case "PORTFOLIO_ALLOCATION":
		fallthrough
	case "PORTFOLIO_ORDER":
		fallthrough
	case "CORPORATE_ACTION":
		fallthrough
	case "ACCOUNT_VALUATION":
		fallthrough
	case "INTRADAY_ACCOUNT_VALUATION":
		fallthrough
	case "CASH_TRANSACTION":
		fallthrough
	case "SECURITY_TRANSACTION":
		fallthrough
	case "ACCOUNT_LIQUIDATION":
		fallthrough
	case "ACCOUNT_RETURNS":
		fallthrough
	case "VIRTUAL_CASH_INCREASE":
		fallthrough
	case "VIRTUAL_CASH_DECREASE":
		fallthrough
	case "FEE_COLLECTION":
		*e = CreateWebhookType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateWebhookType: %v", v)
	}
}

type CreateWebhookWebhookCreateRequest struct {
	// Configuration of webhook packages collection.
	Config *Config `json:"config,omitempty"`
	// Title of the webhook for use on tenant side.
	Title string `json:"title"`
	// What kind of events to be sent by the webhook.
	Type []CreateWebhookType `json:"type,omitempty"`
	// The callback URL to be called by the webhook.
	URL string `json:"url"`
}

func (o *CreateWebhookWebhookCreateRequest) GetConfig() *Config {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateWebhookWebhookCreateRequest) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *CreateWebhookWebhookCreateRequest) GetType() []CreateWebhookType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *CreateWebhookWebhookCreateRequest) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

type CreateWebhookRequest struct {
	RequestBody *CreateWebhookWebhookCreateRequest `request:"mediaType=application/json"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-http-header
	Signature string `header:"style=simple,explode=false,name=signature"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-input-http-he
	SignatureInput string `header:"style=simple,explode=false,name=signature-input"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
}

func (c CreateWebhookRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateWebhookRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateWebhookRequest) GetRequestBody() *CreateWebhookWebhookCreateRequest {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *CreateWebhookRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *CreateWebhookRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *CreateWebhookRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *CreateWebhookRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// CreateWebhookConfig - Configuration of webhook packages collection.
type CreateWebhookConfig struct {
	// Maximum time of package collection.
	Delay *string `json:"delay,omitempty"`
	// Maximum package size (bytes)
	MaxPackageSize *int64 `json:"max_package_size,omitempty"`
}

func (o *CreateWebhookConfig) GetDelay() *string {
	if o == nil {
		return nil
	}
	return o.Delay
}

func (o *CreateWebhookConfig) GetMaxPackageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxPackageSize
}

type CreateWebhookWebhooksType string

const (
	CreateWebhookWebhooksTypeAll                      CreateWebhookWebhooksType = "ALL"
	CreateWebhookWebhooksTypeUser                     CreateWebhookWebhooksType = "USER"
	CreateWebhookWebhooksTypeUserCheck                CreateWebhookWebhooksType = "USER_CHECK"
	CreateWebhookWebhooksTypeOrder                    CreateWebhookWebhooksType = "ORDER"
	CreateWebhookWebhooksTypeOrderCancellation        CreateWebhookWebhooksType = "ORDER_CANCELLATION"
	CreateWebhookWebhooksTypeExecution                CreateWebhookWebhooksType = "EXECUTION"
	CreateWebhookWebhooksTypePosition                 CreateWebhookWebhooksType = "POSITION"
	CreateWebhookWebhooksTypeCashBalance              CreateWebhookWebhooksType = "CASH_BALANCE"
	CreateWebhookWebhooksTypeAccount                  CreateWebhookWebhooksType = "ACCOUNT"
	CreateWebhookWebhooksTypeAccountGroup             CreateWebhookWebhooksType = "ACCOUNT_GROUP"
	CreateWebhookWebhooksTypeReport                   CreateWebhookWebhooksType = "REPORT"
	CreateWebhookWebhooksTypeTreasuryReport           CreateWebhookWebhooksType = "TREASURY_REPORT"
	CreateWebhookWebhooksTypeDirectDebit              CreateWebhookWebhooksType = "DIRECT_DEBIT"
	CreateWebhookWebhooksTypeWithdrawal               CreateWebhookWebhooksType = "WITHDRAWAL"
	CreateWebhookWebhooksTypePortfolio                CreateWebhookWebhooksType = "PORTFOLIO"
	CreateWebhookWebhooksTypePortfolioAllocation      CreateWebhookWebhooksType = "PORTFOLIO_ALLOCATION"
	CreateWebhookWebhooksTypePortfolioOrder           CreateWebhookWebhooksType = "PORTFOLIO_ORDER"
	CreateWebhookWebhooksTypeCorporateAction          CreateWebhookWebhooksType = "CORPORATE_ACTION"
	CreateWebhookWebhooksTypeAccountValuation         CreateWebhookWebhooksType = "ACCOUNT_VALUATION"
	CreateWebhookWebhooksTypeIntradayAccountValuation CreateWebhookWebhooksType = "INTRADAY_ACCOUNT_VALUATION"
	CreateWebhookWebhooksTypeCashTransaction          CreateWebhookWebhooksType = "CASH_TRANSACTION"
	CreateWebhookWebhooksTypeSecurityTransaction      CreateWebhookWebhooksType = "SECURITY_TRANSACTION"
	CreateWebhookWebhooksTypeAccountLiquidation       CreateWebhookWebhooksType = "ACCOUNT_LIQUIDATION"
	CreateWebhookWebhooksTypeAccountReturns           CreateWebhookWebhooksType = "ACCOUNT_RETURNS"
	CreateWebhookWebhooksTypeVirtualCashIncrease      CreateWebhookWebhooksType = "VIRTUAL_CASH_INCREASE"
	CreateWebhookWebhooksTypeVirtualCashDecrease      CreateWebhookWebhooksType = "VIRTUAL_CASH_DECREASE"
	CreateWebhookWebhooksTypeFeeCollection            CreateWebhookWebhooksType = "FEE_COLLECTION"
)

func (e CreateWebhookWebhooksType) ToPointer() *CreateWebhookWebhooksType {
	return &e
}

func (e *CreateWebhookWebhooksType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ALL":
		fallthrough
	case "USER":
		fallthrough
	case "USER_CHECK":
		fallthrough
	case "ORDER":
		fallthrough
	case "ORDER_CANCELLATION":
		fallthrough
	case "EXECUTION":
		fallthrough
	case "POSITION":
		fallthrough
	case "CASH_BALANCE":
		fallthrough
	case "ACCOUNT":
		fallthrough
	case "ACCOUNT_GROUP":
		fallthrough
	case "REPORT":
		fallthrough
	case "TREASURY_REPORT":
		fallthrough
	case "DIRECT_DEBIT":
		fallthrough
	case "WITHDRAWAL":
		fallthrough
	case "PORTFOLIO":
		fallthrough
	case "PORTFOLIO_ALLOCATION":
		fallthrough
	case "PORTFOLIO_ORDER":
		fallthrough
	case "CORPORATE_ACTION":
		fallthrough
	case "ACCOUNT_VALUATION":
		fallthrough
	case "INTRADAY_ACCOUNT_VALUATION":
		fallthrough
	case "CASH_TRANSACTION":
		fallthrough
	case "SECURITY_TRANSACTION":
		fallthrough
	case "ACCOUNT_LIQUIDATION":
		fallthrough
	case "ACCOUNT_RETURNS":
		fallthrough
	case "VIRTUAL_CASH_INCREASE":
		fallthrough
	case "VIRTUAL_CASH_DECREASE":
		fallthrough
	case "FEE_COLLECTION":
		*e = CreateWebhookWebhooksType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateWebhookWebhooksType: %v", v)
	}
}

// CreateWebhookWebhook - Returns a webhook subscription object if a valid webhook subscription object ID was provided.
type CreateWebhookWebhook struct {
	// Configuration of webhook packages collection.
	Config CreateWebhookConfig `json:"config"`
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	// Enable/disable webhook.
	Enabled bool `json:"enabled"`
	// Webhook unique identifier.
	ID string `json:"id"`
	// Title of the webhook for use on tenant side.
	Title string `json:"title"`
	// What kind of events to be sent by the webhook.
	Type []CreateWebhookWebhooksType `json:"type"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time `json:"updated_at"`
	// The callback URL to be called by the webhook.
	URL string `json:"url"`
}

func (c CreateWebhookWebhook) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateWebhookWebhook) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateWebhookWebhook) GetConfig() CreateWebhookConfig {
	if o == nil {
		return CreateWebhookConfig{}
	}
	return o.Config
}

func (o *CreateWebhookWebhook) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *CreateWebhookWebhook) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *CreateWebhookWebhook) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateWebhookWebhook) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *CreateWebhookWebhook) GetType() []CreateWebhookWebhooksType {
	if o == nil {
		return []CreateWebhookWebhooksType{}
	}
	return o.Type
}

func (o *CreateWebhookWebhook) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *CreateWebhookWebhook) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

type CreateWebhookResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Returns a webhook subscription object if a valid webhook subscription object ID was provided.
	Webhook *CreateWebhookWebhook
}

func (o *CreateWebhookResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateWebhookResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *CreateWebhookResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateWebhookResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateWebhookResponse) GetWebhook() *CreateWebhookWebhook {
	if o == nil {
		return nil
	}
	return o.Webhook
}
