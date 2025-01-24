package domain

import "time"

type ITrueClient interface {
	PingSuz() (info *PingSuzInfo, err error)
	PingSuzSilent() bool
	AuthGisSuz() error
	Balance(interface{}, int64) error
	SearchGis(target interface{}) error
	CisesList(target interface{}, cises []string) error
	CisesListPost(target interface{}, cises []string) error
	TokenGIS() string
	TokenSUZ() string
	AuthTime() time.Time
	Save(*TrueClient)
	Errors() []string
	PingSuzInfo() *PingSuzInfo
}
