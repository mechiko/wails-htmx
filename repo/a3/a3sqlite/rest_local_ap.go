package a3sqlite

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"

	"firstwails/domain"
	"firstwails/repo/a3/a3sqlite/a3boil"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (a *dbAlcohelp3) LocalRestAp() (out domain.RestApSlice, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()
	var f1Map map[string]*a3boil.Form1Egais = make(map[string]*a3boil.Form1Egais)
	ctx := context.Background()
	out = make(domain.RestApSlice, 0)
	if f1, err := a3boil.Form1Egaiss(qm.Where("id > 0")).All(ctx, a.db); err != nil {
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
			bottling := ""
			if f1, ok := f1Map[ql.ProductInformF1RegID.String]; ok {
				bottling = f1.BottlingDate.String
			}
			qStr := strings.Trim(ql.ProductQuantity.String, " ")
			quantity := 0.0
			if quant, err := strconv.ParseFloat(qStr, 64); err != nil {
				return out, fmt.Errorf("%s ParseFloat ProductQuantity %w", modError, err)
			} else {
				quantity = math.Round(quant*10000) / 10000
			}
			rstLocal := &domain.RestAp{
				Bottling:    bottling,
				Quantity:    quantity,
				QuantityStr: fmt.Sprintf("%012.4f", quantity),
			}
			rstLocal.ConvertFromA3Sqlite(qrowList[i])
			out = append(out, rstLocal)
		}
	}
	return out, nil
}
