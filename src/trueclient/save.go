package trueclient

import "firstwails/domain"

// сохраняем в конфиге и в модели токены и время получения
func (t *trueClient) Save(model *domain.TrueClient) {
	model.AuthTime = t.authTime
	model.TokenGIS = t.tokenGis
	model.TokenSUZ = t.tokenSuz
}
