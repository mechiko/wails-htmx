package usecase

import (
	"firstwails/domain"
	"net/url"
	"time"
)

// не реентерабельная потому что не защищена
// вызывает авторизацию и затем отправляет модель в редуктор при его создании
// для страницы Setup router() Save()
func (u *usecase) InitTrueClient(model domain.TrueClient) (out domain.TrueClient) {
	var layout = "2006-01-02 15:04:05.000000"
	// model := u.reductor.Model()
	// проверяем базу config.db
	tokenGIS := ""
	tokenSUZ := ""
	authTime := time.Time{}
	hashKey := ""
	deviceId := ""
	omsId := ""
	useConfiDB := false
	//  берем из конфига программы дальше поменяется если настроено
	if u.Configuration().TrueClient.HashKey != "" {
		hashKey = u.Configuration().TrueClient.HashKey
	}
	if u.Configuration().TrueClient.DeviceID != "" {
		deviceId = u.Configuration().TrueClient.DeviceID
	}
	if u.Configuration().TrueClient.OmsID != "" {
		omsId = u.Configuration().TrueClient.OmsID
	}
	if u.Configuration().ConfigDB {
		useConfiDB = true
		if u.Repo().IsConfig() {
			// есть база конфиг.дб
			// берем токены и время из нее
			if token, err := u.Repo().DbConfig().Key("token_suz"); err != nil {
				u.Logger().Errorf("%s %s", modError, err.Error())
			} else {
				tokenSUZ = token
			}
			if token, err := u.Repo().DbConfig().Key("token_gis_mt"); err != nil {
				u.Logger().Errorf("%s %s", modError, err.Error())
			} else {
				tokenGIS = token
			}
			if timeLast, err := u.Repo().DbConfig().Key("token_suz_expiration_date"); err != nil {
				u.Logger().Errorf("%s %s", modError, err.Error())
			} else {
				if auTime, err := time.Parse(layout, timeLast); err != nil {
					u.Logger().Errorf("%s %s", modError, err.Error())
				} else {
					authTime = auTime.Add(-10 * time.Hour)
				}
			}
			if hash, err := u.Repo().DbConfig().Key("certificate_thumbprint"); err != nil {
				u.Logger().Errorf("%s %s", modError, err.Error())
			} else {
				hashKey = hash
			}
			if key, err := u.Repo().DbConfig().Key("oms_id"); err != nil {
				u.Logger().Errorf("%s %s", modError, err.Error())
			} else {
				omsId = key
			}
			if key, err := u.Repo().DbConfig().Key("connection_id"); err != nil {
				u.Logger().Errorf("%s %s", modError, err.Error())
			} else {
				deviceId = key
			}
		}
	}
	// true client stand urls
	if u.Configuration().TrueClient.StandGIS == "" {
		_ = u.Config().Set("trueclient.standgis", "markirovka.sandbox.crptech.ru", true)
	}
	if u.Configuration().TrueClient.StandSUZ == "" {
		_ = u.Config().Set("trueclient.standsuz", "suz.sandbox.crptech.ru", true)
	}
	model.StandGIS = url.URL{
		Scheme: "https",
		Host:   u.Configuration().TrueClient.StandGIS,
	}
	model.StandSUZ = url.URL{
		Scheme: "https",
		Host:   u.Configuration().TrueClient.StandSUZ,
	}
	model.TokenGIS = tokenGIS
	model.TokenSUZ = tokenSUZ
	model.LayoutUTC = u.Configuration().TrueClient.LayoutUTC
	model.AuthTime = authTime
	model.HashKey = hashKey
	model.DeviceID = deviceId
	model.OmsID = omsId
	model.UseConfigDB = useConfiDB
	// сохраняем в конфиге идентификаторы и ключ
	_ = u.Config().Set("trueclient.deviceid", model.DeviceID, true)
	_ = u.Config().Set("trueclient.omsid", model.OmsID, true)
	_ = u.Config().Set("trueclient.hashkey", model.HashKey, true)
	return model
}
