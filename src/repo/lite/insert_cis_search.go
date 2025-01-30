package lite

import (
	"context"
	"encoding/json"
	"firstwails/domain"
	"firstwails/repo/lite/liteboil"
	"fmt"
	"strings"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (a *dbLite) CisRequestAll() (out liteboil.CisRequestSlice, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()
	ctx := context.Background()
	if sl, err := liteboil.CisRequests().All(ctx, a.db); err != nil {
		return nil, fmt.Errorf("%w", err)
	} else {
		return sl, nil
	}
}

func (a *dbLite) InsertCisRequest(in []domain.Cises) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()
	ctx := context.Background()
	for i := range in {
		jsonBlob, err := json.Marshal(in[i])
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		cis := &liteboil.CisRequest{
			Cis:      in[i].Cis,
			Status:   in[i].Status,
			Responce: null.NewBytes(jsonBlob, true),
		}
		if err := cis.Insert(ctx, a.db, boil.Infer()); err != nil {
			return fmt.Errorf("%w", err)
		}

	}
	return nil
}

func (a *dbLite) CisRequestDeleteAll() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()

	sql := "delete from cis_request;"
	if _, err = a.db.Exec(sql); err != nil {
		return fmt.Errorf("%s %v", modError, err)
	}

	// ctx := context.Background()
	// if count, err := liteboil.CisRequests().DeleteAll(ctx, a.db); err != nil {
	// 	return fmt.Errorf("%w", err)
	// } else {
	// 	a.Logger().Debugf("cis_request delete %d records", count)
	// }
	return nil
}

func (a *dbLite) InsertCisRequestPost(in []domain.CisesPost) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()
	ctx := context.Background()
	for i := range in {
		// jsonBlob, err := json.Marshal(in[i])
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		cis := &liteboil.CisRequest{
			Cis:    in[i].Result.Cis,
			Status: strings.ToLower(in[i].Result.Status),
			// Responce: null.NewBytes(jsonBlob, true),
		}
		if err := cis.Insert(ctx, a.db, boil.Infer()); err != nil {
			return fmt.Errorf("%w", err)
		}

	}
	return nil
}
