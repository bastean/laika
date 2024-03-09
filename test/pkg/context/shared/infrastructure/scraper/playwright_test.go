package scraper_test

import (
	"os"
	"testing"

	"github.com/bastean/laika/pkg/context/shared/domain/model"
	"github.com/bastean/laika/pkg/context/shared/infrastructure/scraper"
	"github.com/stretchr/testify/suite"
)

type PlaywrightTestSuite struct {
	suite.Suite
	sut     model.Scraper
	testUrl string
}

func (suite *PlaywrightTestSuite) SetupTest() {
	suite.sut, _ = scraper.NewPlaywright(&scraper.PlaywrightOptions{Headless: true})
	suite.testUrl = os.Getenv("TEST_URL")
}

func (suite *PlaywrightTestSuite) TestGetContent() {
	notExpected := ""

	actual := suite.sut.GetContent(suite.testUrl)

	suite.NotEqualValues(notExpected, actual)
}

func (suite *PlaywrightTestSuite) TestGetLinks() {
	notExpected := []string{}

	actual := suite.sut.GetLinks(suite.testUrl)

	suite.NotEqualValues(len(notExpected), len(actual))
}

func TestPlaywrightSuite(t *testing.T) {
	suite.Run(t, new(PlaywrightTestSuite))
}
