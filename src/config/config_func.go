package config

import (
	"firstwails/domain"
	"fmt"
	"os"
)

// это структура конфигурации потому используем пустой интерфейс он принимает все
func (c *Config) Get() interface{} {
	return c.configuration
}

func (c *Config) Unmarshal(cfg *domain.Configuration) error {
	if err := c.viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	return nil
}

// тут функции надстройки возможно и не так полезные, хотя для записи значений в конфиг используются
// Viper пишет только то что прочитал изначально из конфига, если меняем что то
// то надо усатнавливать через Set а не просто в Configuration
func (c *Config) Save() error {
	if _, err := os.Stat(c.configFileName); os.IsNotExist(err) {
		if err := c.viper.SafeWriteConfig(); err != nil {
			return fmt.Errorf("%s %w", modError, err)
		}
	} else {
		if err := c.viper.WriteConfig(); err != nil {
			return fmt.Errorf("%s %w", modError, err)
		}
	}
	return nil
}

func (c *Config) SaveAs(fn string) error {
	err := c.viper.WriteConfigAs(fn)
	if err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	return nil
}

func (c *Config) SaveSafe() error {
	err := c.viper.SafeWriteConfig()
	if err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	return nil
}

// GetString("datastore.metric.host")
func (c *Config) GetByName(name string) interface{} {
	return c.viper.Get(name)
}

// записываем ключ и его значение, затем обновляем структуру Config этими значениями
func (c *Config) Set(key string, value interface{}, save ...bool) error {
	c.viper.Set(key, value)
	if err := c.viper.Unmarshal(c.configuration); err != nil {
		return fmt.Errorf("Viper.Unmarshal(c.Configuration) %w", err)
	}
	if len(save) > 0 {
		saving := save[0]
		if saving {
			if err := c.Save(); err != nil {
				return fmt.Errorf("%s %w", modError, err)
			}
		}
	}
	return nil
}
