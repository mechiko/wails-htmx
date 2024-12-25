package main

import (
	_ "embed"
	"firstwails/alcogo"
	"firstwails/dbsrc"
	"firstwails/webapp"
	"firstwails/zaplog"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

// var version = "0.0.0"
var fileExe string

func init() {
	fileExe = os.Args[0]
	dir, _ := filepath.Abs(filepath.Dir(fileExe))
	os.Chdir(dir)
	zaplog.InitializeLogger()
}

func main() {

	loger := zaplog.Logger
	loger.Debug("zaplog started")

	db := dbsrc.New(loger)
	defer db.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"authorization", "Content-Type"},
		AllowCredentials: true,
		AllowMethods:     []string{echo.OPTIONS, echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	webApp := webapp.NewWebApp(zaplog.Logger, e)
	e.Renderer = webApp

	e.GET("/page", webApp.CurrentPageIndex)

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "firstwails",
		Width:     1040,
		Height:    768,
		MinWidth:  200,
		MinHeight: 200,
		// MaxWidth:      1280,
		// MaxHeight:     800,
		DisableResize: false,
		Fullscreen:    false,
		Frameless:     false,
		// CSSDragProperty:   "widows",
		// CSSDragValue:      "1",
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		AssetServer: &assetserver.Options{
			Assets: alcogo.Assets,
			Middleware: func(next http.Handler) http.Handler {
				// устанавливаем обработку not found на предлагаемую по умолчанию wails
				// это произойдет когда наш роутер не найдет нужного
				e.RouteNotFound("/*", func(c echo.Context) error {
					next.ServeHTTP(c.Response(), c.Request())
					return echo.NewHTTPError(http.StatusInternalServerError, "not found")
				})
				return e
			},
			// Handler: e,
		},
		Menu:             webApp.ApplicationMenu(),
		Logger:           nil,
		LogLevel:         logger.DEBUG,
		OnStartup:        webApp.Startup,
		OnDomReady:       webApp.DomReady,
		OnBeforeClose:    webApp.BeforeClose,
		OnShutdown:       webApp.Shutdown,
		WindowStartState: options.Normal,
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
			// DisableFramelessWindowDecorations: false,
			WebviewUserDataPath: "",
			ZoomFactor:          1.0,
		},
		Debug: options.Debug{
			OpenInspectorOnStartup: true,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
