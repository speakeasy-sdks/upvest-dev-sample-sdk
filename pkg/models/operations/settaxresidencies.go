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

// MissingTinReason - Reason why TIN is missing
// * TIN_NOT_YET_ASSIGNED - Indicates that the tax identification number has not yet been assigned by the tax authorities. A common example is, that a user has moved to a country and thus became taxable, but that the tax authorities have not yet assigned the TIN to this user.
// * COUNTRY_HAS_NO_TIN - Indicates that the specific country does not provide a TIN.
// * OTHER_REASONS - Applies in case of other reasons - i.e. when a user does not have the TIN at hand. Note this may cause additional inquiries by our customer service team.
type MissingTinReason string

const (
	MissingTinReasonTinNotYetAssigned MissingTinReason = "TIN_NOT_YET_ASSIGNED"
	MissingTinReasonCountryHasNoTin   MissingTinReason = "COUNTRY_HAS_NO_TIN"
	MissingTinReasonOtherReasons      MissingTinReason = "OTHER_REASONS"
)

func (e MissingTinReason) ToPointer() *MissingTinReason {
	return &e
}

func (e *MissingTinReason) UnmarshalJSON(data []byte) error {
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
		*e = MissingTinReason(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MissingTinReason: %v", v)
	}
}

type WithoutTaxIdentifierNumber struct {
	// Country code. [ISO 3166 alpha-2 Codes](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country string `json:"country"`
	// Reason why TIN is missing
	// * TIN_NOT_YET_ASSIGNED - Indicates that the tax identification number has not yet been assigned by the tax authorities. A common example is, that a user has moved to a country and thus became taxable, but that the tax authorities have not yet assigned the TIN to this user.
	// * COUNTRY_HAS_NO_TIN - Indicates that the specific country does not provide a TIN.
	// * OTHER_REASONS - Applies in case of other reasons - i.e. when a user does not have the TIN at hand. Note this may cause additional inquiries by our customer service team.
	MissingTinReason MissingTinReason `json:"missing_tin_reason"`
}

func (o *WithoutTaxIdentifierNumber) GetCountry() string {
	if o == nil {
		return ""
	}
	return o.Country
}

func (o *WithoutTaxIdentifierNumber) GetMissingTinReason() MissingTinReason {
	if o == nil {
		return MissingTinReason("")
	}
	return o.MissingTinReason
}

type WithTaxIdentifierNumber struct {
	// Country code. [ISO 3166 alpha-2 Codes](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country string `json:"country"`
	// Tax identifier number
	TaxIdentifierNumber string `json:"tax_identifier_number"`
}

func (o *WithTaxIdentifierNumber) GetCountry() string {
	if o == nil {
		return ""
	}
	return o.Country
}

func (o *WithTaxIdentifierNumber) GetTaxIdentifierNumber() string {
	if o == nil {
		return ""
	}
	return o.TaxIdentifierNumber
}

type TaxResidencyForCreateRequestType string

const (
	TaxResidencyForCreateRequestTypeWithTaxIdentifierNumber    TaxResidencyForCreateRequestType = "With tax identifier number"
	TaxResidencyForCreateRequestTypeWithoutTaxIdentifierNumber TaxResidencyForCreateRequestType = "Without tax identifier number"
)

type TaxResidencyForCreateRequest struct {
	WithTaxIdentifierNumber    *WithTaxIdentifierNumber
	WithoutTaxIdentifierNumber *WithoutTaxIdentifierNumber

	Type TaxResidencyForCreateRequestType
}

func CreateTaxResidencyForCreateRequestWithTaxIdentifierNumber(withTaxIdentifierNumber WithTaxIdentifierNumber) TaxResidencyForCreateRequest {
	typ := TaxResidencyForCreateRequestTypeWithTaxIdentifierNumber

	return TaxResidencyForCreateRequest{
		WithTaxIdentifierNumber: &withTaxIdentifierNumber,
		Type:                    typ,
	}
}

func CreateTaxResidencyForCreateRequestWithoutTaxIdentifierNumber(withoutTaxIdentifierNumber WithoutTaxIdentifierNumber) TaxResidencyForCreateRequest {
	typ := TaxResidencyForCreateRequestTypeWithoutTaxIdentifierNumber

	return TaxResidencyForCreateRequest{
		WithoutTaxIdentifierNumber: &withoutTaxIdentifierNumber,
		Type:                       typ,
	}
}

func (u *TaxResidencyForCreateRequest) UnmarshalJSON(data []byte) error {

	withTaxIdentifierNumber := WithTaxIdentifierNumber{}
	if err := utils.UnmarshalJSON(data, &withTaxIdentifierNumber, "", true, true); err == nil {
		u.WithTaxIdentifierNumber = &withTaxIdentifierNumber
		u.Type = TaxResidencyForCreateRequestTypeWithTaxIdentifierNumber
		return nil
	}

	withoutTaxIdentifierNumber := WithoutTaxIdentifierNumber{}
	if err := utils.UnmarshalJSON(data, &withoutTaxIdentifierNumber, "", true, true); err == nil {
		u.WithoutTaxIdentifierNumber = &withoutTaxIdentifierNumber
		u.Type = TaxResidencyForCreateRequestTypeWithoutTaxIdentifierNumber
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u TaxResidencyForCreateRequest) MarshalJSON() ([]byte, error) {
	if u.WithTaxIdentifierNumber != nil {
		return utils.MarshalJSON(u.WithTaxIdentifierNumber, "", true)
	}

	if u.WithoutTaxIdentifierNumber != nil {
		return utils.MarshalJSON(u.WithoutTaxIdentifierNumber, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type SetTaxResidenciesTaxResidenciesSetRequest struct {
	TaxResidencies []TaxResidencyForCreateRequest `json:"tax_residencies"`
}

func (o *SetTaxResidenciesTaxResidenciesSetRequest) GetTaxResidencies() []TaxResidencyForCreateRequest {
	if o == nil {
		return []TaxResidencyForCreateRequest{}
	}
	return o.TaxResidencies
}

type SetTaxResidenciesRequest struct {
	RequestBody *SetTaxResidenciesTaxResidenciesSetRequest `request:"mediaType=application/json"`
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
	UserID         string `pathParam:"style=simple,explode=false,name=user_id"`
}

func (s SetTaxResidenciesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SetTaxResidenciesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SetTaxResidenciesRequest) GetRequestBody() *SetTaxResidenciesTaxResidenciesSetRequest {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *SetTaxResidenciesRequest) GetIdempotencyKey() string {
	if o == nil {
		return ""
	}
	return o.IdempotencyKey
}

func (o *SetTaxResidenciesRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *SetTaxResidenciesRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *SetTaxResidenciesRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *SetTaxResidenciesRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

func (o *SetTaxResidenciesRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

// SetTaxResidenciesStatus - Tax residency status
// * PENDING - It indicates that the tax residency records are not yet processed by Upvest.
// * ACTIVE - It indicates that tax residency records are processed, and the tax residency record is the one in use.
type SetTaxResidenciesStatus string

const (
	SetTaxResidenciesStatusPending SetTaxResidenciesStatus = "PENDING"
	SetTaxResidenciesStatusActive  SetTaxResidenciesStatus = "ACTIVE"
)

func (e SetTaxResidenciesStatus) ToPointer() *SetTaxResidenciesStatus {
	return &e
}

func (e *SetTaxResidenciesStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PENDING":
		fallthrough
	case "ACTIVE":
		*e = SetTaxResidenciesStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SetTaxResidenciesStatus: %v", v)
	}
}

// SetTaxResidenciesMissingTinReason - Reason why TIN is missing
// * TIN_NOT_YET_ASSIGNED - Indicates that the tax identification number has not yet been assigned by the tax authorities. A common example is, that a user has moved to a country and thus became taxable, but that the tax authorities have not yet assigned the TIN to this user.
// * COUNTRY_HAS_NO_TIN - Indicates that the specific country does not provide a TIN.
// * OTHER_REASONS - Applies in case of other reasons - i.e. when a user does not have the TIN at hand. Note this may cause additional inquiries by our customer service team.
type SetTaxResidenciesMissingTinReason string

const (
	SetTaxResidenciesMissingTinReasonTinNotYetAssigned SetTaxResidenciesMissingTinReason = "TIN_NOT_YET_ASSIGNED"
	SetTaxResidenciesMissingTinReasonCountryHasNoTin   SetTaxResidenciesMissingTinReason = "COUNTRY_HAS_NO_TIN"
	SetTaxResidenciesMissingTinReasonOtherReasons      SetTaxResidenciesMissingTinReason = "OTHER_REASONS"
)

func (e SetTaxResidenciesMissingTinReason) ToPointer() *SetTaxResidenciesMissingTinReason {
	return &e
}

func (e *SetTaxResidenciesMissingTinReason) UnmarshalJSON(data []byte) error {
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
		*e = SetTaxResidenciesMissingTinReason(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SetTaxResidenciesMissingTinReason: %v", v)
	}
}

type SetTaxResidenciesWithoutTaxIdentifierNumber struct {
	// Country code. [ISO 3166 alpha-2 Codes](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country string `json:"country"`
	// Reason why TIN is missing
	// * TIN_NOT_YET_ASSIGNED - Indicates that the tax identification number has not yet been assigned by the tax authorities. A common example is, that a user has moved to a country and thus became taxable, but that the tax authorities have not yet assigned the TIN to this user.
	// * COUNTRY_HAS_NO_TIN - Indicates that the specific country does not provide a TIN.
	// * OTHER_REASONS - Applies in case of other reasons - i.e. when a user does not have the TIN at hand. Note this may cause additional inquiries by our customer service team.
	MissingTinReason SetTaxResidenciesMissingTinReason `json:"missing_tin_reason"`
}

func (o *SetTaxResidenciesWithoutTaxIdentifierNumber) GetCountry() string {
	if o == nil {
		return ""
	}
	return o.Country
}

func (o *SetTaxResidenciesWithoutTaxIdentifierNumber) GetMissingTinReason() SetTaxResidenciesMissingTinReason {
	if o == nil {
		return SetTaxResidenciesMissingTinReason("")
	}
	return o.MissingTinReason
}

type SetTaxResidenciesWithTaxIdentifierNumber struct {
	// Country code. [ISO 3166 alpha-2 Codes](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country string `json:"country"`
	// Tax identifier number
	TaxIdentifierNumber string `json:"tax_identifier_number"`
}

func (o *SetTaxResidenciesWithTaxIdentifierNumber) GetCountry() string {
	if o == nil {
		return ""
	}
	return o.Country
}

func (o *SetTaxResidenciesWithTaxIdentifierNumber) GetTaxIdentifierNumber() string {
	if o == nil {
		return ""
	}
	return o.TaxIdentifierNumber
}

type SetTaxResidenciesTaxResidencyType string

const (
	SetTaxResidenciesTaxResidencyTypeSetTaxResidenciesWithTaxIdentifierNumber    SetTaxResidenciesTaxResidencyType = "set_tax_residencies_With tax identifier number"
	SetTaxResidenciesTaxResidencyTypeSetTaxResidenciesWithoutTaxIdentifierNumber SetTaxResidenciesTaxResidencyType = "set_tax_residencies_Without tax identifier number"
)

type SetTaxResidenciesTaxResidency struct {
	SetTaxResidenciesWithTaxIdentifierNumber    *SetTaxResidenciesWithTaxIdentifierNumber
	SetTaxResidenciesWithoutTaxIdentifierNumber *SetTaxResidenciesWithoutTaxIdentifierNumber

	Type SetTaxResidenciesTaxResidencyType
}

func CreateSetTaxResidenciesTaxResidencySetTaxResidenciesWithTaxIdentifierNumber(setTaxResidenciesWithTaxIdentifierNumber SetTaxResidenciesWithTaxIdentifierNumber) SetTaxResidenciesTaxResidency {
	typ := SetTaxResidenciesTaxResidencyTypeSetTaxResidenciesWithTaxIdentifierNumber

	return SetTaxResidenciesTaxResidency{
		SetTaxResidenciesWithTaxIdentifierNumber: &setTaxResidenciesWithTaxIdentifierNumber,
		Type:                                     typ,
	}
}

func CreateSetTaxResidenciesTaxResidencySetTaxResidenciesWithoutTaxIdentifierNumber(setTaxResidenciesWithoutTaxIdentifierNumber SetTaxResidenciesWithoutTaxIdentifierNumber) SetTaxResidenciesTaxResidency {
	typ := SetTaxResidenciesTaxResidencyTypeSetTaxResidenciesWithoutTaxIdentifierNumber

	return SetTaxResidenciesTaxResidency{
		SetTaxResidenciesWithoutTaxIdentifierNumber: &setTaxResidenciesWithoutTaxIdentifierNumber,
		Type: typ,
	}
}

func (u *SetTaxResidenciesTaxResidency) UnmarshalJSON(data []byte) error {

	setTaxResidenciesWithTaxIdentifierNumber := SetTaxResidenciesWithTaxIdentifierNumber{}
	if err := utils.UnmarshalJSON(data, &setTaxResidenciesWithTaxIdentifierNumber, "", true, true); err == nil {
		u.SetTaxResidenciesWithTaxIdentifierNumber = &setTaxResidenciesWithTaxIdentifierNumber
		u.Type = SetTaxResidenciesTaxResidencyTypeSetTaxResidenciesWithTaxIdentifierNumber
		return nil
	}

	setTaxResidenciesWithoutTaxIdentifierNumber := SetTaxResidenciesWithoutTaxIdentifierNumber{}
	if err := utils.UnmarshalJSON(data, &setTaxResidenciesWithoutTaxIdentifierNumber, "", true, true); err == nil {
		u.SetTaxResidenciesWithoutTaxIdentifierNumber = &setTaxResidenciesWithoutTaxIdentifierNumber
		u.Type = SetTaxResidenciesTaxResidencyTypeSetTaxResidenciesWithoutTaxIdentifierNumber
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SetTaxResidenciesTaxResidency) MarshalJSON() ([]byte, error) {
	if u.SetTaxResidenciesWithTaxIdentifierNumber != nil {
		return utils.MarshalJSON(u.SetTaxResidenciesWithTaxIdentifierNumber, "", true)
	}

	if u.SetTaxResidenciesWithoutTaxIdentifierNumber != nil {
		return utils.MarshalJSON(u.SetTaxResidenciesWithoutTaxIdentifierNumber, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// SetTaxResidenciesTaxResidencyRecord - User tax residencies
type SetTaxResidenciesTaxResidencyRecord struct {
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	// Tax residency status
	// * PENDING - It indicates that the tax residency records are not yet processed by Upvest.
	// * ACTIVE - It indicates that tax residency records are processed, and the tax residency record is the one in use.
	Status         SetTaxResidenciesStatus         `json:"status"`
	TaxResidencies []SetTaxResidenciesTaxResidency `json:"tax_residencies"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time `json:"updated_at"`
}

func (s SetTaxResidenciesTaxResidencyRecord) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SetTaxResidenciesTaxResidencyRecord) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SetTaxResidenciesTaxResidencyRecord) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *SetTaxResidenciesTaxResidencyRecord) GetStatus() SetTaxResidenciesStatus {
	if o == nil {
		return SetTaxResidenciesStatus("")
	}
	return o.Status
}

func (o *SetTaxResidenciesTaxResidencyRecord) GetTaxResidencies() []SetTaxResidenciesTaxResidency {
	if o == nil {
		return []SetTaxResidenciesTaxResidency{}
	}
	return o.TaxResidencies
}

func (o *SetTaxResidenciesTaxResidencyRecord) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

type SetTaxResidenciesResponse struct {
	// User tax residencies
	TwoHundredApplicationJSONTaxResidencyRecord *SetTaxResidenciesTaxResidencyRecord
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *SetTaxResidenciesResponse) GetTwoHundredApplicationJSONTaxResidencyRecord() *SetTaxResidenciesTaxResidencyRecord {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONTaxResidencyRecord
}

func (o *SetTaxResidenciesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SetTaxResidenciesResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *SetTaxResidenciesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SetTaxResidenciesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
