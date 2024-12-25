select
    tt.doc_number as docnumber,
    concat(SUBSTRING(tt.doc_date, 9, 2), '.', SUBSTRING(tt.doc_date, 6, 2), '.', SUBSTRING(tt.doc_date, 1, 4)) as docdate,
    tp.product_full_name as product,
    tp.product_code as code,
    case tp.producer_type when 'FO' then 'Импортер' else 'Производитель' end as 'type',
    case tp.producer_type when 'FO' then tt.shipper_short_name else tp.producer_short_name end as short,
    case tp.producer_type when 'FO' then tt.shipper_inn else tp.producer_inn end as inn,
    case tp.producer_type when 'FO' then tt.shipper_kpp else tp.producer_kpp end as kpp,
    case tp.producer_type when 'FO' then tt.shipper_client_reg_id else tp.producer_client_reg_id end as regid,
	coalesce((SELECT TOP 1 CAST(product_alc_volume AS float) FROM form1_egais WHERE product_inform_f1_reg_id = tp.product_inform_f1_reg_id ), 0.0) as 'alc_volume_fact',
	coalesce((SELECT TOP 1 gtd_number FROM form1_egais WHERE product_inform_f1_reg_id = tp.product_inform_f1_reg_id ), '') as 'gtd',
    cast(case tp.product_unit_type 
        when 'Packed' then COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) 
	    else 0 end as int) as 'counts',
  case tp.product_unit_type 
    when 'Unpacked' then COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) 
    else 0 end as 'dal',
  round(case tp.product_unit_type 
    when 'Unpacked' then COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) 
    else  COALESCE(cast(tac.product_quantity as float),cast(tp.product_quantity as float),0) * cast(tp.product_capacity as float) * 0.1 end
    ,4) as 'volume'
FROM ttn tt join ttn_products tp on tp.id_ttn = tt.id
  join ttn_form2 tf on tf.id_ttn = tt.id
  left join (select id, id_ttn, act_type, act_date from ttn_acts where id in (select max(ID) from ttn_acts group by id_ttn)) takt on takt.id_ttn = tt.id
  left join (select id, id_ttn, act_type from ttn_acts where id in (select max(ID) from ttn_acts group by id_ttn)) ta on ta.id_ttn = tt.id and ta.act_type = 'Расхождение'
  left join ttn_acts_content tac on tac.id_ttn_acts = ta.id and tac.product_identity = tp.product_identity
  left join ttn_acts_tickets tat on tat.id_ttn_acts = ta.id
  left join form1_egais fe on fe.product_inform_f1_reg_id = tp.product_inform_f1_reg_id 
where tt.ttn_type in ('Исходящий', 'Импорт')
  and tt.consignee_client_reg_id = '{{.Fsrar}}'
  and tt.doc_type <> N'Возврат от меня'
  and tf.doc_fix_date >= '{{.Start}}'
  and tf.doc_fix_date <= '{{.End}}'
  and (takt.act_type <> 'Отказ' or takt.act_type is null)
  and (tat.ticket_type <> 'Отказ' or tat.ticket_type is null)
;
