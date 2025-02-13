package domain

import (
	"net/url"
	"time"
)

type PingSuzInfo struct {
	OmsId      string `json:"omsId"`
	ApiVersion string `json:"apiVersion"`
	OmsVersion string `json:"omsVersion"`
}

type TrueClient struct {
	Title       string
	StandGIS    url.URL
	StandSUZ    url.URL
	TokenGIS    string
	TokenSUZ    string
	AuthTime    time.Time
	LayoutUTC   string
	HashKey     string
	DeviceID    string
	OmsID       string
	UseConfigDB bool
	Errors      []string
	PingSuz     *PingSuzInfo
	Validates   map[string]string
	MyStore     map[string]string
}
