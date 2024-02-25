package saveData_test

import (
	"testing"

	"github.com/bastean/laika/pkg/context/application/saveData"
	"github.com/bastean/laika/test/pkg/context/domain/aggregate"
	"github.com/bastean/laika/test/pkg/context/infrastructure/mock/repository"
	"github.com/stretchr/testify/suite"
)

type SaveDataTestSuite struct {
	suite.Suite
	repository *repository.RepositoryMock
	sut        *saveData.SaveData
}

func (suite *SaveDataTestSuite) SetupTest() {
	suite.repository = new(repository.RepositoryMock)
	suite.sut = saveData.NewSaveData(suite.repository)
}

func (suite *SaveDataTestSuite) TestSaveData() {
	data := aggregate.Create()

	suite.repository.On("Save", data)

	suite.sut.Run(data)

	suite.repository.AssertCalled(suite.T(), "Save", data)
}

func TestSaveDataSuite(t *testing.T) {
	suite.Run(t, new(SaveDataTestSuite))
}
