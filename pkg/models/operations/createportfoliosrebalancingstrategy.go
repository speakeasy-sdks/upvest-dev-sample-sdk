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

// CreatePortfoliosRebalancingStrategyType - The type of the strategy used in the request.
// * DRIFT - Trigger by drift percentage
// * SCHEDULED - Trigger by scheduled date
type CreatePortfoliosRebalancingStrategyType string

const (
	CreatePortfoliosRebalancingStrategyTypeDrift     CreatePortfoliosRebalancingStrategyType = "DRIFT"
	CreatePortfoliosRebalancingStrategyTypeScheduled CreatePortfoliosRebalancingStrategyType = "SCHEDULED"
)

func (e CreatePortfoliosRebalancingStrategyType) ToPointer() *CreatePortfoliosRebalancingStrategyType {
	return &e
}

func (e *CreatePortfoliosRebalancingStrategyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DRIFT":
		fallthrough
	case "SCHEDULED":
		*e = CreatePortfoliosRebalancingStrategyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreatePortfoliosRebalancingStrategyType: %v", v)
	}
}

type Conditions struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	// The type of the ID used in the request.
	// * ISIN - International Securities Identification Number
	// * UPVEST - UPVEST's unique instrument identifier
	Name string `json:"name"`
	// The type of the strategy used in the request.
	// * DRIFT - Trigger by drift percentage
	// * SCHEDULED - Trigger by scheduled date
	Type CreatePortfoliosRebalancingStrategyType `json:"type"`
}

func (c Conditions) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *Conditions) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Conditions) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *Conditions) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Conditions) GetType() CreatePortfoliosRebalancingStrategyType {
	if o == nil {
		return CreatePortfoliosRebalancingStrategyType("")
	}
	return o.Type
}

type CreatePortfoliosRebalancingStrategyPortfoliosRebalancingStrategyRequest struct {
	// List of conditions
	Conditions []Conditions `json:"conditions"`
	// Strategy name
	Name string `json:"name"`
}

func (o *CreatePortfoliosRebalancingStrategyPortfoliosRebalancingStrategyRequest) GetConditions() []Conditions {
	if o == nil {
		return []Conditions{}
	}
	return o.Conditions
}

func (o *CreatePortfoliosRebalancingStrategyPortfoliosRebalancingStrategyRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type CreatePortfoliosRebalancingStrategyRequest struct {
	RequestBody *CreatePortfoliosRebalancingStrategyPortfoliosRebalancingStrategyRequest `request:"mediaType=application/json"`
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

func (c CreatePortfoliosRebalancingStrategyRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreatePortfoliosRebalancingStrategyRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreatePortfoliosRebalancingStrategyRequest) GetRequestBody() *CreatePortfoliosRebalancingStrategyPortfoliosRebalancingStrategyRequest {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *CreatePortfoliosRebalancingStrategyRequest) GetIdempotencyKey() string {
	if o == nil {
		return ""
	}
	return o.IdempotencyKey
}

func (o *CreatePortfoliosRebalancingStrategyRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *CreatePortfoliosRebalancingStrategyRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *CreatePortfoliosRebalancingStrategyRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *CreatePortfoliosRebalancingStrategyRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// CreatePortfoliosRebalancingStrategyPortfoliosType - The type of the strategy used in the request.
// * DRIFT - Trigger by drift percentage
// * SCHEDULED - Trigger by scheduled date
type CreatePortfoliosRebalancingStrategyPortfoliosType string

const (
	CreatePortfoliosRebalancingStrategyPortfoliosTypeDrift     CreatePortfoliosRebalancingStrategyPortfoliosType = "DRIFT"
	CreatePortfoliosRebalancingStrategyPortfoliosTypeScheduled CreatePortfoliosRebalancingStrategyPortfoliosType = "SCHEDULED"
)

func (e CreatePortfoliosRebalancingStrategyPortfoliosType) ToPointer() *CreatePortfoliosRebalancingStrategyPortfoliosType {
	return &e
}

func (e *CreatePortfoliosRebalancingStrategyPortfoliosType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DRIFT":
		fallthrough
	case "SCHEDULED":
		*e = CreatePortfoliosRebalancingStrategyPortfoliosType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreatePortfoliosRebalancingStrategyPortfoliosType: %v", v)
	}
}

type CreatePortfoliosRebalancingStrategyConditions struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	// The type of the ID used in the request.
	// * ISIN - International Securities Identification Number
	// * UPVEST - UPVEST's unique instrument identifier
	Name string `json:"name"`
	// The type of the strategy used in the request.
	// * DRIFT - Trigger by drift percentage
	// * SCHEDULED - Trigger by scheduled date
	Type CreatePortfoliosRebalancingStrategyPortfoliosType `json:"type"`
}

func (c CreatePortfoliosRebalancingStrategyConditions) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreatePortfoliosRebalancingStrategyConditions) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreatePortfoliosRebalancingStrategyConditions) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *CreatePortfoliosRebalancingStrategyConditions) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreatePortfoliosRebalancingStrategyConditions) GetType() CreatePortfoliosRebalancingStrategyPortfoliosType {
	if o == nil {
		return CreatePortfoliosRebalancingStrategyPortfoliosType("")
	}
	return o.Type
}

// CreatePortfoliosRebalancingStrategyPortfoliosRebalancingStrategy - Portfolios rebalancing strategy
type CreatePortfoliosRebalancingStrategyPortfoliosRebalancingStrategy struct {
	// List of conditions
	Conditions []CreatePortfoliosRebalancingStrategyConditions `json:"conditions"`
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	ID        string    `json:"id"`
	// Strategy name
	Name string `json:"name"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time `json:"updated_at"`
}

func (c CreatePortfoliosRebalancingStrategyPortfoliosRebalancingStrategy) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreatePortfoliosRebalancingStrategyPortfoliosRebalancingStrategy) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreatePortfoliosRebalancingStrategyPortfoliosRebalancingStrategy) GetConditions() []CreatePortfoliosRebalancingStrategyConditions {
	if o == nil {
		return []CreatePortfoliosRebalancingStrategyConditions{}
	}
	return o.Conditions
}

func (o *CreatePortfoliosRebalancingStrategyPortfoliosRebalancingStrategy) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *CreatePortfoliosRebalancingStrategyPortfoliosRebalancingStrategy) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreatePortfoliosRebalancingStrategyPortfoliosRebalancingStrategy) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreatePortfoliosRebalancingStrategyPortfoliosRebalancingStrategy) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

type CreatePortfoliosRebalancingStrategyResponse struct {
	// Portfolios rebalancing strategy
	TwoHundredApplicationJSONPortfoliosRebalancingStrategy *CreatePortfoliosRebalancingStrategyPortfoliosRebalancingStrategy
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreatePortfoliosRebalancingStrategyResponse) GetTwoHundredApplicationJSONPortfoliosRebalancingStrategy() *CreatePortfoliosRebalancingStrategyPortfoliosRebalancingStrategy {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONPortfoliosRebalancingStrategy
}

func (o *CreatePortfoliosRebalancingStrategyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreatePortfoliosRebalancingStrategyResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreatePortfoliosRebalancingStrategyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreatePortfoliosRebalancingStrategyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
