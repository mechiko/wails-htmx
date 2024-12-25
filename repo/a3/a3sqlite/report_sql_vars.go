package a3sqlite

import (
	_ "embed"
)

//go:embed sql/report_acts.sql
var reportActSql string

//go:embed sql/report_incoming_return.sql
var reportIncomingReturnSql string

//go:embed sql/report_incoming.sql
var reportIncomingSql string

//go:embed sql/report_outgoing_return.sql
var reportOutgoingReturnSql string

//go:embed sql/report_outgoing.sql
var reportOutgoingSql string

//go:embed sql/report_produce_product.sql
var reportProductSql string

//go:embed sql/report_produce_resource.sql
var reportResorceSql string
