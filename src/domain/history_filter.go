package domain

import (
	"fmt"
	"reflect"
)

type IFilter interface {
	ByName(name string) (out string)
	SetByName(name string, value string) (err error)
	Clear()
}

type IDataSet interface {
	ByName(name string) (out string)
}

type Filter struct {
	Type            string
	DocType         string
	DocId           string
	DocContentId    string
	DocNumber       string
	DocDate         string
	DocReason       string
	DocSource       string
	DocRegId        string
	PartnerFsrarId  string
	FixNumber       string
	FixDate         string
	FixDateASIIU    string
	UnitType        string
	Quantity        string
	QuantityFact    string
	FullName        string
	AlcVolume       string
	AlcVolumeFact   string
	Code            string
	AlcCode         string
	Capacity        string
	ProducerFsrarId string
	Form1           string
	Form2           string
	BottlingDate    string
	Status          string
	Counts          string
	Dal             string
	Vol             string
	BVol            string
	Summ            string
}

func NewFilter() *Filter {
	filter := &Filter{}
	return filter
}

// все поля фильтра строки
func (f *Filter) ByName(name string) (out string) {
	defer func() {
		if r := recover(); r != nil {
			out = fmt.Sprintf("%v", r)
		}
	}()
	valueField := reflect.ValueOf(f).Elem().FieldByName(name)
	return valueField.String()
}

// все поля фильтра строки
func (f *Filter) SetByName(name string, value string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	reflect.ValueOf(f).Elem().FieldByName(name).SetString(value)
	return nil
}

func (f *Filter) Clear() {
	f.Type = ""
	f.DocType = ""
	f.DocId = ""
	f.DocContentId = ""
	f.DocNumber = ""
	f.DocDate = ""
	f.DocReason = ""
	f.DocSource = ""
	f.DocRegId = ""
	f.PartnerFsrarId = ""
	f.FixNumber = ""
	f.FixDate = ""
	f.FixDateASIIU = ""
	f.UnitType = ""
	f.Quantity = ""
	f.QuantityFact = ""
	f.FullName = ""
	f.AlcVolume = ""
	f.AlcVolumeFact = ""
	f.Code = ""
	f.AlcCode = ""
	f.Capacity = ""
	f.ProducerFsrarId = ""
	f.Form1 = ""
	f.Form2 = ""
	f.BottlingDate = ""
	f.Status = ""
	f.Counts = ""
	f.Dal = ""
	f.Vol = ""
	f.BVol = ""
	f.Summ = ""
}
