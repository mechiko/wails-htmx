package lite

import (
	_ "embed"
	"encoding/json"
	"fmt"

	"firstwails/entity"
)

//go:embed sql/update_license.sql
var UpdateLicenseSql string

func (r *dbLite) UpdateLicense(l *entity.License) (err error) {
	defer func() {
		r.сloseDB()
		if r := recover(); r != nil {
			err = fmt.Errorf("panic %v", r)
		}
	}()
	if l.AdrTypeStruct == nil {
		l.AdrTypeStruct = &entity.AdrType{Country: "643"}
		l.AdrTypeJson = ""
	}
	if str, err := json.Marshal(l.AdrTypeStruct); err != nil {
		l.AdrTypeJson = ""
		r.GetLogger().Errorf("report:UpdateLicense %w", err)
	} else {
		l.AdrTypeJson = string(str)
	}
	if _, err := r.db.Exec(UpdateLicenseSql, l.ClientRegID, l.Vid, l.Vid00, l.P000000000011, l.P000000000012, l.P000000000013, l.P000000000014, l.IsRetail, l.IsWholesale, l.IsProducer, l.AdrTypeJson); err != nil {
		return fmt.Errorf("report:UpdateLicense %w", err)
	}
	return nil
}
