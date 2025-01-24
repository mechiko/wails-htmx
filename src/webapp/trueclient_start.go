package webapp

import (
	"firstwails/domain"
	"firstwails/trueclient"
	"firstwails/usecase"
	"fmt"
)

// перед созданием труеклиента
// берем данные из config.db если настроено (потому что они там обновляются в любое время)
func (a *webapp) StartTrueClient(model *domain.Model) domain.ITrueClient {
	defer func() {
		if r := recover(); r != nil {
			errStr := fmt.Sprintf("%s panic %v", modError, r)
			a.Logger().Error(errStr)
		}
	}()
	// получаем обновляем модель для trueclient
	// всегда инициализируем вдруг изменились конфиг.дб алкохелпа
	model.TrueClient = usecase.New(a).InitTrueClient(model.TrueClient)
	// создаем клиента и если надо авторизуемся
	// состояние модели обновляется автоматически
	tc := trueclient.New(a, &model.TrueClient)
	if len(tc.Errors()) > 0 {
		model.Error = append(model.Error, tc.Errors()...)
		a.Logger().Errorf("%s trueclient authorised errors %d", modError, len(tc.Errors()))
	}
	// после авторизации обновляем модель токенами и временем
	a.UpdateTrueClientModel(model.TrueClient, "webapp.startTrueClient")

	return tc
}
