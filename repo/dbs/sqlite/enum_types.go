package sqlite

type UriType int

const (
	RwModeWithCreate UriType = iota
	RoMode
	RwMode
)

func (s UriType) String() string {
	switch s {
	case RwModeWithCreate:
		return "?cache=shared&mode=rwc"
	case RoMode:
		return "?cache=shared&mode=ro"
	case RwMode:
		return "?cache=shared&mode=rw"
	}
	return "?cache=shared&mode=rw"
}

type CheckMode int

const (
	NoMatter CheckMode = iota
	Version
	MinVersion
)

func (s CheckMode) String() string {
	switch s {
	case Version:
		return "версия проверятся c миграцией"
	case MinVersion:
		return "версия проверятся на минимальное значение"
	case NoMatter:
		return "ни каких проверок"
	}
	return "не проверятся"
}
