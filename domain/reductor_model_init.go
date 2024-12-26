package domain

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
	Home: home{
		Export:  "export",
		Browser: "",
		Output:  "output",
	},
}
