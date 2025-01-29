package domain

// State
// 0 начальное
// 1 завершена обработка
type gtins struct {
	State               int
	Title               string
	File                string
	GtinIn              []string // список GTIN для запроса
	A3Gtins             []string
	Chunks              int           // куски
	GtinOut             GtinInfoSlice // список GTIN полученных
	GtinStatus          map[string]int
	ExcelChunkSize      int      // размер куска для выгрузки в файл ексель
	IsConnectedTrueZnak bool     // есть подключение к ЧЗ
	IsTrueZnakA3        bool     // подключена БД ЧЗ А3
	OrderId             int      // номер заказа в ЧЗ А3
	UtilisationId       int      // номер отчета нанесения в ЧЗ А3
	Progress            int      // прогресс опроса
	Errors              []string // массив ошибок

}

type GtinInfo struct {
	Gtin string
	Info string
}

type GtinInfoSlice []*GtinInfo
