package a3mssql

import (
	"context"
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"

	"firstwails/domain"
	"firstwails/repo/a3/a3mssql/a3boil"

	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type formParty struct {
	Form1 string `boil:"form1"`
	Party string `boil:"party"`
}

//go:embed sql/form1_party.sql
var form1PartySql string

func (a *dbAlcohelp3) LocalRestParty() (out domain.RestMap, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()
	var f1Map map[string]*a3boil.Form1Egais = make(map[string]*a3boil.Form1Egais)
	ctx := context.Background()
	out = make(domain.RestMap)

	forms := make([]formParty, 0)
	if err := queries.Raw(form1PartySql).Bind(ctx, a.db, &forms); err != nil {
		return out, fmt.Errorf("%s %w", modError, err)
	}
	formsMap := map[string]string{}
	// эта мапа для ловли дублей партий их исключаем иначе не поймешь как действовать
	// номер партии составной кодАп:партия ибо уникальность партии внутри кода АП
	partyMap := map[string]string{}
	for i := range forms {
		if form1, ok := partyMap[forms[i].Party]; ok {
			return out, fmt.Errorf("%s одинаковая партия %s для разных справок %s %s", modError, forms[i].Party, forms[i].Form1, form1)
		}
		partyMap[forms[i].Party] = forms[i].Form1
		formsMap[forms[i].Form1] = forms[i].Party
	}
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
			// формируем запись для остатка из строк самого остатка и справки 1 для даты розлива
			party, ok := formsMap[ql.ProductInformF1RegID.String]
			if !ok {
				return out, fmt.Errorf("%s для справки %s ненайдена партия", modError, ql.ProductInformF1RegID.String)
			}
			rstParty := &domain.RestLocal{
				Bottling:    bottling,
				Quantity:    quantity,
				QuantityStr: fmt.Sprintf("%012.4f", quantity),
				Party:       party,
			}
			rstParty.ConvertFromA3Mssql(qrowList[i])
			if slice, ok := out[party]; !ok {
				slice = make(domain.RestLocalSlice, 0)
				slice = append(slice, rstParty)
				out[party] = slice
			} else {
				slice = append(slice, rstParty)
				out[party] = slice
			}
		}
	}
	return out, nil
}
