package main

import (
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// applicationMenu creates a menu that we will use for the
// application
func (a *App) applicationMenu() *menu.Menu {

	// We can use the same callback for multiple menu items
	handleRadio := func(cbdata *menu.CallbackData) {
		// The menu item that was selected can be accessed through the
		// callback data
		println("radio checked: ", cbdata.MenuItem.Label)
	}

	// We can define menu items and reuse them.
	// Their state will stay in sync.
	checkbox := &menu.MenuItem{
		Label:   "Checked",
		Type:    menu.CheckboxType,
		Checked: true,
		Click: func(cbdata *menu.CallbackData) {
			fmt.Printf("checked = %v\n", cbdata.MenuItem.Checked)
		},
	}

	// Radio groups are automatically made by adjacent radio menu items
	radioMenu1 := menu.Radio("White", false, nil, handleRadio)
	radioMenu2 := menu.Radio("Green", true, nil, handleRadio)
	radioMenu3 := menu.Radio("Red", false, nil, handleRadio)

	secretMenu := menu.SubMenu("Secret Menu",
		menu.NewMenuFromItems(
			menu.Text("test", nil, nil),
			radioMenu1,
			radioMenu2,
			radioMenu3,
			checkbox,
		),
	)

	// Menu items can be hidden
	secretMenu.Hidden = true

	// Menus can be updated by callbacks. Updates will get reflected after
	// calling `runtime.Menu.UpdateApplicationMenu()`
	toggleSecretMenu := func(cbdata *menu.CallbackData) {
		secretMenu.Hidden = !secretMenu.Hidden
		if secretMenu.Hidden {
			cbdata.MenuItem.Label = "Show secret menu...😉"
		} else {
			cbdata.MenuItem.Label = "Hide secret menu! 😲"
		}
		runtime.MenuUpdateApplicationMenu(a.ctx)
	}

	return menu.NewMenuFromItems(
		menu.SubMenu("File", menu.NewMenuFromItems(
			menu.Text("&Open", keys.CmdOrCtrl("o"), func(cbdata *menu.CallbackData) {
				println("Open Called!")
				filename, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
					DefaultDirectory: "",
					DefaultFilename:  "",
					Title:            "Selct yer fil",
					Filters:          nil,
					// AllowFiles:                 false,
					// AllowDirectories:           false,
					ShowHiddenFiles:            false,
					CanCreateDirectories:       false,
					ResolvesAliases:            false,
					TreatPackagesAsDirectories: false,
				})
				if err != nil {
					return
				}
				_, err = runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
					Type:          "info",
					Title:         "Hello",
					Message:       filename,
					Buttons:       nil,
					DefaultButton: "",
					CancelButton:  "",
					Icon:          []byte{},
				})
				if err != nil {
					return
				}
			}),
			menu.Separator(),
			menu.SubMenu("Submenu",
				menu.NewMenuFromItems(
					&menu.MenuItem{
						Label:    "Disabled",
						Type:     menu.TextType,
						Disabled: true,
					},
					&menu.MenuItem{
						Label:  "Hidden",
						Hidden: true,
						Type:   menu.TextType,
					},
					checkbox,
					checkbox,
					checkbox,
					radioMenu1,
					radioMenu2,
					radioMenu3,
					menu.Separator(),
					radioMenu1,
					radioMenu2,
					radioMenu3,
					&menu.MenuItem{
						Label: "Hover over me",
						Type:  menu.TextType,
					},
					menu.Text("Text Menu Item", nil, nil),
				),
			),
			menu.Text("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
				runtime.Quit(a.ctx)
			}),
		)),
		menu.SubMenu("Tests", menu.NewMenuFromItems(
			menu.Text("Show secret menu...😉", nil, toggleSecretMenu),
			menu.Separator(),
			secretMenu,
		),
		),
	)

}
