package space

import "fmt"

// ChatMessageTextContent — класс с типом текстового сообщения
const ChatMessageTextContent = "ChatMessage.Text"

// SendMessageRequest — структура запроса на отправку сообщения
type SendMessageRequest struct {
	Content sendMessageContent `json:"content"`
	Channel string             `json:"channel"`
}

// sendMessageContent — контент сообщения
type sendMessageContent struct {
	ClassName string `json:"className"`
	Text      string `json:"text"`
}

// NewSendMessageRequest — конструктор нового сообщения
func NewSendMessageRequest(message, channelID string) SendMessageRequest {
	return SendMessageRequest{
		Content: sendMessageContent{
			ClassName: ChatMessageTextContent,
			Text:      message,
		},
		Channel: fmt.Sprintf("channel:id:%s", channelID),
	}
}
