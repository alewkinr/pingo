package space

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	neturl "net/url"

	"github.com/go-pkgz/requester"
	"github.com/go-pkgz/requester/middleware"
	"github.com/pkg/errors"
)

const (
	// name — навзание системы для которой клиент
	name = "space"
	// scheme — схема канала отправки сообщения
	scheme = "space"
)

// API — клиент для работы с API Space
type API struct {
	// clientName — навазние клиента
	clientName string
	// baseURL — базовый URL до API
	baseURL string
	// rq — HTTP клиент для работы с API
	rq *requester.Requester
}

// Option — набор опций для усовершенствования API
type Option func(api *API)

// WithTimeout устанавливает таймаут запроса.
func WithTimeout(t time.Duration) Option {
	return func(space *API) {
		space.rq.Client().Timeout = t
	}
}

// NewClient — создаем новый клиент
func NewClient(spaceHost string, token string, options ...Option) *API {
	const requesterTimeout = time.Second * 5
	api := API{
		clientName: name,
		baseURL:    spaceHost + "/api/http",
		rq: requester.New(
			http.Client{Timeout: requesterTimeout},
			middleware.JSON,
			middleware.Header("Authorization", fmt.Sprintf("Bearer %s", token)),
		),
	}

	for i := range options {
		options[i](&api)
	}

	return &api
}

// doRequest — обертка для выполнения запросов
func (space *API) doRequest(ctx context.Context, method, url string, body []byte) (*http.Response, error) {
	requestURL := space.baseURL + url
	request, makeRequestErr := http.NewRequestWithContext(ctx, method, requestURL, bytes.NewReader(body))
	if makeRequestErr != nil {
		return nil, errors.Wrap(makeRequestErr, "make http request")
	}

	response, doRequestErr := space.rq.Do(request)
	if doRequestErr != nil {
		return nil, errors.Wrap(doRequestErr, "do http request")
	}

	if response.StatusCode != http.StatusOK {
		return response, errors.New("not OK status code response")
	}
	return response, nil
}

// parseDestination — парсит "space:" по аналогии с "mailto:". Парсим URL и возвращаем channelID
func (space *API) parseDestination(destination string) (string, error) {
	u, parseURLErr := neturl.Parse(destination)
	if parseURLErr != nil {
		return "", errors.Wrap(parseURLErr, "parse destination")
	}

	if u.Scheme != scheme {
		return "", fmt.Errorf("схема %s не поддерживается, должно быть указано %s", u.Scheme, scheme)
	}

	return u.Opaque, nil
}

// Name — геттер для названия
func (space *API) Name() string {
	return space.clientName
}

// SendMessage — отправляем сообщение стандартной структуры
func (space *API) SendMessage(ctx context.Context, destination, message string) error {
	channelID, parseDestErr := space.parseDestination(destination)
	if parseDestErr != nil {
		return parseDestErr
	}

	url := "/chats/messages/send-message"
	payload := NewSendMessageRequest(message, channelID)

	body, marshalBodyErr := json.Marshal(payload)
	if marshalBodyErr != nil {
		return errors.Wrap(marshalBodyErr, "marshal request payload")
	}
	response, doRequestErr := space.doRequest(ctx, http.MethodPost, url, body)
	if doRequestErr != nil {
		return errors.Wrap(doRequestErr, fmt.Sprintf("space api request with response: %+v", response))
	}
	_ = response.Body.Close()
	return nil
}
