// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

// ListPortfoliosConfigurationsPortfoliosResponse504Error - Gateway Timeout. The service gateway has reached its internal timeout.
type ListPortfoliosConfigurationsPortfoliosResponse504Error struct {
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

var _ error = &ListPortfoliosConfigurationsPortfoliosResponse504Error{}

func (e *ListPortfoliosConfigurationsPortfoliosResponse504Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ListPortfoliosConfigurationsPortfoliosResponse503Error - Service Unavailable. The service handling for this request cannot be reached at this time.
type ListPortfoliosConfigurationsPortfoliosResponse503Error struct {
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

var _ error = &ListPortfoliosConfigurationsPortfoliosResponse503Error{}

func (e *ListPortfoliosConfigurationsPortfoliosResponse503Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ListPortfoliosConfigurationsPortfoliosResponse500Error - Internal Server Error. The service encountered an unexpected error.
type ListPortfoliosConfigurationsPortfoliosResponse500Error struct {
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

var _ error = &ListPortfoliosConfigurationsPortfoliosResponse500Error{}

func (e *ListPortfoliosConfigurationsPortfoliosResponse500Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ListPortfoliosConfigurationsPortfoliosResponse429Error - Too Many Requests. The caller has exceeded their quota for the time period and has been throttled.
type ListPortfoliosConfigurationsPortfoliosResponse429Error struct {
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

var _ error = &ListPortfoliosConfigurationsPortfoliosResponse429Error{}

func (e *ListPortfoliosConfigurationsPortfoliosResponse429Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ListPortfoliosConfigurationsPortfoliosResponse406Error - Not Acceptable. The resource does not have a current representation that would be acceptable to the user agent. "Accept" header defined unsupported value.
type ListPortfoliosConfigurationsPortfoliosResponse406Error struct {
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

var _ error = &ListPortfoliosConfigurationsPortfoliosResponse406Error{}

func (e *ListPortfoliosConfigurationsPortfoliosResponse406Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ListPortfoliosConfigurationsPortfoliosResponse405Error - Method Not Allowed. The requested method is not allowed on the requested resource.
type ListPortfoliosConfigurationsPortfoliosResponse405Error struct {
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

var _ error = &ListPortfoliosConfigurationsPortfoliosResponse405Error{}

func (e *ListPortfoliosConfigurationsPortfoliosResponse405Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ListPortfoliosConfigurationsPortfoliosResponse404Error - Not Found. The requested resource could not be found.
type ListPortfoliosConfigurationsPortfoliosResponse404Error struct {
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

var _ error = &ListPortfoliosConfigurationsPortfoliosResponse404Error{}

func (e *ListPortfoliosConfigurationsPortfoliosResponse404Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ListPortfoliosConfigurationsPortfoliosResponseError - Forbidden. The caller has been authenticated but is not allowed to take the requested action.
type ListPortfoliosConfigurationsPortfoliosResponseError struct {
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

var _ error = &ListPortfoliosConfigurationsPortfoliosResponseError{}

func (e *ListPortfoliosConfigurationsPortfoliosResponseError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ListPortfoliosConfigurationsPortfoliosError - Unauthorized. The caller has not been authenticated.
type ListPortfoliosConfigurationsPortfoliosError struct {
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

var _ error = &ListPortfoliosConfigurationsPortfoliosError{}

func (e *ListPortfoliosConfigurationsPortfoliosError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ListPortfoliosConfigurationsError - Bad Request. The incoming request had a malformed parameter/object.
type ListPortfoliosConfigurationsError struct {
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

var _ error = &ListPortfoliosConfigurationsError{}

func (e *ListPortfoliosConfigurationsError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
