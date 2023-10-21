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

type RetrieveOrderExecutionRequest struct {
	ExecutionID string `pathParam:"style=simple,explode=false,name=execution_id"`
	OrderID     string `pathParam:"style=simple,explode=false,name=order_id"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-http-header
	Signature string `header:"style=simple,explode=false,name=signature"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-input-http-he
	SignatureInput string `header:"style=simple,explode=false,name=signature-input"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
}

func (r RetrieveOrderExecutionRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrieveOrderExecutionRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrieveOrderExecutionRequest) GetExecutionID() string {
	if o == nil {
		return ""
	}
	return o.ExecutionID
}

func (o *RetrieveOrderExecutionRequest) GetOrderID() string {
	if o == nil {
		return ""
	}
	return o.OrderID
}

func (o *RetrieveOrderExecutionRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *RetrieveOrderExecutionRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *RetrieveOrderExecutionRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *RetrieveOrderExecutionRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// RetrieveOrderExecutionOrderExecutionCurrency - Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
// * EUR - Euro
type RetrieveOrderExecutionOrderExecutionCurrency string

const (
	RetrieveOrderExecutionOrderExecutionCurrencyEur RetrieveOrderExecutionOrderExecutionCurrency = "EUR"
)

func (e RetrieveOrderExecutionOrderExecutionCurrency) ToPointer() *RetrieveOrderExecutionOrderExecutionCurrency {
	return &e
}

func (e *RetrieveOrderExecutionOrderExecutionCurrency) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EUR":
		*e = RetrieveOrderExecutionOrderExecutionCurrency(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveOrderExecutionOrderExecutionCurrency: %v", v)
	}
}

// RetrieveOrderExecutionOrderExecutionSide - Side of the execution.
// * BUY -
// * SELL -
type RetrieveOrderExecutionOrderExecutionSide string

const (
	RetrieveOrderExecutionOrderExecutionSideBuy  RetrieveOrderExecutionOrderExecutionSide = "BUY"
	RetrieveOrderExecutionOrderExecutionSideSell RetrieveOrderExecutionOrderExecutionSide = "SELL"
)

func (e RetrieveOrderExecutionOrderExecutionSide) ToPointer() *RetrieveOrderExecutionOrderExecutionSide {
	return &e
}

func (e *RetrieveOrderExecutionOrderExecutionSide) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BUY":
		fallthrough
	case "SELL":
		*e = RetrieveOrderExecutionOrderExecutionSide(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveOrderExecutionOrderExecutionSide: %v", v)
	}
}

// RetrieveOrderExecutionOrderExecutionStatus - Execution status of the Execution.
// * FILLED -
// * SETTLED -
// * CANCELLED -
type RetrieveOrderExecutionOrderExecutionStatus string

const (
	RetrieveOrderExecutionOrderExecutionStatusFilled    RetrieveOrderExecutionOrderExecutionStatus = "FILLED"
	RetrieveOrderExecutionOrderExecutionStatusSettled   RetrieveOrderExecutionOrderExecutionStatus = "SETTLED"
	RetrieveOrderExecutionOrderExecutionStatusCancelled RetrieveOrderExecutionOrderExecutionStatus = "CANCELLED"
)

func (e RetrieveOrderExecutionOrderExecutionStatus) ToPointer() *RetrieveOrderExecutionOrderExecutionStatus {
	return &e
}

func (e *RetrieveOrderExecutionOrderExecutionStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "FILLED":
		fallthrough
	case "SETTLED":
		fallthrough
	case "CANCELLED":
		*e = RetrieveOrderExecutionOrderExecutionStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveOrderExecutionOrderExecutionStatus: %v", v)
	}
}

// RetrieveOrderExecutionOrderExecutionTaxType - Tax type
// * TOTAL -
type RetrieveOrderExecutionOrderExecutionTaxType string

const (
	RetrieveOrderExecutionOrderExecutionTaxTypeTotal RetrieveOrderExecutionOrderExecutionTaxType = "TOTAL"
)

func (e RetrieveOrderExecutionOrderExecutionTaxType) ToPointer() *RetrieveOrderExecutionOrderExecutionTaxType {
	return &e
}

func (e *RetrieveOrderExecutionOrderExecutionTaxType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TOTAL":
		*e = RetrieveOrderExecutionOrderExecutionTaxType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveOrderExecutionOrderExecutionTaxType: %v", v)
	}
}

type RetrieveOrderExecutionOrderExecutionTax struct {
	Amount string `json:"amount"`
	// Tax type
	// * TOTAL -
	Type *RetrieveOrderExecutionOrderExecutionTaxType `default:"TOTAL" json:"type"`
}

func (r RetrieveOrderExecutionOrderExecutionTax) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrieveOrderExecutionOrderExecutionTax) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrieveOrderExecutionOrderExecutionTax) GetAmount() string {
	if o == nil {
		return ""
	}
	return o.Amount
}

func (o *RetrieveOrderExecutionOrderExecutionTax) GetType() *RetrieveOrderExecutionOrderExecutionTaxType {
	if o == nil {
		return nil
	}
	return o.Type
}

// RetrieveOrderExecutionOrderExecution - OK
type RetrieveOrderExecutionOrderExecution struct {
	CashAmount string `json:"cash_amount"`
	// Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
	// * EUR - Euro
	Currency       *RetrieveOrderExecutionOrderExecutionCurrency `default:"EUR" json:"currency"`
	ID             string                                        `json:"id"`
	OrderID        string                                        `json:"order_id"`
	Price          string                                        `json:"price"`
	SettlementDate *string                                       `json:"settlement_date,omitempty"`
	ShareQuantity  string                                        `json:"share_quantity"`
	// Side of the execution.
	// * BUY -
	// * SELL -
	Side RetrieveOrderExecutionOrderExecutionSide `json:"side"`
	// Execution status of the Execution.
	// * FILLED -
	// * SETTLED -
	// * CANCELLED -
	Status          RetrieveOrderExecutionOrderExecutionStatus `json:"status"`
	Taxes           []RetrieveOrderExecutionOrderExecutionTax  `json:"taxes"`
	TransactionTime time.Time                                  `json:"transaction_time"`
}

func (r RetrieveOrderExecutionOrderExecution) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrieveOrderExecutionOrderExecution) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrieveOrderExecutionOrderExecution) GetCashAmount() string {
	if o == nil {
		return ""
	}
	return o.CashAmount
}

func (o *RetrieveOrderExecutionOrderExecution) GetCurrency() *RetrieveOrderExecutionOrderExecutionCurrency {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *RetrieveOrderExecutionOrderExecution) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveOrderExecutionOrderExecution) GetOrderID() string {
	if o == nil {
		return ""
	}
	return o.OrderID
}

func (o *RetrieveOrderExecutionOrderExecution) GetPrice() string {
	if o == nil {
		return ""
	}
	return o.Price
}

func (o *RetrieveOrderExecutionOrderExecution) GetSettlementDate() *string {
	if o == nil {
		return nil
	}
	return o.SettlementDate
}

func (o *RetrieveOrderExecutionOrderExecution) GetShareQuantity() string {
	if o == nil {
		return ""
	}
	return o.ShareQuantity
}

func (o *RetrieveOrderExecutionOrderExecution) GetSide() RetrieveOrderExecutionOrderExecutionSide {
	if o == nil {
		return RetrieveOrderExecutionOrderExecutionSide("")
	}
	return o.Side
}

func (o *RetrieveOrderExecutionOrderExecution) GetStatus() RetrieveOrderExecutionOrderExecutionStatus {
	if o == nil {
		return RetrieveOrderExecutionOrderExecutionStatus("")
	}
	return o.Status
}

func (o *RetrieveOrderExecutionOrderExecution) GetTaxes() []RetrieveOrderExecutionOrderExecutionTax {
	if o == nil {
		return []RetrieveOrderExecutionOrderExecutionTax{}
	}
	return o.Taxes
}

func (o *RetrieveOrderExecutionOrderExecution) GetTransactionTime() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.TransactionTime
}

type RetrieveOrderExecutionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// OK
	OrderExecution *RetrieveOrderExecutionOrderExecution
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RetrieveOrderExecutionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RetrieveOrderExecutionResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *RetrieveOrderExecutionResponse) GetOrderExecution() *RetrieveOrderExecutionOrderExecution {
	if o == nil {
		return nil
	}
	return o.OrderExecution
}

func (o *RetrieveOrderExecutionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RetrieveOrderExecutionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}