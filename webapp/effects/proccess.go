package effects

import "firstwails/domain"

func (rdc *effects) proccessMessage(msg domain.Message) {
	rdc.mutex.Lock()
	defer rdc.mutex.Unlock()

	rdc.logger.Debugf("%s proccess %s from %s len chain %d", modError, msg.Cmd, msg.Sender, len(rdc.in))
	switch msg.Cmd {
	case "home":
		// отправим состояние в активную страницу если не будет ее то проигнорируется
	case "startup":
		// отправим состояние в активную страницу если не будет ее то проигнорируется
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
