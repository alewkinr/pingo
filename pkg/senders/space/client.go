package space

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-pkgz/requester"
	"github.com/go-pkgz/requester/middleware"
	"github.com/pkg/errors"
)

// name — навзание системы для которой клиент
const name = "Jet Brains Space"

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
	requestUrl := space.baseURL + url
	request, makeRequestErr := http.NewRequestWithContext(ctx, method, requestUrl, bytes.NewReader(body))
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
