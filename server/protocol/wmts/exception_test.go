package wmts

import (
	"github.com/stretchr/testify/assert"
	"github.com/tcoupin/rok4go/protocol/http"
	nethttp "net/http"
	"testing"
)

func TestExceptionCode(t *testing.T) {
	ec := TESTING
	assert.Equal(t, nethttp.StatusTeapot, ec.HttpCode())
	assert.Equal(t, "Testing", ec.String())
}

func TestWMTSException(t *testing.T) {
	ex := MissingParameterValue("foobar")
	r := ex.HTTPResponse()
	expected := http.NewHTTPResponse()
	expected.Code(ex.Code.HttpCode())
	expected.Header()["Content-type"] = "application/xml"
	expected.BodyString(ex.String())
	assert.Equal(t, expected, r)
}

func TestMissingParameterValue(t *testing.T) {
	ex := MissingParameterValue("foobar")

	assert.Equal(t, MISSING_PARAMETER_VALUE, ex.Code)
	assert.Equal(t, nethttp.StatusBadRequest, ex.Code.HttpCode())
	assert.Equal(t, "foobar", ex.Locator)
	assert.Equal(t, "Parameter foobar is missing", ex.Text)
}

func TestInvalidParameterValue(t *testing.T) {
	ex := InvalidParameterValue("foobar", "barfoo")

	assert.Equal(t, INVALID_PARAMETER_VALUE, ex.Code)
	assert.Equal(t, nethttp.StatusBadRequest, ex.Code.HttpCode())
	assert.Equal(t, "foobar", ex.Locator)
	assert.Equal(t, "barfoo", ex.Text)
}

func TestNoApplicableCode(t *testing.T) {
	ex := NoApplicableCode("foobar")

	assert.Equal(t, NO_APPLICABLE_CODE, ex.Code)
	assert.Equal(t, nethttp.StatusInternalServerError, ex.Code.HttpCode())
	assert.Empty(t, ex.Locator)
	assert.Equal(t, "foobar", ex.Text)
}
