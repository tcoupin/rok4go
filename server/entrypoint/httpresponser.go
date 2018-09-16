package entrypoint

import (
	myhttp "github.com/tcoupin/rok4go/server/protocol/http"
	"net/http"
)

type HTTPResponser = func(req *http.Request) myhttp.HTTPWriter

func HandleHTTPResponse(next HTTPResponser) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		resp := next(req)
		resp.Write(res)
	}
}
