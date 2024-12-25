package scrolling

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// start приходит из страницы по запросу htmx и равен началу следующего блока
func (s *scrolling) BlocksHandler(c echo.Context) error {
	hx := HxRequestHeaderFromRequest(c.Request())
	if !hx.HxRequest {
		// сбрасываем фильтр если запрос индекса без htmx
		s.filter.Clear()
	}
	startStr := c.QueryParam("start")
	start, err := strconv.Atoi(startStr)
	if err != nil {
		start = 0
	}
	if s.eof {
		s.eof = true
		blocksData := ScrollingBlocks{
			Columns: s.columns,
			Start:   start,
			Next:    0,
			More:    false,
			Blocks:  nil,
		}
		return c.Render(http.StatusOK, "blocks", blocksData)
	}
	blocks := s.query(s.filter, s.pageSize, start)
	next := start + s.pageSize
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
	return c.Render(http.StatusOK, template, blocksData)
}
