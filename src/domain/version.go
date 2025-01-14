package domain

// эти устанавливать лучше в батнике при компиляции возможно
const ExeVersion string = "0.0.1"

// версия БД хранится в пакете repo/version/create
// она зашита в исходном коде пакета и меняется по мере изменения структуры БД

type VersionDB interface {
	Upgrade() error
	Version() string
	WhatsNew() string
}
