-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- CIS таблица для запроса кодов ЧЗ
CREATE TABLE if not exists cis_request (
	cis TEXT NOT NULL DEFAULT '',
  status TEXT NOT NULL DEFAULT '',
  responce BLOB,
  produced text NOT NULL DEFAULT '',
  expired  text NOT NULL DEFAULT '',
	PRIMARY KEY (cis)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
