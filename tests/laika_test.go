package main_test

import (
	"os"
	"testing"

	"github.com/bastean/laika"
	"github.com/stretchr/testify/suite"
)

type LaikaTestSuite struct {
	suite.Suite
	laika *laika.Laika
}

func (suite *LaikaTestSuite) SetupTest() {
	urls := []string{os.Getenv("TEST_URL")}
	suite.laika = laika.Sniff(urls)
}

func (suite *LaikaTestSuite) TestSniffEmails() {
	expected := []string{"email@example.com"}
	found := suite.laika.Emails()
	suite.EqualValues(expected, found)
}

func TestLaikaSuite(t *testing.T) {
	suite.Run(t, new(LaikaTestSuite))
}
