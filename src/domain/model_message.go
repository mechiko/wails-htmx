package domain

// пакет состояния всей программы
// интерфейс Reader для доступа только чтения
type Model struct {
	Error    []string
	IsZnak   bool
	IsA3     bool
	IsConfig bool
	Gui      gui
	Menu     menu
	// page models
	DbInfo     dbInfo
	Stats      stats
	Gtins      gtins
	TrueClient TrueClient
	Finder     finder
}

// формат сообщения
type Message struct {
	Sender string
	Cmd    string
	Model  *Model
}
