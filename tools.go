package main

import (
	neturl "net/url"
	"time"

	"github.com/alewkinr/pingo/internal/config"
	"github.com/alewkinr/pingo/internal/pingo"
	"github.com/alewkinr/pingo/pkg/sender/email"
	"github.com/alewkinr/pingo/pkg/sender/slack"
	"github.com/alewkinr/pingo/pkg/sender/space"
	"github.com/alewkinr/pingo/pkg/sender/telegram"
)

// makeSenders — встраиваем отправщики пинг-сообщений
func makeSenders(settings *config.Config) map[pingo.Template]pingo.Sender {
	clients := make(map[pingo.Template]pingo.Sender, 0)

	if !settings.Notify.SpaceSettings.IsEmpty() {
		httpClientTimeout := time.Second * 5
		spaceAPI := space.NewClient(settings.Notify.SpaceSettings.Host, settings.Notify.SpaceSettings.Token, space.WithTimeout(httpClientTimeout))

		templates := listTemplatesBySender(settings.TemplatesConfig.Templates, "space")
		for _, tmpl := range templates {
			clients[tmpl] = spaceAPI
		}
	}

	if settings.Notify.SlackToken != "" {
		slackAPI := slack.NewSlack(settings.Notify.SlackToken)
		templates := listTemplatesBySender(settings.TemplatesConfig.Templates, "slack")
		for _, tmpl := range templates {
			clients[tmpl] = slackAPI
		}
	}

	if settings.Notify.TelegramToken != "" {
		telegramAPI := telegram.NewTelegram(settings.Notify.TelegramToken)
		templates := listTemplatesBySender(settings.TemplatesConfig.Templates, "telegram")
		for _, tmpl := range templates {
			clients[tmpl] = telegramAPI
		}
	}

	if !settings.Notify.EmailSettings.IsEmpty() {
		emailSMTPClient := email.NewEmail(settings.Notify.EmailSettings.Host, settings.Notify.EmailSettings.Port, settings.Notify.EmailSettings.Username, settings.Notify.EmailSettings.Password)
		templates := listTemplatesBySender(settings.TemplatesConfig.Templates, "mailto")
		for _, tmpl := range templates {
			clients[tmpl] = emailSMTPClient
		}
	}

	return clients
}

// listTemplatesBySender — фильтруем шаблоны по сендеру
func listTemplatesBySender(templates map[string]pingo.Template, senderScheme string) []pingo.Template {
	senders := make([]pingo.Template, 0)
	for _, tmpl := range templates {
		u, parseURLErr := neturl.Parse(tmpl.GetDestination())
		if parseURLErr != nil {
			continue
		}

		if u.Scheme == senderScheme {
			senders = append(senders, tmpl)
		}
	}

	if len(senders) == 0 {
		return nil
	}

	return senders
}
