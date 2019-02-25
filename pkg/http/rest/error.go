package rest

import (
	"net/http"

	"github.com/go-chi/render"
)

// HTTPError is http error accessor
type HTTPError interface {
	HTTPError() *ErrResponse
}

// ErrResponse is an error response renderer
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

// Render pushes error response to a http.ResponseWriter
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode())

	return nil
}

// StatusCode returns error http status code
func (e *ErrResponse) StatusCode() int {
	return e.HTTPStatusCode
}

// Description returns error message or status text
func (e *ErrResponse) Description() string {
	if e.ErrorText != "" {
		return e.ErrorText
	}

	return e.StatusText
}

// Error implement error
func (e *ErrResponse) Error() string {
	return e.Description()
}

// Cause returns parent error
func (e *ErrResponse) Cause() error {
	return e.Err
}

// HTTPError implements
func (e *ErrResponse) HTTPError() *ErrResponse {
	return e
}
