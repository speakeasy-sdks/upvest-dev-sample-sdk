// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/utils"
	"net/http"
	"time"
)

type CreateIdentifierSecurity struct {
	OauthClientCredentials string `security:"scheme,type=oauth2,name=Authorization"`
}

func (o *CreateIdentifierSecurity) GetOauthClientCredentials() string {
	if o == nil {
		return ""
	}
	return o.OauthClientCredentials
}

type CreateIdentifierIdentifierCreateRequest struct {
	// Identifier value. See [guide here](/guides/users_and_accounts/identifiers.md)
	Identifier string `json:"identifier"`
	// Country code. [ISO 3166 alpha-2 Codes](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	IssuingCountry string `json:"issuing_country"`
	// Identifier type.
	// * NATIONAL_ID -
	Type *string `default:"NATIONAL_ID" json:"type"`
}

func (c CreateIdentifierIdentifierCreateRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateIdentifierIdentifierCreateRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateIdentifierIdentifierCreateRequest) GetIdentifier() string {
	if o == nil {
		return ""
	}
	return o.Identifier
}

func (o *CreateIdentifierIdentifierCreateRequest) GetIssuingCountry() string {
	if o == nil {
		return ""
	}
	return o.IssuingCountry
}

func (o *CreateIdentifierIdentifierCreateRequest) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

type CreateIdentifierRequest struct {
	RequestBody *CreateIdentifierIdentifierCreateRequest `request:"mediaType=application/json"`
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

func (c CreateIdentifierRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateIdentifierRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateIdentifierRequest) GetRequestBody() *CreateIdentifierIdentifierCreateRequest {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *CreateIdentifierRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *CreateIdentifierRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *CreateIdentifierRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *CreateIdentifierRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

func (o *CreateIdentifierRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

// CreateIdentifierIdentifier - The user identifier is created.
type CreateIdentifierIdentifier struct {
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

func (c CreateIdentifierIdentifier) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateIdentifierIdentifier) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateIdentifierIdentifier) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *CreateIdentifierIdentifier) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateIdentifierIdentifier) GetIdentifier() string {
	if o == nil {
		return ""
	}
	return o.Identifier
}

func (o *CreateIdentifierIdentifier) GetIdentifierStandard() string {
	if o == nil {
		return ""
	}
	return o.IdentifierStandard
}

func (o *CreateIdentifierIdentifier) GetIssuingCountry() string {
	if o == nil {
		return ""
	}
	return o.IssuingCountry
}

func (o *CreateIdentifierIdentifier) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *CreateIdentifierIdentifier) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *CreateIdentifierIdentifier) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type CreateIdentifierResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// The user identifier is created.
	Identifier *CreateIdentifierIdentifier
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateIdentifierResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateIdentifierResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *CreateIdentifierResponse) GetIdentifier() *CreateIdentifierIdentifier {
	if o == nil {
		return nil
	}
	return o.Identifier
}

func (o *CreateIdentifierResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateIdentifierResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
