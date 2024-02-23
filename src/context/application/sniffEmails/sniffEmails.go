package sniffEmails

import (
	"net/mail"
	"strings"

	"github.com/bastean/laika/src/context/domain/aggregate"
	"github.com/bastean/laika/src/context/domain/repository"
	"github.com/bastean/laika/src/context/domain/service"
)

type SniffEmails struct {
	Repository repository.Repository
}

func (sniffEmails *SniffEmails) Run() {
	// TODO(refactor): sniff emails

	laika := sniffEmails.Repository.Read()

	laikaFound := new(aggregate.Laika)
	laikaFound.Sniffed = make(map[string][]*aggregate.Data)

	for sniffed, values := range laika.Sniffed {
		for _, data := range values {
			rawEmails := []string{}

			rawEmails = append(rawEmails, service.PlainEmailRegexp.FindAllString(data.Content, -1)...)

			for _, concat := range service.JsConcatenationRegexp.FindAllString(data.Content, -1) {
				letters := service.JsConcatenationRemoveRegexp.FindAllString(concat, -1)
				rawEmail := strings.Join(letters, "")
				rawEmails = append(rawEmails, rawEmail)
			}

			contentWithoutComments := service.HtmlCommentRegexp.ReplaceAllString(data.Content, "")

			rawEmails = append(rawEmails, service.PlainEmailRegexp.FindAllString(contentWithoutComments, -1)...)

			parsedEmails := []string{}

			for _, rawEmail := range rawEmails {
				email, err := mail.ParseAddress(rawEmail)

				if err != nil {
					continue
				}

				parsedEmails = append(parsedEmails, email.Address)
			}

			emails := []string{}
			isNotPresent := true

			for _, parsedEmail := range parsedEmails {
				isNotPresent = true

				for _, email := range emails {
					if parsedEmail == email {
						isNotPresent = false
						break
					}

				}

				if isNotPresent {
					emails = append(emails, parsedEmail)
				}
			}

			laikaFound.Sniffed[sniffed] = append(laikaFound.Sniffed[sniffed], &aggregate.Data{Source: data.Source, Content: data.Content, Found: map[string][]string{"Emails": emails}})
		}
	}

	sniffEmails.Repository.Save(laikaFound)
}

func NewSniffEmails(repository repository.Repository) *SniffEmails {
	return &SniffEmails{
		repository,
	}
}
