package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config описывает структуру конфига
type Config struct {
	// Environment — окружение в котором запущено приложение
	Environment Environment `required:"true" envconfig:"ENVIRONMENT"`

	// Notify – конфигурация для нотификаций
	Notify *Notify

	// TemplatesConfig — конфигурация шаблонов, полученная из файла
	TemplatesConfig *TemplatesConfig
}

// InitConfig возвращает конфиг
func InitConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)

	return &cfg, err
}

// MustInitConfig возвращает конфиг или паникует при ошибке
func MustInitConfig() *Config {
	cfg, err := InitConfig()
	if err != nil {
		panic(err)
	}

	cfg.TemplatesConfig = mustParseTemplatesConfig()
	return cfg
}
