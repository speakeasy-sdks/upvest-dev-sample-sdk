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

// ListPortfoliosAllocationsSort - Sort the result by `id`.
type ListPortfoliosAllocationsSort string

const (
	ListPortfoliosAllocationsSortID ListPortfoliosAllocationsSort = "id"
)

func (e ListPortfoliosAllocationsSort) ToPointer() *ListPortfoliosAllocationsSort {
	return &e
}

func (e *ListPortfoliosAllocationsSort) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "id":
		*e = ListPortfoliosAllocationsSort(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListPortfoliosAllocationsSort: %v", v)
	}
}

type ListPortfoliosAllocationsRequest struct {
	// Returns portfolio allocations with dates up until this date (UTC)
	EndDate *string `queryParam:"style=form,explode=true,name=end_date"`
	// Filters portfolio allocations containing the instruments ID's
	InstrumentIds *string `queryParam:"style=form,explode=true,name=instrument_ids"`
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
	Sort *ListPortfoliosAllocationsSort `default:"id" queryParam:"style=form,explode=true,name=sort"`
	// Returns portfolio allocations with dates starting from and including this date (UTC)
	StartDate *string `queryParam:"style=form,explode=true,name=start_date"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
}

func (l ListPortfoliosAllocationsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListPortfoliosAllocationsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListPortfoliosAllocationsRequest) GetEndDate() *string {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *ListPortfoliosAllocationsRequest) GetInstrumentIds() *string {
	if o == nil {
		return nil
	}
	return o.InstrumentIds
}

func (o *ListPortfoliosAllocationsRequest) GetLimit() *int {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListPortfoliosAllocationsRequest) GetOffset() *int {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListPortfoliosAllocationsRequest) GetOrder() *shared.Order {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListPortfoliosAllocationsRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *ListPortfoliosAllocationsRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *ListPortfoliosAllocationsRequest) GetSort() *ListPortfoliosAllocationsSort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListPortfoliosAllocationsRequest) GetStartDate() *string {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *ListPortfoliosAllocationsRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *ListPortfoliosAllocationsRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocationAllocationInstrumentIDType - The type of the ID used in the request.
// * ISIN - International Securities Identification Number
// * UPVEST - UPVEST's unique instrument identifier
type ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocationAllocationInstrumentIDType string

const (
	ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocationAllocationInstrumentIDTypeIsin   ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocationAllocationInstrumentIDType = "ISIN"
	ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocationAllocationInstrumentIDTypeUpvest ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocationAllocationInstrumentIDType = "UPVEST"
)

func (e ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocationAllocationInstrumentIDType) ToPointer() *ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocationAllocationInstrumentIDType {
	return &e
}

func (e *ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocationAllocationInstrumentIDType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ISIN":
		fallthrough
	case "UPVEST":
		*e = ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocationAllocationInstrumentIDType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocationAllocationInstrumentIDType: %v", v)
	}
}

type ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocationAllocation struct {
	InstrumentID string `json:"instrument_id"`
	// The type of the ID used in the request.
	// * ISIN - International Securities Identification Number
	// * UPVEST - UPVEST's unique instrument identifier
	InstrumentIDType *ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocationAllocationInstrumentIDType `default:"ISIN" json:"instrument_id_type"`
	// Instrument allocation weight
	Weight string `json:"weight"`
}

func (l ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocationAllocation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocationAllocation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocationAllocation) GetInstrumentID() string {
	if o == nil {
		return ""
	}
	return o.InstrumentID
}

func (o *ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocationAllocation) GetInstrumentIDType() *ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocationAllocationInstrumentIDType {
	if o == nil {
		return nil
	}
	return o.InstrumentIDType
}

func (o *ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocationAllocation) GetWeight() string {
	if o == nil {
		return ""
	}
	return o.Weight
}

type ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocation struct {
	// List of portfolios allocations
	Allocation []ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocationAllocation `json:"allocation"`
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	ID        string    `json:"id"`
	// Allocation name
	Name string `json:"name"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time `json:"updated_at"`
}

func (l ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocation) GetAllocation() []ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocationAllocation {
	if o == nil {
		return []ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocationAllocation{}
	}
	return o.Allocation
}

func (o *ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocation) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocation) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocation) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocation) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

// ListPortfoliosAllocationsPortfoliosAllocationsListResponseMetaOrder - The ordering of the response.
// * ASC - Ascending order
// * DESC - Descending order
type ListPortfoliosAllocationsPortfoliosAllocationsListResponseMetaOrder string

const (
	ListPortfoliosAllocationsPortfoliosAllocationsListResponseMetaOrderAsc  ListPortfoliosAllocationsPortfoliosAllocationsListResponseMetaOrder = "ASC"
	ListPortfoliosAllocationsPortfoliosAllocationsListResponseMetaOrderDesc ListPortfoliosAllocationsPortfoliosAllocationsListResponseMetaOrder = "DESC"
)

func (e ListPortfoliosAllocationsPortfoliosAllocationsListResponseMetaOrder) ToPointer() *ListPortfoliosAllocationsPortfoliosAllocationsListResponseMetaOrder {
	return &e
}

func (e *ListPortfoliosAllocationsPortfoliosAllocationsListResponseMetaOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = ListPortfoliosAllocationsPortfoliosAllocationsListResponseMetaOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListPortfoliosAllocationsPortfoliosAllocationsListResponseMetaOrder: %v", v)
	}
}

type ListPortfoliosAllocationsPortfoliosAllocationsListResponseMeta struct {
	// Count of the resources returned in the response.
	Count int64 `json:"count"`
	// Total limit of the response.
	Limit int64 `json:"limit"`
	// Amount of resource to offset in the response.
	Offset int64 `json:"offset"`
	// The ordering of the response.
	// * ASC - Ascending order
	// * DESC - Descending order
	Order *ListPortfoliosAllocationsPortfoliosAllocationsListResponseMetaOrder `json:"order,omitempty"`
	// The field that the list is sorted by.
	Sort *string `json:"sort,omitempty"`
	// Total count of all the resources.
	TotalCount int64 `json:"total_count"`
}

func (o *ListPortfoliosAllocationsPortfoliosAllocationsListResponseMeta) GetCount() int64 {
	if o == nil {
		return 0
	}
	return o.Count
}

func (o *ListPortfoliosAllocationsPortfoliosAllocationsListResponseMeta) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

func (o *ListPortfoliosAllocationsPortfoliosAllocationsListResponseMeta) GetOffset() int64 {
	if o == nil {
		return 0
	}
	return o.Offset
}

func (o *ListPortfoliosAllocationsPortfoliosAllocationsListResponseMeta) GetOrder() *ListPortfoliosAllocationsPortfoliosAllocationsListResponseMetaOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListPortfoliosAllocationsPortfoliosAllocationsListResponseMeta) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListPortfoliosAllocationsPortfoliosAllocationsListResponseMeta) GetTotalCount() int64 {
	if o == nil {
		return 0
	}
	return o.TotalCount
}

// ListPortfoliosAllocationsPortfoliosAllocationsListResponse - Portfolios allocations
type ListPortfoliosAllocationsPortfoliosAllocationsListResponse struct {
	Data []ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocation `json:"data"`
	Meta ListPortfoliosAllocationsPortfoliosAllocationsListResponseMeta                   `json:"meta"`
}

func (o *ListPortfoliosAllocationsPortfoliosAllocationsListResponse) GetData() []ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocation {
	if o == nil {
		return []ListPortfoliosAllocationsPortfoliosAllocationsListResponsePortfoliosAllocation{}
	}
	return o.Data
}

func (o *ListPortfoliosAllocationsPortfoliosAllocationsListResponse) GetMeta() ListPortfoliosAllocationsPortfoliosAllocationsListResponseMeta {
	if o == nil {
		return ListPortfoliosAllocationsPortfoliosAllocationsListResponseMeta{}
	}
	return o.Meta
}

type ListPortfoliosAllocationsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// Portfolios allocations
	PortfoliosAllocationsListResponse *ListPortfoliosAllocationsPortfoliosAllocationsListResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListPortfoliosAllocationsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListPortfoliosAllocationsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *ListPortfoliosAllocationsResponse) GetPortfoliosAllocationsListResponse() *ListPortfoliosAllocationsPortfoliosAllocationsListResponse {
	if o == nil {
		return nil
	}
	return o.PortfoliosAllocationsListResponse
}

func (o *ListPortfoliosAllocationsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListPortfoliosAllocationsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}