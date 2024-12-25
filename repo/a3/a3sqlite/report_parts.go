package a3sqlite

import (
	"bytes"
	"context"
	"fmt"
	"text/template"

	"firstwails/domain"
	"firstwails/zaplog"

	"github.com/volatiletech/sqlboiler/v4/queries"
)

func (a *dbAlcohelp3) ReportParts(part string, start, end string) (reps domain.ReportItemA3Slice, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()

	startEnd := struct {
		Start string
		End   string
	}{
		Start: start,
		End:   end,
	}
	ctx := context.Background()
	reps = make(domain.ReportItemA3Slice, 0)
	query := ""
	switch part {
	case "act":
		query = reportActSql
	// case "chargeon":
	// 	query = historyChargeOnSql
	// case "import":
	// 	query = historyImportSql
	case "incoming":
		query = reportIncomingSql
	case "incomingreturn":
		query = reportIncomingReturnSql
	case "outgoing":
		query = reportOutgoingSql
	case "outgoingreturn":
		query = reportOutgoingReturnSql
	case "production":
		query = reportProductSql
	case "resource":
		query = reportResorceSql
	default:
		return reps, fmt.Errorf("%s part [%s] ошибочна (возможны act,incoming,incomingreturn,outgoing,outgoingreturn,production,resource,import,balance)", modError, part)
	}
	sqlQuery := ""
	if tmpl, err := template.New("sqlQuery").Parse(query); err != nil {
		return reps, fmt.Errorf("%s %w", modError, err)
	} else {
		var tpl bytes.Buffer
		if err := tmpl.Execute(&tpl, startEnd); err != nil {
			return reps, fmt.Errorf("%s %w", modError, err)
		} else {
			sqlQuery = tpl.String()
		}
	}
	zaplog.SqlDump.Info(sqlQuery)
	if err := queries.Raw(sqlQuery).Bind(ctx, a.db, &reps); err != nil {
		return reps, fmt.Errorf("%s %w", modError, err)
	} else {
		return reps, nil
	}
}
