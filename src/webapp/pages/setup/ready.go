package setup

import (
	"bytes"
	"firstwails/webapp/htmxutil"

	"github.com/labstack/echo/v4"
)

func (t *page) Ready(c echo.Context) error {
	if t.ReloadActivePage() {
		// c.Response().Header().Set("HX-Refresh", "true")
		t.SetReloadActivePage(false)
		var buf bytes.Buffer
		model := t.Reductor().Model()
		// model = usecase.New(t).DbInfoModel(model)
		if err := t.Render(&buf, "index", &model, c); err != nil {
			t.Logger().Errorf("%s %s", modError, err.Error())
			c.NoContent(204)
			return nil
		}
		c.HTML(200, buf.String())
		return nil
	}
	hx := htmxutil.HxRequestHeaderFromRequest(c.Request())
	if !hx.HxRequest {
		// готовность DOM когда не надо обновлять страницу и прилетел пинг от нее
		t.DomReadyHttp()
	}
	// без контента свап не производится
	c.NoContent(204)
	return nil
}
