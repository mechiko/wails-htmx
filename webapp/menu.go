package webapp

import (
	"github.com/wailsapp/wails/v2/pkg/menu"
)

func (a *webapp) ApplicationMenu() *menu.Menu {
	reportMenu := menu.SubMenu("Отчеты", menu.NewMenuFromItems(
		menu.Text("Статистика", nil, func(cbdata *menu.CallbackData) {
			a.SetActivePage("stats", true)
		}),
	))
	// homeMenu := &menu.MenuItem{
	// 	Label: "Домой",
	// 	Type:  menu.TextType,
	// 	Click: func(cbdata *menu.CallbackData) {
	// 		a.SetActivePage("stats")
	// 	},
	// }
	appMenu := menu.NewMenuFromItems(reportMenu)
	return appMenu
}
