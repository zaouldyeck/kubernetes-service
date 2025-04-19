package web

// MidHandler is a handler function that runs boilerplate code
// before or after another Handler.
type MidHandler func(Handler) Handler

// wrapMiddleware creates a new handler by wrapping middleware
// around a final handler.
func wrapMiddleware(mw []MidHandler, handler Handler) Handler {
	// Loop backwards through array of middlwares, in order	to set
	// correct order of execution for incoming requests.
	for i := len(mw) - 1; i >= 0; i-- {
		mwFunc := mw[i]
		if mwFunc != nil {
			handler = mwFunc(handler)
		}
	}

	return handler
}
