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
	case "startup":
		mdl := *msg.Model
		rdc.model = &mdl
		msg := domain.Message{
			Sender: "reductor.Startup",
			Cmd:    "home",
			Model:  rdc.model,
		}
		rdc.Effects().ChanIn() <- msg
	case "home":
		mdl := *msg.Model
		_, newModel := rdc.UpdaterGUI("home", mdl)
		rdc.model = &newModel
	}
}
