package main

import (
	"context"
	"time"

	"github.com/alewkinr/pingo/internal/config"
	"github.com/alewkinr/pingo/internal/pingo"
	"github.com/alewkinr/pingo/pkg/log"
	"github.com/alewkinr/pingo/pkg/message"
	"github.com/alewkinr/pingo/pkg/senders/email"
	"github.com/alewkinr/pingo/pkg/senders/slack"
	"github.com/alewkinr/pingo/pkg/senders/space"
	"github.com/alewkinr/pingo/pkg/senders/telegram"
	"github.com/alewkinr/pingo/pkg/trigger"
)

// Handler – обработчик для запросов Yandex.Cloud Functions
func Handler(ctx context.Context, r *trigger.TimerTriggerRequest) (*FunctionResponse, error) {
	settings := config.MustInitConfig()
	logger := log.SetUpLogging()

	senders := makeSenders(settings)
	pinger := pingo.NewPingo(logger, senders...)

	pinger.Ping(message.DailyReminder)
	return &FunctionResponse{}, nil
}

func main() {
	_, _ = Handler(context.Background(), &trigger.TimerTriggerRequest{})
}

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
