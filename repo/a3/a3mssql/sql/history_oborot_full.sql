-- расход и возвраты от меня
select
  tf.doc_fix_date as 'fix_date',
  COALESCE(takt.act_type, 'Проведён') as 'status',
  'ТТН исходящие' as 'type',
  IIF(tt.doc_type='Возврат от меня','Возврат покупателю', 'Расход') as 'doc_type',
  tt.id as 'doc_id',
  tp.id as 'content_id',
  tt.doc_number as 'doc_number',
  tt.doc_date as 'doc_date',
  coalesce((SELECT TOP 1 concat(ttn_number,' ',ttn_date) FROM form2_egais WHERE product_inform_f2_reg_id = tp.product_inform_f2_reg_id ),concat('отсутствует ',tp.product_inform_f2_reg_id)) as 'doc_source',
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
  coalesce((SELECT TOP 1 CAST(product_alc_volume AS float) FROM form1_egais WHERE product_inform_f1_reg_id = tp.product_inform_f1_reg_id ), 0.0) as 'alc_volume_fact',
  coalesce((SELECT TOP 1 product_inform_f1_reg_id FROM form1_egais WHERE product_inform_f1_reg_id = tp.product_inform_f1_reg_id ), concat('отсутствует ', tp.product_inform_f1_reg_id)) as 'form1',
  coalesce((SELECT TOP 1 product_inform_f2_reg_id FROM form2_egais WHERE product_inform_f2_reg_id = tp.product_inform_f2_reg_id ), concat('отсутствует ', tp.product_inform_f2_reg_id)) as 'form2',
  COALESCE((SELECT TOP 1 doc_fix_number FROM ttn_form2 WHERE id_ttn = tt.id), tt.doc_date) as 'fix_number',
  COALESCE((SELECT TOP 1 reg_id FROM tickets t WHERE transport_id = tt.reply_id ), '') as 'doc_reg_id',
  coalesce((SELECT TOP 1 bottling_date FROM form1_egais WHERE product_inform_f1_reg_id = tp.product_inform_f1_reg_id ), concat('отсутствует ', tp.product_inform_f1_reg_id)) as 'bottling_date',
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
  join ttn_form2 tf on tf.id_ttn = tt.id
  left join (select id, id_ttn, act_type, act_date from ttn_acts where id in (select max(ID) from ttn_acts group by id_ttn)) takt on takt.id_ttn = tt.id
  left join (select id, id_ttn, act_type from ttn_acts where id in (select max(ID) from ttn_acts group by id_ttn)) ta on ta.id_ttn = tt.id and ta.act_type = 'Расхождение'
  left join ttn_acts_tickets tat on tat.id_ttn_acts = ta.id
  left join ttn_acts_content tac on tac.id_ttn_acts = ta.id and tac.product_identity = tp.product_identity
where tt.ttn_type in ('Исходящий', 'Импорт')
  and tf.doc_fix_date >= '{{.Start}}'
  and tf.doc_fix_date <= '{{.End}}'
  and (takt.act_type <> 'Отказ' or takt.act_type is null)
  and (tat.ticket_type <> 'Отказ' or tat.ticket_type is null)
union all
-- приход и возврат мне 
select
  tf.doc_fix_date as 'fix_date',
  COALESCE(takt.act_type, 'Проведён') as 'status',
  'ТТН входящие' as 'type',
  IIF(tt.doc_type='Возврат от меня','Возврат от покупателя', 'Приход') as 'doc_type',
  tt.id as 'doc_id',
  tp.id as 'content_id',
  tt.doc_number as 'doc_number',
  tt.doc_date as 'doc_date',
  coalesce((SELECT TOP 1 concat(ttn_number,' ',ttn_date) FROM form2_egais WHERE product_inform_f2_reg_id = tp.product_inform_f2_reg_id ),concat('отсутствует ',tp.product_inform_f2_reg_id)) as 'doc_source',
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
  coalesce((SELECT TOP 1 CAST(product_alc_volume AS float) FROM form1_egais WHERE product_inform_f1_reg_id = tp.product_inform_f1_reg_id ), 0.0) as 'alc_volume_fact',
  coalesce((SELECT TOP 1 product_inform_f1_reg_id FROM form1_egais WHERE product_inform_f1_reg_id = tp.product_inform_f1_reg_id ), concat('отсутствует ', tp.product_inform_f1_reg_id)) as 'form1',
  coalesce((SELECT TOP 1 product_inform_f2_reg_id FROM form2_egais WHERE product_inform_f2_reg_id = tp.product_inform_f2_reg_id ), concat('отсутствует ', tp.product_inform_f2_reg_id)) as 'form2',
  COALESCE((SELECT TOP 1 doc_fix_number FROM ttn_form2 WHERE id_ttn = tt.id), tt.doc_date) as 'fix_number',
  COALESCE((SELECT TOP 1 reg_id FROM tickets t WHERE transport_id = tt.reply_id ), '') as 'doc_reg_id',
  coalesce((SELECT TOP 1 bottling_date FROM form1_egais WHERE product_inform_f1_reg_id = tp.product_inform_f1_reg_id ), concat('отсутствует ', tp.product_inform_f1_reg_id)) as 'bottling_date',
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
  join ttn_form2 tf on tf.id_ttn = tt.id
  left join (select id, id_ttn, act_type, act_date from ttn_acts where id in (select max(ID) from ttn_acts group by id_ttn)) takt on takt.id_ttn = tt.id
  left join (select id, id_ttn, act_type from ttn_acts where id in (select max(ID) from ttn_acts group by id_ttn)) ta on ta.id_ttn = tt.id and ta.act_type = 'Расхождение'
  left join ttn_acts_tickets tat on tat.id_ttn_acts = ta.id
  left join ttn_acts_content tac on tac.id_ttn_acts = ta.id and tac.product_identity = tp.product_identity
where tt.ttn_type = 'Входящий'
  and tf.doc_fix_date >= '{{.Start}}'
  and tf.doc_fix_date <= '{{.End}}'
  and (takt.act_type <> 'Отказ' or takt.act_type is null)
  and (tat.ticket_type <> 'Отказ' or tat.ticket_type is null)
order by fix_date, tp.product_code, tp.product_alc_code, doc_id, content_id
;
