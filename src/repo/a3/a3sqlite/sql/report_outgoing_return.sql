-- ТТН Возврат из Расхода
WITH act_last_diff AS (select ta1.product_iddomain, ta1.product_quantity, ta2.id_ttn, ta2.act_type, ta2.act_date, ta2.id 
from ttn_acts_content ta1 JOIN ttn_acts ta2 ON ta1.id_ttn_acts = ta2.id 
where ta2.id in (select max(ID) from ttn_acts where act_date >= '{{.Start}}' and act_date <= '{{.End}}' group by id_ttn)),
act_last AS (select DISTINCT id_ttn, act_type, act_date, LAST_VALUE(id) over (PARTITION BY id_ttn order by act_date ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING) AS 'id'
from ttn_acts where act_date >= '{{.Start}}' and act_date <= '{{.End}}')
select
  tp.product_alc_code as 'alc_code',
  coalesce((SELECT CAST(product_alc_volume AS float) FROM form1_egais WHERE product_inform_f1_reg_id = tp.product_inform_f1_reg_id ), 0.0) as 'alc_volume',
  tp.product_inform_f1_reg_id as 'alc_volume_f1',
  tt.consignee_region_code AS 'region_code',
  tt.consignee_client_reg_id AS 'fsrarid',
  cast(sum(case tp.product_unit_type when 'Packed' then COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) else 0 end) as int) as 'counts',
  round(sum(case tp.product_unit_type when 'Unpacked' then COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) else 0 end),4) as 'dal',
  round(sum(case tp.product_unit_type when 'Unpacked' then COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) else  COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) * cast(tp.product_capacity as float) * 0.1 end),4) as 'volume'
FROM ttn tt join ttn_products tp on tp.id_ttn = tt.id
  join ttn_form2 tf on tf.id_ttn = tt.id
  left join act_last_diff tac on tac.id_ttn = tt.id AND  tac.product_iddomain = tp.product_iddomain
  left join act_last takt on takt.id_ttn = tt.id
  left join ttn_acts_tickets tat on tat.id_ttn_acts = tac.id
where tt.ttn_type in ('Исходящий', 'Импорт')
  and tt.doc_type = 'Возврат от меня'
  and tf.doc_fix_date >= '{{.Start}}'
  and tf.doc_fix_date <= '{{.End}}'
  and (takt.act_type <> 'Отказ' or takt.act_type is null)
  and (tat.ticket_type <> 'Отказ' or tat.ticket_type is null)
group by tt.consignee_region_code, tt.consignee_client_reg_id, tp.product_alc_code, tp.product_inform_f1_reg_id
;
