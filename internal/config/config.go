package config

import (
	"github.com/alewkinr/pingo/pkg/validation"
	"github.com/kelseyhightower/envconfig"
)

// Config описывает структуру конфига
type Config struct {
	// Environment — окружение в котором запущено приложение
	Environment Environment `envconfig:"ENVIRONMENT" validate:"required,oneof=development staging production"`
	// RemoteConfigURL — URL для скачивания конфигурации шаблонов
	RemoteConfigURL string `envconfig:"REMOTE_CONFIG_URL" validate:"omitempty,url"`
	// Notify – конфигурация для нотификаций
	Notify *Notify `validate:"required"`
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

	cfg.TemplatesConfig = NewTemplatesConfig()

	if cfg.RemoteConfigURL != "" {
		cfg.TemplatesConfig.MustInitRemote(cfg.RemoteConfigURL)
	}

	if cfg.RemoteConfigURL == "" {
		cfg.TemplatesConfig.MustInitLocal("./templates.yaml")
	}

	// валидируем конфиг
	v := validation.NewPlayground()
	validationErr := v.Validate(cfg)
	if validationErr != nil {
		panic(validationErr)
	}

	return cfg
}
