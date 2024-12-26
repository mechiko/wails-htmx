package domain

var TomlConfig = []byte(`
# This is a TOML document.
name=""
hostname = "localhost"
saveconfig = false
utmhost = "localhost"
utmport = "8080"
browser = ""
output = "output"
input = ''

[mssql]
driver = "mssql"
connection = "server=localhost;encrypt=disable;user id=;port=1433"

[sqlite]
driver = "sqlite"
# connectionuri = "?mode=rwc&_journal_mode=WAL"

[layouts]
timelayout = "2006-01-02T15:04:05-0700"
timelayoutclear = "2006.01.02 15:04:05"
timelayoutday = "2006.01.02"
timelayoututc = "2006-01-02T15:04:05"

[application]
license = "f7bc886d-bbcd-4ce9-845f-1209d87d406d"
console = false
fsrarid = ""
dbtype = "sqlite"
scantimer = 30
startpage = "home"

# конфигурации для баз данных
# пустые имена файлов означают попытку поиска бд
[config]
driver = "sqlite"

[alcohelp3]
driver = "sqlite"
dbname = ""

[trueznak]
driver = "sqlite"

[alcogo4]
driver = "sqlite"

`)

type Configuration struct {
	Name        string `json:"name"`
	Hostname    string `json:"hostname"`
	HostPort    string `json:"hostport"`
	UtmHost     string `json:"utmhost"`
	UtmPort     string `json:"utmport"`
	SaveConfig  bool   `json:"saveconfig"` // сохранять конфиг при запуске приложения если его нет отключено!
	Output      string
	Export      string
	Browser     string              `json:"browser"`
	Application appConfiguration    `json:"application"`
	Layouts     layoutConfiguration `json:"layouts"`
	// используется для хранения имен бд эти имена зашиты одни в конфиг.дб другие в entity.DbName для sqlite3
	Sqlite databaseConfiguration `json:"sqlite"`
	Mssql  databaseConfiguration `json:"mssql"`
	// описатели БД рефактор
	Config    databaseConfiguration `json:"config"`
	Alcohelp3 databaseConfiguration `json:"alcohelp3"`
	TrueZnak  databaseConfiguration `json:"trueznak"`
	Alcogo4   databaseConfiguration `json:"alcogo4"`
}

type layoutConfiguration struct {
	TimeLayout      string `json:"timelayout"`
	TimeLayoutClear string `json:"timelayoutclear"`
	TimeLayoutDay   string `json:"timelayoutday"`
	TimeLayoutUTC   string `json:"timelayoututc"`
}

type databaseConfiguration struct {
	Connection     string `json:"connection"`
	Driver         string `json:"driver"`
	AlcoHelpDbName string `json:"alcohelpdbname"`
	ConfigDbName   string `json:"configdbname"`
	ZnakDbName     string `json:"znakdbname"`
	ReportDbName   string `json:"reportdbname"`
	DbName         string `json:"dbname"`
	File           string `json:"file"`
	Timeout        int    `json:"timeout"`
	User           string `json:"user"`
	Pass           string `json:"pass"`
	Host           string `json:"host"`
	Port           string `json:"port"`
}

type appConfiguration struct {
	Pwd          string `json:"pwd"`
	Console      bool   `json:"console"`
	Disconnected bool   `json:"disconnected"`
	Fsrarid      string `json:"fsrarid"`
	DbType       string `json:"dbtype"`
	License      string `json:"license"`
	ScanTimer    int    `json:"scantimer"`
	StartPage    string `json:"startpage"`
}
