package domain

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
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
	ActivePageTitle() string
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
	StartTrueClient(*Model) ITrueClient
	// Render(w io.Writer, name string, data interface{}, c echo.Context) error
	StartDateString() string
	EndDateString() string
	SetFsrarID(string)
	FsrarID() string
	ServerError(w http.ResponseWriter, r *http.Request, err error)
	ClientError(w http.ResponseWriter, r *http.Request, status int, err error)
	Open(url string)
	SessionManager() *scs.SessionManager
}
