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

// UpdateWebhookConfig - Configuration of webhook packages collection.
type UpdateWebhookConfig struct {
	// Maximum time of package collection.
	Delay *string `json:"delay,omitempty"`
	// Maximum package size (bytes)
	MaxPackageSize *int64 `json:"max_package_size,omitempty"`
}

func (o *UpdateWebhookConfig) GetDelay() *string {
	if o == nil {
		return nil
	}
	return o.Delay
}

func (o *UpdateWebhookConfig) GetMaxPackageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxPackageSize
}

type UpdateWebhookType string

const (
	UpdateWebhookTypeAll                      UpdateWebhookType = "ALL"
	UpdateWebhookTypeUser                     UpdateWebhookType = "USER"
	UpdateWebhookTypeUserCheck                UpdateWebhookType = "USER_CHECK"
	UpdateWebhookTypeOrder                    UpdateWebhookType = "ORDER"
	UpdateWebhookTypeOrderCancellation        UpdateWebhookType = "ORDER_CANCELLATION"
	UpdateWebhookTypeExecution                UpdateWebhookType = "EXECUTION"
	UpdateWebhookTypePosition                 UpdateWebhookType = "POSITION"
	UpdateWebhookTypeCashBalance              UpdateWebhookType = "CASH_BALANCE"
	UpdateWebhookTypeAccount                  UpdateWebhookType = "ACCOUNT"
	UpdateWebhookTypeAccountGroup             UpdateWebhookType = "ACCOUNT_GROUP"
	UpdateWebhookTypeReport                   UpdateWebhookType = "REPORT"
	UpdateWebhookTypeTreasuryReport           UpdateWebhookType = "TREASURY_REPORT"
	UpdateWebhookTypeDirectDebit              UpdateWebhookType = "DIRECT_DEBIT"
	UpdateWebhookTypeWithdrawal               UpdateWebhookType = "WITHDRAWAL"
	UpdateWebhookTypePortfolio                UpdateWebhookType = "PORTFOLIO"
	UpdateWebhookTypePortfolioAllocation      UpdateWebhookType = "PORTFOLIO_ALLOCATION"
	UpdateWebhookTypePortfolioOrder           UpdateWebhookType = "PORTFOLIO_ORDER"
	UpdateWebhookTypeCorporateAction          UpdateWebhookType = "CORPORATE_ACTION"
	UpdateWebhookTypeAccountValuation         UpdateWebhookType = "ACCOUNT_VALUATION"
	UpdateWebhookTypeIntradayAccountValuation UpdateWebhookType = "INTRADAY_ACCOUNT_VALUATION"
	UpdateWebhookTypeCashTransaction          UpdateWebhookType = "CASH_TRANSACTION"
	UpdateWebhookTypeSecurityTransaction      UpdateWebhookType = "SECURITY_TRANSACTION"
	UpdateWebhookTypeAccountLiquidation       UpdateWebhookType = "ACCOUNT_LIQUIDATION"
	UpdateWebhookTypeAccountReturns           UpdateWebhookType = "ACCOUNT_RETURNS"
	UpdateWebhookTypeVirtualCashIncrease      UpdateWebhookType = "VIRTUAL_CASH_INCREASE"
	UpdateWebhookTypeVirtualCashDecrease      UpdateWebhookType = "VIRTUAL_CASH_DECREASE"
	UpdateWebhookTypeFeeCollection            UpdateWebhookType = "FEE_COLLECTION"
)

func (e UpdateWebhookType) ToPointer() *UpdateWebhookType {
	return &e
}

func (e *UpdateWebhookType) UnmarshalJSON(data []byte) error {
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
		*e = UpdateWebhookType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateWebhookType: %v", v)
	}
}

type UpdateWebhookWebhookUpdateRequest struct {
	// Configuration of webhook packages collection.
	Config *UpdateWebhookConfig `json:"config,omitempty"`
	// Enable/disable webhook.
	Enabled *bool `json:"enabled,omitempty"`
	// Title of the webhook for use on tenant side.
	Title *string `json:"title,omitempty"`
	// What kind of events to be sent by the webhook.
	Type []UpdateWebhookType `json:"type,omitempty"`
	// The callback URL to be called by the webhook.
	URL *string `json:"url,omitempty"`
}

func (o *UpdateWebhookWebhookUpdateRequest) GetConfig() *UpdateWebhookConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *UpdateWebhookWebhookUpdateRequest) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *UpdateWebhookWebhookUpdateRequest) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *UpdateWebhookWebhookUpdateRequest) GetType() []UpdateWebhookType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *UpdateWebhookWebhookUpdateRequest) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

type UpdateWebhookRequest struct {
	RequestBody *UpdateWebhookWebhookUpdateRequest `request:"mediaType=application/json"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-http-header
	Signature string `header:"style=simple,explode=false,name=signature"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-input-http-he
	SignatureInput string `header:"style=simple,explode=false,name=signature-input"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
	// Webhook identifier
	WebhookID string `pathParam:"style=simple,explode=false,name=webhook_id"`
}

func (u UpdateWebhookRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateWebhookRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateWebhookRequest) GetRequestBody() *UpdateWebhookWebhookUpdateRequest {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *UpdateWebhookRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *UpdateWebhookRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *UpdateWebhookRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *UpdateWebhookRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

func (o *UpdateWebhookRequest) GetWebhookID() string {
	if o == nil {
		return ""
	}
	return o.WebhookID
}

// UpdateWebhookWebhooksConfig - Configuration of webhook packages collection.
type UpdateWebhookWebhooksConfig struct {
	// Maximum time of package collection.
	Delay *string `json:"delay,omitempty"`
	// Maximum package size (bytes)
	MaxPackageSize *int64 `json:"max_package_size,omitempty"`
}

func (o *UpdateWebhookWebhooksConfig) GetDelay() *string {
	if o == nil {
		return nil
	}
	return o.Delay
}

func (o *UpdateWebhookWebhooksConfig) GetMaxPackageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxPackageSize
}

type UpdateWebhookWebhooksType string

const (
	UpdateWebhookWebhooksTypeAll                      UpdateWebhookWebhooksType = "ALL"
	UpdateWebhookWebhooksTypeUser                     UpdateWebhookWebhooksType = "USER"
	UpdateWebhookWebhooksTypeUserCheck                UpdateWebhookWebhooksType = "USER_CHECK"
	UpdateWebhookWebhooksTypeOrder                    UpdateWebhookWebhooksType = "ORDER"
	UpdateWebhookWebhooksTypeOrderCancellation        UpdateWebhookWebhooksType = "ORDER_CANCELLATION"
	UpdateWebhookWebhooksTypeExecution                UpdateWebhookWebhooksType = "EXECUTION"
	UpdateWebhookWebhooksTypePosition                 UpdateWebhookWebhooksType = "POSITION"
	UpdateWebhookWebhooksTypeCashBalance              UpdateWebhookWebhooksType = "CASH_BALANCE"
	UpdateWebhookWebhooksTypeAccount                  UpdateWebhookWebhooksType = "ACCOUNT"
	UpdateWebhookWebhooksTypeAccountGroup             UpdateWebhookWebhooksType = "ACCOUNT_GROUP"
	UpdateWebhookWebhooksTypeReport                   UpdateWebhookWebhooksType = "REPORT"
	UpdateWebhookWebhooksTypeTreasuryReport           UpdateWebhookWebhooksType = "TREASURY_REPORT"
	UpdateWebhookWebhooksTypeDirectDebit              UpdateWebhookWebhooksType = "DIRECT_DEBIT"
	UpdateWebhookWebhooksTypeWithdrawal               UpdateWebhookWebhooksType = "WITHDRAWAL"
	UpdateWebhookWebhooksTypePortfolio                UpdateWebhookWebhooksType = "PORTFOLIO"
	UpdateWebhookWebhooksTypePortfolioAllocation      UpdateWebhookWebhooksType = "PORTFOLIO_ALLOCATION"
	UpdateWebhookWebhooksTypePortfolioOrder           UpdateWebhookWebhooksType = "PORTFOLIO_ORDER"
	UpdateWebhookWebhooksTypeCorporateAction          UpdateWebhookWebhooksType = "CORPORATE_ACTION"
	UpdateWebhookWebhooksTypeAccountValuation         UpdateWebhookWebhooksType = "ACCOUNT_VALUATION"
	UpdateWebhookWebhooksTypeIntradayAccountValuation UpdateWebhookWebhooksType = "INTRADAY_ACCOUNT_VALUATION"
	UpdateWebhookWebhooksTypeCashTransaction          UpdateWebhookWebhooksType = "CASH_TRANSACTION"
	UpdateWebhookWebhooksTypeSecurityTransaction      UpdateWebhookWebhooksType = "SECURITY_TRANSACTION"
	UpdateWebhookWebhooksTypeAccountLiquidation       UpdateWebhookWebhooksType = "ACCOUNT_LIQUIDATION"
	UpdateWebhookWebhooksTypeAccountReturns           UpdateWebhookWebhooksType = "ACCOUNT_RETURNS"
	UpdateWebhookWebhooksTypeVirtualCashIncrease      UpdateWebhookWebhooksType = "VIRTUAL_CASH_INCREASE"
	UpdateWebhookWebhooksTypeVirtualCashDecrease      UpdateWebhookWebhooksType = "VIRTUAL_CASH_DECREASE"
	UpdateWebhookWebhooksTypeFeeCollection            UpdateWebhookWebhooksType = "FEE_COLLECTION"
)

func (e UpdateWebhookWebhooksType) ToPointer() *UpdateWebhookWebhooksType {
	return &e
}

func (e *UpdateWebhookWebhooksType) UnmarshalJSON(data []byte) error {
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
		*e = UpdateWebhookWebhooksType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateWebhookWebhooksType: %v", v)
	}
}

// UpdateWebhookWebhook - Returns a webhook subscription object if a valid webhook subscription object ID was provided.
type UpdateWebhookWebhook struct {
	// Configuration of webhook packages collection.
	Config UpdateWebhookWebhooksConfig `json:"config"`
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	// Enable/disable webhook.
	Enabled bool `json:"enabled"`
	// Webhook unique identifier.
	ID string `json:"id"`
	// Title of the webhook for use on tenant side.
	Title string `json:"title"`
	// What kind of events to be sent by the webhook.
	Type []UpdateWebhookWebhooksType `json:"type"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time `json:"updated_at"`
	// The callback URL to be called by the webhook.
	URL string `json:"url"`
}

func (u UpdateWebhookWebhook) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateWebhookWebhook) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateWebhookWebhook) GetConfig() UpdateWebhookWebhooksConfig {
	if o == nil {
		return UpdateWebhookWebhooksConfig{}
	}
	return o.Config
}

func (o *UpdateWebhookWebhook) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *UpdateWebhookWebhook) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *UpdateWebhookWebhook) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateWebhookWebhook) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *UpdateWebhookWebhook) GetType() []UpdateWebhookWebhooksType {
	if o == nil {
		return []UpdateWebhookWebhooksType{}
	}
	return o.Type
}

func (o *UpdateWebhookWebhook) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *UpdateWebhookWebhook) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

type UpdateWebhookResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Returns a webhook subscription object if a valid webhook subscription object ID was provided.
	Webhook *UpdateWebhookWebhook
}

func (o *UpdateWebhookResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateWebhookResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *UpdateWebhookResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateWebhookResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateWebhookResponse) GetWebhook() *UpdateWebhookWebhook {
	if o == nil {
		return nil
	}
	return o.Webhook
}
