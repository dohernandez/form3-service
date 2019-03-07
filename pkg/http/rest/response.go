package rest

import "net/http"

type emptyRenderer int

func (er emptyRenderer) Render(w http.ResponseWriter, _ *http.Request) error {
	code := int(er)

	w.WriteHeader(code)

	return nil
}

// NoContent is an empty response
const NoContent = emptyRenderer(http.StatusNoContent)
