package http

import (
	"bytes"
	"net/http"
)

type HTTPResponse struct {
	code   int
	header map[string]string
	body   bytes.Buffer
}

func NewHTTPResponse() HTTPResponse {
	resp := HTTPResponse{}
	resp.code = 200
	resp.header = make(map[string]string)
	resp.header["Content-Type"] = "text/plain"
	return resp
}

func (r *HTTPResponse) Code(code int) {
	r.code = code
}

func (r HTTPResponse) Header() map[string]string {
	return r.header
}

func (r *HTTPResponse) BodyString(body string) {
	r.body.WriteString(body)
}

func (r *HTTPResponse) BodyByte(body []byte) {
	r.body.Write(body)
}

func (r HTTPResponse) Write(res http.ResponseWriter) {
	for k, v := range r.Header() {
		res.Header().Set(k, v)
	}
	res.WriteHeader(r.code)
	r.body.WriteTo(res)
}
