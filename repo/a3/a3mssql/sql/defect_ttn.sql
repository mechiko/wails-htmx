-- проблемные ТТН
WITH act_last_diff AS (select ta1.product_iddomain, ta1.product_quantity, ta2.id_ttn, ta2.act_type, ta2.act_date, ta2.id 
from ttn_acts_content ta1 JOIN ttn_acts ta2 ON ta1.id_ttn_acts = ta2.id 
where ta2.id in (select max(ID) from ttn_acts group by id_ttn)),
act_last AS (SELECT ta.id_ttn, ta.act_type, ta.act_date, ta.id FROM ttn_acts ta WHERE EXISTS (SELECT LAST_VALUE(id) over (PARTITION BY id_ttn order by id ROWS BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING) AS 'id' FROM ttn_acts WHERE id_ttn = ta.id_ttn))
select 
  tt.consignee_client_reg_id as 'client_reg_id',
  tt.consignee_short_name as 'short_name',
  tt.doc_number as 'doc_number', 
  tt.doc_date as 'doc_date', 
  tf.doc_fix_date as 'fix_date',
  COALESCE(takt.act_type, '') as 'last_act_type',
  COALESCE(takt.act_date, '') as 'last_act_date',
  COALESCE(tac.act_date, '') as 'act_diff_date',
  COALESCE(tat.ticket_type , '') as 'act_diff_reject',
  sum(round(case tp.product_unit_type 
    when 'Unpacked' then COALESCE(cast(tp.product_quantity as float),0) 
    else  COALESCE(cast(tp.product_quantity as float),0) * cast(tp.product_capacity as float) * 0.1 end
    ,4)) as 'volume',
  sum(IIF(tat.ticket_type = 'Подтверждение', round(case tp.product_unit_type 
    when 'Unpacked' then COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) 
    else  COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) * cast(tp.product_capacity as float) * 0.1 end
    ,4), round(case tp.product_unit_type 
    when 'Unpacked' then COALESCE(cast(tp.product_quantity as float),0) 
    else  COALESCE(cast(tp.product_quantity as float),0) * cast(tp.product_capacity as float) * 0.1 end
    ,4))) as 'diff'
from ttn tt
  join ttn_form2 tf on tf.id_ttn = tt.id
  join ttn_products tp on tp.id_ttn = tt.id
  left join act_last_diff tac on tac.id_ttn = tt.id AND  tac.product_iddomain = tp.product_iddomain
  left join ttn_acts_tickets tat on tat.id_ttn_acts = tac.id
  left join act_last takt ON takt.id_ttn = tt.id 
where tt.ttn_type in ('Исходящий', 'Импорт')
  and tf.doc_fix_date >= '{{.Start}}'
  and tf.doc_fix_date <= '{{.End}}'
  and ((takt.act_type in ('Расхождение', 'Отказ', '')) or (tf.doc_fix_date <> tt.doc_date))
group by tt.id, tt.doc_number, tt.doc_date, tf.doc_fix_date, takt.act_type, takt.act_date, tac.act_date, tt.consignee_client_reg_id, tt.consignee_short_name, tat.ticket_type
;
