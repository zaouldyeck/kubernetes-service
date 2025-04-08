package checkapi

import (
	"net/http"
)

// Routes adds routes for checkapi.
func Routes(mux *http.ServeMux) {
	mux.HandleFunc("GET /liveness", liveness)
	mux.HandleFunc("GET /readiness", readiness)
}
