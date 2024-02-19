package laika

import (
	"encoding/json"
	"os"

	"github.com/bastean/laika/sniff"
)

type Url struct {
	Path    string
	Content string
}

type Data struct {
	Path   string
	Emails []string
}

type Laika struct {
	Urls    map[string][]*Url
	Sniffed map[string][]*Data
}

func (laika *Laika) Dump(filename string) error {
	data, err := json.Marshal(laika)

	if err != nil {
		return err
	}

	err = os.WriteFile(filename+".json", data, 0644)

	if err != nil {
		return err
	}

	return nil
}

func (laika *Laika) Emails() []string {
	emails := []string{}

	for url, paths := range laika.Urls {
		for _, path := range paths {
			sniffedEmails := sniff.Emails(path.Content)

			if len(sniffedEmails) == 0 {
				continue
			}

			laika.Sniffed[url] = append(laika.Sniffed[url], &Data{Path: path.Path, Emails: sniffedEmails})

			emails = append(emails, sniffedEmails...)
		}
	}

	return emails
}

func Content(baseUrl, url string, content map[string][]*Url) {
	route, err := sniff.Url(url)

	if err != nil {
		return
	}

	for _, path := range content[baseUrl] {
		if route.Path == path.Path {
			return
		}
	}

	html, err := sniff.Html(url)

	if err != nil || html == "" {
		return
	}

	content[baseUrl] = append(content[baseUrl], &Url{Path: route.Path, Content: html})

	links := sniff.Links(route.Scheme, route.Host, html)

	for _, link := range links {
		Content(baseUrl, link, content)
	}
}

func Sniff(urls []string) *Laika {
	content := make(map[string][]*Url)

	for _, url := range urls {
		Content(url, url, content)
	}

	return &Laika{
		Urls:    content,
		Sniffed: make(map[string][]*Data),
	}
}
