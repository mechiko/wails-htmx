package webapp

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (a *webapp) Route() {
	a.echo.GET("/page", a.CurrentPageIndex)
	a.echo.GET("/ready", a.Ready)
	// a.echo.GET("/sse", a.HandleSSE)
}

// хэндлер для обработки запроса GET отображения текущей страницы s.activePage
func (s *webapp) CurrentPageIndex(c echo.Context) error {
	model := s.reductor.Model()
	return c.Render(http.StatusOK, s.activePage, &model)
}

// func (a *webapp) HandleSSE(c echo.Context) error { // longer variant with disconnect logic
// 	a.Logger().Debugf("The client is connected: %v", c.RealIP())
// 	go func() {
// 		<-c.Request().Context().Done() // Received Browser Disconnection
// 		a.Logger().Debugf("The client is disconnected: %v", c.RealIP())
// 	}()
// 	a.sse.ServeHTTP(c.Response(), c.Request())
// 	return nil
// }

func (a *webapp) Ready(c echo.Context) error {
	// без контента свап не производится
	a.DomReadyHttp()
	a.readyPingLastTime = time.Now()
	c.NoContent(204)
	return nil
}
