package domain

// создаем начальную модель для редуктора
// потом будет вызов usecase.InitModel()
var InitModel Model = Model{
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
}
