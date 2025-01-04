package webapp

import (
	"context"
	"fmt"
)

func (a *webapp) Run(ctx context.Context) error {
	go func() {
		a.reductor.RegisterGui(a.ReductorUpdater)
		a.reductor.Run(ctx)
	}()
	go func() {
		a.effects.Run(ctx)
	}()
	port := a.Configuration().HostPort
	host := a.Configuration().Hostname
	proto := `http://`
	url := fmt.Sprintf("%s%s:%s", proto, host, port)
	a.Open(url)
	<-ctx.Done()
	return nil
}
