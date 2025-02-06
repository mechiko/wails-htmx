package ginserver

import (
	"firstwails/domain"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func New(a domain.IApp, port string) {
	// Create and configure the HTTP server.
	s := http.Server{
		Addr:         fmt.Sprintf("%s:%d", "127.0.0.1", port),
		Handler:      routes.UrlPatterns(a),
		ErrorLog:     slog.NewLogLogger(a.Logger.Handler(), slog.LevelError),
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
	}

	a.Logger.Info("Starting app", "addr", s.Addr)

	if err := s.ListenAndServe(); err != nil {
		a.Logger.Error(err.Error())
		os.Exit(1)
	}
}
