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

type RetrieveAccountGroupRequest struct {
	AccountGroupID string `pathParam:"style=simple,explode=false,name=account_group_id"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-http-header
	Signature string `header:"style=simple,explode=false,name=signature"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-input-http-he
	SignatureInput string `header:"style=simple,explode=false,name=signature-input"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
}

func (r RetrieveAccountGroupRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrieveAccountGroupRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrieveAccountGroupRequest) GetAccountGroupID() string {
	if o == nil {
		return ""
	}
	return o.AccountGroupID
}

func (o *RetrieveAccountGroupRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *RetrieveAccountGroupRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *RetrieveAccountGroupRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *RetrieveAccountGroupRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// RetrieveAccountGroupAccountGroupStatus - Status of the account group
// * PENDING_APPROVAL - Account group approval is pending - the account group is visible through our API but cannot be acted on.
// * ACTIVE - Account group is active - full functionality of the Investment API is accessible.
// * CLOSING - Account group is closing.
// * CLOSED - Account group is closed.
// * LOCKED - Account group is locked for all actions.
type RetrieveAccountGroupAccountGroupStatus string

const (
	RetrieveAccountGroupAccountGroupStatusPendingApproval RetrieveAccountGroupAccountGroupStatus = "PENDING_APPROVAL"
	RetrieveAccountGroupAccountGroupStatusActive          RetrieveAccountGroupAccountGroupStatus = "ACTIVE"
	RetrieveAccountGroupAccountGroupStatusClosing         RetrieveAccountGroupAccountGroupStatus = "CLOSING"
	RetrieveAccountGroupAccountGroupStatusClosed          RetrieveAccountGroupAccountGroupStatus = "CLOSED"
	RetrieveAccountGroupAccountGroupStatusLocked          RetrieveAccountGroupAccountGroupStatus = "LOCKED"
)

func (e RetrieveAccountGroupAccountGroupStatus) ToPointer() *RetrieveAccountGroupAccountGroupStatus {
	return &e
}

func (e *RetrieveAccountGroupAccountGroupStatus) UnmarshalJSON(data []byte) error {
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
		*e = RetrieveAccountGroupAccountGroupStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveAccountGroupAccountGroupStatus: %v", v)
	}
}

// RetrieveAccountGroupAccountGroupType - Account group type.
// * PERSONAL - Account group of a person holding assets on their own behalf.
// * LEGAL_ENTITY - Account group of a legal entity holding assets on behalf of their users.
type RetrieveAccountGroupAccountGroupType string

const (
	RetrieveAccountGroupAccountGroupTypePersonal    RetrieveAccountGroupAccountGroupType = "PERSONAL"
	RetrieveAccountGroupAccountGroupTypeLegalEntity RetrieveAccountGroupAccountGroupType = "LEGAL_ENTITY"
)

func (e RetrieveAccountGroupAccountGroupType) ToPointer() *RetrieveAccountGroupAccountGroupType {
	return &e
}

func (e *RetrieveAccountGroupAccountGroupType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PERSONAL":
		fallthrough
	case "LEGAL_ENTITY":
		*e = RetrieveAccountGroupAccountGroupType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveAccountGroupAccountGroupType: %v", v)
	}
}

// RetrieveAccountGroupAccountGroupUsersType - Relation type
// * OWNER -
type RetrieveAccountGroupAccountGroupUsersType string

const (
	RetrieveAccountGroupAccountGroupUsersTypeOwner RetrieveAccountGroupAccountGroupUsersType = "OWNER"
)

func (e RetrieveAccountGroupAccountGroupUsersType) ToPointer() *RetrieveAccountGroupAccountGroupUsersType {
	return &e
}

func (e *RetrieveAccountGroupAccountGroupUsersType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OWNER":
		*e = RetrieveAccountGroupAccountGroupUsersType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveAccountGroupAccountGroupUsersType: %v", v)
	}
}

type RetrieveAccountGroupAccountGroupUsers struct {
	// User unique identifier.
	ID *string `json:"id,omitempty"`
	// Relation type
	// * OWNER -
	Type *RetrieveAccountGroupAccountGroupUsersType `default:"OWNER" json:"type"`
}

func (r RetrieveAccountGroupAccountGroupUsers) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrieveAccountGroupAccountGroupUsers) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrieveAccountGroupAccountGroupUsers) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RetrieveAccountGroupAccountGroupUsers) GetType() *RetrieveAccountGroupAccountGroupUsersType {
	if o == nil {
		return nil
	}
	return o.Type
}

// RetrieveAccountGroupAccountGroup - OK
type RetrieveAccountGroupAccountGroup struct {
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
	Status RetrieveAccountGroupAccountGroupStatus `json:"status"`
	// Account group type.
	// * PERSONAL - Account group of a person holding assets on their own behalf.
	// * LEGAL_ENTITY - Account group of a legal entity holding assets on behalf of their users.
	Type RetrieveAccountGroupAccountGroupType `json:"type"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time                               `json:"updated_at"`
	Users     []RetrieveAccountGroupAccountGroupUsers `json:"users"`
}

func (r RetrieveAccountGroupAccountGroup) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrieveAccountGroupAccountGroup) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrieveAccountGroupAccountGroup) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *RetrieveAccountGroupAccountGroup) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveAccountGroupAccountGroup) GetSecuritiesAccountNumber() string {
	if o == nil {
		return ""
	}
	return o.SecuritiesAccountNumber
}

func (o *RetrieveAccountGroupAccountGroup) GetStatus() RetrieveAccountGroupAccountGroupStatus {
	if o == nil {
		return RetrieveAccountGroupAccountGroupStatus("")
	}
	return o.Status
}

func (o *RetrieveAccountGroupAccountGroup) GetType() RetrieveAccountGroupAccountGroupType {
	if o == nil {
		return RetrieveAccountGroupAccountGroupType("")
	}
	return o.Type
}

func (o *RetrieveAccountGroupAccountGroup) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *RetrieveAccountGroupAccountGroup) GetUsers() []RetrieveAccountGroupAccountGroupUsers {
	if o == nil {
		return []RetrieveAccountGroupAccountGroupUsers{}
	}
	return o.Users
}

type RetrieveAccountGroupResponse struct {
	// OK
	AccountGroup *RetrieveAccountGroupAccountGroup
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RetrieveAccountGroupResponse) GetAccountGroup() *RetrieveAccountGroupAccountGroup {
	if o == nil {
		return nil
	}
	return o.AccountGroup
}

func (o *RetrieveAccountGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RetrieveAccountGroupResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *RetrieveAccountGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RetrieveAccountGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
