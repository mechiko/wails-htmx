package scrolling

import (
	"html/template"
	"io"
	"log"

	"github.com/labstack/echo/v4"
)

// путь к файлам шаблонов модуля для локальной отладки только
const defaultSrc = `scrolling\views`

type Templates struct {
	src      string
	template *template.Template
	logger   *log.Logger
}

// шаблоны парсятся однажды
func NewTemplate(logger *log.Logger, src string) *Templates {
	t := &Templates{
		src: src,
	}
	t.logger = logger
	if src == "" {
		t.src = defaultSrc
	}
	tt := template.New("")
	template.Must(tt.New("style").Parse(styleTmpl))
	template.Must(tt.New("blocks").Funcs(funcMapHtml).Parse(blocksTmpl))
	template.Must(tt.New("blocks-index").Funcs(funcMapHtml).Parse(blocksIndexTmpl))
	template.Must(tt.New("input-query").Funcs(funcMapHtml).Parse(inputQueryTmpl))
	template.Must(tt.New("select-columns").Funcs(funcMapHtml).Parse(selectColumns))
	t.template = tt
	return t
}

// шаблоны парсятся каждый раз
func NewOnDemand(src string) *Templates {
	t := &Templates{
		src: src,
	}
	return t
}

// такой вариант парсит при создании
func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.template.ExecuteTemplate(w, name, data)
}

// такой вариант на ходу парсит шаблоны для отладки удобней
func (t *Templates) DoRender(w io.Writer, name string, data interface{}, c echo.Context) error {
	tt := template.New("")
	if _, err := tt.New("style").Parse(fileGetContents(t.src + "\\style.html")); err != nil {
		t.logger.Printf(err.Error())
	}
	if _, err := tt.New("blocks").Funcs(funcMapHtml).Parse(fileGetContents(t.src + "\\blocks.html")); err != nil {
		t.logger.Printf(err.Error())
	}
	if _, err := tt.New("blocks-index").Funcs(funcMapHtml).Parse(fileGetContents(t.src + "\\blocks-index.html")); err != nil {
		t.logger.Printf(err.Error())
	}
	if _, err := tt.New("input-query").Funcs(funcMapHtml).Parse(fileGetContents(t.src + "\\input-query.html")); err != nil {
		t.logger.Printf(err.Error())
	}
	if _, err := tt.New("select-columns").Funcs(funcMapHtml).Parse(fileGetContents(t.src + "\\select-columns.html")); err != nil {
		t.logger.Printf(err.Error())
	}
	return tt.ExecuteTemplate(w, name, data)
}
