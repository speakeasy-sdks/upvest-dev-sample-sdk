// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

// ListUserAccountGroupsUsersResponse504Error - Gateway Timeout. The service gateway has reached its internal timeout.
type ListUserAccountGroupsUsersResponse504Error struct {
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

var _ error = &ListUserAccountGroupsUsersResponse504Error{}

func (e *ListUserAccountGroupsUsersResponse504Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ListUserAccountGroupsUsersResponse503Error - Service Unavailable. The service handling for this request cannot be reached at this time.
type ListUserAccountGroupsUsersResponse503Error struct {
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

var _ error = &ListUserAccountGroupsUsersResponse503Error{}

func (e *ListUserAccountGroupsUsersResponse503Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ListUserAccountGroupsUsersResponse500Error - Internal Server Error. The service encountered an unexpected error.
type ListUserAccountGroupsUsersResponse500Error struct {
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

var _ error = &ListUserAccountGroupsUsersResponse500Error{}

func (e *ListUserAccountGroupsUsersResponse500Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ListUserAccountGroupsUsersResponse429Error - Too Many Requests. The caller has exceeded their quota for the time period and has been throttled.
type ListUserAccountGroupsUsersResponse429Error struct {
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

var _ error = &ListUserAccountGroupsUsersResponse429Error{}

func (e *ListUserAccountGroupsUsersResponse429Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ListUserAccountGroupsUsersResponse406Error - Not Acceptable. The resource does not have a current representation that would be acceptable to the user agent. "Accept" header defined unsupported value.
type ListUserAccountGroupsUsersResponse406Error struct {
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

var _ error = &ListUserAccountGroupsUsersResponse406Error{}

func (e *ListUserAccountGroupsUsersResponse406Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ListUserAccountGroupsUsersResponse404Error - Not Found. The requested resource could not be found.
type ListUserAccountGroupsUsersResponse404Error struct {
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

var _ error = &ListUserAccountGroupsUsersResponse404Error{}

func (e *ListUserAccountGroupsUsersResponse404Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ListUserAccountGroupsUsersResponseError - Forbidden. The caller has been authenticated but is not allowed to take the requested action.
type ListUserAccountGroupsUsersResponseError struct {
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

var _ error = &ListUserAccountGroupsUsersResponseError{}

func (e *ListUserAccountGroupsUsersResponseError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ListUserAccountGroupsUsersError - Unauthorized. The caller has not been authenticated.
type ListUserAccountGroupsUsersError struct {
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

var _ error = &ListUserAccountGroupsUsersError{}

func (e *ListUserAccountGroupsUsersError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ListUserAccountGroupsError - Bad Request. The incoming request had a malformed parameter/object.
type ListUserAccountGroupsError struct {
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

var _ error = &ListUserAccountGroupsError{}

func (e *ListUserAccountGroupsError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
