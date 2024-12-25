-- ресурсы в производстве
select
  pp.resource_alc_code as 'alc_code',
  pp.resource_alc_volume as 'alc_volume',
  pp.producer_region_code AS 'region_code',
  pp.producer_client_reg_id AS 'fsrarid',
  total(case pp.resource_unit_type when "Packed" then pp.resource_quantity else 0 end) as 'counts',
  printf('%0.3f', total(case pp.resource_unit_type when "Unpacked" then pp.resource_quantity else 0 end)) as 'dal',
  printf('%0.3f', total(case pp.resource_unit_type when 'Unpacked' then pp.resource_quantity else pp.resource_quantity * pp.resource_capacity * 0.1 end)) as 'volume'
FROM production_reports tt join production_resources pp on pp.id_production_reports = tt.id
where 
  tt.status in ('Проведён')  
  and tt.doc_date >= '{{.Start}}'
  and tt.doc_date <= '{{.End}}'
group by pp.resource_alc_code, pp.resource_alc_volume
;
