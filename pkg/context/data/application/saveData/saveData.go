package saveData

import (
	"github.com/bastean/laika/pkg/context/shared/domain/aggregate"
	"github.com/bastean/laika/pkg/context/shared/domain/repository"
)

type SaveData struct {
	Repository repository.Repository
}

func (save *SaveData) Run(data *aggregate.Data) {
	save.Repository.Save(data)
}

func NewSaveData(repository repository.Repository) *SaveData {
	return &SaveData{
		repository,
	}
}
