package dbsrc

import (
	"bytes"
	_ "embed"
	"os"
	"text/template"

	"firstwails/domain"
)

//go:embed ttn.sql
var ttnSql string

// проглатываем
func (a *dbsrc) TTN(filter interface{}, limit, offset int) (out []domain.IDataSet) {
	defer func() {
		if r := recover(); r != nil {
			a.logger.Printf("panic db %v", r)
			out = nil
		}
	}()
	out = make([]domain.IDataSet, 0)
	out1 := make(domain.TTNOriginSlice, 0)
	filterOk, ok := filter.(*domain.TTNFilter)
	if !ok {
		a.logger.Println("error bas filter type")
		return nil
	}
	startEnd := struct {
		Limit  int
		Offset int
		Filter *domain.TTNFilter
	}{
		Limit:  limit,
		Offset: offset,
		Filter: filterOk,
	}
	templateQuery := ttnSql
	sqlQuery := ""
	if tmpl, err := template.New("sqlQuery").Parse(templateQuery); err != nil {
		a.logger.Printf("error db %s", err.Error())
		return nil
	} else {
		var tpl bytes.Buffer
		if err := tmpl.Execute(&tpl, startEnd); err != nil {
			a.logger.Printf("error db %s", err.Error())
			return nil
		} else {
			sqlQuery = tpl.String()
		}
	}
	os.WriteFile("dump.sql", []byte(sqlQuery), 0644)
	if err := a.db.Select(&out1, sqlQuery); err != nil {
		a.logger.Printf("error db %s", err.Error())
		return nil
	}
	for i := range out1 {
		out = append(out, out1[i])
	}
	return out
}
