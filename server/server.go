package server

import (
	"net/http"

	"github.com/tcoupin/rok4go/objects"
	"github.com/tcoupin/rok4go/server/entrypoint"
)

// NewServer create server using provided Config
func NewServer(listen string, config *objects.Config) *http.Server {
	mux := http.NewServeMux()

	mux.Handle("/", entrypoint.NewUIHandler())
	mux.Handle("/wmts/", entrypoint.NewWMTSHandler(config))
	mux.Handle("/api/", http.RedirectHandler("/api/v1/", http.StatusFound))
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", entrypoint.NewAPIV1Handler(config)))

	server := &http.Server{Addr: listen, Handler: mux}

	return server
}
