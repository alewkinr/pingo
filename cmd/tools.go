package main

import (
	"time"

	"github.com/alewkinr/pingo/internal/config"
	"github.com/alewkinr/pingo/internal/pingo"
	"github.com/alewkinr/pingo/pkg/sender/email"
	"github.com/alewkinr/pingo/pkg/sender/slack"
	"github.com/alewkinr/pingo/pkg/sender/space"
	"github.com/alewkinr/pingo/pkg/sender/telegram"
)

// makeSenders — создаем отправщиков пинг-сообщения
func makeSenders(settings *config.Config) []pingo.Sender {
	clients := make([]pingo.Sender, 0)

	if settings.Notify.SpaceSettings != nil {
		httpClientTimeout := time.Second * 5
		spaceAPI := space.NewClient(settings.Notify.SpaceSettings.Host, settings.Notify.SpaceSettings.Token, space.WithTimeout(httpClientTimeout))
		clients = append(clients, spaceAPI)
	}

	if settings.Notify.SlackToken != "" {
		slackAPI := slack.NewSlack(settings.Notify.SlackToken)
		clients = append(clients, slackAPI)
	}

	if settings.Notify.TelegramToken != "" {
		telegramAPI := telegram.NewTelegram(settings.Notify.TelegramToken)
		clients = append(clients, telegramAPI)
	}

	if settings.Notify.EmailSettings != nil {
		emailSMTPClient := email.NewEmail(settings.Notify.EmailSettings.Host, settings.Notify.EmailSettings.Port, settings.Notify.EmailSettings.Username, settings.Notify.EmailSettings.Password)
		clients = append(clients, emailSMTPClient)
	}

	return clients
}
