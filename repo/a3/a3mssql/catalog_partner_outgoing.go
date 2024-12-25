package a3mssql

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"text/template"

	"firstwails/domain"
	"firstwails/zaplog"

	"github.com/volatiletech/sqlboiler/v4/queries"
)

//go:embed sql/partner_outgoing.sql
var partnerOutgoingSql string

func (a *dbAlcohelp3) CatalogPartnerOutgoing(start, end string) (out domain.PartnersOriginSlice, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()
	ctx := context.Background()
	out = make(domain.PartnersOriginSlice, 0)
	startEnd := struct {
		Start string
		End   string
	}{
		Start: start,
		End:   end,
	}
	sqlQuery := ""
	if tmpl, err := template.New("sqlQuery").Parse(partnerOutgoingSql); err != nil {
		return out, fmt.Errorf("%s %w", modError, err)
	} else {
		var tpl bytes.Buffer
		if err := tmpl.Execute(&tpl, startEnd); err != nil {
			return out, fmt.Errorf("%s %w", modError, err)
		} else {
			sqlQuery = tpl.String()
		}
	}
	zaplog.SqlDump.Info(sqlQuery)
	if err := queries.Raw(sqlQuery).Bind(ctx, a.db, &out); err != nil {
		return out, fmt.Errorf("%s %w", modError, err)
	}
	return out, nil
}

// func (a *dbAlcohelp3) CatalogPartnerOutgoing(start, end string) (list domain.PartnersOriginSlice, err error) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			err = fmt.Errorf("%s panic %v", modError, r)
// 		}
// 		a.сloseDB()
// 	}()
// 	ctx := context.Background()
// 	list = make(domain.PartnersOriginSlice, 0)
// 	where := `ID in (select distinct max(ID) from partners_egais group by client_reg_id)  AND pe.client_reg_id IN (SELECT DISTINCT tt.consignee_client_reg_id
// 	FROM ttn tt where tt.status in ('Подтверждён','Проведён') and tt.ttn_type in ('Исходящий', 'Импорт')
// 		and tt.doc_date >= ?
// 		and tt.doc_date <= ?
// 	ORDER BY 1)`
// 	if sl, err := a3boil.PartnersEgaiss(qm.Where(where, start, end)).All(ctx, a.db); err != nil {
// 		return list, fmt.Errorf("%s %w", modError, err)
// 	} else {
// 		for i := range sl {
// 			// partner := &domain.PartnersOrigin{}
// 			partner := new(domain.PartnersOrigin)
// 			partner.ConvertFromA3Mssql(sl[i])
// 			list = append(list, partner)
// 		}
// 		return list, nil
// 	}
// }
