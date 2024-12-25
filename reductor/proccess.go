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
	case "init":
		// чисто устанавливает модель в редуктор
		// разыменование делает новый экземпляр
		mdl := *msg.Model
		// новый экземпляр используем как текущую модель
		rdc.model = &mdl
	// case "home":
	// 	// прилетает сюда из вида home
	// 	mdl := *rdc.model
	// 	// создаем еффект заменяем модель на текующую
	// 	msg.Model = &mdl
	// 	// это меняет только в msg работает :)
	// 	// msg.Model.Gui.MainWindow.Title = "test change"
	// 	rdc.effects.ChanIn() <- msg
	// 	// из эффекта возвращается в "page"
	// case "znak":
	// 	// прилетает сюда из вида znak
	// 	mdl := *rdc.model
	// 	msg.Model = &mdl
	// 	rdc.effects.ChanIn() <- msg
	// 	// из эффекта возвращается в "page"
	case "page":
		mdl := *msg.Model
		_, newModel := rdc.UpdaterGUI("page", mdl)
		rdc.model = &newModel
	case "reload":
		mdl := *msg.Model
		_, newModel := rdc.UpdaterGUI("reload", mdl)
		rdc.model = &newModel
	// case "first":
	// 	// прилетает сюда когда впервые выставляется меню
	// 	// этакая общая инициализация
	// 	mdl := *rdc.model
	// 	msg.Model = &mdl
	// 	rdc.effects.ChanIn() <- msg
	// потом улетает в all
	// case "all":
	// 	mdl := *msg.Model
	// 	_, newModel := rdc.UpdaterGUI("all", mdl)
	// 	rdc.model = &newModel
	// 	msg := domain.Message{
	// 		Sender: "reductor.all",
	// 		Cmd:    "znak",
	// 		Model:  rdc.model,
	// 	}
	// 	rdc.effects.ChanIn() <- msg
	case "all":
		mdl := *msg.Model
		_, _ = rdc.UpdaterGUI("all", mdl)
		rdc.model.Gui.MainWindow.StatusBar.Fsrarid = msg.Model.Gui.MainWindow.StatusBar.Fsrarid
		rdc.model.Gui.MainWindow.StatusBar.Utm = msg.Model.Gui.MainWindow.StatusBar.Utm
		rdc.model.Gui.MainWindow.StatusBar.License = msg.Model.Gui.MainWindow.StatusBar.License
	case "statusbar_fsrarid":
		mdl := *msg.Model
		_, _ = rdc.UpdaterGUI("statusbar", mdl)
		// только статус обновляем поля модели редуктора
		rdc.model.Gui.MainWindow.StatusBar.Fsrarid = msg.Model.Gui.MainWindow.StatusBar.Fsrarid
		rdc.model.Gui.MainWindow.StatusBar.Utm = msg.Model.Gui.MainWindow.StatusBar.Utm
	case "statusbar_license":
		mdl := *msg.Model
		_, _ = rdc.UpdaterGUI("statusbar", mdl)
		// только статус обновляем поля модели редуктора
		rdc.model.Gui.MainWindow.StatusBar.License = msg.Model.Gui.MainWindow.StatusBar.License
	case "statusbar_scan":
		mdl := *msg.Model
		_, _ = rdc.UpdaterGUI("statusbar", mdl)
		// только статус обновляем поля модели редуктора
		rdc.model.Gui.MainWindow.StatusBar.Scan = msg.Model.Gui.MainWindow.StatusBar.Scan
	case "statusbar_utm":
		mdl := *msg.Model
		_, _ = rdc.UpdaterGUI("statusbar", mdl)
		// только статус обновляем поля модели редуктора
		rdc.model.Gui.MainWindow.StatusBar.Utm = msg.Model.Gui.MainWindow.StatusBar.Utm
		rdc.model.Gui.MainWindow.StatusBar.Fsrarid = msg.Model.Gui.MainWindow.StatusBar.Fsrarid
	}
}
