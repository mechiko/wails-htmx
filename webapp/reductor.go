package webapp

import (
	"firstwails/domain"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *webapp) initReductor() {
	if a.reductor == nil {
		panic("webapp panic reductor nil")
	}
	m := domain.InitModel
	m.Home.UtmHost = a.configuration.UtmHost
	m.Home.UtmPort = a.configuration.UtmPort
	m.Gui.MainWindow.StatusBar.Fsrarid = a.configuration.Application.Fsrarid
	m.Home.Output = a.configuration.Output
	m.Home.Export = a.configuration.Export
	m.Gui.MainWindow.StatusBar.Utm = false
	msg := domain.Message{
		Sender: "webapp.initReductor",
		Cmd:    "init",
		Model:  &m,
	}
	a.reductor.ChanIn() <- msg
}

// на входе команда и модель
// на выходе новая команда и новая модель
// пока на выходе всегда пустая команда и таже модель что на входе
func (a *webapp) UpdateGUI(cmd string, model domain.Model) (string, domain.Model) {
	defer func() {
		if r := recover(); r != nil {
			a.logger.Errorf("%s Reductor UpdateGUI %s panic %v", modError, cmd, r)
		}
	}()
	// обновляем по модели текущую страницу
	runtime.WindowReload(a.ctx)
	return "", model
}
