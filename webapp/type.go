package webapp

import (
	"context"
	"firstwails/domain"
	"firstwails/webapp/effects"
	"firstwails/webapp/pages"
	"firstwails/webapp/reductor"
	"fmt"
	"io"
	"net/http"
	"time"

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
	pwd           string
	startTime     time.Time
	endTime       time.Time
	repo          domain.Repo
	readyDOM      bool
}

const modError = "webapp"

var _ domain.IApp = &webapp{}

func NewWebApp(logger *zap.SugaredLogger, e *echo.Echo, pwd string) *webapp {
	sc := &webapp{}
	sc.pwd = pwd
	sc.logger = logger
	sc.uuid = uuid.New().String()
	sc.configuration = &domain.Configuration{}
	if err := sc.initConfig(); err != nil {
		panic(fmt.Sprintf("%s NewWebApp() %s", modError, err.Error()))
	}
	sc.pages = pages.New(sc, e)
	sc.activePage = "home"
	sc.echo = e
	if err := sc.pages.InitPages(); err != nil {
		panic(fmt.Sprintf("%s NewWebApp() %s", modError, err.Error()))
	}
	// сначала эффекты они прописываются в редукторе
	sc.effects = effects.New(sc)
	sc.reductor = reductor.New(sc, sc.effects)
	sc.Route()
	sc.initDateMn()
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
	s.logger.Debugf("echo %s render %s", modError, name)
	return s.pages.RenderPage(w, name, data, c)
}

// startup is called at application startup
func (a *webapp) Startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
	a.logger.Debug("Startup!")
	msg := domain.Message{
		Sender: "webapp.Startup",
		Cmd:    "startup",
		Model:  nil,
	}
	a.Effects().ChanIn() <- msg
}

// domReady is called after front-end resources have been loaded
func (a *webapp) DomReady(ctx context.Context) {
	// Add your action here
	a.readyDOM = true
	a.logger.Debugf("DomReady [readyDOM:%v]", a.readyDOM)
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
	if err := a.echo.Shutdown(ctx); err != nil {
		a.logger.Errorf("%s echo shutdown error %s", modError, err.Error())
	}
}

// shutdown is called at application termination
func (a *webapp) OnShutdown() {
	// Perform your teardown here
	a.Shutdown(a.ctx)
}

func (a *webapp) Reductor() domain.Reductor {
	return a.reductor
}

func (a *webapp) Effects() domain.Effects {
	return a.effects
}

func (a *webapp) Route() {
	a.echo.GET("/page", a.CurrentPageIndex)
}

func (a *webapp) initDateMn() {
	loc, _ := time.LoadLocation("Europe/Moscow")
	t := time.Now().In(loc)
	_, m, _ := t.Date()
	a.startTime = time.Date(t.Year(), m, 1, 1, 0, 0, 0, loc)
	a.endTime = a.startTime.AddDate(0, 1, -1)
}

func (a *webapp) NowDateString() string {
	n := time.Now()
	return fmt.Sprintf("%4d.%02d.%02d %02d:%02d:%02d", n.Local().Year(), n.Local().Month(), n.Local().Day(), n.Local().Hour(), n.Local().Minute(), n.Local().Second())
}

func (a *webapp) StartDateString() string {
	return fmt.Sprintf("%4d.%02d.%02d", a.startTime.Local().Year(), a.startTime.Local().Month(), a.startTime.Local().Day())
}

func (a *webapp) EndDateString() string {
	return fmt.Sprintf("%4d.%02d.%02d", a.endTime.Local().Year(), a.endTime.Local().Month(), a.endTime.Local().Day())
}

func (a *webapp) SetStartDate(d time.Time) {
	a.startTime = d
}

func (a *webapp) SetEndDate(d time.Time) {
	a.endTime = d
}

func (a *webapp) StartDate() time.Time {
	return a.startTime
}

func (a *webapp) EndDate() time.Time {
	return a.endTime
}

func (a *webapp) SetRepo(repo domain.Repo) {
	a.repo = repo
}

func (a *webapp) FsrarID() string {
	return a.configuration.Application.Fsrarid
}

func (a *webapp) SetFsrarID(id string) {
	a.Config().Set("application.fsrarid", id, true)
}

func (a *webapp) Logger() *zap.SugaredLogger {
	return a.logger
}

func (a *webapp) Pwd() string {
	return a.pwd
}

func (a *webapp) Repo() domain.Repo {
	return a.repo
}
