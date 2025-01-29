package pages

import (
	"firstwails/domain"
	"firstwails/webapp/pages/dbinfo"
	"firstwails/webapp/pages/gtins"
	"firstwails/webapp/pages/setup"
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
	infoDbInfoPage := &domain.PageInfo{
		Name:      "dbinfo",
		Url:       "dbinfo",
		MenuTitle: "Инфо",
		Desc:      "иноформация по бд",
		Svg:       dbInfoPage.Svg(),
	}
	pgs.AddPage(infoDbInfoPage, dbInfoPage.Render)

	statPage := stats.New(pgs)
	if err := statPage.Route(pgs.echo); err != nil {
		return fmt.Errorf("%s InitPages %w", modError, err)
	}
	infoStatPage := &domain.PageInfo{
		Name:      "stats",
		Url:       "stats",
		MenuTitle: "Статистика",
		Desc:      "статистика по КМ",
		Svg:       statPage.Svg(),
	}
	pgs.AddPage(infoStatPage, statPage.Render)

	// setup page
	setupPage := setup.New(pgs)
	// инициализируем маршруты страницы в АПИ
	if err := setupPage.Route(pgs.echo); err != nil {
		return fmt.Errorf("%s InitPages %w", modError, err)
	}
	infoSetupPage := &domain.PageInfo{
		Name:      "setup",
		Url:       "setup",
		MenuTitle: "Настройка",
		Desc:      "настройка ЧЗ",
		Svg:       setupPage.Svg(),
	}
	pgs.AddPage(infoSetupPage, setupPage.Render)

	gtinsPage := gtins.New(pgs)
	if err := gtinsPage.Route(pgs.echo); err != nil {
		return fmt.Errorf("%s InitPages %w", modError, err)
	}
	infoGtinsPage := &domain.PageInfo{
		Name:      "gtins",
		Url:       "gtins",
		MenuTitle: "Справочник GTIN",
		Desc:      "информация по GTIN",
		Svg:       gtinsPage.Svg(),
	}
	pgs.AddPage(infoGtinsPage, gtinsPage.Render)
	return nil
}
