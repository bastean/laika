package saveData

import (
	"github.com/bastean/laika/pkg/context/domain/aggregate"
	"github.com/bastean/laika/pkg/context/domain/repository"
)

type SaveData struct {
	Repository repository.Repository
}

func (save *SaveData) Run(laika *aggregate.Laika) {
	save.Repository.Save(laika)
}

func NewSaveData(repository repository.Repository) *SaveData {
	return &SaveData{
		repository,
	}
}
