package api

import (
	"net/http"

	myhttp "github.com/tcoupin/rok4go/server/protocol/http"
)

type apiResponse struct {
	Code    int
	Status  string
	Message string
}

func response(code int, message string) (json myhttp.JSONResponse) {
	json = myhttp.NewJSONResponse()
	json.Code(code)
	json.SetData(apiResponse{
		Code:    code,
		Status:  http.StatusText(code),
		Message: message,
	})
	return
}

// BadRequest returns JSONRespnse with status BadRequest
func BadRequest(message string) myhttp.JSONResponse {
	return response(http.StatusBadRequest, message)
}

// InternalServerError returns JSONRespnse with status InternalServerError
func InternalServerError(message string) myhttp.JSONResponse {
	return response(http.StatusInternalServerError, message)
}

// MethodNotAllowed returns JSONRespnse with status MethodNotAllowed
func MethodNotAllowed(message string) myhttp.JSONResponse {
	return response(http.StatusMethodNotAllowed, message)
}

// Ok returns JSONRespnse with status OK
func Ok(message string) myhttp.JSONResponse {
	return response(http.StatusOK, message)
}
