package persistence_test

import (
	"os"
	"testing"

	"github.com/bastean/laika/pkg/context/shared/domain/repository"
	"github.com/bastean/laika/pkg/context/shared/infrastructure/persistence"
	aggregateMother "github.com/bastean/laika/test/pkg/context/shared/domain/aggregate"
	"github.com/stretchr/testify/suite"
)

type LocalJsonTestSuite struct {
	suite.Suite
	sut      repository.Repository
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
	data := aggregateMother.Create()

	suite.NotPanics(func() { suite.sut.Save(data) })
}

func (suite *LocalJsonTestSuite) TestRead() {
	expected := aggregateMother.Create()

	suite.NotPanics(func() { suite.sut.Save(expected) })

	actual, err := suite.sut.Read()

	suite.NoError(err)

	suite.EqualValues(expected, actual)
}

func TestLocalJsonSuite(t *testing.T) {
	suite.Run(t, new(LocalJsonTestSuite))
}
