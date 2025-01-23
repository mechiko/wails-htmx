package webapp

import (
	"firstwails/domain"
	"firstwails/trueclient"
	"firstwails/usecase"
	"fmt"
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
	// получаем обновляем модель для trueclient
	model.TrueClient = usecase.New(a).InitTrueClient(model.TrueClient)
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
