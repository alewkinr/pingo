package pingo

import (
	"context"

	"github.com/alewkinr/pingo/pkg/message"
	"github.com/sirupsen/logrus"
)

// Pingo — сервис для пингов
type Pingo struct {
	// logger — логгер
	logger *logrus.Logger
	// senders — список клиентов для отправки пингов
	senders []Sender
}

// NewPingo — создаем новый пингер
func NewPingo(l *logrus.Logger, senders ...Sender) *Pingo {
	return &Pingo{
		logger:  l,
		senders: senders,
	}
}

// Ping — отправить пинг-сообщение
func (p Pingo) Ping(message *message.Template) {
	for _, sender := range p.senders {
		sendError := sender.SendMessage(context.Background(), message.Channel(), message.Text())
		if sendError != nil {
			p.logger.WithFields(logrus.Fields{
				"sender": sender.Name(),
				"error":  sendError,
			}).Error("send message")
		}
	}
}
