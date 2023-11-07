// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/utils"
	"net/http"
	"time"
)

type CreatePortfoliosConfigurationPortfoliosConfigurationCreateRequest struct {
	// Account unique identifier.
	AccountID    string `json:"account_id"`
	AllocationID string `json:"allocation_id"`
	// List of rebalancing strategy ids
	RebalancingStrategyIds []string `json:"rebalancing_strategy_ids,omitempty"`
}

func (o *CreatePortfoliosConfigurationPortfoliosConfigurationCreateRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *CreatePortfoliosConfigurationPortfoliosConfigurationCreateRequest) GetAllocationID() string {
	if o == nil {
		return ""
	}
	return o.AllocationID
}

func (o *CreatePortfoliosConfigurationPortfoliosConfigurationCreateRequest) GetRebalancingStrategyIds() []string {
	if o == nil {
		return nil
	}
	return o.RebalancingStrategyIds
}

type CreatePortfoliosConfigurationRequest struct {
	RequestBody *CreatePortfoliosConfigurationPortfoliosConfigurationCreateRequest `request:"mediaType=application/json"`
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

func (c CreatePortfoliosConfigurationRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreatePortfoliosConfigurationRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreatePortfoliosConfigurationRequest) GetRequestBody() *CreatePortfoliosConfigurationPortfoliosConfigurationCreateRequest {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *CreatePortfoliosConfigurationRequest) GetIdempotencyKey() string {
	if o == nil {
		return ""
	}
	return o.IdempotencyKey
}

func (o *CreatePortfoliosConfigurationRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *CreatePortfoliosConfigurationRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *CreatePortfoliosConfigurationRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *CreatePortfoliosConfigurationRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// CreatePortfoliosConfigurationPortfoliosConfiguration - Portfolios configuration
type CreatePortfoliosConfigurationPortfoliosConfiguration struct {
	// Account unique identifier.
	AccountID    string `json:"account_id"`
	AllocationID string `json:"allocation_id"`
	// Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	CreatedAt time.Time `json:"created_at"`
	// List of rebalancing strategy ids
	RebalancingStrategyIds []string `json:"rebalancing_strategy_ids,omitempty"`
	// Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)
	UpdatedAt time.Time `json:"updated_at"`
}

func (c CreatePortfoliosConfigurationPortfoliosConfiguration) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreatePortfoliosConfigurationPortfoliosConfiguration) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreatePortfoliosConfigurationPortfoliosConfiguration) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *CreatePortfoliosConfigurationPortfoliosConfiguration) GetAllocationID() string {
	if o == nil {
		return ""
	}
	return o.AllocationID
}

func (o *CreatePortfoliosConfigurationPortfoliosConfiguration) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *CreatePortfoliosConfigurationPortfoliosConfiguration) GetRebalancingStrategyIds() []string {
	if o == nil {
		return nil
	}
	return o.RebalancingStrategyIds
}

func (o *CreatePortfoliosConfigurationPortfoliosConfiguration) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

type CreatePortfoliosConfigurationResponse struct {
	// Portfolios configuration
	TwoHundredApplicationJSONPortfoliosConfiguration *CreatePortfoliosConfigurationPortfoliosConfiguration
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreatePortfoliosConfigurationResponse) GetTwoHundredApplicationJSONPortfoliosConfiguration() *CreatePortfoliosConfigurationPortfoliosConfiguration {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONPortfoliosConfiguration
}

func (o *CreatePortfoliosConfigurationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreatePortfoliosConfigurationResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreatePortfoliosConfigurationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreatePortfoliosConfigurationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
