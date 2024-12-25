-- здесь все лицензии БД
CREATE TABLE if not exists licenses (
  client_reg_id TEXT NOT NULL PRIMARY KEY,
  vid TEXT NOT NULL DEFAULT(''), -- Вид деятельности, указанный в лицензии текстом
  vid00 TEXT NOT NULL DEFAULT('04'), -- Вид деятельности для владельца БД тут 01 02 03 04 05 для декларации 33
  P000000000011 TEXT NOT NULL DEFAULT(''), -- Серия лицензии, номер
  P000000000012 TEXT NOT NULL DEFAULT(''), -- Дата выдачи
  P000000000013 TEXT NOT NULL DEFAULT(''), -- Дата окончания
  P000000000014 TEXT NOT NULL DEFAULT(''), -- Кем выдана
  is_retail     integer NOT NULL DEFAULT(0), -- розница
  is_wholesale  integer NOT NULL DEFAULT(0), -- опт
  is_producer   integer NOT NULL DEFAULT(0), -- производитель
  adr_type TEXT NOT NULL DEFAULT('') -- АдпТип xsd строка JSON
);
