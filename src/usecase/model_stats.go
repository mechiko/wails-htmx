package usecase

import (
	"firstwails/domain"
)

// var reentranceStatsModelFlag int64

// вызывается в Render() страницы stats
func (u *usecase) StatsModel(model domain.Model) domain.Model {
	model.Stats.IsTrueZnakA3 = u.Repo().IsZnak()
	return model
}
