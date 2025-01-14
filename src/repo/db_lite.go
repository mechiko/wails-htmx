package repo

import (
	"fmt"

	"firstwails/repo/lite/version"
)

// при инициализации приложения этот метод вызывается однажды и прописывает объект доступа
// к базе данных, далее проверяет версию БД возможна ошибка и нужно выходить из приложения
func (r *repository) lite() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("repo:lite panic %v", r)
		}
	}()
	return version.CheckVersionDb(r)
}
