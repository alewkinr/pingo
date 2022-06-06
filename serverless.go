package main

import (
	"context"

	"github.com/alewkinr/pingo/cmd"
	"github.com/alewkinr/pingo/internal/config"
	"github.com/alewkinr/pingo/internal/pingo"
	"github.com/alewkinr/pingo/pkg/log"
	"github.com/alewkinr/pingo/pkg/trigger"
)

// Handler – обработчик для запросов Yandex.Cloud Functions
//nolint:unparam,deadcode
func Handler(ctx context.Context, r *trigger.TimerRequest) (struct{}, error) {
	settings := config.MustInitConfig()
	logger := log.SetUpLogging()

	senders := cmd.MakeSenders(settings)

	pinger := pingo.NewPingo(logger, senders)

	pingErr := pinger.Run(ctx, settings.TemplatesConfig.Templates)
	if pingErr != nil {
		return struct{}{}, pingErr
	}
	return struct{}{}, nil
}
