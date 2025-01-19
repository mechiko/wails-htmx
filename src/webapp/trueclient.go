package webapp

import (
	"firstwails/domain"
	"firstwails/trueclient"
	"net/url"
	"time"
)

// перед созданием труеклиента
// всегда берем данные из config.db потому что они там обновляются в любое время
// и могут поменяться
func (a *webapp) StartTrueClient(model domain.Model) domain.ITrueClient {
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
	} else {
		// если нет config.db берем из конфига программы
		if a.Configuration().TrueClient.HashKey != "" {
			hashKey = a.Configuration().TrueClient.HashKey
		}
		if a.Configuration().TrueClient.DeviceID != "" {
			model.TrueClient.DeviceID = a.Configuration().TrueClient.DeviceID
		}
		if a.Configuration().TrueClient.OmsID != "" {
			model.TrueClient.OmsID = a.Configuration().TrueClient.OmsID
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
	a.Logger().Debugf("%s trueclient authorised", modError)
	// после авторизации обновляем модель
	model.TrueClient.TokenGIS = tc.TokenGIS()
	model.TrueClient.TokenSUZ = tc.TokenSUZ()
	model.TrueClient.AuthTime = tc.AuthTime()
	a.UpdateModel(model)
	return tc
}

// делаем пинг суз
func (a *webapp) StartTrueClientSuz(model domain.Model) domain.ITrueClient {
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
	} else {
		// если нет config.db берем из конфига программы
		if a.Configuration().TrueClient.HashKey != "" {
			hashKey = a.Configuration().TrueClient.HashKey
		}
		if a.Configuration().TrueClient.DeviceID != "" {
			model.TrueClient.DeviceID = a.Configuration().TrueClient.DeviceID
		}
		if a.Configuration().TrueClient.OmsID != "" {
			model.TrueClient.OmsID = a.Configuration().TrueClient.OmsID
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
	} else {
		// если нет ошибок сделаем пинг суз
		if info, err := tc.PingSuz(); err != nil {
			model.Error = append(model.Error, err.Error())
		} else {
			model.TrueClient.PingSuz = info
		}
	}
	a.Logger().Debugf("%s trueclient authorised", modError)
	// после авторизации обновляем модель
	model.TrueClient.TokenGIS = tc.TokenGIS()
	model.TrueClient.TokenSUZ = tc.TokenSUZ()
	model.TrueClient.AuthTime = tc.AuthTime()
	a.UpdateModel(model)
	return tc
}
