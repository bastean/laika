package repository

import (
	"github.com/bastean/laika/pkg/context/domain/aggregate"
	aggregateMother "github.com/bastean/laika/test/pkg/context/domain/aggregate"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (mock *RepositoryMock) Save(laika *aggregate.Laika) {
	mock.Called(laika)
}

func (mock *RepositoryMock) Read() *aggregate.Laika {
	return aggregateMother.Create()
}
