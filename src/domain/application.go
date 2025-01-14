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
}
