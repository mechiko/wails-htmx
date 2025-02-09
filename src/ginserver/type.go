package ginserver

import (
	"context"
	"firstwails/domain"
	"firstwails/ginserver/routes"
	"firstwails/zaplog"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"go.uber.org/zap"
)

const (
	_defaultReadTimeout     = 120 * time.Second
	_defaultWriteTimeout    = 120 * time.Second
	_defaultAddr            = "127.0.0.1:8088"
	_defaultShutdownTimeout = 1 * time.Second
)

// Server -.
type server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
	// templates is a map for caching parsed templates, allowing efficient rendering and reuse within the application.
	templates map[string]*template.Template
}

func New(a domain.IApp, port string) *server {
	addr := fmt.Sprintf("%s:%s", "127.0.0.1", port)
	if port == "" {
		addr = _defaultAddr
	}
	// Create and configure the HTTP server.
	s := &http.Server{
		Addr:         addr,
		Handler:      nil,
		ErrorLog:     zap.NewStdLog(zaplog.EchoSugar),
		IdleTimeout:  time.Minute,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
	}
	ss := &server{
		server:          s,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
	}
	routes.UrlPatterns(a)
	return ss
}

func (s *server) Start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

// Notify -.
func (s *server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()
	err := s.server.Shutdown(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *server) Handler() http.Handler {
	return s.server.Handler
}

func (s *server) SetHandler(h http.Handler) {
	s.server.Handler = h
}
