// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

// ListDirectDebitsPaymentsResponse504Error - Gateway Timeout. The service gateway has reached its internal timeout.
type ListDirectDebitsPaymentsResponse504Error struct {
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

var _ error = &ListDirectDebitsPaymentsResponse504Error{}

func (e *ListDirectDebitsPaymentsResponse504Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ListDirectDebitsPaymentsResponse503Error - Service Unavailable. The service handling for this request cannot be reached at this time.
type ListDirectDebitsPaymentsResponse503Error struct {
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

var _ error = &ListDirectDebitsPaymentsResponse503Error{}

func (e *ListDirectDebitsPaymentsResponse503Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ListDirectDebitsPaymentsResponse500Error - Internal Server Error. The service encountered an unexpected error.
type ListDirectDebitsPaymentsResponse500Error struct {
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

var _ error = &ListDirectDebitsPaymentsResponse500Error{}

func (e *ListDirectDebitsPaymentsResponse500Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ListDirectDebitsPaymentsResponse429Error - Too Many Requests. The caller has exceeded their quota for the time period and has been throttled.
type ListDirectDebitsPaymentsResponse429Error struct {
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

var _ error = &ListDirectDebitsPaymentsResponse429Error{}

func (e *ListDirectDebitsPaymentsResponse429Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ListDirectDebitsPaymentsResponse406Error - Not Acceptable. The resource does not have a current representation that would be acceptable to the user agent. "Accept" header defined unsupported value.
type ListDirectDebitsPaymentsResponse406Error struct {
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

var _ error = &ListDirectDebitsPaymentsResponse406Error{}

func (e *ListDirectDebitsPaymentsResponse406Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ListDirectDebitsPaymentsResponse404Error - Not Found. The requested resource could not be found.
type ListDirectDebitsPaymentsResponse404Error struct {
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

var _ error = &ListDirectDebitsPaymentsResponse404Error{}

func (e *ListDirectDebitsPaymentsResponse404Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ListDirectDebitsPaymentsResponseError - Forbidden. The caller has been authenticated but is not allowed to take the requested action.
type ListDirectDebitsPaymentsResponseError struct {
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

var _ error = &ListDirectDebitsPaymentsResponseError{}

func (e *ListDirectDebitsPaymentsResponseError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ListDirectDebitsPaymentsError - Unauthorized. The caller has not been authenticated.
type ListDirectDebitsPaymentsError struct {
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

var _ error = &ListDirectDebitsPaymentsError{}

func (e *ListDirectDebitsPaymentsError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ListDirectDebitsError - Bad Request. The incoming request had a malformed parameter/object.
type ListDirectDebitsError struct {
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

var _ error = &ListDirectDebitsError{}

func (e *ListDirectDebitsError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
