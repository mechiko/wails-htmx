package webapp

import (
	"github.com/labstack/echo/v4"
)

func (a *webapp) customHTTPErrorHandler(err error, c echo.Context) {
	c.Echo().DefaultHTTPErrorHandler(err, c)
	a.Logger().Errorf("%s echo error %s", modError, err.Error())
}
