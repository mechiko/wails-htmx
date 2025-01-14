package usecase

import (
	"firstwails/domain"
	"fmt"
	"sync/atomic"
)

var reentranceStatsModelFlag int64

func (u *usecase) StatsModel(model domain.Model) domain.Model {
	defer func() {
		if r := recover(); r != nil {
			u.Logger().Error(fmt.Errorf("%s panic %v", modError, r))
		}
	}()
	go func() {
		if atomic.CompareAndSwapInt64(&reentranceStatsModelFlag, 0, 1) {
			defer atomic.StoreInt64(&reentranceStatsModelFlag, 0)
		} else {
			u.Logger().Errorf("%s reenter StatsModel()", modError)
			return
		}
		// эта часть исполняется только в одиночку,
		// повторный вызов не будет выполнять если запущена уже
		model.Stats.Title = "Inits Title"
		msg := domain.Message{}
		msg.Cmd = "stats"
		msg.Sender = "usecase.StatsModel"
		msg.Model = &model
		u.Reductor().ChanIn() <- msg
	}()

	return model
}
