package laika

import (
	"github.com/bastean/laika/pkg/context/application/createEmptyData"
	"github.com/bastean/laika/pkg/context/application/readData"
	"github.com/bastean/laika/pkg/context/application/saveData"
	"github.com/bastean/laika/pkg/context/application/sniffContentFromUrl"
	"github.com/bastean/laika/pkg/context/application/sniffEmails"
	"github.com/bastean/laika/pkg/context/domain/aggregate"
	"github.com/bastean/laika/pkg/context/domain/repository"
	"github.com/bastean/laika/pkg/context/infrastructure/persistence"
)

type Laika struct {
	Data                *aggregate.Laika
	Store               repository.Repository
	SaveData            *saveData.SaveData
	SniffContentFromUrl *sniffContentFromUrl.SniffContentFromUrl
	SniffEmails         *sniffEmails.SniffEmails
}

func (laika *Laika) SetStore(persistence repository.Repository) {
	laika.Store = persistence
	laika.SaveData = saveData.NewSaveData(persistence)
}

func (laika *Laika) SaveSniffed() {
	if laika.Store != nil {
		laika.SaveData.Run(laika.Data)
	}
}

func (laika *Laika) ContentFromUrls(urls []string) {
	for _, url := range urls {
		laika.SniffContentFromUrl.Run(laika.Data, url)
	}
}

func (laika *Laika) EmailsFromContent() {
	laika.SniffEmails.Run(laika.Data)
}

func (laika *Laika) SniffedEmails() []string {
	emails := []string{}

	for _, sniffed := range laika.Data.Sniffed {
		for _, data := range sniffed {
			emails = append(emails, data.Found["Emails"]...)
		}
	}

	return emails
}

func NewEmptyData() *aggregate.Laika {
	return createEmptyData.NewCreateEmptyData().Run()
}

func NewInMemoryStore() repository.Repository {
	return nil
}

func NewLocalJsonStore(path, filename string) *persistence.LocalJson {
	return persistence.NewLocalJson(path, filename)
}

func ReadDataFromStore(persistence repository.Repository) *aggregate.Laika {
	return readData.NewReadData(persistence).Run()
}

func New(data *aggregate.Laika) *Laika {
	return &Laika{
		Data:                data,
		SniffContentFromUrl: sniffContentFromUrl.NewSniffContentFromUrl(),
		SniffEmails:         sniffEmails.NewSniffEmails(),
	}
}
