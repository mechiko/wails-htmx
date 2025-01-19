package setup

import (
	"firstwails/domain"
	"firstwails/utility"
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

const modError = "home"

// путь к файлам шаблонов модуля для локальной отладки только
// для пути используется в utility этого пакета
// путь filepath.Join(domain.RootPathTemplateDebug, defaultSrc, filename)
const defaultSrc = `setup\templates`

const defaultTemplateName = "index"

type page struct {
	domain.IApp
	src          string
	template     *template.Template
	templateName string
	reloadPage   bool
}

func New(app domain.IApp) *page {
	t := &page{
		IApp:         app,
		templateName: defaultTemplateName,
		src:          defaultSrc,
	}
	if domain.Mode != "development" {
		t.Production()
	}
	return t
}

// шаблоны парсятся однажды в режиме продакшн
func (t *page) Production() {
	tt := template.New("")
	template.Must(tt.New("index").Funcs(funcMapHtml).Parse(indexTmpl))
	template.Must(tt.New("page").Funcs(funcMapHtml).Parse(pageTmpl))
	template.Must(tt.New("pagesave").Funcs(funcMapHtml).Parse(pageSaveTmpl))
	template.Must(tt.New("omsid").Funcs(funcMapHtml).Parse(omsIdTmpl))
	template.Must(tt.New("deviceid").Funcs(funcMapHtml).Parse(deviceIdTmpl))
	template.Must(tt.New("configdb").Funcs(funcMapHtml).Parse(configDbTmpl))
	template.Must(tt.New("modal").Funcs(funcMapHtml).Parse(modalTmpl))
	t.template = tt
}

// такой вариант парсит при создании
func (t *page) Render(w io.Writer, templateName string, data interface{}, c echo.Context) error {
	if templateName == "" {
		templateName = defaultTemplateName
	}
	if domain.Mode == "development" {
		return t.DoRender(w, templateName, data, c)
	} else {
		if t.template == nil {
			return fmt.Errorf("%s template is nil", modError)
		}
		err := t.template.ExecuteTemplate(w, templateName, data)
		return err
	}
}

// такой вариант на ходу парсит шаблоны для отладки
func (t *page) DoRender(w io.Writer, templateName string, data interface{}, c echo.Context) (err error) {
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
	template.Must(tt.New("index").Funcs(funcMapHtml).Parse(utility.FilePackageContents("index.html")))
	template.Must(tt.New("page").Funcs(funcMapHtml).Parse(utility.FilePackageContents("page.html")))
	template.Must(tt.New("pagesave").Funcs(funcMapHtml).Parse(utility.FilePackageContents("pagesave.html")))
	template.Must(tt.New("modal").Funcs(funcMapHtml).Parse(utility.FilePackageContents("modal.html")))
	template.Must(tt.New("omsid").Funcs(funcMapHtml).Parse(utility.FilePackageContents("omsid.html")))
	template.Must(tt.New("deviceid").Funcs(funcMapHtml).Parse(utility.FilePackageContents("deviceid.html")))
	template.Must(tt.New("configdb").Funcs(funcMapHtml).Parse(utility.FilePackageContents("configdb.html")))
	return tt.ExecuteTemplate(w, templateName, data)
}

func (t *page) SetReload(b bool) {
	t.reloadPage = b
}
