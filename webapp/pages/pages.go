package pages

import (
	"firstwails/webapp/pages/home"
	"firstwails/webapp/pages/stats"
	"fmt"
)

func (pgs *Pages) InitPages() error {
	// шаблон парсится при запуске
	// homePage := home.NewPage(pgs.logger, "", "")
	// шаблон парсится каждый раз при обращении
	homePage := home.NewOnDemand(pgs, "", "")
	if err := homePage.Route(pgs.echo); err != nil {
		return fmt.Errorf("%s InitPages %w", modError, err)
	}
	pgs.AddPage("home", homePage.DoRender)
	statPage := stats.NewOnDemand(pgs, "", "")
	if err := statPage.Route(pgs.echo); err != nil {
		return fmt.Errorf("%s InitPages %w", modError, err)
	}
	pgs.AddPage("stats", statPage.DoRender)
	return nil
}
