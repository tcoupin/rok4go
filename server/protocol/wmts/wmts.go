package wmts

import (
	"fmt"
	"github.com/tcoupin/rok4go/server/protocol/http"
	nethttp "net/http"
)

func CheckCommons(req *nethttp.Request) (http.HTTPResponse, bool) {
	// Service
	if req.Form.Get("service") == "" {
		return MissingParameterValue("service").HTTPResponse(), true
	}
	if req.Form.Get("service") != "WMTS" {
		return InvalidParameterValue("service", fmt.Sprintf("Service %s is not supported", req.Form.Get("service"))).HTTPResponse(), true
	}
	// Request
	if req.Form.Get("request") == "" {
		return MissingParameterValue("request").HTTPResponse(), true
	}
	if req.Form.Get("request") != "GetCapabilities" && req.Form.Get("request") != "GetTile" {
		return InvalidParameterValue("request", fmt.Sprintf("Request %s is not supported", req.Form.Get("request"))).HTTPResponse(), true
	}
	return http.HTTPResponse{}, false
}

func GetCapabilities(req *nethttp.Request) http.HTTPResponse {
	capabilities := http.NewHTTPResponse()
	capabilities.BodyString("GetCap here")
	return capabilities
}

func CheckGetTile(req *nethttp.Request) (http.HTTPResponse, bool) {
	return http.HTTPResponse{}, false
}

func GetTile(req *nethttp.Request) http.HTTPResponse {
	tile := http.NewHTTPResponse()
	tile.BodyString("GetTile here")
	return tile
}
