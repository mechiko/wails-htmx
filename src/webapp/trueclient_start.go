package webapp

import (
	"firstwails/domain"
	"firstwails/trueclient"
	"fmt"
	"net/url"
	"time"
)

// перед созданием труеклиента
// берем данные из config.db если настроено (потому что они там обновляются в любое время)
func (a *webapp) StartTrueClient(model domain.Model) domain.ITrueClient {
	defer func() {
		if r := recover(); r != nil {
			errStr := fmt.Sprintf("%s panic %v", modError, r)
			a.Logger().Error(errStr)
		}
	}()
	// layout config.db
	var layout = "2006-01-02 15:04:05.000000"
	// model := a.reductor.Model()
	// проверяем базу config.db
	tokenGIS := ""
	tokenSUZ := ""
	authTime := time.Time{}
	hashKey := ""
	deviceId := ""
	omsId := ""
	//  берем из конфига программы дальше поменяется если настроено
	if a.Configuration().TrueClient.HashKey != "" {
		hashKey = a.Configuration().TrueClient.HashKey
	}
	if a.Configuration().TrueClient.DeviceID != "" {
		deviceId = a.Configuration().TrueClient.DeviceID
	}
	if a.Configuration().TrueClient.OmsID != "" {
		omsId = a.Configuration().TrueClient.OmsID
	}
	if a.Configuration().ConfigDB {
		if a.Repo().IsConfig() {
			// есть база конфиг.дб
			// берем токены и время из нее
			if token, err := a.Repo().DbConfig().Key("token_suz"); err != nil {
				a.logger.Errorf("%s %s", modError, err.Error())
			} else {
				tokenSUZ = token
			}
			if token, err := a.Repo().DbConfig().Key("token_gis_mt"); err != nil {
				a.logger.Errorf("%s %s", modError, err.Error())
			} else {
				tokenGIS = token
			}
			if timeLast, err := a.Repo().DbConfig().Key("token_suz_expiration_date"); err != nil {
				a.logger.Errorf("%s %s", modError, err.Error())
			} else {
				if auTime, err := time.Parse(layout, timeLast); err != nil {
					a.logger.Errorf("%s %s", modError, err.Error())
				} else {
					authTime = auTime
				}
			}
			if hash, err := a.Repo().DbConfig().Key("certificate_thumbprint"); err != nil {
				a.logger.Errorf("%s %s", modError, err.Error())
			} else {
				hashKey = hash
			}
			if key, err := a.Repo().DbConfig().Key("oms_id"); err != nil {
				a.logger.Errorf("%s %s", modError, err.Error())
			} else {
				omsId = key
			}
			if key, err := a.Repo().DbConfig().Key("connection_id"); err != nil {
				a.logger.Errorf("%s %s", modError, err.Error())
			} else {
				deviceId = key
			}
		}
	}
	// true client stand urls
	if a.Configuration().TrueClient.StandGIS == "" {
		_ = a.Config().Set("trueclient.standgis", "markirovka.sandbox.crptech.ru", true)
	}
	if a.Configuration().TrueClient.StandSUZ == "" {
		_ = a.Config().Set("trueclient.standsuz", "suz.sandbox.crptech.ru", true)
	}
	model.TrueClient.StandGIS = url.URL{
		Scheme: "https",
		Host:   a.Configuration().TrueClient.StandGIS,
	}
	model.TrueClient.StandSUZ = url.URL{
		Scheme: "https",
		Host:   a.Configuration().TrueClient.StandSUZ,
	}
	model.TrueClient.TokenGIS = tokenGIS
	model.TrueClient.TokenSUZ = tokenSUZ
	model.TrueClient.LayoutUTC = a.Configuration().TrueClient.LayoutUTC
	model.TrueClient.AuthTime = authTime
	model.TrueClient.HashKey = hashKey
	model.TrueClient.DeviceID = deviceId
	model.TrueClient.OmsID = omsId
	// создаем клиента и если надо авторизуемся
	// там же сохраняется в конфиг если происходит обновление токенов
	tc := trueclient.New(a, model.TrueClient)
	if len(tc.Errors()) > 0 {
		model.Error = append(model.Error, tc.Errors()...)
		a.Logger().Errorf("%s trueclient authorised errors %d", modError, len(tc.Errors()))
	}
	// после авторизации обновляем модель токенами и временем
	model.TrueClient.TokenGIS = tc.TokenGIS()
	model.TrueClient.TokenSUZ = tc.TokenSUZ()
	model.TrueClient.AuthTime = tc.AuthTime()
	a.UpdateTrueClientModel(model.TrueClient, "webapp.starttrueclient")

	// сохраняем в конфиге идентификаторы и ключ
	_ = a.Config().Set("trueclient.deviceid", model.TrueClient.DeviceID, true)
	_ = a.Config().Set("trueclient.omsid", model.TrueClient.OmsID, true)
	_ = a.Config().Set("trueclient.hashkey", model.TrueClient.HashKey, true)

	return tc
}
