package a3mssql

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"

	"firstwails/domain"
	"firstwails/repo/a3/a3mssql/a3boil"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (a *dbAlcohelp3) LocalRest() (out domain.RestMap, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()
	var f1Map map[string]*a3boil.Form1Egais = make(map[string]*a3boil.Form1Egais)
	ctx := context.Background()
	out = make(domain.RestMap)
	if f1, err := a3boil.Form1Egaiss().All(ctx, a.db); err != nil {
		return out, fmt.Errorf("%s %w", modError, err)
	} else {
		// заполняем мапу справками F1
		for i, f1Info := range f1 {
			if _, ok := f1Map[f1Info.ProductInformF1RegID.String]; !ok {
				f1Map[f1Info.ProductInformF1RegID.String] = f1[i]
			}
		}
	}
	if qrowList, err := a3boil.RestsApLocals(qm.OrderBy("product_alc_code, product_inform_f1_reg_id, product_quantity")).All(ctx, a.db); err != nil {
		return out, fmt.Errorf("%s %w", modError, err)
	} else {
		for i, ql := range qrowList {
			bottling := "9999.12.31"
			// a.GetLogger().Debugf("rest %d %s %s", i, ql.ProductAlcCode.String, ql.ProductInformF1RegID.String)
			if f1, ok := f1Map[ql.ProductInformF1RegID.String]; ok {
				// a.GetLogger().Debugf("   f1 bottling %s", i, f1.BottlingDate.String)
				bottling = f1.BottlingDate.String
			}
			qStr := strings.Trim(ql.ProductQuantity.String, " ")
			quantity := 0.0
			if quant, err := strconv.ParseFloat(qStr, 64); err != nil {
				return out, fmt.Errorf("%s ParseFloat ProductQuantity %w", modError, err)
			} else {
				// math.Round(x*100)/100
				// quantity = quant
				quantity = math.Round(quant*10000) / 10000
			}
			// формируем запись для остатка из строк самого остатка и справки 1 для даты розлива
			rstLocal := &domain.RestLocal{
				Bottling:    bottling,
				Quantity:    quantity,
				QuantityStr: fmt.Sprintf("%012.4f", quantity),
			}
			rstLocal.ConvertFromA3Mssql(qrowList[i])
			if slice, ok := out[ql.ProductAlcCode.String]; !ok {
				slice = make(domain.RestLocalSlice, 0)
				slice = append(slice, rstLocal)
				out[ql.ProductAlcCode.String] = slice
			} else {
				slice = append(slice, rstLocal)
				out[ql.ProductAlcCode.String] = slice
			}
		}
	}
	return out, nil
}
