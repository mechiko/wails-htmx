package lite

import (
	_ "embed"
	"encoding/json"
	"fmt"

	"firstwails/entity"
)

//go:embed sql/license.sql
var SelectLicenseSql string

func (r *dbLte) GetLicense() (out []entity.License, err error) {
	defer func() {
		// обязательно, иначе лочится рессурс репозитория для указателя на db в DbService
		r.сloseDB()
		if r := recover(); r != nil {
			err = fmt.Errorf("dbreport panic %v", r)
		}
	}()

	out = make([]entity.License, 0)
	if err := r.db.Select(&out, SelectLicenseSql); err != nil {
		return out, fmt.Errorf("dbreport:GetLicense %w", err)
	}
	for i, l := range out {
		adr := &entity.AdrType{Country: "643"}
		if !json.Valid([]byte(l.AdrTypeJson)) {
			r.GetLogger().Errorf("dbreport:GetLicense not valid json %s", l.AdrTypeJson)
			l.AdrTypeStruct = adr
			out[i] = l
		} else {
			if err := json.Unmarshal([]byte(l.AdrTypeJson), adr); err != nil {
				// return out, fmt.Errorf("dbreport:GetLicense %w", err)
				r.GetLogger().Errorf("dbreport:GetLicense %w", err)
			}
			l.AdrTypeStruct = adr
			out[i] = l
		}
	}
	return out, nil
}

func (r *dbReport) GetLicenseWithVids() (out *entity.LicenseList, err error) {
	defer func() {
		// обязательно, иначе лочится рессурс репозитория для указателя на db в DbService
		r.сloseDB()
		if r := recover(); r != nil {
			err = fmt.Errorf("dbreport panic %v", r)
		}
	}()

	m := make(entity.LicenseActivityMap)
	out = &entity.LicenseList{
		Items:    make([]*entity.License, 0),
		Activity: m,
	}
	acivity := make([]entity.LicenseActivity, 0)
	query := `select * from licenses_vid lv ;`
	if err := r.db.Select(&acivity, query); err != nil {
		return out, fmt.Errorf("dbreport:GetLicense %w", err)
	} else {
		for _, a := range acivity {
			m[a.Description] = a.ID
		}
	}
	out.Activity = m
	if err := r.db.Select(&out.Items, SelectLicenseSql); err != nil {
		return out, fmt.Errorf("dbreport:GetLicense %w", err)
	}
	for i, l := range out.Items {
		adr := &entity.AdrType{Country: "643"}
		if l.AdrTypeJson != "" {
			if err := json.Unmarshal([]byte(l.AdrTypeJson), adr); err != nil {
				return out, fmt.Errorf("dbreport:GetLicense %w", err)
			}
			l.AdrTypeStruct = adr
			out.Items[i] = l
		} else {
			l.AdrTypeStruct = adr
			out.Items[i] = l
		}
	}
	return out, nil
}
