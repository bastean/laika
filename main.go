package laika

import (
	"log"
	"slices"

	"github.com/bastean/laika/pkg/context/data/application/createEmptyData"
	"github.com/bastean/laika/pkg/context/data/application/readData"
	"github.com/bastean/laika/pkg/context/data/application/saveData"
	"github.com/bastean/laika/pkg/context/shared/domain/aggregate"
	"github.com/bastean/laika/pkg/context/shared/domain/model"
	"github.com/bastean/laika/pkg/context/shared/domain/repository"
	"github.com/bastean/laika/pkg/context/shared/infrastructure/persistence"
	"github.com/bastean/laika/pkg/context/shared/infrastructure/scraper"
	"github.com/bastean/laika/pkg/context/website/application/sniffFromUrls"
)

type Laika struct {
	Data          aggregate.Data
	Store         repository.Repository
	SaveData      *saveData.SaveData
	Scraper       model.Scraper
	SniffFromUrls *sniffFromUrls.SniffFromUrls
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

func (laika *Laika) FromUrls(urls []string) {
	laika.SniffFromUrls.Run(laika.Data, urls, &sniffFromUrls.SniffFromUrlsOptions{FollowLinks: true})
}

func (laika *Laika) SniffedEmails() []string {
	emails := []string{}

	for _, sniffed := range laika.Data {
		for _, data := range sniffed {
			emails = append(emails, data.Emails...)
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

func NewEmptyData() aggregate.Data {
	return createEmptyData.NewCreateEmptyData().Run()
}

func NewInMemoryStore() repository.Repository {
	return nil
}

func NewLocalJsonStore(path, filename string) repository.Repository {
	return persistence.NewLocalJson(path, filename)
}

func ReadDataFromStore(persistence repository.Repository) (aggregate.Data, error) {
	return readData.NewReadData(persistence).Run()
}

func New(data aggregate.Data) *Laika {
	scraper, err := scraper.NewPlaywright(&scraper.PlaywrightOptions{Headless: true})

	if err != nil {
		log.Fatalln(err)
	}

	return &Laika{
		Data:          data,
		Scraper:       scraper,
		SniffFromUrls: sniffFromUrls.NewSniffFromUrls(scraper),
	}
}
