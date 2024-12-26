-- постановка на баланс
select
  COALESCE((SELECT TOP 1 SUBSTRING(ticket_date,1,10) FROM tickets t WHERE transport_id = tt.reply_id ), tt.doc_date) as 'fix_date',
  --'' as 'fix_date_asiiu',
  'Постановка на баланс' as 'type',
  'На баланс' as 'doc_type',
  tt.id as 'doc_id',
  tp.id as 'content_id',
  tt.doc_number as 'doc_number',
  tt.doc_date as 'doc_date',
  '' as 'doc_source',
  'на баланс' as 'doc_reason',
  tp.product_unit_type as 'product_unit_type',
  COALESCE(cast(tp.product_quantity as float),0) as 'quantity',
  COALESCE(cast(tp.product_quantity as float),0) as 'quantity_fact',
  tp.product_full_name as 'full_name',
  tp.product_code as 'code',
  tp.product_alc_code as 'alc_code',
  cast(tp.product_alc_volume as float) as 'alc_volume',
  cast(tp.product_capacity as float) as 'capacity',
  '' as 'partner_fsrar_id',
  tp.producer_client_reg_id as 'producer_fsrar_id',
  cast(tp.product_alc_volume as float) as 'alc_volume_fact',
  coalesce((SELECT TOP 1 product_inform_f1_reg_id FROM form1_egais WHERE product_inform_f1_reg_id = tfc.product_inform_f1_reg_id), concat('отсутствует ', tfc.product_inform_f1_reg_id)) as 'form1',
  COALESCE(tfc.product_inform_f2_reg_id, 'отсутствует') as 'form2',
  '' as 'fix_number',
  COALESCE(tf.doc_reg_id,'отсутствует') as 'doc_reg_id',
  coalesce((SELECT TOP 1 bottling_date FROM form1_egais WHERE product_inform_f1_reg_id = tfc.product_inform_f1_reg_id), 'отсутствует') as 'bottling_date',
  cast(case tp.product_unit_type 
        when 'Packed' then cast(tp.product_quantity as float) 
	    else 0 end as int) as 'counts',
  case tp.product_unit_type 
    when 'Unpacked' then cast(tp.product_quantity as float)
    else 0 end as 'dal',
  round(case tp.product_unit_type 
    when 'Unpacked' then cast(tp.product_quantity as float) 
    else  cast(tp.product_quantity as float) * cast(tp.product_capacity as float) * 0.1 end
    ,4) as 'volume',
    0.0 as 'summ'
FROM charge_on_acts tt join charge_on_products tp on tp.id_charge_on_acts = tt.id
  left join charge_on_form1 tf on tf.id_charge_on_acts = tt.id
  left join charge_on_form1_content tfc on tfc.id_charge_on_form1 = tf.id and tfc.product_iddomain = tp.product_iddomain 
where 
	tt.status in ('Проведён') 
  and tt.doc_date >= '{{.Start}}'
  and tt.doc_date <= '{{.End}}'
order by fix_date, tp.product_code, tp.product_alc_code, doc_id, content_id
;
