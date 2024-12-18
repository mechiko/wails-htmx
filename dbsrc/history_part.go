package dbsrc

import (
	"bytes"
	_ "embed"
	"fmt"
	"text/template"

	"firstwails/domain"
)

//go:embed history_full.sql
var historyFullSql string

func (a *dbsrc) HistoryPartOrigin() (out domain.HistoryItemSlice, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic %v", r)
		}
		a.db.Close()
	}()
	out = make(domain.HistoryItemSlice, 0)
	startEnd := struct {
		Start string
		End   string
	}{
		Start: "2024.06.01",
		End:   "2024.06.31",
	}
	templateQuery := historyFullSql
	sqlQuery := ""
	if tmpl, err := template.New("sqlQuery").Parse(templateQuery); err != nil {
		return out, fmt.Errorf("%w", err)
	} else {
		var tpl bytes.Buffer
		if err := tmpl.Execute(&tpl, startEnd); err != nil {
			return out, fmt.Errorf("%w", err)
		} else {
			sqlQuery = tpl.String()
		}
	}
	if err := a.db.Select(&out, sqlQuery); err != nil {
		return out, fmt.Errorf("%w", err)
	}
	for i, hi := range out {
		out[i].ID = int64(i + 1)
		if hi.AlcVolumeFact > 0 {
			out[i].BVol = hi.Vol * hi.AlcVolumeFact / 100
		} else {
			out[i].BVol = hi.Vol * hi.AlcVolume / 100
		}
	}
	return out, nil
}
