package entrypoint

import (
	"net/http"

	myhttp "github.com/tcoupin/rok4go/server/protocol/http"
)

// HTTPResponser is an interface
type HTTPResponser = func(req *http.Request) myhttp.HTTPWriter

// HandleHTTPResponse write HTTP response
func HandleHTTPResponse(next HTTPResponser) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		resp := next(req)
		resp.Write(res)
	}
}
