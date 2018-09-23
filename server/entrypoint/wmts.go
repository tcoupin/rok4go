package entrypoint

import (
	"fmt"
	nethttp "net/http"
	"strings"

	"github.com/tcoupin/rok4go/objects"
	"github.com/tcoupin/rok4go/server/protocol/http"
	"github.com/tcoupin/rok4go/server/protocol/wmts"
)

// WMTSHandler is an handle for WMTS protocol
type WMTSHandler struct {
	config *objects.Config
}

// NewWMTSHandler creates handler for WMTS
func NewWMTSHandler(config *objects.Config) WMTSHandler {
	return WMTSHandler{config: config}
}

// ServeHTTP handles request
func (w WMTSHandler) ServeHTTP(res nethttp.ResponseWriter, req *nethttp.Request) {
	if err := req.ParseForm(); err != nil {
		wmts.NoApplicableCode(fmt.Sprint(err)).HTTPResponse().Write(res)
		return
	}

	// Parameters key lowercase
	for k, v := range req.Form {
		req.Form[strings.ToLower(k)] = v
	}

	c := make(chan http.HTTPResponse)
	go processRequest(req, c)
	resp := <-c
	resp.Write(res)
}

func processRequest(req *nethttp.Request, c chan http.HTTPResponse) {
	if resp, mustReturn := wmts.CheckCommons(req); mustReturn {
		c <- resp
		return
	}

	if req.Form.Get("request") == "GetCapabilities" {
		c <- wmts.GetCapabilities(req)
		return
	} else if req.Form.Get("request") == "GetTile" {
		if resp, mustReturn := wmts.CheckGetTile(req); mustReturn {
			c <- resp
			return
		}
		c <- wmts.GetTile(req)
		return
	}
}
