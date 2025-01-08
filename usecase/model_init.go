package usecase

import "firstwails/domain"

func (u *usecase) InitModel(model domain.Model) domain.Model {
	model.DbInfo.UtmHost = u.Configuration().UtmHost
	model.DbInfo.UtmPort = u.Configuration().UtmPort
	model.DbInfo.Output = u.Configuration().Output
	model.DbInfo.Export = u.Configuration().Export
	model.Gui.MainWindow.StatusBar.Utm = false
	model.Gui.MainWindow.StatusBar.Fsrarid = u.Configuration().Application.Fsrarid
	return model
}
