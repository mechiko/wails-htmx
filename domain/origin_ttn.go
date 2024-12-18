package domain

import (
	"fmt"
	"reflect"
	"strconv"
)

type TTNOrigin struct {
	ID                   int    `db:"id"`
	CreateDate           string `db:"create_date"`
	TTNType              string `db:"ttn_type"`
	DocIdentity          string `db:"doc_identity"`
	DocType              string `db:"doc_type"`
	DocNumber            string `db:"doc_number"`
	DocDate              string `db:"doc_date"`
	DocShippingDate      string `db:"doc_shipping_date"`
	DocBase              string `db:"doc_base"`
	DocComment           string `db:"doc_comment"`
	ShipperType          string `db:"shipper_type"`
	ShipperClientRegID   string `db:"shipper_client_reg_id"`
	ShipperInn           string `db:"shipper_inn"`
	ShipperKPP           string `db:"shipper_kpp"`
	ShipperFullName      string `db:"shipper_full_name"`
	ShipperShortName     string `db:"shipper_short_name"`
	ShipperCountryCode   string `db:"shipper_country_code"`
	ShipperRegionCode    string `db:"shipper_region_code"`
	ShipperDescription   string `db:"shipper_description"`
	ConsigneeType        string `db:"consignee_type"`
	ConsigneeClientRegID string `db:"consignee_client_reg_id"`
	ConsigneeInn         string `db:"consignee_inn"`
	ConsigneeKPP         string `db:"consignee_kpp"`
	ConsigneeFullName    string `db:"consignee_full_name"`
	ConsigneeShortName   string `db:"consignee_short_name"`
	ConsigneeCountryCode string `db:"consignee_country_code"`
	ConsigneeRegionCode  string `db:"consignee_region_code"`
	ConsigneeDescription string `db:"consignee_description"`
	TranType             string `db:"tran_type"`
	TransportCompany     string `db:"transport_company"`
	TransportCustomer    string `db:"transport_customer"`
	TransportOwnership   string `db:"transport_ownership"`
	TransportType        string `db:"transport_type"`
	TransportDriver      string `db:"transport_driver"`
	TransportTrailer     string `db:"transport_trailer"`
	TransportRegNumber   string `db:"transport_reg_number"`
	TransportForwarder   string `db:"transport_forwarder"`
	TransportLoadPoint   string `db:"transport_load_point"`
	TransportUnloadPoint string `db:"transport_unload_point"`
	TransportRedirect    string `db:"transport_redirect"`
	Version              string `db:"version"`
	State                string `db:"state"`
	Status               string `db:"status"`
	ReplyID              string `db:"reply_id"`
	Archive              int    `db:"archive"`
	XML                  string `db:"xml"`
}

type TTNOriginSlice []*TTNOrigin

// все поля фильтра строки
func (f *TTNOrigin) ByName(name string) (out string) {
	defer func() {
		if r := recover(); r != nil {
			out = fmt.Sprintf("%v", r)
		}
	}()
	valueField := reflect.ValueOf(f).Elem().FieldByName(name)
	if valueField.CanInt() {
		return strconv.Itoa(int(valueField.Int()))
	}
	if valueField.CanFloat() {
		return fmt.Sprintf("%.3f", valueField.Float())
	}
	return valueField.String()
}
