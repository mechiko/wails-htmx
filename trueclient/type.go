package trueclient

import (
	"firstwails/domain"
	"net/http"
	"net/url"
	"sync"
	"time"

	"go.uber.org/zap"
)

const modError = "trueclient"

type IApp interface {
	Logger() *zap.SugaredLogger
	Configuration() *domain.Configuration
}

type TrueClient interface {
	PingSuz(target interface{}) error
	AuthGisSuz() error
	Balance(interface{}, int64) error
	SearchGis(target interface{}) error
	CisesList(target interface{}, cises []string) error
}

type trueClient struct {
	IApp
	url        url.URL
	layout     string
	logger     *zap.SugaredLogger
	tokenGis   string // токен авторизации для урла
	tokenSuz   string
	hash       string // кэп
	deviсeId   string
	omsId      string
	httpClient *http.Client
}

// создаем пустышку для проверки всех методов интерфейса
var _ TrueClient = &trueClient{}

var instance *trueClient = nil
var once sync.Once

// клиент http по умолчанию
var trueClientDefault = &http.Client{Timeout: 5 * time.Second}

func New(a IApp, url url.URL, hash string, oms string, device string, tokenGis string, tokenSuz string) TrueClient {
	s := &trueClient{
		IApp:       a,
		httpClient: trueClientDefault,
		layout:     a.Configuration().Layouts.TimeLayout,
		logger:     a.Logger(),
		url:        url,
		tokenGis:   tokenGis,
		tokenSuz:   tokenSuz,
		hash:       hash,
		deviсeId:   device,
		omsId:      oms,
	}
	return s
}
