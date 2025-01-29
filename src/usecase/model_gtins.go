package usecase

import (
	"firstwails/domain"
	"fmt"
	"sync/atomic"
)

var reentranceGtindModelFlag int64

func (u *usecase) GtinsModel(model domain.Model) domain.Model {
	defer func() {
		if r := recover(); r != nil {
			u.Logger().Error(fmt.Errorf("%s panic %v", modError, r))
		}
	}()
	go func() {
		if atomic.CompareAndSwapInt64(&reentranceGtindModelFlag, 0, 1) {
			defer atomic.StoreInt64(&reentranceGtindModelFlag, 0)
		} else {
			u.Logger().Errorf("%s reenter GtinsModel()", modError)
			return
		}
		// эта часть исполняется только в одиночку,
		// повторный вызов не будет выполнять если запущена уже
		if u.Repo().IsZnak() {
			if gtins, err := u.Repo().DbZnak().GtinAll(); err != nil {
				u.Logger().Errorf("%s repo dbznak %s", modError, err.Error())
			} else {
				// model.Gtins.GtinIn = utility.UniqueStringSliceElements(gtins)
				model.Gtins.A3Gtins = gtins
			}
		}
		model.Gtins.Title = "Коды GTIN"
		msg := domain.Message{}
		msg.Cmd = "gtins"
		msg.Sender = "usecase.GtinsModel"
		msg.Model = &model
		u.Reductor().ChanIn() <- msg
	}()

	return model
}
