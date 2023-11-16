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

type UpdateAccountAccountUpdateRequest struct {
	// The name of the account.
	Name *string `json:"name,omitempty"`
}

func (o *UpdateAccountAccountUpdateRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type UpdateAccountRequest struct {
	RequestBody *UpdateAccountAccountUpdateRequest `request:"mediaType=application/json"`
	AccountID   string                             `pathParam:"style=simple,explode=false,name=account_id"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-http-header
	Signature string `header:"style=simple,explode=false,name=signature"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-input-http-he
	SignatureInput string `header:"style=simple,explode=false,name=signature-input"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
}

func (u UpdateAccountRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateAccountRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateAccountRequest) GetRequestBody() *UpdateAccountAccountUpdateRequest {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *UpdateAccountRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *UpdateAccountRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *UpdateAccountRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *UpdateAccountRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *UpdateAccountRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// UpdateAccountStatus - The status of the account
// * PENDING_APPROVAL - Account approval is pending - the account is visible through our API but cannot be acted on.
// * ACTIVE - Account is active - full functionality of the Investment API is accessible.
// * CLOSING - Account is closing - only sell orders or the transfer of positions out are permissible before the account is closed.
// * CLOSED - Account is closed with zero balance successfully.
// * LOCKED - Account is locked for all actions.
type UpdateAccountStatus string

const (
	UpdateAccountStatusPendingApproval UpdateAccountStatus = "PENDING_APPROVAL"
	UpdateAccountStatusActive          UpdateAccountStatus = "ACTIVE"
	UpdateAccountStatusClosing         UpdateAccountStatus = "CLOSING"
	UpdateAccountStatusClosed          UpdateAccountStatus = "CLOSED"
	UpdateAccountStatusLocked          UpdateAccountStatus = "LOCKED"
)

func (e UpdateAccountStatus) ToPointer() *UpdateAccountStatus {
	return &e
}

func (e *UpdateAccountStatus) UnmarshalJSON(data []byte) error {
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
		*e = UpdateAccountStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateAccountStatus: %v", v)
	}
}

// UpdateAccountType - Account type.
// * TRADING - Orders in accounts of this type are created on a specific instrument basis.
// * PORTFOLIO - Orders in accounts of this type are created on a portfolio basis and additional portfolio functionality is available.
type UpdateAccountType string

const (
	UpdateAccountTypeTrading   UpdateAccountType = "TRADING"
	UpdateAccountTypePortfolio UpdateAccountType = "PORTFOLIO"
)

func (e UpdateAccountType) ToPointer() *UpdateAccountType {
	return &e
}

func (e *UpdateAccountType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TRADING":
		fallthrough
	case "PORTFOLIO":
		*e = UpdateAccountType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateAccountType: %v", v)
	}
}

// UpdateAccountAccountsType - Relation type
// * OWNER -
type UpdateAccountAccountsType string

const (
	UpdateAccountAccountsTypeOwner UpdateAccountAccountsType = "OWNER"
)

func (e UpdateAccountAccountsType) ToPointer() *UpdateAccountAccountsType {
	return &e
}

func (e *UpdateAccountAccountsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OWNER":
		*e = UpdateAccountAccountsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateAccountAccountsType: %v", v)
	}
}

type UpdateAccountUsers struct {
	// User unique identifier.
	ID *string `json:"id,omitempty"`
	// Relation type
	// * OWNER -
	Type *UpdateAccountAccountsType `default:"OWNER" json:"type"`
}

func (u UpdateAccountUsers) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateAccountUsers) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateAccountUsers) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateAccountUsers) GetType() *UpdateAccountAccountsType {
	if o == nil {
		return nil
	}
	return o.Type
}

// UpdateAccountAccount - Account updated.
type UpdateAccountAccount struct {
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
	Status UpdateAccountStatus `json:"status"`
	// Account type.
	// * TRADING - Orders in accounts of this type are created on a specific instrument basis.
	// * PORTFOLIO - Orders in accounts of this type are created on a portfolio basis and additional portfolio functionality is available.
	Type UpdateAccountType `json:"type"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time            `json:"updated_at"`
	Users     []UpdateAccountUsers `json:"users"`
}

func (u UpdateAccountAccount) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateAccountAccount) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateAccountAccount) GetAccountGroupID() string {
	if o == nil {
		return ""
	}
	return o.AccountGroupID
}

func (o *UpdateAccountAccount) GetAccountNumber() int64 {
	if o == nil {
		return 0
	}
	return o.AccountNumber
}

func (o *UpdateAccountAccount) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *UpdateAccountAccount) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateAccountAccount) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpdateAccountAccount) GetStatus() UpdateAccountStatus {
	if o == nil {
		return UpdateAccountStatus("")
	}
	return o.Status
}

func (o *UpdateAccountAccount) GetType() UpdateAccountType {
	if o == nil {
		return UpdateAccountType("")
	}
	return o.Type
}

func (o *UpdateAccountAccount) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *UpdateAccountAccount) GetUsers() []UpdateAccountUsers {
	if o == nil {
		return []UpdateAccountUsers{}
	}
	return o.Users
}

type UpdateAccountResponse struct {
	// Account updated.
	TwoHundredApplicationJSONAccount *UpdateAccountAccount
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateAccountResponse) GetTwoHundredApplicationJSONAccount() *UpdateAccountAccount {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONAccount
}

func (o *UpdateAccountResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateAccountResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *UpdateAccountResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateAccountResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
