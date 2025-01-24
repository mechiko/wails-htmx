package trueclient

import (
	"firstwails/domain"
	"firstwails/zaplog"
	"fmt"
	"net"
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

// type TrueClient interface {
// 	PingSuz() (info *domain.PingSuzInfo, err error)
// 	PingSuzSilent() bool
// 	AuthGisSuz() error
// 	Balance(interface{}, int64) error
// 	SearchGis(target interface{}) error
// 	CisesList(target interface{}, cises []string) error
// 	CisesListPost(target interface{}, cises []string) error
// 	TokenGIS() string
// 	TokenSUZ() string
// 	AuthTime() time.Time
// 	Save()
// 	Errors() []string
// 	PingSuzInfo() *domain.PingSuzInfo
// }

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
var _ domain.ITrueClient = &trueClient{}

// клиент http по умолчанию
// var trueClientDefault = &http.Client{Timeout: 20 * time.Second}
// http.Client{
//   Transport: &http.Transport{
//     Dial: net.Dialer{Timeout: 2 * time.Second}.Dial,
//   },
// }

// инициализируем структурой с полями
// проверка необходимиости авторизации и ее выполнение
// паника если ошибки авторизации
// model указатель и изменяется авторизацией
// не отправляется в редуктор
func New(a domain.IApp, model *domain.TrueClient) (s *trueClient) {
	defer func() {
		if r := recover(); r != nil {
			errStr := fmt.Sprintf("%s panic %v", modError, r)
			s.errors = append(s.errors, errStr)
		}
	}()

	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			// Timeout: 20 * time.Second,
		}).Dial,
		// TLSHandshakeTimeout: 20 * time.Second,
	}
	var netClient = &http.Client{
		// Timeout:   time.Second * 120,
		Transport: netTransport,
	}
	s = &trueClient{
		IApp:       a,
		httpClient: netClient,
		layout:     model.LayoutUTC,
		// logger:     a.Logger(),
		logger:   zaplog.TrueSugar,
		urlSUZ:   model.StandSUZ,
		urlGIS:   model.StandGIS,
		tokenGis: model.TokenGIS,
		tokenSuz: model.TokenSUZ,
		hash:     model.HashKey,
		deviсeId: model.DeviceID,
		omsId:    model.OmsID,
		authTime: model.AuthTime,
		errors:   make([]string, 0),
	}
	s.Save(model)
	if (s.deviсeId) == "" {
		panic(fmt.Sprintf("%s %s", modError, "нужна настройка конфигурации"))
	}
	if (s.omsId) == "" {
		panic(fmt.Sprintf("%s %s", modError, "нужна настройка конфигурации"))
	}
	if (s.hash) == "" {
		panic(fmt.Sprintf("%s %s", modError, "нужна настройка конфигурации"))
	}
	// проверяем необходимость авторизации пингом СУЗ
	if s.CheckNeedAuth() {
		if err := s.AuthGisSuz(); err != nil {
			panic(fmt.Sprintf("%s %s", modError, err.Error()))
		}
		// сохраняем конфиг после авторизации
		s.Save(model)
	}
	return s
}

func NewPing(a domain.IApp, model domain.TrueClient) (s *trueClient) {
	defer func() {
		if r := recover(); r != nil {
			errStr := fmt.Sprintf("%s panic %v", modError, r)
			s.errors = append(s.errors, errStr)
		}
	}()

	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			// Timeout: 5 * time.Second,
		}).Dial,
		// TLSHandshakeTimeout: 5 * time.Second,
	}
	var netClient = &http.Client{
		// Timeout:   time.Second * 120,
		Transport: netTransport,
	}
	s = &trueClient{
		IApp:       a,
		httpClient: netClient,
		layout:     model.LayoutUTC,
		// logger:     a.Logger(),
		logger:   zaplog.TrueSugar,
		urlSUZ:   model.StandSUZ,
		urlGIS:   model.StandGIS,
		tokenGis: model.TokenGIS,
		tokenSuz: model.TokenSUZ,
		hash:     model.HashKey,
		deviсeId: model.DeviceID,
		omsId:    model.OmsID,
		authTime: model.AuthTime,
		errors:   make([]string, 0),
	}
	if (s.deviсeId) == "" {
		panic(fmt.Sprintf("%s %s", modError, "нужна настройка конфигурации"))
	}
	if (s.omsId) == "" {
		panic(fmt.Sprintf("%s %s", modError, "нужна настройка конфигурации"))
	}
	if (s.hash) == "" {
		panic(fmt.Sprintf("%s %s", modError, "нужна настройка конфигурации"))
	}
	// проверяем необходимость авторизации пингом СУЗ
	if s.CheckNeedAuthPing() {
		if err := s.AuthGisSuz(); err != nil {
			panic(fmt.Sprintf("%s %s", modError, err.Error()))
		}
		// сохраняем конфиг после авторизации
		s.Save(&model)
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
