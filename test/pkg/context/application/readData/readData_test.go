package readData_test

import (
	"testing"

	"github.com/bastean/laika/pkg/context/application/readData"
	"github.com/bastean/laika/test/pkg/context/domain/aggregate"
	"github.com/bastean/laika/test/pkg/context/infrastructure/mock/repository"
	"github.com/stretchr/testify/suite"
)

type ReadDataTestSuite struct {
	suite.Suite
	repository *repository.RepositoryMock
	sut        *readData.ReadData
}

func (suite *ReadDataTestSuite) SetupTest() {
	suite.repository = new(repository.RepositoryMock)
	suite.sut = readData.NewReadData(suite.repository)
}

func (suite *ReadDataTestSuite) TestReadData() {
	expected := aggregate.Create()
	actual := suite.sut.Run()

	suite.EqualValues(expected, actual)
}

func TestReadDataSuite(t *testing.T) {
	suite.Run(t, new(ReadDataTestSuite))
}
