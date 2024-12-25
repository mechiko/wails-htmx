-- ресурсы в производстве
select
  pp.resource_alc_code as 'alc_code',
  pp.resource_alc_volume as 'alc_volume',
  cast(sum(case pp.resource_unit_type when N'Packed' then cast(pp.resource_quantity as int) else 0 end) as int) as 'counts',
  cast(format(sum(case pp.resource_unit_type when N'Unpacked' then cast(pp.resource_quantity as float) else 0 end),'N') as float) as 'dal',
  cast(format(sum(case pp.resource_unit_type when N'Unpacked' then cast(pp.resource_quantity as float) else cast(pp.resource_quantity as float) * cast(pp.resource_capacity as float) * 0.1 end),'N') as float) as 'volume'
FROM production_reports tt join production_resources pp on pp.id_production_reports = tt.id
where 
  tt.status in ('Проведён')  
  and tt.doc_date >= '{{.Start}}'
  and tt.doc_date <= '{{.End}}'
group by pp.resource_alc_code, pp.resource_alc_volume
;
