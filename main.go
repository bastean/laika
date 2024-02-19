package laika

import (
	"encoding/json"
	"os"

	"github.com/bastean/laika/sniff"
)

type Data struct {
	Path   string
	Emails []string
}

type Laika struct {
	Urls    []string
	Sniffed map[string][]*Data
}

func (laika *Laika) Dump(filename string) error {
	data, err := json.Marshal(laika.Sniffed)

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

	for _, url := range laika.Urls {
		html, err := sniff.Html(url)

		if err != nil {
			continue
		}

		sniffedEmails := sniff.Emails(html)

		if len(sniffedEmails) == 0 {
			continue
		}

		path, err := sniff.Path(url)

		if err != nil {
			continue
		}

		laika.Sniffed[url] = append(laika.Sniffed[url], &Data{Path: path, Emails: sniffedEmails})

		emails = append(emails, sniffedEmails...)
	}

	return emails
}

func Sniff(urls []string) *Laika {
	return &Laika{
		Urls:    urls,
		Sniffed: make(map[string][]*Data),
	}

}
