package entrypoint

import (
	"encoding/json"
	"fmt"
	"github.com/tcoupin/rok4go/objects"
	"github.com/tcoupin/rok4go/server/protocol/api"
	myhttp "github.com/tcoupin/rok4go/server/protocol/http"
	"github.com/tcoupin/rok4go/utils/log"
	"net/http"
)

type APIV1Handler struct {
	config *objects.Config
	mux    *http.ServeMux
}

func NewAPIV1Handler(config *objects.Config) http.Handler {
	h := APIV1Handler{config: config}

	h.mux = http.NewServeMux()
	h.mux.HandleFunc("/config/global", HandleHTTPResponse(h.globalHandler))

	return h.mux
}

func (h APIV1Handler) globalHandler(req *http.Request) myhttp.HTTPWriter {
	if err := req.ParseForm(); err != nil {
		log.ERROR("Fail to parse request params: %v", err)
		return api.InternalServerError("Fail to parse request params")
	}

	if req.Method == http.MethodGet {
		log.DEBUG("GET Global Config")
		json := myhttp.NewJSONResponse()
		json.SetData(h.config.GlobalConfig(req.Form.Get("force") != ""))
		return json

	} else if req.Method == http.MethodPost {
		log.DEBUG("POST Global Config")
		dec := json.NewDecoder(req.Body)
		defer req.Body.Close()
		var newconf objects.GlobalConfig
		if err := dec.Decode(&newconf); err != nil {
			log.ERROR("Fail to parse Global Config from body: %v", err)
			return api.BadRequest(fmt.Sprintf("Fail to parse Global Config from body: %v", err))
		}

		if err := h.config.SetGlobalConfig(&newconf); err != nil {
			log.ERROR("Fail to update Global Config: %v", err)
			return api.BadRequest(fmt.Sprintf("Fail to update Global Config: %v", err))
		}
		return api.Ok("Global config updated")
	}
	return api.MethodNotAllowed("Only GET and POST are allowed")
}
