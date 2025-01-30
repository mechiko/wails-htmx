package usecase

import (
	"firstwails/domain"
	"fmt"
)

// var reentranceStatsModelFlag int64

// вызывается в Render() страницы stats
func (u *usecase) StatsModel(model domain.Model) domain.Model {
	defer func() {
		if r := recover(); r != nil {
			u.Logger().Error(fmt.Errorf("%s panic %v", modError, r))
		}
	}()
	// if atomic.CompareAndSwapInt64(&reentranceStatsModelFlag, 0, 1) {
	// 	defer atomic.StoreInt64(&reentranceStatsModelFlag, 0)
	// } else {
	// 	u.Logger().Errorf("%s reenter StatsModel()", modError)
	// 	return
	// }
	model.Stats.IsTrueZnakA3 = u.Repo().IsZnak()
	return model
}
