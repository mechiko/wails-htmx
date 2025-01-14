package repo

func (r *repository) IsA3() bool {
	return !r.dbs.A3().Absent()
}

func (r *repository) IsConfig() bool {
	return !r.dbs.Config().Absent()
}

func (r *repository) IsZnak() bool {
	return !r.dbs.Znak().Absent()
}
