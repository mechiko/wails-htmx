package usecase

import "firstwails/domain"

func (u *usecase) StatsModel(model domain.Model) domain.Model {
	model.Stats.Title = "Inits Title"
	return model
}
