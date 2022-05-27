package config

import (
	"fmt"
	"io/ioutil"
	path "path/filepath"

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

// mustParseTemplatesConfig — парсим шаблоны для отправки сообщений из yaml конфигурации или паникуем
func mustParseTemplatesConfig() *TemplatesConfig {
	filepath, getAbsPathErr := path.Abs("./" + configFileName)
	if getAbsPathErr != nil {
		panic("parse absolute path to templates config file: " + getAbsPathErr.Error())
	}

	yamlConfigFile, readYamlConfigFileErr := ioutil.ReadFile(filepath)
	if readYamlConfigFileErr != nil {
		panic("read yaml templates config file: " + readYamlConfigFileErr.Error())
	}

	var config TemplatesConfig
	unmarshalErr := yaml.Unmarshal(yamlConfigFile, &config)
	if unmarshalErr != nil {
		panic("unmarshal templates config file: " + unmarshalErr.Error())
	}

	if config.Version != supportedAPIVersion {
		panic(fmt.Sprintf("api version %s is not supported, must be %s", config.Version, supportedAPIVersion))
	}

	return &config
}
