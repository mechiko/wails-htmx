package a3mssql

import (
	_ "embed"
)

//go:embed sql/history_full.sql
var historyFullSql string

//go:embed sql/history_charge_on.sql
var historyChargeOnSql string

//go:embed sql/history_import.sql
var historyImportSql string

//go:embed sql/history_incoming.sql
var historyIncomingSql string

//go:embed sql/history_outgoing.sql
var historyOutgoingSql string

//go:embed sql/history_production_resource.sql
var historyProductionResourceSql string

//go:embed sql/history_production.sql
var historyProductionSql string

//go:embed sql/history_writeoff.sql
var historyWriteOffSql string

//go:embed sql/history_oborot_full.sql
var historyOborotSql string
