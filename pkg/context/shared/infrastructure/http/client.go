package http

import (
	"io"
	"net/http"

	"github.com/bastean/laika/pkg/context/shared/domain/model"
)

type Client struct{}

func (client *Client) Get(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	response, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(response), nil
}

func NewClient() model.HttpClient {
	return new(Client)
}
