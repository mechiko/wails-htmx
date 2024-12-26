package domain

// рефактор
type DbLite interface {
	Version() (string, error)
}
