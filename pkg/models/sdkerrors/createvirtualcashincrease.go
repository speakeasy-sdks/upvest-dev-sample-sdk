// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

// CreateVirtualCashIncreaseVirtualCashBalancesResponse504Error - Gateway Timeout. The service gateway has reached its internal timeout.
type CreateVirtualCashIncreaseVirtualCashBalancesResponse504Error struct {
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
	// A human-readable description of the specific error.
	Detail *string `json:"detail,omitempty"`
	// This optional key may be present, with a unique URI for the specific error; this will often point to an error log for that specific response.
	Instance *string `json:"instance,omitempty"`
	// Correlation ID for the original request.
	RequestID *string `json:"request_id,omitempty"`
	// Transmission of the HTTP status code so that all information can be found in one place, but also to correct changes in the status code due to the use of proxy servers.
	Status int64 `json:"status"`
	// A short, human-readable title for the general error type; the title should not change for given types.
	Title *string `json:"title,omitempty"`
	// URL to a document describing the error condition.
	Type string `json:"type"`
}

var _ error = &CreateVirtualCashIncreaseVirtualCashBalancesResponse504Error{}

func (e *CreateVirtualCashIncreaseVirtualCashBalancesResponse504Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// CreateVirtualCashIncreaseVirtualCashBalancesResponse503Error - Service Unavailable. The service handling for this request cannot be reached at this time.
type CreateVirtualCashIncreaseVirtualCashBalancesResponse503Error struct {
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
	// A human-readable description of the specific error.
	Detail *string `json:"detail,omitempty"`
	// This optional key may be present, with a unique URI for the specific error; this will often point to an error log for that specific response.
	Instance *string `json:"instance,omitempty"`
	// Correlation ID for the original request.
	RequestID *string `json:"request_id,omitempty"`
	// Transmission of the HTTP status code so that all information can be found in one place, but also to correct changes in the status code due to the use of proxy servers.
	Status int64 `json:"status"`
	// A short, human-readable title for the general error type; the title should not change for given types.
	Title *string `json:"title,omitempty"`
	// URL to a document describing the error condition.
	Type string `json:"type"`
}

var _ error = &CreateVirtualCashIncreaseVirtualCashBalancesResponse503Error{}

func (e *CreateVirtualCashIncreaseVirtualCashBalancesResponse503Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// CreateVirtualCashIncreaseVirtualCashBalancesResponse500Error - Internal Server Error. The service encountered an unexpected error.
type CreateVirtualCashIncreaseVirtualCashBalancesResponse500Error struct {
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
	// A human-readable description of the specific error.
	Detail *string `json:"detail,omitempty"`
	// This optional key may be present, with a unique URI for the specific error; this will often point to an error log for that specific response.
	Instance *string `json:"instance,omitempty"`
	// Correlation ID for the original request.
	RequestID *string `json:"request_id,omitempty"`
	// Transmission of the HTTP status code so that all information can be found in one place, but also to correct changes in the status code due to the use of proxy servers.
	Status int64 `json:"status"`
	// A short, human-readable title for the general error type; the title should not change for given types.
	Title *string `json:"title,omitempty"`
	// URL to a document describing the error condition.
	Type string `json:"type"`
}

var _ error = &CreateVirtualCashIncreaseVirtualCashBalancesResponse500Error{}

func (e *CreateVirtualCashIncreaseVirtualCashBalancesResponse500Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// CreateVirtualCashIncreaseVirtualCashBalancesResponse429Error - Too Many Requests. The caller has exceeded their quota for the time period and has been throttled.
type CreateVirtualCashIncreaseVirtualCashBalancesResponse429Error struct {
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
	// A human-readable description of the specific error.
	Detail *string `json:"detail,omitempty"`
	// This optional key may be present, with a unique URI for the specific error; this will often point to an error log for that specific response.
	Instance *string `json:"instance,omitempty"`
	// Correlation ID for the original request.
	RequestID *string `json:"request_id,omitempty"`
	// Transmission of the HTTP status code so that all information can be found in one place, but also to correct changes in the status code due to the use of proxy servers.
	Status int64 `json:"status"`
	// A short, human-readable title for the general error type; the title should not change for given types.
	Title *string `json:"title,omitempty"`
	// URL to a document describing the error condition.
	Type string `json:"type"`
}

var _ error = &CreateVirtualCashIncreaseVirtualCashBalancesResponse429Error{}

func (e *CreateVirtualCashIncreaseVirtualCashBalancesResponse429Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// CreateVirtualCashIncreaseVirtualCashBalancesResponse406Error - Not Acceptable. The resource does not have a current representation that would be acceptable to the user agent. "Accept" header defined unsupported value.
type CreateVirtualCashIncreaseVirtualCashBalancesResponse406Error struct {
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
	// A human-readable description of the specific error.
	Detail *string `json:"detail,omitempty"`
	// This optional key may be present, with a unique URI for the specific error; this will often point to an error log for that specific response.
	Instance *string `json:"instance,omitempty"`
	// Correlation ID for the original request.
	RequestID *string `json:"request_id,omitempty"`
	// Transmission of the HTTP status code so that all information can be found in one place, but also to correct changes in the status code due to the use of proxy servers.
	Status int64 `json:"status"`
	// A short, human-readable title for the general error type; the title should not change for given types.
	Title *string `json:"title,omitempty"`
	// URL to a document describing the error condition.
	Type string `json:"type"`
}

var _ error = &CreateVirtualCashIncreaseVirtualCashBalancesResponse406Error{}

func (e *CreateVirtualCashIncreaseVirtualCashBalancesResponse406Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// CreateVirtualCashIncreaseVirtualCashBalancesResponse404Error - Not Found. The requested resource could not be found.
type CreateVirtualCashIncreaseVirtualCashBalancesResponse404Error struct {
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
	// A human-readable description of the specific error.
	Detail *string `json:"detail,omitempty"`
	// This optional key may be present, with a unique URI for the specific error; this will often point to an error log for that specific response.
	Instance *string `json:"instance,omitempty"`
	// Correlation ID for the original request.
	RequestID *string `json:"request_id,omitempty"`
	// Transmission of the HTTP status code so that all information can be found in one place, but also to correct changes in the status code due to the use of proxy servers.
	Status int64 `json:"status"`
	// A short, human-readable title for the general error type; the title should not change for given types.
	Title *string `json:"title,omitempty"`
	// URL to a document describing the error condition.
	Type string `json:"type"`
}

var _ error = &CreateVirtualCashIncreaseVirtualCashBalancesResponse404Error{}

func (e *CreateVirtualCashIncreaseVirtualCashBalancesResponse404Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// CreateVirtualCashIncreaseVirtualCashBalancesResponseError - Forbidden. The caller has been authenticated but is not allowed to take the requested action.
type CreateVirtualCashIncreaseVirtualCashBalancesResponseError struct {
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
	// A human-readable description of the specific error.
	Detail *string `json:"detail,omitempty"`
	// This optional key may be present, with a unique URI for the specific error; this will often point to an error log for that specific response.
	Instance *string `json:"instance,omitempty"`
	// Correlation ID for the original request.
	RequestID *string `json:"request_id,omitempty"`
	// Transmission of the HTTP status code so that all information can be found in one place, but also to correct changes in the status code due to the use of proxy servers.
	Status int64 `json:"status"`
	// A short, human-readable title for the general error type; the title should not change for given types.
	Title *string `json:"title,omitempty"`
	// URL to a document describing the error condition.
	Type string `json:"type"`
}

var _ error = &CreateVirtualCashIncreaseVirtualCashBalancesResponseError{}

func (e *CreateVirtualCashIncreaseVirtualCashBalancesResponseError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// CreateVirtualCashIncreaseVirtualCashBalancesError - Unauthorized. The caller has not been authenticated.
type CreateVirtualCashIncreaseVirtualCashBalancesError struct {
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
	// A human-readable description of the specific error.
	Detail *string `json:"detail,omitempty"`
	// This optional key may be present, with a unique URI for the specific error; this will often point to an error log for that specific response.
	Instance *string `json:"instance,omitempty"`
	// Correlation ID for the original request.
	RequestID *string `json:"request_id,omitempty"`
	// Transmission of the HTTP status code so that all information can be found in one place, but also to correct changes in the status code due to the use of proxy servers.
	Status int64 `json:"status"`
	// A short, human-readable title for the general error type; the title should not change for given types.
	Title *string `json:"title,omitempty"`
	// URL to a document describing the error condition.
	Type string `json:"type"`
}

var _ error = &CreateVirtualCashIncreaseVirtualCashBalancesError{}

func (e *CreateVirtualCashIncreaseVirtualCashBalancesError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// CreateVirtualCashIncreaseError - Bad Request. The incoming request had a malformed parameter/object.
type CreateVirtualCashIncreaseError struct {
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
	// A human-readable description of the specific error.
	Detail *string `json:"detail,omitempty"`
	// This optional key may be present, with a unique URI for the specific error; this will often point to an error log for that specific response.
	Instance *string `json:"instance,omitempty"`
	// Correlation ID for the original request.
	RequestID *string `json:"request_id,omitempty"`
	// Transmission of the HTTP status code so that all information can be found in one place, but also to correct changes in the status code due to the use of proxy servers.
	Status int64 `json:"status"`
	// A short, human-readable title for the general error type; the title should not change for given types.
	Title *string `json:"title,omitempty"`
	// URL to a document describing the error condition.
	Type string `json:"type"`
}

var _ error = &CreateVirtualCashIncreaseError{}

func (e *CreateVirtualCashIncreaseError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
