package pages

import (
	"firstwails/domain"
	"firstwails/webapp/pages/home"
	"fmt"
	"io"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

const modError = "pages"

type Pages struct {
	registered map[string]domain.RenderHandler
	logger     *zap.SugaredLogger
	echo       *echo.Echo
}

func New(logger *zap.SugaredLogger, e *echo.Echo) *Pages {
	return &Pages{
		logger:     logger,
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
		return pgs.registered[name](w, "", data, c)
	}
	return fmt.Errorf("%s not registered", name)
}

func (pgs *Pages) InitPages() error {
	// шаблон парсится при запуске
	// homePage := home.NewPage(pgs.logger, "", "")
	// шаблон парсится каждый раз при обращении
	homePage := home.NewOnDemand(pgs.logger, "", "")
	if err := homePage.Route(pgs.echo); err != nil {
		return fmt.Errorf("%s InitPages %w", modError, err)
	}
	pgs.AddPage("home", homePage.DoRender)
	return nil
}
