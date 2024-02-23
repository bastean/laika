package contentFromUrl

import (
	"github.com/bastean/laika/src/context/domain/aggregate"
	"github.com/bastean/laika/src/context/domain/repository"
	"github.com/bastean/laika/src/context/domain/service"
)

type ContentFromUrl struct {
	Repository repository.Repository
}

func (content *ContentFromUrl) Run(source string) {
	url := service.ParseUrl(source)
	host := url.Host
	path := url.Path
	laika := content.Repository.Read()

	for _, data := range laika.Sniffed[host] {
		if url.Path == data.Source {
			return
		}
	}

	html := service.ParseHtml(url.String())

	if html == "" {
		return
	}

	laika.Sniffed[host] = append(laika.Sniffed[host], &aggregate.Data{Source: path, Content: html, Found: make(map[string][]string)})

	content.Repository.Save(laika)

	links := service.ParseLinks(url.Scheme, url.Host, html)

	for _, link := range links {
		content.Run(link)
	}
}

func NewContentFromUrl(repository repository.Repository) *ContentFromUrl {
	return &ContentFromUrl{
		repository,
	}
}
