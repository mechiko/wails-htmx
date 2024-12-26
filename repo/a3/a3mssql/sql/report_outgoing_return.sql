-- ТТН Возврат из Расхода
select
  tp.product_alc_code as 'alc_code',
  coalesce((SELECT TOP 1 CAST(product_alc_volume AS float) FROM form1_egais WHERE product_inform_f1_reg_id = tp.product_inform_f1_reg_id ), 0.0) as 'alc_volume',
  tp.product_inform_f1_reg_id as 'alc_volume_f1',
  tt.consignee_region_code AS 'region_code',
  tt.consignee_client_reg_id AS 'fsrarid',
  cast(sum(case tp.product_unit_type when N'Packed' then COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) else 0 end) as int) as 'counts',
  round(sum(case tp.product_unit_type when N'Unpacked' then COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) else 0 end),4) as 'dal',
  round(sum(case tp.product_unit_type when N'Unpacked' then COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) else  COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) * cast(tp.product_capacity as float) * 0.1 end),4) as 'volume'
FROM ttn tt join ttn_products tp on tp.id_ttn = tt.id
  join ttn_form2 tf on tf.id_ttn = tt.id
  left join (select id, id_ttn, act_type, act_date from ttn_acts where id in (select max(ID) from ttn_acts where act_date >= '{{.Start}}' and act_date <= '{{.End}}' group by id_ttn)) takt on takt.id_ttn = tt.id
  left join (select id, id_ttn, act_type from ttn_acts where id in (select max(ID) from ttn_acts where act_date >= '{{.Start}}' and act_date <= '{{.End}}' group by id_ttn)) ta on ta.id_ttn = tt.id and ta.act_type = 'Расхождение'
  --left join ttn_acts_tickets tat on tat.id_ttn_acts = ta.id
  left join ttn_acts_tickets tat on tat.id_ttn_acts = ta.id and tat.ticket_date >= '{{.Start}}' and tat.ticket_date <= '{{.End}}'
  left join ttn_acts_content tac on tac.id_ttn_acts = ta.id and tac.product_iddomain = tp.product_iddomain
where tt.ttn_type in ('Исходящий', 'Импорт')
  and tt.doc_type = N'Возврат от меня'
  and tf.doc_fix_date >= '{{.Start}}'
  and tf.doc_fix_date <= '{{.End}}'
  and (takt.act_type <> 'Отказ' or takt.act_type is null)
  and (tat.ticket_type <> 'Отказ' or tat.ticket_type is null)
group by tt.consignee_region_code, tt.consignee_client_reg_id, tp.product_alc_code, tp.product_inform_f1_reg_id
;
