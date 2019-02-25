package http

import (
	"net/http"
	"os"
	"path/filepath"
)

// DocsHandler struct to handler request to the /docs endpoint...
type DocsHandler struct {
	prefix   string
	docsPath string
}

// NewDocsHandler returns a DocsHandler instance
func NewDocsHandler(prefix, docsPath string) *DocsHandler {
	return &DocsHandler{
		prefix:   prefix,
		docsPath: docsPath,
	}
}

func (h *DocsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	workingDir, _ := os.Getwd()
	filesDir := filepath.Join(workingDir, h.docsPath)
	root := http.Dir(filesDir)

	fs := http.StripPrefix("/docs", http.FileServer(root))
	fs.ServeHTTP(w, r)
}
