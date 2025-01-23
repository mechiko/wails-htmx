package setup

import (
	"bytes"
	"firstwails/domain"

	"github.com/labstack/echo/v4"
)

func (t *page) Ping(c echo.Context) error {
	var pingError error
	pingSuzInfo := &domain.PingSuzInfo{}
	// запускаем клиента чз
	tc := t.StartTrueClient(t.Reductor().Model())
	// проверяем что вышло
	if len(tc.Errors()) > 0 {
		t.Logger().Errorf("%s %s", modError, tc.Errors()[0])
	} else {
		if pingSuzInfo, pingError = tc.PingSuz(); pingError != nil {
			t.Logger().Errorf("%s %s", modError, pingError.Error())
		}
	}
	var buf bytes.Buffer
	model := t.Reductor().Model().TrueClient
	model.PingSuz = pingSuzInfo
	model.Errors = append(model.Errors, tc.Errors()...)
	if pingError != nil {
		model.Errors = append(model.Errors, pingError.Error())
	}
	if err := t.Render(&buf, "page", &model, c); err != nil {
		t.Logger().Errorf("%s %s", modError, err.Error())
		c.NoContent(204)
		return nil
	}
	c.HTML(200, buf.String())
	return nil
}
