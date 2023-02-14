package rocket

import (
	"context"
	"fmt"
	"time"

	neturl "net/url"

	"github.com/badkaktus/gorocket"
	"github.com/pkg/errors"
)

const (
	// name — навзание системы для которой клиент
	name = "rocket"
	// scheme — название схемы для парсинга точки назначения
	scheme = "rocket"
)

// API — клиент для работы с API RocketChat
type API struct {
	// clientName — навазние клиента
	clientName string
	// baseURL — базовый URL до API
	baseURL string
	// rq — HTTP клиент для работы с API
	rq *gorocket.Client
}

// Option — набор опций для усовершенствования API
type Option func(api *API)

// WithTimeout устанавливает таймаут запроса.
func WithTimeout(t time.Duration) Option {
	return func(api *API) {
		api.rq.HTTPClient.Timeout = t
	}
}

// NewClient — создаем новый клиент
func NewClient(rocketChatHost string, userID, token string, options ...Option) *API {
	const requesterTimeout = time.Second * 5
	api := API{
		clientName: name,
		baseURL:    rocketChatHost + "/api",
		rq: gorocket.NewWithOptions(rocketChatHost,
			gorocket.WithUserID(userID),
			gorocket.WithXToken(token),
			gorocket.WithTimeout(requesterTimeout),
		),
	}

	for i := range options {
		options[i](&api)
	}

	return &api
}

// Name — геттер для названия
func (api *API) Name() string {
	return api.clientName
}

// SendMessage — отправляем сообщение стандартной структуры
func (api *API) SendMessage(ctx context.Context, destination, message string) error {
	channelID, parseDestErr := api.parseDestination(destination)
	if parseDestErr != nil {
		return parseDestErr
	}

	msg := &gorocket.Message{
		Channel: channelID,
		Text:    message,
	}
	resp, postMsgErr := api.rq.PostMessage(msg)
	if postMsgErr != nil {
		return errors.Wrap(postMsgErr, fmt.Sprintf("%s api request error: %s", api.Name(), postMsgErr.Error()))
	}

	if !resp.Success {
		return fmt.Errorf("%s api response error: %s:%s", api.Name(), resp.ErrorType, resp.Error)
	}

	return nil
}

// parseDestination — парсит "space:" по аналогии с "mailto:". Парсим URL и возвращаем channelID
func (api *API) parseDestination(destination string) (string, error) {
	u, parseURLErr := neturl.Parse(destination)
	if parseURLErr != nil {
		return "", errors.Wrap(parseURLErr, "parse destination")
	}

	if u.Scheme != scheme {
		return "", fmt.Errorf("схема %s не поддерживается, должно быть указано %s", u.Scheme, scheme)
	}

	return u.Opaque, nil
}
