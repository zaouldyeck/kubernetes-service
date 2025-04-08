// Package mux provides support for binding domain level routes
// to mux.
package mux

import (
	"net/http"

	"github.com/zaouldyeck/kubernetes-service/apis/services/sales/route/sys/checkapi"
)

// WebAPI constructs http.Handler with routes bound.
func WebAPI() *http.ServeMux {
	mux := http.NewServeMux()

	checkapi.Routes(mux)

	return mux
}
