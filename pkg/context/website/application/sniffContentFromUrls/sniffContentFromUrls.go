package sniffContentFromUrls

import (
	"github.com/bastean/laika/pkg/context/shared/domain/aggregate"
	"github.com/bastean/laika/pkg/context/shared/domain/model"
	"github.com/bastean/laika/pkg/context/shared/domain/service"
)

type SniffContentFromUrls struct {
	Client model.HttpClient
}

func (sniff *SniffContentFromUrls) Run(data *aggregate.Data, sources []string) {
	for _, source := range sources {
		url, err := service.ParseUrl(source)

		if err != nil {
			return
		}

		host := url.Host
		path := url.Path

		for _, sniffed := range data.Sniffed[host] {
			if path == sniffed.Source {
				return
			}
		}

		rawHtml, err := sniff.Client.Get(url.String())

		if rawHtml == "" || err != nil {
			return
		}

		html := service.ParseHtml(string(rawHtml))

		if html == "" {
			return
		}

		data.Sniffed[host] = append(data.Sniffed[host], &aggregate.Sniffed{Source: path, Content: html, Found: make(map[string][]string)})

		links := service.ParseLinks(url.Scheme, url.Host, html)

		sniff.Run(data, links)
	}
}

func NewSniffContentFromUrls(client model.HttpClient) *SniffContentFromUrls {
	return &SniffContentFromUrls{
		Client: client,
	}
}
