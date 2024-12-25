package a3mssql

import (
	"bytes"
	"context"
	"fmt"
	"text/template"

	"firstwails/domain"
	"firstwails/zaplog"

	"github.com/volatiletech/sqlboiler/v4/queries"
)

func (a *dbAlcohelp3) HistoryPartOrigin(start, end, part string) (out domain.HistoryItemSlice, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()
	ctx := context.Background()
	out = make(domain.HistoryItemSlice, 0)
	startEnd := struct {
		Start string
		End   string
	}{
		Start: start,
		End:   end,
	}
	templateQuery := ""
	switch part {
	case "all":
		templateQuery = historyFullSql
	case "chargeon":
		templateQuery = historyChargeOnSql
	case "import":
		templateQuery = historyImportSql
	case "incoming":
		templateQuery = historyIncomingSql
	case "outgoing":
		templateQuery = historyOutgoingSql
	case "production":
		templateQuery = historyProductionSql
	case "resource":
		templateQuery = historyProductionResourceSql
	case "writeoff":
		templateQuery = historyWriteOffSql
	case "oborot":
		// templateQuery = historyOborotSql
		templateQuery = historyOutgoingSql
	default:
		return out, fmt.Errorf("%s part [%s] ошибочна (возможны all,chargeon,import,incoming,outgoing,production,resource,writeoff)", modError, part)
	}
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
	for i, hi := range out {
		if hi.AlcVolumeFact > 0 {
			out[i].BVol = hi.Vol * hi.AlcVolumeFact / 100
		} else {
			out[i].BVol = hi.Vol * hi.AlcVolume / 100
		}
	}
	return out, nil
}
