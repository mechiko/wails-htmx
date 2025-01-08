// Package v1 implements routing paths. Each services in own file.
package router

import (
	_ "embed"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"firstwails/domain"
)

const modError = "http:router"

type V1Routes struct {
	app entity.IApp
	su  entity.ServiceUseCase
}

func New(handler *echo.Echo, app entity.IApp) *V1Routes {
	r := &V1Routes{
		app: app,
		su:  services.New(app),
	}
	handler.GET("/healthz", func(c echo.Context) error {
		fmt.Println("GET(/healthz, func(c echo.Context)")
		return c.String(200, "ok")
	})
	handler.Pre(middleware.RemoveTrailingSlash())
	// Routers
	// http://localhost:3600/v1/name
	h := handler.Group("/v1")
	{
		h.GET("/home", r.Home)
		h.GET("/importttn/exam", r.ExamImportTTN)
		h.GET("/importttn/import", r.ImportImportTTN)
		h.GET("/importttn/check", r.CheckImportTTN)
		h.GET("/tranfree", r.TransactionsReportFreeHtml)
		h.GET("/tranfree/:start", r.TransactionsReportFreeHtml)
		h.GET("/tranfree/:start/:end", r.TransactionsReportFreeHtml)
		h.GET("/transactions/excel", r.ExcelTransactionReportFile)
		h.GET("/localrest", r.RestLocalAdGrid)
		h.GET("/localrest/chartaptype", r.RestLocalChartApType)
		h.GET("/history", r.HistoryAdGrid)
		h.GET("/history/writeoff", r.HistoryWriteOff)
		h.GET("/history/production", r.HistoryProduction)
		h.GET("/history/parts", r.HistoryParts)
		h.GET("/history/full", r.HistoryFull)
		h.GET("/maintain/adminreport", r.AdminReport)
		h.GET("/maintain/adminreport/clear", r.AdminReportClear)
		h.GET("/history/oborot", r.HistoryOborot)
		h.GET("/history/oborot/excel", r.ExcelHistoryOborotFile)
		h.GET("/partner/oborot/:fsrarid", r.PartnerOborot)
		h.GET("/partner/oborot/excel/:fsrarid", r.PartnerOborotExcel)
		h.GET("/declaration/xml/orgs/:fsrarid", r.XmlOrgsPartner)
		h.GET("/declaration/xml/oborot/:fsrarid", r.XmlOborotPartner)
	}

	return r
}
