package webapp

import (
	"context"
	"firstwails/domain"
	"firstwails/usecase"
)

// startup is called at application startup WAILS
func (a *webapp) Startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
	mdl := a.Reductor().Model()
	msg := domain.Message{
		Sender: "webapp.StartUp",
		Cmd:    "startup",
		Model:  &mdl,
	}
	a.Effects().ChanIn() <- msg
}

// вызывается из эффектора при запуске
// HTTP
func (a *webapp) StartUp() {
	// Perform your setup here
	mm := usecase.New(a).InitModel(a.Reductor().Model())
	msg := domain.Message{
		Sender: "webapp.StartUp",
		Cmd:    "startup",
		Model:  &mm,
	}
	a.Effects().ChanIn() <- msg
}
