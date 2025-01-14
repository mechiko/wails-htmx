delete from requests WHERE status in ('Отправлен','Ошибка');
--delete from requests WHERE id in (SELECT id from requests WHERE status in ('Отправлен','Ошибка') LIMIT 3);
