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

// ListUserAccountGroupsSort - Sort the result by `created_at`.
type ListUserAccountGroupsSort string

const (
	ListUserAccountGroupsSortCreatedAt ListUserAccountGroupsSort = "created_at"
)

func (e ListUserAccountGroupsSort) ToPointer() *ListUserAccountGroupsSort {
	return &e
}

func (e *ListUserAccountGroupsSort) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		*e = ListUserAccountGroupsSort(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListUserAccountGroupsSort: %v", v)
	}
}

type ListUserAccountGroupsRequest struct {
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
	Sort *ListUserAccountGroupsSort `default:"created_at" queryParam:"style=form,explode=true,name=sort"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
	UserID         string `pathParam:"style=simple,explode=false,name=user_id"`
}

func (l ListUserAccountGroupsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListUserAccountGroupsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListUserAccountGroupsRequest) GetLimit() *int {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListUserAccountGroupsRequest) GetOffset() *int {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListUserAccountGroupsRequest) GetOrder() *shared.Order {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListUserAccountGroupsRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *ListUserAccountGroupsRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *ListUserAccountGroupsRequest) GetSort() *ListUserAccountGroupsSort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListUserAccountGroupsRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *ListUserAccountGroupsRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

func (o *ListUserAccountGroupsRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

// ListUserAccountGroupsAccountGroupsListResponseAccountGroupStatus - Status of the account group
// * PENDING_APPROVAL - Account group approval is pending - the account group is visible through our API but cannot be acted on.
// * ACTIVE - Account group is active - full functionality of the Investment API is accessible.
// * CLOSING - Account group is closing.
// * CLOSED - Account group is closed.
// * LOCKED - Account group is locked for all actions.
type ListUserAccountGroupsAccountGroupsListResponseAccountGroupStatus string

const (
	ListUserAccountGroupsAccountGroupsListResponseAccountGroupStatusPendingApproval ListUserAccountGroupsAccountGroupsListResponseAccountGroupStatus = "PENDING_APPROVAL"
	ListUserAccountGroupsAccountGroupsListResponseAccountGroupStatusActive          ListUserAccountGroupsAccountGroupsListResponseAccountGroupStatus = "ACTIVE"
	ListUserAccountGroupsAccountGroupsListResponseAccountGroupStatusClosing         ListUserAccountGroupsAccountGroupsListResponseAccountGroupStatus = "CLOSING"
	ListUserAccountGroupsAccountGroupsListResponseAccountGroupStatusClosed          ListUserAccountGroupsAccountGroupsListResponseAccountGroupStatus = "CLOSED"
	ListUserAccountGroupsAccountGroupsListResponseAccountGroupStatusLocked          ListUserAccountGroupsAccountGroupsListResponseAccountGroupStatus = "LOCKED"
)

func (e ListUserAccountGroupsAccountGroupsListResponseAccountGroupStatus) ToPointer() *ListUserAccountGroupsAccountGroupsListResponseAccountGroupStatus {
	return &e
}

func (e *ListUserAccountGroupsAccountGroupsListResponseAccountGroupStatus) UnmarshalJSON(data []byte) error {
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
		*e = ListUserAccountGroupsAccountGroupsListResponseAccountGroupStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListUserAccountGroupsAccountGroupsListResponseAccountGroupStatus: %v", v)
	}
}

// ListUserAccountGroupsAccountGroupsListResponseAccountGroupType - Account group type.
// * PERSONAL - Account group of a person holding assets on their own behalf.
// * LEGAL_ENTITY - Account group of a legal entity holding assets on behalf of their users.
type ListUserAccountGroupsAccountGroupsListResponseAccountGroupType string

const (
	ListUserAccountGroupsAccountGroupsListResponseAccountGroupTypePersonal    ListUserAccountGroupsAccountGroupsListResponseAccountGroupType = "PERSONAL"
	ListUserAccountGroupsAccountGroupsListResponseAccountGroupTypeLegalEntity ListUserAccountGroupsAccountGroupsListResponseAccountGroupType = "LEGAL_ENTITY"
)

func (e ListUserAccountGroupsAccountGroupsListResponseAccountGroupType) ToPointer() *ListUserAccountGroupsAccountGroupsListResponseAccountGroupType {
	return &e
}

func (e *ListUserAccountGroupsAccountGroupsListResponseAccountGroupType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PERSONAL":
		fallthrough
	case "LEGAL_ENTITY":
		*e = ListUserAccountGroupsAccountGroupsListResponseAccountGroupType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListUserAccountGroupsAccountGroupsListResponseAccountGroupType: %v", v)
	}
}

// ListUserAccountGroupsAccountGroupsListResponseAccountGroupUsersType - Relation type
// * OWNER -
type ListUserAccountGroupsAccountGroupsListResponseAccountGroupUsersType string

const (
	ListUserAccountGroupsAccountGroupsListResponseAccountGroupUsersTypeOwner ListUserAccountGroupsAccountGroupsListResponseAccountGroupUsersType = "OWNER"
)

func (e ListUserAccountGroupsAccountGroupsListResponseAccountGroupUsersType) ToPointer() *ListUserAccountGroupsAccountGroupsListResponseAccountGroupUsersType {
	return &e
}

func (e *ListUserAccountGroupsAccountGroupsListResponseAccountGroupUsersType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OWNER":
		*e = ListUserAccountGroupsAccountGroupsListResponseAccountGroupUsersType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListUserAccountGroupsAccountGroupsListResponseAccountGroupUsersType: %v", v)
	}
}

type ListUserAccountGroupsAccountGroupsListResponseAccountGroupUsers struct {
	// User unique identifier.
	ID *string `json:"id,omitempty"`
	// Relation type
	// * OWNER -
	Type *ListUserAccountGroupsAccountGroupsListResponseAccountGroupUsersType `default:"OWNER" json:"type"`
}

func (l ListUserAccountGroupsAccountGroupsListResponseAccountGroupUsers) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListUserAccountGroupsAccountGroupsListResponseAccountGroupUsers) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListUserAccountGroupsAccountGroupsListResponseAccountGroupUsers) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ListUserAccountGroupsAccountGroupsListResponseAccountGroupUsers) GetType() *ListUserAccountGroupsAccountGroupsListResponseAccountGroupUsersType {
	if o == nil {
		return nil
	}
	return o.Type
}

type ListUserAccountGroupsAccountGroupsListResponseAccountGroup struct {
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
	Status ListUserAccountGroupsAccountGroupsListResponseAccountGroupStatus `json:"status"`
	// Account group type.
	// * PERSONAL - Account group of a person holding assets on their own behalf.
	// * LEGAL_ENTITY - Account group of a legal entity holding assets on behalf of their users.
	Type ListUserAccountGroupsAccountGroupsListResponseAccountGroupType `json:"type"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time                                                         `json:"updated_at"`
	Users     []ListUserAccountGroupsAccountGroupsListResponseAccountGroupUsers `json:"users"`
}

func (l ListUserAccountGroupsAccountGroupsListResponseAccountGroup) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListUserAccountGroupsAccountGroupsListResponseAccountGroup) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListUserAccountGroupsAccountGroupsListResponseAccountGroup) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *ListUserAccountGroupsAccountGroupsListResponseAccountGroup) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ListUserAccountGroupsAccountGroupsListResponseAccountGroup) GetSecuritiesAccountNumber() string {
	if o == nil {
		return ""
	}
	return o.SecuritiesAccountNumber
}

func (o *ListUserAccountGroupsAccountGroupsListResponseAccountGroup) GetStatus() ListUserAccountGroupsAccountGroupsListResponseAccountGroupStatus {
	if o == nil {
		return ListUserAccountGroupsAccountGroupsListResponseAccountGroupStatus("")
	}
	return o.Status
}

func (o *ListUserAccountGroupsAccountGroupsListResponseAccountGroup) GetType() ListUserAccountGroupsAccountGroupsListResponseAccountGroupType {
	if o == nil {
		return ListUserAccountGroupsAccountGroupsListResponseAccountGroupType("")
	}
	return o.Type
}

func (o *ListUserAccountGroupsAccountGroupsListResponseAccountGroup) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *ListUserAccountGroupsAccountGroupsListResponseAccountGroup) GetUsers() []ListUserAccountGroupsAccountGroupsListResponseAccountGroupUsers {
	if o == nil {
		return []ListUserAccountGroupsAccountGroupsListResponseAccountGroupUsers{}
	}
	return o.Users
}

// ListUserAccountGroupsAccountGroupsListResponseMetaOrder - The ordering of the response.
// * ASC - Ascending order
// * DESC - Descending order
type ListUserAccountGroupsAccountGroupsListResponseMetaOrder string

const (
	ListUserAccountGroupsAccountGroupsListResponseMetaOrderAsc  ListUserAccountGroupsAccountGroupsListResponseMetaOrder = "ASC"
	ListUserAccountGroupsAccountGroupsListResponseMetaOrderDesc ListUserAccountGroupsAccountGroupsListResponseMetaOrder = "DESC"
)

func (e ListUserAccountGroupsAccountGroupsListResponseMetaOrder) ToPointer() *ListUserAccountGroupsAccountGroupsListResponseMetaOrder {
	return &e
}

func (e *ListUserAccountGroupsAccountGroupsListResponseMetaOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = ListUserAccountGroupsAccountGroupsListResponseMetaOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListUserAccountGroupsAccountGroupsListResponseMetaOrder: %v", v)
	}
}

type ListUserAccountGroupsAccountGroupsListResponseMeta struct {
	// Count of the resources returned in the response.
	Count int64 `json:"count"`
	// Total limit of the response.
	Limit int64 `json:"limit"`
	// Amount of resource to offset in the response.
	Offset int64 `json:"offset"`
	// The ordering of the response.
	// * ASC - Ascending order
	// * DESC - Descending order
	Order *ListUserAccountGroupsAccountGroupsListResponseMetaOrder `json:"order,omitempty"`
	// The field that the list is sorted by.
	Sort *string `json:"sort,omitempty"`
	// Total count of all the resources.
	TotalCount int64 `json:"total_count"`
}

func (o *ListUserAccountGroupsAccountGroupsListResponseMeta) GetCount() int64 {
	if o == nil {
		return 0
	}
	return o.Count
}

func (o *ListUserAccountGroupsAccountGroupsListResponseMeta) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

func (o *ListUserAccountGroupsAccountGroupsListResponseMeta) GetOffset() int64 {
	if o == nil {
		return 0
	}
	return o.Offset
}

func (o *ListUserAccountGroupsAccountGroupsListResponseMeta) GetOrder() *ListUserAccountGroupsAccountGroupsListResponseMetaOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListUserAccountGroupsAccountGroupsListResponseMeta) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListUserAccountGroupsAccountGroupsListResponseMeta) GetTotalCount() int64 {
	if o == nil {
		return 0
	}
	return o.TotalCount
}

// ListUserAccountGroupsAccountGroupsListResponse - OK
type ListUserAccountGroupsAccountGroupsListResponse struct {
	Data []ListUserAccountGroupsAccountGroupsListResponseAccountGroup `json:"data"`
	Meta ListUserAccountGroupsAccountGroupsListResponseMeta           `json:"meta"`
}

func (o *ListUserAccountGroupsAccountGroupsListResponse) GetData() []ListUserAccountGroupsAccountGroupsListResponseAccountGroup {
	if o == nil {
		return []ListUserAccountGroupsAccountGroupsListResponseAccountGroup{}
	}
	return o.Data
}

func (o *ListUserAccountGroupsAccountGroupsListResponse) GetMeta() ListUserAccountGroupsAccountGroupsListResponseMeta {
	if o == nil {
		return ListUserAccountGroupsAccountGroupsListResponseMeta{}
	}
	return o.Meta
}

type ListUserAccountGroupsResponse struct {
	// OK
	AccountGroupsListResponse *ListUserAccountGroupsAccountGroupsListResponse
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListUserAccountGroupsResponse) GetAccountGroupsListResponse() *ListUserAccountGroupsAccountGroupsListResponse {
	if o == nil {
		return nil
	}
	return o.AccountGroupsListResponse
}

func (o *ListUserAccountGroupsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListUserAccountGroupsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *ListUserAccountGroupsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListUserAccountGroupsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}