package usecase

import (
	"firstwails/domain"
)

const modError = "usecase"

// type IApp interface {
// 	Logger() *zap.SugaredLogger
// 	Configuration() *domain.Configuration
// 	Reductor() domain.Reductor
// 	Effects() domain.Effects
// 	Repo() domain.Repo
// }

type usecase struct {
	domain.IApp
}

func New(app domain.IApp) *usecase {
	return &usecase{
		IApp: app,
	}
}
