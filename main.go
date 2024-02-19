package laika

import (
	"github.com/bastean/laika/pkg/sniff"
)

type laika struct {
	Urls []string
}

func (laika *laika) Emails() []string {
	emails := []string{}

	for _, url := range laika.Urls {
		sniff.Emails(url)
	}

	return emails
}

func Sniff(urls []string) *laika {
	return &laika{
		Urls: urls,
	}
}
