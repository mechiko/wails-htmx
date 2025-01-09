package webapp

import (
	"context"
	"firstwails/domain"
)

// startup is called at application startup WAILS
func (a *webapp) Startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
	a.logger.Debug("Startup!")
	msg := domain.Message{
		Sender: "webapp.Startup",
		Cmd:    "startup",
		Model:  nil,
	}
	a.Effects().ChanIn() <- msg
}

// вызывается из эффектора при запуске
// HTTP
func (a *webapp) StartUp() {
	// Perform your setup here
	a.logger.Debug("StartUp HTTP!")
	msg := domain.Message{
		Sender: "webapp.StartUp",
		Cmd:    "startup",
		Model:  nil,
	}
	a.Effects().ChanIn() <- msg
}
