package config

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	path "path/filepath"
	"time"

	"github.com/alewkinr/pingo/internal/pingo"
	"gopkg.in/yaml.v3"
)

const (
	// supportedAPIVersion — поддерживаемая версия API
	supportedAPIVersion = "1"
)

// TemplatesConfig — схема конфигурации шаблонов
type TemplatesConfig struct {
	Version   string                    `yaml:"version" validate:"required"`
	Templates map[string]pingo.Template `yaml:"templates" validate:"required"`
}

// NewTemplatesConfig — конструктор конфигурации шаблонов
func NewTemplatesConfig() *TemplatesConfig {
	return &TemplatesConfig{}
}

// InitLocal — парсим шаблоны для отправки сообщений из ЛОКАЛЬНОГО файла-конфига или паникуем
func (cfg *TemplatesConfig) InitLocal(configFile string) error {
	filepath, getAbsPathErr := path.Abs(configFile)
	if getAbsPathErr != nil {
		return fmt.Errorf("parse absolute path to templates config file: %s", getAbsPathErr.Error())
	}

	yamlConfigFile, readYamlConfigFileErr := ioutil.ReadFile(filepath)
	if readYamlConfigFileErr != nil {
		return fmt.Errorf("read yaml templates config file: %s", readYamlConfigFileErr.Error())
	}

	return cfg.parseConfig(yamlConfigFile)
}

// MustInitLocal — парсим шаблоны для отправки сообщений из ЛОКАЛЬНОГО файла-конфига или паникуем
// TODO: откзааться от паникующих функций
func (cfg *TemplatesConfig) MustInitLocal(configFile string) {
	filepath, getAbsPathErr := path.Abs(configFile)
	if getAbsPathErr != nil {
		panic("parse absolute path to templates config file: " + getAbsPathErr.Error())
	}

	yamlConfigFile, readYamlConfigFileErr := ioutil.ReadFile(filepath)
	if readYamlConfigFileErr != nil {
		panic("read yaml templates config file: " + readYamlConfigFileErr.Error())
	}

	cfg.mustParseConfig(yamlConfigFile)
}

// MustInitRemote — парсим шаблоны для отправки сообщений из УДАЛЕННОГО файла-конфига или паникуем
func (cfg *TemplatesConfig) MustInitRemote(remoteURL string) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFunc()

	request, makeRequestErr := http.NewRequestWithContext(ctx, http.MethodGet, remoteURL, nil)
	if makeRequestErr != nil {
		panic("make new http request: " + makeRequestErr.Error())
	}
	request.Header.Add("Content-Type", "text/yaml; charset=utf-8")
	response, doRequestErr := http.DefaultClient.Do(request)
	if doRequestErr != nil {
		panic("do http request: " + doRequestErr.Error())
	}
	defer func() { _ = response.Body.Close() }()

	if response.StatusCode != http.StatusOK {
		panic(fmt.Errorf("unexpected http status code: %d, with body: %v", response.StatusCode, response.Body))
	}

	body, readBodyErr := ioutil.ReadAll(response.Body)
	if readBodyErr != nil {
		panic("read response body: " + readBodyErr.Error())
	}

	cfg.mustParseConfig(body)
}

// parseConfig — парсим шаблоны из конфигурации или паникуем
func (cfg *TemplatesConfig) parseConfig(yamlConfig []byte) error {
	unmarshalErr := yaml.Unmarshal(yamlConfig, &cfg)
	if unmarshalErr != nil {
		return fmt.Errorf("unmarshal templates config file: %s", unmarshalErr.Error())
	}

	if cfg.Version != supportedAPIVersion {
		return fmt.Errorf("api version %s is not supported, must be %s", cfg.Version, supportedAPIVersion)
	}

	return nil
}

// mustParseConfig — парсим шаблоны из конфигурации или паникуем
// TODO: отказаться от паникующих функций
func (cfg *TemplatesConfig) mustParseConfig(yamlConfig []byte) {
	unmarshalErr := yaml.Unmarshal(yamlConfig, &cfg)
	if unmarshalErr != nil {
		panic("unmarshal templates config file: " + unmarshalErr.Error())
	}

	if cfg.Version != supportedAPIVersion {
		panic(fmt.Sprintf("api version %s is not supported, must be %s", cfg.Version, supportedAPIVersion))
	}
}
