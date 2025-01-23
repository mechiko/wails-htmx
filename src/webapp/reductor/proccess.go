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
	rdc.logger.Debugf("%s excel chunk size %s", modError, rdc.model.Stats.ExcelChunkSize)
	// если в сообщении модель nil берем текущую
	if msg.Model == nil {
		msg.Model = rdc.model
	}
	switch msg.Cmd {
	case "model":
		// set model to reductor
		mdl := *msg.Model
		rdc.model = &mdl
	case "trueclient":
		// set model to reductor only trueclient struct
		mdl := *msg.Model
		rdc.model.TrueClient = mdl.TrueClient
	case "startup":
		mdl := *msg.Model
		rdc.model = &mdl
	case "dbinfo":
		mdl := *msg.Model
		_, newModel := rdc.UpdaterGUI("dbinfo", mdl)
		rdc.model = &newModel
	case "stats":
		mdl := *msg.Model
		_, newModel := rdc.UpdaterGUI("stats", mdl)
		rdc.model = &newModel
	case "setup":
		mdl := *msg.Model
		_, newModel := rdc.UpdaterGUI("setup", mdl)
		rdc.model = &newModel
	default:
		rdc.Logger().Errorf("%s cmd %s not found", modError, msg.Cmd)
	}
}
