package readData_test

import (
	"testing"

	"github.com/bastean/laika/pkg/context/data/application/readData"
	aggregateMother "github.com/bastean/laika/test/pkg/context/shared/domain/aggregate"
	repositoryMock "github.com/bastean/laika/test/pkg/context/shared/infrastructure/mock/repository"
	"github.com/stretchr/testify/suite"
)

type ReadDataTestSuite struct {
	suite.Suite
	sut        *readData.ReadData
	repository *repositoryMock.RepositoryMock
}

func (suite *ReadDataTestSuite) SetupTest() {
	suite.repository = new(repositoryMock.RepositoryMock)
	suite.sut = readData.NewReadData(suite.repository)
}

func (suite *ReadDataTestSuite) TestReadData() {
	expected := aggregateMother.Create()
	actual, err := suite.sut.Run()

	suite.NoError(err)
	suite.EqualValues(expected, actual)
}

func TestReadDataSuite(t *testing.T) {
	suite.Run(t, new(ReadDataTestSuite))
}
