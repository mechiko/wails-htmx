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
		// usecase делает обновление редуктора по завершению запроса
		// выполняется как подзадача потому модель обновляется там в ней
		_ = usecase.New(rdc).StatsModel(rdc.Reductor().Model())
	case "gtins":
		// usecase делает обновление редуктора по завершению запроса
		// выполняется как подзадача потому модель обновляется там в ней
		_ = usecase.New(rdc).GtinsModel(rdc.Reductor().Model())
	case "setup":
		// usecase делает обновление редуктора по завершению запроса
		// выполняется как подзадача потому модель обновляется там в ней
		_ = usecase.New(rdc).SetupModel(rdc.Reductor().Model())
	case "dbinfo":
		msg.Cmd = "dbinfo"
		msg.Sender = "effects.dbinfo"
		mm := usecase.New(rdc).DbInfoModel(rdc.Reductor().Model())
		msg.Model = &mm
		rdc.Reductor().ChanIn() <- msg
	case "startup":
		// событие из webapp.StartUp
		msg.Cmd = "startup"
		msg.Sender = "effects.startup"
		msg.Model = msgIn.Model
		rdc.Reductor().ChanIn() <- msg
	default:
		rdc.Logger().Errorf("%s cmd %s not found", modError, msgIn.Cmd)
	}
}
