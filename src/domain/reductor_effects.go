package domain

import "context"

type FuncUpdaterGUI func(string, Model) (string, Model)

type Reductor interface {
	ChanIn() chan Message
	RegisterGui(FuncUpdaterGUI)
	Model() Model
	Run(ctx context.Context) error
}

type Effects interface {
	ChanIn() chan Message
	Run(ctx context.Context) error
}
