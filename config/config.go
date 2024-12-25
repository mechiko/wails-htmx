package config

import (
	"firstwails/domain"
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	domain.IConfig
	logger         *zap.SugaredLogger
	viper          *viper.Viper
	configuration  interface{}
	configFileName string
}

var TomlConfig []byte

const modError = "pkg:config"

// создание нового экземпляра конфига реентерабельно возвращается интерфейс
// configuration - указатель на структуру куда unmarshall конфигурация
func New(logger *zap.SugaredLogger, cfgName string, configuration interface{}) (cfg *Config, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
	}()

	configName := domain.Name
	if cfgName != "" {
		configName = cfgName
	}
	viperOrigin := viper.GetViper()
	configFileName := configName + ".toml"

	viperOrigin.SetConfigName(configName)
	viperOrigin.SetConfigType("toml")
	viperOrigin.AddConfigPath(".")

	if err := viperOrigin.MergeConfig(strings.NewReader(string(TomlConfig))); err != nil {
		return nil, fmt.Errorf("%s %w", modError, err)
	}

	if err := viper.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("%s %w", modError, err)
		} else {
			logger.Errorf("%s Config file ('%s') not found", modError, configFileName)
		}
	}

	if err := viper.Unmarshal(configuration); err != nil {
		return nil, fmt.Errorf("%s %w", modError, err)
	}

	cfg = &Config{
		logger:         logger,
		configuration:  configuration,
		configFileName: configFileName,
		viper:          viperOrigin,
	}
	viperOrigin.SafeWriteConfig()

	return cfg, nil
}
