CREATE TABLE if not exists alc_code_rest (
	alc_code TEXT NOT NULL DEFAULT '',
  alc_volume TEXT NOT NULL DEFAULT '',
	date_rest TEXT NOT NULL DEFAULT '',
  volume REAL NOT NULL DEFAULT 0,
	PRIMARY KEY (alc_code, alc_volume, date_rest)
);

-- таблица вид ап
CREATE TABLE if not exists vid_ap (
  vcode TEXT NOT NULL PRIMARY KEY,
  deleted INTEGER NOT NULL DEFAULT 0,
  name TEXT NOT NULL DEFAULT (''),
  typ TEXT NOT NULL DEFAULT (''),
  okved2 TEXT NOT NULL DEFAULT (''),
  other TEXT NOT NULL DEFAULT ('')
);

-- здесь лицензии для Контрагентов декларации 33 раздела Справочники
CREATE TABLE if not exists licenses (
  client_reg_id TEXT NOT NULL PRIMARY KEY,
  P0000000000111 TEXT NOT NULL DEFAULT(''), -- Вид деятельности, указанный в лицензии 01 02 03 04 05
  P000000000011 TEXT NOT NULL DEFAULT(''), -- Серия лицензии, номер
  P000000000012 TEXT NOT NULL DEFAULT(''), -- Дата выдачи
  P000000000013 TEXT NOT NULL DEFAULT(''), -- Дата окончания
  P000000000014 TEXT NOT NULL DEFAULT(''), -- Кем выдана
  is_retail     integer NOT NULL DEFAULT(0), -- розница
  is_wholesale  integer NOT NULL DEFAULT(0), -- опт
  is_producer   integer NOT NULL DEFAULT(0) -- производитель
);


-- <xs:element name="ВидДеят">
-- <xs:annotation>
-- <xs:documentation>Вид деятельности по лицензии 
-- 02 - Производство, хранение и поставки произведенного этилового спирта питьевого (относится к алкогольной продукции, для районов Крайнего Севера) 
-- 04 - Производство, хранение и поставки произведенных вин
-- </xs:documentation>
-- </xs:annotation>
-- <xs:simpleType>
-- <xs:restriction base="xs:string">
-- <xs:length value="2"/>
-- <xs:enumeration value="02"/>
-- <xs:enumeration value="04"/>
-- </xs:restriction>
-- </xs:simpleType>
-- </xs:element>
-- P0000000000111
-- 01 - Производство, хранение и поставки произведенного этилового спирта, в том числе денатурированного, 
-- 02 - Производство, хранение и поставки произведенного этилового спирта питьевого (относится к алкогольной продукции, для районов Крайнего Севера), 
-- 03 - Производство, хранение и поставки произведенных спиртных напитков, 
-- 04 - Производство, хранение и поставки произведенных вин, 
-- 05 - Производство, хранение и поставки произведенной спиртосодержащей пищевой продукции
CREATE TABLE if not exists licenses_vid (
  id TEXT NOT NULL PRIMARY KEY,
  description   TEXT NOT NULL DEFAULT(0)
);

CREATE TABLE if not exists country (
  code TEXT NOT NULL PRIMARY KEY,
  name TEXT NOT NULL DEFAULT ('')
);

