package config

// Notify — конфигурация для нотификаций
type Notify struct {
	// SpaceSettings – конфигурация Jet Brains Space
	SpaceSettings *Space
	// SlackToken — API токен для работы с API Slack
	SlackToken string `envconfig:"SLACK_TOKEN"`
	// TelegramToken — API токен для работы с API Telegram
	TelegramToken string `envconfig:"TELEGRAM_TOKEN"`
	// EmailSettings — настройки для уведомлений по Email
	EmailSettings *SMTP
	// RocketChatSettings — настройки для уведомлений в RocketChat
	RocketChatSettings *RocketChat
}

// SMTP — настройки SMTP сервера
type SMTP struct {
	// Host — SMTP хост
	Host string `envconfig:"SMTP_HOST"`
	// Port — SMTP порт
	Port int `envconfig:"SMTP_PORT" validate:"required_with=Host"`
	// Username — пользователь SMTP-сервера
	Username string `envconfig:"SMTP_USERNAME" validate:"required_with=Host"`
	// Password — пароль от SMTP сервера
	Password string `envconfig:"SMTP_PASSWORD" validate:"required_with=Host"`
}

// IsEmpty — проверяем, что SMTP настройки пустые
func (email SMTP) IsEmpty() bool {
	return email == SMTP{}
}

// Space — конфигурация для работы с JetBrains Space
type Space struct {
	// Host — адрес до инсталляции
	Host string `envconfig:"SPACE_HOST"`
	// Token — токен доступа к API
	Token string `envconfig:"SPACE_TOKEN" validate:"required_with=Host"`
}

// IsEmpty — проверяем, что Space настройки пустые
func (s Space) IsEmpty() bool {
	return s == Space{}
}

// RocketChat — конфигурация для работы с RocketChat
type RocketChat struct {
	// Host — адрес до инсталляции
	Host string `envconfig:"ROCKETCHAT_HOST"`
	// Token — токен доступа к API
	Token string `envconfig:"ROCKETCHAT_TOKEN" validate:"required_with=Host"`
	// UserID — пользователь для авторизации
	UserID string `envconfig:"ROCKETCHAT_USER_ID" validate:"required_with=Host"`
}

// IsEmpty — проверяем, что RocketChat настройки пустые
func (s RocketChat) IsEmpty() bool {
	return s == RocketChat{}
}
