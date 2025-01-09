package domain

import (
	"net/url"
	"time"
)

type trueClient struct {
	StandGIS  url.URL
	StandSUZ  url.URL
	TokenGIS  string
	TokenSUZ  string
	AuthTime  time.Time
	LayoutUTC string
	HashKey   string
}
