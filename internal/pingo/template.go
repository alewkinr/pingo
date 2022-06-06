package pingo

// Template — структура, описывающая шаблон
type Template struct {
	Destination string `yaml:"destination" validate:"required"`
	Text        string `yaml:"text" validate:"required"`
}

// GetDestination – геттер пункта назначения сообщения
func (t *Template) GetDestination() string {
	return t.Destination
}

// GetText – геттер текста сообщения
func (t *Template) GetText() string {
	return t.Text
}
