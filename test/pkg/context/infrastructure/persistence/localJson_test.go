package persistence_test

import (
	"os"
	"testing"

	"github.com/bastean/laika/pkg/context/infrastructure/persistence"
	"github.com/bastean/laika/test/pkg/context/domain/aggregate"
	"github.com/stretchr/testify/suite"
)

type LocalJsonTestSuite struct {
	suite.Suite
	sut      *persistence.LocalJson
	path     string
	filename string
}

func (suite *LocalJsonTestSuite) SetupTest() {
	suite.path = "."
	suite.filename = "laika"

	suite.sut = persistence.NewLocalJson(suite.path, suite.filename)
}

func (suite *LocalJsonTestSuite) TearDownTest() {
	os.Remove(suite.filename + ".json")
}

func (suite *LocalJsonTestSuite) TestSave() {
	data := aggregate.Create()

	suite.NotPanics(func() { suite.sut.Save(data) })
}

func (suite *LocalJsonTestSuite) TestRead() {
	expected := aggregate.Create()

	suite.NotPanics(func() { suite.sut.Save(expected) })

	actual, err := suite.sut.Read()

	suite.NoError(err)
	suite.EqualValues(expected, actual)
}

func TestLocalJsonSuite(t *testing.T) {
	suite.Run(t, new(LocalJsonTestSuite))
}
