package checkapi

import (
	"github.com/zaouldyeck/kubernetes-service/foundation/web"
)

// Routes adds routes for checkapi.
func Routes(app *web.App) {
	app.HandleFunc("GET /liveness", liveness)
	app.HandleFunc("GET /readiness", readiness)
}
