package sniffContentFromUrl

import (
	"github.com/bastean/laika/pkg/context/domain/aggregate"
	"github.com/bastean/laika/pkg/context/domain/service"
)

type SniffContentFromUrl struct{}

func (sniffContent *SniffContentFromUrl) Run(data *aggregate.Laika, source string) {
	url := service.ParseUrl(source)
	host := url.Host
	path := url.Path

	for _, data := range data.Sniffed[host] {
		if url.Path == data.Source {
			return
		}
	}

	html := service.ParseHtml(url.String())

	if html == "" {
		return
	}

	data.Sniffed[host] = append(data.Sniffed[host], &aggregate.Data{Source: path, Content: html, Found: make(map[string][]string)})

	links := service.ParseLinks(url.Scheme, url.Host, html)

	for _, link := range links {
		sniffContent.Run(data, link)
	}
}

func NewSniffContentFromUrl() *SniffContentFromUrl {
	return new(SniffContentFromUrl)
}
