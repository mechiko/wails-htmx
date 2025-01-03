package pages

import (
	"firstwails/domain"
	"fmt"
	"io"

	"github.com/labstack/echo/v4"
)

const modError = "pages"

// type IApp interface {
// 	Logger() *zap.SugaredLogger
// 	Config() domain.IConfig
// 	Configuration() *domain.Configuration
// 	Reductor() domain.Reductor
// 	Effects() domain.Effects
// 	Repo() domain.Repo
// }

type Pages struct {
	domain.IApp
	registered map[string]domain.RenderHandler
	echo       *echo.Echo
}

func New(app domain.IApp, e *echo.Echo) *Pages {
	return &Pages{
		IApp:       app,
		registered: make(map[string]domain.RenderHandler),
		echo:       e,
	}
}

func (pgs *Pages) AddPage(name string, f domain.RenderHandler) {
	pgs.registered[name] = f
}

// рендер страницы по зарегистрированному имени name страницы
func (pgs *Pages) RenderPage(w io.Writer, name string, data interface{}, c echo.Context) error {
	if _, ok := pgs.registered[name]; ok {
		// передаем templateName пустой для отображения дефолтового шаблона страницы
		model := pgs.Reductor().Model()
		return pgs.registered[name](w, "", &model, c)
	}
	return fmt.Errorf("%s not registered", name)
}
