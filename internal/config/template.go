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
	// configFileName — название для файла конфигурации с шаблонами
	configFileName = "templates.yaml"
	// supportedAPIVersion — поддерживаемая версия API
	supportedAPIVersion = "1"
)

// TemplatesConfig — схема конфигурации шаблонов
type TemplatesConfig struct {
	Version   string                    `yaml:"version"`
	Templates map[string]pingo.Template `yaml:"templates"`
}

// NewTemplatesConfig — конструктор конфигурации шаблонов
func NewTemplatesConfig() *TemplatesConfig {
	return &TemplatesConfig{}
}

// MustInitLocal — парсим шаблоны для отправки сообщений из ЛОКАЛЬНОГО файла-конфига или паникуем
func (cfg *TemplatesConfig) MustInitLocal() {
	filepath, getAbsPathErr := path.Abs("./" + configFileName)
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

// mustParseConfig — парсим шаблоны из конфигурации или паникуем
func (cfg *TemplatesConfig) mustParseConfig(yamlConfig []byte) {
	unmarshalErr := yaml.Unmarshal(yamlConfig, &cfg)
	if unmarshalErr != nil {
		panic("unmarshal templates config file: " + unmarshalErr.Error())
	}

	if cfg.Version != supportedAPIVersion {
		panic(fmt.Sprintf("api version %s is not supported, must be %s", cfg.Version, supportedAPIVersion))
	}
}
