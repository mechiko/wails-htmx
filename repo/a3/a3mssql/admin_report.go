package a3mssql

import (
	"context"
	"fmt"

	"firstwails/domain"

	"github.com/volatiletech/sqlboiler/v4/queries"
)

func (a *dbAlcohelp3) AdminReport() (out *domain.AdminReport, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()

	ctx := context.Background()
	out = &domain.AdminReport{}
	forms := make(domain.FormDoubleSlice, 0)
	sql := `SELECT doc_reg_id as id, COUNT(*) as total FROM production_form1 GROUP BY id_production_reports, doc_reg_id HAVING count(*) > 1;`
	if err := queries.Raw(sql).Bind(ctx, a.db, &forms); err != nil {
		return out, fmt.Errorf("%s %w", modError, err)
	}
	if len(forms) > 0 {
		out.IsDoubleFormProduce = true
		out.IsDoubleFormProduceRow = forms
	}

	forms = nil
	sql = `select doc_reg_id as id, COUNT(*) as total from import_form1 if2 group by id_import_reports, doc_reg_id HAVING count(*) > 1;`
	if err := queries.Raw(sql).Bind(ctx, a.db, &forms); err != nil {
		return out, fmt.Errorf("%s %w", modError, err)
	}
	if len(forms) > 0 {
		out.IsDoubleFormImport = true
		out.IsDoubleFormImportRow = forms
	}

	forms = nil
	sql = `select doc_reg_id as id, COUNT(*) as total from ttn_form2 tf GROUP BY id_ttn, doc_reg_id  HAVING count(*) > 1;`
	if err := queries.Raw(sql).Bind(ctx, a.db, &forms); err != nil {
		return out, fmt.Errorf("%s %w", modError, err)
	}
	if len(forms) > 0 {
		out.IsDoubleFormTtn = true
		out.IsDoubleFormTtnRow = forms
	}

	forms = nil
	sql = `select product_inform_f1_reg_id as id, COUNT(*) as total from form1_egais GROUP BY product_inform_f1_reg_id  HAVING count(*) > 1;`
	if err := queries.Raw(sql).Bind(ctx, a.db, &forms); err != nil {
		return out, fmt.Errorf("%s %w", modError, err)
	}
	if len(forms) > 0 {
		out.IsDoubleForm1 = true
		out.IsDoubleForm1Row = forms
	}

	forms = nil
	sql = `select product_inform_f2_reg_id as id, COUNT(*) as total from form2_egais GROUP BY product_inform_f2_reg_id  HAVING count(*) > 1;`
	if err := queries.Raw(sql).Bind(ctx, a.db, &forms); err != nil {
		return out, fmt.Errorf("%s %w", modError, err)
	}
	if len(forms) > 0 {
		out.IsDoubleForm2 = true
		out.IsDoubleForm2Row = forms
	}

	absent := make(domain.AbsentItemSlice, 0)
	sql = `select DISTINCT tp.product_inform_f1_reg_id as 'id'
  FROM ttn tt join ttn_products tp on tp.id_ttn = tt.id
  where 
  tt.doc_date >= '2024.10.01'
  and tp.product_inform_f1_reg_id not in (select product_inform_f1_reg_id from form1_egais)
  ;`
	if err := queries.Raw(sql).Bind(ctx, a.db, &absent); err != nil {
		return out, fmt.Errorf("%s %w", modError, err)
	}
	out.AbsentForm1 = append(out.AbsentForm1, absent...)

	absent = nil
	sql = `select DISTINCT tp.product_inform_f2_reg_id as 'id'
  FROM ttn tt join ttn_products tp on tp.id_ttn = tt.id
  where tt.ttn_type in ('Исходящий', 'Импорт')
  and tt.doc_date >= '2024.10.01'
  and tp.product_inform_f2_reg_id not in (select product_inform_f2_reg_id from form2_egais)
  ;`
	if err := queries.Raw(sql).Bind(ctx, a.db, &absent); err != nil {
		return out, fmt.Errorf("%s %w", modError, err)
	}
	out.AbsentForm2 = append(out.AbsentForm2, absent...)

	absent = nil
	sql = `select DISTINCT tt.consignee_client_reg_id as 'id'
  FROM ttn tt where tt.doc_date >= '2024.10.01' and tt.consignee_client_reg_id not in (select client_reg_id from partners_egais)
  ;`
	if err := queries.Raw(sql).Bind(ctx, a.db, &absent); err != nil {
		return out, fmt.Errorf("%s %w", modError, err)
	}
	out.AbsentClient = append(out.AbsentClient, absent...)
	return out, nil
}
