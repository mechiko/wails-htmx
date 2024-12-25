-- акты
select
  pp.product_alc_code as 'alc_code',
  ifnull(fe.product_alc_volume,concat('missing ',pp.product_inform_f1_reg_id)) as 'alc_volume',
  -- pp.producer_region_code AS 'region_code',
  -- pp.producer_client_reg_id AS 'fsrarid',
  total(case pp.product_unit_type when "Packed" then pp.product_quantity else '' end) as 'counts',
  printf('%0.3f', total(case pp.product_unit_type when "Unpacked" then pp.product_quantity else '' end)) as 'dal',
  printf('%0.3f', total(case pp.product_unit_type when 'Unpacked' then pp.product_quantity else pp.product_quantity * pp.product_capacity * 0.1 end)) as 'volume'
FROM write_off_acts tt join write_off_products pp on pp.id_write_off_acts = tt.id
left join form1_egais fe on fe.product_inform_f1_reg_id = pp.product_inform_f1_reg_id 
where 
  tt.archive = 0 
  and tt.status in ('Проведён')  
  and tt.doc_date >= '{{.Start}}'
  and tt.doc_date <= '{{.End}}'
group by pp.product_alc_code, fe.product_alc_volume
;
