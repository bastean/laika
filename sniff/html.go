package sniff

import (
	"html"
	"io"
	"net/http"
)

func Html(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	content, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return html.UnescapeString(string(content)), nil
}
