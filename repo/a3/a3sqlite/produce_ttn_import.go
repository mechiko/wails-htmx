package a3sqlite

import (
	"context"
	"fmt"

	"firstwails/domain"
	"firstwails/repo/a3/a3sqlite/a3boil"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (a *dbAlcohelp3) ProduceTTN(ttn *domain.TTNImport, products domain.TTNProductSlice) (err error) {
	ctx := context.Background()
	tx, err := a.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = fmt.Errorf("panic %v", r)
		}
		a.сloseDB()
	}()

	ttnNew := &a3boil.TTN{}
	ttn.ConvertToTTNSqlite(ttnNew)
	// вставляем шапку ТТН
	if err = ttnNew.Insert(ctx, tx, boil.Infer()); err != nil {
		tx.Rollback()
		return fmt.Errorf("откатываем из-за ошибки %w", err)
	}
	// если есть вставляем продукты
	// при ошибке вставки откатываем всю ТТН
	for i := range products {
		product := &a3boil.TTNProduct{}
		products[i].ConvertToA3Sqlite(product)
		product.IDTTN = null.Int64From(ttnNew.ID)
		if err = product.Insert(ctx, tx, boil.Infer()); err != nil {
			tx.Rollback()
			return fmt.Errorf("откатываем из-за ошибки %w", err)
		}
	}
	tx.Commit()
	return nil
}
