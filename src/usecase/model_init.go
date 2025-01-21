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
	// model = u.TrueClientConfig(model)
	// if len(model.Error) != 0 {
	// 	u.SetActivePage("setup")
	// }
	// меняем настройку использовать config.db если эта бд не доступна
	if u.Configuration().ConfigDB {
		if !u.Repo().IsConfig() {
			_ = u.Config().Set("configdb", false, true)
		}
	}
	trueModel := u.InitTrueClient(model.TrueClient)
	model.TrueClient = trueModel
	return model
}
