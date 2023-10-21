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

// CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCreateRequestCurrency - Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
// * EUR - Euro
type CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCreateRequestCurrency string

const (
	CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCreateRequestCurrencyEur CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCreateRequestCurrency = "EUR"
)

func (e CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCreateRequestCurrency) ToPointer() *CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCreateRequestCurrency {
	return &e
}

func (e *CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCreateRequestCurrency) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EUR":
		*e = CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCreateRequestCurrency(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCreateRequestCurrency: %v", v)
	}
}

type CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCreateRequest struct {
	// Account group unique identifier.
	AccountGroupID string `json:"account_group_id"`
	Amount         string `json:"amount"`
	// Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
	// * EUR - Euro
	Currency *CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCreateRequestCurrency `default:"EUR" json:"currency"`
}

func (c CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCreateRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCreateRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCreateRequest) GetAccountGroupID() string {
	if o == nil {
		return ""
	}
	return o.AccountGroupID
}

func (o *CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCreateRequest) GetAmount() string {
	if o == nil {
		return ""
	}
	return o.Amount
}

func (o *CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCreateRequest) GetCurrency() *CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCreateRequestCurrency {
	if o == nil {
		return nil
	}
	return o.Currency
}

type CreateVirtualCashDecreaseRequest struct {
	RequestBody *CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCreateRequest `request:"mediaType=application/json"`
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

func (c CreateVirtualCashDecreaseRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateVirtualCashDecreaseRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateVirtualCashDecreaseRequest) GetRequestBody() *CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCreateRequest {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *CreateVirtualCashDecreaseRequest) GetIdempotencyKey() string {
	if o == nil {
		return ""
	}
	return o.IdempotencyKey
}

func (o *CreateVirtualCashDecreaseRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *CreateVirtualCashDecreaseRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *CreateVirtualCashDecreaseRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *CreateVirtualCashDecreaseRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCurrency - Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
// * EUR - Euro
type CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCurrency string

const (
	CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCurrencyEur CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCurrency = "EUR"
)

func (e CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCurrency) ToPointer() *CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCurrency {
	return &e
}

func (e *CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCurrency) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EUR":
		*e = CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCurrency(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCurrency: %v", v)
	}
}

// CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseStatus - Status of the virtual cash
// * ISSUED - Virtual cash decrease is created.
// * CONFIRMED - Virtual cash decrease was successfully processed.
// * QUEUED - Virtual cash decrease was queued.
// * CANCELLED - Virtual cash decrease was cancelled.
type CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseStatus string

const (
	CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseStatusIssued    CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseStatus = "ISSUED"
	CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseStatusConfirmed CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseStatus = "CONFIRMED"
	CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseStatusQueued    CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseStatus = "QUEUED"
	CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseStatusCancelled CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseStatus = "CANCELLED"
)

func (e CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseStatus) ToPointer() *CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseStatus {
	return &e
}

func (e *CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ISSUED":
		fallthrough
	case "CONFIRMED":
		fallthrough
	case "QUEUED":
		fallthrough
	case "CANCELLED":
		*e = CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseStatus: %v", v)
	}
}

// CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecrease - Virtual Cash Balances Decrease
type CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecrease struct {
	// Account group unique identifier.
	AccountGroupID string `json:"account_group_id"`
	Amount         string `json:"amount"`
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	// Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
	// * EUR - Euro
	Currency *CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCurrency `default:"EUR" json:"currency"`
	// Virtual cash unique identifier
	ID string `json:"id"`
	// Status of the virtual cash
	// * ISSUED - Virtual cash decrease is created.
	// * CONFIRMED - Virtual cash decrease was successfully processed.
	// * QUEUED - Virtual cash decrease was queued.
	// * CANCELLED - Virtual cash decrease was cancelled.
	Status CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseStatus `json:"status"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time `json:"updated_at"`
}

func (c CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecrease) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecrease) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecrease) GetAccountGroupID() string {
	if o == nil {
		return ""
	}
	return o.AccountGroupID
}

func (o *CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecrease) GetAmount() string {
	if o == nil {
		return ""
	}
	return o.Amount
}

func (o *CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecrease) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecrease) GetCurrency() *CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseCurrency {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecrease) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecrease) GetStatus() CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseStatus {
	if o == nil {
		return CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecreaseStatus("")
	}
	return o.Status
}

func (o *CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecrease) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

type CreateVirtualCashDecreaseResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Virtual Cash Balances Decrease
	VirtualCashBalanceVirtualCashDecrease *CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecrease
}

func (o *CreateVirtualCashDecreaseResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateVirtualCashDecreaseResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreateVirtualCashDecreaseResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateVirtualCashDecreaseResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateVirtualCashDecreaseResponse) GetVirtualCashBalanceVirtualCashDecrease() *CreateVirtualCashDecreaseVirtualCashBalanceVirtualCashDecrease {
	if o == nil {
		return nil
	}
	return o.VirtualCashBalanceVirtualCashDecrease
}