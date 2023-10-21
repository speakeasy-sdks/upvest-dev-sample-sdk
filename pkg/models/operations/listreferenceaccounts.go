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

// ListReferenceAccountsOrder - Sort order of the result list if the `sort` parameter is specified. Use `ASC` for ascending or `DESC` for descending sort order.
type ListReferenceAccountsOrder string

const (
	ListReferenceAccountsOrderAsc  ListReferenceAccountsOrder = "ASC"
	ListReferenceAccountsOrderDesc ListReferenceAccountsOrder = "DESC"
)

func (e ListReferenceAccountsOrder) ToPointer() *ListReferenceAccountsOrder {
	return &e
}

func (e *ListReferenceAccountsOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = ListReferenceAccountsOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListReferenceAccountsOrder: %v", v)
	}
}

// ListReferenceAccountsSort - Field of resource to sort by
type ListReferenceAccountsSort string

const (
	ListReferenceAccountsSortID        ListReferenceAccountsSort = "id"
	ListReferenceAccountsSortCreatedAt ListReferenceAccountsSort = "created_at"
)

func (e ListReferenceAccountsSort) ToPointer() *ListReferenceAccountsSort {
	return &e
}

func (e *ListReferenceAccountsSort) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "id":
		fallthrough
	case "created_at":
		*e = ListReferenceAccountsSort(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListReferenceAccountsSort: %v", v)
	}
}

type ListReferenceAccountsRequest struct {
	// Use the `limit` argument to specify the maximum number of items returned.
	Limit *int `default:"100" queryParam:"style=form,explode=true,name=limit"`
	// Use the `offset` argument to specify where in the list of results to start when returning items for a particular query.
	Offset *int `queryParam:"style=form,explode=true,name=offset"`
	// Sort order of the result list if the `sort` parameter is specified. Use `ASC` for ascending or `DESC` for descending sort order.
	Order *ListReferenceAccountsOrder `default:"ASC" queryParam:"style=form,explode=true,name=order"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-http-header
	Signature string `header:"style=simple,explode=false,name=signature"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-input-http-he
	SignatureInput string `header:"style=simple,explode=false,name=signature-input"`
	// Field of resource to sort by
	Sort *ListReferenceAccountsSort `default:"created_at" queryParam:"style=form,explode=true,name=sort"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
	UserID         string `pathParam:"style=simple,explode=false,name=user_id"`
}

func (l ListReferenceAccountsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListReferenceAccountsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListReferenceAccountsRequest) GetLimit() *int {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListReferenceAccountsRequest) GetOffset() *int {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListReferenceAccountsRequest) GetOrder() *ListReferenceAccountsOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListReferenceAccountsRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *ListReferenceAccountsRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *ListReferenceAccountsRequest) GetSort() *ListReferenceAccountsSort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListReferenceAccountsRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *ListReferenceAccountsRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

func (o *ListReferenceAccountsRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type ListReferenceAccountsUsersListResponseReferenceAccount struct {
	// First and last name of the reference account owner
	AccountOwner string `json:"account_owner"`
	// Business Identifier Code (also known as SWIFT-BIC, BIC, SWIFT ID or SWIFT code) [ISO 9362](https://en.wikipedia.org/wiki/ISO_9362).
	Bic string `json:"bic"`
	// Timestamp of when user validated the reference account
	ConfirmedAt time.Time `json:"confirmed_at"`
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	// International Bank Account Number [IBAN](https://en.wikipedia.org/wiki/International_Bank_Account_Number).
	Iban string `json:"iban"`
	// Reference account unique identifier.
	ID string `json:"id"`
	// Human-readable name of the reference bank account
	Name string `json:"name"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time `json:"updated_at"`
	// User unique identifier.
	UserID string `json:"user_id"`
}

func (l ListReferenceAccountsUsersListResponseReferenceAccount) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListReferenceAccountsUsersListResponseReferenceAccount) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListReferenceAccountsUsersListResponseReferenceAccount) GetAccountOwner() string {
	if o == nil {
		return ""
	}
	return o.AccountOwner
}

func (o *ListReferenceAccountsUsersListResponseReferenceAccount) GetBic() string {
	if o == nil {
		return ""
	}
	return o.Bic
}

func (o *ListReferenceAccountsUsersListResponseReferenceAccount) GetConfirmedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.ConfirmedAt
}

func (o *ListReferenceAccountsUsersListResponseReferenceAccount) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *ListReferenceAccountsUsersListResponseReferenceAccount) GetIban() string {
	if o == nil {
		return ""
	}
	return o.Iban
}

func (o *ListReferenceAccountsUsersListResponseReferenceAccount) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ListReferenceAccountsUsersListResponseReferenceAccount) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ListReferenceAccountsUsersListResponseReferenceAccount) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *ListReferenceAccountsUsersListResponseReferenceAccount) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

// ListReferenceAccountsUsersListResponseMetaOrder - The ordering of the response.
// * ASC - Ascending order
// * DESC - Descending order
type ListReferenceAccountsUsersListResponseMetaOrder string

const (
	ListReferenceAccountsUsersListResponseMetaOrderAsc  ListReferenceAccountsUsersListResponseMetaOrder = "ASC"
	ListReferenceAccountsUsersListResponseMetaOrderDesc ListReferenceAccountsUsersListResponseMetaOrder = "DESC"
)

func (e ListReferenceAccountsUsersListResponseMetaOrder) ToPointer() *ListReferenceAccountsUsersListResponseMetaOrder {
	return &e
}

func (e *ListReferenceAccountsUsersListResponseMetaOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = ListReferenceAccountsUsersListResponseMetaOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListReferenceAccountsUsersListResponseMetaOrder: %v", v)
	}
}

type ListReferenceAccountsUsersListResponseMeta struct {
	// Count of the resources returned in the response.
	Count int64 `json:"count"`
	// Total limit of the response.
	Limit int64 `json:"limit"`
	// Amount of resource to offset in the response.
	Offset int64 `json:"offset"`
	// The ordering of the response.
	// * ASC - Ascending order
	// * DESC - Descending order
	Order *ListReferenceAccountsUsersListResponseMetaOrder `json:"order,omitempty"`
	// The field that the list is sorted by.
	Sort *string `json:"sort,omitempty"`
	// Total count of all the resources.
	TotalCount int64 `json:"total_count"`
}

func (o *ListReferenceAccountsUsersListResponseMeta) GetCount() int64 {
	if o == nil {
		return 0
	}
	return o.Count
}

func (o *ListReferenceAccountsUsersListResponseMeta) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

func (o *ListReferenceAccountsUsersListResponseMeta) GetOffset() int64 {
	if o == nil {
		return 0
	}
	return o.Offset
}

func (o *ListReferenceAccountsUsersListResponseMeta) GetOrder() *ListReferenceAccountsUsersListResponseMetaOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListReferenceAccountsUsersListResponseMeta) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListReferenceAccountsUsersListResponseMeta) GetTotalCount() int64 {
	if o == nil {
		return 0
	}
	return o.TotalCount
}

// ListReferenceAccountsUsersListResponse - OK
type ListReferenceAccountsUsersListResponse struct {
	Data []ListReferenceAccountsUsersListResponseReferenceAccount `json:"data"`
	Meta ListReferenceAccountsUsersListResponseMeta               `json:"meta"`
}

func (o *ListReferenceAccountsUsersListResponse) GetData() []ListReferenceAccountsUsersListResponseReferenceAccount {
	if o == nil {
		return []ListReferenceAccountsUsersListResponseReferenceAccount{}
	}
	return o.Data
}

func (o *ListReferenceAccountsUsersListResponse) GetMeta() ListReferenceAccountsUsersListResponseMeta {
	if o == nil {
		return ListReferenceAccountsUsersListResponseMeta{}
	}
	return o.Meta
}

type ListReferenceAccountsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	UsersListResponse *ListReferenceAccountsUsersListResponse
}

func (o *ListReferenceAccountsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListReferenceAccountsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *ListReferenceAccountsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListReferenceAccountsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListReferenceAccountsResponse) GetUsersListResponse() *ListReferenceAccountsUsersListResponse {
	if o == nil {
		return nil
	}
	return o.UsersListResponse
}