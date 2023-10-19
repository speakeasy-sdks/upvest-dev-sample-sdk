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

type RetrieveAccountRequest struct {
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-http-header
	Signature string `header:"style=simple,explode=false,name=signature"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-input-http-he
	SignatureInput string `header:"style=simple,explode=false,name=signature-input"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
}

func (r RetrieveAccountRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrieveAccountRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrieveAccountRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *RetrieveAccountRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *RetrieveAccountRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *RetrieveAccountRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *RetrieveAccountRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// RetrieveAccountAccountStatus - The status of the account
// * PENDING_APPROVAL - Account approval is pending - the account is visible through our API but cannot be acted on.
// * ACTIVE - Account is active - full functionality of the Investment API is accessible.
// * CLOSING - Account is closing - only sell orders or the transfer of positions out are permissible before the account is closed.
// * CLOSED - Account is closed with zero balance successfully.
// * LOCKED - Account is locked for all actions.
type RetrieveAccountAccountStatus string

const (
	RetrieveAccountAccountStatusPendingApproval RetrieveAccountAccountStatus = "PENDING_APPROVAL"
	RetrieveAccountAccountStatusActive          RetrieveAccountAccountStatus = "ACTIVE"
	RetrieveAccountAccountStatusClosing         RetrieveAccountAccountStatus = "CLOSING"
	RetrieveAccountAccountStatusClosed          RetrieveAccountAccountStatus = "CLOSED"
	RetrieveAccountAccountStatusLocked          RetrieveAccountAccountStatus = "LOCKED"
)

func (e RetrieveAccountAccountStatus) ToPointer() *RetrieveAccountAccountStatus {
	return &e
}

func (e *RetrieveAccountAccountStatus) UnmarshalJSON(data []byte) error {
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
		*e = RetrieveAccountAccountStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveAccountAccountStatus: %v", v)
	}
}

// RetrieveAccountAccountType - Account type.
// * TRADING - Orders in accounts of this type are created on a specific instrument basis.
// * PORTFOLIO - Orders in accounts of this type are created on a portfolio basis and additional portfolio functionality is available.
type RetrieveAccountAccountType string

const (
	RetrieveAccountAccountTypeTrading   RetrieveAccountAccountType = "TRADING"
	RetrieveAccountAccountTypePortfolio RetrieveAccountAccountType = "PORTFOLIO"
)

func (e RetrieveAccountAccountType) ToPointer() *RetrieveAccountAccountType {
	return &e
}

func (e *RetrieveAccountAccountType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TRADING":
		fallthrough
	case "PORTFOLIO":
		*e = RetrieveAccountAccountType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveAccountAccountType: %v", v)
	}
}

// RetrieveAccountAccountUsersType - Relation type
// * OWNER -
type RetrieveAccountAccountUsersType string

const (
	RetrieveAccountAccountUsersTypeOwner RetrieveAccountAccountUsersType = "OWNER"
)

func (e RetrieveAccountAccountUsersType) ToPointer() *RetrieveAccountAccountUsersType {
	return &e
}

func (e *RetrieveAccountAccountUsersType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OWNER":
		*e = RetrieveAccountAccountUsersType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveAccountAccountUsersType: %v", v)
	}
}

type RetrieveAccountAccountUsers struct {
	// User unique identifier.
	ID *string `json:"id,omitempty"`
	// Relation type
	// * OWNER -
	Type *RetrieveAccountAccountUsersType `default:"OWNER" json:"type"`
}

func (r RetrieveAccountAccountUsers) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrieveAccountAccountUsers) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrieveAccountAccountUsers) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RetrieveAccountAccountUsers) GetType() *RetrieveAccountAccountUsersType {
	if o == nil {
		return nil
	}
	return o.Type
}

// RetrieveAccountAccount - OK
type RetrieveAccountAccount struct {
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
	Status RetrieveAccountAccountStatus `json:"status"`
	// Account type.
	// * TRADING - Orders in accounts of this type are created on a specific instrument basis.
	// * PORTFOLIO - Orders in accounts of this type are created on a portfolio basis and additional portfolio functionality is available.
	Type RetrieveAccountAccountType `json:"type"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time                     `json:"updated_at"`
	Users     []RetrieveAccountAccountUsers `json:"users"`
}

func (r RetrieveAccountAccount) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrieveAccountAccount) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrieveAccountAccount) GetAccountGroupID() string {
	if o == nil {
		return ""
	}
	return o.AccountGroupID
}

func (o *RetrieveAccountAccount) GetAccountNumber() int64 {
	if o == nil {
		return 0
	}
	return o.AccountNumber
}

func (o *RetrieveAccountAccount) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *RetrieveAccountAccount) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveAccountAccount) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *RetrieveAccountAccount) GetStatus() RetrieveAccountAccountStatus {
	if o == nil {
		return RetrieveAccountAccountStatus("")
	}
	return o.Status
}

func (o *RetrieveAccountAccount) GetType() RetrieveAccountAccountType {
	if o == nil {
		return RetrieveAccountAccountType("")
	}
	return o.Type
}

func (o *RetrieveAccountAccount) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *RetrieveAccountAccount) GetUsers() []RetrieveAccountAccountUsers {
	if o == nil {
		return []RetrieveAccountAccountUsers{}
	}
	return o.Users
}

type RetrieveAccountResponse struct {
	// OK
	Account *RetrieveAccountAccount
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RetrieveAccountResponse) GetAccount() *RetrieveAccountAccount {
	if o == nil {
		return nil
	}
	return o.Account
}

func (o *RetrieveAccountResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RetrieveAccountResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *RetrieveAccountResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RetrieveAccountResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
