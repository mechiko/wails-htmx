package domain

import (
	"go.uber.org/zap"
)

type IApp interface {
	Logger() *zap.SugaredLogger
	Config() IConfig
	Reductor() Reductor
	Effects() Effects
}
