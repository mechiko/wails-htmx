package pages

import (
	"firstwails/webapp/pages/dbinfo"
	"firstwails/webapp/pages/stats"
	"fmt"
)

// устанавливаем страницы по именам для перехода по URL
func (pgs *Pages) InitPages() error {
	// шаблон парсится при запуске
	// dbInfoPage := home.NewPage(pgs.logger, "", "")
	// шаблон парсится каждый раз при обращении
	dbInfoPage := dbinfo.New(pgs)
	// инициализируем маршруты страницы в АПИ
	if err := dbInfoPage.Route(pgs.echo); err != nil {
		return fmt.Errorf("%s InitPages %w", modError, err)
	}
	pgs.AddPage("dbinfo", dbInfoPage.Render)
	statPage := stats.NewOnDemand(pgs, "", "")
	if err := statPage.Route(pgs.echo); err != nil {
		return fmt.Errorf("%s InitPages %w", modError, err)
	}
	pgs.AddPage("stats", statPage.DoRender)
	return nil
}
