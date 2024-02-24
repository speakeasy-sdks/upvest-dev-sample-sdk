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

type RetrieveMandateSecurity struct {
	OauthClientCredentials string `security:"scheme,type=oauth2,name=Authorization"`
}

func (o *RetrieveMandateSecurity) GetOauthClientCredentials() string {
	if o == nil {
		return ""
	}
	return o.OauthClientCredentials
}

type RetrieveMandateRequest struct {
	MandateID string `pathParam:"style=simple,explode=false,name=mandate_id"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-http-header
	Signature string `header:"style=simple,explode=false,name=signature"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-input-http-he
	SignatureInput string `header:"style=simple,explode=false,name=signature-input"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
}

func (r RetrieveMandateRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrieveMandateRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrieveMandateRequest) GetMandateID() string {
	if o == nil {
		return ""
	}
	return o.MandateID
}

func (o *RetrieveMandateRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *RetrieveMandateRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *RetrieveMandateRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *RetrieveMandateRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// RetrieveMandateAddress - Address. Must not be a P.O. box or c/o address.
type RetrieveMandateAddress struct {
	// First address line of the address.
	AddressLine1 string `json:"address_line1"`
	// Second address line of the address.
	AddressLine2 *string `json:"address_line2,omitempty"`
	City         string  `json:"city"`
	// Country code. [ISO 3166 alpha-2 Codes](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country string `json:"country"`
	// Postal code (postcode, PIN or ZIP code)
	Postcode string `json:"postcode"`
	// State, province, county. [ISO 3166 alpha-2 Codes](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	State *string `json:"state,omitempty"`
}

func (o *RetrieveMandateAddress) GetAddressLine1() string {
	if o == nil {
		return ""
	}
	return o.AddressLine1
}

func (o *RetrieveMandateAddress) GetAddressLine2() *string {
	if o == nil {
		return nil
	}
	return o.AddressLine2
}

func (o *RetrieveMandateAddress) GetCity() string {
	if o == nil {
		return ""
	}
	return o.City
}

func (o *RetrieveMandateAddress) GetCountry() string {
	if o == nil {
		return ""
	}
	return o.Country
}

func (o *RetrieveMandateAddress) GetPostcode() string {
	if o == nil {
		return ""
	}
	return o.Postcode
}

func (o *RetrieveMandateAddress) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

// RetrieveMandateType - Type of mandate.
// * RECURRENT -
type RetrieveMandateType string

const (
	RetrieveMandateTypeRecurrent RetrieveMandateType = "RECURRENT"
)

func (e RetrieveMandateType) ToPointer() *RetrieveMandateType {
	return &e
}

func (e *RetrieveMandateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "RECURRENT":
		*e = RetrieveMandateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveMandateType: %v", v)
	}
}

// RetrieveMandateDirectDebitMandate - Mandate
type RetrieveMandateDirectDebitMandate struct {
	// Business Identifier Code (also known as SWIFT-BIC, BIC, SWIFT ID or SWIFT code) [ISO 9362](https://en.wikipedia.org/wiki/ISO_9362).
	Bic string `json:"bic"`
	// Timestamp of when user validated the mandate
	ConfirmedAt time.Time `json:"confirmed_at"`
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	// Address. Must not be a P.O. box or c/o address.
	CreditorAddress RetrieveMandateAddress `json:"creditor_address"`
	// Banking identifier of the creditor.
	CreditorID string `json:"creditor_id"`
	// Name of the creditor on the mandate.
	CreditorName string `json:"creditor_name"`
	// International Bank Account Number [IBAN](https://en.wikipedia.org/wiki/International_Bank_Account_Number).
	Iban string `json:"iban"`
	// Direct Debit Mandate unique identifier.
	ID string `json:"id"`
	// Type of mandate.
	// * RECURRENT -
	Type *RetrieveMandateType `default:"RECURRENT" json:"type"`
	// User unique identifier.
	UserID string `json:"user_id"`
}

func (r RetrieveMandateDirectDebitMandate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrieveMandateDirectDebitMandate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrieveMandateDirectDebitMandate) GetBic() string {
	if o == nil {
		return ""
	}
	return o.Bic
}

func (o *RetrieveMandateDirectDebitMandate) GetConfirmedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.ConfirmedAt
}

func (o *RetrieveMandateDirectDebitMandate) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *RetrieveMandateDirectDebitMandate) GetCreditorAddress() RetrieveMandateAddress {
	if o == nil {
		return RetrieveMandateAddress{}
	}
	return o.CreditorAddress
}

func (o *RetrieveMandateDirectDebitMandate) GetCreditorID() string {
	if o == nil {
		return ""
	}
	return o.CreditorID
}

func (o *RetrieveMandateDirectDebitMandate) GetCreditorName() string {
	if o == nil {
		return ""
	}
	return o.CreditorName
}

func (o *RetrieveMandateDirectDebitMandate) GetIban() string {
	if o == nil {
		return ""
	}
	return o.Iban
}

func (o *RetrieveMandateDirectDebitMandate) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveMandateDirectDebitMandate) GetType() *RetrieveMandateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *RetrieveMandateDirectDebitMandate) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type RetrieveMandateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Mandate
	DirectDebitMandate *RetrieveMandateDirectDebitMandate
	Headers            map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RetrieveMandateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RetrieveMandateResponse) GetDirectDebitMandate() *RetrieveMandateDirectDebitMandate {
	if o == nil {
		return nil
	}
	return o.DirectDebitMandate
}

func (o *RetrieveMandateResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *RetrieveMandateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RetrieveMandateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
