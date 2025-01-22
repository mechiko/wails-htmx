// Package routes provided the application routes
package routes

import (
	"firstwails/domain"
	"firstwails/embeded"
	"firstwails/httpserver/middleware"
	"net/http"

	"github.com/justinas/alice"
)

// UrlPatterns sets up and returns an HTTP handler with predefined routing patterns for the given application.
func UrlPatterns(app domain.IApp) http.Handler {

	// Create a new HTTP request multiplexer to manage URL-to-handler routing.
	mux := http.NewServeMux()

	// Instantiate the middleware chain for handling requests with recovery capabilities.
	m := middleware.NewMiddleware(app)

	// Setup dynamic middleware
	dynamic := alice.New(m.CsrfToken)

	// Serve static and media files under /static/ and /uploads/ path.
	mux.Handle("GET /assets/", http.FileServer(embeded.GetFileSystem()))

	// Public routes for pages
	mux.Handle("GET /{$}", dynamic.Then())

	// Setup standard middleware
	standardMiddleware := alice.New(m.Recover, m.Logging, m.Headers)

	return standardMiddleware.Then(mux)
}
