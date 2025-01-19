package domain

import (
	"context"
)

type Repo interface {
	Start(context.Context) error
	Shutdown()
	// получаем интерфейсы к базам данных через которые осуществляем доступ из приложения
	DbLite() DbLite
	DbZnak() DbZnak
	DbA3() DbAlcohelp3
	DbConfig() DbConfig
	Dbs() Dbs
	IsA3() bool
	IsZnak() bool
	IsConfig() bool
}
