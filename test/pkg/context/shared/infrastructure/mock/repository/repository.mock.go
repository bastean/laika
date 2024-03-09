package repositoryMock

import (
	"github.com/bastean/laika/pkg/context/shared/domain/aggregate"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (mock *RepositoryMock) Save(data aggregate.Data) error {
	mock.Called(data)
	return nil
}

func (mock *RepositoryMock) Read() (aggregate.Data, error) {
	args := mock.Called()
	return args.Get(0).(aggregate.Data), nil
}
