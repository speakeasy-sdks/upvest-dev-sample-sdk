// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

// DeleteReferenceAccountReferenceAccountsResponse504Error - Gateway Timeout. The service gateway has reached its internal timeout.
type DeleteReferenceAccountReferenceAccountsResponse504Error struct {
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

var _ error = &DeleteReferenceAccountReferenceAccountsResponse504Error{}

func (e *DeleteReferenceAccountReferenceAccountsResponse504Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// DeleteReferenceAccountReferenceAccountsResponse503Error - Service Unavailable. The service handling for this request cannot be reached at this time.
type DeleteReferenceAccountReferenceAccountsResponse503Error struct {
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

var _ error = &DeleteReferenceAccountReferenceAccountsResponse503Error{}

func (e *DeleteReferenceAccountReferenceAccountsResponse503Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// DeleteReferenceAccountReferenceAccountsResponse500Error - Internal Server Error. The service encountered an unexpected error.
type DeleteReferenceAccountReferenceAccountsResponse500Error struct {
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

var _ error = &DeleteReferenceAccountReferenceAccountsResponse500Error{}

func (e *DeleteReferenceAccountReferenceAccountsResponse500Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// DeleteReferenceAccountReferenceAccountsResponse429Error - Too Many Requests. The caller has exceeded their quota for the time period and has been throttled.
type DeleteReferenceAccountReferenceAccountsResponse429Error struct {
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

var _ error = &DeleteReferenceAccountReferenceAccountsResponse429Error{}

func (e *DeleteReferenceAccountReferenceAccountsResponse429Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// DeleteReferenceAccountReferenceAccountsResponseError - Not Found. The requested resource could not be found.
type DeleteReferenceAccountReferenceAccountsResponseError struct {
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

var _ error = &DeleteReferenceAccountReferenceAccountsResponseError{}

func (e *DeleteReferenceAccountReferenceAccountsResponseError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// DeleteReferenceAccountReferenceAccountsError - Forbidden. The caller has been authenticated but is not allowed to take the requested action.
type DeleteReferenceAccountReferenceAccountsError struct {
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

var _ error = &DeleteReferenceAccountReferenceAccountsError{}

func (e *DeleteReferenceAccountReferenceAccountsError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// DeleteReferenceAccountError - Unauthorized. The caller has not been authenticated.
type DeleteReferenceAccountError struct {
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

var _ error = &DeleteReferenceAccountError{}

func (e *DeleteReferenceAccountError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
