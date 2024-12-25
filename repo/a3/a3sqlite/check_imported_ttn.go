package a3sqlite

import (
	"fmt"

	"firstwails/entity"
)

const sqlImportedCheck = `SELECT 
	t.doc_number as 'number',
	t.doc_date as 'date',
	(SELECT count(*) FROM ttn_products tp WHERE tp.id_ttn = t.id) as 'positions',
	(SELECT sum(cast(tp.product_quantity as float) * cast(tp.product_price as float)) FROM ttn_products tp WHERE tp.id_ttn = t.id) as 'summa',
	(SELECT sum(cast(tp.product_quantity as float)) FROM ttn_products tp WHERE tp.id_ttn = t.id AND tp.product_unit_type = 'Unpacked') AS 'unpacked',
	(SELECT sum(cast(tp.product_quantity as float)) FROM ttn_products tp WHERE tp.id_ttn = t.id AND tp.product_unit_type <> 'Unpacked') AS 'packages',
	(SELECT sum(iif(tp.product_unit_type='Unpacked', tp.product_quantity, cast(tp.product_capacity as float) * cast(tp.product_quantity as float))) FROM ttn_products tp WHERE tp.id_ttn = t.id) as 'volume'
FROM ttn t
where t.doc_date = ?
;`

// ищем любой тип с таким же номером и датой со state <> 'Создан'
// если такие есть то возвращается ID
// ошибка если только реально ошибка обращения к БД
func (a *dbAlcohelp3) CheckImportedTtn(date string) (result entity.ImportedTtnInfoSlice, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic FindImportTtn %v", modError, r)
		}
		a.сloseDB()
	}()
	result = make(entity.ImportedTtnInfoSlice, 0)
	rows, err := a.db.Query(sqlImportedCheck, date)
	if err != nil {
		return result, fmt.Errorf("%s %w", modError, err)
	}
	for rows.Next() {
		newInfo := entity.ImportedTtnInfo{}
		if err := rows.Scan(&newInfo.Number, &newInfo.DocDate, &newInfo.Positions, &newInfo.Summa, &newInfo.UnPackages, &newInfo.Packages, &newInfo.Volume); err != nil {
			return result, fmt.Errorf("%s %w", modError, err)
		}
		result = append(result, &newInfo)
	}
	return result, nil
}
