package cmd

import (
	"context"
	"fmt"

	"github.com/alewkinr/pingo/internal/config"
	"github.com/alewkinr/pingo/internal/pingo"
	"github.com/alewkinr/pingo/pkg/validation"
	"github.com/sirupsen/logrus"
)

const (
	// ExitNoError — выход из приложения без обшибки
	ExitNoError = iota
	// ExitWithErr — выход из приложения с ошибкой
	ExitWithErr
)

// CLI — структура интерфейса командой строки
type CLI struct {
	// l — логгер для приложения
	l *logrus.Logger
	// v — валидатор для приложения
	v validation.Validator
	// cfg — конфигурация приложения
	cfg *config.Config
}

// NewCLI — создаем CLI-приложение
func NewCLI(logger *logrus.Logger, v validation.Validator, cfg *config.Config) *CLI {
	return &CLI{
		l:   logger,
		v:   v,
		cfg: cfg,
	}
}

// L — геттер для логгера
func (c *CLI) L() *logrus.Logger {
	return c.l
}

// V — геттер для валидатора
func (c *CLI) V() validation.Validator {
	return c.v
}

// Setting — геттер настроек прилжения
func (c *CLI) Setting() *config.Config {
	return c.cfg
}

// Send — команда отправки сообщения произвольного содержания в произвольный канал назначения
func (c *CLI) Send(destination, text string) (int, error) {
	tmpTemplateName := "cli_send_template_once"
	tmpSendTemplate := pingo.Template{Destination: destination, Text: text}
	c.Setting().TemplatesConfig.Templates[tmpTemplateName] = tmpSendTemplate

	if validationErr := c.V().Validate(c.Setting()); validationErr != nil {
		return ExitWithErr, validationErr
	}

	pinger := pingo.NewPingo(c.L(), MakeSenders(c.Setting()))
	runErr := pinger.Run(context.TODO(), map[string]pingo.Template{tmpTemplateName: tmpSendTemplate})
	if runErr != nil {
		return ExitWithErr, runErr
	}

	return ExitNoError, runErr
}

// SendTemplate — отпавка шаблона сообщения
func (c *CLI) SendTemplate(configFilePath, templateName string) (int, error) {
	if initCfgErr := c.Setting().TemplatesConfig.InitLocal(configFilePath); initCfgErr != nil {
		return ExitWithErr, initCfgErr
	}

	if validationErr := c.V().Validate(c.Setting()); validationErr != nil {
		return ExitWithErr, validationErr
	}

	template, ok := c.Setting().TemplatesConfig.Templates[templateName]
	if !ok {
		return ExitWithErr, fmt.Errorf("missing template %s in config. Check %s file", template, configFilePath)
	}
	pinger := pingo.NewPingo(c.L(), MakeSenders(c.Setting()))
	runErr := pinger.Run(context.TODO(), map[string]pingo.Template{templateName: template})
	if runErr != nil {
		return ExitWithErr, runErr
	}

	return ExitNoError, runErr
}
