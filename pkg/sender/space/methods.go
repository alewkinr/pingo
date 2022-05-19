package space

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// Name — геттер для названия
func (space *API) Name() string {
	return space.clientName
}

// SendMessage — отправляем сообщение стандартной структуры
func (space *API) SendMessage(ctx context.Context, destination, message string) error {
	url := fmt.Sprintf("/chats/channels/%s/messages", destination)
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
