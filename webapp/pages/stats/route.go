package stats

import "github.com/labstack/echo/v4"

func (t *page) Route(e *echo.Echo) error {
	e.GET("/stats/home", t.Home)
	return nil
}

// переходим на страницу Home
func (t *page) Home(c echo.Context) error {
	t.SetActivePage("home", false)
	c.String(204, "")
	c.Response().Header().Add("HX-Refresh", "true")
	return nil
}
