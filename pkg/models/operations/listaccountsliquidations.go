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

// ListAccountsLiquidationsQueryParamSort - Sort the result by `id`.
type ListAccountsLiquidationsQueryParamSort string

const (
	ListAccountsLiquidationsQueryParamSortID ListAccountsLiquidationsQueryParamSort = "id"
)

func (e ListAccountsLiquidationsQueryParamSort) ToPointer() *ListAccountsLiquidationsQueryParamSort {
	return &e
}

func (e *ListAccountsLiquidationsQueryParamSort) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "id":
		*e = ListAccountsLiquidationsQueryParamSort(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountsLiquidationsQueryParamSort: %v", v)
	}
}

type ListAccountsLiquidationsRequest struct {
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// Returns accounts liquidations created up until this date (UTC)
	EndDate *string `queryParam:"style=form,explode=true,name=end_date"`
	// Use the `limit` argument to specify the maximum number of items returned.
	Limit *int `default:"100" queryParam:"style=form,explode=true,name=limit"`
	// Use the `offset` argument to specify where in the list of results to start when returning items for a particular query.
	Offset *int `queryParam:"style=form,explode=true,name=offset"`
	// Sort order of the result list if the `sort` parameter is specified. Use `ASC` for ascending or `DESC` for descending sort order.
	Order *shared.Order `default:"ASC" queryParam:"style=form,explode=true,name=order"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-http-header
	Signature string `header:"style=simple,explode=false,name=signature"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-input-http-he
	SignatureInput string `header:"style=simple,explode=false,name=signature-input"`
	// Sort the result by `id`.
	Sort *ListAccountsLiquidationsQueryParamSort `default:"id" queryParam:"style=form,explode=true,name=sort"`
	// Returns accounts liquidations created starting from and including this date (UTC)
	StartDate *string `queryParam:"style=form,explode=true,name=start_date"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
}

func (l ListAccountsLiquidationsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountsLiquidationsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAccountsLiquidationsRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *ListAccountsLiquidationsRequest) GetEndDate() *string {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *ListAccountsLiquidationsRequest) GetLimit() *int {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAccountsLiquidationsRequest) GetOffset() *int {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListAccountsLiquidationsRequest) GetOrder() *shared.Order {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListAccountsLiquidationsRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *ListAccountsLiquidationsRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *ListAccountsLiquidationsRequest) GetSort() *ListAccountsLiquidationsQueryParamSort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListAccountsLiquidationsRequest) GetStartDate() *string {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *ListAccountsLiquidationsRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *ListAccountsLiquidationsRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// ListAccountsLiquidationsCurrency - Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
// * EUR - Euro
type ListAccountsLiquidationsCurrency string

const (
	ListAccountsLiquidationsCurrencyEur ListAccountsLiquidationsCurrency = "EUR"
)

func (e ListAccountsLiquidationsCurrency) ToPointer() *ListAccountsLiquidationsCurrency {
	return &e
}

func (e *ListAccountsLiquidationsCurrency) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EUR":
		*e = ListAccountsLiquidationsCurrency(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountsLiquidationsCurrency: %v", v)
	}
}

// ListAccountsLiquidationsLiquidationsStatus - Execution status of the Account liquidation order.
// * NEW -
// * PROCESSING -
// * FILLED -
// * CANCELLED -
type ListAccountsLiquidationsLiquidationsStatus string

const (
	ListAccountsLiquidationsLiquidationsStatusNew        ListAccountsLiquidationsLiquidationsStatus = "NEW"
	ListAccountsLiquidationsLiquidationsStatusProcessing ListAccountsLiquidationsLiquidationsStatus = "PROCESSING"
	ListAccountsLiquidationsLiquidationsStatusFilled     ListAccountsLiquidationsLiquidationsStatus = "FILLED"
	ListAccountsLiquidationsLiquidationsStatusCancelled  ListAccountsLiquidationsLiquidationsStatus = "CANCELLED"
)

func (e ListAccountsLiquidationsLiquidationsStatus) ToPointer() *ListAccountsLiquidationsLiquidationsStatus {
	return &e
}

func (e *ListAccountsLiquidationsLiquidationsStatus) UnmarshalJSON(data []byte) error {
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
		*e = ListAccountsLiquidationsLiquidationsStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountsLiquidationsLiquidationsStatus: %v", v)
	}
}

type ListAccountsLiquidationsLiquidationsAccountLiquidation struct {
	ID string `json:"id"`
	// Side of the order.
	// * SELL -
	Side *string `default:"SELL" json:"side"`
	// Execution status of the Account liquidation order.
	// * NEW -
	// * PROCESSING -
	// * FILLED -
	// * CANCELLED -
	Status ListAccountsLiquidationsLiquidationsStatus `json:"status"`
}

func (l ListAccountsLiquidationsLiquidationsAccountLiquidation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountsLiquidationsLiquidationsAccountLiquidation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAccountsLiquidationsLiquidationsAccountLiquidation) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ListAccountsLiquidationsLiquidationsAccountLiquidation) GetSide() *string {
	if o == nil {
		return nil
	}
	return o.Side
}

func (o *ListAccountsLiquidationsLiquidationsAccountLiquidation) GetStatus() ListAccountsLiquidationsLiquidationsStatus {
	if o == nil {
		return ListAccountsLiquidationsLiquidationsStatus("")
	}
	return o.Status
}

// ListAccountsLiquidationsStatus - Execution status of the Account liquidation.
// * NEW -
// * PROCESSING -
// * FILLED -
// * CANCELLED -
// * SETTLED -
type ListAccountsLiquidationsStatus string

const (
	ListAccountsLiquidationsStatusNew        ListAccountsLiquidationsStatus = "NEW"
	ListAccountsLiquidationsStatusProcessing ListAccountsLiquidationsStatus = "PROCESSING"
	ListAccountsLiquidationsStatusFilled     ListAccountsLiquidationsStatus = "FILLED"
	ListAccountsLiquidationsStatusCancelled  ListAccountsLiquidationsStatus = "CANCELLED"
	ListAccountsLiquidationsStatusSettled    ListAccountsLiquidationsStatus = "SETTLED"
)

func (e ListAccountsLiquidationsStatus) ToPointer() *ListAccountsLiquidationsStatus {
	return &e
}

func (e *ListAccountsLiquidationsStatus) UnmarshalJSON(data []byte) error {
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
		*e = ListAccountsLiquidationsStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountsLiquidationsStatus: %v", v)
	}
}

type ListAccountsLiquidationsAccountLiquidation struct {
	// Account unique identifier.
	AccountID  string `json:"account_id"`
	CashAmount string `json:"cash_amount"`
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	// Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
	// * EUR - Euro
	Currency *ListAccountsLiquidationsCurrency `default:"EUR" json:"currency"`
	ID       string                            `json:"id"`
	// Position liquidation orders associated with this account liquidation
	Orders []ListAccountsLiquidationsLiquidationsAccountLiquidation `json:"orders"`
	// Execution status of the Account liquidation.
	// * NEW -
	// * PROCESSING -
	// * FILLED -
	// * CANCELLED -
	// * SETTLED -
	Status ListAccountsLiquidationsStatus `json:"status"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time `json:"updated_at"`
	// User unique identifier.
	UserID *string `json:"user_id,omitempty"`
}

func (l ListAccountsLiquidationsAccountLiquidation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountsLiquidationsAccountLiquidation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAccountsLiquidationsAccountLiquidation) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *ListAccountsLiquidationsAccountLiquidation) GetCashAmount() string {
	if o == nil {
		return ""
	}
	return o.CashAmount
}

func (o *ListAccountsLiquidationsAccountLiquidation) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *ListAccountsLiquidationsAccountLiquidation) GetCurrency() *ListAccountsLiquidationsCurrency {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *ListAccountsLiquidationsAccountLiquidation) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ListAccountsLiquidationsAccountLiquidation) GetOrders() []ListAccountsLiquidationsLiquidationsAccountLiquidation {
	if o == nil {
		return []ListAccountsLiquidationsLiquidationsAccountLiquidation{}
	}
	return o.Orders
}

func (o *ListAccountsLiquidationsAccountLiquidation) GetStatus() ListAccountsLiquidationsStatus {
	if o == nil {
		return ListAccountsLiquidationsStatus("")
	}
	return o.Status
}

func (o *ListAccountsLiquidationsAccountLiquidation) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *ListAccountsLiquidationsAccountLiquidation) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

// ListAccountsLiquidationsOrder - The ordering of the response.
// * ASC - Ascending order
// * DESC - Descending order
type ListAccountsLiquidationsOrder string

const (
	ListAccountsLiquidationsOrderAsc  ListAccountsLiquidationsOrder = "ASC"
	ListAccountsLiquidationsOrderDesc ListAccountsLiquidationsOrder = "DESC"
)

func (e ListAccountsLiquidationsOrder) ToPointer() *ListAccountsLiquidationsOrder {
	return &e
}

func (e *ListAccountsLiquidationsOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = ListAccountsLiquidationsOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountsLiquidationsOrder: %v", v)
	}
}

type ListAccountsLiquidationsMeta struct {
	// Count of the resources returned in the response.
	Count int64 `json:"count"`
	// Total limit of the response.
	Limit int64 `json:"limit"`
	// Amount of resource to offset in the response.
	Offset int64 `json:"offset"`
	// The ordering of the response.
	// * ASC - Ascending order
	// * DESC - Descending order
	Order *ListAccountsLiquidationsOrder `json:"order,omitempty"`
	// The field that the list is sorted by.
	Sort *string `json:"sort,omitempty"`
	// Total count of all the resources.
	TotalCount int64 `json:"total_count"`
}

func (o *ListAccountsLiquidationsMeta) GetCount() int64 {
	if o == nil {
		return 0
	}
	return o.Count
}

func (o *ListAccountsLiquidationsMeta) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

func (o *ListAccountsLiquidationsMeta) GetOffset() int64 {
	if o == nil {
		return 0
	}
	return o.Offset
}

func (o *ListAccountsLiquidationsMeta) GetOrder() *ListAccountsLiquidationsOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListAccountsLiquidationsMeta) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListAccountsLiquidationsMeta) GetTotalCount() int64 {
	if o == nil {
		return 0
	}
	return o.TotalCount
}

// ListAccountsLiquidationsPortfoliosOrdersListResponse - Accounts liquidations
type ListAccountsLiquidationsPortfoliosOrdersListResponse struct {
	Data []ListAccountsLiquidationsAccountLiquidation `json:"data"`
	Meta ListAccountsLiquidationsMeta                 `json:"meta"`
}

func (o *ListAccountsLiquidationsPortfoliosOrdersListResponse) GetData() []ListAccountsLiquidationsAccountLiquidation {
	if o == nil {
		return []ListAccountsLiquidationsAccountLiquidation{}
	}
	return o.Data
}

func (o *ListAccountsLiquidationsPortfoliosOrdersListResponse) GetMeta() ListAccountsLiquidationsMeta {
	if o == nil {
		return ListAccountsLiquidationsMeta{}
	}
	return o.Meta
}

type ListAccountsLiquidationsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// Accounts liquidations
	PortfoliosOrdersListResponse *ListAccountsLiquidationsPortfoliosOrdersListResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListAccountsLiquidationsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAccountsLiquidationsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *ListAccountsLiquidationsResponse) GetPortfoliosOrdersListResponse() *ListAccountsLiquidationsPortfoliosOrdersListResponse {
	if o == nil {
		return nil
	}
	return o.PortfoliosOrdersListResponse
}

func (o *ListAccountsLiquidationsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAccountsLiquidationsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
