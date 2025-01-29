package domain

type DbZnak interface {
	GtinAll() (out []string, err error)
}
