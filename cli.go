package main

import (
	"context"

	"github.com/alewkinr/pingo/pkg/trigger"
)

func main() {
	_, _ = Handler(context.Background(), &trigger.TimerRequest{})
}
