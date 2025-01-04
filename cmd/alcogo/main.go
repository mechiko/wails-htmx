package main

import (
	"context"
	_ "embed"
	"firstwails/checkdbg"
	"firstwails/domain"
	"firstwails/httpserver"
	"firstwails/repo"
	"firstwails/utility"
	"firstwails/webapp"
	"firstwails/zaplog"
	"fmt"
	"os"
	"path/filepath"
	"time"

	// "github.com/brpaz/echozap"
	"github.com/karagenc/zap4echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/r3labs/sse/v2"
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

	serverSse := sse.New()       // create SSE broadcaster server
	serverSse.AutoReplay = false // do not replay messages for each new subscriber that connects
	_ = serverSse.CreateStream("ready")
	go func(s *sse.Server) {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				s.Publish("ready", &sse.Event{
					Data: []byte("time: " + time.Now().Format(time.RFC3339Nano)),
				})
			}
		}
	}(serverSse)

	e := echo.New()
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

	// инитим роутер для http, конфиг и прочее
	webApp := webapp.NewWebApp(zaplog.LoggerShugar, e, serverSse, dir)
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

	port := webApp.Configuration().HostPort
	if port == "" || port == "auto" {
		if portFree, err := utility.GetFreePort(); err == nil {
			port = fmt.Sprintf("%d", portFree)
			webApp.Config().Set("hostport", port, true)
		}
	}
	loger.Infof("http port %s", port)
	host := webApp.Configuration().Hostname
	if host == "" {
		host = "127.0.0.1"
	}
	httpServer := httpserver.New(e, httpserver.HostPort(host, port))
	group.Go(func() error {
		go func() {
			<-groupCtx.Done()
			loger.Debugf("%s получен сигнал завершения контекста группы в HTTP", modError)
			if err := httpServer.Shutdown(); err != nil {
				loger.Debugf("%s Stopped http server with error:", modError, err)
			}
		}()
		httpServer.Start()
		// по ошибке сервера возвращаем в группу код ошибки
		return <-httpServer.Notify()
	})

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
	// ожидание завершения всех в группе
	if err := group.Wait(); err != nil {
		panic(err.Error())
	} else {
		fmt.Println("game over!")
	}
}
