package entrypoint

import (
	"fmt"
	"github.com/tcoupin/rok4go/objects"
	"github.com/tcoupin/rok4go/server/protocol/http"
	"github.com/tcoupin/rok4go/server/protocol/wmts"
	nethttp "net/http"
	"strings"
)

type WMTSHandler struct {
	config *objects.Config
}

func NewWMTSHandler(config *objects.Config) WMTSHandler {
	return WMTSHandler{config: config}
}

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
