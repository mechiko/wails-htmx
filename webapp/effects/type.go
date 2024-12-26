package effects

import (
	"firstwails/domain"
	"firstwails/zaplog"
	"sync"

	"go.uber.org/zap"
)

const modError = "pkg:effects"

type IApp interface {
	Logger() *zap.SugaredLogger
	Repo() domain.Repo
	Reductor() domain.Reductor
	Effects() domain.Effects
	Configuration() *domain.Configuration
}

type effects struct {
	IApp
	mutex  sync.Mutex
	in     chan domain.Message
	logger *zap.SugaredLogger
}

var _ domain.Effects = &effects{}

// регистрируем тут функцию обновления ГУЯ
// создаем начальную модель которую будем тут и соблюдать
func New(app IApp) *effects {
	return &effects{
		IApp:   app,
		in:     make(chan domain.Message, 5),
		logger: zaplog.Reductor.Sugar(),
	}
}

func (rdc *effects) ChanIn() chan domain.Message {
	return rdc.in
}

func (rdc *effects) Logger() *zap.SugaredLogger {
	return rdc.logger
}
