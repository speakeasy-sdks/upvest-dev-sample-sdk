// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/types"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/utils"
	"net/http"
	"time"
)

type RetrieveFeeCollectionRequest struct {
	FeeCollectionID string `pathParam:"style=simple,explode=false,name=fee_collection_id"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-http-header
	Signature string `header:"style=simple,explode=false,name=signature"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-input-http-he
	SignatureInput string `header:"style=simple,explode=false,name=signature-input"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
}

func (r RetrieveFeeCollectionRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrieveFeeCollectionRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrieveFeeCollectionRequest) GetFeeCollectionID() string {
	if o == nil {
		return ""
	}
	return o.FeeCollectionID
}

func (o *RetrieveFeeCollectionRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *RetrieveFeeCollectionRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *RetrieveFeeCollectionRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *RetrieveFeeCollectionRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// RetrieveFeeCollectionCurrency - Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
// * EUR - Euro
type RetrieveFeeCollectionCurrency string

const (
	RetrieveFeeCollectionCurrencyEur RetrieveFeeCollectionCurrency = "EUR"
)

func (e RetrieveFeeCollectionCurrency) ToPointer() *RetrieveFeeCollectionCurrency {
	return &e
}

func (e *RetrieveFeeCollectionCurrency) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EUR":
		*e = RetrieveFeeCollectionCurrency(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveFeeCollectionCurrency: %v", v)
	}
}

type RetrieveFeeCollectionProcessedAmount struct {
	CashBalance *string `json:"cash_balance,omitempty"`
	SellToCover *string `json:"sell_to_cover,omitempty"`
}

func (o *RetrieveFeeCollectionProcessedAmount) GetCashBalance() *string {
	if o == nil {
		return nil
	}
	return o.CashBalance
}

func (o *RetrieveFeeCollectionProcessedAmount) GetSellToCover() *string {
	if o == nil {
		return nil
	}
	return o.SellToCover
}

// RetrieveFeeCollectionStatus - Status of the fee collection
// * PROCESSING - Fee collection is in progress.
// * FINALISED - Fees have been collected from the account and the funds has been transferred to the client.
// * CANCELLED - Fee collection has been cancelled.
type RetrieveFeeCollectionStatus string

const (
	RetrieveFeeCollectionStatusProcessing RetrieveFeeCollectionStatus = "PROCESSING"
	RetrieveFeeCollectionStatusFinalised  RetrieveFeeCollectionStatus = "FINALISED"
	RetrieveFeeCollectionStatusCancelled  RetrieveFeeCollectionStatus = "CANCELLED"
)

func (e RetrieveFeeCollectionStatus) ToPointer() *RetrieveFeeCollectionStatus {
	return &e
}

func (e *RetrieveFeeCollectionStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PROCESSING":
		fallthrough
	case "FINALISED":
		fallthrough
	case "CANCELLED":
		*e = RetrieveFeeCollectionStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveFeeCollectionStatus: %v", v)
	}
}

// RetrieveFeeCollectionType - Type of the fee collection
// * SERVICE_FEE - Service fee intake in a pre-defined cadence (e.g. monthly)
// * SERVICE_FEE_LIQUIDATION - Service fee intake as a result of a Portfolio liquidation
type RetrieveFeeCollectionType string

const (
	RetrieveFeeCollectionTypeServiceFee            RetrieveFeeCollectionType = "SERVICE_FEE"
	RetrieveFeeCollectionTypeServiceFeeLiquidation RetrieveFeeCollectionType = "SERVICE_FEE_LIQUIDATION"
)

func (e RetrieveFeeCollectionType) ToPointer() *RetrieveFeeCollectionType {
	return &e
}

func (e *RetrieveFeeCollectionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SERVICE_FEE":
		fallthrough
	case "SERVICE_FEE_LIQUIDATION":
		*e = RetrieveFeeCollectionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveFeeCollectionType: %v", v)
	}
}

// RetrieveFeeCollectionFeeCollection - OK
type RetrieveFeeCollectionFeeCollection struct {
	// Account group unique identifier.
	AccountGroupID string `json:"account_group_id"`
	// Account unique identifier.
	AccountID        string `json:"account_id"`
	CollectionAmount string `json:"collection_amount"`
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	// Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
	// * EUR - Euro
	Currency *RetrieveFeeCollectionCurrency `default:"EUR" json:"currency"`
	// Fee collection unique identifier.
	ID string `json:"id"`
	// End date of the fee collection period in YYYY-MM-DD format. [RFC 3339, section 5.6](https://json-schema.org/draft/2020-12/json-schema-validation.html#RFC3339) RFC 3339
	PeriodEnd types.Date `json:"period_end"`
	// Start date of the fee collection period in YYYY-MM-DD format. [RFC 3339, section 5.6](https://json-schema.org/draft/2020-12/json-schema-validation.html#RFC3339) RFC 3339
	PeriodStart     types.Date                           `json:"period_start"`
	ProcessedAmount RetrieveFeeCollectionProcessedAmount `json:"processed_amount"`
	// Status of the fee collection
	// * PROCESSING - Fee collection is in progress.
	// * FINALISED - Fees have been collected from the account and the funds has been transferred to the client.
	// * CANCELLED - Fee collection has been cancelled.
	Status RetrieveFeeCollectionStatus `json:"status"`
	// Type of the fee collection
	// * SERVICE_FEE - Service fee intake in a pre-defined cadence (e.g. monthly)
	// * SERVICE_FEE_LIQUIDATION - Service fee intake as a result of a Portfolio liquidation
	Type RetrieveFeeCollectionType `json:"type"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time `json:"updated_at"`
}

func (r RetrieveFeeCollectionFeeCollection) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrieveFeeCollectionFeeCollection) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrieveFeeCollectionFeeCollection) GetAccountGroupID() string {
	if o == nil {
		return ""
	}
	return o.AccountGroupID
}

func (o *RetrieveFeeCollectionFeeCollection) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *RetrieveFeeCollectionFeeCollection) GetCollectionAmount() string {
	if o == nil {
		return ""
	}
	return o.CollectionAmount
}

func (o *RetrieveFeeCollectionFeeCollection) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *RetrieveFeeCollectionFeeCollection) GetCurrency() *RetrieveFeeCollectionCurrency {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *RetrieveFeeCollectionFeeCollection) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveFeeCollectionFeeCollection) GetPeriodEnd() types.Date {
	if o == nil {
		return types.Date{}
	}
	return o.PeriodEnd
}

func (o *RetrieveFeeCollectionFeeCollection) GetPeriodStart() types.Date {
	if o == nil {
		return types.Date{}
	}
	return o.PeriodStart
}

func (o *RetrieveFeeCollectionFeeCollection) GetProcessedAmount() RetrieveFeeCollectionProcessedAmount {
	if o == nil {
		return RetrieveFeeCollectionProcessedAmount{}
	}
	return o.ProcessedAmount
}

func (o *RetrieveFeeCollectionFeeCollection) GetStatus() RetrieveFeeCollectionStatus {
	if o == nil {
		return RetrieveFeeCollectionStatus("")
	}
	return o.Status
}

func (o *RetrieveFeeCollectionFeeCollection) GetType() RetrieveFeeCollectionType {
	if o == nil {
		return RetrieveFeeCollectionType("")
	}
	return o.Type
}

func (o *RetrieveFeeCollectionFeeCollection) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

type RetrieveFeeCollectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	FeeCollection *RetrieveFeeCollectionFeeCollection
	Headers       map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RetrieveFeeCollectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RetrieveFeeCollectionResponse) GetFeeCollection() *RetrieveFeeCollectionFeeCollection {
	if o == nil {
		return nil
	}
	return o.FeeCollection
}

func (o *RetrieveFeeCollectionResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *RetrieveFeeCollectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RetrieveFeeCollectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
