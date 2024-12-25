package domain

// пакет состояния всей программы
// интерфейс Reader для доступа только чтения
type Model struct {
	Gui  gui
	Home app
}

// формат сообщения
type Message struct {
	Sender string
	Cmd    string
	Model  *Model
}
