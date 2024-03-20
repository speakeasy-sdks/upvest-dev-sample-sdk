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

// CreatePortfoliosOrderCurrency - Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
// * EUR - Euro
type CreatePortfoliosOrderCurrency string

const (
	CreatePortfoliosOrderCurrencyEur CreatePortfoliosOrderCurrency = "EUR"
)

func (e CreatePortfoliosOrderCurrency) ToPointer() *CreatePortfoliosOrderCurrency {
	return &e
}

func (e *CreatePortfoliosOrderCurrency) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EUR":
		*e = CreatePortfoliosOrderCurrency(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreatePortfoliosOrderCurrency: %v", v)
	}
}

// CreatePortfoliosOrderSide - Side of the portfolio order.
// * BUY -
// * SELL -
type CreatePortfoliosOrderSide string

const (
	CreatePortfoliosOrderSideBuy  CreatePortfoliosOrderSide = "BUY"
	CreatePortfoliosOrderSideSell CreatePortfoliosOrderSide = "SELL"
)

func (e CreatePortfoliosOrderSide) ToPointer() *CreatePortfoliosOrderSide {
	return &e
}

func (e *CreatePortfoliosOrderSide) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BUY":
		fallthrough
	case "SELL":
		*e = CreatePortfoliosOrderSide(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreatePortfoliosOrderSide: %v", v)
	}
}

type CreatePortfoliosOrderPortfoliosOrderPlaceRequest struct {
	// Account unique identifier.
	AccountID  string `json:"account_id"`
	CashAmount string `json:"cash_amount"`
	// Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
	// * EUR - Euro
	Currency *CreatePortfoliosOrderCurrency `default:"EUR" json:"currency"`
	// Cash amount is post-tax value
	PostTax *bool `default:"false" json:"post_tax"`
	// Side of the portfolio order.
	// * BUY -
	// * SELL -
	Side CreatePortfoliosOrderSide `json:"side"`
	// User unique identifier.
	UserID string `json:"user_id"`
}

func (c CreatePortfoliosOrderPortfoliosOrderPlaceRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreatePortfoliosOrderPortfoliosOrderPlaceRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreatePortfoliosOrderPortfoliosOrderPlaceRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *CreatePortfoliosOrderPortfoliosOrderPlaceRequest) GetCashAmount() string {
	if o == nil {
		return ""
	}
	return o.CashAmount
}

func (o *CreatePortfoliosOrderPortfoliosOrderPlaceRequest) GetCurrency() *CreatePortfoliosOrderCurrency {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *CreatePortfoliosOrderPortfoliosOrderPlaceRequest) GetPostTax() *bool {
	if o == nil {
		return nil
	}
	return o.PostTax
}

func (o *CreatePortfoliosOrderPortfoliosOrderPlaceRequest) GetSide() CreatePortfoliosOrderSide {
	if o == nil {
		return CreatePortfoliosOrderSide("")
	}
	return o.Side
}

func (o *CreatePortfoliosOrderPortfoliosOrderPlaceRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type CreatePortfoliosOrderRequest struct {
	RequestBody *CreatePortfoliosOrderPortfoliosOrderPlaceRequest `request:"mediaType=application/json"`
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

func (c CreatePortfoliosOrderRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreatePortfoliosOrderRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreatePortfoliosOrderRequest) GetRequestBody() *CreatePortfoliosOrderPortfoliosOrderPlaceRequest {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *CreatePortfoliosOrderRequest) GetIdempotencyKey() string {
	if o == nil {
		return ""
	}
	return o.IdempotencyKey
}

func (o *CreatePortfoliosOrderRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *CreatePortfoliosOrderRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *CreatePortfoliosOrderRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *CreatePortfoliosOrderRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// CreatePortfoliosOrderPortfoliosCurrency - Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
// * EUR - Euro
type CreatePortfoliosOrderPortfoliosCurrency string

const (
	CreatePortfoliosOrderPortfoliosCurrencyEur CreatePortfoliosOrderPortfoliosCurrency = "EUR"
)

func (e CreatePortfoliosOrderPortfoliosCurrency) ToPointer() *CreatePortfoliosOrderPortfoliosCurrency {
	return &e
}

func (e *CreatePortfoliosOrderPortfoliosCurrency) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EUR":
		*e = CreatePortfoliosOrderPortfoliosCurrency(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreatePortfoliosOrderPortfoliosCurrency: %v", v)
	}
}

// CreatePortfoliosOrderPortfoliosSide - Side of the portfolio order.
// * BUY -
// * SELL -
type CreatePortfoliosOrderPortfoliosSide string

const (
	CreatePortfoliosOrderPortfoliosSideBuy  CreatePortfoliosOrderPortfoliosSide = "BUY"
	CreatePortfoliosOrderPortfoliosSideSell CreatePortfoliosOrderPortfoliosSide = "SELL"
)

func (e CreatePortfoliosOrderPortfoliosSide) ToPointer() *CreatePortfoliosOrderPortfoliosSide {
	return &e
}

func (e *CreatePortfoliosOrderPortfoliosSide) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BUY":
		fallthrough
	case "SELL":
		*e = CreatePortfoliosOrderPortfoliosSide(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreatePortfoliosOrderPortfoliosSide: %v", v)
	}
}

// CreatePortfoliosOrderPortfoliosStatus - The execution status of the order.
// * NEW -
// * PROCESSING -
// * FILLED -
// * CANCELLED -
type CreatePortfoliosOrderPortfoliosStatus string

const (
	CreatePortfoliosOrderPortfoliosStatusNew        CreatePortfoliosOrderPortfoliosStatus = "NEW"
	CreatePortfoliosOrderPortfoliosStatusProcessing CreatePortfoliosOrderPortfoliosStatus = "PROCESSING"
	CreatePortfoliosOrderPortfoliosStatusFilled     CreatePortfoliosOrderPortfoliosStatus = "FILLED"
	CreatePortfoliosOrderPortfoliosStatusCancelled  CreatePortfoliosOrderPortfoliosStatus = "CANCELLED"
)

func (e CreatePortfoliosOrderPortfoliosStatus) ToPointer() *CreatePortfoliosOrderPortfoliosStatus {
	return &e
}

func (e *CreatePortfoliosOrderPortfoliosStatus) UnmarshalJSON(data []byte) error {
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
		*e = CreatePortfoliosOrderPortfoliosStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreatePortfoliosOrderPortfoliosStatus: %v", v)
	}
}

type PortfoliosOrder struct {
	ID string `json:"id"`
	// Side of the portfolio order.
	// * BUY -
	// * SELL -
	Side CreatePortfoliosOrderPortfoliosSide `json:"side"`
	// The execution status of the order.
	// * NEW -
	// * PROCESSING -
	// * FILLED -
	// * CANCELLED -
	Status CreatePortfoliosOrderPortfoliosStatus `json:"status"`
}

func (o *PortfoliosOrder) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PortfoliosOrder) GetSide() CreatePortfoliosOrderPortfoliosSide {
	if o == nil {
		return CreatePortfoliosOrderPortfoliosSide("")
	}
	return o.Side
}

func (o *PortfoliosOrder) GetStatus() CreatePortfoliosOrderPortfoliosStatus {
	if o == nil {
		return CreatePortfoliosOrderPortfoliosStatus("")
	}
	return o.Status
}

// CreatePortfoliosOrderStatus - Execution status of the Portfolio Order.
// * NEW -
// * PROCESSING -
// * FILLED -
// * SETTLED -
// * CANCELLED -
type CreatePortfoliosOrderStatus string

const (
	CreatePortfoliosOrderStatusNew        CreatePortfoliosOrderStatus = "NEW"
	CreatePortfoliosOrderStatusProcessing CreatePortfoliosOrderStatus = "PROCESSING"
	CreatePortfoliosOrderStatusFilled     CreatePortfoliosOrderStatus = "FILLED"
	CreatePortfoliosOrderStatusSettled    CreatePortfoliosOrderStatus = "SETTLED"
	CreatePortfoliosOrderStatusCancelled  CreatePortfoliosOrderStatus = "CANCELLED"
)

func (e CreatePortfoliosOrderStatus) ToPointer() *CreatePortfoliosOrderStatus {
	return &e
}

func (e *CreatePortfoliosOrderStatus) UnmarshalJSON(data []byte) error {
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
	case "SETTLED":
		fallthrough
	case "CANCELLED":
		*e = CreatePortfoliosOrderStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreatePortfoliosOrderStatus: %v", v)
	}
}

// CreatePortfoliosOrderType - Type of the Portfolio Order.
// * BUY -
// * SELL -
// * REBALANCING -
type CreatePortfoliosOrderType string

const (
	CreatePortfoliosOrderTypeBuy         CreatePortfoliosOrderType = "BUY"
	CreatePortfoliosOrderTypeSell        CreatePortfoliosOrderType = "SELL"
	CreatePortfoliosOrderTypeRebalancing CreatePortfoliosOrderType = "REBALANCING"
)

func (e CreatePortfoliosOrderType) ToPointer() *CreatePortfoliosOrderType {
	return &e
}

func (e *CreatePortfoliosOrderType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BUY":
		fallthrough
	case "SELL":
		fallthrough
	case "REBALANCING":
		*e = CreatePortfoliosOrderType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreatePortfoliosOrderType: %v", v)
	}
}

// CreatePortfoliosOrderPortfoliosOrder - Portfolios order
type CreatePortfoliosOrderPortfoliosOrder struct {
	// Account unique identifier.
	AccountID    string  `json:"account_id"`
	AllocationID *string `json:"allocation_id,omitempty"`
	CashAmount   string  `json:"cash_amount"`
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	// Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
	// * EUR - Euro
	Currency *CreatePortfoliosOrderPortfoliosCurrency `default:"EUR" json:"currency"`
	ID       string                                   `json:"id"`
	// Orders associated with this portfolio order
	Orders []PortfoliosOrder `json:"orders"`
	// Cash amount is post-tax value
	PostTax *bool `default:"false" json:"post_tax"`
	// Execution status of the Portfolio Order.
	// * NEW -
	// * PROCESSING -
	// * FILLED -
	// * SETTLED -
	// * CANCELLED -
	Status CreatePortfoliosOrderStatus `json:"status"`
	// Type of the Portfolio Order.
	// * BUY -
	// * SELL -
	// * REBALANCING -
	Type *CreatePortfoliosOrderType `json:"type,omitempty"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time `json:"updated_at"`
	// User unique identifier.
	UserID *string `json:"user_id,omitempty"`
}

func (c CreatePortfoliosOrderPortfoliosOrder) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreatePortfoliosOrderPortfoliosOrder) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreatePortfoliosOrderPortfoliosOrder) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *CreatePortfoliosOrderPortfoliosOrder) GetAllocationID() *string {
	if o == nil {
		return nil
	}
	return o.AllocationID
}

func (o *CreatePortfoliosOrderPortfoliosOrder) GetCashAmount() string {
	if o == nil {
		return ""
	}
	return o.CashAmount
}

func (o *CreatePortfoliosOrderPortfoliosOrder) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *CreatePortfoliosOrderPortfoliosOrder) GetCurrency() *CreatePortfoliosOrderPortfoliosCurrency {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *CreatePortfoliosOrderPortfoliosOrder) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreatePortfoliosOrderPortfoliosOrder) GetOrders() []PortfoliosOrder {
	if o == nil {
		return []PortfoliosOrder{}
	}
	return o.Orders
}

func (o *CreatePortfoliosOrderPortfoliosOrder) GetPostTax() *bool {
	if o == nil {
		return nil
	}
	return o.PostTax
}

func (o *CreatePortfoliosOrderPortfoliosOrder) GetStatus() CreatePortfoliosOrderStatus {
	if o == nil {
		return CreatePortfoliosOrderStatus("")
	}
	return o.Status
}

func (o *CreatePortfoliosOrderPortfoliosOrder) GetType() *CreatePortfoliosOrderType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *CreatePortfoliosOrderPortfoliosOrder) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *CreatePortfoliosOrderPortfoliosOrder) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

type CreatePortfoliosOrderResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// Portfolios order
	PortfoliosOrder *CreatePortfoliosOrderPortfoliosOrder
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreatePortfoliosOrderResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreatePortfoliosOrderResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *CreatePortfoliosOrderResponse) GetPortfoliosOrder() *CreatePortfoliosOrderPortfoliosOrder {
	if o == nil {
		return nil
	}
	return o.PortfoliosOrder
}

func (o *CreatePortfoliosOrderResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreatePortfoliosOrderResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
