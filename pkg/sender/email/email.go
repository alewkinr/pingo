package email

import (
	"context"

	"github.com/go-pkgz/notify"
)

// name — навзание системы для которой клиент
const name = "Email"

// Email — сендер для Email
type Email struct {
	// clientName — навазние клиента
	clientName string
	// rq — клиент для работы с Email
	rq *notify.Email
}

// NewEmail – констурктор отправщика для Email
func NewEmail(host string, port int, username, password string) *Email {
	email := notify.NewEmail(notify.SMTPParams{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	})
	return &Email{clientName: name, rq: email}
}

// Name — геттер для названия системы для сендера
func (s Email) Name() string {
	return s.clientName
}

// SendMessage — отправка сообщения
// destination — mailto:"John Wayne"<john@example.org>?subject=test-subj&from="Notifier"<notify@example.org>
func (s Email) SendMessage(ctx context.Context, destination, message string) error {
	return s.rq.Send(ctx, destination, message)
}
