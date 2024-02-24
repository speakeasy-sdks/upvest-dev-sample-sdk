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

type CreateAccountLiquidationSecurity struct {
	OauthClientCredentials string `security:"scheme,type=oauth2,name=Authorization"`
}

func (o *CreateAccountLiquidationSecurity) GetOauthClientCredentials() string {
	if o == nil {
		return ""
	}
	return o.OauthClientCredentials
}

type CreateAccountLiquidationAccountLiquidationRequest struct {
	// User unique identifier.
	UserID string `json:"user_id"`
}

func (o *CreateAccountLiquidationAccountLiquidationRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type CreateAccountLiquidationRequest struct {
	RequestBody *CreateAccountLiquidationAccountLiquidationRequest `request:"mediaType=application/json"`
	AccountID   string                                             `pathParam:"style=simple,explode=false,name=account_id"`
	// A UUID to be used as an idempotency key.  This prevents a duplicate request from being replayed.
	// https://docs.upvest.co/concepts/api_concepts/idempotency
	//
	IdempotencyKey string `header:"style=simple,explode=false,name=idempotency-key"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-http-header
	Signature string `header:"style=simple,explode=false,name=signature"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-input-http-he
	SignatureInput string `header:"style=simple,explode=false,name=signature-input"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
}

func (c CreateAccountLiquidationRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateAccountLiquidationRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateAccountLiquidationRequest) GetRequestBody() *CreateAccountLiquidationAccountLiquidationRequest {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *CreateAccountLiquidationRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *CreateAccountLiquidationRequest) GetIdempotencyKey() string {
	if o == nil {
		return ""
	}
	return o.IdempotencyKey
}

func (o *CreateAccountLiquidationRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *CreateAccountLiquidationRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *CreateAccountLiquidationRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *CreateAccountLiquidationRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// CreateAccountLiquidationCurrency - Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
// * EUR - Euro
type CreateAccountLiquidationCurrency string

const (
	CreateAccountLiquidationCurrencyEur CreateAccountLiquidationCurrency = "EUR"
)

func (e CreateAccountLiquidationCurrency) ToPointer() *CreateAccountLiquidationCurrency {
	return &e
}

func (e *CreateAccountLiquidationCurrency) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EUR":
		*e = CreateAccountLiquidationCurrency(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAccountLiquidationCurrency: %v", v)
	}
}

// CreateAccountLiquidationLiquidationsStatus - Execution status of the Account liquidation order.
// * NEW -
// * PROCESSING -
// * FILLED -
// * CANCELLED -
type CreateAccountLiquidationLiquidationsStatus string

const (
	CreateAccountLiquidationLiquidationsStatusNew        CreateAccountLiquidationLiquidationsStatus = "NEW"
	CreateAccountLiquidationLiquidationsStatusProcessing CreateAccountLiquidationLiquidationsStatus = "PROCESSING"
	CreateAccountLiquidationLiquidationsStatusFilled     CreateAccountLiquidationLiquidationsStatus = "FILLED"
	CreateAccountLiquidationLiquidationsStatusCancelled  CreateAccountLiquidationLiquidationsStatus = "CANCELLED"
)

func (e CreateAccountLiquidationLiquidationsStatus) ToPointer() *CreateAccountLiquidationLiquidationsStatus {
	return &e
}

func (e *CreateAccountLiquidationLiquidationsStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NEW":
		fallthrough
	case "PROCESSING":
		fallthrough
	case "FILLED":
		fallthrough
	case "CANCELLED":
		*e = CreateAccountLiquidationLiquidationsStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAccountLiquidationLiquidationsStatus: %v", v)
	}
}

type AccountLiquidation struct {
	ID string `json:"id"`
	// Side of the order.
	// * SELL -
	Side *string `default:"SELL" json:"side"`
	// Execution status of the Account liquidation order.
	// * NEW -
	// * PROCESSING -
	// * FILLED -
	// * CANCELLED -
	Status CreateAccountLiquidationLiquidationsStatus `json:"status"`
}

func (a AccountLiquidation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountLiquidation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountLiquidation) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AccountLiquidation) GetSide() *string {
	if o == nil {
		return nil
	}
	return o.Side
}

func (o *AccountLiquidation) GetStatus() CreateAccountLiquidationLiquidationsStatus {
	if o == nil {
		return CreateAccountLiquidationLiquidationsStatus("")
	}
	return o.Status
}

// CreateAccountLiquidationStatus - Execution status of the Account liquidation.
// * NEW -
// * PROCESSING -
// * FILLED -
// * CANCELLED -
// * SETTLED -
type CreateAccountLiquidationStatus string

const (
	CreateAccountLiquidationStatusNew        CreateAccountLiquidationStatus = "NEW"
	CreateAccountLiquidationStatusProcessing CreateAccountLiquidationStatus = "PROCESSING"
	CreateAccountLiquidationStatusFilled     CreateAccountLiquidationStatus = "FILLED"
	CreateAccountLiquidationStatusCancelled  CreateAccountLiquidationStatus = "CANCELLED"
	CreateAccountLiquidationStatusSettled    CreateAccountLiquidationStatus = "SETTLED"
)

func (e CreateAccountLiquidationStatus) ToPointer() *CreateAccountLiquidationStatus {
	return &e
}

func (e *CreateAccountLiquidationStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NEW":
		fallthrough
	case "PROCESSING":
		fallthrough
	case "FILLED":
		fallthrough
	case "CANCELLED":
		fallthrough
	case "SETTLED":
		*e = CreateAccountLiquidationStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAccountLiquidationStatus: %v", v)
	}
}

// CreateAccountLiquidationAccountLiquidation - Account liquidation object
type CreateAccountLiquidationAccountLiquidation struct {
	// Account unique identifier.
	AccountID  string `json:"account_id"`
	CashAmount string `json:"cash_amount"`
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	// Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
	// * EUR - Euro
	Currency *CreateAccountLiquidationCurrency `default:"EUR" json:"currency"`
	ID       string                            `json:"id"`
	// Position liquidation orders associated with this account liquidation
	Orders []AccountLiquidation `json:"orders"`
	// Execution status of the Account liquidation.
	// * NEW -
	// * PROCESSING -
	// * FILLED -
	// * CANCELLED -
	// * SETTLED -
	Status CreateAccountLiquidationStatus `json:"status"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time `json:"updated_at"`
	// User unique identifier.
	UserID *string `json:"user_id,omitempty"`
}

func (c CreateAccountLiquidationAccountLiquidation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateAccountLiquidationAccountLiquidation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateAccountLiquidationAccountLiquidation) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *CreateAccountLiquidationAccountLiquidation) GetCashAmount() string {
	if o == nil {
		return ""
	}
	return o.CashAmount
}

func (o *CreateAccountLiquidationAccountLiquidation) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *CreateAccountLiquidationAccountLiquidation) GetCurrency() *CreateAccountLiquidationCurrency {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *CreateAccountLiquidationAccountLiquidation) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateAccountLiquidationAccountLiquidation) GetOrders() []AccountLiquidation {
	if o == nil {
		return []AccountLiquidation{}
	}
	return o.Orders
}

func (o *CreateAccountLiquidationAccountLiquidation) GetStatus() CreateAccountLiquidationStatus {
	if o == nil {
		return CreateAccountLiquidationStatus("")
	}
	return o.Status
}

func (o *CreateAccountLiquidationAccountLiquidation) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *CreateAccountLiquidationAccountLiquidation) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

type CreateAccountLiquidationResponse struct {
	// Account liquidation object
	AccountLiquidation *CreateAccountLiquidationAccountLiquidation
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateAccountLiquidationResponse) GetAccountLiquidation() *CreateAccountLiquidationAccountLiquidation {
	if o == nil {
		return nil
	}
	return o.AccountLiquidation
}

func (o *CreateAccountLiquidationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAccountLiquidationResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *CreateAccountLiquidationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAccountLiquidationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
