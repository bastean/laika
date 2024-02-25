package createEmptyData_test

import (
	"testing"

	"github.com/bastean/laika/pkg/context/application/createEmptyData"
	"github.com/bastean/laika/test/pkg/context/domain/aggregate"
	"github.com/stretchr/testify/suite"
)

type CreateEmptyDataTestSuite struct {
	suite.Suite
}

func (suite *CreateEmptyDataTestSuite) SetupTest() {}

func (suite *CreateEmptyDataTestSuite) TestCreateEmptyData() {
	expected := aggregate.Create()
	actual := createEmptyData.NewCreateEmptyData().Run()

	suite.EqualValues(expected, actual)
}

func TestCreateEmptyDataSuite(t *testing.T) {
	suite.Run(t, new(CreateEmptyDataTestSuite))
}
