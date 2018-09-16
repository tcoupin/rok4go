package http

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	data         interface{}
	httpresponse HTTPResponse
}

type emptystruct struct{}

func NewJSONResponse() JSONResponse {
	resp := JSONResponse{httpresponse: NewHTTPResponse()}
	return resp
}

func (r *JSONResponse) Code(code int) {
	r.httpresponse.Code(code)
}

func (r JSONResponse) Header() map[string]string {
	return r.httpresponse.Header()
}

func (r JSONResponse) Data() interface{} {
	return r.data
}

func (r *JSONResponse) SetData(data interface{}) {
	r.data = data
}

func (r JSONResponse) Write(res http.ResponseWriter) {
	r.httpresponse.Header()["Content-type"] = "application/json"

	if r.data == nil {
		r.data = emptystruct{}
	}
	bb := &bytes.Buffer{}
	enc := json.NewEncoder(bb)
	enc.Encode(r.data)
	r.httpresponse.BodyByte(bb.Bytes())
	r.httpresponse.Write(res)
}
