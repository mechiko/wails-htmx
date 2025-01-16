//go:build windows

package domain

const RootPathTemplateDebug = `E:\src\goproj\!!firstwails\src`

var (
	ConfigPath       = "./truestat"
	DbPath           = "./truestat"
	LogPath          = "./truestat"
	Supported        = true
	Windows          = true
	Linux            = false
	PosixUserUIDGUID = 1002
	PosixChownPath   = 0755
	PosixChownFile   = 0644
)
