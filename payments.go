// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package upvestdevsamplesdk

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/internal/hooks"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/utils"
	"io"
	"net/http"
	"net/url"
)

// Payments - All payments related paths
type Payments struct {
	sdkConfiguration sdkConfiguration
}

func newPayments(sdkConfig sdkConfiguration) *Payments {
	return &Payments{
		sdkConfiguration: sdkConfig,
	}
}

// CancelCashWithdrawal - Cancel withdrawal by ID
// Cancels a withdrawal specified by its ID. It is only possible to cancel a withdrawal if it has the status `NEW`.
func (s *Payments) CancelCashWithdrawal(ctx context.Context, request operations.CancelCashWithdrawalRequest, security operations.CancelCashWithdrawalSecurity) (*operations.CancelCashWithdrawalResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "cancel_cash_withdrawal",
		SecuritySource: withSecurity(security),
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/payments/withdrawals/{withdrawal_id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/problem+json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	utils.PopulateHeaders(ctx, req, request)

	if err := utils.PopulateSecurity(ctx, req, withSecurity(security)); err != nil {
		return nil, err
	}

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"401", "403", "404", "422", "429", "4XX", "500", "503", "504", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.CancelCashWithdrawalResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 204:
		res.Headers = httpRes.Header

	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 429:
		fallthrough
	case httpRes.StatusCode == 500:
		fallthrough
	case httpRes.StatusCode == 503:
		fallthrough
	case httpRes.StatusCode == 504:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			var out sdkerrors.CancelCashWithdrawalError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			out.RawResponse = httpRes
			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 422:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			var out sdkerrors.CancelCashWithdrawalPaymentsError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			out.RawResponse = httpRes
			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}

// CreateCashWithdrawal - Trigger a withdrawal
// Trigger a withdrawal
func (s *Payments) CreateCashWithdrawal(ctx context.Context, request operations.CreateCashWithdrawalRequest, security operations.CreateCashWithdrawalSecurity, opts ...operations.Option) (*operations.CreateCashWithdrawalResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "create_cash_withdrawal",
		SecuritySource: withSecurity(security),
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionAcceptHeaderOverride,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := url.JoinPath(baseURL, "/payments/withdrawals")
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, true, "RequestBody", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if o.AcceptHeaderOverride != nil {
		req.Header.Set("Accept", string(*o.AcceptHeaderOverride))
	} else {
		req.Header.Set("Accept", "application/json;q=1, application/problem+json;q=0")
	}

	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request)

	if err := utils.PopulateSecurity(ctx, req, withSecurity(security)); err != nil {
		return nil, err
	}

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"400", "401", "403", "404", "406", "429", "4XX", "500", "503", "504", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.CreateCashWithdrawalResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out operations.CreateCashWithdrawalPaymentsWithdrawal
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.PaymentsWithdrawal = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 406:
		fallthrough
	case httpRes.StatusCode == 429:
		fallthrough
	case httpRes.StatusCode == 500:
		fallthrough
	case httpRes.StatusCode == 503:
		fallthrough
	case httpRes.StatusCode == 504:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			var out sdkerrors.CreateCashWithdrawalError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			out.RawResponse = httpRes
			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}

// CreateDirectDebit - Trigger a direct debit
// Trigger a direct debit
func (s *Payments) CreateDirectDebit(ctx context.Context, request operations.CreateDirectDebitRequest, security operations.CreateDirectDebitSecurity, opts ...operations.Option) (*operations.CreateDirectDebitResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "create_direct_debit",
		SecuritySource: withSecurity(security),
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionAcceptHeaderOverride,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := url.JoinPath(baseURL, "/payments/direct_debits")
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, true, "RequestBody", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if o.AcceptHeaderOverride != nil {
		req.Header.Set("Accept", string(*o.AcceptHeaderOverride))
	} else {
		req.Header.Set("Accept", "application/json;q=1, application/problem+json;q=0")
	}

	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request)

	if err := utils.PopulateSecurity(ctx, req, withSecurity(security)); err != nil {
		return nil, err
	}

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"400", "401", "403", "404", "406", "429", "4XX", "500", "503", "504", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.CreateDirectDebitResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out operations.CreateDirectDebitResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Object = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 406:
		fallthrough
	case httpRes.StatusCode == 429:
		fallthrough
	case httpRes.StatusCode == 500:
		fallthrough
	case httpRes.StatusCode == 503:
		fallthrough
	case httpRes.StatusCode == 504:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			var out sdkerrors.CreateDirectDebitError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			out.RawResponse = httpRes
			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}

// ListCashWithdrawals - List withdrawals
// List withdrawals
func (s *Payments) ListCashWithdrawals(ctx context.Context, request operations.ListCashWithdrawalsRequest, security operations.ListCashWithdrawalsSecurity, opts ...operations.Option) (*operations.ListCashWithdrawalsResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "list_cash_withdrawals",
		SecuritySource: withSecurity(security),
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionAcceptHeaderOverride,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/account_groups/{account_group_id}/payments/withdrawals", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if o.AcceptHeaderOverride != nil {
		req.Header.Set("Accept", string(*o.AcceptHeaderOverride))
	} else {
		req.Header.Set("Accept", "application/json;q=1, application/problem+json;q=0")
	}

	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	utils.PopulateHeaders(ctx, req, request)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	if err := utils.PopulateSecurity(ctx, req, withSecurity(security)); err != nil {
		return nil, err
	}

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"400", "401", "403", "404", "406", "429", "4XX", "500", "503", "504", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.ListCashWithdrawalsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out operations.ListCashWithdrawalsWithdrawalsListResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.WithdrawalsListResponse = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 406:
		fallthrough
	case httpRes.StatusCode == 429:
		fallthrough
	case httpRes.StatusCode == 500:
		fallthrough
	case httpRes.StatusCode == 503:
		fallthrough
	case httpRes.StatusCode == 504:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			var out sdkerrors.ListCashWithdrawalsError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			out.RawResponse = httpRes
			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}

// ListDirectDebits - List direct debits
// List direct debits
func (s *Payments) ListDirectDebits(ctx context.Context, request operations.ListDirectDebitsRequest, security operations.ListDirectDebitsSecurity, opts ...operations.Option) (*operations.ListDirectDebitsResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "list_direct_debits",
		SecuritySource: withSecurity(security),
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionAcceptHeaderOverride,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/account_groups/{account_group_id}/payments/direct_debits", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if o.AcceptHeaderOverride != nil {
		req.Header.Set("Accept", string(*o.AcceptHeaderOverride))
	} else {
		req.Header.Set("Accept", "application/json;q=1, application/problem+json;q=0")
	}

	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	utils.PopulateHeaders(ctx, req, request)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	if err := utils.PopulateSecurity(ctx, req, withSecurity(security)); err != nil {
		return nil, err
	}

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"400", "401", "403", "404", "406", "429", "4XX", "500", "503", "504", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.ListDirectDebitsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out operations.ListDirectDebitsPaymentsDirectDebitsListResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.PaymentsDirectDebitsListResponse = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 406:
		fallthrough
	case httpRes.StatusCode == 429:
		fallthrough
	case httpRes.StatusCode == 500:
		fallthrough
	case httpRes.StatusCode == 503:
		fallthrough
	case httpRes.StatusCode == 504:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			var out sdkerrors.ListDirectDebitsError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			out.RawResponse = httpRes
			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}

// RetrieveCashWithdrawal - Retrieve withdrawal
// Retrieve withdrawal
func (s *Payments) RetrieveCashWithdrawal(ctx context.Context, request operations.RetrieveCashWithdrawalRequest, security operations.RetrieveCashWithdrawalSecurity, opts ...operations.Option) (*operations.RetrieveCashWithdrawalResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "retrieve_cash_withdrawal",
		SecuritySource: withSecurity(security),
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionAcceptHeaderOverride,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/payments/withdrawals/{withdrawal_id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if o.AcceptHeaderOverride != nil {
		req.Header.Set("Accept", string(*o.AcceptHeaderOverride))
	} else {
		req.Header.Set("Accept", "application/json;q=1, application/problem+json;q=0")
	}

	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	utils.PopulateHeaders(ctx, req, request)

	if err := utils.PopulateSecurity(ctx, req, withSecurity(security)); err != nil {
		return nil, err
	}

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"401", "403", "404", "406", "429", "4XX", "500", "503", "504", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.RetrieveCashWithdrawalResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out operations.RetrieveCashWithdrawalPaymentsWithdrawal
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.PaymentsWithdrawal = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 406:
		fallthrough
	case httpRes.StatusCode == 429:
		fallthrough
	case httpRes.StatusCode == 500:
		fallthrough
	case httpRes.StatusCode == 503:
		fallthrough
	case httpRes.StatusCode == 504:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			var out sdkerrors.RetrieveCashWithdrawalError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			out.RawResponse = httpRes
			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}

// RetrieveDirectDebit - Retrieve a direct debit
// Retrieve a direct debit
func (s *Payments) RetrieveDirectDebit(ctx context.Context, request operations.RetrieveDirectDebitRequest, security operations.RetrieveDirectDebitSecurity, opts ...operations.Option) (*operations.RetrieveDirectDebitResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "retrieve_direct_debit",
		SecuritySource: withSecurity(security),
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionAcceptHeaderOverride,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/payments/direct_debits/{direct_debit_id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if o.AcceptHeaderOverride != nil {
		req.Header.Set("Accept", string(*o.AcceptHeaderOverride))
	} else {
		req.Header.Set("Accept", "application/json;q=1, application/problem+json;q=0")
	}

	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	utils.PopulateHeaders(ctx, req, request)

	if err := utils.PopulateSecurity(ctx, req, withSecurity(security)); err != nil {
		return nil, err
	}

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"401", "403", "404", "406", "429", "4XX", "500", "503", "504", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.RetrieveDirectDebitResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out operations.RetrieveDirectDebitResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Object = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 406:
		fallthrough
	case httpRes.StatusCode == 429:
		fallthrough
	case httpRes.StatusCode == 500:
		fallthrough
	case httpRes.StatusCode == 503:
		fallthrough
	case httpRes.StatusCode == 504:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			var out sdkerrors.RetrieveDirectDebitError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			out.RawResponse = httpRes
			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}
