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

// CreateAccountRequestBodyType - Account type.
// * TRADING - Orders in accounts of this type are created on a specific instrument basis.
// * PORTFOLIO - Orders in accounts of this type are created on a portfolio basis and additional portfolio functionality is available.
type CreateAccountRequestBodyType string

const (
	CreateAccountRequestBodyTypeTrading   CreateAccountRequestBodyType = "TRADING"
	CreateAccountRequestBodyTypePortfolio CreateAccountRequestBodyType = "PORTFOLIO"
)

func (e CreateAccountRequestBodyType) ToPointer() *CreateAccountRequestBodyType {
	return &e
}

func (e *CreateAccountRequestBodyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TRADING":
		fallthrough
	case "PORTFOLIO":
		*e = CreateAccountRequestBodyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAccountRequestBodyType: %v", v)
	}
}

type CreateAccountRequestBody struct {
	// Account group unique identifier.
	AccountGroupID string `json:"account_group_id"`
	// The name of the account.
	Name *string `json:"name,omitempty"`
	// Account type.
	// * TRADING - Orders in accounts of this type are created on a specific instrument basis.
	// * PORTFOLIO - Orders in accounts of this type are created on a portfolio basis and additional portfolio functionality is available.
	Type CreateAccountRequestBodyType `json:"type"`
	// User unique identifier.
	UserID string `json:"user_id"`
}

func (o *CreateAccountRequestBody) GetAccountGroupID() string {
	if o == nil {
		return ""
	}
	return o.AccountGroupID
}

func (o *CreateAccountRequestBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CreateAccountRequestBody) GetType() CreateAccountRequestBodyType {
	if o == nil {
		return CreateAccountRequestBodyType("")
	}
	return o.Type
}

func (o *CreateAccountRequestBody) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type CreateAccountRequest struct {
	RequestBody *CreateAccountRequestBody `request:"mediaType=application/json"`
	// A UUID to be used as an idempotency key.  This prevents a duplicate request from being replayed.
	// https://docs.upvest.co/concepts/api_concepts/idempotency
	//
	IdempotencyKey string `header:"style=simple,explode=false,name=idempotency-key"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-http-header
	Signature string `header:"style=simple,explode=false,name=signature"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-input-http-he
	SignatureInput string `header:"style=simple,explode=false,name=signature-input"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
}

func (c CreateAccountRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateAccountRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateAccountRequest) GetRequestBody() *CreateAccountRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *CreateAccountRequest) GetIdempotencyKey() string {
	if o == nil {
		return ""
	}
	return o.IdempotencyKey
}

func (o *CreateAccountRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *CreateAccountRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *CreateAccountRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *CreateAccountRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// CreateAccountAccountStatus - The status of the account
// * PENDING_APPROVAL - Account approval is pending - the account is visible through our API but cannot be acted on.
// * ACTIVE - Account is active - full functionality of the Investment API is accessible.
// * CLOSING - Account is closing - only sell orders or the transfer of positions out are permissible before the account is closed.
// * CLOSED - Account is closed with zero balance successfully.
// * LOCKED - Account is locked for all actions.
type CreateAccountAccountStatus string

const (
	CreateAccountAccountStatusPendingApproval CreateAccountAccountStatus = "PENDING_APPROVAL"
	CreateAccountAccountStatusActive          CreateAccountAccountStatus = "ACTIVE"
	CreateAccountAccountStatusClosing         CreateAccountAccountStatus = "CLOSING"
	CreateAccountAccountStatusClosed          CreateAccountAccountStatus = "CLOSED"
	CreateAccountAccountStatusLocked          CreateAccountAccountStatus = "LOCKED"
)

func (e CreateAccountAccountStatus) ToPointer() *CreateAccountAccountStatus {
	return &e
}

func (e *CreateAccountAccountStatus) UnmarshalJSON(data []byte) error {
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
		*e = CreateAccountAccountStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAccountAccountStatus: %v", v)
	}
}

// CreateAccountAccountType - Account type.
// * TRADING - Orders in accounts of this type are created on a specific instrument basis.
// * PORTFOLIO - Orders in accounts of this type are created on a portfolio basis and additional portfolio functionality is available.
type CreateAccountAccountType string

const (
	CreateAccountAccountTypeTrading   CreateAccountAccountType = "TRADING"
	CreateAccountAccountTypePortfolio CreateAccountAccountType = "PORTFOLIO"
)

func (e CreateAccountAccountType) ToPointer() *CreateAccountAccountType {
	return &e
}

func (e *CreateAccountAccountType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TRADING":
		fallthrough
	case "PORTFOLIO":
		*e = CreateAccountAccountType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAccountAccountType: %v", v)
	}
}

// CreateAccountAccountUsersType - Relation type
// * OWNER -
type CreateAccountAccountUsersType string

const (
	CreateAccountAccountUsersTypeOwner CreateAccountAccountUsersType = "OWNER"
)

func (e CreateAccountAccountUsersType) ToPointer() *CreateAccountAccountUsersType {
	return &e
}

func (e *CreateAccountAccountUsersType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OWNER":
		*e = CreateAccountAccountUsersType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAccountAccountUsersType: %v", v)
	}
}

type CreateAccountAccountUsers struct {
	// User unique identifier.
	ID *string `json:"id,omitempty"`
	// Relation type
	// * OWNER -
	Type *CreateAccountAccountUsersType `default:"OWNER" json:"type"`
}

func (c CreateAccountAccountUsers) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateAccountAccountUsers) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateAccountAccountUsers) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateAccountAccountUsers) GetType() *CreateAccountAccountUsersType {
	if o == nil {
		return nil
	}
	return o.Type
}

// CreateAccountAccount - Account created.
type CreateAccountAccount struct {
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
	Status CreateAccountAccountStatus `json:"status"`
	// Account type.
	// * TRADING - Orders in accounts of this type are created on a specific instrument basis.
	// * PORTFOLIO - Orders in accounts of this type are created on a portfolio basis and additional portfolio functionality is available.
	Type CreateAccountAccountType `json:"type"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time                   `json:"updated_at"`
	Users     []CreateAccountAccountUsers `json:"users"`
}

func (c CreateAccountAccount) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateAccountAccount) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateAccountAccount) GetAccountGroupID() string {
	if o == nil {
		return ""
	}
	return o.AccountGroupID
}

func (o *CreateAccountAccount) GetAccountNumber() int64 {
	if o == nil {
		return 0
	}
	return o.AccountNumber
}

func (o *CreateAccountAccount) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *CreateAccountAccount) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateAccountAccount) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateAccountAccount) GetStatus() CreateAccountAccountStatus {
	if o == nil {
		return CreateAccountAccountStatus("")
	}
	return o.Status
}

func (o *CreateAccountAccount) GetType() CreateAccountAccountType {
	if o == nil {
		return CreateAccountAccountType("")
	}
	return o.Type
}

func (o *CreateAccountAccount) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *CreateAccountAccount) GetUsers() []CreateAccountAccountUsers {
	if o == nil {
		return []CreateAccountAccountUsers{}
	}
	return o.Users
}

type CreateAccountResponse struct {
	// Account created.
	Account *CreateAccountAccount
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateAccountResponse) GetAccount() *CreateAccountAccount {
	if o == nil {
		return nil
	}
	return o.Account
}

func (o *CreateAccountResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAccountResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreateAccountResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAccountResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
