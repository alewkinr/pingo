package pingo

// Template — структура, описывающая шаблон
type Template struct {
	Destination string `yaml:"destination"`
	Text        string `yaml:"text"`
}

// GetDestination – геттер пункта назначения сообщения
func (t *Template) GetDestination() string {
	return t.Destination
}

// GetText – геттер текста сообщения
func (t *Template) GetText() string {
	return t.Text
}
