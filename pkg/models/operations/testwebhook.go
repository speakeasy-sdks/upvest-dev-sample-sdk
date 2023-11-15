// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/utils"
	"net/http"
)

type TestWebhookRequest struct {
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-http-header
	Signature string `header:"style=simple,explode=false,name=signature"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-input-http-he
	SignatureInput string `header:"style=simple,explode=false,name=signature-input"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
	// Webhook identifier
	WebhookID string `pathParam:"style=simple,explode=false,name=webhook_id"`
}

func (t TestWebhookRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TestWebhookRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TestWebhookRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *TestWebhookRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *TestWebhookRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *TestWebhookRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

func (o *TestWebhookRequest) GetWebhookID() string {
	if o == nil {
		return ""
	}
	return o.WebhookID
}

type TestWebhookResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *TestWebhookResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TestWebhookResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *TestWebhookResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TestWebhookResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
