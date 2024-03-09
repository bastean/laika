package sniffContentFromUrls_test

import (
	"testing"

	"github.com/bastean/laika/pkg/context/website/application/sniffFromUrls"
	aggregateMother "github.com/bastean/laika/test/pkg/context/shared/domain/aggregate"
	scraperMock "github.com/bastean/laika/test/pkg/context/shared/infrastructure/mock/scraper"
	"github.com/stretchr/testify/suite"
)

type SniffFromUrlsTestSuite struct {
	suite.Suite
	sut     *sniffFromUrls.SniffFromUrls
	scraper *scraperMock.ScraperMock
}

func (suite *SniffFromUrlsTestSuite) SetupTest() {
	suite.scraper = new(scraperMock.ScraperMock)

	suite.sut = sniffFromUrls.NewSniffFromUrls(suite.scraper)
}

func (suite *SniffFromUrlsTestSuite) TestSniffFromUrls() {
	data := aggregateMother.Create()

	url := "http://example.com/"

	response := `<span>email@example.com</span>`

	links := []string{}

	email := "email@example.com"

	urls := []string{url}

	suite.scraper.On("GetContent", url).Return(response)

	suite.scraper.On("GetLinks", url).Return(links)

	suite.sut.Run(data, urls, &sniffFromUrls.SniffFromUrlsOptions{FollowLinks: true})

	suite.scraper.AssertCalled(suite.T(), "GetContent", url)

	suite.scraper.AssertCalled(suite.T(), "GetLinks", url)

	expected := email

	actual := data["example.com"][0].Emails[0]

	suite.EqualValues(expected, actual)
}

func TestSniffContentFromUrlsSuite(t *testing.T) {
	suite.Run(t, new(SniffFromUrlsTestSuite))
}
