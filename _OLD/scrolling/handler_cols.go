package scrolling

import (
	"github.com/labstack/echo/v4"
)

// текст для поиска приходит из страницы по запросу htmx
func (s *scrolling) ColsHandler(c echo.Context) error {
	cols := c.QueryParam("cols")
	id := c.QueryParam("id")
	if col := s.columns.Column(id); col != nil {
		if cols == "" {
			col.SetVisible(false)
		} else {
			col.SetVisible(true)
		}
	}
	s.logger.Printf("id = %s cols = %s\n", id, cols)
	return c.String(204, "")
}

func (s *scrolling) ColsSubmit(c echo.Context) error {
	s.logger.Println("ColsSubmit")
	c.Response().Header().Add("HX-Refresh", "true")
	return c.String(204, "")
}
