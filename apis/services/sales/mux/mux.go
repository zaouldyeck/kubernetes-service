// Package mux provides support for binding domain level routes
// to mux.
package mux

import (
	"os"

	"github.com/zaouldyeck/kubernetes-service/apis/services/sales/route/sys/checkapi"
	"github.com/zaouldyeck/kubernetes-service/foundation/web"
)

// WebAPI constructs http.Handler with routes bound.
func WebAPI(shutdown chan os.Signal) *web.App {
	mux := web.NewApp(shutdown)

	checkapi.Routes(mux)

	return mux
}
