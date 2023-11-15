// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/utils"
	"net/http"
)

type RetrieveCashBalanceRequest struct {
	AccountGroupID string `pathParam:"style=simple,explode=false,name=account_group_id"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-http-header
	Signature string `header:"style=simple,explode=false,name=signature"`
	// https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-input-http-he
	SignatureInput string `header:"style=simple,explode=false,name=signature-input"`
	// Upvest API version (Note: Do not include quotation marks)
	UpvestAPIVersion *shared.APIVersion `default:"1" header:"style=simple,explode=false,name=upvest-api-version"`
	// Tenant Client ID
	UpvestClientID string `header:"style=simple,explode=false,name=upvest-client-id"`
}

func (r RetrieveCashBalanceRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrieveCashBalanceRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrieveCashBalanceRequest) GetAccountGroupID() string {
	if o == nil {
		return ""
	}
	return o.AccountGroupID
}

func (o *RetrieveCashBalanceRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *RetrieveCashBalanceRequest) GetSignatureInput() string {
	if o == nil {
		return ""
	}
	return o.SignatureInput
}

func (o *RetrieveCashBalanceRequest) GetUpvestAPIVersion() *shared.APIVersion {
	if o == nil {
		return nil
	}
	return o.UpvestAPIVersion
}

func (o *RetrieveCashBalanceRequest) GetUpvestClientID() string {
	if o == nil {
		return ""
	}
	return o.UpvestClientID
}

// RetrieveCashBalanceCurrency - Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
// * EUR - Euro
type RetrieveCashBalanceCurrency string

const (
	RetrieveCashBalanceCurrencyEur RetrieveCashBalanceCurrency = "EUR"
)

func (e RetrieveCashBalanceCurrency) ToPointer() *RetrieveCashBalanceCurrency {
	return &e
}

func (e *RetrieveCashBalanceCurrency) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EUR":
		*e = RetrieveCashBalanceCurrency(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveCashBalanceCurrency: %v", v)
	}
}

// RetrieveCashBalanceResponseBody - Response
type RetrieveCashBalanceResponseBody struct {
	// Account group unique identifier.
	AccountGroupID         string `json:"account_group_id"`
	AvailableForTrading    string `json:"available_for_trading"`
	AvailableForWithdrawal string `json:"available_for_withdrawal"`
	Balance                string `json:"balance"`
	// Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
	// * EUR - Euro
	Currency          *RetrieveCashBalanceCurrency `default:"EUR" json:"currency"`
	LockedForTrading  string                       `json:"locked_for_trading"`
	PendingSettlement string                       `json:"pending_settlement"`
}

func (r RetrieveCashBalanceResponseBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrieveCashBalanceResponseBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrieveCashBalanceResponseBody) GetAccountGroupID() string {
	if o == nil {
		return ""
	}
	return o.AccountGroupID
}

func (o *RetrieveCashBalanceResponseBody) GetAvailableForTrading() string {
	if o == nil {
		return ""
	}
	return o.AvailableForTrading
}

func (o *RetrieveCashBalanceResponseBody) GetAvailableForWithdrawal() string {
	if o == nil {
		return ""
	}
	return o.AvailableForWithdrawal
}

func (o *RetrieveCashBalanceResponseBody) GetBalance() string {
	if o == nil {
		return ""
	}
	return o.Balance
}

func (o *RetrieveCashBalanceResponseBody) GetCurrency() *RetrieveCashBalanceCurrency {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *RetrieveCashBalanceResponseBody) GetLockedForTrading() string {
	if o == nil {
		return ""
	}
	return o.LockedForTrading
}

func (o *RetrieveCashBalanceResponseBody) GetPendingSettlement() string {
	if o == nil {
		return ""
	}
	return o.PendingSettlement
}

type RetrieveCashBalanceResponse struct {
	// Response
	TwoHundredApplicationJSONObject *RetrieveCashBalanceResponseBody
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RetrieveCashBalanceResponse) GetTwoHundredApplicationJSONObject() *RetrieveCashBalanceResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *RetrieveCashBalanceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RetrieveCashBalanceResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *RetrieveCashBalanceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RetrieveCashBalanceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
