package usecase

import (
	"firstwails/domain"
)

// вызывается из main затем как то будет передана в редуктор для начальной модели
// из конфига и дефолтных значения
func (u *usecase) InitModel(model domain.Model) domain.Model {
	model.DbInfo.UtmHost = u.Configuration().UtmHost
	model.DbInfo.UtmPort = u.Configuration().UtmPort
	model.DbInfo.Output = u.Configuration().Output
	model.DbInfo.Export = u.Configuration().Export
	model.Gui.MainWindow.StatusBar.Utm = false
	model.Gui.MainWindow.StatusBar.Fsrarid = u.Configuration().Application.Fsrarid
	// create menu state
	model = u.MenuModel(model)
	return model
}
