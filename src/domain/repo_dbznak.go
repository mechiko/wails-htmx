package domain

type DbZnak interface {
	GtinAll() (out []string, err error)
	AttachLite(liteDbName string, status string, gtin string) (int64, error)
}
