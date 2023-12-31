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

// CreateVirtualCashIncreaseCurrency - Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
// * EUR - Euro
type CreateVirtualCashIncreaseCurrency string

const (
	CreateVirtualCashIncreaseCurrencyEur CreateVirtualCashIncreaseCurrency = "EUR"
)

func (e CreateVirtualCashIncreaseCurrency) ToPointer() *CreateVirtualCashIncreaseCurrency {
	return &e
}

func (e *CreateVirtualCashIncreaseCurrency) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EUR":
		*e = CreateVirtualCashIncreaseCurrency(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateVirtualCashIncreaseCurrency: %v", v)
	}
}

type CreateVirtualCashIncreaseVirtualCashBalanceVirtualCashIncreaseCreateRequest struct {
	// Account group unique identifier.
	AccountGroupID string `json:"account_group_id"`
	Amount         string `json:"amount"`
	// Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
	// * EUR - Euro
	Currency *CreateVirtualCashIncreaseCurrency `default:"EUR" json:"currency"`
}

func (c CreateVirtualCashIncreaseVirtualCashBalanceVirtualCashIncreaseCreateRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateVirtualCashIncreaseVirtualCashBalanceVirtualCashIncreaseCreateRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateVirtualCashIncreaseVirtualCashBalanceVirtualCashIncreaseCreateRequest) GetAccountGroupID() string {
	if o == nil {
		return ""
	}
	return o.AccountGroupID
}

func (o *CreateVirtualCashIncreaseVirtualCashBalanceVirtualCashIncreaseCreateRequest) GetAmount() string {
	if o == nil {
		return ""
	}
	return o.Amount
}

func (o *CreateVirtualCashIncreaseVirtualCashBalanceVirtualCashIncreaseCreateRequest) GetCurrency() *CreateVirtualCashIncreaseCurrency {
	if o == nil {
		return nil
	}
	return o.Currency
}

type CreateVirtualCashIncreaseRequest struct {
	RequestBody *CreateVirtualCashIncreaseVirtualCashBalanceVirtualCashIncreaseCreateRequest `request:"mediaType=application/json"`
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

func (c CreateVirtualCashIncreaseRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateVirtualCashIncreaseRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateVirtualCashIncreaseRequest) GetRequestBody() *CreateVirtualCashIncreaseVirtualCashBalanceVirtualCashIncreaseCreateRequest {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *CreateVirtualCashIncreaseRequest) GetIdempotencyKey() string {
	if o == nil {
		return ""
	}
	return o.IdempotencyKey
}

func (o *CreateVirtualCashIncreaseRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *CreateVirtualCashIncreaseRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *CreateVirtualCashIncreaseRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *CreateVirtualCashIncreaseRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// CreateVirtualCashIncreaseVirtualCashBalancesCurrency - Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
// * EUR - Euro
type CreateVirtualCashIncreaseVirtualCashBalancesCurrency string

const (
	CreateVirtualCashIncreaseVirtualCashBalancesCurrencyEur CreateVirtualCashIncreaseVirtualCashBalancesCurrency = "EUR"
)

func (e CreateVirtualCashIncreaseVirtualCashBalancesCurrency) ToPointer() *CreateVirtualCashIncreaseVirtualCashBalancesCurrency {
	return &e
}

func (e *CreateVirtualCashIncreaseVirtualCashBalancesCurrency) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EUR":
		*e = CreateVirtualCashIncreaseVirtualCashBalancesCurrency(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateVirtualCashIncreaseVirtualCashBalancesCurrency: %v", v)
	}
}

// CreateVirtualCashIncreaseStatus - Status of the virtual cash
// * ISSUED - Virtual cash increase is created.
// * CONFIRMED - Virtual cash increase was successfully processed.
type CreateVirtualCashIncreaseStatus string

const (
	CreateVirtualCashIncreaseStatusIssued    CreateVirtualCashIncreaseStatus = "ISSUED"
	CreateVirtualCashIncreaseStatusConfirmed CreateVirtualCashIncreaseStatus = "CONFIRMED"
)

func (e CreateVirtualCashIncreaseStatus) ToPointer() *CreateVirtualCashIncreaseStatus {
	return &e
}

func (e *CreateVirtualCashIncreaseStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ISSUED":
		fallthrough
	case "CONFIRMED":
		*e = CreateVirtualCashIncreaseStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateVirtualCashIncreaseStatus: %v", v)
	}
}

// CreateVirtualCashIncreaseVirtualCashBalanceVirtualCashIncrease - Virtual Cash Balances Increase
type CreateVirtualCashIncreaseVirtualCashBalanceVirtualCashIncrease struct {
	// Account group unique identifier.
	AccountGroupID string `json:"account_group_id"`
	Amount         string `json:"amount"`
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	// Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
	// * EUR - Euro
	Currency *CreateVirtualCashIncreaseVirtualCashBalancesCurrency `default:"EUR" json:"currency"`
	// Virtual cash unique identifier
	ID string `json:"id"`
	// Status of the virtual cash
	// * ISSUED - Virtual cash increase is created.
	// * CONFIRMED - Virtual cash increase was successfully processed.
	Status CreateVirtualCashIncreaseStatus `json:"status"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time `json:"updated_at"`
}

func (c CreateVirtualCashIncreaseVirtualCashBalanceVirtualCashIncrease) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateVirtualCashIncreaseVirtualCashBalanceVirtualCashIncrease) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateVirtualCashIncreaseVirtualCashBalanceVirtualCashIncrease) GetAccountGroupID() string {
	if o == nil {
		return ""
	}
	return o.AccountGroupID
}

func (o *CreateVirtualCashIncreaseVirtualCashBalanceVirtualCashIncrease) GetAmount() string {
	if o == nil {
		return ""
	}
	return o.Amount
}

func (o *CreateVirtualCashIncreaseVirtualCashBalanceVirtualCashIncrease) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *CreateVirtualCashIncreaseVirtualCashBalanceVirtualCashIncrease) GetCurrency() *CreateVirtualCashIncreaseVirtualCashBalancesCurrency {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *CreateVirtualCashIncreaseVirtualCashBalanceVirtualCashIncrease) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateVirtualCashIncreaseVirtualCashBalanceVirtualCashIncrease) GetStatus() CreateVirtualCashIncreaseStatus {
	if o == nil {
		return CreateVirtualCashIncreaseStatus("")
	}
	return o.Status
}

func (o *CreateVirtualCashIncreaseVirtualCashBalanceVirtualCashIncrease) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

type CreateVirtualCashIncreaseResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Virtual Cash Balances Increase
	VirtualCashBalanceVirtualCashIncrease *CreateVirtualCashIncreaseVirtualCashBalanceVirtualCashIncrease
}

func (o *CreateVirtualCashIncreaseResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateVirtualCashIncreaseResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *CreateVirtualCashIncreaseResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateVirtualCashIncreaseResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateVirtualCashIncreaseResponse) GetVirtualCashBalanceVirtualCashIncrease() *CreateVirtualCashIncreaseVirtualCashBalanceVirtualCashIncrease {
	if o == nil {
		return nil
	}
	return o.VirtualCashBalanceVirtualCashIncrease
}
