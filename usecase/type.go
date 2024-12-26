package usecase

import (
	"firstwails/domain"

	"go.uber.org/zap"
)

type IApp interface {
	Logger() *zap.SugaredLogger
	Configuration() *domain.Configuration
	Reductor() domain.Reductor
	Effects() domain.Effects
	Repo() domain.Repo
}

type usecase struct {
	IApp
}

func New(app IApp) *usecase {
	return &usecase{
		IApp: app,
	}
}
