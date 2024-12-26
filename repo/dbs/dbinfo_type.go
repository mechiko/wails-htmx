package dbs

import (
	"firstwails/domain"
)

// name имя файла без пути
// file полное имя файла и путь
// name имя БД
// driver драйвер
// uri строка соединения
// absent отсутствует БД (надо делать пинг видимо)
type dbInfo struct {
	file       string
	name       string
	driver     string
	connection string
	mode       domain.UriMode
	checkMode  domain.CheckMode
	absent     bool
	dbService  domain.IDbService
}

var _ domain.DbInfo = &dbInfo{}
