// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/types"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/utils"
	"net/http"
	"time"
)

// ListAccountValuationHistoryQueryParamOrder - Sort order of the result list if the `sort` parameter is specified. Use `ASC` for ascending or `DESC` for descending sort order.
type ListAccountValuationHistoryQueryParamOrder string

const (
	ListAccountValuationHistoryQueryParamOrderAsc  ListAccountValuationHistoryQueryParamOrder = "ASC"
	ListAccountValuationHistoryQueryParamOrderDesc ListAccountValuationHistoryQueryParamOrder = "DESC"
)

func (e ListAccountValuationHistoryQueryParamOrder) ToPointer() *ListAccountValuationHistoryQueryParamOrder {
	return &e
}

func (e *ListAccountValuationHistoryQueryParamOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = ListAccountValuationHistoryQueryParamOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountValuationHistoryQueryParamOrder: %v", v)
	}
}

// ListAccountValuationHistoryQueryParamSort - Sort the result by `valuation_time`.
type ListAccountValuationHistoryQueryParamSort string

const (
	ListAccountValuationHistoryQueryParamSortValuationTime ListAccountValuationHistoryQueryParamSort = "valuation_time"
)

func (e ListAccountValuationHistoryQueryParamSort) ToPointer() *ListAccountValuationHistoryQueryParamSort {
	return &e
}

func (e *ListAccountValuationHistoryQueryParamSort) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "valuation_time":
		*e = ListAccountValuationHistoryQueryParamSort(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountValuationHistoryQueryParamSort: %v", v)
	}
}

type ListAccountValuationHistoryRequest struct {
	AccountID string      `pathParam:"style=simple,explode=false,name=account_id"`
	EndDate   *types.Date `queryParam:"style=form,explode=true,name=end_date"`
	// Use the `limit` argument to specify the maximum number of items returned.
	Limit *int `default:"100" queryParam:"style=form,explode=true,name=limit"`
	// Use the `offset` argument to specify where in the list of results to start when returning items for a particular query.
	Offset *int `queryParam:"style=form,explode=true,name=offset"`
	// Sort order of the result list if the `sort` parameter is specified. Use `ASC` for ascending or `DESC` for descending sort order.
	Order *ListAccountValuationHistoryQueryParamOrder `default:"ASC" queryParam:"style=form,explode=true,name=order"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-http-header
	Signature string `header:"style=simple,explode=false,name=signature"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-input-http-he
	SignatureInput string `header:"style=simple,explode=false,name=signature-input"`
	// Sort the result by `valuation_time`.
	Sort      *ListAccountValuationHistoryQueryParamSort `default:"valuation_time" queryParam:"style=form,explode=true,name=sort"`
	StartDate *types.Date                                `queryParam:"style=form,explode=true,name=start_date"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
}

func (l ListAccountValuationHistoryRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountValuationHistoryRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAccountValuationHistoryRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *ListAccountValuationHistoryRequest) GetEndDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *ListAccountValuationHistoryRequest) GetLimit() *int {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAccountValuationHistoryRequest) GetOffset() *int {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListAccountValuationHistoryRequest) GetOrder() *ListAccountValuationHistoryQueryParamOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListAccountValuationHistoryRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *ListAccountValuationHistoryRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *ListAccountValuationHistoryRequest) GetSort() *ListAccountValuationHistoryQueryParamSort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListAccountValuationHistoryRequest) GetStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *ListAccountValuationHistoryRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *ListAccountValuationHistoryRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// ListAccountValuationHistoryPriceQuality - Price quality used for the calculation of the account valuation.
// * EOD - end of day price
type ListAccountValuationHistoryPriceQuality string

const (
	ListAccountValuationHistoryPriceQualityEod ListAccountValuationHistoryPriceQuality = "EOD"
)

func (e ListAccountValuationHistoryPriceQuality) ToPointer() *ListAccountValuationHistoryPriceQuality {
	return &e
}

func (e *ListAccountValuationHistoryPriceQuality) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EOD":
		*e = ListAccountValuationHistoryPriceQuality(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountValuationHistoryPriceQuality: %v", v)
	}
}

// ListAccountValuationHistoryInstrument - Entity representing the financial instrument.
type ListAccountValuationHistoryInstrument struct {
	// International securities identification number defined by [ISO 6166](https://en.wikipedia.org/wiki/International_Securities_Identification_Number).
	Isin *string `json:"isin,omitempty"`
	// String representing the instrument internal identifier.
	UUID string `json:"uuid"`
}

func (o *ListAccountValuationHistoryInstrument) GetIsin() *string {
	if o == nil {
		return nil
	}
	return o.Isin
}

func (o *ListAccountValuationHistoryInstrument) GetUUID() string {
	if o == nil {
		return ""
	}
	return o.UUID
}

// ListAccountValuationHistoryValuationsCurrency - Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
// * EUR - Euro
type ListAccountValuationHistoryValuationsCurrency string

const (
	ListAccountValuationHistoryValuationsCurrencyEur ListAccountValuationHistoryValuationsCurrency = "EUR"
)

func (e ListAccountValuationHistoryValuationsCurrency) ToPointer() *ListAccountValuationHistoryValuationsCurrency {
	return &e
}

func (e *ListAccountValuationHistoryValuationsCurrency) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EUR":
		*e = ListAccountValuationHistoryValuationsCurrency(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountValuationHistoryValuationsCurrency: %v", v)
	}
}

// ListAccountValuationHistorySecurityPositionValue - Entity representing the monetary value by amount and currency, and the time of the price used.
type ListAccountValuationHistorySecurityPositionValue struct {
	Amount string `json:"amount"`
	// Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
	// * EUR - Euro
	Currency *ListAccountValuationHistoryValuationsCurrency `default:"EUR" json:"currency"`
	// The date and time of the price used for the calculation. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	PriceTime time.Time `json:"price_time"`
}

func (l ListAccountValuationHistorySecurityPositionValue) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountValuationHistorySecurityPositionValue) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAccountValuationHistorySecurityPositionValue) GetAmount() string {
	if o == nil {
		return ""
	}
	return o.Amount
}

func (o *ListAccountValuationHistorySecurityPositionValue) GetCurrency() *ListAccountValuationHistoryValuationsCurrency {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *ListAccountValuationHistorySecurityPositionValue) GetPriceTime() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.PriceTime
}

type ListAccountValuationHistoryAccountValuationSecurityPosition struct {
	// Entity representing the financial instrument.
	Instrument ListAccountValuationHistoryInstrument             `json:"instrument"`
	Quantity   string                                            `json:"quantity"`
	Value      *ListAccountValuationHistorySecurityPositionValue `json:"value"`
	Weight     *string                                           `json:"weight"`
}

func (o *ListAccountValuationHistoryAccountValuationSecurityPosition) GetInstrument() ListAccountValuationHistoryInstrument {
	if o == nil {
		return ListAccountValuationHistoryInstrument{}
	}
	return o.Instrument
}

func (o *ListAccountValuationHistoryAccountValuationSecurityPosition) GetQuantity() string {
	if o == nil {
		return ""
	}
	return o.Quantity
}

func (o *ListAccountValuationHistoryAccountValuationSecurityPosition) GetValue() *ListAccountValuationHistorySecurityPositionValue {
	if o == nil {
		return nil
	}
	return o.Value
}

func (o *ListAccountValuationHistoryAccountValuationSecurityPosition) GetWeight() *string {
	if o == nil {
		return nil
	}
	return o.Weight
}

// ListAccountValuationHistoryCurrency - Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
// * EUR - Euro
type ListAccountValuationHistoryCurrency string

const (
	ListAccountValuationHistoryCurrencyEur ListAccountValuationHistoryCurrency = "EUR"
)

func (e ListAccountValuationHistoryCurrency) ToPointer() *ListAccountValuationHistoryCurrency {
	return &e
}

func (e *ListAccountValuationHistoryCurrency) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EUR":
		*e = ListAccountValuationHistoryCurrency(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountValuationHistoryCurrency: %v", v)
	}
}

// ListAccountValuationHistoryTotalSecurityValue - Entity representing the monetary value by amount and currency.
type ListAccountValuationHistoryTotalSecurityValue struct {
	Amount string `json:"amount"`
	// Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
	// * EUR - Euro
	Currency *ListAccountValuationHistoryCurrency `default:"EUR" json:"currency"`
}

func (l ListAccountValuationHistoryTotalSecurityValue) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountValuationHistoryTotalSecurityValue) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAccountValuationHistoryTotalSecurityValue) GetAmount() string {
	if o == nil {
		return ""
	}
	return o.Amount
}

func (o *ListAccountValuationHistoryTotalSecurityValue) GetCurrency() *ListAccountValuationHistoryCurrency {
	if o == nil {
		return nil
	}
	return o.Currency
}

type AccountValuation struct {
	// Account unique identifier.
	AccountID string `json:"account_id"`
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	// Account valuation unique identifier.
	ID string `json:"id"`
	// Price quality used for the calculation of the account valuation.
	// * EOD - end of day price
	PriceQuality *ListAccountValuationHistoryPriceQuality `default:"EOD" json:"price_quality"`
	// Positions associated with this account valuation.
	SecurityPositions []ListAccountValuationHistoryAccountValuationSecurityPosition `json:"security_positions,omitempty"`
	// Entity representing the monetary value by amount and currency.
	TotalSecurityValue ListAccountValuationHistoryTotalSecurityValue `json:"total_security_value"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time `json:"updated_at"`
	// Date and time as of which the value was calculated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	ValuationTime time.Time `json:"valuation_time"`
}

func (a AccountValuation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountValuation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountValuation) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *AccountValuation) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *AccountValuation) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AccountValuation) GetPriceQuality() *ListAccountValuationHistoryPriceQuality {
	if o == nil {
		return nil
	}
	return o.PriceQuality
}

func (o *AccountValuation) GetSecurityPositions() []ListAccountValuationHistoryAccountValuationSecurityPosition {
	if o == nil {
		return nil
	}
	return o.SecurityPositions
}

func (o *AccountValuation) GetTotalSecurityValue() ListAccountValuationHistoryTotalSecurityValue {
	if o == nil {
		return ListAccountValuationHistoryTotalSecurityValue{}
	}
	return o.TotalSecurityValue
}

func (o *AccountValuation) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *AccountValuation) GetValuationTime() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.ValuationTime
}

// ListAccountValuationHistoryOrder - The ordering of the response.
// * ASC - Ascending order
// * DESC - Descending order
type ListAccountValuationHistoryOrder string

const (
	ListAccountValuationHistoryOrderAsc  ListAccountValuationHistoryOrder = "ASC"
	ListAccountValuationHistoryOrderDesc ListAccountValuationHistoryOrder = "DESC"
)

func (e ListAccountValuationHistoryOrder) ToPointer() *ListAccountValuationHistoryOrder {
	return &e
}

func (e *ListAccountValuationHistoryOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = ListAccountValuationHistoryOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountValuationHistoryOrder: %v", v)
	}
}

type ListAccountValuationHistoryMeta struct {
	// Count of the resources returned in the response.
	Count int64 `json:"count"`
	// Total limit of the response.
	Limit int64 `json:"limit"`
	// Amount of resource to offset in the response.
	Offset int64 `json:"offset"`
	// The ordering of the response.
	// * ASC - Ascending order
	// * DESC - Descending order
	Order *ListAccountValuationHistoryOrder `json:"order,omitempty"`
	// The field that the list is sorted by.
	Sort *string `json:"sort,omitempty"`
	// Total count of all the resources.
	TotalCount int64 `json:"total_count"`
}

func (o *ListAccountValuationHistoryMeta) GetCount() int64 {
	if o == nil {
		return 0
	}
	return o.Count
}

func (o *ListAccountValuationHistoryMeta) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

func (o *ListAccountValuationHistoryMeta) GetOffset() int64 {
	if o == nil {
		return 0
	}
	return o.Offset
}

func (o *ListAccountValuationHistoryMeta) GetOrder() *ListAccountValuationHistoryOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListAccountValuationHistoryMeta) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListAccountValuationHistoryMeta) GetTotalCount() int64 {
	if o == nil {
		return 0
	}
	return o.TotalCount
}

// ListAccountValuationHistoryAccountValuationListResponse - Valuations
type ListAccountValuationHistoryAccountValuationListResponse struct {
	Data []AccountValuation              `json:"data"`
	Meta ListAccountValuationHistoryMeta `json:"meta"`
}

func (o *ListAccountValuationHistoryAccountValuationListResponse) GetData() []AccountValuation {
	if o == nil {
		return []AccountValuation{}
	}
	return o.Data
}

func (o *ListAccountValuationHistoryAccountValuationListResponse) GetMeta() ListAccountValuationHistoryMeta {
	if o == nil {
		return ListAccountValuationHistoryMeta{}
	}
	return o.Meta
}

type ListAccountValuationHistoryResponse struct {
	// Valuations
	AccountValuationListResponse *ListAccountValuationHistoryAccountValuationListResponse
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListAccountValuationHistoryResponse) GetAccountValuationListResponse() *ListAccountValuationHistoryAccountValuationListResponse {
	if o == nil {
		return nil
	}
	return o.AccountValuationListResponse
}

func (o *ListAccountValuationHistoryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAccountValuationHistoryResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *ListAccountValuationHistoryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAccountValuationHistoryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
