package usecase

import (
	"firstwails/domain"
)

// var reentranceStatsModelFlag int64

// вызывается каждый раз как забирается модел в Reductor
func (u *usecase) EveryModel(model domain.Model) domain.Model {
	model.Stats.IsTrueZnakA3 = u.Repo().IsZnak()
	model.IsA3 = u.Repo().IsA3()
	model.IsConfig = u.Repo().IsConfig()
	model.IsZnak = u.Repo().IsZnak()
	return model
}
