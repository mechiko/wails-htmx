package dbinfo

import (
	"firstwails/domain"
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

const modError = "home"

// путь к файлам шаблонов модуля для локальной отладки только
const defaultSrc = `C:\!src\wails-htmx\webapp\pages\dbinfo\templates`

// const defaultSrc = `E:\src\goproj\!!firstwails\webapp\pages\dbinfo\templates`

const defaultTemplateName = "index"

// type IApp interface {
// 	Logger() *zap.SugaredLogger
// 	Config() domain.IConfig
// 	Configuration() *domain.Configuration
// 	Reductor() domain.Reductor
// 	Effects() domain.Effects
// 	Repo() domain.Repo
// }

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
	template.Must(tt.New("style").Parse(styleTmpl))
	template.Must(tt.New("index").Funcs(funcMapHtml).Parse(indexTmpl))
	template.Must(tt.New("indexerror").Funcs(funcMapHtml).Parse(indexErrorTmpl))
	template.Must(tt.New("navbar").Funcs(funcMapHtml).Parse(navbarTmpl))
	template.Must(tt.New("page").Funcs(funcMapHtml).Parse(pageTmpl))
	template.Must(tt.New("error").Funcs(funcMapHtml).Parse(errorTmpl))
	template.Must(tt.New("footer").Funcs(funcMapHtml).Parse(footerTmpl))
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
	template.Must(tt.New("style").Parse(fileGetContents(t.src + "\\style.html")))
	template.Must(tt.New("index").Funcs(funcMapHtml).Parse(fileGetContents(t.src + "\\index.html")))
	template.Must(tt.New("indexerror").Funcs(funcMapHtml).Parse(fileGetContents(t.src + "\\indexerror.html")))
	template.Must(tt.New("navbar").Funcs(funcMapHtml).Parse(fileGetContents(t.src + "\\navbar.html")))
	template.Must(tt.New("page").Funcs(funcMapHtml).Parse(fileGetContents(t.src + "\\page.html")))
	template.Must(tt.New("error").Funcs(funcMapHtml).Parse(fileGetContents(t.src + "\\error.html")))
	template.Must(tt.New("footer").Funcs(funcMapHtml).Parse(fileGetContents(t.src + "\\footer.html")))
	// if _, err := tt.New("style").Parse(fileGetContents(t.src + "\\style.html")); err != nil {
	// 	t.Logger().Error(err)
	// }
	// if _, err := tt.New("index").Funcs(funcMapHtml).Parse(fileGetContents(t.src + "\\index.html")); err != nil {
	// 	t.Logger().Error(err)
	// }
	// if _, err := tt.New("indexerror").Funcs(funcMapHtml).Parse(fileGetContents(t.src + "\\indexerror.html")); err != nil {
	// 	t.Logger().Error(err)
	// }
	// if _, err := tt.New("navbar").Funcs(funcMapHtml).Parse(fileGetContents(t.src + "\\navbar.html")); err != nil {
	// 	t.Logger().Error(err)
	// }
	// if _, err := tt.New("page").Funcs(funcMapHtml).Parse(fileGetContents(t.src + "\\page.html")); err != nil {
	// 	t.Logger().Error(err)
	// }
	// if _, err := tt.New("error").Funcs(funcMapHtml).Parse(fileGetContents(t.src + "\\error.html")); err != nil {
	// 	t.Logger().Error(err)
	// }
	// if _, err := tt.New("footer").Funcs(funcMapHtml).Parse(fileGetContents(t.src + "\\footer.html")); err != nil {
	// 	t.Logger().Error(err)
	// }
	// err = tt.ExecuteTemplate(w, templateName, data)
	return tt.ExecuteTemplate(w, templateName, data)
}

func (t *page) Model(model domain.Model) domain.Model {
	return model
}
