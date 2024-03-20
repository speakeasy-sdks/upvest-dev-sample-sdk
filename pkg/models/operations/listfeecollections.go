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

type ListFeeCollectionsRequest struct {
	AccountGroupID string `queryParam:"style=form,explode=true,name=account_group_id"`
	AccountID      string `queryParam:"style=form,explode=true,name=account_id"`
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
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
}

func (l ListFeeCollectionsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListFeeCollectionsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListFeeCollectionsRequest) GetAccountGroupID() string {
	if o == nil {
		return ""
	}
	return o.AccountGroupID
}

func (o *ListFeeCollectionsRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *ListFeeCollectionsRequest) GetLimit() *int {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListFeeCollectionsRequest) GetOffset() *int {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListFeeCollectionsRequest) GetOrder() *shared.Order {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListFeeCollectionsRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *ListFeeCollectionsRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *ListFeeCollectionsRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *ListFeeCollectionsRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// ListFeeCollectionsCurrency - Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
// * EUR - Euro
type ListFeeCollectionsCurrency string

const (
	ListFeeCollectionsCurrencyEur ListFeeCollectionsCurrency = "EUR"
)

func (e ListFeeCollectionsCurrency) ToPointer() *ListFeeCollectionsCurrency {
	return &e
}

func (e *ListFeeCollectionsCurrency) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EUR":
		*e = ListFeeCollectionsCurrency(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListFeeCollectionsCurrency: %v", v)
	}
}

type ListFeeCollectionsProcessedAmount struct {
	CashBalance *string `json:"cash_balance,omitempty"`
	SellToCover *string `json:"sell_to_cover,omitempty"`
}

func (o *ListFeeCollectionsProcessedAmount) GetCashBalance() *string {
	if o == nil {
		return nil
	}
	return o.CashBalance
}

func (o *ListFeeCollectionsProcessedAmount) GetSellToCover() *string {
	if o == nil {
		return nil
	}
	return o.SellToCover
}

// ListFeeCollectionsStatus - Status of the fee collection
// * PROCESSING - Fee collection is in progress.
// * FINALISED - Fees have been collected from the account and the funds has been transferred to the client.
// * CANCELLED - Fee collection has been cancelled.
type ListFeeCollectionsStatus string

const (
	ListFeeCollectionsStatusProcessing ListFeeCollectionsStatus = "PROCESSING"
	ListFeeCollectionsStatusFinalised  ListFeeCollectionsStatus = "FINALISED"
	ListFeeCollectionsStatusCancelled  ListFeeCollectionsStatus = "CANCELLED"
)

func (e ListFeeCollectionsStatus) ToPointer() *ListFeeCollectionsStatus {
	return &e
}

func (e *ListFeeCollectionsStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PROCESSING":
		fallthrough
	case "FINALISED":
		fallthrough
	case "CANCELLED":
		*e = ListFeeCollectionsStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListFeeCollectionsStatus: %v", v)
	}
}

// ListFeeCollectionsType - Type of the fee collection
// * SERVICE_FEE - Service fee intake in a pre-defined cadence (e.g. monthly)
// * SERVICE_FEE_LIQUIDATION - Service fee intake as a result of a Portfolio liquidation
type ListFeeCollectionsType string

const (
	ListFeeCollectionsTypeServiceFee            ListFeeCollectionsType = "SERVICE_FEE"
	ListFeeCollectionsTypeServiceFeeLiquidation ListFeeCollectionsType = "SERVICE_FEE_LIQUIDATION"
)

func (e ListFeeCollectionsType) ToPointer() *ListFeeCollectionsType {
	return &e
}

func (e *ListFeeCollectionsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SERVICE_FEE":
		fallthrough
	case "SERVICE_FEE_LIQUIDATION":
		*e = ListFeeCollectionsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListFeeCollectionsType: %v", v)
	}
}

type FeeCollection struct {
	// Account group unique identifier.
	AccountGroupID string `json:"account_group_id"`
	// Account unique identifier.
	AccountID        string `json:"account_id"`
	CollectionAmount string `json:"collection_amount"`
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	// Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
	// * EUR - Euro
	Currency *ListFeeCollectionsCurrency `default:"EUR" json:"currency"`
	// Fee collection unique identifier.
	ID string `json:"id"`
	// End date of the fee collection period in YYYY-MM-DD format. [RFC 3339, section 5.6](https://json-schema.org/draft/2020-12/json-schema-validation.html#RFC3339) RFC 3339
	PeriodEnd types.Date `json:"period_end"`
	// Start date of the fee collection period in YYYY-MM-DD format. [RFC 3339, section 5.6](https://json-schema.org/draft/2020-12/json-schema-validation.html#RFC3339) RFC 3339
	PeriodStart     types.Date                        `json:"period_start"`
	ProcessedAmount ListFeeCollectionsProcessedAmount `json:"processed_amount"`
	// Status of the fee collection
	// * PROCESSING - Fee collection is in progress.
	// * FINALISED - Fees have been collected from the account and the funds has been transferred to the client.
	// * CANCELLED - Fee collection has been cancelled.
	Status ListFeeCollectionsStatus `json:"status"`
	// Type of the fee collection
	// * SERVICE_FEE - Service fee intake in a pre-defined cadence (e.g. monthly)
	// * SERVICE_FEE_LIQUIDATION - Service fee intake as a result of a Portfolio liquidation
	Type ListFeeCollectionsType `json:"type"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time `json:"updated_at"`
}

func (f FeeCollection) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *FeeCollection) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *FeeCollection) GetAccountGroupID() string {
	if o == nil {
		return ""
	}
	return o.AccountGroupID
}

func (o *FeeCollection) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *FeeCollection) GetCollectionAmount() string {
	if o == nil {
		return ""
	}
	return o.CollectionAmount
}

func (o *FeeCollection) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *FeeCollection) GetCurrency() *ListFeeCollectionsCurrency {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *FeeCollection) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *FeeCollection) GetPeriodEnd() types.Date {
	if o == nil {
		return types.Date{}
	}
	return o.PeriodEnd
}

func (o *FeeCollection) GetPeriodStart() types.Date {
	if o == nil {
		return types.Date{}
	}
	return o.PeriodStart
}

func (o *FeeCollection) GetProcessedAmount() ListFeeCollectionsProcessedAmount {
	if o == nil {
		return ListFeeCollectionsProcessedAmount{}
	}
	return o.ProcessedAmount
}

func (o *FeeCollection) GetStatus() ListFeeCollectionsStatus {
	if o == nil {
		return ListFeeCollectionsStatus("")
	}
	return o.Status
}

func (o *FeeCollection) GetType() ListFeeCollectionsType {
	if o == nil {
		return ListFeeCollectionsType("")
	}
	return o.Type
}

func (o *FeeCollection) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

// ListFeeCollectionsOrder - The ordering of the response.
// * ASC - Ascending order
// * DESC - Descending order
type ListFeeCollectionsOrder string

const (
	ListFeeCollectionsOrderAsc  ListFeeCollectionsOrder = "ASC"
	ListFeeCollectionsOrderDesc ListFeeCollectionsOrder = "DESC"
)

func (e ListFeeCollectionsOrder) ToPointer() *ListFeeCollectionsOrder {
	return &e
}

func (e *ListFeeCollectionsOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = ListFeeCollectionsOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListFeeCollectionsOrder: %v", v)
	}
}

type ListFeeCollectionsMeta struct {
	// Count of the resources returned in the response.
	Count int64 `json:"count"`
	// Total limit of the response.
	Limit int64 `json:"limit"`
	// Amount of resource to offset in the response.
	Offset int64 `json:"offset"`
	// The ordering of the response.
	// * ASC - Ascending order
	// * DESC - Descending order
	Order *ListFeeCollectionsOrder `json:"order,omitempty"`
	// The field that the list is sorted by.
	Sort *string `json:"sort,omitempty"`
	// Total count of all the resources.
	TotalCount int64 `json:"total_count"`
}

func (o *ListFeeCollectionsMeta) GetCount() int64 {
	if o == nil {
		return 0
	}
	return o.Count
}

func (o *ListFeeCollectionsMeta) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

func (o *ListFeeCollectionsMeta) GetOffset() int64 {
	if o == nil {
		return 0
	}
	return o.Offset
}

func (o *ListFeeCollectionsMeta) GetOrder() *ListFeeCollectionsOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListFeeCollectionsMeta) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListFeeCollectionsMeta) GetTotalCount() int64 {
	if o == nil {
		return 0
	}
	return o.TotalCount
}

// ListFeeCollectionsFeeCollectionListResponse - OK
type ListFeeCollectionsFeeCollectionListResponse struct {
	Data []FeeCollection        `json:"data"`
	Meta ListFeeCollectionsMeta `json:"meta"`
}

func (o *ListFeeCollectionsFeeCollectionListResponse) GetData() []FeeCollection {
	if o == nil {
		return []FeeCollection{}
	}
	return o.Data
}

func (o *ListFeeCollectionsFeeCollectionListResponse) GetMeta() ListFeeCollectionsMeta {
	if o == nil {
		return ListFeeCollectionsMeta{}
	}
	return o.Meta
}

type ListFeeCollectionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	FeeCollectionListResponse *ListFeeCollectionsFeeCollectionListResponse
	Headers                   map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListFeeCollectionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListFeeCollectionsResponse) GetFeeCollectionListResponse() *ListFeeCollectionsFeeCollectionListResponse {
	if o == nil {
		return nil
	}
	return o.FeeCollectionListResponse
}

func (o *ListFeeCollectionsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *ListFeeCollectionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListFeeCollectionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
