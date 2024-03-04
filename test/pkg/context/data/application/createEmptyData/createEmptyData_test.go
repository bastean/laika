package createEmptyData_test

import (
	"testing"

	"github.com/bastean/laika/pkg/context/data/application/createEmptyData"
	aggregateMother "github.com/bastean/laika/test/pkg/context/shared/domain/aggregate"
	"github.com/stretchr/testify/suite"
)

type CreateEmptyDataTestSuite struct {
	suite.Suite
	sut *createEmptyData.CreateEmptyData
}

func (suite *CreateEmptyDataTestSuite) SetupTest() {
	suite.sut = createEmptyData.NewCreateEmptyData()
}

func (suite *CreateEmptyDataTestSuite) TestCreateEmptyData() {
	expected := aggregateMother.Create()
	actual := suite.sut.Run()

	suite.EqualValues(expected, actual)
}

func TestCreateEmptyDataSuite(t *testing.T) {
	suite.Run(t, new(CreateEmptyDataTestSuite))
}
