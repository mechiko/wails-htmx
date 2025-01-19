package domain

type DbConfig interface {
	Key(key string) (out string, err error)
}
