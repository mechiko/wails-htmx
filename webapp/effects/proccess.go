package effects

import (
	"firstwails/domain"
	"firstwails/usecase"
)

func (rdc *effects) proccessMessage(msgIn domain.Message) {
	rdc.mutex.Lock()
	defer rdc.mutex.Unlock()

	rdc.logger.Debugf("%s proccess %s from %s len chain %d", modError, msgIn.Cmd, msgIn.Sender, len(rdc.in))
	msg := domain.Message{}
	switch msgIn.Cmd {
	case "stats":
		msg.Cmd = "stats"
		msg.Sender = "effects.stats"
		// mm := usecase.New(rdc).MenuModel(rdc.Reductor().Model())
		mm := usecase.New(rdc).StatsModel(rdc.Reductor().Model())
		msg.Model = &mm
		rdc.Reductor().ChanIn() <- msg
	case "dbinfo":
		msg.Cmd = "dbinfo"
		msg.Sender = "effects.dbinfo"
		mm := usecase.New(rdc).DbInfoModel(rdc.Reductor().Model())
		msg.Model = &mm
		rdc.Reductor().ChanIn() <- msg
	case "startup":
		// msg.Cmd = rdc.Configuration().Application.StartPage
		msg.Cmd = "startup"
		msg.Sender = "effects.startup"
		mm := usecase.New(rdc).InitModel(rdc.Reductor().Model())
		msg.Model = &mm
		rdc.Reductor().ChanIn() <- msg
	default:
		rdc.Logger().Errorf("%s cmd %s not found", modError, msg.Cmd)
	}
}
