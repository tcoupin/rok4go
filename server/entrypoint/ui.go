package entrypoint

import (
	"github.com/tcoupin/rok4go/resources"
	"net/http"
)

func NewUIHandler() http.Handler {
	return http.FileServer(resources.Assets)
}
