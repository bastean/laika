package httpMock

import (
	"github.com/stretchr/testify/mock"
)

type ClientMock struct {
	mock.Mock
	Response string
}

func (mock *ClientMock) Get(url string) (string, error) {
	mock.Called(url)
	return mock.Response, nil
}
