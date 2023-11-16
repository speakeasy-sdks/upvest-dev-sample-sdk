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

// ListDirectDebitsQueryParamOrder - Sort order of the result list if the `sort` parameter is specified. Use `ASC` for ascending or `DESC` for descending sort order.
type ListDirectDebitsQueryParamOrder string

const (
	ListDirectDebitsQueryParamOrderAsc  ListDirectDebitsQueryParamOrder = "ASC"
	ListDirectDebitsQueryParamOrderDesc ListDirectDebitsQueryParamOrder = "DESC"
)

func (e ListDirectDebitsQueryParamOrder) ToPointer() *ListDirectDebitsQueryParamOrder {
	return &e
}

func (e *ListDirectDebitsQueryParamOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = ListDirectDebitsQueryParamOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListDirectDebitsQueryParamOrder: %v", v)
	}
}

// ListDirectDebitsQueryParamSort - Field of resource to sort by
type ListDirectDebitsQueryParamSort string

const (
	ListDirectDebitsQueryParamSortID        ListDirectDebitsQueryParamSort = "id"
	ListDirectDebitsQueryParamSortCreatedAt ListDirectDebitsQueryParamSort = "created_at"
)

func (e ListDirectDebitsQueryParamSort) ToPointer() *ListDirectDebitsQueryParamSort {
	return &e
}

func (e *ListDirectDebitsQueryParamSort) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "id":
		fallthrough
	case "created_at":
		*e = ListDirectDebitsQueryParamSort(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListDirectDebitsQueryParamSort: %v", v)
	}
}

type ListDirectDebitsRequest struct {
	AccountGroupID string `pathParam:"style=simple,explode=false,name=account_group_id"`
	// Use the `limit` argument to specify the maximum number of items returned.
	Limit *int `default:"100" queryParam:"style=form,explode=true,name=limit"`
	// Use the `offset` argument to specify where in the list of results to start when returning items for a particular query.
	Offset *int `queryParam:"style=form,explode=true,name=offset"`
	// Sort order of the result list if the `sort` parameter is specified. Use `ASC` for ascending or `DESC` for descending sort order.
	Order *ListDirectDebitsQueryParamOrder `default:"ASC" queryParam:"style=form,explode=true,name=order"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-http-header
	Signature string `header:"style=simple,explode=false,name=signature"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-input-http-he
	SignatureInput string `header:"style=simple,explode=false,name=signature-input"`
	// Field of resource to sort by
	Sort *ListDirectDebitsQueryParamSort `default:"created_at" queryParam:"style=form,explode=true,name=sort"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
}

func (l ListDirectDebitsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListDirectDebitsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListDirectDebitsRequest) GetAccountGroupID() string {
	if o == nil {
		return ""
	}
	return o.AccountGroupID
}

func (o *ListDirectDebitsRequest) GetLimit() *int {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListDirectDebitsRequest) GetOffset() *int {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListDirectDebitsRequest) GetOrder() *ListDirectDebitsQueryParamOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListDirectDebitsRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *ListDirectDebitsRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *ListDirectDebitsRequest) GetSort() *ListDirectDebitsQueryParamSort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListDirectDebitsRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *ListDirectDebitsRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// ListDirectDebitsCurrency - Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
// * EUR - Euro
type ListDirectDebitsCurrency string

const (
	ListDirectDebitsCurrencyEur ListDirectDebitsCurrency = "EUR"
)

func (e ListDirectDebitsCurrency) ToPointer() *ListDirectDebitsCurrency {
	return &e
}

func (e *ListDirectDebitsCurrency) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EUR":
		*e = ListDirectDebitsCurrency(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListDirectDebitsCurrency: %v", v)
	}
}

// ListDirectDebitsStatus - Status of the direct debit
// * NEW - Direct debit is created but not started processing.
// * PROCESSING - Direct debit is in processing.
// * CONFIRMED - Direct debit was successfully processed.
// * CANCELLED - Direct debit was cancelled.
type ListDirectDebitsStatus string

const (
	ListDirectDebitsStatusNew        ListDirectDebitsStatus = "NEW"
	ListDirectDebitsStatusProcessing ListDirectDebitsStatus = "PROCESSING"
	ListDirectDebitsStatusConfirmed  ListDirectDebitsStatus = "CONFIRMED"
	ListDirectDebitsStatusCancelled  ListDirectDebitsStatus = "CANCELLED"
)

func (e ListDirectDebitsStatus) ToPointer() *ListDirectDebitsStatus {
	return &e
}

func (e *ListDirectDebitsStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NEW":
		fallthrough
	case "PROCESSING":
		fallthrough
	case "CONFIRMED":
		fallthrough
	case "CANCELLED":
		*e = ListDirectDebitsStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListDirectDebitsStatus: %v", v)
	}
}

type ListDirectDebitsData struct {
	// Account group unique identifier.
	AccountGroupID string `json:"account_group_id"`
	CashAmount     string `json:"cash_amount"`
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	// Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
	// * EUR - Euro
	Currency *ListDirectDebitsCurrency `default:"EUR" json:"currency"`
	// Direct debit funding request unique identifier
	ID string `json:"id"`
	// Direct Debit Mandate unique identifier.
	MandateID string `json:"mandate_id"`
	// Payment reference the end user will see in their bank statement for the corresponding direct debit booking (“Verwendungszweck”)
	RemittanceInformation *string `json:"remittance_information,omitempty"`
	// Status of the direct debit
	// * NEW - Direct debit is created but not started processing.
	// * PROCESSING - Direct debit is in processing.
	// * CONFIRMED - Direct debit was successfully processed.
	// * CANCELLED - Direct debit was cancelled.
	Status *ListDirectDebitsStatus `json:"status,omitempty"`
	// User unique identifier.
	UserID string `json:"user_id"`
}

func (l ListDirectDebitsData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListDirectDebitsData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListDirectDebitsData) GetAccountGroupID() string {
	if o == nil {
		return ""
	}
	return o.AccountGroupID
}

func (o *ListDirectDebitsData) GetCashAmount() string {
	if o == nil {
		return ""
	}
	return o.CashAmount
}

func (o *ListDirectDebitsData) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *ListDirectDebitsData) GetCurrency() *ListDirectDebitsCurrency {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *ListDirectDebitsData) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ListDirectDebitsData) GetMandateID() string {
	if o == nil {
		return ""
	}
	return o.MandateID
}

func (o *ListDirectDebitsData) GetRemittanceInformation() *string {
	if o == nil {
		return nil
	}
	return o.RemittanceInformation
}

func (o *ListDirectDebitsData) GetStatus() *ListDirectDebitsStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *ListDirectDebitsData) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

// ListDirectDebitsOrder - The ordering of the response.
// * ASC - Ascending order
// * DESC - Descending order
type ListDirectDebitsOrder string

const (
	ListDirectDebitsOrderAsc  ListDirectDebitsOrder = "ASC"
	ListDirectDebitsOrderDesc ListDirectDebitsOrder = "DESC"
)

func (e ListDirectDebitsOrder) ToPointer() *ListDirectDebitsOrder {
	return &e
}

func (e *ListDirectDebitsOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = ListDirectDebitsOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListDirectDebitsOrder: %v", v)
	}
}

type ListDirectDebitsMeta struct {
	// Count of the resources returned in the response.
	Count int64 `json:"count"`
	// Total limit of the response.
	Limit int64 `json:"limit"`
	// Amount of resource to offset in the response.
	Offset int64 `json:"offset"`
	// The ordering of the response.
	// * ASC - Ascending order
	// * DESC - Descending order
	Order *ListDirectDebitsOrder `json:"order,omitempty"`
	// The field that the list is sorted by.
	Sort *string `json:"sort,omitempty"`
	// Total count of all the resources.
	TotalCount int64 `json:"total_count"`
}

func (o *ListDirectDebitsMeta) GetCount() int64 {
	if o == nil {
		return 0
	}
	return o.Count
}

func (o *ListDirectDebitsMeta) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

func (o *ListDirectDebitsMeta) GetOffset() int64 {
	if o == nil {
		return 0
	}
	return o.Offset
}

func (o *ListDirectDebitsMeta) GetOrder() *ListDirectDebitsOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListDirectDebitsMeta) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListDirectDebitsMeta) GetTotalCount() int64 {
	if o == nil {
		return 0
	}
	return o.TotalCount
}

// ListDirectDebitsPaymentsDirectDebitsListResponse - Direct debits list
type ListDirectDebitsPaymentsDirectDebitsListResponse struct {
	Data []ListDirectDebitsData `json:"data"`
	Meta ListDirectDebitsMeta   `json:"meta"`
}

func (o *ListDirectDebitsPaymentsDirectDebitsListResponse) GetData() []ListDirectDebitsData {
	if o == nil {
		return []ListDirectDebitsData{}
	}
	return o.Data
}

func (o *ListDirectDebitsPaymentsDirectDebitsListResponse) GetMeta() ListDirectDebitsMeta {
	if o == nil {
		return ListDirectDebitsMeta{}
	}
	return o.Meta
}

type ListDirectDebitsResponse struct {
	// Direct debits list
	TwoHundredApplicationJSONPaymentsDirectDebitsListResponse *ListDirectDebitsPaymentsDirectDebitsListResponse
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListDirectDebitsResponse) GetTwoHundredApplicationJSONPaymentsDirectDebitsListResponse() *ListDirectDebitsPaymentsDirectDebitsListResponse {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONPaymentsDirectDebitsListResponse
}

func (o *ListDirectDebitsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListDirectDebitsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *ListDirectDebitsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListDirectDebitsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
