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

//go:embed sql/xml_producer_shipper_out_by_partner.sql
var xmlProducerShipperByPartnerSql string

func (a *dbAlcohelp3) XmlProducerShipperByPartner(start, end, fsrarId string) (out domain.XmlPartnerSlice, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()
	if start == "" {
		start = a.StartDateString()
	}
	if end == "" {
		end = a.EndDateString()
	}
	ctx := context.Background()
	out = make(domain.XmlPartnerSlice, 0)
	startEnd := struct {
		Start string
		End   string
		Fsrar string
	}{
		Start: start,
		End:   end,
		Fsrar: fsrarId,
	}
	templateQuery := xmlProducerShipperByPartnerSql
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
