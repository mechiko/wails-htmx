WITH act_last_diff AS (select ta1.product_identity, ta1.product_quantity, ta2.id_ttn, ta2.act_type, ta2.act_date, ta2.id 
from ttn_acts_content ta1 JOIN ttn_acts ta2 ON ta1.id_ttn_acts = ta2.id 
where ta2.id in (select max(ID) from ttn_acts group by id_ttn))
-- расход и возвраты от меня
select
  COALESCE((SELECT doc_fix_date FROM ttn_form2 WHERE id_ttn = tt.id), tt.doc_date) as 'fix_date',
  'ТТН' as 'type',
  IIF(tt.doc_type='Возврат от меня','Возврат покупателю', 'Расход') as 'doc_type',
  tt.id as 'doc_id',
  tp.id as 'content_id',
  tt.doc_number as 'doc_number',
  tt.doc_date as 'doc_date',
  coalesce((SELECT concat(ttn_number,' ',ttn_date) FROM form2_egais WHERE product_inform_f2_reg_id = tp.product_inform_f2_reg_id),concat('отсутствует ',tp.product_inform_f2_reg_id)) as 'doc_source',
  IIF(tac.id > 0 AND round(cast(tac.product_quantity as float),3) <> round(cast(tp.product_quantity as float),3),'отгрузка с расхождением','отгрузка') as 'doc_reason',
  tp.product_unit_type as 'product_unit_type',
  COALESCE(cast(tp.product_quantity as float),0) as 'quantity',
  COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) as 'quantity_fact',
  tp.product_full_name as 'full_name',
  tp.product_code as 'code',
  tp.product_alc_code as 'alc_code',
  cast(tp.product_alc_volume as float) as 'alc_volume',
  cast(tp.product_capacity as float) as 'capacity',
  tt.consignee_client_reg_id as 'partner_fsrar_id',
  tp.producer_client_reg_id as 'producer_fsrar_id',
  coalesce((SELECT CAST(product_alc_volume AS float) FROM form1_egais WHERE product_inform_f1_reg_id = tp.product_inform_f1_reg_id ), 0.0) as 'alc_volume_fact',
  coalesce((SELECT product_inform_f1_reg_id FROM form1_egais WHERE product_inform_f1_reg_id = tp.product_inform_f1_reg_id ), concat('отсутствует ', tp.product_inform_f1_reg_id)) as 'form1',
  coalesce((SELECT product_inform_f2_reg_id FROM form2_egais WHERE product_inform_f2_reg_id = tp.product_inform_f2_reg_id ), concat('отсутствует ', tp.product_inform_f2_reg_id)) as 'form2',
  COALESCE((SELECT doc_fix_number FROM ttn_form2 WHERE id_ttn = tt.id), tt.doc_date) as 'fix_number',
  COALESCE((SELECT doc_reg_id FROM ttn_form2 WHERE id_ttn = tt.id), '') as 'doc_reg_id',
  coalesce((SELECT bottling_date FROM form1_egais WHERE product_inform_f1_reg_id = tp.product_inform_f1_reg_id ), concat('отсутствует ', tp.product_inform_f1_reg_id)) as 'bottling_date',
  cast(case tp.product_unit_type 
        when 'Packed' then COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) 
	    else 0 end as int) as 'counts',
  case tp.product_unit_type 
    when 'Unpacked' then COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) 
    else 0 end as 'dal',
  round(case tp.product_unit_type 
    when 'Unpacked' then COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) 
    else  COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) * cast(tp.product_capacity as float) * 0.1 end
    ,4) as 'volume',
  COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) * cast(tp.product_price as float) as 'summ'
FROM ttn tt 
  join ttn_products tp on tp.id_ttn = tt.id
  left join act_last_diff tac on tac.id_ttn = tt.id AND tac.product_identity = tp.product_identity
where tt.status in ('Подтверждён','Проведён') 
  and tt.ttn_type in ('Исходящий', 'Импорт')
  and tt.doc_date >= '{{.Start}}'
  and tt.doc_date <= '{{.End}}'
union all
-- приход и возврат мне 
select
  COALESCE((SELECT doc_fix_date FROM ttn_form2 WHERE id_ttn = tt.id ), tt.doc_date) as 'fix_date',
  'ТТН' as 'type',
  IIF(tt.doc_type='Возврат от меня','Возврат от покупателя', 'Приход') as 'doc_type',
  tt.id as 'doc_id',
  tp.id as 'content_id',
  tt.doc_number as 'doc_number',
  tt.doc_date as 'doc_date',
  coalesce((SELECT concat(ttn_number,' ',ttn_date) FROM form2_egais WHERE product_inform_f2_reg_id = tp.product_inform_f2_reg_id ),concat('отсутствует ',tp.product_inform_f2_reg_id)) as 'doc_source',
  IIF(tac.id > 0 AND round(cast(tac.product_quantity as float),3) <> round(cast(tp.product_quantity as float),3),'поступление с расхождением','поступление') as 'doc_reason',
  tp.product_unit_type as 'product_unit_type',
  COALESCE(cast(tp.product_quantity as float),0) as 'quantity',
  COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) as 'quantity_fact',
  tp.product_full_name as 'full_name',
  tp.product_code as 'code',
  tp.product_alc_code as 'alc_code',
  cast(tp.product_alc_volume as float) as 'alc_volume',
  cast(tp.product_capacity as float) as 'capacity',
  tt.shipper_client_reg_id as 'partner_fsrar_id',
  tp.producer_client_reg_id as 'producer_fsrar_id',
  coalesce((SELECT CAST(product_alc_volume AS float) FROM form1_egais WHERE product_inform_f1_reg_id = tp.product_inform_f1_reg_id ), 0.0) as 'alc_volume_fact',
  coalesce((SELECT product_inform_f1_reg_id FROM form1_egais WHERE product_inform_f1_reg_id = tp.product_inform_f1_reg_id ), concat('отсутствует ', tp.product_inform_f1_reg_id)) as 'form1',
  coalesce((SELECT product_inform_f2_reg_id FROM form2_egais WHERE product_inform_f2_reg_id = tp.product_inform_f2_reg_id ), concat('отсутствует ', tp.product_inform_f2_reg_id)) as 'form2',
  COALESCE((SELECT doc_fix_number FROM ttn_form2 WHERE id_ttn = tt.id), tt.doc_date) as 'fix_number',
  COALESCE((SELECT doc_reg_id FROM ttn_form2 WHERE id_ttn = tt.id), '') as 'doc_reg_id',
  coalesce((SELECT bottling_date FROM form1_egais WHERE product_inform_f1_reg_id = tp.product_inform_f1_reg_id ), concat('отсутствует ', tp.product_inform_f1_reg_id)) as 'bottling_date',
  cast(case tp.product_unit_type 
        when 'Packed' then COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) 
	    else 0 end as int) as 'counts',
  case tp.product_unit_type 
    when 'Unpacked' then COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) 
    else 0 end as 'dal',
  round(case tp.product_unit_type 
    when 'Unpacked' then COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) 
    else  COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) * cast(tp.product_capacity as float) * 0.1 end
    ,4) as 'volume',
  COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) * cast(tp.product_price as float) as 'summ'
FROM ttn tt 
  join ttn_products tp on tp.id_ttn = tt.id
  left join act_last_diff tac on tac.id_ttn = tt.id AND tac.product_identity = tp.product_identity
where tt.status in ('Подтверждён','Проведён') 
  and tt.ttn_type in ('Входящий')
  and tt.doc_date >= '{{.Start}}'
  and tt.doc_date <= '{{.End}}'
UNION ALL 
-- производство не до конца понимаю что с датой фиксации ... по сути это дата производства ... но для пивняков обычных это дата тикета...
select
  --COALESCE((SELECT SUBSTRING(ticket_date,1,10) FROM tickets t WHERE transport_id = tt.reply_id ), tt.doc_date) as 'fix_date',
  tt.doc_produced_date as 'fix_date_asiiu',
  'Производство' as 'type',
  'Приход' as 'doc_type',
  tt.id as 'doc_id',
  tp.id as 'content_id',
  tt.doc_number as 'doc_number',
  tt.doc_date as 'doc_date',
  '' as 'doc_source',
  'производство' as 'doc_reason',
  tp.product_unit_type as 'product_unit_type',
  COALESCE(cast(tp.product_quantity as float),0) as 'quantity',
  COALESCE(cast(tp.product_quantity as float),0) as 'quantity_fact',
  tp.product_full_name as 'full_name',
  tp.product_code as 'code',
  tp.product_alc_code as 'alc_code',
  cast(tp.product_alc_volume as float) as 'alc_volume',
  cast(tp.product_capacity as float) as 'capacity',
  '' as 'partner_fsrar_id',
  tt.producer_client_reg_id as 'producer_fsrar_id',
  cast(tp.product_alc_volume as float) as 'alc_volume_fact',
  coalesce((SELECT product_inform_f1_reg_id FROM form1_egais WHERE product_inform_f1_reg_id = tfc.product_inform_f1_reg_id ), concat('отсутствует ', tfc.product_inform_f1_reg_id)) as 'form1',
  COALESCE(tfc.product_inform_f2_reg_id, 'отсутствует') as 'form2',
  '' as 'fix_number',
  COALESCE(tf.doc_reg_id,'отсутствует') as 'doc_reg_id',
  tt.doc_produced_date as 'bottling_date',
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
FROM production_reports tt join production_products tp on tp.id_production_reports = tt.id
  left join production_form1 tf on tf.id_production_reports = tt.id
  left join production_form1_content tfc on tfc.id_production_form1 = tf.id and tfc.product_identity = tp.product_identity 
where 
  tt.status in ('Проведён') 
  and tt.doc_produced_date >= '{{.Start}}'
  and tt.doc_produced_date <= '{{.End}}'
union ALL 
-- импорт
select
  COALESCE((SELECT SUBSTRING(ticket_date,1,10) FROM tickets t WHERE transport_id = tt.reply_id), tt.doc_date) as 'fix_date',
  'Импорт' as 'type',
  'Приход' as 'doc_type',
  tt.id as 'doc_id',
  tp.id as 'content_id',
  tt.doc_number as 'doc_number',
  tt.doc_date as 'doc_date',
  '' as 'doc_source',
  'импорт' as 'doc_reason',
  tp.product_unit_type as 'product_unit_type',
  COALESCE(cast(tp.product_quantity as float),0) as 'quantity',
  COALESCE(cast(tp.product_quantity as float),0) as 'quantity_fact',
  tp.product_full_name as 'full_name',
  tp.product_code as 'code',
  tp.product_alc_code as 'alc_code',
  cast(tp.product_alc_volume as float) as 'alc_volume',
  cast(tp.product_capacity as float) as 'capacity',
  '' as 'partner_fsrar_id',
  tt.importer_client_reg_id as 'producer_fsrar_id',
  cast(tp.product_alc_volume as float) as 'alc_volume_fact',
  coalesce((SELECT product_inform_f1_reg_id FROM form1_egais WHERE product_inform_f1_reg_id = tfc.product_inform_f1_reg_id), concat('отсутствует ', tfc.product_inform_f1_reg_id)) as 'form1',
  tfc.product_inform_f2_reg_id as 'form2',
  '' as 'fix_number',
  COALESCE(tf.doc_reg_id,'отсутствует') as 'doc_reg_id',
  tt.gtd_date as 'bottling_date',
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
FROM import_reports tt join import_products tp on tp.id_import_reports = tt.id
  left join import_form1 tf on tf.id_import_reports = tt.id
  left join import_form1_content tfc on tfc.id_import_form1 = tf.id and tfc.product_identity = tp.product_identity 
where 
  tt.status in ('Проведён') 
  and tt.doc_date >= '{{.Start}}'
  and tt.doc_date <= '{{.End}}'
UNION ALL 
-- акты списания
SELECT
  COALESCE((SELECT SUBSTRING(ticket_date,1,10) FROM tickets t WHERE transport_id = tt.reply_id), tt.doc_date) as 'fix_date',
  'Акт' as 'type',
  'Списание' as 'doc_type',
  tt.id as 'doc_id',
  tp.id as 'content_id',
  tt.doc_number as 'doc_number',
  tt.doc_date as 'doc_date',
  '' as 'doc_source',
  tt.doc_type as 'doc_reason',
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
  coalesce((SELECT CAST(product_alc_volume AS float) FROM form1_egais WHERE product_inform_f1_reg_id = tp.product_inform_f1_reg_id ), 0.0) as 'alc_volume_fact',
  coalesce((SELECT product_inform_f1_reg_id FROM form1_egais WHERE product_inform_f1_reg_id = tp.product_inform_f1_reg_id ), concat('отсутствует ', tp.product_inform_f1_reg_id)) as 'form1',
  coalesce((SELECT product_inform_f2_reg_id FROM form2_egais WHERE product_inform_f2_reg_id = tp.product_inform_f2_reg_id ), concat('отсутствует ', tp.product_inform_f2_reg_id)) as 'form2',
  '' as 'fix_number',
  COALESCE((SELECT reg_id FROM tickets t WHERE transport_id = tt.reply_id ), '') as 'doc_reg_id',
  coalesce((SELECT bottling_date FROM form1_egais WHERE product_inform_f1_reg_id = tp.product_inform_f1_reg_id ), concat('отсутствует ', tp.product_inform_f1_reg_id)) as 'bottling_date',
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
    cast(tp.product_price as float) as 'summ'
FROM write_off_acts tt 
  join write_off_products tp on tp.id_write_off_acts = tt.id
where 
  tt.status in ('Проведён') 
  and tt.doc_date >= '{{.Start}}'
  and tt.doc_date <= '{{.End}}'
UNION ALL
-- постановка на баланс
select
  COALESCE((SELECT SUBSTRING(ticket_date,1,10) FROM tickets t WHERE transport_id = tt.reply_id ), tt.doc_date) as 'fix_date',
  --'' as 'fix_date_asiiu',
  'Постановка на баланс' as 'type',
  'Приход' as 'doc_type',
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
  coalesce((SELECT product_inform_f1_reg_id FROM form1_egais WHERE product_inform_f1_reg_id = tfc.product_inform_f1_reg_id), concat('отсутствует ', tfc.product_inform_f1_reg_id)) as 'form1',
  COALESCE(tfc.product_inform_f2_reg_id, 'отсутствует') as 'form2',
  '' as 'fix_number',
  COALESCE(tf.doc_reg_id,'отсутствует') as 'doc_reg_id',
  coalesce((SELECT bottling_date FROM form1_egais WHERE product_inform_f1_reg_id = tfc.product_inform_f1_reg_id ), 'отсутствует') as 'bottling_date',
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
  left join charge_on_form1_content tfc on tfc.id_charge_on_form1 = tf.id and tfc.product_identity = tp.product_identity 
where 
	tt.status in ('Проведён') 
  and tt.doc_date >= '{{.Start}}'
  and tt.doc_date <= '{{.End}}'
UNION ALL
-- производство ресурсы не до конца понимаю что с датой фиксации ... по сути это дата производства ... но для пивняков обычных это дата тикета...
select
  COALESCE((SELECT SUBSTRING(ticket_date,1,10) FROM tickets t WHERE transport_id = tt.reply_id), tt.doc_date) as 'fix_date',
  -- tt.doc_produced_date as 'fix_date_asiiu',
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
  cast(tp.resource_capacity as float) as 'capacity',
  '' as 'partner_fsrar_id',
  tp.producer_client_reg_id as 'producer_fsrar_id',
  coalesce((SELECT CAST(product_alc_volume AS float) FROM form1_egais WHERE product_inform_f1_reg_id = tp.resource_inform_f1_reg_id ), 0.0) as 'alc_volume_fact',
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
  and tt.doc_produced_date >= '{{.Start}}'
  and tt.doc_produced_date <= '{{.End}}'
order by fix_date, tp.product_code, tp.product_alc_code, doc_id, content_id
;
