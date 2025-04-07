// Package mux provides support for binding domain level routes
// to mux.
package mux

import (
	"encoding/json"
	"net/http"
)

// WebAPI constructs http.Handler with routes bound.
func WebAPI() *http.ServeMux {
	mux := http.NewServeMux()

	h := func(w http.ResponseWriter, r *http.Request) {
		status := struct {
			Status string
		}{
			Status: "OK",
		}

		json.NewEncoder(w).Encode(status)
	}

	mux.HandleFunc("GET /test", h)

	return mux
}
