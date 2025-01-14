package usecase

import (
	"firstwails/domain"
	"firstwails/trueclient"
	"firstwails/utility"
	"fmt"
	"time"
)

// не реентерабельная потому что не защищена
// вызывает авторизацию и затем отправляет модель в редуктор при его создании
// для страницы Setup router() Save()
func (u *usecase) TrueClientSearch(model domain.Model) (out domain.Model) {
	defer func() {
		if r := recover(); r != nil {
			errStr := fmt.Sprintf("%s panic %v", modError, r)
			u.Logger().Error(errStr)
			out.Stats.Errors = append(out.Stats.Errors, errStr)
			out = model
		}
	}()
	// создаем клиента и если надо авторизуемся
	// там же сохраняется в конфиг если происходит обновление токенов
	tc := trueclient.New(u, model.TrueClient)
	if len(tc.Errors()) > 0 {
		model.Stats.Errors = append(model.Stats.Errors, tc.Errors()...)
		u.Logger().Debugf("%s trueclient authorised errors %d", modError, len(tc.Errors()))
		return model
	}
	u.Logger().Debugf("%s trueclient authorised", modError)
	// после авторизации обновляем модель

	chunkSize := 1000
	allCises := utility.ReadTextStringArray(model.Stats.File)
	chunks := utility.SplitStringSlice2Chunks(allCises, chunkSize)
	u.Logger().Debugf("%s len allCises %d chunks %d", modError, len(allCises), len(chunks))
	if err := u.Repo().DbLite().CisRequestDeleteAll(); err != nil {
		model.Stats.Errors = append(model.Stats.Errors, tc.Errors()...)
		u.Logger().Debugf("%s trueclient authorised errors %d", modError, len(tc.Errors()))
		return model
	}
	model.Stats.CisStatus = make(map[string]int)
	model.Stats.CisList = make(domain.CisSlice, 0)
	start := time.Now()
	for _, chunk := range chunks {
		cisResponce := []domain.CisesPost{}
		if err := tc.CisesListPost(&cisResponce, chunk); err != nil {
			model.Stats.Errors = append(model.Stats.Errors, err.Error())
			u.Logger().Debugf("%s tc.CisesList %s", modError, err.Error())
			return model
		}
		if err := u.Repo().DbLite().InsertCisRequestPost(cisResponce); err != nil {
			model.Stats.Errors = append(model.Stats.Errors, tc.Errors()...)
			u.Logger().Debugf("%s InsertCisRequest %s", modError, err.Error())
			return model
		}
		for _, cisItem := range cisResponce {
			if count, ok := model.Stats.CisStatus[cisItem.Result.Status]; !ok {
				model.Stats.CisStatus[cisItem.Result.Status] = 1
			} else {
				model.Stats.CisStatus[cisItem.Result.Status] = count + 1
			}
			model.Stats.CisList = append(model.Stats.CisList, &domain.Cis{Cis: cisItem.Result.Cis, Status: cisItem.Result.Status})
		}
	}
	u.Logger().Debugf("%s CisRequest time since %v", modError, time.Since(start))
	u.Logger().Debugf("%s CisList %d", modError, len(model.Stats.CisList))
	// обновляем страницу статс
	// msg := domain.Message{}
	// msg.Cmd = "stats"
	// msg.Sender = "usecase.TrueClientSearch"
	// msg.Model = &model
	// u.Reductor().ChanIn() <- msg
	return model
}
