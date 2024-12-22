package main

import (
	"embed"
	"firstwails/dbsrc"
	"firstwails/domain"
	"firstwails/scrolling"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist components
var assets embed.FS

//go:embed build/appicon.png
var icon []byte
var version = "0.0.0"

func main() {

	loger := log.New(os.Stdout, "", log.LstdFlags)
	loger.Println("starting!")

	db := dbsrc.New(loger)
	defer db.Close()

	cols := scrolling.NewColumns()
	cols.SetMaxWidth(5000)
	initCols(cols)
	filter := domain.NewTTNFilter()
	sc := scrolling.NewScrolling(loger, db.TTN, cols, filter, 50)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"authorization", "Content-Type"},
		AllowCredentials: true,
		AllowMethods:     []string{echo.OPTIONS, echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Renderer = sc
	// e.GET("/index.html", sc.BlocksHandler)
	e.GET("/blocks", sc.BlocksHandler)
	e.GET("/input", sc.InputHandler)
	e.GET("/cols", sc.ColsHandler)
	e.GET("/colssubmit", sc.ColsSubmit)
	// e.GET("/", sc.BlocksHandler)
	// e.GET("/*", sc.BlocksHandler)
	// go func() { e.Logger.Fatal(e.Start("127.0.0.1:8888")) }()

	// Create an instance of the app structure and custom Middleware
	app := NewApp()

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
			Assets: assets,
			Middleware: func(next http.Handler) http.Handler {
				// устанавливаем обработку not found на предлагаемую по умолчанию wails
				// это произойдет когда наш роутер не найдет нужного
				e.RouteNotFound("/*", func(c echo.Context) error {
					next.ServeHTTP(c.Response(), c.Request())
					return nil
				})
				return e
			},
			// Handler: e,
		},
		Menu:             app.applicationMenu(),
		Logger:           nil,
		LogLevel:         logger.DEBUG,
		OnStartup:        app.startup,
		OnDomReady:       app.domReady,
		OnBeforeClose:    app.beforeClose,
		OnShutdown:       app.shutdown,
		WindowStartState: options.Normal,
		Bind: []interface{}{
			app,
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
			// DisableFramelessWindowDecorations: false,
			WebviewUserDataPath: "",
			ZoomFactor:          1.0,
		},
		// Mac platform specific options
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "firstwails",
				Message: "",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
