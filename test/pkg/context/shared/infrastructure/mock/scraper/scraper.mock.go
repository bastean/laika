package scraperMock

import (
	"github.com/stretchr/testify/mock"
)

type ScraperMock struct {
	mock.Mock
}

func (mock *ScraperMock) GetContent(source string) string {
	args := mock.Called(source)
	return args.Get(0).(string)
}

func (mock *ScraperMock) GetLinks(source string) []string {
	args := mock.Called(source)
	return args.Get(0).([]string)
}
