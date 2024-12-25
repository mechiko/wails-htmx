package scrolling

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// текст для поиска приходит из страницы по запросу htmx
func (s *scrolling) InputHandler(c echo.Context) error {
	inputText := c.QueryParam("query")
	field := c.QueryParam("id")
	s.filter.SetByName(field, inputText)
	start := 0
	next := start + s.pageSize
	blocks := s.query(s.filter, s.pageSize, start)
	// массив фильтрованных пуст
	if len(blocks) == 0 {
		next = 0
		s.eof = true
		blocksData := ScrollingBlocks{
			Columns: s.columns,
			Start:   start,
			Next:    0,
			More:    false,
			Blocks:  blocks,
		}
		return c.Render(http.StatusOK, "blocks", blocksData)
	}

	template := "blocks"
	if start == 0 {
		template = "blocks-index"
	}
	blocksData := ScrollingBlocks{
		Columns: s.columns,
		Start:   start,
		Next:    next,
		More:    next > 0, // next контролируется и тут проверяем start плюс шаг
		Blocks:  blocks,
	}
	s.logger.Printf("template:%s start:%d next:%d mode:%v filtered:%d\n", template, blocksData.Start, blocksData.Next, blocksData.More, len(blocks))
	return c.Render(http.StatusOK, "blocks", blocksData)
}
