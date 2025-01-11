package usecase

import (
	"firstwails/domain"
	"firstwails/trueclient"
	"fmt"
	"net/url"
	"time"
)

// не реентерабельная потому что не защищена
// вызывает авторизацию и затем отправляет модель в редуктор при его создании
// в конструкторе NewWebApp()
func (u *usecase) TrueClientConfig(model domain.Model) (out domain.Model) {
	defer func() {
		if r := recover(); r != nil {
			u.Logger().Error(fmt.Errorf("%s panic %v", modError, r))
			out = model
		}
	}()
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
		if authTime, err := time.Parse(model.TrueClient.LayoutUTC, u.Configuration().TrueClient.AuthTime); err != nil {
			model.TrueClient.AuthTime = time.Time{}
		} else {
			model.TrueClient.AuthTime = authTime
		}
	}
	model.TrueClient.HashKey = u.Configuration().TrueClient.HashKey
	model.TrueClient.DeviceID = u.Configuration().TrueClient.DeviceID
	model.TrueClient.OmsID = u.Configuration().TrueClient.OmsID
	// создаем клиента и если надо авторизуемся
	// там же сохраняется в конфиг если происходит обновление токенов
	tc := trueclient.New(u, model.TrueClient)
	if len(tc.Errors()) > 0 {
		model.Error = append(model.Error, tc.Errors()...)
		u.Logger().Debugf("%s trueclient authorised errors %d", modError, len(tc.Errors()))
	}
	u.Logger().Debugf("%s trueclient authorised", modError)
	// после авторизации обновляем модель
	model.TrueClient.TokenGIS = tc.TokenGIS()
	model.TrueClient.TokenSUZ = tc.TokenSUZ()
	model.TrueClient.AuthTime = tc.AuthTime()
	model.TrueClient.PingSuz = tc.PingSuzInfo()

	return model
}
