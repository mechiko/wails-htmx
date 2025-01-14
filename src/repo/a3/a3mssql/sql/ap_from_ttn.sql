SELECT DISTINCT 
coalesce(cast(product_full_name as varchar(500)), '') as 'product_full_name',
coalesce(cast(product_capacity as varchar(500)), '') as 'product_capacity',
coalesce(cast(product_alc_volume as varchar), '') as 'product_alc_volume',
coalesce(cast(product_alc_code as varchar), '') as 'product_alc_code',
coalesce(cast(product_code as varchar), '') as 'product_code',
coalesce(cast(product_unit_type as varchar), '') as 'product_unit_type',
coalesce(cast(producer_type as varchar), '') as 'producer_type',
coalesce(cast(producer_client_reg_id as varchar), '') as 'producer_client_reg_id',
coalesce(cast(producer_inn as varchar), '') as 'producer_inn',
coalesce(cast(producer_kpp as varchar), '') as 'producer_kpp',
coalesce(cast(producer_full_name as varchar(500)), '') as 'producer_full_name',
coalesce(cast(producer_short_name as varchar(500)), '') as 'producer_short_name',
coalesce(cast(producer_country_code as varchar), '') as 'producer_country_code',
coalesce(cast(producer_region_code as varchar), '') as 'producer_region_code',
coalesce(cast(producer_description as varchar(500)), '') as 'producer_description'
FROM ttn_products where id in (select max(id) from ttn_products tp where SUBSTRING(product_alc_code,1,2) <> '00' group by product_alc_code)
order by product_alc_code
;
