package trueclient

import (
	"firstwails/domain"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"go.uber.org/zap"
)

const modError = "trueclient"

// type IApp interface {
// 	Logger() *zap.SugaredLogger
// 	Configuration() *domain.Configuration
// }

type TrueClient interface {
	PingSuz() (info *domain.PingSuzInfo, err error)
	PingSuzSilent() bool
	AuthGisSuz() error
	Balance(interface{}, int64) error
	SearchGis(target interface{}) error
	CisesList(target interface{}, cises []string) error
	TokenGIS() string
	TokenSUZ() string
	AuthTime() time.Time
	Save()
	Errors() []string
	PingSuzInfo() *domain.PingSuzInfo
}

type trueClient struct {
	domain.IApp
	// url        url.URL
	urlSUZ     url.URL
	urlGIS     url.URL
	layout     string
	logger     *zap.SugaredLogger
	tokenGis   string // токен авторизации для урла
	tokenSuz   string
	hash       string // кэп
	deviсeId   string
	omsId      string
	httpClient *http.Client
	authTime   time.Time
	errors     []string
	pingSUZ    *domain.PingSuzInfo
}

// создаем пустышку для проверки всех методов интерфейса
var _ TrueClient = &trueClient{}

// клиент http по умолчанию
var trueClientDefault = &http.Client{Timeout: 5 * time.Second}

// инициализируем структурой с полями
// проверка необходимиости авторизации и ее выполнение
// паника если ошибки авторизации
func New(a domain.IApp, model domain.TrueClient) TrueClient {
	s := &trueClient{
		IApp:       a,
		httpClient: trueClientDefault,
		layout:     model.LayoutUTC,
		logger:     a.Logger(),
		urlSUZ:     model.StandSUZ,
		urlGIS:     model.StandGIS,
		tokenGis:   model.TokenGIS,
		tokenSuz:   model.TokenSUZ,
		hash:       model.HashKey,
		deviсeId:   model.DeviceID,
		omsId:      model.OmsID,
		authTime:   model.AuthTime,
		errors:     make([]string, 0),
	}
	if s.hash == "" {
		panic(fmt.Sprintf("%s panic %s", modError, "нужен хэш ключа в конфигурации"))
	}
	if s.CheckNeedAuth() {
		if err := s.AuthGisSuz(); err != nil {
			panic(fmt.Sprintf("%s panic %s", modError, err.Error()))
		}
		// сохраняем конфиг после авторизации
		s.Save()
	}
	return s
}

func (t *trueClient) TokenGIS() string {
	return t.tokenGis
}

func (t *trueClient) TokenSUZ() string {
	return t.tokenSuz
}

func (t *trueClient) AuthTime() time.Time {
	return t.authTime
}

func (t *trueClient) Errors() []string {
	return t.errors
}

func (t *trueClient) AddError(err error) {
	t.errors = append(t.errors, err.Error())
}

func (t *trueClient) PingSuzInfo() *domain.PingSuzInfo {
	return t.pingSUZ
}
