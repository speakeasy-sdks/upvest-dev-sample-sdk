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

// ListAccountOrdersOrder - Sort order of the result list if the `sort` parameter is specified. Use `ASC` for ascending or `DESC` for descending sort order.
type ListAccountOrdersOrder string

const (
	ListAccountOrdersOrderAsc  ListAccountOrdersOrder = "ASC"
	ListAccountOrdersOrderDesc ListAccountOrdersOrder = "DESC"
)

func (e ListAccountOrdersOrder) ToPointer() *ListAccountOrdersOrder {
	return &e
}

func (e *ListAccountOrdersOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = ListAccountOrdersOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountOrdersOrder: %v", v)
	}
}

type ListAccountOrdersRequest struct {
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// Use the `limit` argument to specify the maximum number of items returned.
	Limit *int `default:"100" queryParam:"style=form,explode=true,name=limit"`
	// Use the `offset` argument to specify where in the list of results to start when returning items for a particular query.
	Offset *int `queryParam:"style=form,explode=true,name=offset"`
	// Sort order of the result list if the `sort` parameter is specified. Use `ASC` for ascending or `DESC` for descending sort order.
	Order *ListAccountOrdersOrder `default:"ASC" queryParam:"style=form,explode=true,name=order"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-http-header
	Signature string `header:"style=simple,explode=false,name=signature"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-input-http-he
	SignatureInput string `header:"style=simple,explode=false,name=signature-input"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
}

func (l ListAccountOrdersRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountOrdersRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAccountOrdersRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *ListAccountOrdersRequest) GetLimit() *int {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAccountOrdersRequest) GetOffset() *int {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListAccountOrdersRequest) GetOrder() *ListAccountOrdersOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListAccountOrdersRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *ListAccountOrdersRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *ListAccountOrdersRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *ListAccountOrdersRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// ListAccountOrdersOrdersListResponseOrderCancellationReason - Reason for Order cancellation. The field is present in case the Order has a status of CANCELLED.
// * CANCELLED_BY_CLIENT -
// * CANCELLED_BY_UPVEST_OPERATIONS -
// * CANCELLED_BY_TRADING_PARTNER -
// * CANCELLED_BY_UPVEST_PLATFORM -
type ListAccountOrdersOrdersListResponseOrderCancellationReason string

const (
	ListAccountOrdersOrdersListResponseOrderCancellationReasonCancelledByClient           ListAccountOrdersOrdersListResponseOrderCancellationReason = "CANCELLED_BY_CLIENT"
	ListAccountOrdersOrdersListResponseOrderCancellationReasonCancelledByUpvestOperations ListAccountOrdersOrdersListResponseOrderCancellationReason = "CANCELLED_BY_UPVEST_OPERATIONS"
	ListAccountOrdersOrdersListResponseOrderCancellationReasonCancelledByTradingPartner   ListAccountOrdersOrdersListResponseOrderCancellationReason = "CANCELLED_BY_TRADING_PARTNER"
	ListAccountOrdersOrdersListResponseOrderCancellationReasonCancelledByUpvestPlatform   ListAccountOrdersOrdersListResponseOrderCancellationReason = "CANCELLED_BY_UPVEST_PLATFORM"
)

func (e ListAccountOrdersOrdersListResponseOrderCancellationReason) ToPointer() *ListAccountOrdersOrdersListResponseOrderCancellationReason {
	return &e
}

func (e *ListAccountOrdersOrdersListResponseOrderCancellationReason) UnmarshalJSON(data []byte) error {
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
		*e = ListAccountOrdersOrdersListResponseOrderCancellationReason(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountOrdersOrdersListResponseOrderCancellationReason: %v", v)
	}
}

// ListAccountOrdersOrdersListResponseOrderCurrency - The currency for the order.
type ListAccountOrdersOrdersListResponseOrderCurrency string

const (
	ListAccountOrdersOrdersListResponseOrderCurrencyEur ListAccountOrdersOrdersListResponseOrderCurrency = "EUR"
)

func (e ListAccountOrdersOrdersListResponseOrderCurrency) ToPointer() *ListAccountOrdersOrdersListResponseOrderCurrency {
	return &e
}

func (e *ListAccountOrdersOrdersListResponseOrderCurrency) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EUR":
		*e = ListAccountOrdersOrdersListResponseOrderCurrency(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountOrdersOrdersListResponseOrderCurrency: %v", v)
	}
}

// ListAccountOrdersOrdersListResponseOrderExecutionFlow - Execution flow that the order processing goes through. If no value is specified, the default value is assumed - `STRAIGHT_THROUGH`.
// * STRAIGHT_THROUGH -
// * BLOCK -
type ListAccountOrdersOrdersListResponseOrderExecutionFlow string

const (
	ListAccountOrdersOrdersListResponseOrderExecutionFlowStraightThrough ListAccountOrdersOrdersListResponseOrderExecutionFlow = "STRAIGHT_THROUGH"
	ListAccountOrdersOrdersListResponseOrderExecutionFlowBlock           ListAccountOrdersOrdersListResponseOrderExecutionFlow = "BLOCK"
)

func (e ListAccountOrdersOrdersListResponseOrderExecutionFlow) ToPointer() *ListAccountOrdersOrdersListResponseOrderExecutionFlow {
	return &e
}

func (e *ListAccountOrdersOrdersListResponseOrderExecutionFlow) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "STRAIGHT_THROUGH":
		fallthrough
	case "BLOCK":
		*e = ListAccountOrdersOrdersListResponseOrderExecutionFlow(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountOrdersOrdersListResponseOrderExecutionFlow: %v", v)
	}
}

// ListAccountOrdersOrdersListResponseOrderOrderExecutionCurrency - Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
// * EUR - Euro
type ListAccountOrdersOrdersListResponseOrderOrderExecutionCurrency string

const (
	ListAccountOrdersOrdersListResponseOrderOrderExecutionCurrencyEur ListAccountOrdersOrdersListResponseOrderOrderExecutionCurrency = "EUR"
)

func (e ListAccountOrdersOrdersListResponseOrderOrderExecutionCurrency) ToPointer() *ListAccountOrdersOrdersListResponseOrderOrderExecutionCurrency {
	return &e
}

func (e *ListAccountOrdersOrdersListResponseOrderOrderExecutionCurrency) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EUR":
		*e = ListAccountOrdersOrdersListResponseOrderOrderExecutionCurrency(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountOrdersOrdersListResponseOrderOrderExecutionCurrency: %v", v)
	}
}

// ListAccountOrdersOrdersListResponseOrderOrderExecutionSide - Side of the execution.
// * BUY -
// * SELL -
type ListAccountOrdersOrdersListResponseOrderOrderExecutionSide string

const (
	ListAccountOrdersOrdersListResponseOrderOrderExecutionSideBuy  ListAccountOrdersOrdersListResponseOrderOrderExecutionSide = "BUY"
	ListAccountOrdersOrdersListResponseOrderOrderExecutionSideSell ListAccountOrdersOrdersListResponseOrderOrderExecutionSide = "SELL"
)

func (e ListAccountOrdersOrdersListResponseOrderOrderExecutionSide) ToPointer() *ListAccountOrdersOrdersListResponseOrderOrderExecutionSide {
	return &e
}

func (e *ListAccountOrdersOrdersListResponseOrderOrderExecutionSide) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BUY":
		fallthrough
	case "SELL":
		*e = ListAccountOrdersOrdersListResponseOrderOrderExecutionSide(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountOrdersOrdersListResponseOrderOrderExecutionSide: %v", v)
	}
}

// ListAccountOrdersOrdersListResponseOrderOrderExecutionStatus - Execution status of the Execution.
// * FILLED -
// * SETTLED -
// * CANCELLED -
type ListAccountOrdersOrdersListResponseOrderOrderExecutionStatus string

const (
	ListAccountOrdersOrdersListResponseOrderOrderExecutionStatusFilled    ListAccountOrdersOrdersListResponseOrderOrderExecutionStatus = "FILLED"
	ListAccountOrdersOrdersListResponseOrderOrderExecutionStatusSettled   ListAccountOrdersOrdersListResponseOrderOrderExecutionStatus = "SETTLED"
	ListAccountOrdersOrdersListResponseOrderOrderExecutionStatusCancelled ListAccountOrdersOrdersListResponseOrderOrderExecutionStatus = "CANCELLED"
)

func (e ListAccountOrdersOrdersListResponseOrderOrderExecutionStatus) ToPointer() *ListAccountOrdersOrdersListResponseOrderOrderExecutionStatus {
	return &e
}

func (e *ListAccountOrdersOrdersListResponseOrderOrderExecutionStatus) UnmarshalJSON(data []byte) error {
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
		*e = ListAccountOrdersOrdersListResponseOrderOrderExecutionStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountOrdersOrdersListResponseOrderOrderExecutionStatus: %v", v)
	}
}

// ListAccountOrdersOrdersListResponseOrderOrderExecutionTaxType - Tax type
// * TOTAL -
type ListAccountOrdersOrdersListResponseOrderOrderExecutionTaxType string

const (
	ListAccountOrdersOrdersListResponseOrderOrderExecutionTaxTypeTotal ListAccountOrdersOrdersListResponseOrderOrderExecutionTaxType = "TOTAL"
)

func (e ListAccountOrdersOrdersListResponseOrderOrderExecutionTaxType) ToPointer() *ListAccountOrdersOrdersListResponseOrderOrderExecutionTaxType {
	return &e
}

func (e *ListAccountOrdersOrdersListResponseOrderOrderExecutionTaxType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TOTAL":
		*e = ListAccountOrdersOrdersListResponseOrderOrderExecutionTaxType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountOrdersOrdersListResponseOrderOrderExecutionTaxType: %v", v)
	}
}

type ListAccountOrdersOrdersListResponseOrderOrderExecutionTax struct {
	Amount string `json:"amount"`
	// Tax type
	// * TOTAL -
	Type *ListAccountOrdersOrdersListResponseOrderOrderExecutionTaxType `default:"TOTAL" json:"type"`
}

func (l ListAccountOrdersOrdersListResponseOrderOrderExecutionTax) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountOrdersOrdersListResponseOrderOrderExecutionTax) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAccountOrdersOrdersListResponseOrderOrderExecutionTax) GetAmount() string {
	if o == nil {
		return ""
	}
	return o.Amount
}

func (o *ListAccountOrdersOrdersListResponseOrderOrderExecutionTax) GetType() *ListAccountOrdersOrdersListResponseOrderOrderExecutionTaxType {
	if o == nil {
		return nil
	}
	return o.Type
}

type ListAccountOrdersOrdersListResponseOrderOrderExecution struct {
	CashAmount string `json:"cash_amount"`
	// Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
	// * EUR - Euro
	Currency       *ListAccountOrdersOrdersListResponseOrderOrderExecutionCurrency `default:"EUR" json:"currency"`
	ID             string                                                          `json:"id"`
	OrderID        string                                                          `json:"order_id"`
	Price          string                                                          `json:"price"`
	SettlementDate *string                                                         `json:"settlement_date,omitempty"`
	ShareQuantity  string                                                          `json:"share_quantity"`
	// Side of the execution.
	// * BUY -
	// * SELL -
	Side ListAccountOrdersOrdersListResponseOrderOrderExecutionSide `json:"side"`
	// Execution status of the Execution.
	// * FILLED -
	// * SETTLED -
	// * CANCELLED -
	Status          ListAccountOrdersOrdersListResponseOrderOrderExecutionStatus `json:"status"`
	Taxes           []ListAccountOrdersOrdersListResponseOrderOrderExecutionTax  `json:"taxes"`
	TransactionTime time.Time                                                    `json:"transaction_time"`
}

func (l ListAccountOrdersOrdersListResponseOrderOrderExecution) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountOrdersOrdersListResponseOrderOrderExecution) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAccountOrdersOrdersListResponseOrderOrderExecution) GetCashAmount() string {
	if o == nil {
		return ""
	}
	return o.CashAmount
}

func (o *ListAccountOrdersOrdersListResponseOrderOrderExecution) GetCurrency() *ListAccountOrdersOrdersListResponseOrderOrderExecutionCurrency {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *ListAccountOrdersOrdersListResponseOrderOrderExecution) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ListAccountOrdersOrdersListResponseOrderOrderExecution) GetOrderID() string {
	if o == nil {
		return ""
	}
	return o.OrderID
}

func (o *ListAccountOrdersOrdersListResponseOrderOrderExecution) GetPrice() string {
	if o == nil {
		return ""
	}
	return o.Price
}

func (o *ListAccountOrdersOrdersListResponseOrderOrderExecution) GetSettlementDate() *string {
	if o == nil {
		return nil
	}
	return o.SettlementDate
}

func (o *ListAccountOrdersOrdersListResponseOrderOrderExecution) GetShareQuantity() string {
	if o == nil {
		return ""
	}
	return o.ShareQuantity
}

func (o *ListAccountOrdersOrdersListResponseOrderOrderExecution) GetSide() ListAccountOrdersOrdersListResponseOrderOrderExecutionSide {
	if o == nil {
		return ListAccountOrdersOrdersListResponseOrderOrderExecutionSide("")
	}
	return o.Side
}

func (o *ListAccountOrdersOrdersListResponseOrderOrderExecution) GetStatus() ListAccountOrdersOrdersListResponseOrderOrderExecutionStatus {
	if o == nil {
		return ListAccountOrdersOrdersListResponseOrderOrderExecutionStatus("")
	}
	return o.Status
}

func (o *ListAccountOrdersOrdersListResponseOrderOrderExecution) GetTaxes() []ListAccountOrdersOrdersListResponseOrderOrderExecutionTax {
	if o == nil {
		return []ListAccountOrdersOrdersListResponseOrderOrderExecutionTax{}
	}
	return o.Taxes
}

func (o *ListAccountOrdersOrdersListResponseOrderOrderExecution) GetTransactionTime() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.TransactionTime
}

// ListAccountOrdersOrdersListResponseOrderInitiationFlow - Initiation flow used during order creation, i.e. what triggered the order.
// * API -
// * PORTFOLIO -
// * CASH_DIVIDEND_REINVESTMENT -
// * PORTFOLIO_REBALANCING -
// * SELL_TO_COVER_FEES -
// * SELL_TO_COVER_TAXES -
// * ACCOUNT_LIQUIDATION -
// * UPVEST_OPERATIONS -
type ListAccountOrdersOrdersListResponseOrderInitiationFlow string

const (
	ListAccountOrdersOrdersListResponseOrderInitiationFlowAPI                      ListAccountOrdersOrdersListResponseOrderInitiationFlow = "API"
	ListAccountOrdersOrdersListResponseOrderInitiationFlowPortfolio                ListAccountOrdersOrdersListResponseOrderInitiationFlow = "PORTFOLIO"
	ListAccountOrdersOrdersListResponseOrderInitiationFlowCashDividendReinvestment ListAccountOrdersOrdersListResponseOrderInitiationFlow = "CASH_DIVIDEND_REINVESTMENT"
	ListAccountOrdersOrdersListResponseOrderInitiationFlowPortfolioRebalancing     ListAccountOrdersOrdersListResponseOrderInitiationFlow = "PORTFOLIO_REBALANCING"
	ListAccountOrdersOrdersListResponseOrderInitiationFlowSellToCoverFees          ListAccountOrdersOrdersListResponseOrderInitiationFlow = "SELL_TO_COVER_FEES"
	ListAccountOrdersOrdersListResponseOrderInitiationFlowSellToCoverTaxes         ListAccountOrdersOrdersListResponseOrderInitiationFlow = "SELL_TO_COVER_TAXES"
	ListAccountOrdersOrdersListResponseOrderInitiationFlowAccountLiquidation       ListAccountOrdersOrdersListResponseOrderInitiationFlow = "ACCOUNT_LIQUIDATION"
	ListAccountOrdersOrdersListResponseOrderInitiationFlowUpvestOperations         ListAccountOrdersOrdersListResponseOrderInitiationFlow = "UPVEST_OPERATIONS"
)

func (e ListAccountOrdersOrdersListResponseOrderInitiationFlow) ToPointer() *ListAccountOrdersOrdersListResponseOrderInitiationFlow {
	return &e
}

func (e *ListAccountOrdersOrdersListResponseOrderInitiationFlow) UnmarshalJSON(data []byte) error {
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
		*e = ListAccountOrdersOrdersListResponseOrderInitiationFlow(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountOrdersOrdersListResponseOrderInitiationFlow: %v", v)
	}
}

// ListAccountOrdersOrdersListResponseOrderInstrumentIDType - The type of the ID used in the request.
// * ISIN -
type ListAccountOrdersOrdersListResponseOrderInstrumentIDType string

const (
	ListAccountOrdersOrdersListResponseOrderInstrumentIDTypeIsin ListAccountOrdersOrdersListResponseOrderInstrumentIDType = "ISIN"
)

func (e ListAccountOrdersOrdersListResponseOrderInstrumentIDType) ToPointer() *ListAccountOrdersOrdersListResponseOrderInstrumentIDType {
	return &e
}

func (e *ListAccountOrdersOrdersListResponseOrderInstrumentIDType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ISIN":
		*e = ListAccountOrdersOrdersListResponseOrderInstrumentIDType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountOrdersOrdersListResponseOrderInstrumentIDType: %v", v)
	}
}

// ListAccountOrdersOrdersListResponseOrderOrderType - Type of the order.
// * MARKET -
// * LIMIT -
// * STOP -
type ListAccountOrdersOrdersListResponseOrderOrderType string

const (
	ListAccountOrdersOrdersListResponseOrderOrderTypeMarket ListAccountOrdersOrdersListResponseOrderOrderType = "MARKET"
	ListAccountOrdersOrdersListResponseOrderOrderTypeLimit  ListAccountOrdersOrdersListResponseOrderOrderType = "LIMIT"
	ListAccountOrdersOrdersListResponseOrderOrderTypeStop   ListAccountOrdersOrdersListResponseOrderOrderType = "STOP"
)

func (e ListAccountOrdersOrdersListResponseOrderOrderType) ToPointer() *ListAccountOrdersOrdersListResponseOrderOrderType {
	return &e
}

func (e *ListAccountOrdersOrdersListResponseOrderOrderType) UnmarshalJSON(data []byte) error {
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
		*e = ListAccountOrdersOrdersListResponseOrderOrderType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountOrdersOrdersListResponseOrderOrderType: %v", v)
	}
}

// ListAccountOrdersOrdersListResponseOrderSide - Side of the order.
// * BUY -
// * SELL -
type ListAccountOrdersOrdersListResponseOrderSide string

const (
	ListAccountOrdersOrdersListResponseOrderSideBuy  ListAccountOrdersOrdersListResponseOrderSide = "BUY"
	ListAccountOrdersOrdersListResponseOrderSideSell ListAccountOrdersOrdersListResponseOrderSide = "SELL"
)

func (e ListAccountOrdersOrdersListResponseOrderSide) ToPointer() *ListAccountOrdersOrdersListResponseOrderSide {
	return &e
}

func (e *ListAccountOrdersOrdersListResponseOrderSide) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BUY":
		fallthrough
	case "SELL":
		*e = ListAccountOrdersOrdersListResponseOrderSide(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountOrdersOrdersListResponseOrderSide: %v", v)
	}
}

// ListAccountOrdersOrdersListResponseOrderStatus - The execution status of the order.
// * NEW -
// * PROCESSING -
// * FILLED -
// * CANCELLED -
type ListAccountOrdersOrdersListResponseOrderStatus string

const (
	ListAccountOrdersOrdersListResponseOrderStatusNew        ListAccountOrdersOrdersListResponseOrderStatus = "NEW"
	ListAccountOrdersOrdersListResponseOrderStatusProcessing ListAccountOrdersOrdersListResponseOrderStatus = "PROCESSING"
	ListAccountOrdersOrdersListResponseOrderStatusFilled     ListAccountOrdersOrdersListResponseOrderStatus = "FILLED"
	ListAccountOrdersOrdersListResponseOrderStatusCancelled  ListAccountOrdersOrdersListResponseOrderStatus = "CANCELLED"
)

func (e ListAccountOrdersOrdersListResponseOrderStatus) ToPointer() *ListAccountOrdersOrdersListResponseOrderStatus {
	return &e
}

func (e *ListAccountOrdersOrdersListResponseOrderStatus) UnmarshalJSON(data []byte) error {
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
		*e = ListAccountOrdersOrdersListResponseOrderStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountOrdersOrdersListResponseOrderStatus: %v", v)
	}
}

type ListAccountOrdersOrdersListResponseOrder struct {
	// The ID of the account that owns the order
	AccountID string `json:"account_id"`
	// Reason for Order cancellation. The field is present in case the Order has a status of CANCELLED.
	// * CANCELLED_BY_CLIENT -
	// * CANCELLED_BY_UPVEST_OPERATIONS -
	// * CANCELLED_BY_TRADING_PARTNER -
	// * CANCELLED_BY_UPVEST_PLATFORM -
	CancellationReason *ListAccountOrdersOrdersListResponseOrderCancellationReason `json:"cancellation_reason,omitempty"`
	CashAmount         string                                                      `json:"cash_amount"`
	// An ID provided by the client
	ClientReference string `json:"client_reference"`
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time                                         `json:"created_at"`
	Currency  *ListAccountOrdersOrdersListResponseOrderCurrency `default:"EUR" json:"currency"`
	// Execution flow that the order processing goes through. If no value is specified, the default value is assumed - `STRAIGHT_THROUGH`.
	// * STRAIGHT_THROUGH -
	// * BLOCK -
	ExecutionFlow *ListAccountOrdersOrdersListResponseOrderExecutionFlow `json:"execution_flow,omitempty"`
	// Order executions associated with this order
	Executions []ListAccountOrdersOrdersListResponseOrderOrderExecution `json:"executions"`
	ExpiryDate *string                                                  `json:"expiry_date,omitempty"`
	Fee        string                                                   `json:"fee"`
	ID         string                                                   `json:"id"`
	// Initiation flow used during order creation, i.e. what triggered the order.
	// * API -
	// * PORTFOLIO -
	// * CASH_DIVIDEND_REINVESTMENT -
	// * PORTFOLIO_REBALANCING -
	// * SELL_TO_COVER_FEES -
	// * SELL_TO_COVER_TAXES -
	// * ACCOUNT_LIQUIDATION -
	// * UPVEST_OPERATIONS -
	InitiationFlow ListAccountOrdersOrdersListResponseOrderInitiationFlow `json:"initiation_flow"`
	// International securities identification number defined by [ISO 6166](https://en.wikipedia.org/wiki/International_Securities_Identification_Number).
	InstrumentID string `json:"instrument_id"`
	// The type of the ID used in the request.
	// * ISIN -
	InstrumentIDType *ListAccountOrdersOrdersListResponseOrderInstrumentIDType `default:"ISIN" json:"instrument_id_type"`
	LimitPrice       *string                                                   `json:"limit_price,omitempty"`
	// Type of the order.
	// * MARKET -
	// * LIMIT -
	// * STOP -
	OrderType ListAccountOrdersOrdersListResponseOrderOrderType `json:"order_type"`
	Quantity  string                                            `json:"quantity"`
	// Side of the order.
	// * BUY -
	// * SELL -
	Side ListAccountOrdersOrdersListResponseOrderSide `json:"side"`
	// The execution status of the order.
	// * NEW -
	// * PROCESSING -
	// * FILLED -
	// * CANCELLED -
	Status    ListAccountOrdersOrdersListResponseOrderStatus `json:"status"`
	StopPrice *string                                        `json:"stop_price,omitempty"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time `json:"updated_at"`
	// The ID of the user
	UserID string `json:"user_id"`
	// Only applicable if the user has failed the instrument fit check for the instrument type being ordered. True if the user has acknowledged their willingness to trade.
	UserInstrumentFitAcknowledgement *bool `json:"user_instrument_fit_acknowledgement,omitempty"`
}

func (l ListAccountOrdersOrdersListResponseOrder) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountOrdersOrdersListResponseOrder) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAccountOrdersOrdersListResponseOrder) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *ListAccountOrdersOrdersListResponseOrder) GetCancellationReason() *ListAccountOrdersOrdersListResponseOrderCancellationReason {
	if o == nil {
		return nil
	}
	return o.CancellationReason
}

func (o *ListAccountOrdersOrdersListResponseOrder) GetCashAmount() string {
	if o == nil {
		return ""
	}
	return o.CashAmount
}

func (o *ListAccountOrdersOrdersListResponseOrder) GetClientReference() string {
	if o == nil {
		return ""
	}
	return o.ClientReference
}

func (o *ListAccountOrdersOrdersListResponseOrder) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *ListAccountOrdersOrdersListResponseOrder) GetCurrency() *ListAccountOrdersOrdersListResponseOrderCurrency {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *ListAccountOrdersOrdersListResponseOrder) GetExecutionFlow() *ListAccountOrdersOrdersListResponseOrderExecutionFlow {
	if o == nil {
		return nil
	}
	return o.ExecutionFlow
}

func (o *ListAccountOrdersOrdersListResponseOrder) GetExecutions() []ListAccountOrdersOrdersListResponseOrderOrderExecution {
	if o == nil {
		return []ListAccountOrdersOrdersListResponseOrderOrderExecution{}
	}
	return o.Executions
}

func (o *ListAccountOrdersOrdersListResponseOrder) GetExpiryDate() *string {
	if o == nil {
		return nil
	}
	return o.ExpiryDate
}

func (o *ListAccountOrdersOrdersListResponseOrder) GetFee() string {
	if o == nil {
		return ""
	}
	return o.Fee
}

func (o *ListAccountOrdersOrdersListResponseOrder) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ListAccountOrdersOrdersListResponseOrder) GetInitiationFlow() ListAccountOrdersOrdersListResponseOrderInitiationFlow {
	if o == nil {
		return ListAccountOrdersOrdersListResponseOrderInitiationFlow("")
	}
	return o.InitiationFlow
}

func (o *ListAccountOrdersOrdersListResponseOrder) GetInstrumentID() string {
	if o == nil {
		return ""
	}
	return o.InstrumentID
}

func (o *ListAccountOrdersOrdersListResponseOrder) GetInstrumentIDType() *ListAccountOrdersOrdersListResponseOrderInstrumentIDType {
	if o == nil {
		return nil
	}
	return o.InstrumentIDType
}

func (o *ListAccountOrdersOrdersListResponseOrder) GetLimitPrice() *string {
	if o == nil {
		return nil
	}
	return o.LimitPrice
}

func (o *ListAccountOrdersOrdersListResponseOrder) GetOrderType() ListAccountOrdersOrdersListResponseOrderOrderType {
	if o == nil {
		return ListAccountOrdersOrdersListResponseOrderOrderType("")
	}
	return o.OrderType
}

func (o *ListAccountOrdersOrdersListResponseOrder) GetQuantity() string {
	if o == nil {
		return ""
	}
	return o.Quantity
}

func (o *ListAccountOrdersOrdersListResponseOrder) GetSide() ListAccountOrdersOrdersListResponseOrderSide {
	if o == nil {
		return ListAccountOrdersOrdersListResponseOrderSide("")
	}
	return o.Side
}

func (o *ListAccountOrdersOrdersListResponseOrder) GetStatus() ListAccountOrdersOrdersListResponseOrderStatus {
	if o == nil {
		return ListAccountOrdersOrdersListResponseOrderStatus("")
	}
	return o.Status
}

func (o *ListAccountOrdersOrdersListResponseOrder) GetStopPrice() *string {
	if o == nil {
		return nil
	}
	return o.StopPrice
}

func (o *ListAccountOrdersOrdersListResponseOrder) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *ListAccountOrdersOrdersListResponseOrder) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *ListAccountOrdersOrdersListResponseOrder) GetUserInstrumentFitAcknowledgement() *bool {
	if o == nil {
		return nil
	}
	return o.UserInstrumentFitAcknowledgement
}

// ListAccountOrdersOrdersListResponseMetaOrder - The ordering of the response.
// * ASC - Ascending order
// * DESC - Descending order
type ListAccountOrdersOrdersListResponseMetaOrder string

const (
	ListAccountOrdersOrdersListResponseMetaOrderAsc  ListAccountOrdersOrdersListResponseMetaOrder = "ASC"
	ListAccountOrdersOrdersListResponseMetaOrderDesc ListAccountOrdersOrdersListResponseMetaOrder = "DESC"
)

func (e ListAccountOrdersOrdersListResponseMetaOrder) ToPointer() *ListAccountOrdersOrdersListResponseMetaOrder {
	return &e
}

func (e *ListAccountOrdersOrdersListResponseMetaOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = ListAccountOrdersOrdersListResponseMetaOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountOrdersOrdersListResponseMetaOrder: %v", v)
	}
}

type ListAccountOrdersOrdersListResponseMeta struct {
	// Count of the resources returned in the response.
	Count int64 `json:"count"`
	// Total limit of the response.
	Limit int64 `json:"limit"`
	// Amount of resource to offset in the response.
	Offset int64 `json:"offset"`
	// The ordering of the response.
	// * ASC - Ascending order
	// * DESC - Descending order
	Order *ListAccountOrdersOrdersListResponseMetaOrder `json:"order,omitempty"`
	// The field that the list is sorted by.
	Sort *string `json:"sort,omitempty"`
	// Total count of all the resources.
	TotalCount int64 `json:"total_count"`
}

func (o *ListAccountOrdersOrdersListResponseMeta) GetCount() int64 {
	if o == nil {
		return 0
	}
	return o.Count
}

func (o *ListAccountOrdersOrdersListResponseMeta) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

func (o *ListAccountOrdersOrdersListResponseMeta) GetOffset() int64 {
	if o == nil {
		return 0
	}
	return o.Offset
}

func (o *ListAccountOrdersOrdersListResponseMeta) GetOrder() *ListAccountOrdersOrdersListResponseMetaOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListAccountOrdersOrdersListResponseMeta) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListAccountOrdersOrdersListResponseMeta) GetTotalCount() int64 {
	if o == nil {
		return 0
	}
	return o.TotalCount
}

// ListAccountOrdersOrdersListResponse - OK
type ListAccountOrdersOrdersListResponse struct {
	Data []ListAccountOrdersOrdersListResponseOrder `json:"data"`
	Meta ListAccountOrdersOrdersListResponseMeta    `json:"meta"`
}

func (o *ListAccountOrdersOrdersListResponse) GetData() []ListAccountOrdersOrdersListResponseOrder {
	if o == nil {
		return []ListAccountOrdersOrdersListResponseOrder{}
	}
	return o.Data
}

func (o *ListAccountOrdersOrdersListResponse) GetMeta() ListAccountOrdersOrdersListResponseMeta {
	if o == nil {
		return ListAccountOrdersOrdersListResponseMeta{}
	}
	return o.Meta
}

type ListAccountOrdersResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// OK
	OrdersListResponse *ListAccountOrdersOrdersListResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListAccountOrdersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAccountOrdersResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *ListAccountOrdersResponse) GetOrdersListResponse() *ListAccountOrdersOrdersListResponse {
	if o == nil {
		return nil
	}
	return o.OrdersListResponse
}

func (o *ListAccountOrdersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAccountOrdersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
