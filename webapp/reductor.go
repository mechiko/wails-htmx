package webapp

import (
	"firstwails/domain"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// // на входе команда и модель
// // на выходе новая команда и новая модель
// // пока на выходе всегда пустая команда и таже модель что на входе
// func (a *webapp) UpdateGUI(cmd string, model domain.Model) (string, domain.Model) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			a.logger.Errorf("%s Reductor UpdateGUI %s panic %v", modError, cmd, r)
// 		}
// 	}()
// 	// обновляем по модели текущую страницу
// 	runtime.WindowReload(a.ctx)
// 	return "", model
// }

// вызывается в редукторе для обновления страницы по имени и модели
// wails
func (a *webapp) ReductorUpdater(cmd string, model domain.Model) (string, domain.Model) {
	// пока не будет сигнала DomReady(ctx context.Context)
	// при перезагрузке окна он еще раз прилетает
	a.logger.Debugf("webapp ReductorUpdater [readyDOM:%v]", a.readyDOM)
	if !a.readyDOM {
		return "", model
	}
	runtime.WindowReload(a.ctx)
	return "", model
}

// вызывается в редукторе для обновления страницы по имени и модели
// HTTP
func (a *webapp) ReductorUpdaterHttp(cmd string, model domain.Model) (string, domain.Model) {
	a.logger.Debugf("webapp ReductorUpdaterHttp [readyDOM:%v]", a.readyDOM)
	if !a.readyDOM {
		return "", model
	}
	a.SetReloadActivePage(true)
	return "", model
}

func (a *webapp) SetActivePage(page string, reductor bool) {
	if reductor {
		model := a.reductor.Model()
		msg := domain.Message{
			Sender: "webapp.SetActivePage",
			Cmd:    page,
			Model:  &model,
		}
		a.Effects().ChanIn() <- msg
	}
	a.activePage = page
}

func (a *webapp) SetReloadActivePage(b bool) {
	a.reloadActivePage = b
}

func (a *webapp) ReloadActivePage() bool {
	return a.reloadActivePage
}
