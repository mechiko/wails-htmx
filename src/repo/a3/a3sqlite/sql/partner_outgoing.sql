SELECT * 
FROM partners_egais pe 
WHERE ID in (select distinct max(ID) from partners_egais group by client_reg_id)
  AND pe.client_reg_id IN (SELECT DISTINCT tt.consignee_client_reg_id 
FROM ttn tt 
where 
  tt.status in ('Подтверждён','Проведён') 
  and tt.ttn_type in ('Исходящий', 'Импорт')
  and tt.doc_date >= '{{.Start}}'
  and tt.doc_date <= '{{.End}}'
)
;
