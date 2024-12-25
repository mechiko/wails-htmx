package a3sqlite

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

//go:embed sql/history_partner.sql
var historyPartnerSql string

func (a *dbAlcohelp3) HistoryPartner(start, end, fsrarId string) (out domain.HistoryItemSlice, err error) {
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
		Fsrar string
	}{
		Start: start,
		End:   end,
		Fsrar: fsrarId,
	}
	templateQuery := historyPartnerSql
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
