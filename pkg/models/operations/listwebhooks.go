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

// ListWebhooksOrder - Sort order of the result list if the `sort` parameter is specified. Use `ASC` for ascending or `DESC` for descending sort order.
type ListWebhooksOrder string

const (
	ListWebhooksOrderAsc  ListWebhooksOrder = "ASC"
	ListWebhooksOrderDesc ListWebhooksOrder = "DESC"
)

func (e ListWebhooksOrder) ToPointer() *ListWebhooksOrder {
	return &e
}

func (e *ListWebhooksOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = ListWebhooksOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListWebhooksOrder: %v", v)
	}
}

// ListWebhooksSort - Sort the result by `created_at`, `updated_at`, `title`, `url`, or `enabled`.
type ListWebhooksSort string

const (
	ListWebhooksSortCreatedAt ListWebhooksSort = "created_at"
	ListWebhooksSortUpdatedAt ListWebhooksSort = "updated_at"
	ListWebhooksSortTitle     ListWebhooksSort = "title"
	ListWebhooksSortURL       ListWebhooksSort = "url"
	ListWebhooksSortEnabled   ListWebhooksSort = "enabled"
)

func (e ListWebhooksSort) ToPointer() *ListWebhooksSort {
	return &e
}

func (e *ListWebhooksSort) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		fallthrough
	case "updated_at":
		fallthrough
	case "title":
		fallthrough
	case "url":
		fallthrough
	case "enabled":
		*e = ListWebhooksSort(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListWebhooksSort: %v", v)
	}
}

type ListWebhooksRequest struct {
	// Use the `limit` argument to specify the maximum number of items returned.
	Limit *int `default:"100" queryParam:"style=form,explode=true,name=limit"`
	// Use the `offset` argument to specify where in the list of results to start when returning items for a particular query.
	Offset *int `queryParam:"style=form,explode=true,name=offset"`
	// Sort order of the result list if the `sort` parameter is specified. Use `ASC` for ascending or `DESC` for descending sort order.
	Order *ListWebhooksOrder `default:"ASC" queryParam:"style=form,explode=true,name=order"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-http-header
	Signature string `header:"style=simple,explode=false,name=signature"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-input-http-he
	SignatureInput string `header:"style=simple,explode=false,name=signature-input"`
	// Sort the result by `created_at`, `updated_at`, `title`, `url`, or `enabled`.
	Sort *ListWebhooksSort `default:"created_at" queryParam:"style=form,explode=true,name=sort"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
}

func (l ListWebhooksRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListWebhooksRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListWebhooksRequest) GetLimit() *int {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListWebhooksRequest) GetOffset() *int {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListWebhooksRequest) GetOrder() *ListWebhooksOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListWebhooksRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *ListWebhooksRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *ListWebhooksRequest) GetSort() *ListWebhooksSort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListWebhooksRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *ListWebhooksRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// ListWebhooksWebhooksListResponseWebhookConfig - Configuration of webhook packages collection.
type ListWebhooksWebhooksListResponseWebhookConfig struct {
	// Maximum time of package collection.
	Delay *string `json:"delay,omitempty"`
	// Maximum package size (bytes)
	MaxPackageSize *int64 `json:"max_package_size,omitempty"`
}

func (o *ListWebhooksWebhooksListResponseWebhookConfig) GetDelay() *string {
	if o == nil {
		return nil
	}
	return o.Delay
}

func (o *ListWebhooksWebhooksListResponseWebhookConfig) GetMaxPackageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxPackageSize
}

type ListWebhooksWebhooksListResponseWebhookType string

const (
	ListWebhooksWebhooksListResponseWebhookTypeAll                      ListWebhooksWebhooksListResponseWebhookType = "ALL"
	ListWebhooksWebhooksListResponseWebhookTypeUser                     ListWebhooksWebhooksListResponseWebhookType = "USER"
	ListWebhooksWebhooksListResponseWebhookTypeUserCheck                ListWebhooksWebhooksListResponseWebhookType = "USER_CHECK"
	ListWebhooksWebhooksListResponseWebhookTypeOrder                    ListWebhooksWebhooksListResponseWebhookType = "ORDER"
	ListWebhooksWebhooksListResponseWebhookTypeOrderCancellation        ListWebhooksWebhooksListResponseWebhookType = "ORDER_CANCELLATION"
	ListWebhooksWebhooksListResponseWebhookTypeExecution                ListWebhooksWebhooksListResponseWebhookType = "EXECUTION"
	ListWebhooksWebhooksListResponseWebhookTypePosition                 ListWebhooksWebhooksListResponseWebhookType = "POSITION"
	ListWebhooksWebhooksListResponseWebhookTypeCashBalance              ListWebhooksWebhooksListResponseWebhookType = "CASH_BALANCE"
	ListWebhooksWebhooksListResponseWebhookTypeAccount                  ListWebhooksWebhooksListResponseWebhookType = "ACCOUNT"
	ListWebhooksWebhooksListResponseWebhookTypeAccountGroup             ListWebhooksWebhooksListResponseWebhookType = "ACCOUNT_GROUP"
	ListWebhooksWebhooksListResponseWebhookTypeReport                   ListWebhooksWebhooksListResponseWebhookType = "REPORT"
	ListWebhooksWebhooksListResponseWebhookTypeTreasuryReport           ListWebhooksWebhooksListResponseWebhookType = "TREASURY_REPORT"
	ListWebhooksWebhooksListResponseWebhookTypeDirectDebit              ListWebhooksWebhooksListResponseWebhookType = "DIRECT_DEBIT"
	ListWebhooksWebhooksListResponseWebhookTypeWithdrawal               ListWebhooksWebhooksListResponseWebhookType = "WITHDRAWAL"
	ListWebhooksWebhooksListResponseWebhookTypePortfolio                ListWebhooksWebhooksListResponseWebhookType = "PORTFOLIO"
	ListWebhooksWebhooksListResponseWebhookTypePortfolioAllocation      ListWebhooksWebhooksListResponseWebhookType = "PORTFOLIO_ALLOCATION"
	ListWebhooksWebhooksListResponseWebhookTypePortfolioOrder           ListWebhooksWebhooksListResponseWebhookType = "PORTFOLIO_ORDER"
	ListWebhooksWebhooksListResponseWebhookTypeCorporateAction          ListWebhooksWebhooksListResponseWebhookType = "CORPORATE_ACTION"
	ListWebhooksWebhooksListResponseWebhookTypeAccountValuation         ListWebhooksWebhooksListResponseWebhookType = "ACCOUNT_VALUATION"
	ListWebhooksWebhooksListResponseWebhookTypeIntradayAccountValuation ListWebhooksWebhooksListResponseWebhookType = "INTRADAY_ACCOUNT_VALUATION"
	ListWebhooksWebhooksListResponseWebhookTypeCashTransaction          ListWebhooksWebhooksListResponseWebhookType = "CASH_TRANSACTION"
	ListWebhooksWebhooksListResponseWebhookTypeSecurityTransaction      ListWebhooksWebhooksListResponseWebhookType = "SECURITY_TRANSACTION"
	ListWebhooksWebhooksListResponseWebhookTypeAccountLiquidation       ListWebhooksWebhooksListResponseWebhookType = "ACCOUNT_LIQUIDATION"
	ListWebhooksWebhooksListResponseWebhookTypeAccountReturns           ListWebhooksWebhooksListResponseWebhookType = "ACCOUNT_RETURNS"
	ListWebhooksWebhooksListResponseWebhookTypeVirtualCashIncrease      ListWebhooksWebhooksListResponseWebhookType = "VIRTUAL_CASH_INCREASE"
	ListWebhooksWebhooksListResponseWebhookTypeVirtualCashDecrease      ListWebhooksWebhooksListResponseWebhookType = "VIRTUAL_CASH_DECREASE"
	ListWebhooksWebhooksListResponseWebhookTypeFeeCollection            ListWebhooksWebhooksListResponseWebhookType = "FEE_COLLECTION"
)

func (e ListWebhooksWebhooksListResponseWebhookType) ToPointer() *ListWebhooksWebhooksListResponseWebhookType {
	return &e
}

func (e *ListWebhooksWebhooksListResponseWebhookType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ALL":
		fallthrough
	case "USER":
		fallthrough
	case "USER_CHECK":
		fallthrough
	case "ORDER":
		fallthrough
	case "ORDER_CANCELLATION":
		fallthrough
	case "EXECUTION":
		fallthrough
	case "POSITION":
		fallthrough
	case "CASH_BALANCE":
		fallthrough
	case "ACCOUNT":
		fallthrough
	case "ACCOUNT_GROUP":
		fallthrough
	case "REPORT":
		fallthrough
	case "TREASURY_REPORT":
		fallthrough
	case "DIRECT_DEBIT":
		fallthrough
	case "WITHDRAWAL":
		fallthrough
	case "PORTFOLIO":
		fallthrough
	case "PORTFOLIO_ALLOCATION":
		fallthrough
	case "PORTFOLIO_ORDER":
		fallthrough
	case "CORPORATE_ACTION":
		fallthrough
	case "ACCOUNT_VALUATION":
		fallthrough
	case "INTRADAY_ACCOUNT_VALUATION":
		fallthrough
	case "CASH_TRANSACTION":
		fallthrough
	case "SECURITY_TRANSACTION":
		fallthrough
	case "ACCOUNT_LIQUIDATION":
		fallthrough
	case "ACCOUNT_RETURNS":
		fallthrough
	case "VIRTUAL_CASH_INCREASE":
		fallthrough
	case "VIRTUAL_CASH_DECREASE":
		fallthrough
	case "FEE_COLLECTION":
		*e = ListWebhooksWebhooksListResponseWebhookType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListWebhooksWebhooksListResponseWebhookType: %v", v)
	}
}

type ListWebhooksWebhooksListResponseWebhook struct {
	// Configuration of webhook packages collection.
	Config ListWebhooksWebhooksListResponseWebhookConfig `json:"config"`
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	// Enable/disable webhook.
	Enabled bool `json:"enabled"`
	// Webhook unique identifier.
	ID string `json:"id"`
	// Title of the webhook for use on tenant side.
	Title string `json:"title"`
	// What kind of events to be sent by the webhook.
	Type []ListWebhooksWebhooksListResponseWebhookType `json:"type"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time `json:"updated_at"`
	// The callback URL to be called by the webhook.
	URL string `json:"url"`
}

func (l ListWebhooksWebhooksListResponseWebhook) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListWebhooksWebhooksListResponseWebhook) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListWebhooksWebhooksListResponseWebhook) GetConfig() ListWebhooksWebhooksListResponseWebhookConfig {
	if o == nil {
		return ListWebhooksWebhooksListResponseWebhookConfig{}
	}
	return o.Config
}

func (o *ListWebhooksWebhooksListResponseWebhook) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *ListWebhooksWebhooksListResponseWebhook) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *ListWebhooksWebhooksListResponseWebhook) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ListWebhooksWebhooksListResponseWebhook) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *ListWebhooksWebhooksListResponseWebhook) GetType() []ListWebhooksWebhooksListResponseWebhookType {
	if o == nil {
		return []ListWebhooksWebhooksListResponseWebhookType{}
	}
	return o.Type
}

func (o *ListWebhooksWebhooksListResponseWebhook) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *ListWebhooksWebhooksListResponseWebhook) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

// ListWebhooksWebhooksListResponseMetaOrder - The ordering of the response.
// * ASC - Ascending order
// * DESC - Descending order
type ListWebhooksWebhooksListResponseMetaOrder string

const (
	ListWebhooksWebhooksListResponseMetaOrderAsc  ListWebhooksWebhooksListResponseMetaOrder = "ASC"
	ListWebhooksWebhooksListResponseMetaOrderDesc ListWebhooksWebhooksListResponseMetaOrder = "DESC"
)

func (e ListWebhooksWebhooksListResponseMetaOrder) ToPointer() *ListWebhooksWebhooksListResponseMetaOrder {
	return &e
}

func (e *ListWebhooksWebhooksListResponseMetaOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = ListWebhooksWebhooksListResponseMetaOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListWebhooksWebhooksListResponseMetaOrder: %v", v)
	}
}

type ListWebhooksWebhooksListResponseMeta struct {
	// Count of the resources returned in the response.
	Count int64 `json:"count"`
	// Total limit of the response.
	Limit int64 `json:"limit"`
	// Amount of resource to offset in the response.
	Offset int64 `json:"offset"`
	// The ordering of the response.
	// * ASC - Ascending order
	// * DESC - Descending order
	Order *ListWebhooksWebhooksListResponseMetaOrder `json:"order,omitempty"`
	// The field that the list is sorted by.
	Sort *string `json:"sort,omitempty"`
	// Total count of all the resources.
	TotalCount int64 `json:"total_count"`
}

func (o *ListWebhooksWebhooksListResponseMeta) GetCount() int64 {
	if o == nil {
		return 0
	}
	return o.Count
}

func (o *ListWebhooksWebhooksListResponseMeta) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

func (o *ListWebhooksWebhooksListResponseMeta) GetOffset() int64 {
	if o == nil {
		return 0
	}
	return o.Offset
}

func (o *ListWebhooksWebhooksListResponseMeta) GetOrder() *ListWebhooksWebhooksListResponseMetaOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListWebhooksWebhooksListResponseMeta) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListWebhooksWebhooksListResponseMeta) GetTotalCount() int64 {
	if o == nil {
		return 0
	}
	return o.TotalCount
}

// ListWebhooksWebhooksListResponse - An object with a data property that contains an array of webhook subscription objects.
type ListWebhooksWebhooksListResponse struct {
	Data []ListWebhooksWebhooksListResponseWebhook `json:"data"`
	Meta ListWebhooksWebhooksListResponseMeta      `json:"meta"`
}

func (o *ListWebhooksWebhooksListResponse) GetData() []ListWebhooksWebhooksListResponseWebhook {
	if o == nil {
		return []ListWebhooksWebhooksListResponseWebhook{}
	}
	return o.Data
}

func (o *ListWebhooksWebhooksListResponse) GetMeta() ListWebhooksWebhooksListResponseMeta {
	if o == nil {
		return ListWebhooksWebhooksListResponseMeta{}
	}
	return o.Meta
}

type ListWebhooksResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// An object with a data property that contains an array of webhook subscription objects.
	WebhooksListResponse *ListWebhooksWebhooksListResponse
}

func (o *ListWebhooksResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListWebhooksResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *ListWebhooksResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListWebhooksResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListWebhooksResponse) GetWebhooksListResponse() *ListWebhooksWebhooksListResponse {
	if o == nil {
		return nil
	}
	return o.WebhooksListResponse
}
