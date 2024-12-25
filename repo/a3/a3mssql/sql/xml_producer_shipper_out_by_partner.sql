-- все производители по исходящим ТТН и контрагенту за период
WITH producer_numbered AS (
  SELECT
	IIF(producer_type in ('UL','TS'),'Производитель',IIF(producer_type = 'FO','Импортер','')) as 'title',
    reg_id,producer_type AS org_type, inn, kpp,full_name, short_name, country_code, region_code, description,
    ROW_NUMBER() OVER(PARTITION BY reg_id ORDER BY row1 DESC) AS row_number
  FROM (
select DISTINCT ROW_NUMBER() OVER (ORDER BY (SELECT 1)) AS 'row1',
  COALESCE(fa.producer_client_reg_id, tp.producer_client_reg_id, '') as reg_id,
  COALESCE(fa.producer_type, tp.producer_type, '') as producer_type,
  COALESCE(fa.producer_inn, tp.producer_inn, '') as inn,
  COALESCE(fa.producer_kpp , tp.producer_kpp, '') as kpp,
  cast(COALESCE(fa.producer_full_name , tp.producer_full_name, '') as varchar(500)) as full_name,
  COALESCE(fa.producer_short_name , tp.producer_short_name, '') as short_name,
  COALESCE(fa.producer_country_code, tp.producer_country_code, '') as country_code,
  COALESCE(fa.producer_region_code , tp.producer_region_code, '') as region_code,
  cast(COALESCE(fa.producer_description , tp.producer_description, '') as varchar(500)) as description
FROM
  ttn tt join ttn_products tp on tt.id = tp.id_ttn
  left join form1_egais fa on fa.product_inform_f1_reg_id = tp.product_inform_f1_reg_id
where
tt.archive = 0 
and tt.status in ('Подтверждён','Проведён') 
and tt.ttn_type in ('Исходящий', 'Импорт') 
and tt.consignee_client_reg_id = '{{.Fsrar}}'
and tt.doc_date >= '{{.Start}}'
and tt.doc_date <= '{{.End}}'
) u1 ), shipper_numbered AS (
  SELECT 'Поставщик' as 'title', reg_id, shipper_type AS org_type, inn, kpp, full_name, short_name, country_code, region_code, description,
    ROW_NUMBER() OVER(PARTITION BY reg_id ORDER BY row2 DESC) AS row_number
  FROM (
select DISTINCT ROW_NUMBER() OVER (ORDER BY (SELECT 1)) AS 'row2',
  COALESCE(fb.shipper_client_reg_id, tt.shipper_client_reg_id) as reg_id,
  COALESCE(fb.shipper_type , tt.shipper_type) as shipper_type,
  COALESCE(fb.shipper_inn , tt.shipper_inn) as inn,
  COALESCE(fb.shipper_kpp , tt.shipper_kpp) as kpp,
  cast(COALESCE(fb.shipper_full_name, tt.shipper_full_name) as varchar(500)) as full_name,
  COALESCE(fb.shipper_short_name, tt.shipper_short_name) as short_name,
  COALESCE(fb.shipper_country_code, tt.shipper_country_code) as country_code,
  COALESCE(fb.shipper_region_code, tt.shipper_region_code) as region_code,
  cast(COALESCE(fb.shipper_description, tt.shipper_description) as varchar(500)) as description
FROM
  ttn tt join ttn_products tp on tt.id = tp.id_ttn
  left join form2_egais fb on fb.product_inform_f2_reg_id = tp.product_inform_f2_reg_id
where
tt.archive = 0 
and tt.status in ('Подтверждён','Проведён') 
and tt.ttn_type in ('Исходящий', 'Импорт') 
and tt.consignee_client_reg_id = '{{.Fsrar}}'
and tt.doc_date >= '{{.Start}}'
and tt.doc_date <= '{{.End}}'
) u2 )
SELECT stuff((select ',' + title from (select title,reg_id,org_type, inn, kpp, full_name, short_name, country_code, region_code, description from producer_numbered
where row_number = 1
union all
select title,reg_id,org_type, inn, kpp, full_name, short_name, country_code, region_code, description from shipper_numbered
where row_number = 1
) w3 for xml path ('')), 1, 1, '') as 'title',
  reg_id,org_type, inn, kpp, full_name, short_name, country_code, region_code, description
FROM (
select title,reg_id,org_type, inn, kpp, full_name, short_name, country_code, region_code, description from producer_numbered
where row_number = 1
union all
select title,reg_id,org_type, inn, kpp, full_name, short_name, country_code, region_code, description from shipper_numbered
where row_number = 1
) w2 group by reg_id,org_type, inn, kpp, full_name, short_name, country_code, region_code, description
;
