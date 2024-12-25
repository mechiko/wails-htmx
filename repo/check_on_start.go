package repo

import "fmt"

// эта проверка запускается в режиме r.App.Configuration().Debug
// видимо это будет устаревшее.. все проверки теперь internal\checkdbg\checks.go
func (r *repository) CheckOnStart() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	return nil
}
