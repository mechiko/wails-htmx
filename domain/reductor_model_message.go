package domain

// пакет состояния всей программы
// интерфейс Reader для доступа только чтения
type Model struct {
	Gui  gui
	Home home
}

// формат сообщения
type Message struct {
	Sender string
	Cmd    string
	Model  *Model
}
