package usecase

import (
	"firstwails/domain"
	"firstwails/trueclient/mystore"
	"fmt"
	"sync/atomic"
)

var reentranceSetupModelFlag int64

func (u *usecase) SetupModel(model domain.Model) domain.Model {
	defer func() {
		if r := recover(); r != nil {
			u.Logger().Error(fmt.Errorf("%s panic %v", modError, r))
		}
	}()
	go func() {
		if atomic.CompareAndSwapInt64(&reentranceSetupModelFlag, 0, 1) {
			defer atomic.StoreInt64(&reentranceSetupModelFlag, 0)
		} else {
			u.Logger().Errorf("%s reenter SetupModel()", modError)
			return
		}
		// эта часть исполняется только в одиночку,
		// повторный вызов не будет выполнять если запущена уже
		if storeMy, err := mystore.List(u.Logger()); err == nil {
			model.TrueClient.MyStore = storeMy
		} else {
			u.Logger().Errorf("%s %s", modError, err.Error())
		}
		msg := domain.Message{}
		msg.Cmd = "setup"
		msg.Sender = "usecase.SetupModel"
		msg.Model = &model
		u.Reductor().ChanIn() <- msg
	}()

	return model
}
