package webapp

import "firstwails/domain"

func (a *webapp) PagesInfo() domain.ArrPageInfo {
	if a.pages == nil {
		return nil
	}
	return a.pages.Infos()
}
