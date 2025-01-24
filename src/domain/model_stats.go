package domain

// State
// 0 начальное
// 1 выбран файл
// 2 запущена обработка
// 3 завершена обработка
type stats struct {
	State               int
	Title               string
	File                string
	CisIn               []string // список CIS для запроса
	Chunks              int      // куски
	CisOut              CisSlice // список CIS полученных
	CisStatus           map[string]int
	ExcelChunkSize      int      // размер куска для выгрузки в файл ексель
	IsConnectedTrueZnak bool     // есть подключение к ЧЗ
	IsTrueZnakA3        bool     // подключена БД ЧЗ А3
	OrderId             int      // номер заказа в ЧЗ А3
	UtilisationId       int      // номер отчета нанесения в ЧЗ А3
	Progress            int      // прогресс опроса
	Errors              []string // массив ошибок

}

type Cis struct {
	Cis    string
	Status string
}

type CisSlice []*Cis
