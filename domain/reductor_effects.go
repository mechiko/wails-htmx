package domain

type FuncUpdaterGUI func(string, Model) (string, Model)

type Reductor interface {
	ChanIn() chan Message
	RegisterGui(FuncUpdaterGUI)
	Model() Model
}

type Effects interface {
	ChanIn() chan Message
}
