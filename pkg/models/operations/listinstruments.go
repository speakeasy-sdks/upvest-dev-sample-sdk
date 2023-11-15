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

// Order - Sort order of the result list if the `sort` parameter is specified. Use `ASC` for ascending or `DESC` for descending sort order.
type Order string

const (
	OrderAsc  Order = "ASC"
	OrderDesc Order = "DESC"
)

func (e Order) ToPointer() *Order {
	return &e
}

func (e *Order) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = Order(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Order: %v", v)
	}
}

// ListInstrumentsQueryParamSort - Sort the result by `created_at` or `updated_at`.
type ListInstrumentsQueryParamSort string

const (
	ListInstrumentsQueryParamSortCreatedAt ListInstrumentsQueryParamSort = "created_at"
	ListInstrumentsQueryParamSortUpdatedAt ListInstrumentsQueryParamSort = "updated_at"
)

func (e ListInstrumentsQueryParamSort) ToPointer() *ListInstrumentsQueryParamSort {
	return &e
}

func (e *ListInstrumentsQueryParamSort) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		fallthrough
	case "updated_at":
		*e = ListInstrumentsQueryParamSort(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListInstrumentsQueryParamSort: %v", v)
	}
}

// InstrumentTradingStatus - Instrument trading status
// * ACTIVE - The instrument can currently be traded on the Upvest platform.
// * INACTIVE - The instrument cannot currently be traded on the Upvest platform.
type InstrumentTradingStatus string

const (
	InstrumentTradingStatusActive   InstrumentTradingStatus = "ACTIVE"
	InstrumentTradingStatusInactive InstrumentTradingStatus = "INACTIVE"
)

func (e InstrumentTradingStatus) ToPointer() *InstrumentTradingStatus {
	return &e
}

func (e *InstrumentTradingStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ACTIVE":
		fallthrough
	case "INACTIVE":
		*e = InstrumentTradingStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InstrumentTradingStatus: %v", v)
	}
}

type ListInstrumentsRequest struct {
	// Use the `limit` argument to specify the maximum number of items returned.
	Limit *int `default:"100" queryParam:"style=form,explode=true,name=limit"`
	// Use the `offset` argument to specify where in the list of results to start when returning items for a particular query.
	Offset *int `queryParam:"style=form,explode=true,name=offset"`
	// Sort order of the result list if the `sort` parameter is specified. Use `ASC` for ascending or `DESC` for descending sort order.
	Order *Order `default:"ASC" queryParam:"style=form,explode=true,name=order"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-http-header
	Signature string `header:"style=simple,explode=false,name=signature"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-input-http-he
	SignatureInput string `header:"style=simple,explode=false,name=signature-input"`
	// Sort the result by `created_at` or `updated_at`.
	Sort *ListInstrumentsQueryParamSort `default:"created_at" queryParam:"style=form,explode=true,name=sort"`
	// Filters the list to only show instruments with a certain status (e.g. only instruments that can be currently traded).
	TradingStatus *InstrumentTradingStatus `queryParam:"style=form,explode=true,name=trading_status"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
}

func (l ListInstrumentsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListInstrumentsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListInstrumentsRequest) GetLimit() *int {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListInstrumentsRequest) GetOffset() *int {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListInstrumentsRequest) GetOrder() *Order {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListInstrumentsRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *ListInstrumentsRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *ListInstrumentsRequest) GetSort() *ListInstrumentsQueryParamSort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListInstrumentsRequest) GetTradingStatus() *InstrumentTradingStatus {
	if o == nil {
		return nil
	}
	return o.TradingStatus
}

func (o *ListInstrumentsRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *ListInstrumentsRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// ListInstrumentsInstrumentTradingStatus - Instrument trading status
// * ACTIVE - The instrument can currently be traded on the Upvest platform.
// * INACTIVE - The instrument cannot currently be traded on the Upvest platform.
type ListInstrumentsInstrumentTradingStatus string

const (
	ListInstrumentsInstrumentTradingStatusActive   ListInstrumentsInstrumentTradingStatus = "ACTIVE"
	ListInstrumentsInstrumentTradingStatusInactive ListInstrumentsInstrumentTradingStatus = "INACTIVE"
)

func (e ListInstrumentsInstrumentTradingStatus) ToPointer() *ListInstrumentsInstrumentTradingStatus {
	return &e
}

func (e *ListInstrumentsInstrumentTradingStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ACTIVE":
		fallthrough
	case "INACTIVE":
		*e = ListInstrumentsInstrumentTradingStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListInstrumentsInstrumentTradingStatus: %v", v)
	}
}

type Data struct {
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	// Determines whether the platform can handle fractional investments within this instrument.
	FractionalTrading bool `json:"fractional_trading"`
	// Instrument unique identifier.
	ID string `json:"id"`
	// International securities identification number defined by [ISO 6166](https://en.wikipedia.org/wiki/International_Securities_Identification_Number).
	Isin *string `json:"isin,omitempty"`
	// Instrument name
	Name string `json:"name"`
	// Instrument trading status
	// * ACTIVE - The instrument can currently be traded on the Upvest platform.
	// * INACTIVE - The instrument cannot currently be traded on the Upvest platform.
	TradingStatus ListInstrumentsInstrumentTradingStatus `json:"trading_status"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time `json:"updated_at"`
	// German securities identification code known as [Wertpapierkennnummer](https://en.wikipedia.org/wiki/Wertpapierkennnummer).
	Wkn *string `json:"wkn,omitempty"`
}

func (d Data) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *Data) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Data) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *Data) GetFractionalTrading() bool {
	if o == nil {
		return false
	}
	return o.FractionalTrading
}

func (o *Data) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Data) GetIsin() *string {
	if o == nil {
		return nil
	}
	return o.Isin
}

func (o *Data) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Data) GetTradingStatus() ListInstrumentsInstrumentTradingStatus {
	if o == nil {
		return ListInstrumentsInstrumentTradingStatus("")
	}
	return o.TradingStatus
}

func (o *Data) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *Data) GetWkn() *string {
	if o == nil {
		return nil
	}
	return o.Wkn
}

// ListInstrumentsOrder - The ordering of the response.
// * ASC - Ascending order
// * DESC - Descending order
type ListInstrumentsOrder string

const (
	ListInstrumentsOrderAsc  ListInstrumentsOrder = "ASC"
	ListInstrumentsOrderDesc ListInstrumentsOrder = "DESC"
)

func (e ListInstrumentsOrder) ToPointer() *ListInstrumentsOrder {
	return &e
}

func (e *ListInstrumentsOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = ListInstrumentsOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListInstrumentsOrder: %v", v)
	}
}

type ListInstrumentsMeta struct {
	// Count of the resources returned in the response.
	Count int64 `json:"count"`
	// Total limit of the response.
	Limit int64 `json:"limit"`
	// Amount of resource to offset in the response.
	Offset int64 `json:"offset"`
	// The ordering of the response.
	// * ASC - Ascending order
	// * DESC - Descending order
	Order *ListInstrumentsOrder `json:"order,omitempty"`
	// The field that the list is sorted by.
	Sort *string `json:"sort,omitempty"`
	// Total count of all the resources.
	TotalCount int64 `json:"total_count"`
}

func (o *ListInstrumentsMeta) GetCount() int64 {
	if o == nil {
		return 0
	}
	return o.Count
}

func (o *ListInstrumentsMeta) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

func (o *ListInstrumentsMeta) GetOffset() int64 {
	if o == nil {
		return 0
	}
	return o.Offset
}

func (o *ListInstrumentsMeta) GetOrder() *ListInstrumentsOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListInstrumentsMeta) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListInstrumentsMeta) GetTotalCount() int64 {
	if o == nil {
		return 0
	}
	return o.TotalCount
}

// ListInstrumentsInstrumentsListResponse - Instruments list
type ListInstrumentsInstrumentsListResponse struct {
	Data []Data              `json:"data"`
	Meta ListInstrumentsMeta `json:"meta"`
}

func (o *ListInstrumentsInstrumentsListResponse) GetData() []Data {
	if o == nil {
		return []Data{}
	}
	return o.Data
}

func (o *ListInstrumentsInstrumentsListResponse) GetMeta() ListInstrumentsMeta {
	if o == nil {
		return ListInstrumentsMeta{}
	}
	return o.Meta
}

type ListInstrumentsResponse struct {
	// Instruments list
	TwoHundredApplicationJSONInstrumentsListResponse *ListInstrumentsInstrumentsListResponse
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListInstrumentsResponse) GetTwoHundredApplicationJSONInstrumentsListResponse() *ListInstrumentsInstrumentsListResponse {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONInstrumentsListResponse
}

func (o *ListInstrumentsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListInstrumentsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *ListInstrumentsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListInstrumentsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
