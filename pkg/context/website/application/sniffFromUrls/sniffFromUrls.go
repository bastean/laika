package sniffFromUrls

import (
	"log"

	"github.com/bastean/laika/pkg/context/shared/domain/aggregate"
	"github.com/bastean/laika/pkg/context/shared/domain/model"
	"github.com/bastean/laika/pkg/context/shared/domain/service"
)

type SniffFromUrls struct {
	Scraper model.Scraper
}

type SniffFromUrlsOptions struct {
	FollowLinks bool
}

func (sniff *SniffFromUrls) Run(data aggregate.Data, sources []string, option *SniffFromUrlsOptions) {
	for _, source := range sources {
		url, err := service.ParseUrl(source)

		if err != nil {
			return
		}

		host := url.Host
		path := url.Path

		for _, sniffed := range data {
			for _, source := range sniffed {
				if path == source.Source {
					return
				}
			}
		}

		rawHtml := sniff.Scraper.GetContent(url.String())

		if rawHtml == "" {
			return
		}

		html := service.ParseHtml(rawHtml)

		emails := service.SniffEmails(html)

		data[host] = append(data[host], &aggregate.Sniffed{Source: path, Emails: emails})

		log.Printf("%s | Emails: %d\n", url, len(emails))

		if option.FollowLinks {
			links := sniff.Scraper.GetLinks(url.String())

			if len(links) == 0 {
				return
			}

			links = service.ParseLinks(url.Scheme, host, links)

			sniff.Run(data, links, option)
		}
	}
}

func NewSniffFromUrls(scraper model.Scraper) *SniffFromUrls {
	return &SniffFromUrls{
		Scraper: scraper,
	}
}
