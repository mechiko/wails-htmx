package config

// конфигурация на случай сбоя инициализации приложения
// или как пример конфига

var tomlConfig2 = []byte(`
# This is a TOML document.

`)

type DefaultConfiguration struct {
	Name     string                `json:"name"`
	Hostname string                `json:"hostname"`
	HostPort string                `json:"hostport"`
	Debug    bool                  `json:"debug"`
	IsAdmin  bool                  `json:"isadmin"`
	Browser  string                `json:"browser"`
	Database databaseConfiguration `json:"database"`
	Gui      guiGonfiguration      `json:"gui"`
}

type guiGonfiguration struct {
	Tray                 bool   `json:"tray"`
	LimitForEmptyFilter  int    `json:"limitforemptyfilter"`
	NoUsePeriodForFilter bool   `json:"nouseperiodforfilter"`
	Output               string `json:"output"`
	HideMainWindow       bool   `json:"hidemainwindow"`
	Export               string `json:"xlsx"`
}

type databaseConfiguration struct {
	ConnectionUri string `json:"connectionuri"`
	Driver        string `json:"driver"`
	DbName        string `json:"dbname"`
	Timeout       int    `json:"timeout"`
}
