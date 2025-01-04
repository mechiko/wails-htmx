package webapp

import (
	"github.com/labstack/echo/v4"
)

func (a *webapp) Route() {
	a.echo.GET("/page", a.CurrentPageIndex)
	a.echo.GET("/sse", a.HandleSSE)
}

func (a *webapp) HandleSSE(c echo.Context) error { // longer variant with disconnect logic
	a.Logger().Debugf("The client is connected: %v", c.RealIP())
	go func() {
		<-c.Request().Context().Done() // Received Browser Disconnection
		a.Logger().Debugf("The client is disconnected: %v", c.RealIP())
	}()
	a.sse.ServeHTTP(c.Response(), c.Request())
	return nil
}
