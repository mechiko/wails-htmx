package a3mssql

import (
	"context"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/queries"
)

// чистим дубли
func (a *dbAlcohelp3) AdminReportRemove() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()

	ctx := context.Background()
	sql := `delete from form2_egais where 
  id in (SELECT id FROM form2_egais fe WHERE product_inform_f2_reg_id IN (select product_inform_f2_reg_id from form2_egais GROUP BY product_inform_f2_reg_id  HAVING COUNT(*) > 1)
  AND id NOT IN (select max(id) from form2_egais GROUP BY product_inform_f2_reg_id  HAVING COUNT(*) > 1))
  ;`
	if _, err := queries.Raw(sql).ExecContext(ctx, a.db); err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}

	sql = `delete from form1_egais where 
  id in (SELECT id FROM form1_egais fe WHERE product_inform_f1_reg_id IN (select product_inform_f1_reg_id from form1_egais GROUP BY product_inform_f1_reg_id  HAVING COUNT(*) > 1)
  AND id NOT IN (select max(id) from form1_egais GROUP BY product_inform_f1_reg_id  HAVING COUNT(*) > 1));`
	if _, err := queries.Raw(sql).ExecContext(ctx, a.db); err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	sql = `delete from production_form1 where 
  id in (SELECT id FROM production_form1 WHERE id_production_reports IN (select id_production_reports from production_form1 GROUP BY id_production_reports  HAVING COUNT(*) > 1)
  AND id NOT IN (select min(id) from production_form1 GROUP BY id_production_reports  HAVING COUNT(*) > 1));`
	if _, err := queries.Raw(sql).ExecContext(ctx, a.db); err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	sql = `delete from ttn_form2 where 
  id in (SELECT id FROM ttn_form2 WHERE id_ttn IN (select id_ttn from ttn_form2 GROUP BY id_ttn  HAVING COUNT(*) > 1)
  AND id NOT IN (select min(id) from ttn_form2 GROUP BY id_ttn  HAVING COUNT(*) > 1));`
	if _, err := queries.Raw(sql).ExecContext(ctx, a.db); err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}

	return nil
}
