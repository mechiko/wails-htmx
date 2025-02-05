package znaksqlite

import (
	"context"
	"firstwails/repo/znak/znaksqlite/znakboil"
	"firstwails/utility"
	"fmt"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (a *dbZnak) AttachLite(liteDbName string, status string, cis string) (out int64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()
	ctx := context.Background()
	if liteDbName == "" {
		return out, fmt.Errorf("%s укажите имя бд lite", modError)
	}
	if !utility.PathOrFileExists(liteDbName) {
		return out, fmt.Errorf("%s бд lite не найдена", modError)
	}
	likeCis := fmt.Sprintf("%s%%", cis)
	cisOrder, err := znakboil.OrderMarkCodesSerialNumbers(qm.Where("code like ?", likeCis)).One(ctx, a.db)
	if err != nil {
		return out, fmt.Errorf("%s %w", modError, err)
	}
	a.Logger().Debugf("cisOrder id=%d", cisOrder.IDOrderMarkCodes.Int64)
	markOrder, err := znakboil.FindOrderMarkCode(ctx, a.db, cisOrder.IDOrderMarkCodes.Int64)
	if err != nil {
		return out, fmt.Errorf("%s %w", modError, err)
	}
	a.Logger().Debugf("markOrder id=%d", markOrder.ID)
	dateNow := time.Now().Local().Format(a.Configuration().Layouts.TimeLayoutClear)
	markOrder.CreateDate = null.StringFrom(dateNow)
	productName := fmt.Sprintf("Дубликат для печати %s", markOrder.ProductName.String)
	markOrder.ProductName = null.StringFrom(productName)
	markOrder.ID = 0

	attachSql := "ATTACH DATABASE ? AS litedb;"
	_, err = a.db.Exec(attachSql, liteDbName)
	if err != nil {
		return out, fmt.Errorf("%s attach db %w", modError, err)
	}
	var countCis int
	sql := `select count(*) from litedb.cis_request where status = ?;`
	err = a.db.Get(&countCis, sql, status)
	if err != nil {
		return out, fmt.Errorf("%s select %w", modError, err)
	}
	markOrder.Quantity = null.Int64From(int64(countCis))
	markOrder.SerialNumberType = null.StringFrom("Оператор")
	if err := markOrder.Insert(ctx, a.db, boil.Infer()); err != nil {
		return out, fmt.Errorf("%s db %w", modError, err)
	}
	out = markOrder.ID
	lenCis := len(cis)
	sql = `INSERT INTO order_mark_codes_serial_numbers (id_order_mark_codes, gtin, serial_number, code, block_id, status) 
  SELECT ?, gtin, serial_number, code, block_id, status  from order_mark_codes_serial_numbers omcsn where 
  SUBSTR(omcsn.code,1,?) in (select cis from litedb.cis_request where status = ?);`
	_, err = a.db.Exec(sql, markOrder.ID, lenCis, status)
	if err != nil {
		return out, fmt.Errorf("%s attach db %w", modError, err)
	}
	return out, nil
}
