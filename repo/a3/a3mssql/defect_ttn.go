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

//go:embed sql/defect_ttn.sql
var defectTtnSql string

func (a *dbAlcohelp3) DeffectTtn(start, end string) (out domain.DefectItemSlice, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()
	ctx := context.Background()
	out = make(domain.DefectItemSlice, 0)
	startEnd := struct {
		Start string
		End   string
	}{
		Start: start,
		End:   end,
	}
	templateQuery := defectTtnSql
	sqlQuery := ""
	if tmpl, err := template.New("sqlQuery").Parse(templateQuery); err != nil {
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
