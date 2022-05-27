package pingo

import (
	"context"

	"github.com/sirupsen/logrus"
)

// Pingo — сервис для пингов
type Pingo struct {
	// logger — логгер
	logger *logrus.Logger
	// senders — клиенты для отправки шаблонных сообщений
	senders map[Template]Sender
}

// NewPingo — создаем новый пингер
func NewPingo(l *logrus.Logger, senders map[Template]Sender) *Pingo {
	return &Pingo{
		logger:  l,
		senders: senders,
	}
}

// Run — отправить пинг-сообщение
func (p Pingo) Run(ctx context.Context, templates map[string]Template) error {
	for templateName, tmpl := range templates {
		templateDestination, templateText := tmpl.GetDestination(), tmpl.GetText()

		sender, ok := p.senders[tmpl]
		if !ok {
			p.logger.WithFields(logrus.Fields{"templateName": templateName}).Error("sender for template not found")
			continue
		}

		sendError := sender.SendMessage(ctx, templateDestination, templateText)
		if sendError != nil {
			p.logger.WithFields(logrus.Fields{
				"templateName": templateName,
				"sender":       sender.Name(),
			}).Error(sendError)
		}
	}

	return nil
}
