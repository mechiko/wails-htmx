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
		mm := usecase.New(rdc).DbInfoModel(rdc.Reductor().Model())
		msg.Model = &mm
		rdc.Reductor().ChanIn() <- msg
	case "dbinfo":
		msg.Cmd = "dbinfo"
		msg.Sender = "effects.dbinfo"
		mm := usecase.New(rdc).DbInfoModel(rdc.Reductor().Model())
		msg.Model = &mm
		rdc.Reductor().ChanIn() <- msg
	case "startup":
		msg.Cmd = rdc.Configuration().Application.StartPage
		msg.Sender = "effects.startup"
		rdc.ChanIn() <- msg
	default:
		rdc.Logger().Errorf("%s cmd %s not found", modError, msg.Cmd)
	}
}
