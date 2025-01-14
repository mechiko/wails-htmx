package domain

// элемент движения алкоголя
type HistoryItem struct {
	ID             int64
	Type           string `db:"type"`       // тип записи ттн производство импорт акт
	DocType        string `db:"doc_type"`   // Расход, Приход или Возврат от покупателя, Возврат Покупателю
	DocId          int    `db:"doc_id"`     // для ссылки на источник документ ттн акт отчет импорта производства
	DocContentId   int    `db:"content_id"` // для ссылки на источник строка документа акта отчета ттн
	DocNumber      string `db:"doc_number"`
	DocDate        string `db:"doc_date"`
	DocReason      string `db:"doc_reason"` // для актов причина списания, для ТТН расхождение если есть (отгрузка с расхождением получение с расхождением)
	DocSource      string `db:"doc_source"`
	DocRegId       string `db:"doc_reg_id"`
	PartnerFsrarId string `db:"partner_fsrar_id"` // фсрар ид партнера потом берем по справочнику
	FixNumber      string `db:"fix_number"`       // номер и дата фиксации опционально может потом потребуется
	FixDate        string `db:"fix_date"`         // номер и дата фиксации опционально может потом потребуется
	FixDateASIIU   string `db:"fix_date_asiiu"`   // номер и дата фиксации опционально может потом потребуется
	// описание товара
	UnitType        string  `db:"product_unit_type"` // тип упаковки
	Quantity        float64 `db:"quantity"`          // количество по документу
	QuantityFact    float64 `db:"quantity_fact"`     // количество фактическое если был акт расхождения
	FullName        string  `db:"full_name"`         // полное имя продукции
	AlcVolume       float64 `db:"alc_volume"`        // градусы
	AlcVolumeFact   float64 `db:"alc_volume_fact"`   // фактический градус по производству по справке А
	Code            string  `db:"code"`              // вид АП
	AlcCode         string  `db:"alc_code"`          // код АП
	Capacity        float64 `db:"capacity"`          // емкость тары для Unpacked старой - 10 литров (1 дал) для новой емкость кеги в литрах
	ProducerFsrarId string  `db:"producer_fsrar_id"` // фсрар ид производителя потом берем по справочнику
	Form1           string  `db:"form1"`             // номер справки 1
	Form2           string  `db:"form2"`             // номер справки 2 по владельцу утм
	BottlingDate    string  `db:"bottling_date"`     // дара розлива
	Status          string  `db:"status"`
	// количество по строке товара
	Counts int64   `db:"counts"` // Packed штуки
	Dal    float64 `db:"dal"`    // Unpacked далы
	Vol    float64 `db:"volume"` // объем далы
	BVol   float64 `db:"bvol"`   // безводный спирт * градусы / 100
	Summ   float64 `db:"summ"`   // сумма по строке Quanity * Price
}

type HistoryItemSlice []*HistoryItem
