package domain

// пакет состояния всей программы
// интерфейс Reader для доступа только чтения
type Model struct {
	Error string
	Gui   gui
	Menu  menu
	// page models
	DbInfo     dbInfo
	Stats      stats
	TrueClient trueClient
}

// формат сообщения
type Message struct {
	Sender string
	Cmd    string
	Model  *Model
}
