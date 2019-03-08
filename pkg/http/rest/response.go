package rest

import "net/http"

type emptyRenderer int

// Render implementation of render.Renderer interface for managing response.
// Error result is always nil there is no error state for this function
// nolint:unparam
func (er emptyRenderer) Render(w http.ResponseWriter, _ *http.Request) error {
	code := int(er)

	w.WriteHeader(code)

	return nil
}

// NoContent is an empty response
const NoContent = emptyRenderer(http.StatusNoContent)
