package reductor

import "firstwails/domain"

func (rdc *reductor) proccessMessage(msg domain.Message) {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		rdc.logger.Error(fmt.Errorf("%s panic %v", modError, r))
	// 	}
	// }()
	rdc.mutex.Lock()
	defer rdc.mutex.Unlock()

	rdc.logger.Debugf("%s proccess %s from %s len chain %d", modError, msg.Cmd, msg.Sender, len(rdc.in))
	// если в сообщении модель nil берем текущую
	if msg.Model == nil {
		msg.Model = rdc.model
	}
	switch msg.Cmd {
	case "model":
		// set model to reductor
		mdl := *msg.Model
		rdc.model = &mdl
	case "startup":
		mdl := *msg.Model
		rdc.model = &mdl
		msg2 := domain.Message{
			Sender: "reductor.StartUp",
			Cmd:    rdc.ActivePage(),
			Model:  msg.Model,
		}
		rdc.Effects().ChanIn() <- msg2
	case "dbinfo":
		mdl := *msg.Model
		_, newModel := rdc.UpdaterGUI("dbinfo", mdl)
		rdc.model = &newModel
	case "stats":
		mdl := *msg.Model
		_, newModel := rdc.UpdaterGUI("stats", mdl)
		rdc.model = &newModel
	default:
		rdc.Logger().Errorf("%s cmd %s not found", modError, msg.Cmd)
	}
}
