package usecase

import "firstwails/domain"

func (u *usecase) MenuModel(model domain.Model) domain.Model {
	pages := u.PagesInfo()
	model.Menu.Items = make(domain.MenuItemSlice, 0)
	for _, page := range pages {
		menuItem := &domain.MenuItem{
			Name:   page.Name,
			Title:  page.MenuTitle,
			Active: u.ActivePage() == page.Name,
			Decs:   page.Desc,
			Svg:    page.Svg,
		}
		model.Menu.Items = append(model.Menu.Items, menuItem)
	}
	return model
}
