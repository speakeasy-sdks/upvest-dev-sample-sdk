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

// ListUserAccountsSort - Sort the result by `created_at`.
type ListUserAccountsSort string

const (
	ListUserAccountsSortCreatedAt ListUserAccountsSort = "created_at"
)

func (e ListUserAccountsSort) ToPointer() *ListUserAccountsSort {
	return &e
}

func (e *ListUserAccountsSort) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		*e = ListUserAccountsSort(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListUserAccountsSort: %v", v)
	}
}

type ListUserAccountsRequest struct {
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
	Sort *ListUserAccountsSort `default:"created_at" queryParam:"style=form,explode=true,name=sort"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
	UserID         string `pathParam:"style=simple,explode=false,name=user_id"`
}

func (l ListUserAccountsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListUserAccountsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListUserAccountsRequest) GetLimit() *int {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListUserAccountsRequest) GetOffset() *int {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListUserAccountsRequest) GetOrder() *shared.Order {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListUserAccountsRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *ListUserAccountsRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *ListUserAccountsRequest) GetSort() *ListUserAccountsSort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListUserAccountsRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *ListUserAccountsRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

func (o *ListUserAccountsRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

// ListUserAccountsAccountsListResponseAccountStatus - The status of the account
// * PENDING_APPROVAL - Account approval is pending - the account is visible through our API but cannot be acted on.
// * ACTIVE - Account is active - full functionality of the Investment API is accessible.
// * CLOSING - Account is closing - only sell orders or the transfer of positions out are permissible before the account is closed.
// * CLOSED - Account is closed with zero balance successfully.
// * LOCKED - Account is locked for all actions.
type ListUserAccountsAccountsListResponseAccountStatus string

const (
	ListUserAccountsAccountsListResponseAccountStatusPendingApproval ListUserAccountsAccountsListResponseAccountStatus = "PENDING_APPROVAL"
	ListUserAccountsAccountsListResponseAccountStatusActive          ListUserAccountsAccountsListResponseAccountStatus = "ACTIVE"
	ListUserAccountsAccountsListResponseAccountStatusClosing         ListUserAccountsAccountsListResponseAccountStatus = "CLOSING"
	ListUserAccountsAccountsListResponseAccountStatusClosed          ListUserAccountsAccountsListResponseAccountStatus = "CLOSED"
	ListUserAccountsAccountsListResponseAccountStatusLocked          ListUserAccountsAccountsListResponseAccountStatus = "LOCKED"
)

func (e ListUserAccountsAccountsListResponseAccountStatus) ToPointer() *ListUserAccountsAccountsListResponseAccountStatus {
	return &e
}

func (e *ListUserAccountsAccountsListResponseAccountStatus) UnmarshalJSON(data []byte) error {
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
		*e = ListUserAccountsAccountsListResponseAccountStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListUserAccountsAccountsListResponseAccountStatus: %v", v)
	}
}

// ListUserAccountsAccountsListResponseAccountType - Account type.
// * TRADING - Orders in accounts of this type are created on a specific instrument basis.
// * PORTFOLIO - Orders in accounts of this type are created on a portfolio basis and additional portfolio functionality is available.
type ListUserAccountsAccountsListResponseAccountType string

const (
	ListUserAccountsAccountsListResponseAccountTypeTrading   ListUserAccountsAccountsListResponseAccountType = "TRADING"
	ListUserAccountsAccountsListResponseAccountTypePortfolio ListUserAccountsAccountsListResponseAccountType = "PORTFOLIO"
)

func (e ListUserAccountsAccountsListResponseAccountType) ToPointer() *ListUserAccountsAccountsListResponseAccountType {
	return &e
}

func (e *ListUserAccountsAccountsListResponseAccountType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TRADING":
		fallthrough
	case "PORTFOLIO":
		*e = ListUserAccountsAccountsListResponseAccountType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListUserAccountsAccountsListResponseAccountType: %v", v)
	}
}

// ListUserAccountsAccountsListResponseAccountUsersType - Relation type
// * OWNER -
type ListUserAccountsAccountsListResponseAccountUsersType string

const (
	ListUserAccountsAccountsListResponseAccountUsersTypeOwner ListUserAccountsAccountsListResponseAccountUsersType = "OWNER"
)

func (e ListUserAccountsAccountsListResponseAccountUsersType) ToPointer() *ListUserAccountsAccountsListResponseAccountUsersType {
	return &e
}

func (e *ListUserAccountsAccountsListResponseAccountUsersType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OWNER":
		*e = ListUserAccountsAccountsListResponseAccountUsersType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListUserAccountsAccountsListResponseAccountUsersType: %v", v)
	}
}

type ListUserAccountsAccountsListResponseAccountUsers struct {
	// User unique identifier.
	ID *string `json:"id,omitempty"`
	// Relation type
	// * OWNER -
	Type *ListUserAccountsAccountsListResponseAccountUsersType `default:"OWNER" json:"type"`
}

func (l ListUserAccountsAccountsListResponseAccountUsers) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListUserAccountsAccountsListResponseAccountUsers) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListUserAccountsAccountsListResponseAccountUsers) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ListUserAccountsAccountsListResponseAccountUsers) GetType() *ListUserAccountsAccountsListResponseAccountUsersType {
	if o == nil {
		return nil
	}
	return o.Type
}

type ListUserAccountsAccountsListResponseAccount struct {
	// Account group unique identifier.
	AccountGroupID string `json:"account_group_id"`
	// The serial account number of the account in the account group.
	AccountNumber int64 `json:"account_number"`
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	// Account unique identifier.
	ID string `json:"id"`
	// The name of the account.
	Name string `json:"name"`
	// The status of the account
	// * PENDING_APPROVAL - Account approval is pending - the account is visible through our API but cannot be acted on.
	// * ACTIVE - Account is active - full functionality of the Investment API is accessible.
	// * CLOSING - Account is closing - only sell orders or the transfer of positions out are permissible before the account is closed.
	// * CLOSED - Account is closed with zero balance successfully.
	// * LOCKED - Account is locked for all actions.
	Status ListUserAccountsAccountsListResponseAccountStatus `json:"status"`
	// Account type.
	// * TRADING - Orders in accounts of this type are created on a specific instrument basis.
	// * PORTFOLIO - Orders in accounts of this type are created on a portfolio basis and additional portfolio functionality is available.
	Type ListUserAccountsAccountsListResponseAccountType `json:"type"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time                                          `json:"updated_at"`
	Users     []ListUserAccountsAccountsListResponseAccountUsers `json:"users"`
}

func (l ListUserAccountsAccountsListResponseAccount) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListUserAccountsAccountsListResponseAccount) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListUserAccountsAccountsListResponseAccount) GetAccountGroupID() string {
	if o == nil {
		return ""
	}
	return o.AccountGroupID
}

func (o *ListUserAccountsAccountsListResponseAccount) GetAccountNumber() int64 {
	if o == nil {
		return 0
	}
	return o.AccountNumber
}

func (o *ListUserAccountsAccountsListResponseAccount) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *ListUserAccountsAccountsListResponseAccount) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ListUserAccountsAccountsListResponseAccount) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ListUserAccountsAccountsListResponseAccount) GetStatus() ListUserAccountsAccountsListResponseAccountStatus {
	if o == nil {
		return ListUserAccountsAccountsListResponseAccountStatus("")
	}
	return o.Status
}

func (o *ListUserAccountsAccountsListResponseAccount) GetType() ListUserAccountsAccountsListResponseAccountType {
	if o == nil {
		return ListUserAccountsAccountsListResponseAccountType("")
	}
	return o.Type
}

func (o *ListUserAccountsAccountsListResponseAccount) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *ListUserAccountsAccountsListResponseAccount) GetUsers() []ListUserAccountsAccountsListResponseAccountUsers {
	if o == nil {
		return []ListUserAccountsAccountsListResponseAccountUsers{}
	}
	return o.Users
}

// ListUserAccountsAccountsListResponseMetaOrder - The ordering of the response.
// * ASC - Ascending order
// * DESC - Descending order
type ListUserAccountsAccountsListResponseMetaOrder string

const (
	ListUserAccountsAccountsListResponseMetaOrderAsc  ListUserAccountsAccountsListResponseMetaOrder = "ASC"
	ListUserAccountsAccountsListResponseMetaOrderDesc ListUserAccountsAccountsListResponseMetaOrder = "DESC"
)

func (e ListUserAccountsAccountsListResponseMetaOrder) ToPointer() *ListUserAccountsAccountsListResponseMetaOrder {
	return &e
}

func (e *ListUserAccountsAccountsListResponseMetaOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = ListUserAccountsAccountsListResponseMetaOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListUserAccountsAccountsListResponseMetaOrder: %v", v)
	}
}

type ListUserAccountsAccountsListResponseMeta struct {
	// Count of the resources returned in the response.
	Count int64 `json:"count"`
	// Total limit of the response.
	Limit int64 `json:"limit"`
	// Amount of resource to offset in the response.
	Offset int64 `json:"offset"`
	// The ordering of the response.
	// * ASC - Ascending order
	// * DESC - Descending order
	Order *ListUserAccountsAccountsListResponseMetaOrder `json:"order,omitempty"`
	// The field that the list is sorted by.
	Sort *string `json:"sort,omitempty"`
	// Total count of all the resources.
	TotalCount int64 `json:"total_count"`
}

func (o *ListUserAccountsAccountsListResponseMeta) GetCount() int64 {
	if o == nil {
		return 0
	}
	return o.Count
}

func (o *ListUserAccountsAccountsListResponseMeta) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

func (o *ListUserAccountsAccountsListResponseMeta) GetOffset() int64 {
	if o == nil {
		return 0
	}
	return o.Offset
}

func (o *ListUserAccountsAccountsListResponseMeta) GetOrder() *ListUserAccountsAccountsListResponseMetaOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListUserAccountsAccountsListResponseMeta) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListUserAccountsAccountsListResponseMeta) GetTotalCount() int64 {
	if o == nil {
		return 0
	}
	return o.TotalCount
}

// ListUserAccountsAccountsListResponse - OK
type ListUserAccountsAccountsListResponse struct {
	Data []ListUserAccountsAccountsListResponseAccount `json:"data"`
	Meta ListUserAccountsAccountsListResponseMeta      `json:"meta"`
}

func (o *ListUserAccountsAccountsListResponse) GetData() []ListUserAccountsAccountsListResponseAccount {
	if o == nil {
		return []ListUserAccountsAccountsListResponseAccount{}
	}
	return o.Data
}

func (o *ListUserAccountsAccountsListResponse) GetMeta() ListUserAccountsAccountsListResponseMeta {
	if o == nil {
		return ListUserAccountsAccountsListResponseMeta{}
	}
	return o.Meta
}

type ListUserAccountsResponse struct {
	// OK
	AccountsListResponse *ListUserAccountsAccountsListResponse
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListUserAccountsResponse) GetAccountsListResponse() *ListUserAccountsAccountsListResponse {
	if o == nil {
		return nil
	}
	return o.AccountsListResponse
}

func (o *ListUserAccountsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListUserAccountsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *ListUserAccountsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListUserAccountsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
