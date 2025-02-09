package middleware

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// CsrfToken is middleware that applies CSRF protection to HTTP requests using nosurf, setting a secure base cookie.
func (m *Middleware) CsrfToken(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
	})

	return csrfHandler
}
