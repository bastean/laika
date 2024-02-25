package main_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type LaikaTestSuite struct {
	suite.Suite
}

func (suite *LaikaTestSuite) SetupTest() {}

func (suite *LaikaTestSuite) TestSniffContentFromUrls() {
	// TODO(test): sniff content from urls

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
