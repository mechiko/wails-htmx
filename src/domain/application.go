package domain

import (
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
	UpdateModel(Model, string)
	UpdateTrueClientModel(TrueClient, string)
	UpdatePage(Model, string, string)
	Pwd() string
	Output() string
	StartTrueClient(model Model) ITrueClient
	// Render(w io.Writer, name string, data interface{}, c echo.Context) error
	StartDateString() string
	EndDateString() string
	SetFsrarID(string)
	FsrarID() string
}
