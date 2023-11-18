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

// Sort the result by `created_at`.
type Sort string

const (
	SortCreatedAt Sort = "created_at"
)

func (e Sort) ToPointer() *Sort {
	return &e
}

func (e *Sort) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		*e = Sort(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Sort: %v", v)
	}
}

type ListAccountGroupsRequest struct {
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
	// Sort the result by `created_at`.
	Sort *Sort `default:"created_at" queryParam:"style=form,explode=true,name=sort"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
}

func (l ListAccountGroupsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountGroupsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAccountGroupsRequest) GetLimit() *int {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAccountGroupsRequest) GetOffset() *int {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListAccountGroupsRequest) GetOrder() *shared.Order {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListAccountGroupsRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *ListAccountGroupsRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *ListAccountGroupsRequest) GetSort() *Sort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListAccountGroupsRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *ListAccountGroupsRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// ListAccountGroupsStatus - Status of the account group
// * PENDING_APPROVAL - Account group approval is pending - the account group is visible through our API but cannot be acted on.
// * ACTIVE - Account group is active - full functionality of the Investment API is accessible.
// * CLOSING - Account group is closing.
// * CLOSED - Account group is closed.
// * LOCKED - Account group is locked for all actions.
type ListAccountGroupsStatus string

const (
	ListAccountGroupsStatusPendingApproval ListAccountGroupsStatus = "PENDING_APPROVAL"
	ListAccountGroupsStatusActive          ListAccountGroupsStatus = "ACTIVE"
	ListAccountGroupsStatusClosing         ListAccountGroupsStatus = "CLOSING"
	ListAccountGroupsStatusClosed          ListAccountGroupsStatus = "CLOSED"
	ListAccountGroupsStatusLocked          ListAccountGroupsStatus = "LOCKED"
)

func (e ListAccountGroupsStatus) ToPointer() *ListAccountGroupsStatus {
	return &e
}

func (e *ListAccountGroupsStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PENDING_APPROVAL":
		fallthrough
	case "ACTIVE":
		fallthrough
	case "CLOSING":
		fallthrough
	case "CLOSED":
		fallthrough
	case "LOCKED":
		*e = ListAccountGroupsStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountGroupsStatus: %v", v)
	}
}

// ListAccountGroupsType - Account group type.
// * PERSONAL - Account group of a person holding assets on their own behalf.
// * LEGAL_ENTITY - Account group of a legal entity holding assets on behalf of their users.
type ListAccountGroupsType string

const (
	ListAccountGroupsTypePersonal    ListAccountGroupsType = "PERSONAL"
	ListAccountGroupsTypeLegalEntity ListAccountGroupsType = "LEGAL_ENTITY"
)

func (e ListAccountGroupsType) ToPointer() *ListAccountGroupsType {
	return &e
}

func (e *ListAccountGroupsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PERSONAL":
		fallthrough
	case "LEGAL_ENTITY":
		*e = ListAccountGroupsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountGroupsType: %v", v)
	}
}

// ListAccountGroupsAccountsType - Relation type
// * OWNER -
type ListAccountGroupsAccountsType string

const (
	ListAccountGroupsAccountsTypeOwner ListAccountGroupsAccountsType = "OWNER"
)

func (e ListAccountGroupsAccountsType) ToPointer() *ListAccountGroupsAccountsType {
	return &e
}

func (e *ListAccountGroupsAccountsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OWNER":
		*e = ListAccountGroupsAccountsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountGroupsAccountsType: %v", v)
	}
}

type ListAccountGroupsUsers struct {
	// User unique identifier.
	ID *string `json:"id,omitempty"`
	// Relation type
	// * OWNER -
	Type *ListAccountGroupsAccountsType `default:"OWNER" json:"type"`
}

func (l ListAccountGroupsUsers) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountGroupsUsers) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAccountGroupsUsers) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ListAccountGroupsUsers) GetType() *ListAccountGroupsAccountsType {
	if o == nil {
		return nil
	}
	return o.Type
}

type AccountGroup struct {
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	// Account group unique identifier.
	ID string `json:"id"`
	// Securities account number
	SecuritiesAccountNumber string `json:"securities_account_number"`
	// Status of the account group
	// * PENDING_APPROVAL - Account group approval is pending - the account group is visible through our API but cannot be acted on.
	// * ACTIVE - Account group is active - full functionality of the Investment API is accessible.
	// * CLOSING - Account group is closing.
	// * CLOSED - Account group is closed.
	// * LOCKED - Account group is locked for all actions.
	Status ListAccountGroupsStatus `json:"status"`
	// Account group type.
	// * PERSONAL - Account group of a person holding assets on their own behalf.
	// * LEGAL_ENTITY - Account group of a legal entity holding assets on behalf of their users.
	Type ListAccountGroupsType `json:"type"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time                `json:"updated_at"`
	Users     []ListAccountGroupsUsers `json:"users"`
}

func (a AccountGroup) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountGroup) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountGroup) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *AccountGroup) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AccountGroup) GetSecuritiesAccountNumber() string {
	if o == nil {
		return ""
	}
	return o.SecuritiesAccountNumber
}

func (o *AccountGroup) GetStatus() ListAccountGroupsStatus {
	if o == nil {
		return ListAccountGroupsStatus("")
	}
	return o.Status
}

func (o *AccountGroup) GetType() ListAccountGroupsType {
	if o == nil {
		return ListAccountGroupsType("")
	}
	return o.Type
}

func (o *AccountGroup) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *AccountGroup) GetUsers() []ListAccountGroupsUsers {
	if o == nil {
		return []ListAccountGroupsUsers{}
	}
	return o.Users
}

// ListAccountGroupsOrder - The ordering of the response.
// * ASC - Ascending order
// * DESC - Descending order
type ListAccountGroupsOrder string

const (
	ListAccountGroupsOrderAsc  ListAccountGroupsOrder = "ASC"
	ListAccountGroupsOrderDesc ListAccountGroupsOrder = "DESC"
)

func (e ListAccountGroupsOrder) ToPointer() *ListAccountGroupsOrder {
	return &e
}

func (e *ListAccountGroupsOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = ListAccountGroupsOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountGroupsOrder: %v", v)
	}
}

type Meta struct {
	// Count of the resources returned in the response.
	Count int64 `json:"count"`
	// Total limit of the response.
	Limit int64 `json:"limit"`
	// Amount of resource to offset in the response.
	Offset int64 `json:"offset"`
	// The ordering of the response.
	// * ASC - Ascending order
	// * DESC - Descending order
	Order *ListAccountGroupsOrder `json:"order,omitempty"`
	// The field that the list is sorted by.
	Sort *string `json:"sort,omitempty"`
	// Total count of all the resources.
	TotalCount int64 `json:"total_count"`
}

func (o *Meta) GetCount() int64 {
	if o == nil {
		return 0
	}
	return o.Count
}

func (o *Meta) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

func (o *Meta) GetOffset() int64 {
	if o == nil {
		return 0
	}
	return o.Offset
}

func (o *Meta) GetOrder() *ListAccountGroupsOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *Meta) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *Meta) GetTotalCount() int64 {
	if o == nil {
		return 0
	}
	return o.TotalCount
}

// ListAccountGroupsAccountGroupsListResponse - OK
type ListAccountGroupsAccountGroupsListResponse struct {
	Data []AccountGroup `json:"data"`
	Meta Meta           `json:"meta"`
}

func (o *ListAccountGroupsAccountGroupsListResponse) GetData() []AccountGroup {
	if o == nil {
		return []AccountGroup{}
	}
	return o.Data
}

func (o *ListAccountGroupsAccountGroupsListResponse) GetMeta() Meta {
	if o == nil {
		return Meta{}
	}
	return o.Meta
}

type ListAccountGroupsResponse struct {
	// OK
	AccountGroupsListResponse *ListAccountGroupsAccountGroupsListResponse
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListAccountGroupsResponse) GetAccountGroupsListResponse() *ListAccountGroupsAccountGroupsListResponse {
	if o == nil {
		return nil
	}
	return o.AccountGroupsListResponse
}

func (o *ListAccountGroupsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAccountGroupsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *ListAccountGroupsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAccountGroupsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
