package sniffEmails

import (
	"github.com/bastean/laika/pkg/context/shared/domain/aggregate"
	"github.com/bastean/laika/pkg/context/shared/domain/service"
)

type SniffEmails struct{}

func (sniff *SniffEmails) Run(data *aggregate.Data) {
	dataFound := new(aggregate.Data)
	dataFound.Sniffed = make(map[string][]*aggregate.Sniffed)

	for source, values := range data.Sniffed {
		for _, sniffed := range values {
			rawEmails := []string{}

			rawEmails = append(rawEmails, service.SniffPlainEmails(sniffed.Content)...)

			rawEmails = append(rawEmails, service.SniffJsConcatenationEmails(sniffed.Content)...)

			contentWithoutComments := service.RemoveHtmlComments(sniffed.Content)

			rawEmails = append(rawEmails, service.SniffPlainEmails(contentWithoutComments)...)

			parsedEmails := service.ParseAddresses(rawEmails)

			emails := service.RemoveDuplicates(parsedEmails)

			dataFound.Sniffed[source] = append(dataFound.Sniffed[source], &aggregate.Sniffed{Source: sniffed.Source, Content: sniffed.Content, Found: map[string][]string{"Emails": emails}})
		}
	}

	data.Sniffed = dataFound.Sniffed
}

func NewSniffEmails() *SniffEmails {
	return new(SniffEmails)
}
