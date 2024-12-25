//go:build !linux && !darwin && !windows && !freebsd

package domain

var (
	DbPath           = ""
	ConfigPath       = ""
	LogPath          = ""
	Supported        = false
	Windows          = false
	Linux            = false
	PosixUserUIDGUID = 1002
	PosixChownPath   = 0755
	PosixChownFile   = 0644
)
