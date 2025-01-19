package domain

import (
	"io"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type IApp interface {
	Logger() *zap.SugaredLogger
	Config() IConfig
	Configuration() *Configuration
	Reductor() Reductor
	Effects() Effects
	Repo() Repo
	SetActivePage(string)
	DomReadyHttp()
	ReloadActivePage() bool
	SetReloadActivePage(bool)
	StartUp()
	PagesInfo() ArrPageInfo
	ActivePage() string
	UpdateModel(Model)
	UpdatePage(Model, string)
	Pwd() string
	Output() string
	StartTrueClient(model Model) ITrueClient
	StartTrueClientSuz(model Model) ITrueClient
	Render(w io.Writer, name string, data interface{}, c echo.Context) error
	StartDateString() string
	EndDateString() string
	SetFsrarID(string)
	FsrarID() string
}
