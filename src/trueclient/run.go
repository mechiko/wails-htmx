package trueclient

import (
	"context"
)

// errgroup
func (t *trueClient) Run(ctx context.Context) error {
	// при создании объекта будет авторизация или паника
	// значит объект был создан и сохраняем при запуске службы значения в конфиг
	t.Save()
	<-ctx.Done()
	return nil
}
