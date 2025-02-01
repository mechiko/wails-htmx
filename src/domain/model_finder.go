package domain

// State
// 0 начальное
// 1 завершена обработка
type finder struct {
	State         int
	Title         string
	File          string
	FormInput     string
	Progress      int      // прогресс опроса
	Errors        []string // массив ошибок
	CisFindInfoIn CisFindInfoSlice
}

type CisFindInfo struct {
	Cis     string
	CisSrc  string
	Code    string
	CodeFNS string
	Name    string
	Root    string
	Dir     string
	Order   int64
}

type CisFindInfoSlice []*CisFindInfo
