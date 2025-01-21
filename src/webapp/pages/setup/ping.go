package setup

import (
	"bytes"

	"github.com/labstack/echo/v4"
)

func (t *page) Ping(c echo.Context) error {
	// запускаем клиента чз
	tc := t.StartTrueClient(t.Reductor().Model())
	// проверяем что вышло
	if len(tc.Errors()) > 0 {
		t.Logger().Errorf("%s %s", modError, tc.Errors()[0])
	}
	var buf bytes.Buffer
	model := t.Reductor().Model().TrueClient
	model.PingSuz = tc.PingSuzInfo()
	if err := t.Render(&buf, "page", &model, c); err != nil {
		t.Logger().Errorf("%s %s", modError, err.Error())
		c.NoContent(204)
		return nil
	}
	c.HTML(200, buf.String())
	return nil
}
