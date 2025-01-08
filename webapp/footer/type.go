package footer

import (
	"firstwails/domain"
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

const modError = "home"

// путь к файлам шаблонов модуля для локальной отладки только
// const defaultSrc = `C:\!src\wails-htmx\webapp\pages\home\templates`
const defaultSrc = `E:\src\goproj\!!firstwails\webapp\header\templates`

const defaultTemplateName = "content"

type page struct {
	domain.IApp
	template     *template.Template
	templateName string
}

func New(app domain.IApp) *page {
	t := &page{
		IApp:         app,
		templateName: defaultTemplateName,
	}
	if domain.Mode != "development" {
		t.Production()
	}
	return t
}

// шаблоны парсятся однажды в режиме продакшн
func (t *page) Production() {
	tt := template.New("")
	template.Must(tt.New("content").Funcs(funcMapHtml).Parse(contentTmpl))
	template.Must(tt.New("modal").Funcs(funcMapHtml).Parse(modalTmpl))
	t.template = tt
}

// шаблоны парсятся каждый раз
func (t *page) Development() {
}

// такой вариант парсит при создании
func (t *page) Render(w io.Writer, templateName string, data interface{}, c echo.Context) error {
	if templateName == "" {
		templateName = defaultTemplateName
	}
	if domain.Mode == "development" {
		return t.doRender(w, templateName, data, c)
	} else {
		if t.template == nil {
			return fmt.Errorf("%s template is nil", modError)
		}
		err := t.template.ExecuteTemplate(w, templateName, data)
		return err
	}
}

// такой вариант на ходу парсит шаблоны для отладки
func (t *page) doRender(w io.Writer, templateName string, data interface{}, c echo.Context) (err error) {
	defer func() {
		// ловим панику тут
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic DoRender %v", modError, r)
		}
	}()
	if templateName == "" {
		templateName = defaultTemplateName
	}
	tt := template.New("")
	// ошибки вызовут panic()
	template.Must(tt.New("content").Funcs(funcMapHtml).Parse(fileGetContents(defaultSrc + "\\content.html")))
	template.Must(tt.New("modal").Funcs(funcMapHtml).Parse(fileGetContents(defaultSrc + "\\modal.html")))
	return tt.ExecuteTemplate(w, templateName, data)
}

func (t *page) Model(model domain.Model) domain.Model {
	return model
}
