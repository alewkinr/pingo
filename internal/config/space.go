package config

// Space — конфигурация для работы с JetBrains Space
type Space struct {
	// Host — адрес до инсталляции
	Host string `required:"true" envconfig:"SPACE_HOST"`
	// Token — токен доступа к API
	Token string `required:"true" envconfig:"SPACE_TOKEN"`
	// DebugChannel — ID канала для отладки отправки сообщений
	DebugChannel string `envconfig:"SPACE_DEBUG_CHANNEL"`
}
