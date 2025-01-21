package effects

import (
	"context"
)

func (rdc *effects) Run(ctx context.Context) error {
	go func() {
		for s := range rdc.in {
			rdc.proccessMessage(s)
		}
	}()

	<-ctx.Done()
	rdc.logger.Infof("%s done", modError)
	close(rdc.in)
	return nil
}
