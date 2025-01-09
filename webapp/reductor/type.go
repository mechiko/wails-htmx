package reductor

import (
	"sync"

	"firstwails/domain"
	"firstwails/zaplog"

	"go.uber.org/zap"
)

const modError = "pkg:reductor"

// type IApp interface {
// 	Logger() *zap.SugaredLogger
// 	Reductor() domain.Reductor
// 	Effects() domain.Effects
// 	Configuration() *domain.Configuration
// }

type reductor struct {
	domain.IApp
	mutex      sync.Mutex
	in         chan domain.Message
	logger     *zap.SugaredLogger
	effects    domain.Effects
	updaterGUI domain.FuncUpdaterGUI
	model      *domain.Model
}

var _ domain.Reductor = &reductor{}

// регистрируем тут функцию обновления ГУЯ
// создаем начальную модель которую будем тут и соблюдать
func New(app domain.IApp, efs domain.Effects) *reductor {
	return &reductor{
		IApp:    app,
		in:      make(chan domain.Message, 5),
		logger:  zaplog.Reductor.Sugar(),
		effects: efs,
		model:   &domain.InitModel,
	}
}

// возвращаем новый объект Модель (разыменовываем внутреннюю)
func (rdc *reductor) Model() domain.Model {
	rdc.mutex.Lock()
	defer rdc.mutex.Unlock()
	// разыменовываем ссылку получаем копию объекта и возвращаем
	mdl := *rdc.model
	return mdl
}

func (rdc *reductor) UpdaterGUI(cmd string, model domain.Model) (cm string, mdl domain.Model) {
	if rdc.updaterGUI != nil {
		go func() {
			rdc.updaterGUI(cmd, model)
		}()
		// return rdc.updaterGUI(cmd, model)
	}
	return "", model
}
