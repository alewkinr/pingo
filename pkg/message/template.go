package message

// Template — шаблон сообщения для отпавки
type Template struct {
	// channel – канал для отправки
	channel string
	// text — текст сообщения
	text string
}

// NewTemplate — конструктор нового сообщения
func NewTemplate(channel string, text string) *Template {
	return &Template{channel: channel, text: text}
}

// Channel – геттер идентификатора канала
func (m *Template) Channel() string {
	return m.channel
}

// Text – геттер текста сообщения
func (m *Template) Text() string {
	return m.text
}
