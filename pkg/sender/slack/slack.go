package slack

import (
	"context"

	"github.com/go-pkgz/notify"
)

// name — навзание системы для которой клиент
const name = "slack"

// Slack — сендер для Slack
type Slack struct {
	// clientName — навазние клиента
	clientName string
	// rq — клиент для работы с Slack
	rq *notify.Slack
}

// NewSlack – констурктор отправщика для Slack
func NewSlack(token string) *Slack {
	s := notify.NewSlack(token)
	return &Slack{clientName: name, rq: s}
}

// Name — геттер для названия системы для сендера
func (s Slack) Name() string {
	return s.clientName
}

// SendMessage — отправка сообщения
func (s Slack) SendMessage(ctx context.Context, channelID, message string) error {
	return s.rq.Send(ctx, channelID, message)
}
