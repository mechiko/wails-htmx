-- CIS таблица для запроса кодов ЧЗ
-- INSERT INTO cis_request VALUES(jsonb('{"country":"Luxembourg","capital":"Luxembourg City","languages":["French","German","Luxembourgish"]}'));
CREATE TABLE if not exists cis_request (
	cis TEXT NOT NULL DEFAULT '',
  status TEXT NOT NULL DEFAULT '',
  responce BLOB,
	PRIMARY KEY (cis)
);
