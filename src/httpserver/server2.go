// Package httpserver implements HTTP server.
package httpserver

import (
	"context"
	"mime"
	"net/http"
	"time"

	// "github.com/brpaz/echozap"

	"firstwails/embeded"
	"firstwails/zaplog"

	"github.com/karagenc/zap4echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server -.
type Server2 struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
	echo            *echo.Echo
}

// New -.
func New2(e http.Handler, opts ...Option) *Server {
	// HTTP Server
	// mime.AddExtensionType(".js", "application/javascript")
	mime.AddExtensionType(".js", "application/javascript; charset=utf-8")
	mime.AddExtensionType(".css", "text/css; charset=utf-8")

	e.Use(middleware.Recover())
	// e.Use(echozap.ZapLogger(zaplog.EchoSugar))
	e.Use(
		zap4echo.Logger(zaplog.EchoSugar),
		zap4echo.Recover(zaplog.EchoSugar),
	)
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}\n",
	// }))
	// e.Logger.SetOutput(os.Stdout)
	// 	e.Logger.SetOutput(ioutil.Discard)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"authorization", "Content-Type"},
		AllowCredentials: true,
		AllowMethods:     []string{echo.OPTIONS, echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.Pre(middleware.Rewrite(map[string]string{
		// "/spa/requestrests/*":   "/",
		// "/spa/requestclients/*": "/",
		// "/test":                 "/",
	}))

	// Prometheus metrics
	// p := prometheus.NewPrometheus("echo", nil)
	// p.Use(e)

	httpServer := &http.Server{
		Handler:      e,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         _defaultAddr,
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
		echo:            e,
	}

	// Custom options
	for _, opt := range opts {
		opt(s)
	}

	assetHandler := http.FileServer(embeded.GetFileSystem())
	e.GET("/*", echo.WrapHandler(assetHandler))

	return s
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()
	err := s.server.Shutdown(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) HandlerEcho() *echo.Echo {
	return s.echo
}
