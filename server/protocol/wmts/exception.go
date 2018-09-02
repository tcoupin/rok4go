package wmts

import (
	"fmt"
	"github.com/tcoupin/rok4go/server/protocol/http"
	nethttp "net/http"
)

type ExceptionCode int

const (
	TESTING ExceptionCode = iota
	MISSING_PARAMETER_VALUE
	INVALID_PARAMETER_VALUE
	NO_APPLICABLE_CODE
)

var code_string = map[ExceptionCode]string{
	TESTING:                 "Testing",
	MISSING_PARAMETER_VALUE: "MissingParameterValue",
	INVALID_PARAMETER_VALUE: "InvalidParameterValue",
	NO_APPLICABLE_CODE:      "NoApplicableCode",
}

var code_httpcode = map[ExceptionCode]int{
	TESTING:                 nethttp.StatusTeapot,
	MISSING_PARAMETER_VALUE: nethttp.StatusBadRequest,
	INVALID_PARAMETER_VALUE: nethttp.StatusBadRequest,
	NO_APPLICABLE_CODE:      nethttp.StatusInternalServerError,
}

func InvalidParameterValue(param string, text string) WMTSException {
	return WMTSException{Code: INVALID_PARAMETER_VALUE, Locator: param, Text: text}
}

func MissingParameterValue(param string) WMTSException {
	return WMTSException{Code: MISSING_PARAMETER_VALUE, Locator: param, Text: "Parameter " + param + " is missing"}
}
func NoApplicableCode(text string) WMTSException {
	return WMTSException{Code: NO_APPLICABLE_CODE, Text: text}
}

func (ex ExceptionCode) String() string {
	return code_string[ex]
}

func (ex ExceptionCode) HttpCode() int {
	return code_httpcode[ex]
}

type WMTSException struct {
	Code    ExceptionCode
	Locator string
	Text    string
}

func (ex WMTSException) String() string {
	locator := ""
	if ex.Locator != "" {
		locator = " locator=\"" + ex.Locator + "\""
	}
	return fmt.Sprintf(
		`<ExceptionReport xmlns="http://www.opengis.net/ows/1.1"><Exception exceptionCode="%s"%s>%s</Exception></ExceptionReport>`,
		ex.Code,
		locator,
		ex.Text,
	)
}

func (ex WMTSException) HTTPResponse() http.HTTPResponse {
	resp := http.NewHTTPResponse()
	resp.Code(ex.Code.HttpCode())
	resp.Header()["Content-type"] = "application/xml"
	resp.BodyString(ex.String())
	return resp
}
