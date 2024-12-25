package reductor

import (
	"context"
)

func (rdc *reductor) Run(ctx context.Context) error {
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
