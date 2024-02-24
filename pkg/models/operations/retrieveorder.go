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

type RetrieveOrderSecurity struct {
	OauthClientCredentials string `security:"scheme,type=oauth2,name=Authorization"`
}

func (o *RetrieveOrderSecurity) GetOauthClientCredentials() string {
	if o == nil {
		return ""
	}
	return o.OauthClientCredentials
}

type RetrieveOrderRequest struct {
	OrderID string `pathParam:"style=simple,explode=false,name=order_id"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-http-header
	Signature string `header:"style=simple,explode=false,name=signature"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-input-http-he
	SignatureInput string `header:"style=simple,explode=false,name=signature-input"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
}

func (r RetrieveOrderRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrieveOrderRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrieveOrderRequest) GetOrderID() string {
	if o == nil {
		return ""
	}
	return o.OrderID
}

func (o *RetrieveOrderRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *RetrieveOrderRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *RetrieveOrderRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *RetrieveOrderRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// RetrieveOrderCancellationReason - Reason for Order cancellation. The field is present in case the Order has a status of CANCELLED.
// * CANCELLED_BY_CLIENT -
// * CANCELLED_BY_UPVEST_OPERATIONS -
// * CANCELLED_BY_TRADING_PARTNER -
// * CANCELLED_BY_UPVEST_PLATFORM -
type RetrieveOrderCancellationReason string

const (
	RetrieveOrderCancellationReasonCancelledByClient           RetrieveOrderCancellationReason = "CANCELLED_BY_CLIENT"
	RetrieveOrderCancellationReasonCancelledByUpvestOperations RetrieveOrderCancellationReason = "CANCELLED_BY_UPVEST_OPERATIONS"
	RetrieveOrderCancellationReasonCancelledByTradingPartner   RetrieveOrderCancellationReason = "CANCELLED_BY_TRADING_PARTNER"
	RetrieveOrderCancellationReasonCancelledByUpvestPlatform   RetrieveOrderCancellationReason = "CANCELLED_BY_UPVEST_PLATFORM"
)

func (e RetrieveOrderCancellationReason) ToPointer() *RetrieveOrderCancellationReason {
	return &e
}

func (e *RetrieveOrderCancellationReason) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CANCELLED_BY_CLIENT":
		fallthrough
	case "CANCELLED_BY_UPVEST_OPERATIONS":
		fallthrough
	case "CANCELLED_BY_TRADING_PARTNER":
		fallthrough
	case "CANCELLED_BY_UPVEST_PLATFORM":
		*e = RetrieveOrderCancellationReason(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveOrderCancellationReason: %v", v)
	}
}

// RetrieveOrderCurrency - The currency for the order.
type RetrieveOrderCurrency string

const (
	RetrieveOrderCurrencyEur RetrieveOrderCurrency = "EUR"
)

func (e RetrieveOrderCurrency) ToPointer() *RetrieveOrderCurrency {
	return &e
}

func (e *RetrieveOrderCurrency) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EUR":
		*e = RetrieveOrderCurrency(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveOrderCurrency: %v", v)
	}
}

// RetrieveOrderExecutionFlow - Execution flow that the order processing goes through. If no value is specified, the default value is assumed - `STRAIGHT_THROUGH`.
// * STRAIGHT_THROUGH -
// * BLOCK -
type RetrieveOrderExecutionFlow string

const (
	RetrieveOrderExecutionFlowStraightThrough RetrieveOrderExecutionFlow = "STRAIGHT_THROUGH"
	RetrieveOrderExecutionFlowBlock           RetrieveOrderExecutionFlow = "BLOCK"
)

func (e RetrieveOrderExecutionFlow) ToPointer() *RetrieveOrderExecutionFlow {
	return &e
}

func (e *RetrieveOrderExecutionFlow) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "STRAIGHT_THROUGH":
		fallthrough
	case "BLOCK":
		*e = RetrieveOrderExecutionFlow(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveOrderExecutionFlow: %v", v)
	}
}

// RetrieveOrderOrdersCurrency - Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
// * EUR - Euro
type RetrieveOrderOrdersCurrency string

const (
	RetrieveOrderOrdersCurrencyEur RetrieveOrderOrdersCurrency = "EUR"
)

func (e RetrieveOrderOrdersCurrency) ToPointer() *RetrieveOrderOrdersCurrency {
	return &e
}

func (e *RetrieveOrderOrdersCurrency) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EUR":
		*e = RetrieveOrderOrdersCurrency(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveOrderOrdersCurrency: %v", v)
	}
}

// RetrieveOrderOrdersSide - Side of the execution.
// * BUY -
// * SELL -
type RetrieveOrderOrdersSide string

const (
	RetrieveOrderOrdersSideBuy  RetrieveOrderOrdersSide = "BUY"
	RetrieveOrderOrdersSideSell RetrieveOrderOrdersSide = "SELL"
)

func (e RetrieveOrderOrdersSide) ToPointer() *RetrieveOrderOrdersSide {
	return &e
}

func (e *RetrieveOrderOrdersSide) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BUY":
		fallthrough
	case "SELL":
		*e = RetrieveOrderOrdersSide(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveOrderOrdersSide: %v", v)
	}
}

// RetrieveOrderOrdersStatus - Execution status of the Execution.
// * FILLED -
// * SETTLED -
// * CANCELLED -
type RetrieveOrderOrdersStatus string

const (
	RetrieveOrderOrdersStatusFilled    RetrieveOrderOrdersStatus = "FILLED"
	RetrieveOrderOrdersStatusSettled   RetrieveOrderOrdersStatus = "SETTLED"
	RetrieveOrderOrdersStatusCancelled RetrieveOrderOrdersStatus = "CANCELLED"
)

func (e RetrieveOrderOrdersStatus) ToPointer() *RetrieveOrderOrdersStatus {
	return &e
}

func (e *RetrieveOrderOrdersStatus) UnmarshalJSON(data []byte) error {
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
		*e = RetrieveOrderOrdersStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveOrderOrdersStatus: %v", v)
	}
}

// RetrieveOrderType - Tax type
// * TOTAL -
type RetrieveOrderType string

const (
	RetrieveOrderTypeTotal RetrieveOrderType = "TOTAL"
)

func (e RetrieveOrderType) ToPointer() *RetrieveOrderType {
	return &e
}

func (e *RetrieveOrderType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TOTAL":
		*e = RetrieveOrderType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveOrderType: %v", v)
	}
}

type RetrieveOrderTax struct {
	Amount string `json:"amount"`
	// Tax type
	// * TOTAL -
	Type *RetrieveOrderType `default:"TOTAL" json:"type"`
}

func (r RetrieveOrderTax) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrieveOrderTax) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrieveOrderTax) GetAmount() string {
	if o == nil {
		return ""
	}
	return o.Amount
}

func (o *RetrieveOrderTax) GetType() *RetrieveOrderType {
	if o == nil {
		return nil
	}
	return o.Type
}

type RetrieveOrderOrderExecution struct {
	CashAmount string `json:"cash_amount"`
	// Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
	// * EUR - Euro
	Currency       *RetrieveOrderOrdersCurrency `default:"EUR" json:"currency"`
	ID             string                       `json:"id"`
	OrderID        string                       `json:"order_id"`
	Price          string                       `json:"price"`
	SettlementDate *string                      `json:"settlement_date,omitempty"`
	ShareQuantity  string                       `json:"share_quantity"`
	// Side of the execution.
	// * BUY -
	// * SELL -
	Side RetrieveOrderOrdersSide `json:"side"`
	// Execution status of the Execution.
	// * FILLED -
	// * SETTLED -
	// * CANCELLED -
	Status          RetrieveOrderOrdersStatus `json:"status"`
	Taxes           []RetrieveOrderTax        `json:"taxes"`
	TransactionTime time.Time                 `json:"transaction_time"`
}

func (r RetrieveOrderOrderExecution) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrieveOrderOrderExecution) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrieveOrderOrderExecution) GetCashAmount() string {
	if o == nil {
		return ""
	}
	return o.CashAmount
}

func (o *RetrieveOrderOrderExecution) GetCurrency() *RetrieveOrderOrdersCurrency {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *RetrieveOrderOrderExecution) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveOrderOrderExecution) GetOrderID() string {
	if o == nil {
		return ""
	}
	return o.OrderID
}

func (o *RetrieveOrderOrderExecution) GetPrice() string {
	if o == nil {
		return ""
	}
	return o.Price
}

func (o *RetrieveOrderOrderExecution) GetSettlementDate() *string {
	if o == nil {
		return nil
	}
	return o.SettlementDate
}

func (o *RetrieveOrderOrderExecution) GetShareQuantity() string {
	if o == nil {
		return ""
	}
	return o.ShareQuantity
}

func (o *RetrieveOrderOrderExecution) GetSide() RetrieveOrderOrdersSide {
	if o == nil {
		return RetrieveOrderOrdersSide("")
	}
	return o.Side
}

func (o *RetrieveOrderOrderExecution) GetStatus() RetrieveOrderOrdersStatus {
	if o == nil {
		return RetrieveOrderOrdersStatus("")
	}
	return o.Status
}

func (o *RetrieveOrderOrderExecution) GetTaxes() []RetrieveOrderTax {
	if o == nil {
		return []RetrieveOrderTax{}
	}
	return o.Taxes
}

func (o *RetrieveOrderOrderExecution) GetTransactionTime() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.TransactionTime
}

// RetrieveOrderInitiationFlow - Initiation flow used during order creation, i.e. what triggered the order.
// * API -
// * PORTFOLIO -
// * CASH_DIVIDEND_REINVESTMENT -
// * PORTFOLIO_REBALANCING -
// * SELL_TO_COVER_FEES -
// * SELL_TO_COVER_TAXES -
// * ACCOUNT_LIQUIDATION -
// * UPVEST_OPERATIONS -
type RetrieveOrderInitiationFlow string

const (
	RetrieveOrderInitiationFlowAPI                      RetrieveOrderInitiationFlow = "API"
	RetrieveOrderInitiationFlowPortfolio                RetrieveOrderInitiationFlow = "PORTFOLIO"
	RetrieveOrderInitiationFlowCashDividendReinvestment RetrieveOrderInitiationFlow = "CASH_DIVIDEND_REINVESTMENT"
	RetrieveOrderInitiationFlowPortfolioRebalancing     RetrieveOrderInitiationFlow = "PORTFOLIO_REBALANCING"
	RetrieveOrderInitiationFlowSellToCoverFees          RetrieveOrderInitiationFlow = "SELL_TO_COVER_FEES"
	RetrieveOrderInitiationFlowSellToCoverTaxes         RetrieveOrderInitiationFlow = "SELL_TO_COVER_TAXES"
	RetrieveOrderInitiationFlowAccountLiquidation       RetrieveOrderInitiationFlow = "ACCOUNT_LIQUIDATION"
	RetrieveOrderInitiationFlowUpvestOperations         RetrieveOrderInitiationFlow = "UPVEST_OPERATIONS"
)

func (e RetrieveOrderInitiationFlow) ToPointer() *RetrieveOrderInitiationFlow {
	return &e
}

func (e *RetrieveOrderInitiationFlow) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "API":
		fallthrough
	case "PORTFOLIO":
		fallthrough
	case "CASH_DIVIDEND_REINVESTMENT":
		fallthrough
	case "PORTFOLIO_REBALANCING":
		fallthrough
	case "SELL_TO_COVER_FEES":
		fallthrough
	case "SELL_TO_COVER_TAXES":
		fallthrough
	case "ACCOUNT_LIQUIDATION":
		fallthrough
	case "UPVEST_OPERATIONS":
		*e = RetrieveOrderInitiationFlow(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveOrderInitiationFlow: %v", v)
	}
}

// RetrieveOrderInstrumentIDType - The type of the ID used in the request.
// * ISIN -
type RetrieveOrderInstrumentIDType string

const (
	RetrieveOrderInstrumentIDTypeIsin RetrieveOrderInstrumentIDType = "ISIN"
)

func (e RetrieveOrderInstrumentIDType) ToPointer() *RetrieveOrderInstrumentIDType {
	return &e
}

func (e *RetrieveOrderInstrumentIDType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ISIN":
		*e = RetrieveOrderInstrumentIDType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveOrderInstrumentIDType: %v", v)
	}
}

// RetrieveOrderOrderType - Type of the order.
// * MARKET -
// * LIMIT -
// * STOP -
type RetrieveOrderOrderType string

const (
	RetrieveOrderOrderTypeMarket RetrieveOrderOrderType = "MARKET"
	RetrieveOrderOrderTypeLimit  RetrieveOrderOrderType = "LIMIT"
	RetrieveOrderOrderTypeStop   RetrieveOrderOrderType = "STOP"
)

func (e RetrieveOrderOrderType) ToPointer() *RetrieveOrderOrderType {
	return &e
}

func (e *RetrieveOrderOrderType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MARKET":
		fallthrough
	case "LIMIT":
		fallthrough
	case "STOP":
		*e = RetrieveOrderOrderType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveOrderOrderType: %v", v)
	}
}

// RetrieveOrderSide - Side of the order.
// * BUY -
// * SELL -
type RetrieveOrderSide string

const (
	RetrieveOrderSideBuy  RetrieveOrderSide = "BUY"
	RetrieveOrderSideSell RetrieveOrderSide = "SELL"
)

func (e RetrieveOrderSide) ToPointer() *RetrieveOrderSide {
	return &e
}

func (e *RetrieveOrderSide) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BUY":
		fallthrough
	case "SELL":
		*e = RetrieveOrderSide(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveOrderSide: %v", v)
	}
}

// RetrieveOrderStatus - The execution status of the order.
// * NEW -
// * PROCESSING -
// * FILLED -
// * CANCELLED -
type RetrieveOrderStatus string

const (
	RetrieveOrderStatusNew        RetrieveOrderStatus = "NEW"
	RetrieveOrderStatusProcessing RetrieveOrderStatus = "PROCESSING"
	RetrieveOrderStatusFilled     RetrieveOrderStatus = "FILLED"
	RetrieveOrderStatusCancelled  RetrieveOrderStatus = "CANCELLED"
)

func (e RetrieveOrderStatus) ToPointer() *RetrieveOrderStatus {
	return &e
}

func (e *RetrieveOrderStatus) UnmarshalJSON(data []byte) error {
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
		*e = RetrieveOrderStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveOrderStatus: %v", v)
	}
}

// RetrieveOrderOrder - OK
type RetrieveOrderOrder struct {
	// The ID of the account that owns the order
	AccountID string `json:"account_id"`
	// Reason for Order cancellation. The field is present in case the Order has a status of CANCELLED.
	// * CANCELLED_BY_CLIENT -
	// * CANCELLED_BY_UPVEST_OPERATIONS -
	// * CANCELLED_BY_TRADING_PARTNER -
	// * CANCELLED_BY_UPVEST_PLATFORM -
	CancellationReason *RetrieveOrderCancellationReason `json:"cancellation_reason,omitempty"`
	CashAmount         string                           `json:"cash_amount"`
	// An ID provided by the client
	ClientReference string `json:"client_reference"`
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time              `json:"created_at"`
	Currency  *RetrieveOrderCurrency `default:"EUR" json:"currency"`
	// Execution flow that the order processing goes through. If no value is specified, the default value is assumed - `STRAIGHT_THROUGH`.
	// * STRAIGHT_THROUGH -
	// * BLOCK -
	ExecutionFlow *RetrieveOrderExecutionFlow `json:"execution_flow,omitempty"`
	// Order executions associated with this order
	Executions []RetrieveOrderOrderExecution `json:"executions"`
	ExpiryDate *string                       `json:"expiry_date,omitempty"`
	Fee        string                        `json:"fee"`
	ID         string                        `json:"id"`
	// Initiation flow used during order creation, i.e. what triggered the order.
	// * API -
	// * PORTFOLIO -
	// * CASH_DIVIDEND_REINVESTMENT -
	// * PORTFOLIO_REBALANCING -
	// * SELL_TO_COVER_FEES -
	// * SELL_TO_COVER_TAXES -
	// * ACCOUNT_LIQUIDATION -
	// * UPVEST_OPERATIONS -
	InitiationFlow RetrieveOrderInitiationFlow `json:"initiation_flow"`
	// International securities identification number defined by [ISO 6166](https://en.wikipedia.org/wiki/International_Securities_Identification_Number).
	InstrumentID string `json:"instrument_id"`
	// The type of the ID used in the request.
	// * ISIN -
	InstrumentIDType *RetrieveOrderInstrumentIDType `default:"ISIN" json:"instrument_id_type"`
	LimitPrice       *string                        `json:"limit_price,omitempty"`
	// Type of the order.
	// * MARKET -
	// * LIMIT -
	// * STOP -
	OrderType RetrieveOrderOrderType `json:"order_type"`
	Quantity  string                 `json:"quantity"`
	// Side of the order.
	// * BUY -
	// * SELL -
	Side RetrieveOrderSide `json:"side"`
	// The execution status of the order.
	// * NEW -
	// * PROCESSING -
	// * FILLED -
	// * CANCELLED -
	Status    RetrieveOrderStatus `json:"status"`
	StopPrice *string             `json:"stop_price,omitempty"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time `json:"updated_at"`
	// The ID of the user
	UserID string `json:"user_id"`
	// Only applicable if the user has failed the instrument fit check for the instrument type being ordered. True if the user has acknowledged their willingness to trade.
	UserInstrumentFitAcknowledgement *bool `json:"user_instrument_fit_acknowledgement,omitempty"`
}

func (r RetrieveOrderOrder) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrieveOrderOrder) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrieveOrderOrder) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *RetrieveOrderOrder) GetCancellationReason() *RetrieveOrderCancellationReason {
	if o == nil {
		return nil
	}
	return o.CancellationReason
}

func (o *RetrieveOrderOrder) GetCashAmount() string {
	if o == nil {
		return ""
	}
	return o.CashAmount
}

func (o *RetrieveOrderOrder) GetClientReference() string {
	if o == nil {
		return ""
	}
	return o.ClientReference
}

func (o *RetrieveOrderOrder) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *RetrieveOrderOrder) GetCurrency() *RetrieveOrderCurrency {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *RetrieveOrderOrder) GetExecutionFlow() *RetrieveOrderExecutionFlow {
	if o == nil {
		return nil
	}
	return o.ExecutionFlow
}

func (o *RetrieveOrderOrder) GetExecutions() []RetrieveOrderOrderExecution {
	if o == nil {
		return []RetrieveOrderOrderExecution{}
	}
	return o.Executions
}

func (o *RetrieveOrderOrder) GetExpiryDate() *string {
	if o == nil {
		return nil
	}
	return o.ExpiryDate
}

func (o *RetrieveOrderOrder) GetFee() string {
	if o == nil {
		return ""
	}
	return o.Fee
}

func (o *RetrieveOrderOrder) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveOrderOrder) GetInitiationFlow() RetrieveOrderInitiationFlow {
	if o == nil {
		return RetrieveOrderInitiationFlow("")
	}
	return o.InitiationFlow
}

func (o *RetrieveOrderOrder) GetInstrumentID() string {
	if o == nil {
		return ""
	}
	return o.InstrumentID
}

func (o *RetrieveOrderOrder) GetInstrumentIDType() *RetrieveOrderInstrumentIDType {
	if o == nil {
		return nil
	}
	return o.InstrumentIDType
}

func (o *RetrieveOrderOrder) GetLimitPrice() *string {
	if o == nil {
		return nil
	}
	return o.LimitPrice
}

func (o *RetrieveOrderOrder) GetOrderType() RetrieveOrderOrderType {
	if o == nil {
		return RetrieveOrderOrderType("")
	}
	return o.OrderType
}

func (o *RetrieveOrderOrder) GetQuantity() string {
	if o == nil {
		return ""
	}
	return o.Quantity
}

func (o *RetrieveOrderOrder) GetSide() RetrieveOrderSide {
	if o == nil {
		return RetrieveOrderSide("")
	}
	return o.Side
}

func (o *RetrieveOrderOrder) GetStatus() RetrieveOrderStatus {
	if o == nil {
		return RetrieveOrderStatus("")
	}
	return o.Status
}

func (o *RetrieveOrderOrder) GetStopPrice() *string {
	if o == nil {
		return nil
	}
	return o.StopPrice
}

func (o *RetrieveOrderOrder) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *RetrieveOrderOrder) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *RetrieveOrderOrder) GetUserInstrumentFitAcknowledgement() *bool {
	if o == nil {
		return nil
	}
	return o.UserInstrumentFitAcknowledgement
}

type RetrieveOrderResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// OK
	Order *RetrieveOrderOrder
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RetrieveOrderResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RetrieveOrderResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *RetrieveOrderResponse) GetOrder() *RetrieveOrderOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *RetrieveOrderResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RetrieveOrderResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
