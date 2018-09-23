package http

import (
	"bytes"
	"net/http"
)

// HTTPResponse store HTTP response
type HTTPResponse struct {
	code   int
	header map[string]string
	body   bytes.Buffer
}

// NewHTTPResponse creates a new HTTPResponse
func NewHTTPResponse() HTTPResponse {
	resp := HTTPResponse{}
	resp.code = 200
	resp.header = make(map[string]string)
	resp.header["Content-Type"] = "text/plain"
	return resp
}

// Code set http status code
func (r *HTTPResponse) Code(code int) {
	r.code = code
}

// Header return hearders map
func (r HTTPResponse) Header() map[string]string {
	return r.header
}

// BodyString writes string to body
func (r *HTTPResponse) BodyString(body string) {
	r.body.WriteString(body)
}

// BodyByte wrties byte array to body
func (r *HTTPResponse) BodyByte(body []byte) {
	r.body.Write(body)
}

// Write converts HTTPResponse to an http response
func (r HTTPResponse) Write(res http.ResponseWriter) {
	for k, v := range r.Header() {
		res.Header().Set(k, v)
	}
	res.WriteHeader(r.code)
	r.body.WriteTo(res)
}
