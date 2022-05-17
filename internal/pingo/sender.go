package pingo

import "context"

// Sender — интерфейс клиента для отправки сообщений
type Sender interface {
	// Name — геттер названия отправщика
	Name() string
	// SendMessage – метод для отправки сообщения
	SendMessage(ctx context.Context, destination, message string) error
}
