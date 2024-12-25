select 
	pfc.product_inform_f1_reg_id as 'form1',
	--pp.product_party as 'party'
 	concat(pp.product_alc_code, ':', pp.product_party) as 'party'
from production_products pp 
  left join production_form1 pf on pf.id_production_reports = pp.id_production_reports 
  left join production_form1_content pfc on pfc.id_production_form1 = pf.id and pfc.product_identity = pp.product_identity 
where pfc.product_inform_f1_reg_id in (select product_inform_f1_reg_id from rests_ap_local ral)  
;
