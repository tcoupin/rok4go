package http

import (
	"net/http"
)

type HTTPWriter interface {
	Write(res http.ResponseWriter)
}
