package domain

var InitModel Model = Model{
	Gui: gui{
		MainWindow: MainWindow{
			Width:  800,
			Height: 500,
			Title:  "Инструменты АлкоХелп 3",
			StatusBar: StatusBar{
				Utm:     false,
				License: false,
				Scan:    false,
				Fsrarid: "",
			},
		},
	},
	Home: app{
		Export:      "export",
		Browser:     "",
		Output:      "output",
		BrowserList: []string{"", "Chrome", "Firefox", "Yandex", "MSEdge"},
	},
}
