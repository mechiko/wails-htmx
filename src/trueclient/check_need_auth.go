package trueclient

import "time"

// возвращает true если авторизация необходима
// проверяется пингом СУЗ
func (t *trueClient) CheckNeedAuth() bool {
	if t.tokenGis == "" {
		return true
	}
	if t.tokenSuz == "" {
		return true
	}
	if t.authTime.IsZero() {
		return true
	}
	// если авторизация была больше 10 часов назад
	sinceTime := time.Since(t.authTime).Hours()
	if time.Duration(sinceTime) > 10 {
		return true
	}
	return false
}

// возвращает true если авторизация необходима
// проверяется пингом СУЗ
func (t *trueClient) CheckNeedAuthPing() bool {
	if t.tokenGis == "" {
		return true
	}
	if t.tokenSuz == "" {
		return true
	}
	if t.authTime.IsZero() {
		return true
	}
	// если авторизация была больше 10 часов назад
	sinceTime := time.Since(t.authTime).Hours()
	if time.Duration(sinceTime) > 10 {
		return true
	}
	if _, err := t.PingSuz(); err != nil {
		return true
	}
	return false
}
