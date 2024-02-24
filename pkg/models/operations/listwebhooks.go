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

type ListWebhooksSecurity struct {
	OauthClientCredentials string `security:"scheme,type=oauth2,name=Authorization"`
}

func (o *ListWebhooksSecurity) GetOauthClientCredentials() string {
	if o == nil {
		return ""
	}
	return o.OauthClientCredentials
}

// ListWebhooksQueryParamOrder - Sort order of the result list if the `sort` parameter is specified. Use `ASC` for ascending or `DESC` for descending sort order.
type ListWebhooksQueryParamOrder string

const (
	ListWebhooksQueryParamOrderAsc  ListWebhooksQueryParamOrder = "ASC"
	ListWebhooksQueryParamOrderDesc ListWebhooksQueryParamOrder = "DESC"
)

func (e ListWebhooksQueryParamOrder) ToPointer() *ListWebhooksQueryParamOrder {
	return &e
}

func (e *ListWebhooksQueryParamOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = ListWebhooksQueryParamOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListWebhooksQueryParamOrder: %v", v)
	}
}

// ListWebhooksQueryParamSort - Sort the result by `created_at`, `updated_at`, `title`, `url`, or `enabled`.
type ListWebhooksQueryParamSort string

const (
	ListWebhooksQueryParamSortCreatedAt ListWebhooksQueryParamSort = "created_at"
	ListWebhooksQueryParamSortUpdatedAt ListWebhooksQueryParamSort = "updated_at"
	ListWebhooksQueryParamSortTitle     ListWebhooksQueryParamSort = "title"
	ListWebhooksQueryParamSortURL       ListWebhooksQueryParamSort = "url"
	ListWebhooksQueryParamSortEnabled   ListWebhooksQueryParamSort = "enabled"
)

func (e ListWebhooksQueryParamSort) ToPointer() *ListWebhooksQueryParamSort {
	return &e
}

func (e *ListWebhooksQueryParamSort) UnmarshalJSON(data []byte) error {
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
		*e = ListWebhooksQueryParamSort(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListWebhooksQueryParamSort: %v", v)
	}
}

type ListWebhooksRequest struct {
	// Use the `limit` argument to specify the maximum number of items returned.
	Limit *int `default:"100" queryParam:"style=form,explode=true,name=limit"`
	// Use the `offset` argument to specify where in the list of results to start when returning items for a particular query.
	Offset *int `queryParam:"style=form,explode=true,name=offset"`
	// Sort order of the result list if the `sort` parameter is specified. Use `ASC` for ascending or `DESC` for descending sort order.
	Order *ListWebhooksQueryParamOrder `default:"ASC" queryParam:"style=form,explode=true,name=order"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-http-header
	Signature string `header:"style=simple,explode=false,name=signature"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-input-http-he
	SignatureInput string `header:"style=simple,explode=false,name=signature-input"`
	// Sort the result by `created_at`, `updated_at`, `title`, `url`, or `enabled`.
	Sort *ListWebhooksQueryParamSort `default:"created_at" queryParam:"style=form,explode=true,name=sort"`
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

func (o *ListWebhooksRequest) GetOrder() *ListWebhooksQueryParamOrder {
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

func (o *ListWebhooksRequest) GetSort() *ListWebhooksQueryParamSort {
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

// ListWebhooksConfig - Configuration of webhook packages collection.
type ListWebhooksConfig struct {
	// Maximum time of package collection.
	Delay *string `json:"delay,omitempty"`
	// Maximum package size (bytes)
	MaxPackageSize *int64 `json:"max_package_size,omitempty"`
}

func (o *ListWebhooksConfig) GetDelay() *string {
	if o == nil {
		return nil
	}
	return o.Delay
}

func (o *ListWebhooksConfig) GetMaxPackageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxPackageSize
}

type ListWebhooksType string

const (
	ListWebhooksTypeAll                      ListWebhooksType = "ALL"
	ListWebhooksTypeUser                     ListWebhooksType = "USER"
	ListWebhooksTypeUserCheck                ListWebhooksType = "USER_CHECK"
	ListWebhooksTypeOrder                    ListWebhooksType = "ORDER"
	ListWebhooksTypeOrderCancellation        ListWebhooksType = "ORDER_CANCELLATION"
	ListWebhooksTypeExecution                ListWebhooksType = "EXECUTION"
	ListWebhooksTypePosition                 ListWebhooksType = "POSITION"
	ListWebhooksTypeCashBalance              ListWebhooksType = "CASH_BALANCE"
	ListWebhooksTypeAccount                  ListWebhooksType = "ACCOUNT"
	ListWebhooksTypeAccountGroup             ListWebhooksType = "ACCOUNT_GROUP"
	ListWebhooksTypeReport                   ListWebhooksType = "REPORT"
	ListWebhooksTypeTreasuryReport           ListWebhooksType = "TREASURY_REPORT"
	ListWebhooksTypeDirectDebit              ListWebhooksType = "DIRECT_DEBIT"
	ListWebhooksTypeWithdrawal               ListWebhooksType = "WITHDRAWAL"
	ListWebhooksTypePortfolio                ListWebhooksType = "PORTFOLIO"
	ListWebhooksTypePortfolioAllocation      ListWebhooksType = "PORTFOLIO_ALLOCATION"
	ListWebhooksTypePortfolioOrder           ListWebhooksType = "PORTFOLIO_ORDER"
	ListWebhooksTypeCorporateAction          ListWebhooksType = "CORPORATE_ACTION"
	ListWebhooksTypeAccountValuation         ListWebhooksType = "ACCOUNT_VALUATION"
	ListWebhooksTypeIntradayAccountValuation ListWebhooksType = "INTRADAY_ACCOUNT_VALUATION"
	ListWebhooksTypeCashTransaction          ListWebhooksType = "CASH_TRANSACTION"
	ListWebhooksTypeSecurityTransaction      ListWebhooksType = "SECURITY_TRANSACTION"
	ListWebhooksTypeAccountLiquidation       ListWebhooksType = "ACCOUNT_LIQUIDATION"
	ListWebhooksTypeAccountReturns           ListWebhooksType = "ACCOUNT_RETURNS"
	ListWebhooksTypeVirtualCashIncrease      ListWebhooksType = "VIRTUAL_CASH_INCREASE"
	ListWebhooksTypeVirtualCashDecrease      ListWebhooksType = "VIRTUAL_CASH_DECREASE"
	ListWebhooksTypeFeeCollection            ListWebhooksType = "FEE_COLLECTION"
)

func (e ListWebhooksType) ToPointer() *ListWebhooksType {
	return &e
}

func (e *ListWebhooksType) UnmarshalJSON(data []byte) error {
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
		*e = ListWebhooksType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListWebhooksType: %v", v)
	}
}

type Webhook struct {
	// Configuration of webhook packages collection.
	Config ListWebhooksConfig `json:"config"`
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	// Enable/disable webhook.
	Enabled bool `json:"enabled"`
	// Webhook unique identifier.
	ID string `json:"id"`
	// Title of the webhook for use on tenant side.
	Title string `json:"title"`
	// What kind of events to be sent by the webhook.
	Type []ListWebhooksType `json:"type"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time `json:"updated_at"`
	// The callback URL to be called by the webhook.
	URL string `json:"url"`
}

func (w Webhook) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *Webhook) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Webhook) GetConfig() ListWebhooksConfig {
	if o == nil {
		return ListWebhooksConfig{}
	}
	return o.Config
}

func (o *Webhook) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *Webhook) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *Webhook) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Webhook) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *Webhook) GetType() []ListWebhooksType {
	if o == nil {
		return []ListWebhooksType{}
	}
	return o.Type
}

func (o *Webhook) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *Webhook) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

// ListWebhooksOrder - The ordering of the response.
// * ASC - Ascending order
// * DESC - Descending order
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

type ListWebhooksMeta struct {
	// Count of the resources returned in the response.
	Count int64 `json:"count"`
	// Total limit of the response.
	Limit int64 `json:"limit"`
	// Amount of resource to offset in the response.
	Offset int64 `json:"offset"`
	// The ordering of the response.
	// * ASC - Ascending order
	// * DESC - Descending order
	Order *ListWebhooksOrder `json:"order,omitempty"`
	// The field that the list is sorted by.
	Sort *string `json:"sort,omitempty"`
	// Total count of all the resources.
	TotalCount int64 `json:"total_count"`
}

func (o *ListWebhooksMeta) GetCount() int64 {
	if o == nil {
		return 0
	}
	return o.Count
}

func (o *ListWebhooksMeta) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

func (o *ListWebhooksMeta) GetOffset() int64 {
	if o == nil {
		return 0
	}
	return o.Offset
}

func (o *ListWebhooksMeta) GetOrder() *ListWebhooksOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListWebhooksMeta) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListWebhooksMeta) GetTotalCount() int64 {
	if o == nil {
		return 0
	}
	return o.TotalCount
}

// ListWebhooksWebhooksListResponse - An object with a data property that contains an array of webhook subscription objects.
type ListWebhooksWebhooksListResponse struct {
	Data []Webhook        `json:"data"`
	Meta ListWebhooksMeta `json:"meta"`
}

func (o *ListWebhooksWebhooksListResponse) GetData() []Webhook {
	if o == nil {
		return []Webhook{}
	}
	return o.Data
}

func (o *ListWebhooksWebhooksListResponse) GetMeta() ListWebhooksMeta {
	if o == nil {
		return ListWebhooksMeta{}
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
		return map[string][]string{}
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
