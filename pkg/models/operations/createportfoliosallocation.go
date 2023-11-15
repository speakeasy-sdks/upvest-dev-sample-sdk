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

// CreatePortfoliosAllocationInstrumentIDType - The type of the ID used in the request.
// * ISIN - International Securities Identification Number
// * UPVEST - UPVEST's unique instrument identifier
type CreatePortfoliosAllocationInstrumentIDType string

const (
	CreatePortfoliosAllocationInstrumentIDTypeIsin   CreatePortfoliosAllocationInstrumentIDType = "ISIN"
	CreatePortfoliosAllocationInstrumentIDTypeUpvest CreatePortfoliosAllocationInstrumentIDType = "UPVEST"
)

func (e CreatePortfoliosAllocationInstrumentIDType) ToPointer() *CreatePortfoliosAllocationInstrumentIDType {
	return &e
}

func (e *CreatePortfoliosAllocationInstrumentIDType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ISIN":
		fallthrough
	case "UPVEST":
		*e = CreatePortfoliosAllocationInstrumentIDType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreatePortfoliosAllocationInstrumentIDType: %v", v)
	}
}

type Allocation struct {
	InstrumentID string `json:"instrument_id"`
	// The type of the ID used in the request.
	// * ISIN - International Securities Identification Number
	// * UPVEST - UPVEST's unique instrument identifier
	InstrumentIDType *CreatePortfoliosAllocationInstrumentIDType `default:"ISIN" json:"instrument_id_type"`
	// Instrument allocation weight
	Weight string `json:"weight"`
}

func (a Allocation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *Allocation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Allocation) GetInstrumentID() string {
	if o == nil {
		return ""
	}
	return o.InstrumentID
}

func (o *Allocation) GetInstrumentIDType() *CreatePortfoliosAllocationInstrumentIDType {
	if o == nil {
		return nil
	}
	return o.InstrumentIDType
}

func (o *Allocation) GetWeight() string {
	if o == nil {
		return ""
	}
	return o.Weight
}

type CreatePortfoliosAllocationPortfoliosAllocationCreateRequest struct {
	// List of portfolios allocations
	Allocation []Allocation `json:"allocation"`
	// Allocation name
	Name *string `json:"name,omitempty"`
}

func (o *CreatePortfoliosAllocationPortfoliosAllocationCreateRequest) GetAllocation() []Allocation {
	if o == nil {
		return []Allocation{}
	}
	return o.Allocation
}

func (o *CreatePortfoliosAllocationPortfoliosAllocationCreateRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type CreatePortfoliosAllocationRequest struct {
	RequestBody *CreatePortfoliosAllocationPortfoliosAllocationCreateRequest `request:"mediaType=application/json"`
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
}

func (c CreatePortfoliosAllocationRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreatePortfoliosAllocationRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreatePortfoliosAllocationRequest) GetRequestBody() *CreatePortfoliosAllocationPortfoliosAllocationCreateRequest {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *CreatePortfoliosAllocationRequest) GetIdempotencyKey() string {
	if o == nil {
		return ""
	}
	return o.IdempotencyKey
}

func (o *CreatePortfoliosAllocationRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *CreatePortfoliosAllocationRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *CreatePortfoliosAllocationRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *CreatePortfoliosAllocationRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// CreatePortfoliosAllocationPortfoliosInstrumentIDType - The type of the ID used in the request.
// * ISIN - International Securities Identification Number
// * UPVEST - UPVEST's unique instrument identifier
type CreatePortfoliosAllocationPortfoliosInstrumentIDType string

const (
	CreatePortfoliosAllocationPortfoliosInstrumentIDTypeIsin   CreatePortfoliosAllocationPortfoliosInstrumentIDType = "ISIN"
	CreatePortfoliosAllocationPortfoliosInstrumentIDTypeUpvest CreatePortfoliosAllocationPortfoliosInstrumentIDType = "UPVEST"
)

func (e CreatePortfoliosAllocationPortfoliosInstrumentIDType) ToPointer() *CreatePortfoliosAllocationPortfoliosInstrumentIDType {
	return &e
}

func (e *CreatePortfoliosAllocationPortfoliosInstrumentIDType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ISIN":
		fallthrough
	case "UPVEST":
		*e = CreatePortfoliosAllocationPortfoliosInstrumentIDType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreatePortfoliosAllocationPortfoliosInstrumentIDType: %v", v)
	}
}

type CreatePortfoliosAllocationAllocation struct {
	InstrumentID string `json:"instrument_id"`
	// The type of the ID used in the request.
	// * ISIN - International Securities Identification Number
	// * UPVEST - UPVEST's unique instrument identifier
	InstrumentIDType *CreatePortfoliosAllocationPortfoliosInstrumentIDType `default:"ISIN" json:"instrument_id_type"`
	// Instrument allocation weight
	Weight string `json:"weight"`
}

func (c CreatePortfoliosAllocationAllocation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreatePortfoliosAllocationAllocation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreatePortfoliosAllocationAllocation) GetInstrumentID() string {
	if o == nil {
		return ""
	}
	return o.InstrumentID
}

func (o *CreatePortfoliosAllocationAllocation) GetInstrumentIDType() *CreatePortfoliosAllocationPortfoliosInstrumentIDType {
	if o == nil {
		return nil
	}
	return o.InstrumentIDType
}

func (o *CreatePortfoliosAllocationAllocation) GetWeight() string {
	if o == nil {
		return ""
	}
	return o.Weight
}

// CreatePortfoliosAllocationPortfoliosAllocation - Portfolios allocation
type CreatePortfoliosAllocationPortfoliosAllocation struct {
	// List of portfolios allocations
	Allocation []CreatePortfoliosAllocationAllocation `json:"allocation"`
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	ID        string    `json:"id"`
	// Allocation name
	Name string `json:"name"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time `json:"updated_at"`
}

func (c CreatePortfoliosAllocationPortfoliosAllocation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreatePortfoliosAllocationPortfoliosAllocation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreatePortfoliosAllocationPortfoliosAllocation) GetAllocation() []CreatePortfoliosAllocationAllocation {
	if o == nil {
		return []CreatePortfoliosAllocationAllocation{}
	}
	return o.Allocation
}

func (o *CreatePortfoliosAllocationPortfoliosAllocation) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *CreatePortfoliosAllocationPortfoliosAllocation) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreatePortfoliosAllocationPortfoliosAllocation) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreatePortfoliosAllocationPortfoliosAllocation) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

type CreatePortfoliosAllocationResponse struct {
	// Portfolios allocation
	TwoHundredApplicationJSONPortfoliosAllocation *CreatePortfoliosAllocationPortfoliosAllocation
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreatePortfoliosAllocationResponse) GetTwoHundredApplicationJSONPortfoliosAllocation() *CreatePortfoliosAllocationPortfoliosAllocation {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONPortfoliosAllocation
}

func (o *CreatePortfoliosAllocationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreatePortfoliosAllocationResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreatePortfoliosAllocationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreatePortfoliosAllocationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
