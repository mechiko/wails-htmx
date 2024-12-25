package webapp

import (
	"firstwails/config"
	"firstwails/domain"
	"fmt"
)

func (a *webapp) initConfig() error {
	// Configuration
	config.TomlConfig = domain.TomlConfig
	tt, err := config.New(a.logger, "", a.configuration)
	if err != nil {
		return fmt.Errorf("app:init %w", err)
	}
	a.config = tt
	return nil
}

func (a *webapp) Config() domain.IConfig {
	return a.config
}

func (a *webapp) Configuration() *domain.Configuration {
	return a.configuration
}
