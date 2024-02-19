package sniff

import (
	"io"
	"net/http"
)

func Html(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	html, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(html), nil
}
