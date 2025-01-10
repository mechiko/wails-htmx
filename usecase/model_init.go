package usecase

import (
	"firstwails/domain"
	"net/url"
	"time"
)

// вызывается из эффектора startup инициализируется модель начальными данными
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
	// true client state
	if u.Configuration().TrueClient.StandGIS == "" {
		_ = u.Config().Set("trueclient.standgis", "markirovka.sandbox.crptech.ru", true)
	}
	if u.Configuration().TrueClient.StandSUZ == "" {
		_ = u.Config().Set("trueclient.standsuz", "suz.sandbox.crptech.ru", true)
	}
	model.TrueClient.StandGIS = url.URL{
		Scheme: "https",
		Host:   u.Configuration().TrueClient.StandGIS,
	}
	model.TrueClient.StandSUZ = url.URL{
		Scheme: "https",
		Host:   u.Configuration().TrueClient.StandSUZ,
	}
	model.TrueClient.TokenGIS = u.Configuration().TrueClient.TokenGIS
	model.TrueClient.TokenSUZ = u.Configuration().TrueClient.TokenSUZ
	model.TrueClient.LayoutUTC = u.Configuration().TrueClient.LayoutUTC
	if model.TrueClient.LayoutUTC == "" {
		_ = u.Config().Set("trueclient.layoututc", "2006-01-02T15:04:05", true)
		// model.TrueClient.LayoutUTC = "2006-01-02T15:04:05"
	}
	// time.Time{} -> IsZero() true
	if u.Configuration().TrueClient.AuthTime == "" {
		model.TrueClient.AuthTime = time.Time{}
	} else {
		if authTime, err := time.Parse(u.Configuration().TrueClient.AuthTime, model.TrueClient.LayoutUTC); err != nil {
			model.TrueClient.AuthTime = time.Time{}
		} else {
			model.TrueClient.AuthTime = authTime
		}
	}
	model.TrueClient.HashKey = u.Configuration().TrueClient.HashKey
	model.TrueClient.DeviceID = u.Configuration().TrueClient.DeviceID
	model.TrueClient.OmsID = u.Configuration().TrueClient.OmsID
	return model
}
