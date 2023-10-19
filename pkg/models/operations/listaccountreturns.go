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

// ListAccountReturnsOrder - Sort order of the result list if the `sort` parameter is specified. Use `ASC` for ascending or `DESC` for descending sort order.
type ListAccountReturnsOrder string

const (
	ListAccountReturnsOrderAsc  ListAccountReturnsOrder = "ASC"
	ListAccountReturnsOrderDesc ListAccountReturnsOrder = "DESC"
)

func (e ListAccountReturnsOrder) ToPointer() *ListAccountReturnsOrder {
	return &e
}

func (e *ListAccountReturnsOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = ListAccountReturnsOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountReturnsOrder: %v", v)
	}
}

// ListAccountReturnsSort - Sort the result by `date`.
type ListAccountReturnsSort string

const (
	ListAccountReturnsSortDate ListAccountReturnsSort = "date"
)

func (e ListAccountReturnsSort) ToPointer() *ListAccountReturnsSort {
	return &e
}

func (e *ListAccountReturnsSort) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "date":
		*e = ListAccountReturnsSort(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountReturnsSort: %v", v)
	}
}

type ListAccountReturnsRequest struct {
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// Returns account returns up until this date (UTC)
	EndDate *string `queryParam:"style=form,explode=true,name=end_date"`
	// Use the `limit` argument to specify the maximum number of items returned.
	Limit *int `default:"100" queryParam:"style=form,explode=true,name=limit"`
	// Use the `offset` argument to specify where in the list of results to start when returning items for a particular query.
	Offset *int `queryParam:"style=form,explode=true,name=offset"`
	// Sort order of the result list if the `sort` parameter is specified. Use `ASC` for ascending or `DESC` for descending sort order.
	Order *ListAccountReturnsOrder `default:"ASC" queryParam:"style=form,explode=true,name=order"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-http-header
	Signature string `header:"style=simple,explode=false,name=signature"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-input-http-he
	SignatureInput string `header:"style=simple,explode=false,name=signature-input"`
	// Sort the result by `date`.
	Sort *ListAccountReturnsSort `default:"date" queryParam:"style=form,explode=true,name=sort"`
	// Returns account returns starting from and including this date (UTC)
	StartDate *string `queryParam:"style=form,explode=true,name=start_date"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
}

func (l ListAccountReturnsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountReturnsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAccountReturnsRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *ListAccountReturnsRequest) GetEndDate() *string {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *ListAccountReturnsRequest) GetLimit() *int {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAccountReturnsRequest) GetOffset() *int {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListAccountReturnsRequest) GetOrder() *ListAccountReturnsOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListAccountReturnsRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *ListAccountReturnsRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *ListAccountReturnsRequest) GetSort() *ListAccountReturnsSort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListAccountReturnsRequest) GetStartDate() *string {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *ListAccountReturnsRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *ListAccountReturnsRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

type ListAccountReturnsAccountReturnListResponseAccountReturnsTwr struct {
	Cumulative *string `json:"cumulative,omitempty"`
	// Date when the cumulative time-weighted returns calculation starts.
	CumulativeStartDate *time.Time `json:"cumulative_start_date,omitempty"`
	Daily               *string    `json:"daily,omitempty"`
}

func (l ListAccountReturnsAccountReturnListResponseAccountReturnsTwr) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountReturnsAccountReturnListResponseAccountReturnsTwr) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAccountReturnsAccountReturnListResponseAccountReturnsTwr) GetCumulative() *string {
	if o == nil {
		return nil
	}
	return o.Cumulative
}

func (o *ListAccountReturnsAccountReturnListResponseAccountReturnsTwr) GetCumulativeStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.CumulativeStartDate
}

func (o *ListAccountReturnsAccountReturnListResponseAccountReturnsTwr) GetDaily() *string {
	if o == nil {
		return nil
	}
	return o.Daily
}

type ListAccountReturnsAccountReturnListResponseAccountReturns struct {
	// Account unique identifier.
	AccountID string `json:"account_id"`
	// Date when returns were calculated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	Date time.Time                                                    `json:"date"`
	Twr  ListAccountReturnsAccountReturnListResponseAccountReturnsTwr `json:"twr"`
}

func (l ListAccountReturnsAccountReturnListResponseAccountReturns) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountReturnsAccountReturnListResponseAccountReturns) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAccountReturnsAccountReturnListResponseAccountReturns) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *ListAccountReturnsAccountReturnListResponseAccountReturns) GetDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Date
}

func (o *ListAccountReturnsAccountReturnListResponseAccountReturns) GetTwr() ListAccountReturnsAccountReturnListResponseAccountReturnsTwr {
	if o == nil {
		return ListAccountReturnsAccountReturnListResponseAccountReturnsTwr{}
	}
	return o.Twr
}

// ListAccountReturnsAccountReturnListResponseMetaOrder - The ordering of the response.
// * ASC - Ascending order
// * DESC - Descending order
type ListAccountReturnsAccountReturnListResponseMetaOrder string

const (
	ListAccountReturnsAccountReturnListResponseMetaOrderAsc  ListAccountReturnsAccountReturnListResponseMetaOrder = "ASC"
	ListAccountReturnsAccountReturnListResponseMetaOrderDesc ListAccountReturnsAccountReturnListResponseMetaOrder = "DESC"
)

func (e ListAccountReturnsAccountReturnListResponseMetaOrder) ToPointer() *ListAccountReturnsAccountReturnListResponseMetaOrder {
	return &e
}

func (e *ListAccountReturnsAccountReturnListResponseMetaOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = ListAccountReturnsAccountReturnListResponseMetaOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountReturnsAccountReturnListResponseMetaOrder: %v", v)
	}
}

type ListAccountReturnsAccountReturnListResponseMeta struct {
	// Count of the resources returned in the response.
	Count int64 `json:"count"`
	// Total limit of the response.
	Limit int64 `json:"limit"`
	// Amount of resource to offset in the response.
	Offset int64 `json:"offset"`
	// The ordering of the response.
	// * ASC - Ascending order
	// * DESC - Descending order
	Order *ListAccountReturnsAccountReturnListResponseMetaOrder `json:"order,omitempty"`
	// The field that the list is sorted by.
	Sort *string `json:"sort,omitempty"`
	// Total count of all the resources.
	TotalCount int64 `json:"total_count"`
}

func (o *ListAccountReturnsAccountReturnListResponseMeta) GetCount() int64 {
	if o == nil {
		return 0
	}
	return o.Count
}

func (o *ListAccountReturnsAccountReturnListResponseMeta) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

func (o *ListAccountReturnsAccountReturnListResponseMeta) GetOffset() int64 {
	if o == nil {
		return 0
	}
	return o.Offset
}

func (o *ListAccountReturnsAccountReturnListResponseMeta) GetOrder() *ListAccountReturnsAccountReturnListResponseMetaOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListAccountReturnsAccountReturnListResponseMeta) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListAccountReturnsAccountReturnListResponseMeta) GetTotalCount() int64 {
	if o == nil {
		return 0
	}
	return o.TotalCount
}

// ListAccountReturnsAccountReturnListResponse - OK
type ListAccountReturnsAccountReturnListResponse struct {
	Data []ListAccountReturnsAccountReturnListResponseAccountReturns `json:"data"`
	Meta ListAccountReturnsAccountReturnListResponseMeta             `json:"meta"`
}

func (o *ListAccountReturnsAccountReturnListResponse) GetData() []ListAccountReturnsAccountReturnListResponseAccountReturns {
	if o == nil {
		return []ListAccountReturnsAccountReturnListResponseAccountReturns{}
	}
	return o.Data
}

func (o *ListAccountReturnsAccountReturnListResponse) GetMeta() ListAccountReturnsAccountReturnListResponseMeta {
	if o == nil {
		return ListAccountReturnsAccountReturnListResponseMeta{}
	}
	return o.Meta
}

type ListAccountReturnsResponse struct {
	// OK
	AccountReturnListResponse *ListAccountReturnsAccountReturnListResponse
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListAccountReturnsResponse) GetAccountReturnListResponse() *ListAccountReturnsAccountReturnListResponse {
	if o == nil {
		return nil
	}
	return o.AccountReturnListResponse
}

func (o *ListAccountReturnsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAccountReturnsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *ListAccountReturnsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAccountReturnsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
