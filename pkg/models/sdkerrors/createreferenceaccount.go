// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

// CreateReferenceAccountReferenceAccountsResponse504Error - Gateway Timeout. The service gateway has reached its internal timeout.
type CreateReferenceAccountReferenceAccountsResponse504Error struct {
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

var _ error = &CreateReferenceAccountReferenceAccountsResponse504Error{}

func (e *CreateReferenceAccountReferenceAccountsResponse504Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// CreateReferenceAccountReferenceAccountsResponse503Error - Service Unavailable. The service handling for this request cannot be reached at this time.
type CreateReferenceAccountReferenceAccountsResponse503Error struct {
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

var _ error = &CreateReferenceAccountReferenceAccountsResponse503Error{}

func (e *CreateReferenceAccountReferenceAccountsResponse503Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// CreateReferenceAccountReferenceAccountsResponse500Error - Internal Server Error. The service encountered an unexpected error.
type CreateReferenceAccountReferenceAccountsResponse500Error struct {
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

var _ error = &CreateReferenceAccountReferenceAccountsResponse500Error{}

func (e *CreateReferenceAccountReferenceAccountsResponse500Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// CreateReferenceAccountReferenceAccountsResponse429Error - Too Many Requests. The caller has exceeded their quota for the time period and has been throttled.
type CreateReferenceAccountReferenceAccountsResponse429Error struct {
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

var _ error = &CreateReferenceAccountReferenceAccountsResponse429Error{}

func (e *CreateReferenceAccountReferenceAccountsResponse429Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// CreateReferenceAccountReferenceAccountsResponse406Error - Not Acceptable. The resource does not have a current representation that would be acceptable to the user agent. "Accept" header defined unsupported value.
type CreateReferenceAccountReferenceAccountsResponse406Error struct {
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

var _ error = &CreateReferenceAccountReferenceAccountsResponse406Error{}

func (e *CreateReferenceAccountReferenceAccountsResponse406Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// CreateReferenceAccountReferenceAccountsResponse404Error - Not Found. The requested resource could not be found.
type CreateReferenceAccountReferenceAccountsResponse404Error struct {
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

var _ error = &CreateReferenceAccountReferenceAccountsResponse404Error{}

func (e *CreateReferenceAccountReferenceAccountsResponse404Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// CreateReferenceAccountReferenceAccountsResponseError - Forbidden. The caller has been authenticated but is not allowed to take the requested action.
type CreateReferenceAccountReferenceAccountsResponseError struct {
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

var _ error = &CreateReferenceAccountReferenceAccountsResponseError{}

func (e *CreateReferenceAccountReferenceAccountsResponseError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// CreateReferenceAccountReferenceAccountsError - Unauthorized. The caller has not been authenticated.
type CreateReferenceAccountReferenceAccountsError struct {
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

var _ error = &CreateReferenceAccountReferenceAccountsError{}

func (e *CreateReferenceAccountReferenceAccountsError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// CreateReferenceAccountError - Bad Request. The incoming request had a malformed parameter/object.
type CreateReferenceAccountError struct {
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

var _ error = &CreateReferenceAccountError{}

func (e *CreateReferenceAccountError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
