package home

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// путь к файлам шаблонов модуля для локальной отладки только
const defaultSrc = `E:\src\goproj\wails\firstwails\pages\home\templates`
const defaultTemplateName = "page"

type page struct {
	src          string
	template     *template.Template
	logger       *zap.SugaredLogger
	templateName string
}

// шаблоны парсятся однажды
// index имя шаблона для страницы по умолчанию (обычно это 'page')
// src путь до файлов шаблона
func NewPage(logger *zap.SugaredLogger, templ string, src string) *page {
	t := &page{
		logger:       logger,
		src:          src,
		templateName: templ,
	}
	if templ == "" {
		t.templateName = defaultTemplateName
	}
	tt := template.New("")
	template.Must(tt.New("style").Parse(styleTmpl))
	template.Must(tt.New("page").Funcs(funcMapHtml).Parse(indexTmpl))
	t.template = tt
	return t
}

// шаблоны парсятся каждый раз
// index имя шаблона по умолчанию
// src путь до файлов шаблона
func NewOnDemand(logger *zap.SugaredLogger, index string, src string) *page {
	t := &page{
		logger:       logger,
		src:          src,
		templateName: index,
	}
	if index == "" {
		t.templateName = defaultTemplateName
	}
	if src == "" {
		t.src = defaultSrc
	}
	return t
}

// такой вариант парсит при создании
func (t *page) Render(w io.Writer, templateName string, data interface{}, c echo.Context) error {
	if templateName == "" {
		templateName = defaultTemplateName
	}
	return t.template.ExecuteTemplate(w, templateName, data)
}

// такой вариант на ходу парсит шаблоны для отладки удобней
func (t *page) DoRender(w io.Writer, templateName string, data interface{}, c echo.Context) error {
	if templateName == "" {
		templateName = defaultTemplateName
	}
	tt := template.New("")
	if _, err := tt.New("style").Parse(fileGetContents(t.src + "\\style.html")); err != nil {
		t.logger.Error(err)
	}
	if _, err := tt.New("page").Funcs(funcMapHtml).Parse(fileGetContents(t.src + "\\index.html")); err != nil {
		t.logger.Error(err)
	}
	return tt.ExecuteTemplate(w, templateName, data)
}
