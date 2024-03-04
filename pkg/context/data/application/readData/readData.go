package readData

import (
	"github.com/bastean/laika/pkg/context/shared/domain/aggregate"
	"github.com/bastean/laika/pkg/context/shared/domain/repository"
)

type ReadData struct {
	Repository repository.Repository
}

func (read *ReadData) Run() (*aggregate.Data, error) {
	return read.Repository.Read()
}

func NewReadData(repository repository.Repository) *ReadData {
	return &ReadData{
		repository,
	}
}
