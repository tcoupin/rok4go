package entrypoint

import (
	"net/http"

	"github.com/tcoupin/rok4go/resources"
)

// NewUIHandler creates handler for UI
func NewUIHandler() http.Handler {
	return http.FileServer(resources.Assets)
}
