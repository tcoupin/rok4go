package api

import (
	myhttp "github.com/tcoupin/rok4go/server/protocol/http"
	"net/http"
)

type APIException struct {
	Code    int
	Status  string
	Message string
}

func exception(code int, message string) (json myhttp.JSONResponse) {
	json = myhttp.NewJSONResponse()
	json.Code(code)
	json.SetData(APIException{
		Code:    code,
		Status:  http.StatusText(code),
		Message: message,
	})
	return
}

func BadRequest(message string) myhttp.JSONResponse {
	return exception(http.StatusBadRequest, message)
}

func InternalServerError(message string) myhttp.JSONResponse {
	return exception(http.StatusInternalServerError, message)
}

func MethodNotAllowed(message string) myhttp.JSONResponse {
	return exception(http.StatusMethodNotAllowed, message)
}

func Ok(message string) myhttp.JSONResponse {
	return exception(http.StatusOK, message)
}
