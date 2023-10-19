// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/utils"
	"net/http"
	"time"
)

type ListUserIdentifiersRequest struct {
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

func (l ListUserIdentifiersRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListUserIdentifiersRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListUserIdentifiersRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *ListUserIdentifiersRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *ListUserIdentifiersRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *ListUserIdentifiersRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

func (o *ListUserIdentifiersRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type ListUserIdentifiersIdentifiersListResponseIdentifier struct {
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

func (l ListUserIdentifiersIdentifiersListResponseIdentifier) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListUserIdentifiersIdentifiersListResponseIdentifier) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListUserIdentifiersIdentifiersListResponseIdentifier) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *ListUserIdentifiersIdentifiersListResponseIdentifier) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ListUserIdentifiersIdentifiersListResponseIdentifier) GetIdentifier() string {
	if o == nil {
		return ""
	}
	return o.Identifier
}

func (o *ListUserIdentifiersIdentifiersListResponseIdentifier) GetIdentifierStandard() string {
	if o == nil {
		return ""
	}
	return o.IdentifierStandard
}

func (o *ListUserIdentifiersIdentifiersListResponseIdentifier) GetIssuingCountry() string {
	if o == nil {
		return ""
	}
	return o.IssuingCountry
}

func (o *ListUserIdentifiersIdentifiersListResponseIdentifier) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ListUserIdentifiersIdentifiersListResponseIdentifier) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *ListUserIdentifiersIdentifiersListResponseIdentifier) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

// ListUserIdentifiersIdentifiersListResponse - OK
type ListUserIdentifiersIdentifiersListResponse struct {
	Data []ListUserIdentifiersIdentifiersListResponseIdentifier `json:"data"`
}

func (o *ListUserIdentifiersIdentifiersListResponse) GetData() []ListUserIdentifiersIdentifiersListResponseIdentifier {
	if o == nil {
		return []ListUserIdentifiersIdentifiersListResponseIdentifier{}
	}
	return o.Data
}

type ListUserIdentifiersResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// OK
	IdentifiersListResponse *ListUserIdentifiersIdentifiersListResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListUserIdentifiersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListUserIdentifiersResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *ListUserIdentifiersResponse) GetIdentifiersListResponse() *ListUserIdentifiersIdentifiersListResponse {
	if o == nil {
		return nil
	}
	return o.IdentifiersListResponse
}

func (o *ListUserIdentifiersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListUserIdentifiersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
