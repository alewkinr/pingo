package main

import (
	"context"
	"github.com/alewkinr/pingo/internal/config"
	"github.com/alewkinr/pingo/internal/trigger"
	"github.com/alewkinr/pingo/pkg/log"
	"github.com/alewkinr/pingo/pkg/message"
	"github.com/alewkinr/pingo/pkg/space"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"time"
)

// Handler – обработчик для запросов Yandex.Cloud Functions
func Handler(ctx context.Context, r *trigger.TimerTriggerRequest) (*FunctionResponse, error) {
	settings := config.MustInitConfig()
	logger := log.SetUpLogging()

	httpClientTimeout := time.Second * 5
	api := space.NewClient(settings.Space.Host, settings.Space.Token, space.WithTimeout(httpClientTimeout))

	channel, text := message.DailyReminder.Channel(), message.DailyReminder.Text()
	if settings.Environment.IsDevelopment() {
		channel = settings.Space.DebugChannel
	}

	sendError := api.SendMessage(ctx, channel, text)
	if sendError != nil {
		logger.WithFields(logrus.Fields{"error": sendError}).Error("space send message")
		return nil, errors.Wrap(sendError, "space send message")
	}
	return &FunctionResponse{}, nil
}

func main() {
	_, _ = Handler(context.Background(), &trigger.TimerTriggerRequest{})
}
