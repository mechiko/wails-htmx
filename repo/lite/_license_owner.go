package lite

import (
	"encoding/json"
	"fmt"

	"firstwails/entity"
)

func (r *dbReport) GetLicenseWithOwner() (out *entity.LicenseList, err error) {
	defer func() {
		// обязательно, иначе лочится рессурс репозитория для указателя на db в DbService
		r.сloseDB()
		if r := recover(); r != nil {
			err = fmt.Errorf("panic %v", r)
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
		return out, fmt.Errorf("report:GetLicense %w", err)
	} else {
		for _, a := range acivity {
			m[a.Description] = a.ID
		}
	}
	out.Activity = m
	// query = `select * from licenses;`
	if err := r.db.Select(&out.Items, SelectLicenseSql); err != nil {
		return out, fmt.Errorf("report:GetLicense %w", err)
	}
	for i := range out.Items {
		license := out.Items[i]
		if partner, err := r.GetDbAlcohelp3().GetPartnerEgaisByFsrarId(license.ClientRegID); err != nil {
			return out, fmt.Errorf("report:GetLicense %w", err)
		} else {
			out.Items[i].Owner = partner
		}
		adr := &entity.AdrType{Country: "643"}
		if license.AdrTypeJson != "" {
			if err := json.Unmarshal([]byte(license.AdrTypeJson), adr); err != nil {
				return out, fmt.Errorf("dbreport:GetLicense %w", err)
			}
			license.AdrTypeStruct = adr
			out.Items[i] = license
		} else {
			license.AdrTypeStruct = adr
			out.Items[i] = license
		}
	}
	return out, nil
}

func (r *dbReport) GetLicenseWithClientMap() (mm entity.LicenseMap, err error) {
	defer func() {
		// обязательно, иначе лочится рессурс репозитория для указателя на db в DbService
		r.сloseDB()
		if r := recover(); r != nil {
			err = fmt.Errorf("panic %v", r)
		}
	}()

	mm = make(entity.LicenseMap)
	out := make([]entity.License, 0)
	if err := r.db.Select(&out, SelectLicenseSql); err != nil {
		return mm, fmt.Errorf("report:GetLicense %w", err)
	}
	for i := range out {
		license := out[i]
		regId := license.ClientRegID
		if partner, err := r.GetDbAlcohelp3().GetPartnerEgaisByFsrarId(regId); err != nil {
			return mm, fmt.Errorf("report:GetLicense %w", err)
		} else {
			license.Owner = partner
		}
		adr := &entity.AdrType{Country: "643"}
		if license.AdrTypeJson != "" {
			if err := json.Unmarshal([]byte(license.AdrTypeJson), adr); err != nil {
				return mm, fmt.Errorf("dbreport:GetLicense %w", err)
			}
			license.AdrTypeStruct = adr
		} else {
			license.AdrTypeStruct = adr
		}
		mm[regId] = &license
	}
	return mm, nil
}
