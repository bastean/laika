package saveData_test

import (
	"testing"

	"github.com/bastean/laika/pkg/context/data/application/saveData"
	aggregateMother "github.com/bastean/laika/test/pkg/context/shared/domain/aggregate"
	repositoryMock "github.com/bastean/laika/test/pkg/context/shared/infrastructure/mock/repository"
	"github.com/stretchr/testify/suite"
)

type SaveDataTestSuite struct {
	suite.Suite
	sut        *saveData.SaveData
	repository *repositoryMock.RepositoryMock
}

func (suite *SaveDataTestSuite) SetupTest() {
	suite.repository = new(repositoryMock.RepositoryMock)
	suite.sut = saveData.NewSaveData(suite.repository)
}

func (suite *SaveDataTestSuite) TestSaveData() {
	data := aggregateMother.Create()

	suite.repository.On("Save", data)

	suite.sut.Run(data)

	suite.repository.AssertCalled(suite.T(), "Save", data)
}

func TestSaveDataSuite(t *testing.T) {
	suite.Run(t, new(SaveDataTestSuite))
}
