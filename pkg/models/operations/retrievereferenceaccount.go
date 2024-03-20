// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/utils"
	"net/http"
	"time"
)

type RetrieveReferenceAccountRequest struct {
	ReferenceAccountID string `pathParam:"style=simple,explode=false,name=reference_account_id"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-http-header
	Signature string `header:"style=simple,explode=false,name=signature"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-input-http-he
	SignatureInput string `header:"style=simple,explode=false,name=signature-input"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
}

func (r RetrieveReferenceAccountRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrieveReferenceAccountRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrieveReferenceAccountRequest) GetReferenceAccountID() string {
	if o == nil {
		return ""
	}
	return o.ReferenceAccountID
}

func (o *RetrieveReferenceAccountRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *RetrieveReferenceAccountRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *RetrieveReferenceAccountRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *RetrieveReferenceAccountRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// RetrieveReferenceAccountReferenceAccount - OK
type RetrieveReferenceAccountReferenceAccount struct {
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

func (r RetrieveReferenceAccountReferenceAccount) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrieveReferenceAccountReferenceAccount) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrieveReferenceAccountReferenceAccount) GetAccountOwner() string {
	if o == nil {
		return ""
	}
	return o.AccountOwner
}

func (o *RetrieveReferenceAccountReferenceAccount) GetBic() string {
	if o == nil {
		return ""
	}
	return o.Bic
}

func (o *RetrieveReferenceAccountReferenceAccount) GetConfirmedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.ConfirmedAt
}

func (o *RetrieveReferenceAccountReferenceAccount) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *RetrieveReferenceAccountReferenceAccount) GetIban() string {
	if o == nil {
		return ""
	}
	return o.Iban
}

func (o *RetrieveReferenceAccountReferenceAccount) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveReferenceAccountReferenceAccount) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *RetrieveReferenceAccountReferenceAccount) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *RetrieveReferenceAccountReferenceAccount) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type RetrieveReferenceAccountResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// OK
	ReferenceAccount *RetrieveReferenceAccountReferenceAccount
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RetrieveReferenceAccountResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RetrieveReferenceAccountResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *RetrieveReferenceAccountResponse) GetReferenceAccount() *RetrieveReferenceAccountReferenceAccount {
	if o == nil {
		return nil
	}
	return o.ReferenceAccount
}

func (o *RetrieveReferenceAccountResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RetrieveReferenceAccountResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
