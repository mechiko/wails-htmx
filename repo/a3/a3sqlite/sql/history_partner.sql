WITH act_last_diff AS (select ta1.product_identity, ta1.product_quantity, ta2.id_ttn, ta2.act_type, ta2.act_date, ta2.id 
from ttn_acts_content ta1 JOIN ttn_acts ta2 ON ta1.id_ttn_acts = ta2.id 
where ta2.id in (select max(ID) from ttn_acts group by id_ttn))
-- расход и возвраты от меня
select
  COALESCE(tf.doc_fix_date, '') as 'fix_date',
  tt.status as 'status',
  'ТТН' as 'type',
  IIF(tt.doc_type='Возврат от меня','Возврат покупателю', 'Расход') as 'doc_type',
  tt.id as 'doc_id',
  tp.id as 'content_id',
  tt.doc_number as 'doc_number',
  tt.doc_date as 'doc_date',
  coalesce(concat(fb.ttn_number,' ',fb.ttn_date),concat('отсутствует ',tp.product_inform_f2_reg_id)) as 'doc_source',
  tp.product_unit_type as 'product_unit_type',
  COALESCE(cast(tp.product_quantity as float),0) as 'quantity',
  COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) as 'quantity_fact',
  tp.product_full_name as 'full_name',
  tp.product_code as 'code',
  tp.product_alc_code as 'alc_code',
  cast(tp.product_alc_volume as float) as 'alc_volume',
  coalesce(CAST(fa.product_alc_volume AS float), 0.0) as 'alc_volume_fact',
  cast(tp.product_capacity as float) as 'capacity',
  tt.consignee_client_reg_id as 'partner_fsrar_id',
  tp.producer_client_reg_id as 'producer_fsrar_id',
  tp.product_inform_f1_reg_id as 'form1',
  tp.product_inform_f2_reg_id as 'form2',
  COALESCE(tf.doc_fix_number, '') as 'fix_number',
  COALESCE(tf.doc_reg_id , '') as 'doc_reg_id',
  coalesce(fa.bottling_date, '') as 'bottling_date',
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
  LEFT join ttn_form2 tf on tf.id_ttn = tt.id
  LEFT join form1_egais fa ON fa.product_inform_f1_reg_id = tp.product_inform_f1_reg_id 
  LEFT join form2_egais fb ON fb.product_inform_f2_reg_id = tp.product_inform_f2_reg_id 
where tt.ttn_type in ('Исходящий', 'Импорт')
  and tt.status IN ('Проведён','Подтверждён')
	and tt.consignee_client_reg_id = '{{.Fsrar}}'
  and tt.doc_date >= '{{.Start}}'
  and tt.doc_date <= '{{.End}}'
union all
-- приход и возврат мне 
select
  COALESCE(tf.doc_fix_date, '') as 'fix_date',
  tt.status as 'status',
  'ТТН' as 'type',
  IIF(tt.doc_type='Возврат от меня','Возврат покупателю', 'Расход') as 'doc_type',
  tt.id as 'doc_id',
  tp.id as 'content_id',
  tt.doc_number as 'doc_number',
  tt.doc_date as 'doc_date',
  coalesce(concat(fb.ttn_number,' ',fb.ttn_date),concat('отсутствует ',tp.product_inform_f2_reg_id)) as 'doc_source',
  tp.product_unit_type as 'product_unit_type',
  COALESCE(cast(tp.product_quantity as float),0) as 'quantity',
  COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) as 'quantity_fact',
  tp.product_full_name as 'full_name',
  tp.product_code as 'code',
  tp.product_alc_code as 'alc_code',
  cast(tp.product_alc_volume as float) as 'alc_volume',
  coalesce(CAST(fa.product_alc_volume AS float), 0.0) as 'alc_volume_fact',
  cast(tp.product_capacity as float) as 'capacity',
  tt.consignee_client_reg_id as 'partner_fsrar_id',
  tp.producer_client_reg_id as 'producer_fsrar_id',
  tp.product_inform_f1_reg_id as 'form1',
  tp.product_inform_f2_reg_id as 'form2',
  COALESCE(tf.doc_fix_number, '') as 'fix_number',
  COALESCE(tf.doc_reg_id , '') as 'doc_reg_id',
  coalesce(fa.bottling_date, '') as 'bottling_date',
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
  LEFT join ttn_form2 tf on tf.id_ttn = tt.id
  LEFT join form1_egais fa ON fa.product_inform_f1_reg_id = tp.product_inform_f1_reg_id 
  LEFT join form2_egais fb ON fb.product_inform_f2_reg_id = tp.product_inform_f2_reg_id 
where tt.ttn_type = 'Входящий'
  and tt.status IN ('Проведён','Подтверждён')
	and tt.shipper_client_reg_id = '{{.Fsrar}}'
  and tt.doc_date >= '{{.Start}}'
  and tt.doc_date <= '{{.End}}'
order by fix_date, tp.product_code, tp.product_alc_code, doc_id, content_id
;
