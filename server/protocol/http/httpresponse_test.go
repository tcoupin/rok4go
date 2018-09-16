package http

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestNewHTTPResponse(t *testing.T) {
	resp := NewHTTPResponse()
	assert.Equal(t, 200, resp.code)
	assert.Equal(t, map[string]string{"Content-Type": "text/plain"}, resp.header)
	assert.Empty(t, resp.body.Bytes())
}

func TestCode(t *testing.T) {
	resp := NewHTTPResponse()
	resp.Code(666)
	assert.Equal(t, 666, resp.code)
}

func TestHeader(t *testing.T) {
	resp := NewHTTPResponse()
	resp.Header()["foo"] = "bar"
	assert.Equal(t, "bar", resp.header["foo"])
}

func TestBodyString(t *testing.T) {
	resp := NewHTTPResponse()
	resp.BodyString("foobar")
	assert.Equal(t, []byte("foobar"), resp.body.Bytes())
}

func TestBodyByte(t *testing.T) {
	resp := NewHTTPResponse()
	resp.BodyByte([]byte("foobar"))
	assert.Equal(t, []byte("foobar"), resp.body.Bytes())
}

func TestWrite(t *testing.T) {
	rw := httptest.NewRecorder()

	resp := NewHTTPResponse()
	resp.BodyString("Todo...")
	resp.Code(666)
	resp.Write(rw)

	assert.Equal(t, 666, rw.Code)

	for k, v := range resp.header {
		assert.Contains(t, rw.Header(), k)
		assert.Equal(t, rw.Header()[k][0], v)
	}

	assert.Equal(t, "Todo...", rw.Body.String())
}
