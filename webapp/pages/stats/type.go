package stats

import (
	"firstwails/domain"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// путь к файлам шаблонов модуля для локальной отладки только
// const defaultSrc = `C:\!src\wails-htmx\webapp\pages\stats\templates`
const defaultSrc = `E:\src\goproj\!!firstwails\webapp\pages\stats\templates`

const defaultTemplateName = "index"

// type IApp interface {
// 	Logger() *zap.SugaredLogger
// 	Config() domain.IConfig
// 	Configuration() *domain.Configuration
// 	Reductor() domain.Reductor
// 	Effects() domain.Effects
// 	Repo() domain.Repo
// 	SetActivePage(page string)
// }

type page struct {
	domain.IApp
	src          string
	template     *template.Template
	templateName string
}

// шаблоны парсятся однажды
// index имя шаблона для страницы по умолчанию (обычно это 'page')
// src путь до файлов шаблона
func NewPage(app domain.IApp, templ string, src string) *page {
	t := &page{
		IApp:         app,
		src:          src,
		templateName: templ,
	}
	if templ == "" {
		t.templateName = defaultTemplateName
	}
	tt := template.New("")
	template.Must(tt.New("style").Parse(styleTmpl))
	template.Must(tt.New("index").Funcs(funcMapHtml).Parse(indexTmpl))
	template.Must(tt.New("navbar").Funcs(funcMapHtml).Parse(navbarTmpl))
	template.Must(tt.New("page").Funcs(funcMapHtml).Parse(pageTmpl))
	template.Must(tt.New("footer").Funcs(funcMapHtml).Parse(footerTmpl))
	t.template = tt
	return t
}

// шаблоны парсятся каждый раз
// index имя шаблона по умолчанию
// src путь до файлов шаблна
func NewOnDemand(app domain.IApp, index string, src string) *page {
	t := &page{
		IApp:         app,
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
		t.Logger().Error(err)
	}
	if _, err := tt.New("index").Funcs(funcMapHtml).Parse(fileGetContents(t.src + "\\index.html")); err != nil {
		t.Logger().Error(err)
	}
	if _, err := tt.New("navbar").Funcs(funcMapHtml).Parse(fileGetContents(t.src + "\\navbar.html")); err != nil {
		t.Logger().Error(err)
	}
	if _, err := tt.New("page").Funcs(funcMapHtml).Parse(fileGetContents(t.src + "\\page.html")); err != nil {
		t.Logger().Error(err)
	}
	if _, err := tt.New("footer").Funcs(funcMapHtml).Parse(fileGetContents(t.src + "\\footer.html")); err != nil {
		t.Logger().Error(err)
	}
	return tt.ExecuteTemplate(w, templateName, data)
}

func (t *page) Model(model domain.Model) domain.Model {
	return model
}
