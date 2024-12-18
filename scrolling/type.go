package scrolling

import (
	"io"
	"log"

	"firstwails/domain"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const defaultPageSize = 1000

// функция возвращает массив указателей на структуру
type fDataSet func(interface{}, int, int) []domain.IDataSet

type scrolling struct {
	id       string // идентификатор для уникальности формы
	template *Templates
	query    fDataSet
	filter   domain.IFilter // структура фильтра
	logger   *log.Logger
	pageSize int
	columns  *Columns
	eof      bool
}

type ScrollingBlocks struct {
	Columns *Columns
	Start   int
	Next    int
	More    bool
	Blocks  []domain.IDataSet
}

func NewScrolling(logger *log.Logger, query fDataSet, cols *Columns, filter domain.IFilter, pageSize int) *scrolling {
	sc := &scrolling{}
	sc.pageSize = defaultPageSize
	if pageSize > 0 {
		sc.pageSize = pageSize
	}
	sc.logger = logger
	sc.id = uuid.New().String()
	sc.template = NewTemplate(logger, "")
	sc.query = query
	sc.filter = filter
	sc.columns = cols
	return sc
}

// такой вариант парсит при создании
func (s *scrolling) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// шаблоны компилируются каждый раз
	return s.template.DoRender(w, name, data, c)
	// когда шаблоны компилируются при запуске
	// return s.template.Render(w, name, data, c)
}

// такой вариант парсит каждый раз при вызове
// func (s *scrolling) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
// 	return s.template.DoRender(w, name, data, c)
// }

func (s *scrolling) InitColumns(listVisible string) {
	s.columns.SetMaxWidth(5000)
	s.columns.AddColumn("Type", "Тип записи истории", "Type", true).SetFiltered(true).ClassValue("").ClassWidth("w-11")
	s.columns.AddColumn("DocType", "Тип документа", "DocType", true).SetFiltered(true).ClassValue("").ClassWidth("w-11")
	s.columns.AddColumn("DocId", "ИД док-та", "DocId", false).SetFiltered(true).ClassValue("").ClassWidth("w-10")
	s.columns.AddColumn("DocContentId", "ИД строки док-та", "DocContentId", false).SetFiltered(true).ClassValue("").ClassWidth("w-10")
	s.columns.AddColumn("DocNumber", "№ док-та", "DocNumber", true).SetFiltered(true).ClassValue("").ClassWidth("w-10")
	s.columns.AddColumn("DocDate", "Дата док-та", "DocDate", true).SetFiltered(true).ClassValue("").ClassWidth("w-10")
	s.columns.AddColumn("DocReason", "Причина", "DocReason", false).SetFiltered(true).ClassValue("").ClassWidth("w-10")
	s.columns.AddColumn("DocSource", "Источник", "DocSource", false).SetFiltered(true).ClassValue("").ClassWidth("w-12")
	s.columns.AddColumn("DocRegId", "РегИД", "DocRegId", true).SetFiltered(true).ClassValue("").ClassWidth("w-12")
	s.columns.AddColumn("PartnerFsrarId", "Партнер ФСРАРИД", "PartnerFsrarId", true).SetFiltered(true).ClassValue("").ClassWidth("w-10")
	s.columns.AddColumn("FixNumber", "Номер фиксации", "FixNumber", false).SetFiltered(true).ClassValue("").ClassWidth("w-13")
	s.columns.AddColumn("FixDate", "Дата фиксации", "FixDate", false).SetFiltered(false).ClassValue("").ClassWidth("w-10")
	s.columns.AddColumn("FixDateASIIU", "FixDateASIIU", "FixDateASIIU", false).SetFiltered(true).ClassValue("").ClassWidth("w-10")
	s.columns.AddColumn("UnitType", "Упаковка", "UnitType", true).SetFiltered(true).ClassValue("").ClassWidth("w-10")
	s.columns.AddColumn("Quantity", "Кол-во", "Quantity", true).SetFiltered(true).ClassValue("").ClassWidth("w-10")
	s.columns.AddColumn("QuantityFact", "Кол-во факт", "QuantityFact", true).SetFiltered(true).ClassValue("").ClassWidth("w-10")
	s.columns.AddColumn("FullName", "Наименование", "FullName", true).SetFiltered(true).ClassValue("").ClassWidth("w-40")
	s.columns.AddColumn("AlcVolume", "Крепость", "AlcVolume", true).SetFiltered(true).ClassValue("").ClassWidth("w-10")
	s.columns.AddColumn("AlcVolumeFact", "Крепость факт", "AlcVolumeFact", false).SetFiltered(true).ClassValue("").ClassWidth("w-10")
	s.columns.AddColumn("Code", "Вид АП", "Code", true).SetFiltered(true).ClassValue("").ClassWidth("w-10")
	s.columns.AddColumn("AlcCode", "Код АП", "AlcCode", true).SetFiltered(true).ClassValue("").ClassWidth("w-10")
	s.columns.AddColumn("Capacity", "Емкость", "Capacity", true).SetFiltered(true).ClassValue("").ClassWidth("w-10")
	s.columns.AddColumn("ProducerFsrarId", "Производитель ФСРАРИД", "ProducerFsrarId", false).SetFiltered(true).ClassValue("").ClassWidth("w-10")
	s.columns.AddColumn("Form1", "Form1", "Form1", true).SetFiltered(true).ClassValue("").ClassWidth("w-10")
	s.columns.AddColumn("Form2", "Form2", "Form2", true).SetFiltered(true).ClassValue("").ClassWidth("w-10")
	s.columns.AddColumn("BottlingDate", "Розлив", "BottlingDate", true).SetFiltered(true).ClassValue("").ClassWidth("w-10")
	s.columns.AddColumn("Status", "Статус", "Status", true).SetFiltered(true).ClassValue("").ClassWidth("w-10")
	s.columns.AddColumn("Counts", "Штуки", "Counts", true).SetFiltered(true).ClassValue("").ClassWidth("w-10")
	s.columns.AddColumn("Dal", "Далы", "Dal", true).SetFiltered(true).ClassValue("").ClassWidth("w-10")
	s.columns.AddColumn("Vol", "Объем", "Vol", true).SetFiltered(true).ClassValue("").ClassWidth("w-10")
	s.columns.AddColumn("BVol", "Безводный", "BVol", false).SetFiltered(true).ClassValue("").ClassWidth("w-10")
	s.columns.AddColumn("Summ", "Сумма", "Summ", false).SetFiltered(true).ClassValue("").ClassWidth("w-10")
}

func (s *scrolling) Columns() *Columns {
	return s.columns
}
