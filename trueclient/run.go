package trueclient

import (
	"context"
)

// errgroup
func (t *trueClient) Run(ctx context.Context) error {
	<-ctx.Done()
	return nil
}
