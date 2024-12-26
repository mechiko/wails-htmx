package webapp

import "context"

func (a *webapp) Run(ctx context.Context) error {
	go func() {
		a.reductor.RegisterGui(a.ReductorUpdater)
		a.reductor.Run(ctx)
	}()
	go func() {
		a.effects.Run(ctx)
	}()
	<-ctx.Done()
	return nil
}
