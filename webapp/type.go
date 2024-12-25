package webapp

import (
	"context"
	"firstwails/domain"
	"firstwails/webapp/pages"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type webapp struct {
	ctx           context.Context
	uuid          string // идентификатор для уникальности формы
	pages         *pages.Pages
	activePage    string
	echo          *echo.Echo
	config        domain.IConfig
	logger        *zap.SugaredLogger
	configuration *domain.Configuration
	reductor      domain.Reductor
	effects       domain.Effects
}

const modError = "webapp"

func NewWebApp(logger *zap.SugaredLogger, e *echo.Echo) *webapp {
	sc := &webapp{}
	sc.logger = logger
	sc.uuid = uuid.New().String()
	if err := sc.initConfig(); err != nil {
		panic(fmt.Sprintf("%s NewWebApp() %s", modError, err.Error()))
	}

	sc.pages = pages.New(logger, e)
	sc.activePage = "home"
	sc.echo = e
	if err := sc.pages.InitPages(); err != nil {
		panic(fmt.Sprintf("%s NewWebApp() %s", modError, err.Error()))
	}
	return sc
}

// хэндлер для обработки запроса GET отображения текущей страницы s.activePage
// sc.pages мап страниц с функциями заполнения
// передаем имя по списку шаблонов зарегистрированных в pages и модель данных для отображения
func (s *webapp) CurrentPageIndex(c echo.Context) error {
	return c.Render(http.StatusOK, s.activePage, &struct{ param string }{param: "page"})
}

// интерфейс для echo заполнения ответа по роутингу
// вызывает для страницы name функцию рендеринга
// data модель отображения
func (s *webapp) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return s.pages.RenderPage(w, name, data, c)
}

// startup is called at application startup
func (a *webapp) Startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a webapp) DomReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *webapp) BeforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *webapp) Shutdown(ctx context.Context) {
	// Perform your teardown here
}
