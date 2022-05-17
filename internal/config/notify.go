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
}

// SMTP — настройки SMTP сервера
type SMTP struct {
	// Host — SMTP хост
	Host string `envconfig:"SMTP_HOST"`
	// Port — SMTP порт
	Port int `envconfig:"SMTP_PORT"`
	// Username — пользователь SMTP-сервера
	Username string `envconfig:"SMTP_USERNAME"`
	// Password — пароль от SMTP сервера
	Password string `envconfig:"SMTP_PASSWORD"`
}

// Space — конфигурация для работы с JetBrains Space
type Space struct {
	// Host — адрес до инсталляции
	Host string `required:"true" envconfig:"SPACE_HOST"`
	// Token — токен доступа к API
	Token string `required:"true" envconfig:"SPACE_TOKEN"`
	// DebugChannel — ID канала для отладки отправки сообщений
	DebugChannel string `envconfig:"SPACE_DEBUG_CHANNEL"`
}
