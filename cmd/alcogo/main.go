package main

import (
	"context"
	_ "embed"
	"firstwails/alcogo"
	"firstwails/checkdbg"
	"firstwails/domain"
	"firstwails/repo"
	"firstwails/utility"
	"firstwails/webapp"
	"firstwails/zaplog"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	// "github.com/brpaz/echozap"
	"github.com/karagenc/zap4echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"golang.org/x/sync/errgroup"
)

const modError = "main"

// var version = "0.0.0"
var fileExe string
var dir string

func init() {
	fileExe = os.Args[0]
	dir, _ = filepath.Abs(filepath.Dir(fileExe))
	os.Chdir(dir)
	utility.PathCreate(domain.ConfigPath)
	utility.PathCreate(domain.LogPath)
	utility.PathCreate(domain.DbPath)
	zaplog.OnStartup()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	group, groupCtx := errgroup.WithContext(ctx)

	// процесс для выключения логгера
	group.Go(func() error {
		go func() {
			<-groupCtx.Done()
			zaplog.OnShutdown()
		}()
		// возврат произойдет только после прерываения контекста
		// чтобы отловить завершение программы и сброс буферов логеров
		return zaplog.Run(groupCtx)
	})

	loger := zaplog.LoggerShugar
	loger.Debug("zaplog started")

	e := echo.New()
	// e.Use(middleware.Logger())
	// e.Use(echozap.ZapLogger(zaplog.EchoSugar))
	e.Use(
		zap4echo.Logger(zaplog.EchoSugar),
		zap4echo.Recover(zaplog.EchoSugar),
	)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"authorization", "Content-Type"},
		AllowCredentials: true,
		AllowMethods:     []string{echo.OPTIONS, echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	webApp := webapp.NewWebApp(zaplog.LoggerShugar, e, dir)
	e.Renderer = webApp

	// инициализируем REPO
	if repo, err := repo.New(webApp); err != nil {
		loger.Errorf("%s %s", modError, err.Error())
		panic(err.Error())
	} else {
		webApp.SetRepo(repo)
		group.Go(func() error {
			go func() {
				<-groupCtx.Done()
				repo.Shutdown()
			}()
			return repo.Start(groupCtx)
		})
	}
	// тесты
	if err := checkdbg.NewChecks(webApp).Run(); err != nil {
		loger.Errorf("check error %v", err)
		// app.MessageBox("Ошибка Checks", err.Error())
		cancel()
	}

	group.Go(func() error {
		go func() {
			<-groupCtx.Done()
			webApp.OnShutdown()
		}()
		// возврат произойдет когда Run завершится
		// он так же зависит от прерывания контекста
		// в Run запускают необходимые фоновые задачи
		return webApp.Run(groupCtx)
	})

	// Create application with options
	if err := wails.Run(&options.App{
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
		LogLevel:         logger.INFO,
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
	}); err != nil {
		loger.Errorf("%s wails error %s", modError, err.Error())
	}
	// все задачи завершаются сразу после запуска и обрабатывают только контекст
	// поэтому по закрытию окна приложения завершаем контекст
	cancel()
	// ожидание завершения всех в группе
	if err := group.Wait(); err != nil {
		panic(err.Error())
	} else {
		fmt.Println("game over!")
	}
}
