-- производство ресурсы не до конца понимаю что с датой фиксации ... по сути это дата производства ... но для пивняков обычных это дата тикета...
select
  COALESCE((SELECT SUBSTRING(ticket_date,1,10) FROM tickets t WHERE transport_id = tt.reply_id), tt.doc_date) as 'fix_date',
  -- tt.doc_produced_date as 'fix_date_asiiu',
  tt.status as 'status',
  'Производство ресурсы' as 'type',
  'Расход' as 'doc_type',
  tt.id as 'doc_id',
  tp.id as 'content_id',
  tt.doc_number as 'doc_number',
  tt.doc_date as 'doc_date',
  '' as 'doc_source',
  'ресурсы' as 'doc_reason',
  tp.resource_unit_type as 'resource_unit_type',
  COALESCE(cast(tp.resource_quantity as float),0) as 'quantity',
  COALESCE(cast(tp.resource_quantity as float),0) as 'quantity_fact',
  tp.resource_full_name as 'full_name',
  tp.resource_code as 'code',
  tp.resource_alc_code as 'alc_code',
  cast(tp.resource_alc_volume as float) as 'alc_volume',
  coalesce((SELECT CAST(product_alc_volume AS float) FROM form1_egais WHERE product_inform_f1_reg_id = tp.resource_inform_f1_reg_id ), 0.0) as 'alc_volume_fact',
  cast(tp.resource_capacity as float) as 'capacity',
  '' as 'partner_fsrar_id',
  tp.producer_client_reg_id as 'producer_fsrar_id',
  coalesce((SELECT product_inform_f1_reg_id FROM form1_egais WHERE product_inform_f1_reg_id = tp.resource_inform_f1_reg_id ), concat('отсутствует ', tp.resource_inform_f1_reg_id)) as 'form1',
  coalesce((SELECT product_inform_f2_reg_id FROM form2_egais WHERE product_inform_f2_reg_id = tp.resource_inform_f2_reg_id ), concat('отсутствует ', tp.resource_inform_f2_reg_id)) as 'form2',
  '' as 'fix_number',
  COALESCE(tf.doc_reg_id,'отсутствует') as 'doc_reg_id',
  coalesce((SELECT bottling_date FROM form1_egais WHERE product_inform_f1_reg_id = tp.resource_inform_f1_reg_id ), concat('отсутствует ', tp.resource_inform_f1_reg_id)) as 'bottling_date',
  cast(case tp.resource_unit_type 
        when 'Packed' then cast(tp.resource_quantity as float) 
	    else 0 end as int) as 'counts',
  case tp.resource_unit_type 
    when 'Unpacked' then cast(tp.resource_quantity as float)
    else 0 end as 'dal',
  round(case tp.resource_unit_type 
    when 'Unpacked' then cast(tp.resource_quantity as float) 
    else  cast(tp.resource_quantity as float) * cast(tp.resource_capacity as float) * 0.1 end
    ,4) as 'volume',
    0.0 as 'summ'
FROM production_reports tt join production_resources tp on tp.id_production_reports = tt.id
  left join production_form1 tf on tf.id_production_reports = tt.id
where 
	tt.status in ('Проведён') 
  -- and tt.doc_date >= '{{.Start}}'
  -- and tt.doc_date <= '{{.End}}'
  and tt.doc_produced_date >= '{{.Start}}'
  and tt.doc_produced_date <= '{{.End}}'
order by fix_date, tp.resource_code, tp.resource_alc_code, doc_id, content_id
;
