package domain

type stats struct {
	Title               string
	File                string
	IsConnectedTrueZnak bool     // есть подключение к ЧЗ
	IsTrueZnakA3        bool     // подключена БД ЧЗ А3
	OrderId             int      // номер заказа в ЧЗ А3
	UtilisationId       int      // номер отчета нанесения в ЧЗ А3
	CisList             CisSlice // список CIS для опроса
	CisStatus           map[string]int
	Progress            int      // прогресс опроса
	Errors              []string // массив ошибок
}

type Cis struct {
	Cis    string
	Status string
}

type CisSlice []*Cis
