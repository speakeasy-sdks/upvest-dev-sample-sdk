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

// ListPortfoliosRebalancingStrategiesQueryParamSort - Sort the result by `id`.
type ListPortfoliosRebalancingStrategiesQueryParamSort string

const (
	ListPortfoliosRebalancingStrategiesQueryParamSortID ListPortfoliosRebalancingStrategiesQueryParamSort = "id"
)

func (e ListPortfoliosRebalancingStrategiesQueryParamSort) ToPointer() *ListPortfoliosRebalancingStrategiesQueryParamSort {
	return &e
}

func (e *ListPortfoliosRebalancingStrategiesQueryParamSort) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "id":
		*e = ListPortfoliosRebalancingStrategiesQueryParamSort(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListPortfoliosRebalancingStrategiesQueryParamSort: %v", v)
	}
}

type ListPortfoliosRebalancingStrategiesRequest struct {
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
	Sort *ListPortfoliosRebalancingStrategiesQueryParamSort `default:"id" queryParam:"style=form,explode=true,name=sort"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
}

func (l ListPortfoliosRebalancingStrategiesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListPortfoliosRebalancingStrategiesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListPortfoliosRebalancingStrategiesRequest) GetLimit() *int {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListPortfoliosRebalancingStrategiesRequest) GetOffset() *int {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListPortfoliosRebalancingStrategiesRequest) GetOrder() *shared.Order {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListPortfoliosRebalancingStrategiesRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *ListPortfoliosRebalancingStrategiesRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *ListPortfoliosRebalancingStrategiesRequest) GetSort() *ListPortfoliosRebalancingStrategiesQueryParamSort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListPortfoliosRebalancingStrategiesRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *ListPortfoliosRebalancingStrategiesRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// ListPortfoliosRebalancingStrategiesType - The type of the strategy used in the request.
// * DRIFT - Trigger by drift percentage
// * SCHEDULED - Trigger by scheduled date
type ListPortfoliosRebalancingStrategiesType string

const (
	ListPortfoliosRebalancingStrategiesTypeDrift     ListPortfoliosRebalancingStrategiesType = "DRIFT"
	ListPortfoliosRebalancingStrategiesTypeScheduled ListPortfoliosRebalancingStrategiesType = "SCHEDULED"
)

func (e ListPortfoliosRebalancingStrategiesType) ToPointer() *ListPortfoliosRebalancingStrategiesType {
	return &e
}

func (e *ListPortfoliosRebalancingStrategiesType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DRIFT":
		fallthrough
	case "SCHEDULED":
		*e = ListPortfoliosRebalancingStrategiesType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListPortfoliosRebalancingStrategiesType: %v", v)
	}
}

type ListPortfoliosRebalancingStrategiesConditions struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	// The type of the ID used in the request.
	// * ISIN - International Securities Identification Number
	// * UPVEST - UPVEST's unique instrument identifier
	Name string `json:"name"`
	// The type of the strategy used in the request.
	// * DRIFT - Trigger by drift percentage
	// * SCHEDULED - Trigger by scheduled date
	Type ListPortfoliosRebalancingStrategiesType `json:"type"`
}

func (l ListPortfoliosRebalancingStrategiesConditions) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListPortfoliosRebalancingStrategiesConditions) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListPortfoliosRebalancingStrategiesConditions) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *ListPortfoliosRebalancingStrategiesConditions) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ListPortfoliosRebalancingStrategiesConditions) GetType() ListPortfoliosRebalancingStrategiesType {
	if o == nil {
		return ListPortfoliosRebalancingStrategiesType("")
	}
	return o.Type
}

type PortfoliosRebalancingStrategy struct {
	// List of conditions
	Conditions []ListPortfoliosRebalancingStrategiesConditions `json:"conditions"`
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	ID        string    `json:"id"`
	// Strategy name
	Name string `json:"name"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time `json:"updated_at"`
}

func (p PortfoliosRebalancingStrategy) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PortfoliosRebalancingStrategy) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PortfoliosRebalancingStrategy) GetConditions() []ListPortfoliosRebalancingStrategiesConditions {
	if o == nil {
		return []ListPortfoliosRebalancingStrategiesConditions{}
	}
	return o.Conditions
}

func (o *PortfoliosRebalancingStrategy) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *PortfoliosRebalancingStrategy) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PortfoliosRebalancingStrategy) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PortfoliosRebalancingStrategy) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

// ListPortfoliosRebalancingStrategiesOrder - The ordering of the response.
// * ASC - Ascending order
// * DESC - Descending order
type ListPortfoliosRebalancingStrategiesOrder string

const (
	ListPortfoliosRebalancingStrategiesOrderAsc  ListPortfoliosRebalancingStrategiesOrder = "ASC"
	ListPortfoliosRebalancingStrategiesOrderDesc ListPortfoliosRebalancingStrategiesOrder = "DESC"
)

func (e ListPortfoliosRebalancingStrategiesOrder) ToPointer() *ListPortfoliosRebalancingStrategiesOrder {
	return &e
}

func (e *ListPortfoliosRebalancingStrategiesOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = ListPortfoliosRebalancingStrategiesOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListPortfoliosRebalancingStrategiesOrder: %v", v)
	}
}

type ListPortfoliosRebalancingStrategiesMeta struct {
	// Count of the resources returned in the response.
	Count int64 `json:"count"`
	// Total limit of the response.
	Limit int64 `json:"limit"`
	// Amount of resource to offset in the response.
	Offset int64 `json:"offset"`
	// The ordering of the response.
	// * ASC - Ascending order
	// * DESC - Descending order
	Order *ListPortfoliosRebalancingStrategiesOrder `json:"order,omitempty"`
	// The field that the list is sorted by.
	Sort *string `json:"sort,omitempty"`
	// Total count of all the resources.
	TotalCount int64 `json:"total_count"`
}

func (o *ListPortfoliosRebalancingStrategiesMeta) GetCount() int64 {
	if o == nil {
		return 0
	}
	return o.Count
}

func (o *ListPortfoliosRebalancingStrategiesMeta) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

func (o *ListPortfoliosRebalancingStrategiesMeta) GetOffset() int64 {
	if o == nil {
		return 0
	}
	return o.Offset
}

func (o *ListPortfoliosRebalancingStrategiesMeta) GetOrder() *ListPortfoliosRebalancingStrategiesOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListPortfoliosRebalancingStrategiesMeta) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListPortfoliosRebalancingStrategiesMeta) GetTotalCount() int64 {
	if o == nil {
		return 0
	}
	return o.TotalCount
}

// ListPortfoliosRebalancingStrategiesPortfoliosRebalancingStrategyListResponse - Portfolios
type ListPortfoliosRebalancingStrategiesPortfoliosRebalancingStrategyListResponse struct {
	Data []PortfoliosRebalancingStrategy         `json:"data"`
	Meta ListPortfoliosRebalancingStrategiesMeta `json:"meta"`
}

func (o *ListPortfoliosRebalancingStrategiesPortfoliosRebalancingStrategyListResponse) GetData() []PortfoliosRebalancingStrategy {
	if o == nil {
		return []PortfoliosRebalancingStrategy{}
	}
	return o.Data
}

func (o *ListPortfoliosRebalancingStrategiesPortfoliosRebalancingStrategyListResponse) GetMeta() ListPortfoliosRebalancingStrategiesMeta {
	if o == nil {
		return ListPortfoliosRebalancingStrategiesMeta{}
	}
	return o.Meta
}

type ListPortfoliosRebalancingStrategiesResponse struct {
	// Portfolios
	TwoHundredApplicationJSONPortfoliosRebalancingStrategyListResponse *ListPortfoliosRebalancingStrategiesPortfoliosRebalancingStrategyListResponse
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListPortfoliosRebalancingStrategiesResponse) GetTwoHundredApplicationJSONPortfoliosRebalancingStrategyListResponse() *ListPortfoliosRebalancingStrategiesPortfoliosRebalancingStrategyListResponse {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONPortfoliosRebalancingStrategyListResponse
}

func (o *ListPortfoliosRebalancingStrategiesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListPortfoliosRebalancingStrategiesResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *ListPortfoliosRebalancingStrategiesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListPortfoliosRebalancingStrategiesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
