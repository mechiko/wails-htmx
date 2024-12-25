//go:build windows

package domain

var (
	ConfigPath       = "./alcogo"
	DbPath           = "./alcogo"
	LogPath          = "./alcogo"
	Supported        = true
	Windows          = true
	Linux            = false
	PosixUserUIDGUID = 1002
	PosixChownPath   = 0755
	PosixChownFile   = 0644
)
