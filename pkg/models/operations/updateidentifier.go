// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/utils"
	"net/http"
	"time"
)

type UpdateIdentifierIdentifierUpdateRequest struct {
	// Identifier value. See [guide here](/guides/users_and_accounts/identifiers.md)
	Identifier string `json:"identifier"`
}

func (o *UpdateIdentifierIdentifierUpdateRequest) GetIdentifier() string {
	if o == nil {
		return ""
	}
	return o.Identifier
}

type UpdateIdentifierRequest struct {
	RequestBody  *UpdateIdentifierIdentifierUpdateRequest `request:"mediaType=application/json"`
	IdentifierID string                                   `pathParam:"style=simple,explode=false,name=identifier_id"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-http-header
	Signature string `header:"style=simple,explode=false,name=signature"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-input-http-he
	SignatureInput string `header:"style=simple,explode=false,name=signature-input"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
	UserID         string `pathParam:"style=simple,explode=false,name=user_id"`
}

func (u UpdateIdentifierRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateIdentifierRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateIdentifierRequest) GetRequestBody() *UpdateIdentifierIdentifierUpdateRequest {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *UpdateIdentifierRequest) GetIdentifierID() string {
	if o == nil {
		return ""
	}
	return o.IdentifierID
}

func (o *UpdateIdentifierRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *UpdateIdentifierRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *UpdateIdentifierRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *UpdateIdentifierRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

func (o *UpdateIdentifierRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

// UpdateIdentifierIdentifier - The user identifier is updated.
type UpdateIdentifierIdentifier struct {
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	// Unique identifier for the user's national ID.
	ID string `json:"id"`
	// Identifier value. See [guide here](/guides/users_and_accounts/identifiers.md)
	Identifier string `json:"identifier"`
	// Identifier standard abbreviation. See [guide here](/guides/users_and_accounts/identifiers.md)
	IdentifierStandard string `json:"identifier_standard"`
	// Country code. [ISO 3166 alpha-2 Codes](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	IssuingCountry string `json:"issuing_country"`
	// Identifier type.
	// * NATIONAL_ID -
	Type *string `default:"NATIONAL_ID" json:"type"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time `json:"updated_at"`
	// User unique identifier.
	UserID string `json:"user_id"`
}

func (u UpdateIdentifierIdentifier) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateIdentifierIdentifier) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateIdentifierIdentifier) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *UpdateIdentifierIdentifier) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateIdentifierIdentifier) GetIdentifier() string {
	if o == nil {
		return ""
	}
	return o.Identifier
}

func (o *UpdateIdentifierIdentifier) GetIdentifierStandard() string {
	if o == nil {
		return ""
	}
	return o.IdentifierStandard
}

func (o *UpdateIdentifierIdentifier) GetIssuingCountry() string {
	if o == nil {
		return ""
	}
	return o.IssuingCountry
}

func (o *UpdateIdentifierIdentifier) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *UpdateIdentifierIdentifier) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *UpdateIdentifierIdentifier) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type UpdateIdentifierResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// The user identifier is updated.
	Identifier *UpdateIdentifierIdentifier
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateIdentifierResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateIdentifierResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *UpdateIdentifierResponse) GetIdentifier() *UpdateIdentifierIdentifier {
	if o == nil {
		return nil
	}
	return o.Identifier
}

func (o *UpdateIdentifierResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateIdentifierResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
