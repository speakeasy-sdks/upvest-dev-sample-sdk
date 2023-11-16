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

// ListPortfolioRebalancingExecutionOrdersQueryParamSort - Sort the result by `status`.
type ListPortfolioRebalancingExecutionOrdersQueryParamSort string

const (
	ListPortfolioRebalancingExecutionOrdersQueryParamSortStatus ListPortfolioRebalancingExecutionOrdersQueryParamSort = "status"
)

func (e ListPortfolioRebalancingExecutionOrdersQueryParamSort) ToPointer() *ListPortfolioRebalancingExecutionOrdersQueryParamSort {
	return &e
}

func (e *ListPortfolioRebalancingExecutionOrdersQueryParamSort) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "status":
		*e = ListPortfolioRebalancingExecutionOrdersQueryParamSort(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListPortfolioRebalancingExecutionOrdersQueryParamSort: %v", v)
	}
}

type ListPortfolioRebalancingExecutionOrdersRequest struct {
	// Returns rebalancing orders with dates up until this date (UTC)
	EndDate     *string `queryParam:"style=form,explode=true,name=end_date"`
	ExecutionID string  `pathParam:"style=simple,explode=false,name=execution_id"`
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
	// Sort the result by `status`.
	Sort *ListPortfolioRebalancingExecutionOrdersQueryParamSort `default:"status" queryParam:"style=form,explode=true,name=sort"`
	// Returns rebalancing orders with dates starting from and including this date (UTC)
	StartDate *string `queryParam:"style=form,explode=true,name=start_date"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
}

func (l ListPortfolioRebalancingExecutionOrdersRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListPortfolioRebalancingExecutionOrdersRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListPortfolioRebalancingExecutionOrdersRequest) GetEndDate() *string {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *ListPortfolioRebalancingExecutionOrdersRequest) GetExecutionID() string {
	if o == nil {
		return ""
	}
	return o.ExecutionID
}

func (o *ListPortfolioRebalancingExecutionOrdersRequest) GetLimit() *int {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListPortfolioRebalancingExecutionOrdersRequest) GetOffset() *int {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListPortfolioRebalancingExecutionOrdersRequest) GetOrder() *shared.Order {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListPortfolioRebalancingExecutionOrdersRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *ListPortfolioRebalancingExecutionOrdersRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *ListPortfolioRebalancingExecutionOrdersRequest) GetSort() *ListPortfolioRebalancingExecutionOrdersQueryParamSort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListPortfolioRebalancingExecutionOrdersRequest) GetStartDate() *string {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *ListPortfolioRebalancingExecutionOrdersRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *ListPortfolioRebalancingExecutionOrdersRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// ListPortfolioRebalancingExecutionOrdersStatus - Execution status of the Portfolio Order.
// * NEW -
// * PROCESSING -
// * FILLED -
// * SETTLED -
// * CANCELLED -
type ListPortfolioRebalancingExecutionOrdersStatus string

const (
	ListPortfolioRebalancingExecutionOrdersStatusNew        ListPortfolioRebalancingExecutionOrdersStatus = "NEW"
	ListPortfolioRebalancingExecutionOrdersStatusProcessing ListPortfolioRebalancingExecutionOrdersStatus = "PROCESSING"
	ListPortfolioRebalancingExecutionOrdersStatusFilled     ListPortfolioRebalancingExecutionOrdersStatus = "FILLED"
	ListPortfolioRebalancingExecutionOrdersStatusSettled    ListPortfolioRebalancingExecutionOrdersStatus = "SETTLED"
	ListPortfolioRebalancingExecutionOrdersStatusCancelled  ListPortfolioRebalancingExecutionOrdersStatus = "CANCELLED"
)

func (e ListPortfolioRebalancingExecutionOrdersStatus) ToPointer() *ListPortfolioRebalancingExecutionOrdersStatus {
	return &e
}

func (e *ListPortfolioRebalancingExecutionOrdersStatus) UnmarshalJSON(data []byte) error {
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
		*e = ListPortfolioRebalancingExecutionOrdersStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListPortfolioRebalancingExecutionOrdersStatus: %v", v)
	}
}

type PortfoliosRebalancingExecutionOrder struct {
	// Account unique identifier.
	AccountID string `json:"account_id"`
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt        time.Time `json:"created_at"`
	PortfolioOrderID string    `json:"portfolio_order_id"`
	// Execution status of the Portfolio Order.
	// * NEW -
	// * PROCESSING -
	// * FILLED -
	// * SETTLED -
	// * CANCELLED -
	Status ListPortfolioRebalancingExecutionOrdersStatus `json:"status"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time `json:"updated_at"`
}

func (p PortfoliosRebalancingExecutionOrder) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PortfoliosRebalancingExecutionOrder) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PortfoliosRebalancingExecutionOrder) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *PortfoliosRebalancingExecutionOrder) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *PortfoliosRebalancingExecutionOrder) GetPortfolioOrderID() string {
	if o == nil {
		return ""
	}
	return o.PortfolioOrderID
}

func (o *PortfoliosRebalancingExecutionOrder) GetStatus() ListPortfolioRebalancingExecutionOrdersStatus {
	if o == nil {
		return ListPortfolioRebalancingExecutionOrdersStatus("")
	}
	return o.Status
}

func (o *PortfoliosRebalancingExecutionOrder) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

// ListPortfolioRebalancingExecutionOrdersOrder - The ordering of the response.
// * ASC - Ascending order
// * DESC - Descending order
type ListPortfolioRebalancingExecutionOrdersOrder string

const (
	ListPortfolioRebalancingExecutionOrdersOrderAsc  ListPortfolioRebalancingExecutionOrdersOrder = "ASC"
	ListPortfolioRebalancingExecutionOrdersOrderDesc ListPortfolioRebalancingExecutionOrdersOrder = "DESC"
)

func (e ListPortfolioRebalancingExecutionOrdersOrder) ToPointer() *ListPortfolioRebalancingExecutionOrdersOrder {
	return &e
}

func (e *ListPortfolioRebalancingExecutionOrdersOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = ListPortfolioRebalancingExecutionOrdersOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListPortfolioRebalancingExecutionOrdersOrder: %v", v)
	}
}

type ListPortfolioRebalancingExecutionOrdersMeta struct {
	// Count of the resources returned in the response.
	Count int64 `json:"count"`
	// Total limit of the response.
	Limit int64 `json:"limit"`
	// Amount of resource to offset in the response.
	Offset int64 `json:"offset"`
	// The ordering of the response.
	// * ASC - Ascending order
	// * DESC - Descending order
	Order *ListPortfolioRebalancingExecutionOrdersOrder `json:"order,omitempty"`
	// The field that the list is sorted by.
	Sort *string `json:"sort,omitempty"`
	// Total count of all the resources.
	TotalCount int64 `json:"total_count"`
}

func (o *ListPortfolioRebalancingExecutionOrdersMeta) GetCount() int64 {
	if o == nil {
		return 0
	}
	return o.Count
}

func (o *ListPortfolioRebalancingExecutionOrdersMeta) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

func (o *ListPortfolioRebalancingExecutionOrdersMeta) GetOffset() int64 {
	if o == nil {
		return 0
	}
	return o.Offset
}

func (o *ListPortfolioRebalancingExecutionOrdersMeta) GetOrder() *ListPortfolioRebalancingExecutionOrdersOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListPortfolioRebalancingExecutionOrdersMeta) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListPortfolioRebalancingExecutionOrdersMeta) GetTotalCount() int64 {
	if o == nil {
		return 0
	}
	return o.TotalCount
}

// ListPortfolioRebalancingExecutionOrdersPortfoliosRebalancingExecutionOrderListResponse - Portfolios
type ListPortfolioRebalancingExecutionOrdersPortfoliosRebalancingExecutionOrderListResponse struct {
	Data []PortfoliosRebalancingExecutionOrder       `json:"data"`
	Meta ListPortfolioRebalancingExecutionOrdersMeta `json:"meta"`
}

func (o *ListPortfolioRebalancingExecutionOrdersPortfoliosRebalancingExecutionOrderListResponse) GetData() []PortfoliosRebalancingExecutionOrder {
	if o == nil {
		return []PortfoliosRebalancingExecutionOrder{}
	}
	return o.Data
}

func (o *ListPortfolioRebalancingExecutionOrdersPortfoliosRebalancingExecutionOrderListResponse) GetMeta() ListPortfolioRebalancingExecutionOrdersMeta {
	if o == nil {
		return ListPortfolioRebalancingExecutionOrdersMeta{}
	}
	return o.Meta
}

type ListPortfolioRebalancingExecutionOrdersResponse struct {
	// Portfolios
	TwoHundredApplicationJSONPortfoliosRebalancingExecutionOrderListResponse *ListPortfolioRebalancingExecutionOrdersPortfoliosRebalancingExecutionOrderListResponse
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListPortfolioRebalancingExecutionOrdersResponse) GetTwoHundredApplicationJSONPortfoliosRebalancingExecutionOrderListResponse() *ListPortfolioRebalancingExecutionOrdersPortfoliosRebalancingExecutionOrderListResponse {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONPortfoliosRebalancingExecutionOrderListResponse
}

func (o *ListPortfolioRebalancingExecutionOrdersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListPortfolioRebalancingExecutionOrdersResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *ListPortfolioRebalancingExecutionOrdersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListPortfolioRebalancingExecutionOrdersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
