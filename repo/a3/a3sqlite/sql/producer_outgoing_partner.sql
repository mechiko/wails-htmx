-- все производители по исходящим ТТН и контрагенту за период
SELECT DISTINCT 
	tp.producer_client_reg_id, 
	tp.producer_inn,
	tp.producer_kpp,
	tp.producer_short_name,
	cast(tp.producer_full_name as varchar(500)) as 'producer_full_name',
	tp.producer_country_code,
	tp.producer_region_code,
	cast(tp.producer_description as varchar(500)) as 'producer_description'
FROM ttn tt 
  join ttn_products tp on tp.id_ttn = tt.id
where 
  tt.status in ('Подтверждён','Проведён') 
  and tt.ttn_type in ('Исходящий', 'Импорт')
	and tt.consignee_client_reg_id = '{{.Fsrar}}'
  and tt.doc_date >= '{{.Start}}'
  and tt.doc_date <= '{{.End}}'
ORDER BY tp.producer_client_reg_id
;
