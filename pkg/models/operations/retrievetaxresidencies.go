// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/utils"
	"net/http"
	"time"
)

type RetrieveTaxResidenciesRequest struct {
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

func (r RetrieveTaxResidenciesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrieveTaxResidenciesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrieveTaxResidenciesRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *RetrieveTaxResidenciesRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *RetrieveTaxResidenciesRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *RetrieveTaxResidenciesRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

func (o *RetrieveTaxResidenciesRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

// RetrieveTaxResidenciesStatus - Tax residency status
// * PENDING - It indicates that the tax residency records are not yet processed by Upvest.
// * ACTIVE - It indicates that tax residency records are processed, and the tax residency record is the one in use.
type RetrieveTaxResidenciesStatus string

const (
	RetrieveTaxResidenciesStatusPending RetrieveTaxResidenciesStatus = "PENDING"
	RetrieveTaxResidenciesStatusActive  RetrieveTaxResidenciesStatus = "ACTIVE"
)

func (e RetrieveTaxResidenciesStatus) ToPointer() *RetrieveTaxResidenciesStatus {
	return &e
}

func (e *RetrieveTaxResidenciesStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PENDING":
		fallthrough
	case "ACTIVE":
		*e = RetrieveTaxResidenciesStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveTaxResidenciesStatus: %v", v)
	}
}

// RetrieveTaxResidenciesMissingTinReason - Reason why TIN is missing
// * TIN_NOT_YET_ASSIGNED - Indicates that the tax identification number has not yet been assigned by the tax authorities. A common example is, that a user has moved to a country and thus became taxable, but that the tax authorities have not yet assigned the TIN to this user.
// * COUNTRY_HAS_NO_TIN - Indicates that the specific country does not provide a TIN.
// * OTHER_REASONS - Applies in case of other reasons - i.e. when a user does not have the TIN at hand. Note this may cause additional inquiries by our customer service team.
type RetrieveTaxResidenciesMissingTinReason string

const (
	RetrieveTaxResidenciesMissingTinReasonTinNotYetAssigned RetrieveTaxResidenciesMissingTinReason = "TIN_NOT_YET_ASSIGNED"
	RetrieveTaxResidenciesMissingTinReasonCountryHasNoTin   RetrieveTaxResidenciesMissingTinReason = "COUNTRY_HAS_NO_TIN"
	RetrieveTaxResidenciesMissingTinReasonOtherReasons      RetrieveTaxResidenciesMissingTinReason = "OTHER_REASONS"
)

func (e RetrieveTaxResidenciesMissingTinReason) ToPointer() *RetrieveTaxResidenciesMissingTinReason {
	return &e
}

func (e *RetrieveTaxResidenciesMissingTinReason) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TIN_NOT_YET_ASSIGNED":
		fallthrough
	case "COUNTRY_HAS_NO_TIN":
		fallthrough
	case "OTHER_REASONS":
		*e = RetrieveTaxResidenciesMissingTinReason(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveTaxResidenciesMissingTinReason: %v", v)
	}
}

type RetrieveTaxResidenciesWithoutTaxIdentifierNumber struct {
	// Country code. [ISO 3166 alpha-2 Codes](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country string `json:"country"`
	// Reason why TIN is missing
	// * TIN_NOT_YET_ASSIGNED - Indicates that the tax identification number has not yet been assigned by the tax authorities. A common example is, that a user has moved to a country and thus became taxable, but that the tax authorities have not yet assigned the TIN to this user.
	// * COUNTRY_HAS_NO_TIN - Indicates that the specific country does not provide a TIN.
	// * OTHER_REASONS - Applies in case of other reasons - i.e. when a user does not have the TIN at hand. Note this may cause additional inquiries by our customer service team.
	MissingTinReason RetrieveTaxResidenciesMissingTinReason `json:"missing_tin_reason"`
}

func (o *RetrieveTaxResidenciesWithoutTaxIdentifierNumber) GetCountry() string {
	if o == nil {
		return ""
	}
	return o.Country
}

func (o *RetrieveTaxResidenciesWithoutTaxIdentifierNumber) GetMissingTinReason() RetrieveTaxResidenciesMissingTinReason {
	if o == nil {
		return RetrieveTaxResidenciesMissingTinReason("")
	}
	return o.MissingTinReason
}

type RetrieveTaxResidenciesWithTaxIdentifierNumber struct {
	// Country code. [ISO 3166 alpha-2 Codes](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country string `json:"country"`
	// Tax identifier number
	TaxIdentifierNumber string `json:"tax_identifier_number"`
}

func (o *RetrieveTaxResidenciesWithTaxIdentifierNumber) GetCountry() string {
	if o == nil {
		return ""
	}
	return o.Country
}

func (o *RetrieveTaxResidenciesWithTaxIdentifierNumber) GetTaxIdentifierNumber() string {
	if o == nil {
		return ""
	}
	return o.TaxIdentifierNumber
}

type TaxResidencyType string

const (
	TaxResidencyTypeRetrieveTaxResidenciesWithTaxIdentifierNumber    TaxResidencyType = "retrieve_tax_residencies_With tax identifier number"
	TaxResidencyTypeRetrieveTaxResidenciesWithoutTaxIdentifierNumber TaxResidencyType = "retrieve_tax_residencies_Without tax identifier number"
)

type TaxResidency struct {
	RetrieveTaxResidenciesWithTaxIdentifierNumber    *RetrieveTaxResidenciesWithTaxIdentifierNumber
	RetrieveTaxResidenciesWithoutTaxIdentifierNumber *RetrieveTaxResidenciesWithoutTaxIdentifierNumber

	Type TaxResidencyType
}

func CreateTaxResidencyRetrieveTaxResidenciesWithTaxIdentifierNumber(retrieveTaxResidenciesWithTaxIdentifierNumber RetrieveTaxResidenciesWithTaxIdentifierNumber) TaxResidency {
	typ := TaxResidencyTypeRetrieveTaxResidenciesWithTaxIdentifierNumber

	return TaxResidency{
		RetrieveTaxResidenciesWithTaxIdentifierNumber: &retrieveTaxResidenciesWithTaxIdentifierNumber,
		Type: typ,
	}
}

func CreateTaxResidencyRetrieveTaxResidenciesWithoutTaxIdentifierNumber(retrieveTaxResidenciesWithoutTaxIdentifierNumber RetrieveTaxResidenciesWithoutTaxIdentifierNumber) TaxResidency {
	typ := TaxResidencyTypeRetrieveTaxResidenciesWithoutTaxIdentifierNumber

	return TaxResidency{
		RetrieveTaxResidenciesWithoutTaxIdentifierNumber: &retrieveTaxResidenciesWithoutTaxIdentifierNumber,
		Type: typ,
	}
}

func (u *TaxResidency) UnmarshalJSON(data []byte) error {

	retrieveTaxResidenciesWithTaxIdentifierNumber := RetrieveTaxResidenciesWithTaxIdentifierNumber{}
	if err := utils.UnmarshalJSON(data, &retrieveTaxResidenciesWithTaxIdentifierNumber, "", true, true); err == nil {
		u.RetrieveTaxResidenciesWithTaxIdentifierNumber = &retrieveTaxResidenciesWithTaxIdentifierNumber
		u.Type = TaxResidencyTypeRetrieveTaxResidenciesWithTaxIdentifierNumber
		return nil
	}

	retrieveTaxResidenciesWithoutTaxIdentifierNumber := RetrieveTaxResidenciesWithoutTaxIdentifierNumber{}
	if err := utils.UnmarshalJSON(data, &retrieveTaxResidenciesWithoutTaxIdentifierNumber, "", true, true); err == nil {
		u.RetrieveTaxResidenciesWithoutTaxIdentifierNumber = &retrieveTaxResidenciesWithoutTaxIdentifierNumber
		u.Type = TaxResidencyTypeRetrieveTaxResidenciesWithoutTaxIdentifierNumber
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u TaxResidency) MarshalJSON() ([]byte, error) {
	if u.RetrieveTaxResidenciesWithTaxIdentifierNumber != nil {
		return utils.MarshalJSON(u.RetrieveTaxResidenciesWithTaxIdentifierNumber, "", true)
	}

	if u.RetrieveTaxResidenciesWithoutTaxIdentifierNumber != nil {
		return utils.MarshalJSON(u.RetrieveTaxResidenciesWithoutTaxIdentifierNumber, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// RetrieveTaxResidenciesTaxResidencyRecord - User tax residencies
type RetrieveTaxResidenciesTaxResidencyRecord struct {
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	// Tax residency status
	// * PENDING - It indicates that the tax residency records are not yet processed by Upvest.
	// * ACTIVE - It indicates that tax residency records are processed, and the tax residency record is the one in use.
	Status         RetrieveTaxResidenciesStatus `json:"status"`
	TaxResidencies []TaxResidency               `json:"tax_residencies"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time `json:"updated_at"`
}

func (r RetrieveTaxResidenciesTaxResidencyRecord) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrieveTaxResidenciesTaxResidencyRecord) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrieveTaxResidenciesTaxResidencyRecord) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *RetrieveTaxResidenciesTaxResidencyRecord) GetStatus() RetrieveTaxResidenciesStatus {
	if o == nil {
		return RetrieveTaxResidenciesStatus("")
	}
	return o.Status
}

func (o *RetrieveTaxResidenciesTaxResidencyRecord) GetTaxResidencies() []TaxResidency {
	if o == nil {
		return []TaxResidency{}
	}
	return o.TaxResidencies
}

func (o *RetrieveTaxResidenciesTaxResidencyRecord) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

type RetrieveTaxResidenciesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// User tax residencies
	TaxResidencyRecord *RetrieveTaxResidenciesTaxResidencyRecord
}

func (o *RetrieveTaxResidenciesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RetrieveTaxResidenciesResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *RetrieveTaxResidenciesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RetrieveTaxResidenciesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RetrieveTaxResidenciesResponse) GetTaxResidencyRecord() *RetrieveTaxResidenciesTaxResidencyRecord {
	if o == nil {
		return nil
	}
	return o.TaxResidencyRecord
}
