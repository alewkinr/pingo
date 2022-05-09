package space

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

// SendMessage — отправляем сообщение стандартной структуры
func (space *API) SendMessage(ctx context.Context, channelID, message string) error {
	url := fmt.Sprintf("/chats/channels/%s/messages", channelID)
	payload := struct {
		Text string `json:"text"`
	}{
		message,
	}

	body, marshalBodyErr := json.Marshal(payload)
	if marshalBodyErr != nil {
		return errors.Wrap(marshalBodyErr, "marshal request payload")
	}
	response, doRequestErr := space.doRequest(ctx, http.MethodPost, url, body)
	if doRequestErr != nil {
		return errors.Wrap(doRequestErr, fmt.Sprintf("space api request with response: %+v", response))
	}
	return nil
}
