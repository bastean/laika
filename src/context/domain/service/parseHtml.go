package service

import (
	"html"
	"io"
	"net/http"
)

func ParseHtml(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	content, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return html.UnescapeString(string(content))
}
