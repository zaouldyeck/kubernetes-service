// Package web. Web framework extension.
package web

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

// A Handler type handles a http request within our framework.
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// App configures context for our http handlers.
type App struct {
	*http.ServeMux
	shutdown chan os.Signal
	mw       []MidHandler
}

// NewApp instantiates an App value, handling routes for the application.
func NewApp(shutdown chan os.Signal, mw ...MidHandler) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		shutdown: shutdown,
		mw:       mw,
	}
}

// HandleFunc sets handler function for a given HTTP method and path for the servemux.
func (a *App) HandleFunc(pattern string, handler Handler, mw ...MidHandler) {
	handler = wrapMiddleware(mw, handler)
	handler = wrapMiddleware(a.mw, handler)

	h := func(w http.ResponseWriter, r *http.Request) {
		if err := handler(r.Context(), w, r); err != nil {
			// ERR HANDLING HERE
			fmt.Println(err)
			return
		}
	}

	a.ServeMux.HandleFunc(pattern, h)
}
