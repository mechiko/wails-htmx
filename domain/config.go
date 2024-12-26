package domain

type IConfig interface {
	Get() interface{}
	Save() error
	SaveAs(fn string) error
	SaveSafe() error
	GetByName(name string) interface{}
	Set(key string, value interface{}, save ...bool) error
	Unmarshal(*Configuration) error
}
