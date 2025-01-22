package webapp

import (
	"net/http"
	"runtime/debug"
)

// ServerError logs the provided server-side error, including the request method, URL, and stack trace, then sends a 500 response.
func (a *webapp) ServerError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		url    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)

	a.Logger().Error(err.Error(), "method", method, "url", url, "trace", trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// ClientError logs a client-side error with method, URL, and status, then sends an HTTP response with the provided status.
func (a *webapp) ClientError(w http.ResponseWriter, r *http.Request, status int, err error) {
	var (
		method = r.Method
		url    = r.URL.RequestURI()
	)
	a.Logger().Error(err.Error(), "method", method, "url", url, "status", status)
	http.Error(w, http.StatusText(status), status)
}
