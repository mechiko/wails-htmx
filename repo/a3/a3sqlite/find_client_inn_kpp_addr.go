package a3sqlite

import (
	"context"
	"fmt"

	"firstwails/domain"
	"firstwails/repo/a3/a3sqlite/a3boil"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// находим последний полученный запрос qm.OrderBy("id desc")
func (a *dbAlcohelp3) FindClient(fsrarId, inn, kpp, adr string) (cl *domain.PartnersEgaisImport, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()
	cl = &domain.PartnersEgaisImport{}
	whereInnKpp := `client_inn = ? and client_kpp = ?`
	whereInnAddr := `client_inn = ? and client_description = ?`
	whereFsrarId := `client_reg_id = ?`
	ctx := context.Background()
	if fsrarId != "" {
		if clMssql, err := a3boil.PartnersEgaiss(qm.Where(whereFsrarId, fsrarId), qm.OrderBy("id desc")).One(ctx, a.db); err != nil {
			return cl, fmt.Errorf("%s FindClient %w", modError, err)
		} else {
			cl.ConvertFromA3Sqlite(clMssql)
			return cl, nil
		}
	}
	// если есть адрес то ищем по паре инн и адрес
	if adr != "" {
		if clMssql, err := a3boil.PartnersEgaiss(qm.Where(whereInnAddr, inn, adr), qm.OrderBy("id desc")).One(ctx, a.db); err != nil {
			return cl, fmt.Errorf("%s FindClient %w", modError, err)
		} else {
			cl.ConvertFromA3Sqlite(clMssql)
			return cl, nil
		}
	}
	// в остальных случаях инн и кпп
	// один из вариантов (inn != "") && (kpp != "") {
	if clMssql, err := a3boil.PartnersEgaiss(qm.Where(whereInnKpp, inn, kpp), qm.OrderBy("id desc")).One(ctx, a.db); err != nil {
		return cl, fmt.Errorf("%s FindClient %w", modError, err)
	} else {
		cl.ConvertFromA3Sqlite(clMssql)
		return cl, nil
	}
}
