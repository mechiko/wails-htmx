-- ап в производстве
select
  pp.product_alc_code as 'alc_code',
  pp.product_alc_volume as 'alc_volume',
  cast(sum(case pp.product_unit_type when N'Packed' then cast(pp.product_quantity as float) else 0 end) as int) as 'counts',
  round(sum(case pp.product_unit_type when N'Unpacked' then cast(pp.product_quantity as float) else 0 end),4) as 'dal',
  round(sum(case pp.product_unit_type when N'Unpacked' then cast(pp.product_quantity as float) else cast(pp.product_quantity as float) * cast(pp.product_capacity as float) * 0.1 end),4) as 'volume'
FROM production_reports tt join production_products pp on pp.id_production_reports = tt.id
where 
  tt.status in (N'Проведён')  
  -- and tt.doc_date >= ?
  -- and tt.doc_date <= ?
  and tt.doc_date >= '{{.Start}}'
  and tt.doc_date <= '{{.End}}'
group by pp.product_alc_code , pp.product_alc_volume 
;
