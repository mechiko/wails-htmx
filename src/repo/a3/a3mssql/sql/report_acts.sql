-- акты
select
  tp.product_alc_code as 'alc_code',
  -- COALESCE(fe.product_alc_volume,concat('missing ',tp.product_inform_f1_reg_id)) as 'alc_volume',
  COALESCE(fe.product_alc_volume,tp.product_inform_f1_reg_id) as 'alc_volume',
  cast(sum(case tp.product_unit_type when N'Packed' then cast(tp.product_quantity as float) else 0 end) as int) as 'counts',
  round(sum(case tp.product_unit_type when N'Unpacked' then cast(tp.product_quantity as float) else 0 end),4) as 'dal',
  round(sum(case tp.product_unit_type when N'Unpacked' then cast(tp.product_quantity as float) else cast(tp.product_quantity as float) * cast(tp.product_capacity as float) * 0.1 end),4) as 'volume'
FROM write_off_acts tt join write_off_products tp on tp.id_write_off_acts = tt.id
left join form1_egais fe on fe.product_inform_f1_reg_id = tp.product_inform_f1_reg_id 
where 
  tt.archive = 0 
  and tt.status in (N'Проведён')  
  and tt.doc_date >= '{{.Start}}'
  and tt.doc_date <= '{{.End}}'
group by tp.product_alc_code, COALESCE(fe.product_alc_volume,tp.product_inform_f1_reg_id)
;
