package domain

import (
	"fmt"
	"reflect"
)

type TTNFilter struct {
	ID                   string
	CreateDate           string
	TTNType              string
	DocIdentity          string
	DocType              string
	DocNumber            string
	DocDate              string
	DocShippingDate      string
	DocBase              string
	DocComment           string
	ShipperType          string
	ShipperClientRegID   string
	ShipperInn           string
	ShipperKPP           string
	ShipperFullName      string
	ShipperShortName     string
	ShipperCountryCode   string
	ShipperRegionCode    string
	ShipperDescription   string
	ConsigneeType        string
	ConsigneeClientRegID string
	ConsigneeInn         string
	ConsigneeKPP         string
	ConsigneeFullName    string
	ConsigneeShortName   string
	ConsigneeCountryCode string
	ConsigneeRegionCode  string
	ConsigneeDescription string
	TranType             string
	TransportCompany     string
	TransportCustomer    string
	TransportOwnership   string
	TransportType        string
	TransportDriver      string
	TransportTrailer     string
	TransportRegNumber   string
	TransportForwarder   string
	TransportLoadPoint   string
	TransportUnloadPoint string
	TransportRedirect    string
	Version              string
	State                string
	Status               string
	ReplyID              string
	Archive              string
}

func NewTTNFilter() *TTNFilter {
	filter := &TTNFilter{}
	return filter
}

// все поля фильтра строки
func (f *TTNFilter) ByName(name string) (out string) {
	defer func() {
		if r := recover(); r != nil {
			out = fmt.Sprintf("%v", r)
		}
	}()
	valueField := reflect.ValueOf(f).Elem().FieldByName(name)
	return valueField.String()
}

// все поля фильтра строки
func (f *TTNFilter) SetByName(name string, value string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	reflect.ValueOf(f).Elem().FieldByName(name).SetString(value)
	return nil
}

func (f *TTNFilter) Clear() {
	f.ID = ""
	f.CreateDate = ""
	f.TTNType = ""
	f.DocIdentity = ""
	f.DocType = ""
	f.DocNumber = ""
	f.DocDate = ""
	f.DocShippingDate = ""
	f.DocBase = ""
	f.DocComment = ""
	f.ShipperType = ""
	f.ShipperClientRegID = ""
	f.ShipperInn = ""
	f.ShipperKPP = ""
	f.ShipperFullName = ""
	f.ShipperShortName = ""
	f.ShipperCountryCode = ""
	f.ShipperRegionCode = ""
	f.ShipperDescription = ""
	f.ConsigneeType = ""
	f.ConsigneeClientRegID = ""
	f.ConsigneeInn = ""
	f.ConsigneeKPP = ""
	f.ConsigneeFullName = ""
	f.ConsigneeShortName = ""
	f.ConsigneeCountryCode = ""
	f.ConsigneeRegionCode = ""
	f.ConsigneeDescription = ""
	f.TranType = ""
	f.TransportCompany = ""
	f.TransportCustomer = ""
	f.TransportOwnership = ""
	f.TransportType = ""
	f.TransportDriver = ""
	f.TransportTrailer = ""
	f.TransportRegNumber = ""
	f.TransportForwarder = ""
	f.TransportLoadPoint = ""
	f.TransportUnloadPoint = ""
	f.TransportRedirect = ""
	f.Version = ""
	f.State = ""
	f.Status = ""
	f.ReplyID = ""
	f.Archive = ""
}
