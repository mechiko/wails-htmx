package domain

// создаем начальную модель для редуктора
// потом будет вызов usecase.InitModel()
var InitModel Model = Model{
	Error: make([]string, 0),
	Gui: gui{
		MainWindow: MainWindow{
			Title: "",
			StatusBar: StatusBar{
				Utm:     false,
				License: false,
				Scan:    false,
				Fsrarid: "",
			},
		},
	},
	DbInfo: dbInfo{
		Export:  "export",
		Browser: "",
		Output:  "output",
	},
	Menu: menu{
		Items: make(MenuItemSlice, 0),
	},
	Stats: stats{
		ExcelChunkSize: 30000,
	},
	TrueClient: TrueClient{},
}
