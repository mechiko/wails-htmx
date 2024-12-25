package effects

import "firstwails/domain"

func (rdc *effects) proccessMessage(msg domain.Message) {
	rdc.mutex.Lock()
	defer rdc.mutex.Unlock()

	rdc.logger.Debugf("%s proccess %s from %s len chain %d", modError, msg.Cmd, msg.Sender, len(rdc.in))
	switch msg.Cmd {
	case "home":
		// отправим состояние в активную страницу если не будет ее то проигнорируется
		msg.Cmd = "page"
		msg.Sender = "effects.home"
		if msg.Model == nil {
			newModel := rdc.Reductor().Model()
			msg.Model = &newModel
		}
		msg.Model.Home.UtmHost = rdc.Configuration().UtmHost
		msg.Model.Home.UtmPort = rdc.Configuration().UtmPort
		msg.Model.Gui.MainWindow.StatusBar.Fsrarid = rdc.Configuration().Application.Fsrarid
		msg.Model.Home.Output = rdc.Configuration().Output
		msg.Model.Home.Export = rdc.Configuration().Export
		msg.Model.Home.DbConfigDesc = rdc.Repo().Dbs().Config().String()
		msg.Model.Home.DbLiteDesc = rdc.Repo().Dbs().Lite().String()
		msg.Model.Home.DbZnakDesc = rdc.Repo().Dbs().Znak().String()
		msg.Model.Home.DbA3Desc = rdc.Repo().Dbs().A3().String()
		rdc.Reductor().ChanIn() <- msg
	case "first":
		// отправим состояние в активную страницу если не будет ее то проигнорируется
		msg.Cmd = "all"
		msg.Sender = "effects.first"
		mm := rdc.Reductor().Model()
		msg.Model = &mm
		msg.Model.Home.UtmHost = rdc.Configuration().UtmHost
		msg.Model.Home.UtmPort = rdc.Configuration().UtmPort
		msg.Model.Home.Output = rdc.Configuration().Output
		msg.Model.Home.Export = rdc.Configuration().Export
		msg.Model.Gui.MainWindow.StatusBar.License = rdc.Licenser().IsActive()
		msg.Model.Gui.MainWindow.StatusBar.Utm = false
		msg.Model.Gui.MainWindow.StatusBar.Fsrarid = rdc.Configuration().Application.Fsrarid
		rdc.Reductor().ChanIn() <- msg
	case "license":
		services.New(rdc.IApp).LicenseState()
	case "utm":
		services.New(rdc.IApp).UtmState()
	}
}
