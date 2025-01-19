package trueclient

import "firstwails/domain"

// сохраняем в конфиге и в модели токены и время получения
func (t *trueClient) Save(model domain.TrueClient) {
	_ = t.IApp.Config().Set("trueclient.deviceid", t.deviсeId, true)
	_ = t.IApp.Config().Set("trueclient.omsid", t.omsId, true)
	_ = t.IApp.Config().Set("trueclient.hashkey", t.hash, true)
	model.AuthTime = t.authTime
	model.TokenGIS = t.tokenGis
	model.TokenSUZ = t.tokenSuz
	t.UpdateTrueClientModel(model, "trueclient.save")
}
