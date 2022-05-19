package main

import (
	"context"

	"github.com/alewkinr/pingo/internal/config"
	"github.com/alewkinr/pingo/internal/pingo"
	"github.com/alewkinr/pingo/pkg/log"
	"github.com/alewkinr/pingo/pkg/message"
	"github.com/alewkinr/pingo/pkg/trigger"
)

// Handler – обработчик для запросов Yandex.Cloud Functions
func Handler(ctx context.Context, r *trigger.TimerRequest) (struct{}, error) {
	settings := config.MustInitConfig()
	logger := log.SetUpLogging()

	senders := makeSenders(settings)
	pinger := pingo.NewPingo(logger, senders...)

	pinger.Ping(ctx, message.DailyReminder)
	return struct{}{}, nil
}
