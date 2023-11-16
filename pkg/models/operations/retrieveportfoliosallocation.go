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

type RetrievePortfoliosAllocationRequest struct {
	AllocationID string `pathParam:"style=simple,explode=false,name=allocation_id"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-http-header
	Signature string `header:"style=simple,explode=false,name=signature"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-input-http-he
	SignatureInput string `header:"style=simple,explode=false,name=signature-input"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
}

func (r RetrievePortfoliosAllocationRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrievePortfoliosAllocationRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrievePortfoliosAllocationRequest) GetAllocationID() string {
	if o == nil {
		return ""
	}
	return o.AllocationID
}

func (o *RetrievePortfoliosAllocationRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *RetrievePortfoliosAllocationRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *RetrievePortfoliosAllocationRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *RetrievePortfoliosAllocationRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// RetrievePortfoliosAllocationInstrumentIDType - The type of the ID used in the request.
// * ISIN - International Securities Identification Number
// * UPVEST - UPVEST's unique instrument identifier
type RetrievePortfoliosAllocationInstrumentIDType string

const (
	RetrievePortfoliosAllocationInstrumentIDTypeIsin   RetrievePortfoliosAllocationInstrumentIDType = "ISIN"
	RetrievePortfoliosAllocationInstrumentIDTypeUpvest RetrievePortfoliosAllocationInstrumentIDType = "UPVEST"
)

func (e RetrievePortfoliosAllocationInstrumentIDType) ToPointer() *RetrievePortfoliosAllocationInstrumentIDType {
	return &e
}

func (e *RetrievePortfoliosAllocationInstrumentIDType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ISIN":
		fallthrough
	case "UPVEST":
		*e = RetrievePortfoliosAllocationInstrumentIDType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrievePortfoliosAllocationInstrumentIDType: %v", v)
	}
}

type RetrievePortfoliosAllocationAllocation struct {
	InstrumentID string `json:"instrument_id"`
	// The type of the ID used in the request.
	// * ISIN - International Securities Identification Number
	// * UPVEST - UPVEST's unique instrument identifier
	InstrumentIDType *RetrievePortfoliosAllocationInstrumentIDType `default:"ISIN" json:"instrument_id_type"`
	// Instrument allocation weight
	Weight string `json:"weight"`
}

func (r RetrievePortfoliosAllocationAllocation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrievePortfoliosAllocationAllocation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrievePortfoliosAllocationAllocation) GetInstrumentID() string {
	if o == nil {
		return ""
	}
	return o.InstrumentID
}

func (o *RetrievePortfoliosAllocationAllocation) GetInstrumentIDType() *RetrievePortfoliosAllocationInstrumentIDType {
	if o == nil {
		return nil
	}
	return o.InstrumentIDType
}

func (o *RetrievePortfoliosAllocationAllocation) GetWeight() string {
	if o == nil {
		return ""
	}
	return o.Weight
}

// RetrievePortfoliosAllocationPortfoliosAllocation - Portfolios allocation
type RetrievePortfoliosAllocationPortfoliosAllocation struct {
	// List of portfolios allocations
	Allocation []RetrievePortfoliosAllocationAllocation `json:"allocation"`
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	ID        string    `json:"id"`
	// Allocation name
	Name string `json:"name"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time `json:"updated_at"`
}

func (r RetrievePortfoliosAllocationPortfoliosAllocation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrievePortfoliosAllocationPortfoliosAllocation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrievePortfoliosAllocationPortfoliosAllocation) GetAllocation() []RetrievePortfoliosAllocationAllocation {
	if o == nil {
		return []RetrievePortfoliosAllocationAllocation{}
	}
	return o.Allocation
}

func (o *RetrievePortfoliosAllocationPortfoliosAllocation) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *RetrievePortfoliosAllocationPortfoliosAllocation) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrievePortfoliosAllocationPortfoliosAllocation) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *RetrievePortfoliosAllocationPortfoliosAllocation) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

type RetrievePortfoliosAllocationResponse struct {
	// Portfolios allocation
	TwoHundredApplicationJSONPortfoliosAllocation *RetrievePortfoliosAllocationPortfoliosAllocation
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RetrievePortfoliosAllocationResponse) GetTwoHundredApplicationJSONPortfoliosAllocation() *RetrievePortfoliosAllocationPortfoliosAllocation {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONPortfoliosAllocation
}

func (o *RetrievePortfoliosAllocationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RetrievePortfoliosAllocationResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *RetrievePortfoliosAllocationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RetrievePortfoliosAllocationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
