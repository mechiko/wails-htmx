package effects

import (
	"firstwails/domain"
	"firstwails/usecase"
)

func (rdc *effects) proccessMessage(msg domain.Message) {
	rdc.mutex.Lock()
	defer rdc.mutex.Unlock()

	rdc.logger.Debugf("%s proccess %s from %s len chain %d", modError, msg.Cmd, msg.Sender, len(rdc.in))
	switch msg.Cmd {
	case "stats":
		msg.Cmd = "stats"
		msg.Sender = "effects.stats"
		mm := usecase.New(rdc).HomeModel(rdc.Reductor().Model())
		msg.Model = &mm
		rdc.Reductor().ChanIn() <- msg
	case "home":
		msg.Cmd = "home"
		msg.Sender = "effects.home"
		mm := usecase.New(rdc).HomeModel(rdc.Reductor().Model())
		msg.Model = &mm
		rdc.Reductor().ChanIn() <- msg
	case "startup":
		msg.Cmd = "startup"
		msg.Sender = "effects.startup"
		mm := rdc.Reductor().Model()
		msg.Model = &mm
		msg.Model.Home.UtmHost = rdc.Configuration().UtmHost
		msg.Model.Home.UtmPort = rdc.Configuration().UtmPort
		msg.Model.Home.Output = rdc.Configuration().Output
		msg.Model.Home.Export = rdc.Configuration().Export
		msg.Model.Gui.MainWindow.StatusBar.Utm = false
		msg.Model.Gui.MainWindow.StatusBar.Fsrarid = rdc.Configuration().Application.Fsrarid
		rdc.Reductor().ChanIn() <- msg
	}
}
