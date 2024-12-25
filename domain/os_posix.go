//go:build linux || darwin || freebsd

package domain

// import "github.com/mechiko/telebot/entity"
// if !entity.supported
var (
	DbPath               = "/var/local/telebot"
	LogPath              = "/var/log/telebot"
	ConfigPath           = "/etc/telebot"
	Supported            = true
	Linux                = true
	Windows              = false
	PosixUserUIDGUID int = 1002
	PosixChownPath   int = 0755
	PosixChownFile   int = 0644
)
