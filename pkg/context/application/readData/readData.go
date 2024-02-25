package readData

import (
	"github.com/bastean/laika/pkg/context/domain/aggregate"
	"github.com/bastean/laika/pkg/context/domain/repository"
)

type ReadData struct {
	Repository repository.Repository
}

func (read *ReadData) Run() *aggregate.Laika {
	return read.Repository.Read()
}

func NewReadData(repository repository.Repository) *ReadData {
	return &ReadData{
		repository,
	}
}
