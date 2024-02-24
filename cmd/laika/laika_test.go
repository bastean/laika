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

func (suite *LaikaTestSuite) TestContentFromUrl() {
	// TODO(test): content from url

	expected := "x"
	found := "y"

	suite.EqualValues(expected, found)
}

func (suite *LaikaTestSuite) TestSniffEmails() {
	// TODO(test): sniff emails

	expected := "x"
	found := "y"

	suite.EqualValues(expected, found)
}

func TestLaikaSuite(t *testing.T) {
	suite.Run(t, new(LaikaTestSuite))
}
