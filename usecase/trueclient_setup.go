package usecase

import (
	"firstwails/domain"
	"firstwails/trueclient"
	"fmt"
)

// не реентерабельная потому что не защищена
// вызывает авторизацию и затем отправляет модель в редуктор при его создании
// для страницы Setup
func (u *usecase) TrueClientSetup(model domain.TrueClient) (out domain.TrueClient) {
	defer func() {
		if r := recover(); r != nil {
			errStr := fmt.Sprintf("%s panic %v", modError, r)
			u.Logger().Error(errStr)
			model.Errors = append(out.Errors, errStr)
			out = model
		}
	}()
	// создаем клиента и если надо авторизуемся
	// там же сохраняется в конфиг если происходит обновление токенов
	tc := trueclient.New(u, model)
	if len(tc.Errors()) > 0 {
		model.Errors = append(model.Errors, tc.Errors()...)
		u.Logger().Debugf("%s trueclient authorised errors %d", modError, len(tc.Errors()))
	}
	u.Logger().Debugf("%s trueclient authorised", modError)
	// после авторизации обновляем модель
	model.TokenGIS = tc.TokenGIS()
	model.TokenSUZ = tc.TokenSUZ()
	model.AuthTime = tc.AuthTime()
	model.PingSuz = tc.PingSuzInfo()
	_ = u.Config().Set("trueclient.deviceid", model.DeviceID, true)
	_ = u.Config().Set("trueclient.omsid", model.OmsID, true)
	_ = u.Config().Set("trueclient.hashkey", model.HashKey, true)

	newModel := u.Reductor().Model()
	newModel.TrueClient = model
	msg := domain.Message{}
	msg.Cmd = "trueclient"
	msg.Sender = "usecase.TrueClientSetup"
	msg.Model = &newModel
	u.Reductor().ChanIn() <- msg
	return model
}
