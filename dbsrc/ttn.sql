select * from ttn 
where 
  id LIKE '%{{.Filter.ID}}%'
  and create_date like '%{{.Filter.CreateDate}}%'
  and ttn_type like '%{{.Filter.TTNType}}%'
  and doc_identity like '%{{.Filter.DocIdentity}}%'
  and doc_type like '%{{.Filter.DocType}}%'
  and doc_number like '%{{.Filter.DocNumber}}%'
  and doc_date like '%{{.Filter.DocDate}}%'
  and doc_shipping_date like '%{{.Filter.DocShippingDate}}%'
  and doc_base like '%{{.Filter.DocBase}}%'
  and doc_comment like '%{{.Filter.DocComment}}%'
  and shipper_type like '%{{.Filter.ShipperType}}%'
  and shipper_client_reg_id like '%{{.Filter.ShipperClientRegID}}%'
  and shipper_inn like '%{{.Filter.ShipperInn}}%'
  and shipper_kpp like '%{{.Filter.ShipperKPP}}%'
  and shipper_full_name like '%{{.Filter.ShipperFullName}}%'
  and shipper_short_name like '%{{.Filter.ShipperShortName}}%'
  and shipper_country_code like '%{{.Filter.ShipperCountryCode}}%'
  and shipper_region_code like '%{{.Filter.ShipperRegionCode}}%'
  and shipper_description like '%{{.Filter.ShipperDescription}}%'
  and consignee_type like '%{{.Filter.ConsigneeType}}%'
  and consignee_client_reg_id like '%{{.Filter.ConsigneeClientRegID}}%'
  and consignee_inn like '%{{.Filter.ConsigneeInn}}%'
  and consignee_kpp like '%{{.Filter.ConsigneeKPP}}%'
  and consignee_full_name like '%{{.Filter.ConsigneeFullName}}%'
  and consignee_short_name like '%{{.Filter.ConsigneeShortName}}%'
  and consignee_country_code like '%{{.Filter.ConsigneeCountryCode}}%'
  and consignee_region_code like '%{{.Filter.ConsigneeRegionCode}}%'
  and consignee_description like '%{{.Filter.ConsigneeDescription}}%'
  and tran_type like '%{{.Filter.TranType}}%'
  and transport_company like '%{{.Filter.TransportCompany}}%'
  and transport_customer like '%{{.Filter.TransportCustomer}}%'
  and transport_ownership like '%{{.Filter.TransportOwnership}}%'
  and transport_type like '%{{.Filter.TransportType}}%'
  and transport_driver like '%{{.Filter.TransportDriver}}%'
  and transport_trailer like '%{{.Filter.TransportTrailer}}%'
  and transport_reg_number like '%{{.Filter.TransportRegNumber}}%'
  and transport_forwarder like '%{{.Filter.TransportForwarder}}%'
  and transport_load_point like '%{{.Filter.TransportLoadPoint}}%'
  and transport_unload_point like '%{{.Filter.TransportUnloadPoint}}%'
  and transport_redirect like '%{{.Filter.TransportRedirect}}%'
  and version like '%{{.Filter.Version}}%'
  and state like '%{{.Filter.State}}%'
  and status like '%{{.Filter.Status}}%'
  and reply_id like '%{{.Filter.ReplyID}}%'
  and archive like '%{{.Filter.Archive}}%'
order by id desc limit {{.Limit}} offset {{.Offset}};