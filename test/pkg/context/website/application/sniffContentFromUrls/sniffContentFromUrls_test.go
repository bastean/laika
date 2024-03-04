package sniffContentFromUrls_test

import (
	"testing"

	"github.com/bastean/laika/pkg/context/website/application/sniffContentFromUrls"
	aggregateMother "github.com/bastean/laika/test/pkg/context/shared/domain/aggregate"
	httpMock "github.com/bastean/laika/test/pkg/context/shared/infrastructure/mock/http"
	"github.com/stretchr/testify/suite"
)

type SniffContentFromUrlsTestSuite struct {
	suite.Suite
	sut    *sniffContentFromUrls.SniffContentFromUrls
	client *httpMock.ClientMock
}

func (suite *SniffContentFromUrlsTestSuite) SetupTest() {
	suite.client = new(httpMock.ClientMock)
	suite.client.Response = `<span>email@example.com</span>`

	suite.sut = sniffContentFromUrls.NewSniffContentFromUrls(suite.client)
}

func (suite *SniffContentFromUrlsTestSuite) TestSniffContentFromUrls() {
	data := aggregateMother.Create()

	url := "http://example.com/"
	urls := []string{url}

	suite.client.On("Get", url)

	suite.sut.Run(data, urls)

	suite.client.AssertCalled(suite.T(), "Get", url)

	expected := suite.client.Response
	actual := data.Sniffed["example.com"][0].Content

	suite.EqualValues(expected, actual)
}

func TestSniffContentFromUrlsSuite(t *testing.T) {
	suite.Run(t, new(SniffContentFromUrlsTestSuite))
}
