package platform

import "errors"

// ErrWrongRequestVersion is a common wrong request error when casting
var ErrWrongRequestVersion = errors.New("wrong request version")

// ErrMismatchRequest is a common mismatch error when casting aggregate root
var ErrMismatchRequest = errors.New("mismatch request")
