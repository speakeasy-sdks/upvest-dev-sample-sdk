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

type CreateCashWithdrawalSecurity struct {
	OauthClientCredentials string `security:"scheme,type=oauth2,name=Authorization"`
}

func (o *CreateCashWithdrawalSecurity) GetOauthClientCredentials() string {
	if o == nil {
		return ""
	}
	return o.OauthClientCredentials
}

type CreateCashWithdrawalPaymentsWithdrawalCreateRequest struct {
	// Account group unique identifier.
	AccountGroupID string `json:"account_group_id"`
	Amount         string `json:"amount"`
	// Reference account unique identifier.
	ReferenceAccountID string `json:"reference_account_id"`
	// Payment reference the end user will see in their bank statement for the corresponding credit transfer booking (“Verwendungszweck”)
	RemittanceInformation *string `json:"remittance_information,omitempty"`
	// User unique identifier.
	UserID string `json:"user_id"`
}

func (o *CreateCashWithdrawalPaymentsWithdrawalCreateRequest) GetAccountGroupID() string {
	if o == nil {
		return ""
	}
	return o.AccountGroupID
}

func (o *CreateCashWithdrawalPaymentsWithdrawalCreateRequest) GetAmount() string {
	if o == nil {
		return ""
	}
	return o.Amount
}

func (o *CreateCashWithdrawalPaymentsWithdrawalCreateRequest) GetReferenceAccountID() string {
	if o == nil {
		return ""
	}
	return o.ReferenceAccountID
}

func (o *CreateCashWithdrawalPaymentsWithdrawalCreateRequest) GetRemittanceInformation() *string {
	if o == nil {
		return nil
	}
	return o.RemittanceInformation
}

func (o *CreateCashWithdrawalPaymentsWithdrawalCreateRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type CreateCashWithdrawalRequest struct {
	RequestBody *CreateCashWithdrawalPaymentsWithdrawalCreateRequest `request:"mediaType=application/json"`
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

func (c CreateCashWithdrawalRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateCashWithdrawalRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateCashWithdrawalRequest) GetRequestBody() *CreateCashWithdrawalPaymentsWithdrawalCreateRequest {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *CreateCashWithdrawalRequest) GetIdempotencyKey() string {
	if o == nil {
		return ""
	}
	return o.IdempotencyKey
}

func (o *CreateCashWithdrawalRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *CreateCashWithdrawalRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *CreateCashWithdrawalRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *CreateCashWithdrawalRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// CreateCashWithdrawalCurrency - Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
// * EUR - Euro
type CreateCashWithdrawalCurrency string

const (
	CreateCashWithdrawalCurrencyEur CreateCashWithdrawalCurrency = "EUR"
)

func (e CreateCashWithdrawalCurrency) ToPointer() *CreateCashWithdrawalCurrency {
	return &e
}

func (e *CreateCashWithdrawalCurrency) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EUR":
		*e = CreateCashWithdrawalCurrency(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateCashWithdrawalCurrency: %v", v)
	}
}

// CreateCashWithdrawalStatus - Status of the withdrawal
// * NEW - Withdrawal is created but not started processing.
// * PROCESSING - Withdrawal is in processing.
// * CONFIRMED - Withdrawal was successfully processed.
// * CANCELLED - Withdrawal was cancelled.
type CreateCashWithdrawalStatus string

const (
	CreateCashWithdrawalStatusNew        CreateCashWithdrawalStatus = "NEW"
	CreateCashWithdrawalStatusProcessing CreateCashWithdrawalStatus = "PROCESSING"
	CreateCashWithdrawalStatusConfirmed  CreateCashWithdrawalStatus = "CONFIRMED"
	CreateCashWithdrawalStatusCancelled  CreateCashWithdrawalStatus = "CANCELLED"
)

func (e CreateCashWithdrawalStatus) ToPointer() *CreateCashWithdrawalStatus {
	return &e
}

func (e *CreateCashWithdrawalStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NEW":
		fallthrough
	case "PROCESSING":
		fallthrough
	case "CONFIRMED":
		fallthrough
	case "CANCELLED":
		*e = CreateCashWithdrawalStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateCashWithdrawalStatus: %v", v)
	}
}

// CreateCashWithdrawalPaymentsWithdrawal - Withdrawal
type CreateCashWithdrawalPaymentsWithdrawal struct {
	// Account group unique identifier.
	AccountGroupID string `json:"account_group_id"`
	Amount         string `json:"amount"`
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	// Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
	// * EUR - Euro
	Currency *CreateCashWithdrawalCurrency `default:"EUR" json:"currency"`
	// Cash withdrawal unique identifier
	ID string `json:"id"`
	// Reference account unique identifier.
	ReferenceAccountID string `json:"reference_account_id"`
	// Payment reference the end user will see in their bank statement for the corresponding credit transfer booking (“Verwendungszweck”)
	RemittanceInformation string `json:"remittance_information"`
	// Status of the withdrawal
	// * NEW - Withdrawal is created but not started processing.
	// * PROCESSING - Withdrawal is in processing.
	// * CONFIRMED - Withdrawal was successfully processed.
	// * CANCELLED - Withdrawal was cancelled.
	Status *CreateCashWithdrawalStatus `json:"status,omitempty"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time `json:"updated_at"`
	// User unique identifier.
	UserID string `json:"user_id"`
}

func (c CreateCashWithdrawalPaymentsWithdrawal) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateCashWithdrawalPaymentsWithdrawal) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateCashWithdrawalPaymentsWithdrawal) GetAccountGroupID() string {
	if o == nil {
		return ""
	}
	return o.AccountGroupID
}

func (o *CreateCashWithdrawalPaymentsWithdrawal) GetAmount() string {
	if o == nil {
		return ""
	}
	return o.Amount
}

func (o *CreateCashWithdrawalPaymentsWithdrawal) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *CreateCashWithdrawalPaymentsWithdrawal) GetCurrency() *CreateCashWithdrawalCurrency {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *CreateCashWithdrawalPaymentsWithdrawal) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateCashWithdrawalPaymentsWithdrawal) GetReferenceAccountID() string {
	if o == nil {
		return ""
	}
	return o.ReferenceAccountID
}

func (o *CreateCashWithdrawalPaymentsWithdrawal) GetRemittanceInformation() string {
	if o == nil {
		return ""
	}
	return o.RemittanceInformation
}

func (o *CreateCashWithdrawalPaymentsWithdrawal) GetStatus() *CreateCashWithdrawalStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *CreateCashWithdrawalPaymentsWithdrawal) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *CreateCashWithdrawalPaymentsWithdrawal) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type CreateCashWithdrawalResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// Withdrawal
	PaymentsWithdrawal *CreateCashWithdrawalPaymentsWithdrawal
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateCashWithdrawalResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateCashWithdrawalResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *CreateCashWithdrawalResponse) GetPaymentsWithdrawal() *CreateCashWithdrawalPaymentsWithdrawal {
	if o == nil {
		return nil
	}
	return o.PaymentsWithdrawal
}

func (o *CreateCashWithdrawalResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateCashWithdrawalResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
