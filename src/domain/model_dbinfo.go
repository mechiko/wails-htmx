package domain

type dbInfo struct {
	Title       string
	Export      string
	Browser     string
	BrowserList []string
	Output      string
	Debug       bool
	UtmHost     string
	UtmPort     string
	DbLite      DbInfo
	DbConfig    DbInfo
	DbZnak      DbInfo
	DbA3        DbInfo
}
