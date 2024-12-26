package usecase

import "firstwails/domain"

func (u *usecase) HomeModel(model domain.Model) domain.Model {
	model.Home.DbA3 = u.Repo().Dbs().A3()
	model.Home.DbLite = u.Repo().Dbs().Lite()
	model.Home.DbConfig = u.Repo().Dbs().Config()
	model.Home.DbZnak = u.Repo().Dbs().Znak()
	return model
}
