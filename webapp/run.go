package webapp

import "context"

func (a *webapp) Run(ctx context.Context) error {
	ch := make(chan struct{})
	go func() {
		<-ctx.Done()
		a.logger.Info("webapp ctx done")
		ch <- struct{}{}
	}()
	go func() {
		a.reductor.Run(ctx)
	}()
	go func() {
		a.effects.Run(ctx)
	}()
	<-ch
	return nil
}
