package telegram

import (
	"context"

	"github.com/go-pkgz/notify"
)

const (
	// name — навзание системы для которой клиент
	name = "Telegram"
	// successAuthMessage — сообщение об успешном подключении сервиса
	successAuthMessage = "🔌 Pingo integration connected successfully!"
	// errorAuthMessage — сообщение о неуспешном подключении
	errorAuthMessage = "⚠️ Pingo integration not connected!"
)

// Telegram — сендер Telegram
type Telegram struct {
	// clientName — навазние клиента
	clientName string
	// rq — клиент для работы с Telegram
	rq *notify.Telegram
}

// NewTelegram – констурктор отправщика для Telegram
func NewTelegram(token string) *Telegram {
	tg, _ := notify.NewTelegram(notify.TelegramParams{
		Token:      token,
		ErrorMsg:   successAuthMessage,
		SuccessMsg: errorAuthMessage,
	})

	return &Telegram{clientName: name, rq: tg}
}

// Name — геттер для названия системы для сендера
func (t Telegram) Name() string {
	return t.clientName
}

// SendMessage — отправка сообщения
func (t Telegram) SendMessage(ctx context.Context, channelID, message string) error {
	return t.rq.Send(ctx, channelID, message)
}
