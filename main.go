package laika

import (
	"slices"

	"github.com/bastean/laika/pkg/context/data/application/createEmptyData"
	"github.com/bastean/laika/pkg/context/data/application/readData"
	"github.com/bastean/laika/pkg/context/data/application/saveData"
	"github.com/bastean/laika/pkg/context/email/application/sniffEmails"
	"github.com/bastean/laika/pkg/context/shared/domain/aggregate"
	"github.com/bastean/laika/pkg/context/shared/domain/repository"
	"github.com/bastean/laika/pkg/context/shared/infrastructure/http"
	"github.com/bastean/laika/pkg/context/shared/infrastructure/persistence"
	"github.com/bastean/laika/pkg/context/website/application/sniffContentFromUrls"
)

type Laika struct {
	Data                 *aggregate.Data
	Store                repository.Repository
	SaveData             *saveData.SaveData
	SniffContentFromUrls *sniffContentFromUrls.SniffContentFromUrls
	SniffEmails          *sniffEmails.SniffEmails
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
	laika.SniffContentFromUrls.Run(laika.Data, urls)
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

	uniques := []string{}

	for _, email := range emails {
		if !slices.Contains(uniques, email) {
			uniques = append(uniques, email)
		}
	}

	return uniques
}

func NewEmptyData() *aggregate.Data {
	return createEmptyData.NewCreateEmptyData().Run()
}

func NewInMemoryStore() repository.Repository {
	return nil
}

func NewLocalJsonStore(path, filename string) repository.Repository {
	return persistence.NewLocalJson(path, filename)
}

func ReadDataFromStore(persistence repository.Repository) (*aggregate.Data, error) {
	return readData.NewReadData(persistence).Run()
}

func New(data *aggregate.Data) *Laika {
	return &Laika{
		Data:                 data,
		SniffContentFromUrls: sniffContentFromUrls.NewSniffContentFromUrls(http.NewClient()),
		SniffEmails:          sniffEmails.NewSniffEmails(),
	}
}
