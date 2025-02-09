package views

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

type views struct {
	domain.IApp
	registered     map[string]domain.RenderHandler
	echo           *echo.Echo
	infoRegistered domain.ArrPageInfo
}

func New(app domain.IApp, e *echo.Echo) *views {
	return &views{
		IApp:           app,
		registered:     make(map[string]domain.RenderHandler),
		echo:           e,
		infoRegistered: make(domain.ArrPageInfo, 0),
	}
}

func (pgs *views) AddPage(info *domain.PageInfo, f domain.RenderHandler) {
	pgs.registered[info.Name] = f
	pgs.infoRegistered = append(pgs.infoRegistered, info)
}

// рендер страницы по зарегистрированному имени name страницы
func (pgs *views) RenderPage(w io.Writer, name string, data interface{}, c echo.Context) error {
	if page, ok := pgs.registered[name]; ok {
		// передаем templateName пустой для отображения дефолтового шаблона страницы
		model := pgs.Reductor().Model()
		err := page(w, "", &model, c)
		if err != nil {
			pgs.Logger().Errorf("page %s error %s", name, err.Error())
		}
		return err
	}
	pgs.Logger().Errorf("%s %s not registered", modError, name)
	return fmt.Errorf("%s %s not registered", modError, name)
}

// мап зарегистрированных страниц
func (pgs *views) Infos() domain.ArrPageInfo {
	return pgs.infoRegistered
}

func (pgs *views) Info(page string) *domain.PageInfo {
	for i, info := range pgs.infoRegistered {
		if info.Name == page {
			return pgs.infoRegistered[i]
		}
	}
	return nil
}
