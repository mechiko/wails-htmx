package usecase

import "firstwails/domain"

func (u *usecase) DbInfoModel(model domain.Model) domain.Model {
	model.DbInfo.DbA3 = u.Repo().Dbs().A3()
	model.DbInfo.DbLite = u.Repo().Dbs().Lite()
	model.DbInfo.DbConfig = u.Repo().Dbs().Config()
	model.DbInfo.DbZnak = u.Repo().Dbs().Znak()
	return model
}
