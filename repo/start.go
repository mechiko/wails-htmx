package repo

import (
	"context"
	"fmt"
)

func (r *repository) Start(ctx context.Context) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	// пустая проверка пока закомментим на будущее может
	// if entity.Mode != "production" {
	// 	r.CheckOnStart()
	// }

	// ожидаем завершения контекста
	<-ctx.Done()
	r.Logger().Infof("завершаем работу репозитория по контексту")
	return nil
}
