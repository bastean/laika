package laika

import (
	"github.com/bastean/laika/internal/app/service"
)

type Laika struct {
	Sources []string
}

func (laika *Laika) ContentFromUrls() {
	for _, source := range laika.Sources {
		service.ContentFromUrl.Run(source)
	}
}

func (laika *Laika) Emails() {
	service.SniffEmails.Run()
}

func Sniff(sources []string) *Laika {
	return &Laika{
		Sources: sources,
	}
}
